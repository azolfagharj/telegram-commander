package bot

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"golang.org/x/net/proxy"

	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/executor"
	"github.com/azolfagharj/telegram-commander/internal/function"
)

const telegramMaxMessageLen = 4096

// App wires config, functions, executor and Telegram handlers.
type App struct {
	Cfg      *config.Config
	Registry *function.Registry
	Exec     executor.Executor
	Index    *Index
	Log      *slog.Logger

	navMu     sync.Mutex
	nav       map[int64]userMenu
	confirmMu sync.Mutex
	confirms  map[int64]confirmWait
}

type userMenu struct {
	nodeID string
	page   int
}

type confirmWait struct {
	buttonID string
	expires  time.Time
}

// NewApp builds an App from loaded pieces.
func NewApp(cfg *config.Config, reg *function.Registry, exec executor.Executor, log *slog.Logger) *App {
	if log == nil {
		log = slog.Default()
	}
	return &App{
		Cfg:      cfg,
		Registry: reg,
		Exec:     exec,
		Index:    BuildIndex(cfg.Buttons, cfg.ButtonsColumns, cfg.PageSize),
		Log:      log,
		nav:      make(map[int64]userMenu),
		confirms: make(map[int64]confirmWait),
	}
}

// HTTPClient builds an HTTP client with optional proxy and TLS skip-verify.
func HTTPClient(cfg config.TelegramConfig) (*http.Client, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.Insecure}, //nolint:gosec // configurable by admin
	}
	if cfg.Proxy.Enabled {
		u, err := url.Parse(cfg.Proxy.URL)
		if err != nil {
			return nil, fmt.Errorf("parse proxy url: %w", err)
		}
		switch u.Scheme {
		case "socks5", "socks5h":
			var auth *proxy.Auth
			if u.User != nil {
				pass, _ := u.User.Password()
				auth = &proxy.Auth{User: u.User.Username(), Password: pass}
			}
			dialer, err := proxy.SOCKS5("tcp", u.Host, auth, proxy.Direct)
			if err != nil {
				return nil, fmt.Errorf("socks5 dialer: %w", err)
			}
			cd, ok := dialer.(proxy.ContextDialer)
			if !ok {
				return nil, fmt.Errorf("socks5 dialer does not support DialContext")
			}
			transport.DialContext = cd.DialContext
		case "http", "https":
			transport.Proxy = http.ProxyURL(u)
		default:
			return nil, fmt.Errorf("unsupported proxy scheme %q", u.Scheme)
		}
	}
	return &http.Client{Transport: transport, Timeout: 0}, nil
}

// NewBot creates a go-telegram/bot instance configured for this app.
func (a *App) NewBot() (*bot.Bot, error) {
	client, err := HTTPClient(a.Cfg.Telegram)
	if err != nil {
		return nil, err
	}
	opts := []bot.Option{
		bot.WithHTTPClient(60*time.Second, client),
		bot.WithServerURL(strings.TrimRight(a.Cfg.Telegram.API, "/")),
		bot.WithErrorsHandler(func(err error) {
			a.Log.Warn("telegram client error", "err", err)
		}),
		bot.WithDefaultHandler(a.defaultHandler),
		bot.WithMessageTextHandler("start", bot.MatchTypeCommandStartOnly, a.handleStart),
		bot.WithMessageTextHandler("help", bot.MatchTypeCommandStartOnly, a.handleHelp),
	}
	if a.Cfg.Telegram.EnableRunCommand {
		opts = append(opts, bot.WithMessageTextHandler("run", bot.MatchTypeCommandStartOnly, a.handleRun))
	}
	return bot.New(a.Cfg.Telegram.BotToken, opts...)
}

// Start runs long polling until ctx is cancelled.
func (a *App) Start(ctx context.Context) error {
	b, err := a.NewBot()
	if err != nil {
		return err
	}
	a.Log.Info("bot starting")
	b.Start(ctx)
	return nil
}

func (a *App) isAllowed(user *models.User) bool {
	if user == nil {
		return false
	}
	uid := strconv.FormatInt(user.ID, 10)
	uname := strings.ToLower(user.Username)
	for _, entry := range a.Cfg.Telegram.AllowedUsers {
		e := strings.TrimSpace(entry)
		if e == "" {
			continue
		}
		if e == uid {
			return true
		}
		if strings.EqualFold(strings.TrimPrefix(e, "@"), uname) && uname != "" {
			return true
		}
	}
	return false
}

func (a *App) denyMessage(user *models.User) string {
	uname := user.Username
	if uname == "" {
		uname = "(none)"
	}
	return fmt.Sprintf(
		"You are not allowed to use this bot.\nYour user_id: %d\nYour username: %s\nContact the admin to get access.",
		user.ID, uname,
	)
}

func (a *App) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}
	if !a.isAllowed(update.Message.From) {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   a.denyMessage(update.Message.From),
		})
		return
	}
	a.handleMenuText(ctx, b, update.Message)
}

func (a *App) handleStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}
	if !a.isAllowed(update.Message.From) {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   a.denyMessage(update.Message.From),
		})
		return
	}
	a.sendMenu(ctx, b, update.Message.Chat.ID, update.Message.From.ID, "", 0)
}

func (a *App) handleHelp(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}
	if !a.isAllowed(update.Message.From) {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   a.denyMessage(update.Message.From),
		})
		return
	}
	text := "telegram-commander\n\n/start - open menu\n/help - show this help"
	if a.Cfg.Telegram.EnableRunCommand {
		text += "\n/run <button name> - run a button by name"
	}
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text,
	})
}

func (a *App) handleRun(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}
	if !a.isAllowed(update.Message.From) {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   a.denyMessage(update.Message.From),
		})
		return
	}
	parts := strings.SplitN(update.Message.Text, " ", 2)
	if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Usage: /run <button name>",
		})
		return
	}
	name := strings.TrimSpace(parts[1])
	node := a.Index.FindByName(name)
	if node == nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Button %q not found", name),
		})
		return
	}
	a.executeButton(ctx, b, update.Message.Chat.ID, update.Message.From, node)
}

func (a *App) handleMenuText(ctx context.Context, b *bot.Bot, msg *models.Message) {
	text := strings.TrimSpace(msg.Text)
	chatID := msg.Chat.ID
	userID := msg.From.ID
	if text == "" || strings.HasPrefix(text, "/") {
		a.replyWithMenu(ctx, b, chatID, userID, "Use /start to open the menu.")
		return
	}

	st := a.getNav(userID)

	switch text {
	case btnHome:
		a.clearConfirm(userID)
		a.sendMenu(ctx, b, chatID, userID, "", 0)
		return
	case btnBack:
		a.clearConfirm(userID)
		parent := ""
		if st.nodeID != "" {
			if n := a.Index.ByID[st.nodeID]; n != nil {
				parent = n.ParentID
			}
		}
		a.sendMenu(ctx, b, chatID, userID, parent, 0)
		return
	case btnYes:
		buttonID, ok := a.consumeConfirm(userID)
		if !ok {
			a.replyWithMenu(ctx, b, chatID, userID, "Confirmation expired. Please try again.")
			return
		}
		node := a.Index.ByID[buttonID]
		if node == nil {
			return
		}
		a.executeButton(ctx, b, chatID, msg.From, node)
		return
	case btnCancel:
		a.clearConfirm(userID)
		a.sendMenu(ctx, b, chatID, userID, st.nodeID, st.page)
		return
	case btnPrev:
		page := st.page - 1
		if page < 0 {
			page = 0
		}
		a.sendMenu(ctx, b, chatID, userID, st.nodeID, page)
		return
	case btnNext:
		a.sendMenu(ctx, b, chatID, userID, st.nodeID, st.page+1)
		return
	}

	node := a.Index.ChildByLabel(st.nodeID, text)
	if node == nil {
		a.replyWithMenu(ctx, b, chatID, userID, "Unknown option. Use /start to open the menu.")
		return
	}
	if node.Type == "category" {
		a.clearConfirm(userID)
		a.sendMenu(ctx, b, chatID, userID, node.ID, 0)
		return
	}
	if node.Confirm {
		a.showConfirm(ctx, b, chatID, userID, node)
		return
	}
	a.executeButton(ctx, b, chatID, msg.From, node)
}

func (a *App) sendMenu(ctx context.Context, b *bot.Bot, chatID, userID int64, nodeID string, page int) {
	kb, title, err := a.Index.KeyboardFor(nodeID, page)
	if err != nil {
		a.Log.Error("keyboard", "err", err)
		nodeID = ""
		page = 0
		kb, title, err = a.Index.KeyboardFor("", 0)
		if err != nil {
			return
		}
	}
	page = a.Index.clampPage(nodeID, page)
	a.setNav(userID, nodeID, page)
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        title,
		ReplyMarkup: kb,
	})
}

func (a *App) replyWithMenu(ctx context.Context, b *bot.Bot, chatID, userID int64, text string) {
	params := &bot.SendMessageParams{ChatID: chatID, Text: text}
	if kb := a.menuKeyboard(userID); kb != nil {
		params.ReplyMarkup = kb
	}
	_, _ = b.SendMessage(ctx, params)
}

func (a *App) menuKeyboard(userID int64) *models.ReplyKeyboardMarkup {
	st := a.getNav(userID)
	kb, _, err := a.Index.KeyboardFor(st.nodeID, st.page)
	if err != nil {
		kb, _, err = a.Index.KeyboardFor("", 0)
		if err != nil {
			return nil
		}
	}
	return kb
}

func (a *App) getNav(userID int64) userMenu {
	a.navMu.Lock()
	defer a.navMu.Unlock()
	return a.nav[userID]
}

func (a *App) setNav(userID int64, nodeID string, page int) {
	a.navMu.Lock()
	a.nav[userID] = userMenu{nodeID: nodeID, page: page}
	a.navMu.Unlock()
}

func (a *App) showConfirm(ctx context.Context, b *bot.Bot, chatID, userID int64, node *Node) {
	a.confirmMu.Lock()
	a.confirms[userID] = confirmWait{
		buttonID: node.ID,
		expires:  time.Now().Add(a.Cfg.ConfirmTTL.Duration),
	}
	a.confirmMu.Unlock()

	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        fmt.Sprintf("Confirm: %s ?", node.Label()),
		ReplyMarkup: ConfirmKeyboard(),
	})
}

func (a *App) consumeConfirm(userID int64) (string, bool) {
	a.confirmMu.Lock()
	defer a.confirmMu.Unlock()
	w, ok := a.confirms[userID]
	delete(a.confirms, userID)
	if !ok || !time.Now().Before(w.expires) {
		return "", false
	}
	return w.buttonID, true
}

func (a *App) clearConfirm(userID int64) {
	a.confirmMu.Lock()
	delete(a.confirms, userID)
	a.confirmMu.Unlock()
}

func (a *App) executeButton(ctx context.Context, b *bot.Bot, chatID int64, user *models.User, node *Node) {
	def, ok := a.Registry.GetCaseInsensitive(node.Function)
	if !ok {
		a.replyWithMenu(ctx, b, chatID, user.ID, fmt.Sprintf("Unknown function %q", node.Function))
		return
	}
	params := function.ButtonParams(node.Raw)
	cmd, err := def.RenderRun(params)
	if err != nil {
		a.replyWithMenu(ctx, b, chatID, user.ID, "Failed to build command: "+err.Error())
		return
	}

	timeout := a.Cfg.Timeout.Duration
	if node.Timeout > 0 {
		timeout = node.Timeout
	}
	workdir := a.Cfg.WorkDir
	if node.WorkDir != "" {
		workdir = node.WorkDir
	}
	env := map[string]string{}
	for k, v := range a.Cfg.Env {
		env[k] = v
	}
	for k, v := range node.Env {
		env[k] = v
	}

	a.replyWithMenu(ctx, b, chatID, user.ID, fmt.Sprintf("Running: %s …", node.Label()))

	res, err := a.Exec.Run(ctx, executor.Spec{
		UserID:   user.ID,
		ButtonID: node.ID,
		Button:   node.Name,
		Command:  cmd,
		Shell:    a.Cfg.Shell,
		WorkDir:  workdir,
		Env:      env,
		Timeout:  timeout,
		MaxBytes: a.Cfg.MaxOutputBytes,
	})

	a.deliverResult(ctx, b, chatID, user.ID, node, res, err)
}

func (a *App) deliverResult(ctx context.Context, b *bot.Bot, chatID, userID int64, node *Node, res executor.Result, err error) {
	var body strings.Builder
	body.WriteString(fmt.Sprintf("Button: %s\n", node.Name))
	body.WriteString(fmt.Sprintf("Exit: %d\n", res.ExitCode))
	body.WriteString(fmt.Sprintf("Duration: %s\n", res.Duration.Round(time.Millisecond)))
	if res.TimedOut {
		body.WriteString("Status: TIMED OUT\n")
	}
	if err != nil && !res.TimedOut {
		body.WriteString("Error: " + err.Error() + "\n")
	}
	if res.Truncated {
		body.WriteString("(output truncated)\n")
	}
	body.WriteString("\n--- stdout ---\n")
	body.WriteString(res.Stdout)
	if res.Stderr != "" {
		body.WriteString("\n--- stderr ---\n")
		body.WriteString(res.Stderr)
	}
	text := body.String()
	kb := a.menuKeyboard(userID)
	if len(text) <= telegramMaxMessageLen {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      chatID,
			Text:        text,
			ReplyMarkup: kb,
		})
		return
	}
	// Send as document when too long.
	_, _ = b.SendDocument(ctx, &bot.SendDocumentParams{
		ChatID: chatID,
		Document: &models.InputFileUpload{
			Filename: "output.txt",
			Data:     strings.NewReader(text),
		},
		Caption:     fmt.Sprintf("%s output (too long for a message)", node.Name),
		ReplyMarkup: kb,
	})
}
