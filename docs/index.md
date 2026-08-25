# Telegram Commander

Telegram Commander is a small Go program that turns a YAML file into a
Telegram bot. The bot shows a menu of buttons. When you tap a button, the
program runs a command on the server where the bot is running and sends the
output back to you in Telegram.

You do not write any code to use it. You describe your menu and your commands
in a configuration file, then start the bot.

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

## What you can do with it

- Build a menu of server actions (restart a service, show disk usage, tail logs)
- Group actions into nested categories
- Ask for confirmation before dangerous actions
- Restrict who can use the bot
- Run through a proxy if Telegram is blocked on your network

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
