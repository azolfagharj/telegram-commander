package config_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/function"
)

func TestParseAndValidateMinimal(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
menu:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	require.Equal(t, "/bin/bash", cfg.Shell)
	require.Equal(t, config.DefaultAPI, cfg.Telegram.API)
	errs := cfg.Validate()
	require.NoError(t, errs.Err())
}

func TestUnknownFieldRejected(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
  unknown_thing: true
menu:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	_, err := config.Parse([]byte(yaml))
	require.Error(t, err)
	require.Contains(t, strings.ToLower(err.Error()), "unknown")
}

func TestDuplicateButtonNames(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
menu:
  - name: A
    type: button
    function: command
    command: "echo 1"
  - name: a
    type: button
    function: command
    command: "echo 2"
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	errs := cfg.Validate()
	require.Error(t, errs.Err())
	require.Contains(t, errs.Error(), "duplicate")
}

func TestMissingCommandForCommandFunction(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
menu:
  - name: Bad
    type: button
    function: command
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	errs := cfg.Validate()
	require.Error(t, errs.Err())
	require.Contains(t, errs.Error(), "command is required")
}

func TestFunctionDirectoryMissingIsHardError(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
function_directory: "/tmp/does-not-exist-telegram-commander-xyz"
menu:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	require.True(t, cfg.FunctionDirectoryWasSet())
	errs := cfg.Validate()
	require.Error(t, errs.Err())
	require.Contains(t, errs.Error(), "does not exist")
}

func TestFunctionDirectoryEmptyIsOk(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
function_directory: ""
menu:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	require.True(t, cfg.FunctionDirectoryWasSet())
	errs := cfg.Validate()
	require.NoError(t, errs.Err())
}

func TestOldButtonsKeyRejected(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
buttons:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	_, err := config.Parse([]byte(yaml))
	require.Error(t, err)
	require.Contains(t, strings.ToLower(err.Error()), "buttons")
	require.Contains(t, strings.ToLower(err.Error()), "not found")
}

func TestOldButtonsColumnsKeyRejected(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
buttons_columns: 2
menu:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	_, err := config.Parse([]byte(yaml))
	require.Error(t, err)
	require.Contains(t, strings.ToLower(err.Error()), "buttons_columns")
	require.Contains(t, strings.ToLower(err.Error()), "not found")
}

func TestRootEnableRunCommandAccepted(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
enable_run_command: true
menu:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	require.True(t, cfg.EnableRunCommand)
	require.NoError(t, cfg.Validate().Err())
}

func TestOldNestedEnableRunCommandRejected(t *testing.T) {
	yaml := `
telegram:
  bot_token: "token"
  allowed_users: ["1"]
  enable_run_command: true
menu:
  - name: Echo
    type: button
    function: command
    command: "echo hi"
`
	_, err := config.Parse([]byte(yaml))
	require.Error(t, err)
	require.Contains(t, strings.ToLower(err.Error()), "enable_run_command")
	require.Contains(t, strings.ToLower(err.Error()), "not found")
}

func TestInlineButtonParamsParsing(t *testing.T) {
	yaml := `
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - name: Check
    type: button
    function: custom
    quoted: "hello world"
    plain: hello
    integer: 100
    decimal: 1.25
    enabled: true
    disabled: false
`
	cfg, err := config.Parse([]byte(yaml))
	require.NoError(t, err)
	require.Equal(t, map[string]string{
		"quoted": "hello world", "plain": "hello", "integer": "100",
		"decimal": "1.25", "enabled": "true", "disabled": "false",
	}, cfg.Menu[0].Params)
}

func TestInlineButtonParamInvalidValues(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "mapping", value: "{nested: value}", want: "scalar"},
		{name: "sequence", value: "[one, two]", want: "scalar"},
		{name: "null", value: "null", want: "must not be null"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml := fmt.Sprintf(`
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - name: Check
    type: button
    function: custom
    extra: %s
`, tt.value)
			_, err := config.Parse([]byte(yaml))
			require.Error(t, err)
			require.Contains(t, err.Error(), "extra")
			require.Contains(t, err.Error(), tt.want)
		})
	}
}

func TestMenuNodeMustBeMapping(t *testing.T) {
	_, err := config.Parse([]byte(`
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - ""
`))
	require.Error(t, err)
	require.Contains(t, err.Error(), "menu node must be a mapping")
}

func TestButtonWithoutInlineParams(t *testing.T) {
	cfg, err := config.Parse([]byte(`
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - name: Echo
    type: button
    function: command
    command: echo hi
`))
	require.NoError(t, err)
	require.Empty(t, cfg.Menu[0].Params)
}

func TestNestedInlineButtonParam(t *testing.T) {
	cfg, err := config.Parse([]byte(`
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - name: Tools
    type: category
    items:
      - name: Ping
        type: button
        function: ping-host
        host: example.com
`))
	require.NoError(t, err)
	require.Equal(t, "example.com", cfg.Menu[0].Items[0].Params["host"])
}

func TestDuplicateButtonKeyRejected(t *testing.T) {
	_, err := config.Parse([]byte(`
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - name: Ping
    name: Again
    type: button
    function: command
    command: echo hi
`))
	require.Error(t, err)
	require.Contains(t, strings.ToLower(err.Error()), "already defined")
}

func TestStrictFieldsDoNotLeak(t *testing.T) {
	tests := []struct {
		name string
		yaml string
		key  string
	}{
		{name: "root", key: "unknown_root", yaml: "unknown_root: true"},
		{name: "telegram", key: "unknown_telegram", yaml: "telegram:\n  bot_token: token\n  allowed_users: [\"1\"]\n  unknown_telegram: true"},
		{name: "logging", key: "unknown_logging", yaml: "logging:\n  unknown_logging: true"},
		{name: "proxy", key: "unknown_proxy", yaml: "telegram:\n  bot_token: token\n  allowed_users: [\"1\"]\n  proxy:\n    unknown_proxy: true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prefix := "telegram:\n  bot_token: token\n  allowed_users: [\"1\"]\n"
			if strings.HasPrefix(tt.yaml, "telegram:") {
				prefix = ""
			}
			_, err := config.Parse([]byte(prefix + tt.yaml + `
menu:
  - name: Echo
    type: button
    function: command
    command: echo hi
`))
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.key)
		})
	}
}

func TestStructuralButtonTypoBecomesUnknownParam(t *testing.T) {
	cfg, err := config.Parse([]byte(`
telegram:
  bot_token: token
  allowed_users: ["1"]
menu:
  - name: Echo
    type: button
    function: command
    command: echo hi
    confrm: true
`))
	require.NoError(t, err)
	errs := function.NewRegistry().ValidateTree(cfg.Menu)
	require.Error(t, errs.Err())
	require.Contains(t, errs.Error(), `unknown parameter "confrm" for function "command"`)
}

func TestCategoryInlineParamsValidation(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]string
		want   []string
	}{
		{name: "none"},
		{name: "one", params: map[string]string{"host": "x"}, want: []string{"host"}},
		{name: "many", params: map[string]string{"zeta": "x", "alpha": "y"}, want: []string{"alpha", "zeta"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.Config{
				Telegram: config.TelegramConfig{BotToken: "token", AllowedUsers: []string{"1"}},
				Menu: []config.ButtonNode{{
					Name: "Tools", Type: "category", Params: tt.params,
					Items: []config.ButtonNode{{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"}},
				}},
			}
			cfg.ApplyDefaults()
			errs := cfg.Validate()
			if len(tt.want) == 0 {
				require.NoError(t, errs.Err())
				return
			}
			require.Error(t, errs.Err())
			for _, key := range tt.want {
				require.Contains(t, errs.Error(), key)
			}
		})
	}
}

func TestUnsupportedButtonTypeStillRejected(t *testing.T) {
	cfg := &config.Config{
		Telegram: config.TelegramConfig{BotToken: "token", AllowedUsers: []string{"1"}},
		Menu:     []config.ButtonNode{{Name: "Bad", Type: "link"}},
	}
	cfg.ApplyDefaults()
	require.Contains(t, cfg.Validate().Error(), `unsupported type "link"`)
}
