// Package logging configures slog handlers from the config logging block.
package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/azolfagharj/telegram-commander/internal/config"
)

// Setup installs the default slog logger from cfg and returns an audit logger.
func Setup(cfg config.LoggingConfig) (audit *slog.Logger, err error) {
	defaultCfg, ok := cfg.Logs["default"]
	if !ok {
		for _, v := range cfg.Logs {
			defaultCfg = v
			ok = true
			break
		}
	}
	if !ok {
		defaultCfg = config.LoggerConfig{
			Level:  "info",
			Format: "console",
			Output: []config.OutputConfig{{Output: "stderr"}},
		}
	}

	handler, err := buildHandler(defaultCfg)
	if err != nil {
		return nil, err
	}
	slog.SetDefault(slog.New(handler))

	auditCfg, hasAudit := cfg.Logs["audit"]
	if !hasAudit {
		auditCfg = config.LoggerConfig{
			Level:  "info",
			Format: "json",
			Output: []config.OutputConfig{{Output: "stderr"}},
		}
	}
	auditHandler, err := buildHandler(auditCfg)
	if err != nil {
		return nil, err
	}
	return slog.New(auditHandler).With("logger", "audit"), nil
}

func buildHandler(cfg config.LoggerConfig) (slog.Handler, error) {
	level := parseLevel(cfg.Level)
	opts := &slog.HandlerOptions{Level: level}

	writers := make([]io.Writer, 0, len(cfg.Output))
	if len(cfg.Output) == 0 {
		writers = append(writers, os.Stderr)
	}
	for _, out := range cfg.Output {
		w, err := openOutput(out)
		if err != nil {
			return nil, err
		}
		if w == nil {
			continue // discard
		}
		writers = append(writers, w)
	}
	if len(writers) == 0 {
		return slog.NewTextHandler(io.Discard, opts), nil
	}
	var w io.Writer
	if len(writers) == 1 {
		w = writers[0]
	} else {
		w = io.MultiWriter(writers...)
	}

	switch strings.ToLower(cfg.Format) {
	case "json":
		return slog.NewJSONHandler(w, opts), nil
	case "console", "":
		return slog.NewTextHandler(w, opts), nil
	default:
		return nil, fmt.Errorf("unsupported log format %q", cfg.Format)
	}
}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

var openFiles sync.Map // path -> *os.File (kept open for process lifetime)

func openOutput(out config.OutputConfig) (io.Writer, error) {
	switch strings.ToLower(out.Output) {
	case "stdout", "":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	case "discard":
		return nil, nil
	case "file":
		if out.File == "" {
			return nil, fmt.Errorf("file output requires file path")
		}
		if v, ok := openFiles.Load(out.File); ok {
			return v.(*os.File), nil
		}
		f, err := os.OpenFile(out.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			return nil, fmt.Errorf("open log file %q: %w", out.File, err)
		}
		openFiles.Store(out.File, f)
		return f, nil
	default:
		return nil, fmt.Errorf("unsupported output %q", out.Output)
	}
}
