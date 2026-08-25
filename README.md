# telegram-commander

A Telegram bot written in Go that reads a YAML config, shows an inline button menu,
and runs host commands (or custom functions) on the server where it runs.

CLI style is inspired by [Caddy](https://caddyserver.com/): `run`, `validate`, `fmt`, `version`, and more.

## Features

- YAML config with strict unknown-field rejection
- Nested category / button menu as Telegram inline keyboards
- Built-in functions: `command`, `script`
- Custom functions as one YAML file each under `function_directory`
- Confirmation prompts, pagination, per-user command queue
- Optional SOCKS5/HTTP proxy for Telegram API
- Structured logging + separate audit logger
- Offline `validate` (optional `--online` token check)

## Quick start

```bash
go build -o telegram-commander ./cmd/telegram-commander

# edit examples/config.yaml and set bot_token + allowed_users
./telegram-commander validate --config examples/config.yaml
./telegram-commander run --config examples/config.yaml
```

`--config` is required and accepts a relative or absolute path.

## Documentation

- [Config reference](docs/config-reference.md)
- [Functions reference](docs/functions-reference.md)
- [CLI reference](docs/cli-reference.md)
- [Install](docs/install.md)

## Example config (excerpt)

```yaml
telegram:
  bot_token: "YOUR_BOT_TOKEN"
  allowed_users:
    - "123456789"
    - "admin_username"

buttons:
  - name: Echo demo
    type: button
    function: command
    command: "echo hello"
```

## License

See repository license file if present.
