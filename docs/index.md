# Telegram Commander

**Control your Linux server from Telegram, one tap at a time.**

Telegram Commander turns a simple YAML file into a Telegram bot with a menu of
buttons. When you tap a button, the program runs a command on the server where
the bot is running and sends the output straight back to your chat.

If you can type it in a terminal, you can put it behind a button — restart a
service, check disk space, tail logs, deploy, run a backup script, ping a host.
You do not write any code: you describe your menu and your commands in a
configuration file, then start the bot.

There is no web panel to secure and no inbound port to open. The bot connects
out to Telegram, and only the users you list can use it.

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

You define these as [buttons](buttons.md), group them into categories, ask for
[confirmation](buttons.md#confirmation) before dangerous actions,
[restrict who can use the bot](config-reference.md#telegram), and run through a
[proxy](config-reference.md#telegram) if Telegram is blocked on your network.

## Where to go next

New here? Follow the pages in order:

1. [Getting started](getting-started.md) — install, set your token, run the bot
2. [Concepts](concepts.md) — the words used in this documentation (buttons, functions, and more)
3. [Configuration](config-reference.md) — every setting, with defaults and examples
4. [Buttons and menus](buttons.md) — how to build your menu tree
5. [Functions](functions-reference.md) — what runs when a button is tapped
6. [CLI](cli-reference.md) — the command line commands
7. [Install and service](install.md) — run it as a systemd service

## Links

- [GitHub repository](https://github.com/azolfagharj/telegram-commander)
- [Latest release](https://github.com/azolfagharj/telegram-commander/releases/latest)

## Support this project

Telegram Commander is free and open source. If it saves you time,
[consider supporting its development](https://azolfagharj.github.io/donate/) —
it helps keep the project alive and maintained.
