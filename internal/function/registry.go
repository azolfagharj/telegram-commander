// Package function loads and validates custom and built-in functions.
package function

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/azolfagharj/telegram-commander/internal/config"
)

var namePattern = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

// ReservedNames cannot be overridden by user-defined function files.
var ReservedNames = map[string]struct{}{
	"command": {},
	"script":  {},
}

// ParamSpec describes a single function parameter.
type ParamSpec struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"` // string | int | bool
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
	Description string `yaml:"description"`
}

// Definition is a function definition loaded from YAML or built-in.
type Definition struct {
	Name   string      `yaml:"name"`
	Run    string      `yaml:"run"`
	Params []ParamSpec `yaml:"params"`
	Source string      `yaml:"-"` // file path or "builtin"
}

// Registry holds all available functions keyed by their exact name.
type Registry struct {
	byName  map[string]*Definition
	byLower map[string]string // lower -> exact name
}

// NewRegistry creates a registry with built-in functions registered.
func NewRegistry() *Registry {
	r := &Registry{
		byName:  make(map[string]*Definition),
		byLower: make(map[string]string),
	}
	r.mustRegister(builtinCommand())
	r.mustRegister(builtinScript())
	return r
}

func (r *Registry) mustRegister(def *Definition) {
	if err := r.register(def); err != nil {
		panic(err)
	}
}

func (r *Registry) register(def *Definition) error {
	if def == nil || def.Name == "" {
		return fmt.Errorf("function name is required")
	}
	if !namePattern.MatchString(def.Name) {
		return fmt.Errorf("invalid function name %q: only letters, digits, '.', '-', '_' allowed", def.Name)
	}
	lower := strings.ToLower(def.Name)
	if existing, ok := r.byLower[lower]; ok {
		return fmt.Errorf("duplicate function name %q (conflicts with %q)", def.Name, existing)
	}
	r.byName[def.Name] = def
	r.byLower[lower] = def.Name
	return nil
}

// Get returns a function by name (case-sensitive lookup of the stored name).
func (r *Registry) Get(name string) (*Definition, bool) {
	def, ok := r.byName[name]
	return def, ok
}

// GetCaseInsensitive finds a function ignoring case.
func (r *Registry) GetCaseInsensitive(name string) (*Definition, bool) {
	exact, ok := r.byLower[strings.ToLower(name)]
	if !ok {
		return nil, false
	}
	return r.byName[exact], true
}

// List returns all definitions sorted by name.
func (r *Registry) List() []*Definition {
	out := make([]*Definition, 0, len(r.byName))
	for _, d := range r.byName {
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func builtinCommand() *Definition {
	return &Definition{
		Name:   "command",
		Source: "builtin",
		Run:    "{{.command}}",
		Params: []ParamSpec{
			{Name: "command", Type: "string", Required: true, Description: "Shell command to execute"},
		},
	}
}

func builtinScript() *Definition {
	return &Definition{
		Name:   "script",
		Source: "builtin",
		Run:    "{{.path}}{{if .args}} {{.args}}{{end}}",
		Params: []ParamSpec{
			{Name: "path", Type: "string", Required: true, Description: "Path to the script"},
			{Name: "args", Type: "string", Required: false, Description: "Optional arguments"},
		},
	}
}

// LoadDirectory recursively loads function YAML files from dir.
// Empty directory is fine. Missing directory should already be caught by config validation
// when the path was set; if called with empty path, it is a no-op.
func (r *Registry) LoadDirectory(dir string) error {
	if strings.TrimSpace(dir) == "" {
		return nil
	}
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}
		def, err := loadFile(path)
		if err != nil {
			return err
		}
		lower := strings.ToLower(def.Name)
		if _, reserved := ReservedNames[lower]; reserved {
			return fmt.Errorf("%s: function name %q is reserved", path, def.Name)
		}
		if err := r.register(def); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		return nil
	})
}

func loadFile(path string) (*Definition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var def Definition
	dec := yaml.NewDecoder(strings.NewReader(string(data)))
	dec.KnownFields(true)
	if err := dec.Decode(&def); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	if def.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if def.Run == "" {
		return nil, fmt.Errorf("run is required")
	}
	for i, p := range def.Params {
		if p.Name == "" {
			return nil, fmt.Errorf("params[%d].name is required", i)
		}
		if !namePattern.MatchString(p.Name) {
			return nil, fmt.Errorf("params[%d].name %q is invalid", i, p.Name)
		}
		switch strings.ToLower(p.Type) {
		case "", "string", "int", "bool":
			if p.Type == "" {
				def.Params[i].Type = "string"
			}
		default:
			return nil, fmt.Errorf("params[%d].type %q is unsupported", i, p.Type)
		}
	}
	def.Source = path
	return &def, nil
}

// ValidateButtonParams checks that a button provides required params for its function.
func (r *Registry) ValidateButtonParams(node config.ButtonNode, path string) config.ValidationErrors {
	var errs config.ValidationErrors
	if strings.ToLower(node.Type) != "button" {
		for i, child := range node.Items {
			childPath := fmt.Sprintf("%s.items[%d]", path, i)
			errs = append(errs, r.ValidateButtonParams(child, childPath)...)
		}
		return errs
	}
	def, ok := r.GetCaseInsensitive(node.Function)
	if !ok {
		errs = append(errs, config.ValidationError{
			Path:    path + ".function",
			Message: fmt.Sprintf("unknown function %q", node.Function),
		})
		return errs
	}
	values := buttonParams(node)
	for _, p := range def.Params {
		v, has := values[p.Name]
		if p.Required && (!has || strings.TrimSpace(v) == "") {
			errs = append(errs, config.ValidationError{
				Path:    path + "." + p.Name,
				Message: fmt.Sprintf("required parameter %q for function %q is missing", p.Name, def.Name),
			})
		}
	}
	return errs
}

// ValidateTree validates function references for the whole button tree.
func (r *Registry) ValidateTree(buttons []config.ButtonNode) config.ValidationErrors {
	var errs config.ValidationErrors
	for i, n := range buttons {
		errs = append(errs, r.ValidateButtonParams(n, fmt.Sprintf("buttons[%d]", i))...)
	}
	return errs
}

// ButtonParams extracts parameter values from a button node.
func ButtonParams(node config.ButtonNode) map[string]string {
	return buttonParams(node)
}

func buttonParams(node config.ButtonNode) map[string]string {
	m := map[string]string{}
	if node.Command != "" {
		m["command"] = node.Command
	}
	if node.Path != "" {
		m["path"] = node.Path
	}
	if node.Args != "" {
		m["args"] = node.Args
	}
	for k, v := range node.Params {
		m[k] = v
	}
	return m
}
