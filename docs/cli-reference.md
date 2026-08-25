# CLI reference

```
telegram-commander <command> [flags]
```

`--config` / `-c` is **required** for commands that load configuration. The path may be relative to the current working directory or absolute.

## Commands

### `run`

Run the bot in the foreground (long polling).

```bash
telegram-commander run --config /etc/telegram-commander/config.yaml
```

There is no hot reload. To apply config changes, restart the process (for example `systemctl restart telegram-commander`).

### `validate`

Validate config, functions, and button references offline.

```bash
telegram-commander validate --config ./config.yaml
telegram-commander validate --config ./config.yaml --online
```

`--online` also calls Telegram `getMe` with the configured token.

### `version`

Print version and Go runtime info.

### `fmt`

Pretty-print a YAML config file.

```bash
telegram-commander fmt --config ./config.yaml
telegram-commander fmt --config ./config.yaml -w
```

### `environ`

Print process environment variables (useful for debugging service units).

### `list-functions`

List built-in and loaded custom functions.

```bash
telegram-commander list-functions --config ./config.yaml
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
