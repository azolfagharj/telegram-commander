// Package config defines the application configuration schema and loading helpers.
package config

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

// Config is the root configuration loaded from a YAML file.
type Config struct {
	Telegram          TelegramConfig    `yaml:"telegram"`
	FunctionDirectory string            `yaml:"function_directory"`
	Shell             string            `yaml:"shell"`
	Timeout           Duration          `yaml:"timeout"`
	MaxOutputBytes    int               `yaml:"max_output_bytes"`
	WorkDir           string            `yaml:"workdir"`
	Env               map[string]string `yaml:"env"`
	MenuColumns       int               `yaml:"menu_columns"`
	PageSize          int               `yaml:"page_size"`
	ConfirmTTL        Duration          `yaml:"confirm_ttl"`
	EnableRunCommand  bool              `yaml:"enable_run_command"`
	Logging           LoggingConfig     `yaml:"logging"`
	Menu              []ButtonNode      `yaml:"menu"`

	// functionDirectorySet is true when the key was present in YAML (even if empty).
	functionDirectorySet bool
}

// TelegramConfig holds Telegram Bot API settings.
type TelegramConfig struct {
	API          string      `yaml:"api"`
	BotToken     string      `yaml:"bot_token"`
	AllowedUsers []string    `yaml:"allowed_users"`
	Proxy        ProxyConfig `yaml:"proxy"`
	Insecure     bool        `yaml:"insecure"`
}

// ProxyConfig configures an optional HTTP/SOCKS proxy for Telegram API calls.
type ProxyConfig struct {
	Enabled bool   `yaml:"enabled"`
	URL     string `yaml:"url"`
}

// ButtonNode is a recursive menu node: either a category or an actionable button.
type ButtonNode struct {
	Name     string            `yaml:"name"`
	Type     string            `yaml:"type"` // category | button
	Icon     string            `yaml:"icon"`
	ID       string            `yaml:"id"`
	Function string            `yaml:"function"`
	Confirm  bool              `yaml:"confirm"`
	Timeout  *Duration         `yaml:"timeout"`
	WorkDir  string            `yaml:"workdir"`
	Env      map[string]string `yaml:"env"`
	Columns  *int              `yaml:"columns"`
	Items    []ButtonNode      `yaml:"items"`

	// Params holds extra key/value pairs passed to the function (e.g. command).
	// Captured via custom unmarshaling of remaining fields on button nodes.
	Params map[string]string `yaml:"-"`

	// Raw fields used during load for function-specific parameters.
	Command string `yaml:"command"`
	Path    string `yaml:"path"`
	Args    string `yaml:"args"`
}

type rawButtonNode ButtonNode

var buttonNodeFields = map[string]struct{}{
	"name": {}, "type": {}, "icon": {}, "id": {}, "function": {},
	"confirm": {}, "timeout": {}, "workdir": {}, "env": {},
	"columns": {}, "items": {}, "command": {}, "path": {}, "args": {},
}

// UnmarshalYAML decodes standard button fields and collects function parameters.
func (b *ButtonNode) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return fmt.Errorf("menu node must be a mapping")
	}

	var raw rawButtonNode
	if err := value.Decode(&raw); err != nil {
		return err
	}
	*b = ButtonNode(raw)

	for i := 0; i+1 < len(value.Content); i += 2 {
		key := value.Content[i].Value
		if _, known := buttonNodeFields[key]; known {
			continue
		}
		param := value.Content[i+1]
		if param.Kind != yaml.ScalarNode {
			return fmt.Errorf("button parameter %q must be a scalar value", key)
		}
		switch param.Tag {
		case "!!str", "!!int", "!!bool", "!!float":
		case "!!null":
			return fmt.Errorf("button parameter %q must not be null", key)
		default:
			return fmt.Errorf("button parameter %q has unsupported scalar type %q", key, param.Tag)
		}
		if b.Params == nil {
			b.Params = make(map[string]string)
		}
		b.Params[key] = param.Value
	}
	return nil
}

// LoggingConfig holds named loggers similar to Caddy's logging block.
type LoggingConfig struct {
	Logs map[string]LoggerConfig `yaml:"logs"`
}

// LoggerConfig configures a single named logger.
type LoggerConfig struct {
	Level  string         `yaml:"level"`
	Format string         `yaml:"format"` // json | console
	Output []OutputConfig `yaml:"output"`
}

// OutputConfig describes where a logger writes.
type OutputConfig struct {
	Output   string `yaml:"output"` // stdout | stderr | file | discard
	File     string `yaml:"file"`
	RollSize int    `yaml:"roll_size_mb"`
	RollKeep int    `yaml:"roll_keep"`
}

// Duration wraps time.Duration for YAML unmarshaling from strings like "30s".
type Duration struct {
	time.Duration
}

// UnmarshalYAML parses a duration string.
func (d *Duration) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	if s == "" {
		d.Duration = 0
		return nil
	}
	parsed, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.Duration = parsed
	return nil
}

// MarshalYAML writes the duration as a string.
func (d Duration) MarshalYAML() (any, error) {
	return d.Duration.String(), nil
}

// Defaults applied when fields are omitted.
const (
	DefaultAPI            = "https://api.telegram.org"
	DefaultShell          = "/bin/bash"
	DefaultTimeout        = 60 * time.Second
	DefaultMaxOutputBytes = 512 * 1024
	DefaultMenuColumns    = 2
	DefaultPageSize       = 8
	DefaultConfirmTTL     = 5 * time.Minute
)

// ApplyDefaults fills zero-value fields with defaults.
func (c *Config) ApplyDefaults() {
	if c.Telegram.API == "" {
		c.Telegram.API = DefaultAPI
	}
	if c.Shell == "" {
		c.Shell = DefaultShell
	}
	if c.Timeout.Duration == 0 {
		c.Timeout.Duration = DefaultTimeout
	}
	if c.MaxOutputBytes == 0 {
		c.MaxOutputBytes = DefaultMaxOutputBytes
	}
	if c.MenuColumns == 0 {
		c.MenuColumns = DefaultMenuColumns
	}
	if c.PageSize == 0 {
		c.PageSize = DefaultPageSize
	}
	if c.ConfirmTTL.Duration == 0 {
		c.ConfirmTTL.Duration = DefaultConfirmTTL
	}
	if c.Logging.Logs == nil {
		c.Logging.Logs = map[string]LoggerConfig{
			"default": {
				Level:  "info",
				Format: "console",
				Output: []OutputConfig{{Output: "stderr"}},
			},
		}
	}
}

// FunctionDirectoryWasSet reports whether function_directory appeared in YAML.
func (c *Config) FunctionDirectoryWasSet() bool {
	return c.functionDirectorySet
}
