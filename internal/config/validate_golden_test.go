package config_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/config"
)

func TestValidationErrorsGolden(t *testing.T) {
	yaml := `
telegram:
  bot_token: ""
  allowed_users: []
menu: []
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	errs := cfg.Validate()
	require.Error(t, errs.Err())

	var buf bytes.Buffer
	for _, e := range errs {
		buf.WriteString(e.Path)
		buf.WriteString(": ")
		buf.WriteString(e.Message)
		buf.WriteByte('\n')
	}
	got := buf.String()
	// Stable subset checks (order of map-like fields may vary; our validator is sequential).
	require.True(t, strings.Contains(got, "telegram.bot_token: bot_token is required"))
	require.True(t, strings.Contains(got, "telegram.allowed_users: at least one allowed user is required"))
	require.True(t, strings.Contains(got, "menu: at least one button or category is required"))
}

func TestCategoryParameterErrorsGolden(t *testing.T) {
	cfg, err := config.Parse([]byte(`
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - name: Tools
    type: category
    zeta: last
    alpha: first
    items:
      - name: Echo
        type: button
        function: command
        command: echo hi
`))
	require.NoError(t, err)
	require.EqualError(t, cfg.Validate().Err(), "menu[0]: category has unsupported parameter keys: alpha, zeta")
}
