package config_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/config"
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
