---
title: Telegram Commander
description: Turn a YAML file into a Telegram bot that runs commands on your Linux server, one tap at a time.
icon: material/cellphone-link
hide:
  - navigation
  - toc
---

# :material-cellphone-link: Control your Linux server from Telegram

**One tap runs a command on your server and sends the output back to your chat.**

Telegram Commander turns a simple YAML [config file](user-guide/concepts/config-file.md)
into a Telegram bot with a menu of [buttons](user-guide/concepts/button.md). Put any
terminal command behind a button and run it from your phone. You do not write any code.

[Get started :material-arrow-right:](user-guide/installation/download-and-run.md){ .md-button .md-button--primary }
[See a config :material-file-code-outline:](user-guide/configuration.md#a-minimal-config){ .md-button }

[Installation](user-guide/installation/download-and-run.md) ·
[Concepts](user-guide/concepts/config-file.md) ·
[Functions](user-guide/functions/index.md) ·
[Configuration](user-guide/configuration.md) ·
[CLI](user-guide/cli.md)

## :material-star-outline: Why use it

<div class="grid cards cols-3" markdown>

-   :material-file-document-outline:{ .middle } __No coding__

    ---

    Describe the menu and commands in one YAML file.

    [:octicons-arrow-right-24: Config file](user-guide/concepts/config-file.md)

-   :material-lan-disconnect:{ .middle } __No open ports__

    ---

    The bot connects out to Telegram. Nothing is exposed to the internet.

    [:octicons-arrow-right-24: How the bot connects](user-guide/concepts/long-polling.md)

-   :material-shield-lock:{ .middle } __Controlled and recorded__

    ---

    Choose who gets the menu, confirm risky actions, and record each run.

    [:octicons-arrow-right-24: Access and confirmation](user-guide/concepts/allowed-users.md)

</div>

## :material-image-multiple-outline: Screenshots

Your menu, a command running, the output that comes back, and typing a command
by hand. Click any picture to see it full size.

![Command output with the System menu open](images/01.jpeg){ width="140" loading=lazy }
![Resource and process buttons](images/02.jpeg){ width="140" loading=lazy }
![Storage and package buttons](images/03.jpeg){ width="140" loading=lazy }
![Network tools and typing a command by hand](images/04.jpeg){ width="140" loading=lazy }
![Command output sent back to the chat](images/05.jpeg){ width="140" loading=lazy }

## :material-arrow-decision-outline: How it works

<div class="grid cards cols-3" markdown>

-   :material-file-document-outline:{ .middle } __1. Write your menu__

    ---

    List your buttons and the command each one runs.

-   :material-rocket-launch:{ .middle } __2. Start the bot__

    ---

    Run it now, or keep it up as a service.

-   :material-gesture-tap-button:{ .middle } __3. Tap and read__

    ---

    Tap a button and read the output in the chat.

</div>

## :material-file-code-outline: A tiny example

This config makes a bot with one button called "Uptime". Tapping it runs the
`uptime` command on the server.

!!! example "This complete config adds one button"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN" # (1)!
      allowed_users:
        - "YOUR_USER_ID" # (2)!

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime" # (3)!
    ```

    1.  Ask BotFather in Telegram for a token when you create the bot.
    2.  Only the accounts listed here can open the menu. You can use a numeric id
        or an `@username`.
    3.  Anything you could type in a terminal goes here.

That is a complete, working configuration. Everything else is optional.

## :material-view-grid-outline: What you can do with it

<div class="grid cards cols-3" markdown>

-   :material-restart:{ .middle } __Manage__

    ---

    Control services and packages.

-   :material-chart-line:{ .middle } __Observe__

    ---

    Check resources, logs, networks, and health.

-   :material-script-text:{ .middle } __Automate__

    ---

    Run deploys, backups, cleanups, and scripts.

    [:octicons-arrow-right-24: Functions](user-guide/functions/index.md)

</div>

!!! tip "Built for a small, trusted group"

    Nothing listens for incoming connections, and only the accounts in
    `allowed_users` get a menu. Everyone who can use the bot can run the
    buttons you defined, so keep that list short.

## :material-hand-pointing-right: Ready to try it?

[Latest release :material-download:](https://github.com/azolfagharj/telegram-commander/releases/latest){ .md-button .md-button--primary }
[Browse the source :fontawesome-brands-github:](https://github.com/azolfagharj/telegram-commander){ .md-button }

Telegram Commander is free and open source. If it saves you time,
[consider supporting its development](https://azolfagharj.github.io/donate/) —
it helps keep the project alive and maintained.
