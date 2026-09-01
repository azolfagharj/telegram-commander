---
title: CLI reference
description: Every Telegram Commander command line option, from running the bot to checking and formatting your config file, and the --config flag they use.
icon: material/console-line
---

# :material-console-line: CLI reference

!!! info "Commands use this form"

    ```text title="Command syntax"
    telegram-commander <command> [flags]
    ```

`--config` / `-c` is **required** for commands that load configuration. The path
may be relative to the current working directory or absolute. What goes in that
file is documented in [Configuration](configuration.md).

If you are just getting started, [Run in CLI](installation/download-and-run.md)
shows these commands in the order you will use them.

## :material-format-list-bulleted-square: Commands

### `run`

Run the bot in the foreground.

!!! info "Start the bot in the foreground"

    ```bash title="Run the bot"
    telegram-commander run --config /path/to/configfile.yaml
    ```

!!! note "Config changes need a restart"

    The bot reads the config once at start. After you edit the file, restart
    the process (for example `systemctl restart telegram-commander`).

### `validate`

Validate config, functions, and button references offline.

!!! info "Check the config offline or online"

    ```bash title="Validate the config"
    telegram-commander validate --config /path/to/configfile.yaml
    telegram-commander validate --config /path/to/configfile.yaml --online
    ```

!!! note "`--online` needs internet"

    With `--online` the check also asks Telegram whether the bot token works,
    so the machine must be able to reach Telegram.

### `version`

Print the program version.

### `fmt`

Pretty-print a YAML config file.

!!! info "Print or save formatted YAML"

    ```bash title="Format a config file"
    telegram-commander fmt --config /path/to/configfile.yaml
    telegram-commander fmt --config /path/to/configfile.yaml -w
    ```

### `environ`

Print process environment variables (useful for debugging service units).

### `list-functions`

List built-in and loaded custom functions. Use it to confirm your custom
function files were found. See [Functions](functions.md).

!!! info "Show every available function"

    ```bash title="List available functions"
    telegram-commander list-functions --config /path/to/configfile.yaml
    ```

### `completion`

Generate shell completion scripts:

!!! info "Choose the shell you use"

    ```bash title="Generate a completion script"
    telegram-commander completion bash
    telegram-commander completion zsh
    telegram-commander completion fish
    telegram-commander completion powershell
    ```

### `manpage`

Write a man page to stdout.

## Related pages

- [Configuration](configuration.md) — the file passed to `--config`
- [Functions](functions.md) — what `list-functions` shows
- [Run as a service](installation/run-as-a-service.md) — run `run` under systemd
