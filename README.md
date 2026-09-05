# Telegram Commander

[![CI](https://github.com/azolfagharj/telegram-commander/actions/workflows/ci.yml/badge.svg)](https://github.com/azolfagharj/telegram-commander/actions) [![Documentation](https://img.shields.io/badge/Documentation-Site-blue?logo=github)](https://telecommander.ir/) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE) [![Donate](https://img.shields.io/badge/Donate-to%20Keep%20This%20Project%20Alive-orange)](https://azolfagharj.github.io/donate/)

**Control your Linux server from Telegram, one tap at a time.**

<p align="center">
  <a href="docs/markdown/images/01.jpeg"><img src="docs/markdown/images/01.jpeg" alt="Command output with the System menu open" width="150"></a>
  <a href="docs/markdown/images/02.jpeg"><img src="docs/markdown/images/02.jpeg" alt="Resource and process buttons" width="150"></a>
  <a href="docs/markdown/images/03.jpeg"><img src="docs/markdown/images/03.jpeg" alt="Storage and package buttons" width="150"></a>
  <a href="docs/markdown/images/04.jpeg"><img src="docs/markdown/images/04.jpeg" alt="Network tools and typing a command by hand" width="150"></a>
  <a href="docs/markdown/images/05.jpeg"><img src="docs/markdown/images/05.jpeg" alt="Command output sent back to the chat" width="150"></a>
</p>

Telegram Commander turns a simple YAML file into a Telegram bot with a menu of
buttons. Each button can run *any* command on the server — restart a service,
check disk space, tail logs, deploy, run a backup script, ping a host — and the
output comes straight back to your chat.

If you can type it in a terminal, you can put it behind a button. No app to
build, no web panel to manage, no inbound port to open: the bot talks to
Telegram over an outbound connection and only the users you list can use it.

The routine work you keep opening an SSH session for — checking a service,
reading a log, seeing how much disk is left — becomes a tap on your phone. SSH
is still there for everything else.

You describe the menu in YAML:

```yaml
menu:
  - name: System
    type: category
    icon: "🖥️"
    items:
      - name: Uptime & Load
        type: button
        icon: "🕒"
        function: command
        command: "uptime && uptime -p"
```

That button shows up in your chat. Tap it, the command runs on the server, and
the output comes back as a message.

## Full documentation

<p align="center">
  <a href="https://telecommander.ir/"><img src="docs/markdown/images/visit-website.svg" alt="Visit the website" width="300" height="56"></a>
</p>

The complete guide is written in plain language and is available in several
languages. Open the one you read:

<a href="https://telecommander.ir/"><img src="docs/markdown/images/flags/gb.svg" alt="" width="20" height="15"> English</a>
·
<a href="https://telecommander.ir/de/"><img src="docs/markdown/images/flags/de.svg" alt="" width="20" height="15"> Deutsch</a>
·
<a href="https://telecommander.ir/fr/"><img src="docs/markdown/images/flags/fr.svg" alt="" width="20" height="15"> Français</a>
·
<a href="https://telecommander.ir/es/"><img src="docs/markdown/images/flags/es.svg" alt="" width="20" height="15"> Español</a>
·
<a href="https://telecommander.ir/ru/"><img src="docs/markdown/images/flags/ru.svg" alt="" width="20" height="15"> Русский</a>
·
<a href="https://telecommander.ir/zh/"><img src="docs/markdown/images/flags/cn.svg" alt="" width="20" height="15"> 简体中文</a>
·
<a href="https://telecommander.ir/fa/"><img src="docs/markdown/images/flags/ir.svg" alt="" width="20" height="15"> فارسی</a>

## What can I do with it?

Anything you can run in a shell. A few common examples:

- Restart or stop a service
- Start and stop containers
- Update system packages
- Read logs and journals
- Check disk space
- Watch CPU and memory
- Ping hosts and test URLs
- Take and restore backups
- Run your own scripts
- Reboot or shut down the host
- Type any command by hand
- And almost anything else

## Why use it?

- **No coding.** Describe the menu and commands in one YAML file.
- **From anywhere.** Open Telegram on your phone and manage the server. No VPN
  into the host.
- **No open ports.** The bot connects out to Telegram. Nothing is exposed to
  the internet.
- **Output in the chat.** The result comes back as a message. You do not need
  an SSH session.
- **Controlled and recorded.** Choose who gets the menu, confirm risky actions,
  and record each run.
- **Nested menus.** Group buttons into categories. Home stays at the top; Back
  takes you up.
- **Reusable functions.** Write a command once, then fill in different values
  on each button.
- **Stays running.** Install it as a service and the bot starts with the host.

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

If you do not know your user id, leave any number in `allowed_users` for now
and start the bot, then send it a message. Because you are not on the list yet,
the bot replies with your own `user_id`. Copy that id into `allowed_users`,
stop the bot, and start it again.

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

The complete, beginner-friendly documentation lives here:

**<https://telecommander.ir/>**

Handy entry points:

- [Run in CLI](https://telecommander.ir/documentation/installation/download-and-run/) — from download to first run
- [Concepts](https://telecommander.ir/documentation/concepts/config-file/) — the vocabulary
- [Configuration](https://telecommander.ir/documentation/configuration/) — every setting
- [Menu](https://telecommander.ir/documentation/concepts/menu/) — build your menu
- [Functions](https://telecommander.ir/documentation/functions/) — built-in and custom functions
- [CLI](https://telecommander.ir/documentation/cli/) — command line reference
- [Run as a service](https://telecommander.ir/documentation/installation/run-as-a-service/) — systemd setup

The same pages live as Markdown in [`docs/markdown/`](docs/markdown/).

## Example configs

In this repository:

- Minimal: [`examples/config-examples/config.minimal.yaml`](examples/config-examples/config.minimal.yaml)
- Full: [`examples/config-examples/config.full.yaml`](examples/config-examples/config.full.yaml)
- Sample functions: [`examples/functions/`](examples/functions/)

In the release archive the same files sit next to the binaries, without the
`examples/` prefix: after extracting you have `config-examples/` and
`functions/` inside the `telegram-commander/` folder, which is why Quick start
above uses `config-examples/config.minimal.yaml`.

## License

Released under the [MIT License](LICENSE).

---

## Support this Project

🤝 **Enjoying this free project?** <a href="https://azolfagharj.github.io/donate/">Consider supporting</a> its development

<a href="https://azolfagharj.github.io/donate/"><img src="https://img.shields.io/badge/Donate-Support%20Development-orange?style=for-the-badge" alt="Donate"></a>
