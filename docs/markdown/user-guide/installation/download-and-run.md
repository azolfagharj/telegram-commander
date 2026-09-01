---
title: Run in CLI
description: Get your bot running from the terminal, step by step, from downloading the release and writing a small config file to tapping the first button.
icon: material/console
---

# :material-console: Run in CLI

This page takes you from nothing to a running bot. No prior experience with the
project is needed. If a word is unclear, check the [Concepts](../concepts/config-file.md)
pages.

## :material-clipboard-check-outline: Before you begin

You need two things from Telegram:

1. **A bot token.** Open a chat with [@BotFather](https://t.me/BotFather),
   send `/newbot`, follow the prompts, and copy the token it gives you. It looks
   like `123456789:AAExampleTokenValue`.
2. **Your numeric user id.** This is a number, not your `@username`. If you do
   not know it, do not worry — the bot will tell you your id the first time you
   message it (see [step 5](#step-5-find-your-user-id-if-needed)).

## :material-download: Step 1: download

Download the release archive and extract it.

!!! example "Download and open the release folder"

    ```bash title="Download and extract the release"
    wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
    tar -xzf telegram-commander.tar.gz
    cd telegram-commander
    ```

Inside the folder you will find:

- `telegram-commander-linux-amd64` and `telegram-commander-linux-arm64` — the program, one per CPU type
- `config-examples/` — ready-made [config files](../concepts/config-file.md) (see [Configuration](../configuration.md))
- `functions/` — example custom [functions](../concepts/function.md) (see [Functions](../functions.md))

## :material-chip: Step 2: pick your binary

!!! info "Which file is for your machine?"

    Most servers and PCs are `amd64` (also called `x86_64`). Small ARM boards
    and some cloud VMs are `arm64`.

    If you are unsure, run `uname -m`: `x86_64` means amd64, `aarch64` means
    arm64.

=== ":fontawesome-brands-linux: AMD64"

    ```bash title="Keep the amd64 program"
    mv telegram-commander-linux-amd64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-arm64
    ```

=== ":fontawesome-brands-linux: ARM64"

    ```bash title="Keep the arm64 program"
    mv telegram-commander-linux-arm64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-amd64
    ```

Now you have a single program named `telegram-commander`.

## :material-file-cog-outline: Step 3: create your config

Copy the minimal example to a working file:

!!! example "Make an editable config"

    ```bash title="Copy the example config"
    cp config-examples/config.minimal.yaml ./config.yaml
    ```

Open `config.yaml` and replace two placeholders:

- `YOUR_BOT_TOKEN` — the token from BotFather
- `YOUR_USER_ID` — your numeric id (or leave it for now and see step 5)

To learn what every setting means, read [Configuration](../configuration.md).

## :material-file-check-outline: Step 4: validate

Always check the config before running. This catches typos and mistakes without
starting the bot.

!!! success "Check that the config works"

    ```bash title="Validate the config"
    ./telegram-commander validate --config config.yaml
    ```

If it prints `Valid configuration`, you are good. If not, it lists exactly what
is wrong and where. See the [CLI page](../cli.md#validate) for details.

## :material-account-search: Step 5: find your user id (if needed)

If you did not know your user id, set only the token in `config.yaml`, put any
number in `allowed_users` for now, then run the bot:

!!! info "Start once to see your user id"

    ```bash title="Run the bot to learn your id"
    ./telegram-commander run --config config.yaml
    ```

Open Telegram, find your bot, and send it any message. Because you are not in
[allowed users](../concepts/allowed-users.md) yet, the bot replies with your
`user_id` and `username`. Copy that id into `allowed_users`, stop the bot with
`Ctrl+C`, and run it again.

This behavior is part of how access control works; see
[Configuration → telegram](../configuration.md#telegram).

## :material-play-circle-outline: Step 6: run

!!! example "Start the bot in the terminal"

    ```bash title="Start the bot"
    ./telegram-commander run --config config.yaml
    ```

Open your bot in Telegram and send `/start`. You should see your menu. Tap a
[button](../concepts/button.md) to run its command.

!!! success "Your bot is live"

    The menu you described in `config.yaml` is now in your chat, and every tap
    runs its command on this machine.

To keep the bot running after you log out of the server, set it up as a service.
See [Run as a service](run-as-a-service.md).

## :material-map-marker-path: What next

- Add more [buttons](../concepts/button.md) and [categories](../concepts/category.md): [Menu](../concepts/menu.md)
- Understand what actually runs: [Functions](../functions.md)
- See every command line option: [CLI](../cli.md)
