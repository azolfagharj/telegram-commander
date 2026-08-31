# Run as a service

This page covers running Telegram Commander as a systemd service and building
from source. For a step-by-step first run in the terminal, see
[Run in CLI](download-and-run.md).

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

The service runs the bot with [long polling](../concepts/long-polling.md). There
is no hot reload — restart the service after you change your
[config file](../concepts/config-file.md).

## Build from source

```bash
git clone https://github.com/azolfagharj/telegram-commander.git
cd telegram-commander
go build -o telegram-commander ./cmd/telegram-commander
```

The minimal config does not set `function_directory`. To try sample functions,
use `config-examples/config.full.yaml` (it points to `./functions`) or set
`function_directory` yourself. Each sample function is described in
[Functions → Custom functions in the release pack](../functions.md#custom-functions-in-the-release-pack).

## Related pages

- [Run in CLI](download-and-run.md) — explained first run
- [Configuration](../configuration.md) — the config file
- [CLI](../cli.md) — `run`, `validate`, and more
