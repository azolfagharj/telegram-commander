package bot_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/bot"
	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/executor"
	"github.com/azolfagharj/telegram-commander/internal/function"
)

type apiServer struct {
	mu         sync.Mutex
	messages   []map[string]any
	deletes    int
	deletedIDs []string
}

func (s *apiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_ = r.ParseMultipartForm(32 << 20)

	path := r.URL.Path
	switch {
	case strings.HasSuffix(path, "/getMe"):
		writeOK(w, map[string]any{
			"id": 1, "is_bot": true, "first_name": "Test", "username": "testbot",
		})
	case strings.HasSuffix(path, "/sendMessage"):
		id := len(s.messages) + 1
		m := map[string]any{
			"chat_id":      r.FormValue("chat_id"),
			"text":         r.FormValue("text"),
			"reply_markup": r.FormValue("reply_markup"),
			"message_id":   id,
			"kind":         "send",
		}
		s.messages = append(s.messages, m)
		writeOK(w, map[string]any{
			"message_id": id,
			"date":       time.Now().Unix(),
			"chat":       map[string]any{"id": 10, "type": "private"},
			"text":       m["text"],
		})
	case strings.HasSuffix(path, "/answerCallbackQuery"):
		writeOK(w, true)
	case strings.HasSuffix(path, "/editMessageText"):
		m := map[string]any{
			"chat_id":    r.FormValue("chat_id"),
			"message_id": r.FormValue("message_id"),
			"text":       r.FormValue("text"),
			"kind":       "edit",
		}
		s.messages = append(s.messages, m)
		writeOK(w, map[string]any{
			"message_id": 1,
			"date":       time.Now().Unix(),
			"chat":       map[string]any{"id": 10, "type": "private"},
			"text":       m["text"],
		})
	case strings.HasSuffix(path, "/deleteMessage"):
		s.deletes++
		s.deletedIDs = append(s.deletedIDs, r.FormValue("message_id"))
		writeOK(w, true)
	default:
		writeOK(w, true)
	}
}

func writeOK(w http.ResponseWriter, result any) {
	_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "result": result})
}

func startCommandUpdate(userID int64, username, text string) *models.Update {
	return &models.Update{
		ID: 1,
		Message: &models.Message{
			ID:   1,
			Date: int(time.Now().Unix()),
			Chat: models.Chat{ID: 10, Type: models.ChatTypePrivate},
			From: &models.User{ID: userID, FirstName: "U", Username: username},
			Text: text,
			Entities: []models.MessageEntity{
				{Type: models.MessageEntityTypeBotCommand, Offset: 0, Length: len(strings.SplitN(text, " ", 2)[0])},
			},
		},
	}
}

func callbackUpdate(userID int64, data string) *models.Update {
	return &models.Update{
		ID: 3,
		CallbackQuery: &models.CallbackQuery{
			ID:   "cb1",
			From: models.User{ID: userID, FirstName: "U", Username: "admin"},
			Data: data,
			Message: models.MaybeInaccessibleMessage{
				Type: models.MaybeInaccessibleMessageTypeMessage,
				Message: &models.Message{
					ID:   1,
					Date: int(time.Now().Unix()),
					Chat: models.Chat{ID: 10, Type: models.ChatTypePrivate},
				},
			},
		},
	}
}

func visibleTexts(s *apiServer) []string {
	var texts []string
	for _, m := range s.messages {
		t, ok := m["text"].(string)
		if !ok || t == "" || t == "\u2060" || t == "\u2800" {
			continue
		}
		texts = append(texts, t)
	}
	return texts
}

func TestUnauthorizedUserGetsDenyMessage(t *testing.T) {
	srv := &apiServer{}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	cfg := &config.Config{
		Telegram: config.TelegramConfig{
			API:          ts.URL,
			BotToken:     "123:ABC",
			AllowedUsers: []string{"999"},
		},
		Buttons: []config.ButtonNode{
			{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"},
		},
	}
	cfg.ApplyDefaults()

	app := bot.NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)
	b, err := app.NewBotWithOptions(
		tgbot.WithHTTPClient(5*time.Second, ts.Client()),
		tgbot.WithServerURL(ts.URL),
		tgbot.WithSkipGetMe(),
		tgbot.WithNotAsyncHandlers(),
	)
	require.NoError(t, err)

	b.ProcessUpdate(context.Background(), startCommandUpdate(1, "intruder", "/start"))

	srv.mu.Lock()
	defer srv.mu.Unlock()
	require.NotEmpty(t, srv.messages)
	text, _ := srv.messages[0]["text"].(string)
	require.Contains(t, text, "not allowed")
	require.Contains(t, text, "1")
	require.Contains(t, text, "intruder")
}

func TestAllowedStartSendsMenu(t *testing.T) {
	srv := &apiServer{}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	cfg := &config.Config{
		Telegram: config.TelegramConfig{
			API:          ts.URL,
			BotToken:     "123:ABC",
			AllowedUsers: []string{"42", "admin"},
		},
		Buttons: []config.ButtonNode{
			{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"},
		},
	}
	cfg.ApplyDefaults()

	app := bot.NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)
	b, err := app.NewBotWithOptions(
		tgbot.WithHTTPClient(5*time.Second, ts.Client()),
		tgbot.WithServerURL(ts.URL),
		tgbot.WithSkipGetMe(),
		tgbot.WithNotAsyncHandlers(),
	)
	require.NoError(t, err)

	b.ProcessUpdate(context.Background(), startCommandUpdate(42, "admin", "/start"))

	srv.mu.Lock()
	defer srv.mu.Unlock()
	require.Contains(t, visibleTexts(srv), "Menu")
}

func TestInlineMenuNavigation(t *testing.T) {
	srv := &apiServer{}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	cfg := &config.Config{
		Telegram: config.TelegramConfig{
			API:          ts.URL,
			BotToken:     "123:ABC",
			AllowedUsers: []string{"42"},
		},
		Buttons: []config.ButtonNode{
			{
				Name: "Cat",
				Type: "category",
				Icon: "📁",
				Items: []config.ButtonNode{
					{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"},
				},
			},
		},
	}
	cfg.ApplyDefaults()

	app := bot.NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)
	b, err := app.NewBotWithOptions(
		tgbot.WithHTTPClient(5*time.Second, ts.Client()),
		tgbot.WithServerURL(ts.URL),
		tgbot.WithSkipGetMe(),
		tgbot.WithNotAsyncHandlers(),
	)
	require.NoError(t, err)

	ctx := context.Background()
	b.ProcessUpdate(ctx, startCommandUpdate(42, "admin", "/start"))
	b.ProcessUpdate(ctx, callbackUpdate(42, "o:"+app.Index.Roots[0]))
	b.ProcessUpdate(ctx, callbackUpdate(42, "h"))

	srv.mu.Lock()
	defer srv.mu.Unlock()
	require.NotEmpty(t, srv.messages)
	texts := visibleTexts(srv)
	require.Contains(t, texts, "Menu")
	foundCat := false
	for _, line := range texts {
		if strings.Contains(line, "Cat") {
			foundCat = true
			break
		}
	}
	require.True(t, foundCat)
	require.Equal(t, "Menu", texts[len(texts)-1])
}

func TestMenuAlwaysSendsFreshMessageAndCleansUpOld(t *testing.T) {
	srv := &apiServer{}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	cfg := &config.Config{
		Telegram: config.TelegramConfig{
			API:          ts.URL,
			BotToken:     "123:ABC",
			AllowedUsers: []string{"42"},
		},
		Buttons: []config.ButtonNode{
			{
				Name: "Cat",
				Type: "category",
				Items: []config.ButtonNode{
					{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"},
				},
			},
		},
	}
	cfg.ApplyDefaults()

	app := bot.NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)
	b, err := app.NewBotWithOptions(
		tgbot.WithHTTPClient(5*time.Second, ts.Client()),
		tgbot.WithServerURL(ts.URL),
		tgbot.WithSkipGetMe(),
		tgbot.WithNotAsyncHandlers(),
	)
	require.NoError(t, err)

	ctx := context.Background()
	b.ProcessUpdate(ctx, startCommandUpdate(42, "admin", "/start"))
	b.ProcessUpdate(ctx, callbackUpdate(42, "o:"+app.Index.Roots[0]))
	b.ProcessUpdate(ctx, callbackUpdate(42, "h"))

	srv.mu.Lock()
	defer srv.mu.Unlock()

	// Every screen (Menu, Cat, Menu) must be a brand new sendMessage call,
	// never an editMessageText, so Telegram never plays its keyboard
	// "menu transition" resize animation that can make button text look
	// cut off or overlapping on Android.
	require.Equal(t, 3, len(visibleTexts(srv)), "each screen should be its own new message")
	// Each screen after the first deletes the previous menu message, so the
	// chat never fills up with old screens.
	require.Equal(t, 2, srv.deletes, "old menu messages should be cleaned up")
}

func TestMenuMessageCarriesHomeReplyKeyboard(t *testing.T) {
	srv := &apiServer{}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	cfg := &config.Config{
		Telegram: config.TelegramConfig{
			API:          ts.URL,
			BotToken:     "123:ABC",
			AllowedUsers: []string{"42"},
		},
		Buttons: []config.ButtonNode{
			{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"},
		},
	}
	cfg.ApplyDefaults()

	app := bot.NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)
	b, err := app.NewBotWithOptions(
		tgbot.WithHTTPClient(5*time.Second, ts.Client()),
		tgbot.WithServerURL(ts.URL),
		tgbot.WithSkipGetMe(),
		tgbot.WithNotAsyncHandlers(),
	)
	require.NoError(t, err)

	ctx := context.Background()
	b.ProcessUpdate(ctx, startCommandUpdate(42, "admin", "/start"))

	srv.mu.Lock()
	defer srv.mu.Unlock()

	// No separate near-blank message should be sent anymore; the Home button
	// now lives in the reply keyboard attached to the menu message itself.
	for _, m := range srv.messages {
		text, _ := m["text"].(string)
		require.NotEqual(t, "\u2800", text, "no blank keyboard-priming message should be sent")
	}

	// The menu message must carry a reply keyboard (shown under the message
	// box) that includes the Home button.
	require.NotEmpty(t, srv.messages)
	rm := fmt.Sprintf("%v", srv.messages[0]["reply_markup"])
	require.Contains(t, rm, "keyboard")
	require.Contains(t, rm, "Home")
}

func TestTypedStartAfterTapSendsNewMessage(t *testing.T) {
	srv := &apiServer{}
	ts := httptest.NewServer(srv)
	defer ts.Close()

	cfg := &config.Config{
		Telegram: config.TelegramConfig{
			API:          ts.URL,
			BotToken:     "123:ABC",
			AllowedUsers: []string{"42"},
		},
		Buttons: []config.ButtonNode{
			{
				Name: "Cat",
				Type: "category",
				Items: []config.ButtonNode{
					{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"},
				},
			},
		},
	}
	cfg.ApplyDefaults()

	app := bot.NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)
	b, err := app.NewBotWithOptions(
		tgbot.WithHTTPClient(5*time.Second, ts.Client()),
		tgbot.WithServerURL(ts.URL),
		tgbot.WithSkipGetMe(),
		tgbot.WithNotAsyncHandlers(),
	)
	require.NoError(t, err)

	ctx := context.Background()
	b.ProcessUpdate(ctx, startCommandUpdate(42, "admin", "/start"))
	b.ProcessUpdate(ctx, callbackUpdate(42, "o:"+app.Index.Roots[0])) // a button tap: edits in place
	b.ProcessUpdate(ctx, startCommandUpdate(42, "admin", "/start"))   // typed again: must be a new message

	srv.mu.Lock()
	defer srv.mu.Unlock()

	last := srv.messages[len(srv.messages)-1]
	require.Equal(t, "send", last["kind"], "the menu after a typed command must be a new message, not an edit of an older one")
	require.Equal(t, "Menu", last["text"])
}
