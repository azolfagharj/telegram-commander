package cli_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/cli"
)

func TestValidateExampleConfig(t *testing.T) {
	root, err := filepath.Abs("../..")
	require.NoError(t, err)
	cfgPath := filepath.Join(root, "examples", "config.yaml")
	require.FileExists(t, cfgPath)

	wd, err := os.Getwd()
	require.NoError(t, err)
	require.NoError(t, os.Chdir(root))
	t.Cleanup(func() { _ = os.Chdir(wd) })

	cmd := cli.NewRoot("1.0.0", "2026-08-25")
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"validate", "--config", "examples/config.yaml"})
	err = cmd.Execute()
	require.NoError(t, err, buf.String())
	require.Contains(t, buf.String(), "Valid configuration")
}

func TestVersion(t *testing.T) {
	cmd := cli.NewRoot("1.0.0", "2026-08-25")
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"version"})
	require.NoError(t, cmd.Execute())
	require.Contains(t, buf.String(), "telegram-commander 1.0.0 (2026-08-25)")
}
