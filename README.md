# Telegram Commander

[![CI](https://github.com/azolfagharj/telegram-commander/actions/workflows/ci.yml/badge.svg)](https://github.com/azolfagharj/telegram-commander/actions) [![Documentation](https://img.shields.io/badge/Documentation-Site-blue?logo=github)](https://azolfagharj.github.io/telegram-commander/) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE) [![Donate](https://img.shields.io/badge/Donate-to%20Keep%20This%20Project%20Alive-orange)](https://azolfagharj.github.io/donate/)

**Control your Linux server from Telegram, one tap at a time.**

Telegram Commander turns a simple YAML file into a Telegram bot with a menu of
buttons. Each button can run *any* command on the server — restart a service,
check disk space, tail logs, deploy, run a backup script, ping a host — and the
output comes straight back to your chat.

If you can type it in a terminal, you can put it behind a button. No app to
build, no web panel to secure, no inbound port to open: the bot talks to
Telegram over an outbound connection and only the users you list can use it.

> **Full documentation:** <https://azolfagharj.github.io/telegram-commander/>
>
> New here? Start with [Run in CLI](https://azolfagharj.github.io/telegram-commander/user-guide/installation/download-and-run/).

## What can I do with it?

Anything you can run in a shell. A few common examples:

- **Manage services** — `systemctl restart nginx`, start/stop/status of anything
- **Check the server** — disk usage, memory, uptime, running processes
- **Read logs** — tail files or pull recent `journalctl` output for a unit
- **Run your own scripts** — deploys, backups, cleanups, health checks
- **Reach out** — curl a health endpoint, ping a host

You define these as **buttons**, group them into **categories**, and optionally
require a confirmation tap before the risky ones run. Values you reuse can live
in small **custom functions** so buttons stay short and consistent.

## Why use it?

- **Zero coding.** Describe the menu and commands in YAML; no plugins to write.
- **Works from anywhere.** Manage the box from your phone, wherever you are.
- **Safe by default.** Only listed users get in; unknown users are told to
  contact the admin. Add `confirm: true` to guard destructive actions.
- **No open ports.** The bot connects out to Telegram; nothing to expose to the internet.
- **Auditable.** A separate audit log records who ran what, the exit code, and duration.

## Features

- YAML config with strict unknown-field rejection (typos become errors)
- Nested category / button menu in Telegram
- Run any shell command via the built-in `command` and `script` functions
- Reusable custom functions, one YAML file each under `function_directory`
- Confirmation prompts, pagination, and a per-user command queue (serial per user)
- Access control by numeric user id or `@username`
- Optional SOCKS5/HTTP proxy for the Telegram API
- Structured logging plus a separate audit logger
- Offline `validate` (with optional `--online` token check)
- CLI: `run`, `validate`, `fmt`, `list-functions`, `version`, and more

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

The complete, beginner-friendly documentation lives on GitHub Pages:

**<https://azolfagharj.github.io/telegram-commander/>**

Handy entry points:

- [Run in CLI](https://azolfagharj.github.io/telegram-commander/user-guide/installation/download-and-run/) — from download to first run
- [Concepts](https://azolfagharj.github.io/telegram-commander/user-guide/concepts/config-file/) — the vocabulary
- [Configuration](https://azolfagharj.github.io/telegram-commander/user-guide/configuration/) — every setting
- [Menu](https://azolfagharj.github.io/telegram-commander/user-guide/concepts/menu/) — build your menu
- [Functions](https://azolfagharj.github.io/telegram-commander/user-guide/functions/) — built-in and custom functions
- [CLI](https://azolfagharj.github.io/telegram-commander/user-guide/cli/) — command line reference
- [Run as a service](https://azolfagharj.github.io/telegram-commander/user-guide/installation/run-as-a-service/) — systemd setup

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
