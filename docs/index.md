# Telegram Commander

**Control your Linux server from Telegram, one tap at a time.**

Telegram Commander turns a simple YAML [config file](user-guide/concepts/config-file.md) into a Telegram bot with a menu of
[buttons](user-guide/concepts/button.md). When you tap a button, the program runs a command on the server where
the bot is running and sends the output straight back to your chat.

If you can type it in a terminal, you can put it behind a button — restart a
service, check disk space, tail logs, deploy, run a backup script, ping a host.
You do not write any code: you describe your menu and your commands in a
configuration file, then start the bot.

There is no web panel to secure and no inbound port to open. The bot connects
out to Telegram through [long polling](user-guide/concepts/long-polling.md), and only the
[allowed users](user-guide/concepts/allowed-users.md) you list can use it.

## A tiny example

This config makes a bot with one button called "Uptime". Tapping it runs the
`uptime` command on the server.

```yaml
telegram:
  bot_token: "YOUR_BOT_TOKEN"
  allowed_users:
    - "YOUR_USER_ID"

buttons:
  - name: Uptime
    type: button
    function: command
    command: "uptime"
```

That is a complete, working configuration. Everything else is optional.

## What can I do with it?

Anything you can run in a shell, for example:

- **Manage services** — `systemctl restart nginx`, start/stop/status of anything
- **Check the server** — disk usage, memory, uptime, running processes
- **Read logs** — tail files or pull recent `journalctl` output for a unit
- **Run your own scripts** — deploys, backups, cleanups, health checks
- **Reach out** — curl a health endpoint, ping a host

You define these as [buttons](user-guide/buttons.md), group them into [categories](user-guide/concepts/category.md), ask for
[confirmation](user-guide/concepts/confirmation.md) before dangerous actions,
[restrict who can use the bot](user-guide/configuration.md#telegram), and run through a
[proxy](user-guide/configuration.md#telegram) if Telegram is blocked on your network.

## Where to go next

New here? Follow the pages in order:

1. [Run in CLI](user-guide/installation/download-and-run.md) — install, set your token, run the bot
2. [Concepts](user-guide/concepts/config-file.md) — the words used in this documentation
3. [Configuration](user-guide/configuration.md) — every setting, with defaults and examples
4. [Buttons and menus](user-guide/buttons.md) — how to build your menu tree
5. [Functions](user-guide/functions.md) — what runs when a button is tapped
6. [CLI](user-guide/cli.md) — the command line commands
7. [Run as a service](user-guide/installation/run-as-a-service.md) — run it as a systemd service

## Links

- [GitHub repository](https://github.com/azolfagharj/telegram-commander)
- [Latest release](https://github.com/azolfagharj/telegram-commander/releases/latest)

## Support this project

Telegram Commander is free and open source. If it saves you time,
[consider supporting its development](https://azolfagharj.github.io/donate/) —
it helps keep the project alive and maintained.
