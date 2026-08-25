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
cd telegram-commander
```

### 2. Rename the binary for your architecture

**Linux amd64:**

```bash
mv telegram-commander-linux-amd64 telegram-commander
chmod +x telegram-commander
```

**Linux arm64:**

```bash
mv telegram-commander-linux-arm64 telegram-commander
chmod +x telegram-commander
```

### 3. Copy minimal config and set token / user id

```bash
cp config-examples/config.minimal.yaml ./config.yaml
```

Edit `config.yaml` and replace `YOUR_BOT_TOKEN` and `YOUR_USER_ID`.

If you do not know your user id, run once with only the token set; unauthorized users get a message that includes their `user_id`.

### 4. Validate and run

```bash
./telegram-commander validate --config config.yaml
./telegram-commander run --config config.yaml
```

## systemd (as root)

After manual testing works, create a unit file yourself (for example `/etc/systemd/system/telegram-commander.service`):

```ini
[Unit]
Description=telegram-commander Telegram bot
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
ExecStart=/path/to/telegram-commander run --config /path/to/configfile.yaml
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Replace `/path/to/telegram-commander` and `/path/to/configfile.yaml` with your real paths, then:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now telegram-commander
```

## Documentation

Site (GitHub Pages): https://azolfagharj.github.io/telegram-commander/

- [Config reference](docs/config-reference.md)
- [Functions reference](docs/functions-reference.md)
- [CLI reference](docs/cli-reference.md)
- [Install](docs/install.md)

## Example configs

- Minimal: [`examples/config-examples/config.minimal.yaml`](examples/config-examples/config.minimal.yaml)
- Full: [`examples/config-examples/config.full.yaml`](examples/config-examples/config.full.yaml)
- Sample functions: [`examples/functions/`](examples/functions/)

## License

See repository license file if present.
