---
title: Run as a service
description: Keep the bot running in the background with systemd, so it starts at boot and comes back after a failure, as root or as a normal user.
icon: material/server
---

# :material-server: Run as a service

This page covers running Telegram Commander as a systemd service. For a
step-by-step first run in the terminal, see [Run in CLI](download-and-run.md).

## :material-cog-outline: systemd (root)

Create the unit file yourself, for example `/etc/systemd/system/telegram-commander.service`:

!!! warning "This example runs every button as root"

    ```ini title="/etc/systemd/system/telegram-commander.service"
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

    With no `User=` line, the service starts as root, so every command in your
    menu has full rights on the machine. Add `User=someone` to the `[Service]`
    section if your buttons do not need that much power.

Replace the placeholder paths, then:

!!! example "Load and start the service"

    ```bash title="Enable, start, and watch the service"
    sudo systemctl daemon-reload
    sudo systemctl enable --now telegram-commander
    sudo systemctl status telegram-commander
    sudo journalctl -u telegram-commander -f
    ```

!!! info "Changes to the config need a restart"

    The service keeps the bot running, but it reads your
    [config file](../concepts/config-file.md) once at start. Run
    `sudo systemctl restart telegram-commander` after you edit it.

## Related pages

- [Run in CLI](download-and-run.md) — explained first run
- [Configuration](../configuration.md) — the config file
- [CLI](../cli.md) — `run`, `validate`, and more
