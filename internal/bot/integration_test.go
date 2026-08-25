package bot_test

import (
	"context"
	"encoding/json"
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
	mu       sync.Mutex
	messages []map[string]any
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
		m := map[string]any{
			"chat_id": r.FormValue("chat_id"),
			"text":    r.FormValue("text"),
		}
		s.messages = append(s.messages, m)
		writeOK(w, map[string]any{
			"message_id": len(s.messages),
			"date":       time.Now().Unix(),
			"chat":       map[string]any{"id": 10, "type": "private"},
			"text":       m["text"],
		})
	case strings.HasSuffix(path, "/answerCallbackQuery"):
		writeOK(w, true)
	case strings.HasSuffix(path, "/editMessageText"):
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

func textUpdate(userID int64, username, text string) *models.Update {
	return &models.Update{
		ID: 2,
		Message: &models.Message{
			ID:   2,
			Date: int(time.Now().Unix()),
			Chat: models.Chat{ID: 10, Type: models.ChatTypePrivate},
			From: &models.User{ID: userID, FirstName: "U", Username: username},
			Text: text,
		},
	}
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
	require.NotEmpty(t, srv.messages)
	text, _ := srv.messages[0]["text"].(string)
	require.Equal(t, "Menu", text)
}

func TestReplyKeyboardNavigation(t *testing.T) {
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
	b.ProcessUpdate(ctx, textUpdate(42, "admin", "📁 Cat"))
	b.ProcessUpdate(ctx, textUpdate(42, "admin", "🏠 Home"))

	srv.mu.Lock()
	defer srv.mu.Unlock()
	require.GreaterOrEqual(t, len(srv.messages), 3)
	require.Equal(t, "Menu", srv.messages[0]["text"])
	require.Contains(t, srv.messages[1]["text"], "Cat")
	require.Equal(t, "Menu", srv.messages[2]["text"])
}
