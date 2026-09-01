// Package cli implements the Cobra-based command line interface.
package cli

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/azolfagharj/telegram-commander/internal/bot"
	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/executor"
	"github.com/azolfagharj/telegram-commander/internal/function"
	"github.com/azolfagharj/telegram-commander/internal/logging"
)

// NewRoot builds the root cobra command with all subcommands.
// version and releaseDate come from cmd/telegram-commander/main.go only.
func NewRoot(version, releaseDate string) *cobra.Command {
	root := &cobra.Command{
		Use:           "telegram-commander",
		Short:         "Telegram bot that runs server commands from a YAML menu",
		Long:          "telegram-commander reads a YAML config, shows an inline button menu in Telegram, and executes configured functions on the host.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(
		newRunCmd(),
		newValidateCmd(),
		newVersionCmd(version, releaseDate),
		newFmtCmd(),
		newEnvironCmd(),
		newListFunctionsCmd(),
		newCompletionCmd(),
		newManpageCmd(),
	)
	return root
}

func requireConfigFlag(cmd *cobra.Command) (string, error) {
	path, err := cmd.Flags().GetString("config")
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(path) == "" {
		return "", fmt.Errorf("--config is required (relative or absolute path)")
	}
	return path, nil
}

func addConfigFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("config", "c", "", "path to config YAML file (required)")
	_ = cmd.MarkFlagRequired("config")
}

func loadAndValidate(path string) (*config.Config, *function.Registry, error) {
	cfg, err := config.Load(path)
	if err != nil {
		return nil, nil, err
	}
	errs := cfg.Validate()

	reg := function.NewRegistry()
	if dir, ok := cfg.AbsFunctionDirectory(); ok {
		if err := reg.LoadDirectory(dir); err != nil {
			errs = append(errs, config.ValidationError{
				Path:    "function_directory",
				Message: err.Error(),
			})
		}
	} else if !cfg.FunctionDirectoryWasSet() || strings.TrimSpace(cfg.FunctionDirectory) == "" {
		fmt.Fprintln(os.Stderr, "info: function_directory not set or empty; using built-in functions only")
	}
	for _, d := range reg.List() {
		if err := d.ValidateTemplate(); err != nil {
			errs = append(errs, config.ValidationError{
				Path:    "functions." + d.Name,
				Message: err.Error(),
			})
		}
	}
	errs = append(errs, reg.ValidateTree(cfg.Menu)...)
	if err := errs.Err(); err != nil {
		return nil, nil, err
	}
	return cfg, reg, nil
}

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the bot in the foreground",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := requireConfigFlag(cmd)
			if err != nil {
				return err
			}
			cfg, reg, err := loadAndValidate(path)
			if err != nil {
				return err
			}
			audit, err := logging.Setup(cfg.Logging)
			if err != nil {
				return fmt.Errorf("logging: %w", err)
			}
			exec := executor.NewShellExecutor(audit)
			app := bot.NewApp(cfg, reg, exec, nil)
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()
			return app.Start(ctx)
		},
	}
	addConfigFlag(cmd)
	return cmd
}

func newValidateCmd() *cobra.Command {
	var online bool
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate the configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := requireConfigFlag(cmd)
			if err != nil {
				return err
			}
			cfg, _, err := loadAndValidate(path)
			if err != nil {
				return err
			}
			if online {
				client, err := bot.HTTPClient(cfg.Telegram)
				if err != nil {
					return err
				}
				api := strings.TrimRight(cfg.Telegram.API, "/")
				url := fmt.Sprintf("%s/bot%s/getMe", api, cfg.Telegram.BotToken)
				resp, err := client.Get(url)
				if err != nil {
					return fmt.Errorf("online check failed: %w", err)
				}
				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)
				if resp.StatusCode != 200 {
					return fmt.Errorf("online check failed: HTTP %d: %s", resp.StatusCode, string(body))
				}
				if !strings.Contains(string(body), `"ok":true`) {
					return fmt.Errorf("online check failed: %s", string(body))
				}
				fmt.Fprintln(cmd.OutOrStdout(), "Valid configuration (online check passed)")
				return nil
			}
			fmt.Fprintln(cmd.OutOrStdout(), "Valid configuration")
			return nil
		},
	}
	addConfigFlag(cmd)
	cmd.Flags().BoolVar(&online, "online", false, "also verify bot_token with Telegram API")
	return cmd
}

func newVersionCmd(version, releaseDate string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "telegram-commander %s (%s)\nGo %s %s/%s\n",
				version, releaseDate, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		},
	}
}

func newFmtCmd() *cobra.Command {
	var write bool
	cmd := &cobra.Command{
		Use:   "fmt",
		Short: "Format a config YAML file",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := requireConfigFlag(cmd)
			if err != nil {
				return err
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			out, err := config.FormatYAML(data)
			if err != nil {
				return err
			}
			if write {
				return os.WriteFile(path, out, 0o644)
			}
			_, err = cmd.OutOrStdout().Write(out)
			return err
		},
	}
	addConfigFlag(cmd)
	cmd.Flags().BoolVarP(&write, "write", "w", false, "write result back to the file")
	return cmd
}

func newEnvironCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "environ",
		Short: "Print environment variables",
		Run: func(cmd *cobra.Command, args []string) {
			for _, e := range os.Environ() {
				fmt.Fprintln(cmd.OutOrStdout(), e)
			}
		},
	}
}

func newListFunctionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-functions",
		Short: "List built-in and configured functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := requireConfigFlag(cmd)
			if err != nil {
				return err
			}
			cfg, err := config.Load(path)
			if err != nil {
				return err
			}
			reg := function.NewRegistry()
			if dir, ok := cfg.AbsFunctionDirectory(); ok {
				if err := reg.LoadDirectory(dir); err != nil {
					return err
				}
			}
			for _, d := range reg.List() {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\tsource=%s\tparams=%d\n", d.Name, d.Source, len(d.Params))
			}
			return nil
		},
	}
	addConfigFlag(cmd)
	return cmd
}

func newCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate shell completion scripts",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return cmd.Root().GenBashCompletion(cmd.OutOrStdout())
			case "zsh":
				return cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			case "fish":
				return cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
			case "powershell":
				return cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
			default:
				return fmt.Errorf("unsupported shell %q", args[0])
			}
		},
	}
}

func newManpageCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "manpage",
		Short: "Generate a man page on stdout",
		RunE: func(cmd *cobra.Command, args []string) error {
			header := &doc.GenManHeader{
				Title:   "TELEGRAM-COMMANDER",
				Section: "1",
			}
			return doc.GenMan(cmd.Root(), header, cmd.OutOrStdout())
		},
	}
}
