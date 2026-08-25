package function_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/function"
)

func TestBuiltinsPresent(t *testing.T) {
	r := function.NewRegistry()
	_, ok := r.Get("command")
	require.True(t, ok)
	_, ok = r.Get("script")
	require.True(t, ok)
}

func TestReservedNameRejected(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cmd.yaml")
	require.NoError(t, os.WriteFile(path, []byte(`
name: command
run: "echo {{.x}}"
params:
  - name: x
    required: true
`), 0o644))
	r := function.NewRegistry()
	err := r.LoadDirectory(dir)
	require.Error(t, err)
	require.Contains(t, err.Error(), "reserved")
}

func TestDuplicateCaseInsensitive(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "a.yaml"), []byte(`
name: MyFunc
run: "echo 1"
`), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "b.yaml"), []byte(`
name: myfunc
run: "echo 2"
`), 0o644))
	r := function.NewRegistry()
	err := r.LoadDirectory(dir)
	require.Error(t, err)
	require.Contains(t, err.Error(), "duplicate")
}

func TestRenderCommand(t *testing.T) {
	r := function.NewRegistry()
	def, ok := r.Get("command")
	require.True(t, ok)
	out, err := def.RenderRun(map[string]string{"command": "echo hi"})
	require.NoError(t, err)
	require.Equal(t, "echo hi", out)
}

func TestValidateTreeUnknownFunction(t *testing.T) {
	r := function.NewRegistry()
	errs := r.ValidateTree([]config.ButtonNode{
		{Name: "X", Type: "button", Function: "nope"},
	})
	require.Error(t, errs.Err())
	require.Contains(t, errs.Error(), "unknown function")
}

func TestLoadCustomFunction(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "echo-script.yaml"), []byte(`
name: echo-script
run: "bash {{.path}}{{if .args}} {{.args}}{{end}}"
params:
  - name: path
    type: string
    required: true
  - name: args
    type: string
    required: false
`), 0o644))
	r := function.NewRegistry()
	require.NoError(t, r.LoadDirectory(dir))
	def, ok := r.Get("echo-script")
	require.True(t, ok)
	out, err := def.RenderRun(map[string]string{"path": "/tmp/x.sh", "args": "--flag"})
	require.NoError(t, err)
	require.Equal(t, "bash /tmp/x.sh --flag", out)
}
