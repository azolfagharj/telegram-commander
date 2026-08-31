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
	nodeID    string
	page      int
	messageID int
	chatID    int64
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
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, a.handleCallback),
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
	if update.CallbackQuery != nil {
		return
	}
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
	chatID := update.Message.Chat.ID
	a.removeReplyKeyboard(ctx, b, chatID)
	a.sendMenu(ctx, b, chatID, update.Message.From.ID, "", 0)
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

func (a *App) handleCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	cq := update.CallbackQuery
	if cq == nil {
		return
	}
	_, _ = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: cq.ID,
		ShowAlert:       false,
	})
	user := cq.From
	if !a.isAllowed(&user) {
		return
	}
	chatID := int64(0)
	if cq.Message.Message != nil {
		chatID = cq.Message.Message.Chat.ID
	}
	if chatID == 0 {
		return
	}
	a.applyAction(ctx, b, chatID, &user, cq.Data)
}

func (a *App) handleMenuText(ctx context.Context, b *bot.Bot, msg *models.Message) {
	text := strings.TrimSpace(msg.Text)
	chatID := msg.Chat.ID
	userID := msg.From.ID
	if text == "" || strings.HasPrefix(text, "/") {
		a.showStatus(ctx, b, chatID, userID, "Use /start to open the menu.")
		return
	}

	switch text {
	case btnHome:
		a.applyAction(ctx, b, chatID, msg.From, cbHome)
		return
	case btnBack:
		a.applyAction(ctx, b, chatID, msg.From, cbBack)
		return
	case btnPrev:
		a.applyAction(ctx, b, chatID, msg.From, cbPrev)
		return
	case btnNext:
		a.applyAction(ctx, b, chatID, msg.From, cbNext)
		return
	case btnYes:
		a.applyAction(ctx, b, chatID, msg.From, cbYes)
		return
	case btnCancel:
		a.applyAction(ctx, b, chatID, msg.From, cbCancel)
		return
	}

	st := a.getNav(userID)
	node := a.Index.ChildByLabel(st.nodeID, text)
	if node == nil {
		a.showStatus(ctx, b, chatID, userID, "Unknown option. Use /start to open the menu.")
		return
	}
	if node.Type == "category" {
		a.applyAction(ctx, b, chatID, msg.From, cbOpen+node.ID)
		return
	}
	a.applyAction(ctx, b, chatID, msg.From, cbRun+node.ID)
}

func (a *App) applyAction(ctx context.Context, b *bot.Bot, chatID int64, user *models.User, data string) {
	userID := user.ID
	st := a.getNav(userID)

	switch {
	case data == cbHome:
		a.clearConfirm(userID)
		a.sendMenu(ctx, b, chatID, userID, "", 0)
	case data == cbBack:
		a.clearConfirm(userID)
		parent := ""
		if st.nodeID != "" {
			if n := a.Index.ByID[st.nodeID]; n != nil {
				parent = n.ParentID
			}
		}
		a.sendMenu(ctx, b, chatID, userID, parent, 0)
	case data == cbPrev:
		page := st.page - 1
		if page < 0 {
			return
		}
		a.sendMenu(ctx, b, chatID, userID, st.nodeID, page)
	case data == cbNext:
		view, err := a.Index.BuildMenu(st.nodeID, st.page)
		if err != nil || !view.HasNext {
			return
		}
		a.sendMenu(ctx, b, chatID, userID, st.nodeID, st.page+1)
	case data == cbYes:
		buttonID, ok := a.consumeConfirm(userID)
		if !ok {
			a.showStatus(ctx, b, chatID, userID, "Confirmation expired. Please try again.")
			return
		}
		node := a.Index.ByID[buttonID]
		if node == nil {
			return
		}
		a.executeButton(ctx, b, chatID, user, node)
	case data == cbCancel:
		a.clearConfirm(userID)
		a.sendMenu(ctx, b, chatID, userID, st.nodeID, st.page)
	case strings.HasPrefix(data, cbOpen):
		id := strings.TrimPrefix(data, cbOpen)
		node := a.Index.ByID[id]
		if node == nil || node.Type != "category" {
			return
		}
		a.clearConfirm(userID)
		a.sendMenu(ctx, b, chatID, userID, node.ID, 0)
	case strings.HasPrefix(data, cbRun):
		id := strings.TrimPrefix(data, cbRun)
		node := a.Index.ByID[id]
		if node == nil || node.Type != "button" {
			return
		}
		if node.Confirm {
			a.showConfirm(ctx, b, chatID, userID, node)
			return
		}
		a.executeButton(ctx, b, chatID, user, node)
	}
}

func (a *App) sendMenu(ctx context.Context, b *bot.Bot, chatID, userID int64, nodeID string, page int) {
	view, err := a.Index.BuildMenu(nodeID, page)
	if err != nil {
		a.Log.Error("keyboard", "err", err)
		nodeID = ""
		page = 0
		view, err = a.Index.BuildMenu("", 0)
		if err != nil {
			return
		}
	}
	page = view.Page
	a.setLocation(userID, nodeID, page)
	a.sendInline(ctx, b, chatID, userID, view.Title, view.Inline)
}

func (a *App) showStatus(ctx context.Context, b *bot.Bot, chatID, userID int64, text string) {
	st := a.getNav(userID)
	view, err := a.Index.BuildMenu(st.nodeID, st.page)
	if err != nil {
		view, err = a.Index.BuildMenu("", 0)
		if err != nil {
			a.sendInline(ctx, b, chatID, userID, text, nil)
			return
		}
	}
	a.sendInline(ctx, b, chatID, userID, text, view.Inline)
}

// sendInline shows one screen. It keeps a single message per chat: if that
// message still exists, it is edited in place; otherwise (first screen, or
// the old message can no longer be edited) a new message is sent.
func (a *App) sendInline(ctx context.Context, b *bot.Bot, chatID, userID int64, text string, inline *models.InlineKeyboardMarkup) {
	st := a.getNav(userID)
	if st.messageID != 0 && st.chatID == chatID {
		params := &bot.EditMessageTextParams{
			ChatID:    chatID,
			MessageID: st.messageID,
			Text:      text,
		}
		if inline != nil {
			params.ReplyMarkup = inline
		}
		if _, err := b.EditMessageText(ctx, params); err == nil || isMessageNotModified(err) {
			return
		}
	}
	params := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	}
	if inline != nil {
		params.ReplyMarkup = inline
	}
	msg, err := b.SendMessage(ctx, params)
	if err != nil || msg == nil {
		a.Log.Error("send menu", "err", err)
		return
	}
	a.setMessageID(userID, chatID, msg.ID)
}

func isMessageNotModified(err error) bool {
	return err != nil && strings.Contains(strings.ToLower(err.Error()), "message is not modified")
}

// removeReplyKeyboard drops a leftover custom keyboard from older versions.
func (a *App) removeReplyKeyboard(ctx context.Context, b *bot.Bot, chatID int64) {
	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "\u2800",
		ReplyMarkup: &models.ReplyKeyboardRemove{
			RemoveKeyboard: true,
		},
	})
	if err != nil || msg == nil {
		return
	}
	_, _ = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    chatID,
		MessageID: msg.ID,
	})
}

func (a *App) getNav(userID int64) userMenu {
	a.navMu.Lock()
	defer a.navMu.Unlock()
	return a.nav[userID]
}

func (a *App) setLocation(userID int64, nodeID string, page int) {
	a.navMu.Lock()
	st := a.nav[userID]
	st.nodeID = nodeID
	st.page = page
	a.nav[userID] = st
	a.navMu.Unlock()
}

func (a *App) setMessageID(userID, chatID int64, messageID int) {
	a.navMu.Lock()
	st := a.nav[userID]
	st.chatID = chatID
	st.messageID = messageID
	a.nav[userID] = st
	a.navMu.Unlock()
}

func (a *App) showConfirm(ctx context.Context, b *bot.Bot, chatID, userID int64, node *Node) {
	a.confirmMu.Lock()
	a.confirms[userID] = confirmWait{
		buttonID: node.ID,
		expires:  time.Now().Add(a.Cfg.ConfirmTTL.Duration),
	}
	a.confirmMu.Unlock()

	st := a.getNav(userID)
	hasBack := st.nodeID != ""
	a.sendInline(ctx, b, chatID, userID, fmt.Sprintf("Confirm: %s ?", node.Label()), ConfirmInlineKeyboard(hasBack))
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
		a.showStatus(ctx, b, chatID, user.ID, fmt.Sprintf("Unknown function %q", node.Function))
		return
	}
	params := function.ButtonParams(node.Raw)
	cmd, err := def.RenderRun(params)
	if err != nil {
		a.showStatus(ctx, b, chatID, user.ID, "Failed to build command: "+err.Error())
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

	a.showStatus(ctx, b, chatID, user.ID, fmt.Sprintf("Running: %s …", node.Label()))

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
	if len(text) > telegramMaxMessageLen {
		text = text[:telegramMaxMessageLen]
	}
	st := a.getNav(userID)
	hasBack := st.nodeID != ""
	a.sendInline(ctx, b, chatID, userID, text, ResultInlineKeyboard(hasBack))
}
