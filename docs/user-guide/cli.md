# CLI reference

```
telegram-commander <command> [flags]
```

`--config` / `-c` is **required** for commands that load configuration. The path
may be relative to the current working directory or absolute. What goes in that
file is documented in [Configuration](configuration.md).

If you are just getting started, [Run in CLI](installation/download-and-run.md)
shows these commands in the order you will use them.

## Commands

### `run`

Run the bot in the foreground.

```bash
telegram-commander run --config /path/to/configfile.yaml
```

To apply config changes, restart the process (for example
`systemctl restart telegram-commander`).

### `validate`

Validate config, functions, and button references offline.

```bash
telegram-commander validate --config /path/to/configfile.yaml
telegram-commander validate --config /path/to/configfile.yaml --online
```

`--online` also checks the bot token with Telegram.

### `version`

Print the program version.

### `fmt`

Pretty-print a YAML config file.

```bash
telegram-commander fmt --config /path/to/configfile.yaml
telegram-commander fmt --config /path/to/configfile.yaml -w
```

### `environ`

Print process environment variables (useful for debugging service units).

### `list-functions`

List built-in and loaded custom functions. Use it to confirm your custom
function files were found. See [Functions](functions.md).

```bash
telegram-commander list-functions --config /path/to/configfile.yaml
```

### `completion`

Generate shell completion scripts:

```bash
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
