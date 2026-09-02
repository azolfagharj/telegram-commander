package function_test

import (
	"fmt"
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

func loadDefinition(t *testing.T, body string) (*function.Registry, *function.Definition) {
	t.Helper()
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(body), 0o644))
	r := function.NewRegistry()
	require.NoError(t, r.LoadDirectory(dir))
	def, ok := r.Get("test-fn")
	require.True(t, ok)
	return r, def
}

func TestReservedButtonFieldParamNamesRejected(t *testing.T) {
	names := []string{"name", "type", "icon", "id", "function", "confirm", "timeout", "workdir", "env", "columns", "items"}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			body := fmt.Sprintf("name: test-fn\nrun: echo\nparams:\n  - name: %s\n", name)
			require.NoError(t, os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(body), 0o644))
			err := function.NewRegistry().LoadDirectory(dir)
			require.Error(t, err)
			require.Contains(t, err.Error(), "conflicts with a button field")
			require.Contains(t, err.Error(), name)
		})
	}
}

func TestCommandPathAndArgsParamNamesAllowed(t *testing.T) {
	_, def := loadDefinition(t, `
name: test-fn
run: "{{.command}} {{.path}} {{.args}}"
params:
  - name: command
  - name: path
  - name: args
`)
	require.NoError(t, def.ValidateTemplate())
}

func TestParamDefaultsMatchTypes(t *testing.T) {
	tests := []struct {
		name      string
		paramType string
		value     string
		wantError bool
	}{
		{name: "bad int", paramType: "int", value: "abc", wantError: true},
		{name: "bad bool", paramType: "bool", value: "yes123", wantError: true},
		{name: "good int", paramType: "int", value: "-12"},
		{name: "good bool", paramType: "bool", value: "false"},
		{name: "good string", paramType: "string", value: "anything"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			body := fmt.Sprintf(`
name: test-fn
run: "{{.value}}"
params:
  - name: value
    type: %s
    default: %q
`, tt.paramType, tt.value)
			require.NoError(t, os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(body), 0o644))
			err := function.NewRegistry().LoadDirectory(dir)
			if tt.wantError {
				require.Error(t, err)
				require.Contains(t, err.Error(), "default")
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParamTypeValidationOnLoad(t *testing.T) {
	t.Run("unsupported", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(dir, "test.yaml"), []byte(`
name: test-fn
run: echo
params:
  - name: value
    type: number
`), 0o644))
		err := function.NewRegistry().LoadDirectory(dir)
		require.Error(t, err)
		require.Contains(t, err.Error(), "unsupported")
	})
	t.Run("empty defaults to string", func(t *testing.T) {
		_, def := loadDefinition(t, `
name: test-fn
run: "{{.value}}"
params:
  - name: value
`)
		require.Equal(t, "string", def.Params[0].Type)
	})
}

func TestValidateButtonParamsRules(t *testing.T) {
	r, _ := loadDefinition(t, `
name: test-fn
run: "{{.required}} {{.number}} {{.enabled}}"
params:
  - name: required
    required: true
  - name: optional
  - name: number
    type: int
  - name: enabled
    type: bool
`)
	tests := []struct {
		name   string
		params map[string]string
		want   []string
	}{
		{name: "valid negative and false", params: map[string]string{"required": "x", "number": "-1", "enabled": "false"}},
		{name: "valid zero", params: map[string]string{"required": "x", "number": "0", "enabled": "true"}},
		{name: "optional missing", params: map[string]string{"required": "x"}},
		{name: "unknown", params: map[string]string{"required": "x", "mystery": "y"}, want: []string{`unknown parameter "mystery"`, `"test-fn"`}},
		{name: "required missing", params: map[string]string{}, want: []string{`required parameter "required"`}},
		{name: "required whitespace", params: map[string]string{"required": "   "}, want: []string{`required parameter "required"`}},
		{name: "bad int", params: map[string]string{"required": "x", "number": "abc"}, want: []string{`parameter "number"`, "must be an int"}},
		{name: "bad bool", params: map[string]string{"required": "x", "enabled": "maybe"}, want: []string{`parameter "enabled"`, "must be a bool"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := r.ValidateButtonParams(config.ButtonNode{
				Name: "Run", Type: "button", Function: "test-fn", Params: tt.params,
			}, "menu[0]")
			if len(tt.want) == 0 {
				require.NoError(t, errs.Err())
				return
			}
			require.Error(t, errs.Err())
			for _, want := range tt.want {
				require.Contains(t, errs.Error(), want)
			}
		})
	}
}

func TestExtraParamOnBuiltinCommandRejected(t *testing.T) {
	errs := function.NewRegistry().ValidateTree([]config.ButtonNode{{
		Name: "Echo", Type: "button", Function: "command", Command: "echo hi",
		Params: map[string]string{"extra": "x"},
	}})
	require.Error(t, errs.Err())
	require.Contains(t, errs.Error(), `unknown parameter "extra" for function "command"`)
}

func TestNestedParamErrorIncludesFullPath(t *testing.T) {
	r, _ := loadDefinition(t, `
name: test-fn
run: "{{.value}}"
params:
  - name: value
    required: true
`)
	errs := r.ValidateTree([]config.ButtonNode{{
		Name: "One", Type: "category", Items: []config.ButtonNode{{
			Name: "Two", Type: "category", Items: []config.ButtonNode{{
				Name: "Run", Type: "button", Function: "test-fn",
			}},
		}},
	}})
	require.Error(t, errs.Err())
	require.Contains(t, errs.Error(), "menu[0].items[0].items[0].value")
}

func TestValidateTemplateFields(t *testing.T) {
	tests := []struct {
		name    string
		run     string
		params  []function.ParamSpec
		wantErr string
	}{
		{name: "undefined direct", run: "{{.x}}", wantErr: `undefined parameter "x"`},
		{name: "undefined if", run: "{{if .x}}yes{{end}}", wantErr: `undefined parameter "x"`},
		{name: "all declared", run: "{{.x}} {{if .enabled}}{{.y}}{{end}}", params: []function.ParamSpec{{Name: "x"}, {Name: "enabled"}, {Name: "y"}}},
		{name: "broken syntax", run: "{{if .x}}", params: []function.ParamSpec{{Name: "x"}}, wantErr: "invalid run template"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			def := function.Definition{Name: "test", Run: tt.run, Params: tt.params}
			err := def.ValidateTemplate()
			if tt.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestReleaseFunctionFilesLoadAndValidate(t *testing.T) {
	r := function.NewRegistry()
	require.NoError(t, r.LoadDirectory(filepath.Join("..", "..", "examples", "functions")))
	names := []string{"curl-url", "disk-path", "echo-script", "journal-unit", "ping-host"}
	for _, name := range names {
		def, ok := r.Get(name)
		require.True(t, ok, name)
		require.NoError(t, def.ValidateTemplate(), name)
	}
}

func TestRenderRunWithParamsDefaultsAndBranches(t *testing.T) {
	def := function.Definition{
		Name: "render",
		Run:  `run {{.value}} {{.count}}{{if .flag}} --flag={{.flag}}{{end}}`,
		Params: []function.ParamSpec{
			{Name: "value", Required: true},
			{Name: "count", Default: "4"},
			{Name: "flag"},
		},
	}
	tests := []struct {
		name   string
		params map[string]string
		want   string
	}{
		{name: "button value and defaults", params: map[string]string{"value": "hello"}, want: "run hello 4"},
		{name: "if branch omitted", params: map[string]string{"value": "hello", "count": "2"}, want: "run hello 2"},
		{name: "if branch included", params: map[string]string{"value": "hello", "flag": "yes"}, want: "run hello 4 --flag=yes"},
		{name: "special characters unchanged", params: map[string]string{"value": `a b;$HOME"'`}, want: `run a b;$HOME"' 4`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := def.RenderRun(tt.params)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
