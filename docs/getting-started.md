# Getting started

This page takes you from nothing to a running bot. No prior experience with the
project is needed. If a word is unclear, check the [Concepts](concepts.md) page.

## Before you begin

You need two things from Telegram:

1. **A bot token.** Open a chat with [@BotFather](https://t.me/BotFather),
   send `/newbot`, follow the prompts, and copy the token it gives you. It looks
   like `123456789:AAExampleTokenValue`.
2. **Your numeric user id.** This is a number, not your `@username`. If you do
   not know it, do not worry — the bot will tell you your id the first time you
   message it (see [step 5](#step-5-find-your-user-id-if-needed)).

## Step 1: download

Download the release archive and extract it.

```bash
wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
tar -xzf telegram-commander.tar.gz
cd telegram-commander
```

Inside the folder you will find:

- `telegram-commander-linux-amd64` and `telegram-commander-linux-arm64` — the program, one per CPU type
- `config-examples/` — ready-made config files (see [Configuration](config-reference.md))
- `functions/` — example custom functions (see [Functions](functions-reference.md))

## Step 2: pick your binary

Most servers and PCs are `amd64` (also called `x86_64`). Small ARM boards and
some cloud VMs are `arm64`. If unsure, run `uname -m`: `x86_64` means amd64,
`aarch64` means arm64.

```bash
# amd64
mv telegram-commander-linux-amd64 telegram-commander
# arm64: mv telegram-commander-linux-arm64 telegram-commander
chmod +x telegram-commander
```

Now you have a single program named `telegram-commander`.

## Step 3: create your config

Copy the minimal example to a working file:

```bash
cp config-examples/config.minimal.yaml ./config.yaml
```

Open `config.yaml` and replace two placeholders:

- `YOUR_BOT_TOKEN` — the token from BotFather
- `YOUR_USER_ID` — your numeric id (or leave it for now and see step 5)

To learn what every setting means, read [Configuration](config-reference.md).

## Step 4: validate

Always check the config before running. This catches typos and mistakes without
starting the bot.

```bash
./telegram-commander validate --config config.yaml
```

If it prints `Valid configuration`, you are good. If not, it lists exactly what
is wrong and where. See the [CLI page](cli-reference.md#validate) for details.

## Step 5: find your user id (if needed)

If you did not know your user id, set only the token in `config.yaml`, put any
number in `allowed_users` for now, then run the bot:

```bash
./telegram-commander run --config config.yaml
```

Open Telegram, find your bot, and send it any message. Because you are not in
`allowed_users` yet, the bot replies with your `user_id` and `username`. Copy
that id into `allowed_users`, stop the bot with `Ctrl+C`, and run it again.

This behavior is part of how access control works; see
[Configuration → telegram](config-reference.md#telegram).

## Step 6: run

```bash
./telegram-commander run --config config.yaml
```

Open your bot in Telegram and send `/start`. You should see your menu. Tap a
button to run its command.

To keep the bot running after you log out of the server, set it up as a service.
See [Install and service](install.md).

## What next

- Add more buttons and categories: [Buttons and menus](buttons.md)
- Understand what actually runs: [Functions](functions-reference.md)
- See every command line option: [CLI](cli-reference.md)
