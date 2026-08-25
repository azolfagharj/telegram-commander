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

### 1. Download and extract

```bash
wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
tar -xzf telegram-commander.tar.gz
```

### 2. Pick your binary

**Linux amd64:**

```bash
cp telegram-commander-linux-amd64 telegram-commander && chmod +x telegram-commander
```

**Linux arm64:**

```bash
cp telegram-commander-linux-arm64 telegram-commander && chmod +x telegram-commander
```

### 3. Set bot token and user id

Edit `config.minimal.yaml` and replace `YOUR_BOT_TOKEN` and `YOUR_USER_ID`.

If you do not know your user id, run once with only the token set; unauthorized users get a message that includes their `user_id`.

### 4. Run

```bash
./telegram-commander validate --config config.minimal.yaml
./telegram-commander run --config config.minimal.yaml
```

## systemd (as root)

No extra Linux user is created. The unit runs as root.

```bash
wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
tar -xzf telegram-commander.tar.gz
sudo mkdir -p /etc/telegram-commander
# amd64:
sudo cp telegram-commander-linux-amd64 /etc/telegram-commander/telegram-commander
# arm64: sudo cp telegram-commander-linux-arm64 /etc/telegram-commander/telegram-commander
sudo chmod +x /etc/telegram-commander/telegram-commander
sudo cp config.minimal.yaml /etc/telegram-commander/config.yaml
# edit /etc/telegram-commander/config.yaml (token + user id)
sudo cp telegram-commander.service /etc/systemd/system/
sudo systemctl daemon-reload && sudo systemctl enable --now telegram-commander
```

Pick the binary line that matches your machine architecture.

## Documentation

Site (GitHub Pages): https://azolfagharj.github.io/telegram-commander/

- [Config reference](docs/config-reference.md)
- [Functions reference](docs/functions-reference.md)
- [CLI reference](docs/cli-reference.md)
- [Install](docs/install.md)

## Example configs

- Minimal: [`examples/config.minimal.yaml`](examples/config.minimal.yaml)
- Full: [`examples/config.full.yaml`](examples/config.full.yaml)

## License

See repository license file if present.
