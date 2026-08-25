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

	confirmMu sync.Mutex
	confirms  map[string]time.Time // key: userID:buttonID
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
		confirms: make(map[string]time.Time),
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
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, a.handleCallback),
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
	// Unknown text: show menu tip.
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Use /start to open the menu.",
	})
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
	kb, title, err := a.Index.KeyboardFor("", 0)
	if err != nil {
		a.Log.Error("build keyboard", "err", err)
		return
	}
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        title,
		ReplyMarkup: kb,
	})
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
	a.executeButton(ctx, b, update.Message.Chat.ID, update.Message.From, node, false)
}

func (a *App) handleCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	cq := update.CallbackQuery
	if cq == nil || cq.From.ID == 0 {
		return
	}
	_, _ = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{CallbackQueryID: cq.ID})

	if !a.isAllowed(&cq.From) {
		chatID := callbackChatID(cq)
		if chatID != 0 {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   a.denyMessage(&cq.From),
			})
		}
		return
	}

	kind, payload := ParseCallback(cq.Data)
	chatID := callbackChatID(cq)
	msgID := callbackMessageID(cq)

	switch kind {
	case "home":
		a.editMenu(ctx, b, chatID, msgID, "", 0)
	case "nav":
		a.editMenu(ctx, b, chatID, msgID, payload, 0)
	case "page":
		parts := strings.SplitN(payload, ":", 2)
		nodeID := parts[0]
		page := 0
		if len(parts) == 2 {
			page, _ = strconv.Atoi(parts[1])
		}
		a.editMenu(ctx, b, chatID, msgID, nodeID, page)
	case "run":
		node := a.Index.ByID[payload]
		if node == nil || node.Type != "button" {
			return
		}
		if node.Confirm {
			a.showConfirm(ctx, b, chatID, msgID, &cq.From, node)
			return
		}
		a.executeButton(ctx, b, chatID, &cq.From, node, true)
	case "confirm":
		if !a.consumeConfirm(cq.From.ID, payload) {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "Confirmation expired. Please try again.",
			})
			return
		}
		node := a.Index.ByID[payload]
		if node == nil {
			return
		}
		a.executeButton(ctx, b, chatID, &cq.From, node, true)
	case "cancel":
		a.clearConfirm(cq.From.ID, payload)
		a.editMenu(ctx, b, chatID, msgID, "", 0)
	}
}

func callbackChatID(cq *models.CallbackQuery) int64 {
	if cq.Message.Message != nil {
		return cq.Message.Message.Chat.ID
	}
	return 0
}

func callbackMessageID(cq *models.CallbackQuery) int {
	if cq.Message.Message != nil {
		return cq.Message.Message.ID
	}
	return 0
}

func (a *App) editMenu(ctx context.Context, b *bot.Bot, chatID int64, msgID int, nodeID string, page int) {
	kb, title, err := a.Index.KeyboardFor(nodeID, page)
	if err != nil {
		a.Log.Error("keyboard", "err", err)
		return
	}
	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   msgID,
		Text:        title,
		ReplyMarkup: kb,
	})
	if err != nil {
		a.Log.Debug("edit message", "err", err)
	}
}

func (a *App) showConfirm(ctx context.Context, b *bot.Bot, chatID int64, msgID int, user *models.User, node *Node) {
	key := confirmKey(user.ID, node.ID)
	a.confirmMu.Lock()
	a.confirms[key] = time.Now().Add(a.Cfg.ConfirmTTL.Duration)
	a.confirmMu.Unlock()

	text := fmt.Sprintf("Confirm: %s ?", node.Label())
	kb := ConfirmKeyboard(node.ID)
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   msgID,
		Text:        text,
		ReplyMarkup: kb,
	})
	if err != nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      chatID,
			Text:        text,
			ReplyMarkup: kb,
		})
	}
}

func confirmKey(userID int64, buttonID string) string {
	return fmt.Sprintf("%d:%s", userID, buttonID)
}

func (a *App) consumeConfirm(userID int64, buttonID string) bool {
	key := confirmKey(userID, buttonID)
	a.confirmMu.Lock()
	defer a.confirmMu.Unlock()
	exp, ok := a.confirms[key]
	delete(a.confirms, key)
	if !ok {
		return false
	}
	return time.Now().Before(exp)
}

func (a *App) clearConfirm(userID int64, buttonID string) {
	a.confirmMu.Lock()
	delete(a.confirms, confirmKey(userID, buttonID))
	a.confirmMu.Unlock()
}

func (a *App) executeButton(ctx context.Context, b *bot.Bot, chatID int64, user *models.User, node *Node, viaCallback bool) {
	_ = viaCallback
	def, ok := a.Registry.GetCaseInsensitive(node.Function)
	if !ok {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   fmt.Sprintf("Unknown function %q", node.Function),
		})
		return
	}
	params := function.ButtonParams(node.Raw)
	cmd, err := def.RenderRun(params)
	if err != nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Failed to build command: " + err.Error(),
		})
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

	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   fmt.Sprintf("Running: %s …", node.Label()),
	})

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

	a.deliverResult(ctx, b, chatID, node, cmd, res, err)
}

func (a *App) deliverResult(ctx context.Context, b *bot.Bot, chatID int64, node *Node, cmd string, res executor.Result, err error) {
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
	if len(text) <= telegramMaxMessageLen {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   text,
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
		Caption: fmt.Sprintf("%s output (too long for a message)", node.Name),
	})
}
