# Install and service

This page covers downloading the release, running as a systemd service, and
building from source. For a slower, explained walkthrough of your first run, see
[Getting started](getting-started.md).

## Quick start

```bash
wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
tar -xzf telegram-commander.tar.gz
cd telegram-commander
```

Inside the extracted folder:

- `telegram-commander-linux-amd64`
- `telegram-commander-linux-arm64`
- `config-examples/` (`config.minimal.yaml`, `config.full.yaml`)
- `functions/` (sample custom functions)

```bash
# amd64
mv telegram-commander-linux-amd64 telegram-commander
# arm64: mv telegram-commander-linux-arm64 telegram-commander
chmod +x telegram-commander

cp config-examples/config.minimal.yaml ./config.yaml
# edit config.yaml (token + user id)

./telegram-commander validate --config config.yaml
./telegram-commander run --config config.yaml
```

Edit `config.yaml` following [Configuration](config-reference.md), and see
[CLI](cli-reference.md) for what `validate` and `run` do.

The minimal config does not set `function_directory`. To try sample functions,
use `config-examples/config.full.yaml` (it points to `./functions`) or set
`function_directory` yourself. Each sample function is described in
[Functions → Custom functions in the release pack](functions-reference.md#custom-functions-in-the-release-pack).

## systemd (root)

Create the unit file yourself, for example `/etc/systemd/system/telegram-commander.service`:

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

Replace the placeholder paths, then:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now telegram-commander
sudo systemctl status telegram-commander
sudo journalctl -u telegram-commander -f
```

## Build from source

```bash
git clone https://github.com/azolfagharj/telegram-commander.git
cd telegram-commander
go build -o telegram-commander ./cmd/telegram-commander
```

## Related pages

- [Getting started](getting-started.md) — explained first run
- [Configuration](config-reference.md) — the config file
- [CLI](cli-reference.md) — `run`, `validate`, and more
