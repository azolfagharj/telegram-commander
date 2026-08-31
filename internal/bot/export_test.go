package bot

import (
	"strings"
	"time"

	"github.com/go-telegram/bot"
)

// NewBotWithOptions creates a bot merging app handlers with extra options (for tests).
func (a *App) NewBotWithOptions(extra ...bot.Option) (*bot.Bot, error) {
	client, err := HTTPClient(a.Cfg.Telegram)
	if err != nil {
		return nil, err
	}
	api := strings.TrimRight(a.Cfg.Telegram.API, "/")
	opts := []bot.Option{
		bot.WithHTTPClient(60*time.Second, client),
		bot.WithServerURL(api),
		bot.WithDefaultHandler(a.defaultHandler),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, a.handleCallback),
		bot.WithMessageTextHandler("start", bot.MatchTypeCommandStartOnly, a.handleStart),
		bot.WithMessageTextHandler("help", bot.MatchTypeCommandStartOnly, a.handleHelp),
	}
	if a.Cfg.Telegram.EnableRunCommand {
		opts = append(opts, bot.WithMessageTextHandler("run", bot.MatchTypeCommandStartOnly, a.handleRun))
	}
	opts = append(opts, extra...)
	return bot.New(a.Cfg.Telegram.BotToken, opts...)
}
