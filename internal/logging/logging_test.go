package logging_test

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/logging"
)

func TestSetupDefault(t *testing.T) {
	audit, err := logging.Setup(config.LoggingConfig{
		Logs: map[string]config.LoggerConfig{
			"default": {
				Level:  "debug",
				Format: "console",
				Output: []config.OutputConfig{{Output: "discard"}},
			},
			"audit": {
				Level:  "info",
				Format: "json",
				Output: []config.OutputConfig{{Output: "discard"}},
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, audit)
	slog.Info("logging ok")
}
