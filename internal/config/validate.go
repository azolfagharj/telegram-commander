package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Validate checks structural rules of the loaded config.
// It does not load functions; that is done by the function package.
func (c *Config) Validate() ValidationErrors {
	var errs ValidationErrors

	if strings.TrimSpace(c.Telegram.BotToken) == "" {
		errs = append(errs, ValidationError{
			Path:    "telegram.bot_token",
			Message: "bot_token is required",
		})
	}
	if c.Telegram.API == "" {
		errs = append(errs, ValidationError{
			Path:    "telegram.api",
			Message: "api must not be empty",
		})
	}
	if len(c.Telegram.AllowedUsers) == 0 {
		errs = append(errs, ValidationError{
			Path:    "telegram.allowed_users",
			Message: "at least one allowed user is required",
		})
	}
	for i, u := range c.Telegram.AllowedUsers {
		if strings.TrimSpace(u) == "" {
			errs = append(errs, ValidationError{
				Path:    fmt.Sprintf("telegram.allowed_users[%d]", i),
				Message: "allowed user must not be empty",
			})
		}
	}
	if c.Telegram.Proxy.Enabled {
		if strings.TrimSpace(c.Telegram.Proxy.URL) == "" {
			errs = append(errs, ValidationError{
				Path:    "telegram.proxy.url",
				Message: "proxy.url is required when proxy.enabled is true",
			})
		}
	}
	if c.Shell == "" {
		errs = append(errs, ValidationError{
			Path:    "shell",
			Message: "shell must not be empty",
		})
	}
	if c.Timeout.Duration <= 0 {
		errs = append(errs, ValidationError{
			Path:    "timeout",
			Message: "timeout must be positive",
		})
	}
	if c.MaxOutputBytes <= 0 {
		errs = append(errs, ValidationError{
			Path:    "max_output_bytes",
			Message: "max_output_bytes must be positive",
		})
	}
	if c.MenuColumns < 1 {
		errs = append(errs, ValidationError{
			Path:    "menu_columns",
			Message: "menu_columns must be >= 1",
		})
	}
	if c.PageSize < 1 {
		errs = append(errs, ValidationError{
			Path:    "page_size",
			Message: "page_size must be >= 1",
		})
	}
	if len(c.Menu) == 0 {
		errs = append(errs, ValidationError{
			Path:    "menu",
			Message: "at least one button or category is required",
		})
	}

	errs = append(errs, validateButtons(c.Menu, "menu")...)
	errs = append(errs, c.validateFunctionDirectory()...)
	errs = append(errs, c.validateLogging()...)

	return errs
}

func (c *Config) validateFunctionDirectory() ValidationErrors {
	var errs ValidationErrors
	if !c.functionDirectorySet || strings.TrimSpace(c.FunctionDirectory) == "" {
		// Missing or empty: info-level only, no hard error.
		return nil
	}
	info, err := os.Stat(c.FunctionDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			errs = append(errs, ValidationError{
				Path:    "function_directory",
				Message: fmt.Sprintf("directory %q does not exist", c.FunctionDirectory),
			})
			return errs
		}
		errs = append(errs, ValidationError{
			Path:    "function_directory",
			Message: fmt.Sprintf("cannot access %q: %v", c.FunctionDirectory, err),
		})
		return errs
	}
	if !info.IsDir() {
		errs = append(errs, ValidationError{
			Path:    "function_directory",
			Message: fmt.Sprintf("%q is not a directory", c.FunctionDirectory),
		})
	}
	return errs
}

func (c *Config) validateLogging() ValidationErrors {
	var errs ValidationErrors
	for name, lg := range c.Logging.Logs {
		prefix := fmt.Sprintf("logging.logs.%s", name)
		level := strings.ToLower(lg.Level)
		switch level {
		case "", "debug", "info", "warn", "warning", "error":
		default:
			errs = append(errs, ValidationError{
				Path:    prefix + ".level",
				Message: fmt.Sprintf("unsupported level %q", lg.Level),
			})
		}
		format := strings.ToLower(lg.Format)
		switch format {
		case "", "json", "console":
		default:
			errs = append(errs, ValidationError{
				Path:    prefix + ".format",
				Message: fmt.Sprintf("unsupported format %q", lg.Format),
			})
		}
		for i, out := range lg.Output {
			op := fmt.Sprintf("%s.output[%d]", prefix, i)
			switch strings.ToLower(out.Output) {
			case "stdout", "stderr", "discard", "":
			case "file":
				if strings.TrimSpace(out.File) == "" {
					errs = append(errs, ValidationError{
						Path:    op + ".file",
						Message: "file path is required when output is file",
					})
				}
			default:
				errs = append(errs, ValidationError{
					Path:    op + ".output",
					Message: fmt.Sprintf("unsupported output %q", out.Output),
				})
			}
		}
	}
	return errs
}

func validateButtons(nodes []ButtonNode, path string) ValidationErrors {
	var errs ValidationErrors
	seen := map[string]struct{}{}
	for i, n := range nodes {
		p := fmt.Sprintf("%s[%d]", path, i)
		nameKey := strings.ToLower(n.Name)
		if strings.TrimSpace(n.Name) == "" {
			errs = append(errs, ValidationError{
				Path:    p + ".name",
				Message: "name is required",
			})
		} else if _, ok := seen[nameKey]; ok {
			errs = append(errs, ValidationError{
				Path:    p + ".name",
				Message: fmt.Sprintf("duplicate name %q in the same level", n.Name),
			})
		} else {
			seen[nameKey] = struct{}{}
		}

		switch strings.ToLower(n.Type) {
		case "category":
			if len(n.Params) > 0 {
				keys := make([]string, 0, len(n.Params))
				for key := range n.Params {
					keys = append(keys, key)
				}
				sort.Strings(keys)
				errs = append(errs, ValidationError{
					Path:    p,
					Message: fmt.Sprintf("category has unsupported parameter keys: %s", strings.Join(keys, ", ")),
				})
			}
			if len(n.Items) == 0 {
				errs = append(errs, ValidationError{
					Path:    p + ".items",
					Message: "category must have at least one item",
				})
			}
			errs = append(errs, validateButtons(n.Items, p+".items")...)
		case "button":
			if strings.TrimSpace(n.Function) == "" {
				errs = append(errs, ValidationError{
					Path:    p + ".function",
					Message: "function is required for buttons",
				})
			}
			if len(n.Items) > 0 {
				errs = append(errs, ValidationError{
					Path:    p + ".items",
					Message: "button must not have items",
				})
			}
			fn := strings.ToLower(n.Function)
			if fn == "command" && strings.TrimSpace(n.Command) == "" {
				errs = append(errs, ValidationError{
					Path:    p + ".command",
					Message: "command is required when function is command",
				})
			}
			if fn == "script" && strings.TrimSpace(n.Path) == "" {
				errs = append(errs, ValidationError{
					Path:    p + ".path",
					Message: "path is required when function is script",
				})
			}
		case "":
			errs = append(errs, ValidationError{
				Path:    p + ".type",
				Message: "type is required (category or button)",
			})
		default:
			errs = append(errs, ValidationError{
				Path:    p + ".type",
				Message: fmt.Sprintf("unsupported type %q (expected category or button)", n.Type),
			})
		}

		if n.Columns != nil && *n.Columns < 1 {
			errs = append(errs, ValidationError{
				Path:    p + ".columns",
				Message: "columns must be >= 1",
			})
		}
		if n.Timeout != nil && n.Timeout.Duration <= 0 {
			errs = append(errs, ValidationError{
				Path:    p + ".timeout",
				Message: "timeout must be positive",
			})
		}
	}
	return errs
}

// AbsFunctionDirectory returns an absolute path for function_directory if set.
func (c *Config) AbsFunctionDirectory() (string, bool) {
	if !c.functionDirectorySet || strings.TrimSpace(c.FunctionDirectory) == "" {
		return "", false
	}
	if filepath.IsAbs(c.FunctionDirectory) {
		return c.FunctionDirectory, true
	}
	abs, err := filepath.Abs(c.FunctionDirectory)
	if err != nil {
		return c.FunctionDirectory, true
	}
	return abs, true
}
