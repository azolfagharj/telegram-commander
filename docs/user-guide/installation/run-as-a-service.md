# Run as a service

This page covers running Telegram Commander as a systemd service. For a
step-by-step first run in the terminal, see [Run in CLI](download-and-run.md).

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

The service keeps the bot running. Restart the service after you change your
[config file](../concepts/config-file.md).

## Related pages

- [Run in CLI](download-and-run.md) — explained first run
- [Configuration](../configuration.md) — the config file
- [CLI](../cli.md) — `run`, `validate`, and more
