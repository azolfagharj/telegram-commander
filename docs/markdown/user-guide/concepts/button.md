---
title: Button
description: A tappable menu item that runs a function. On the button you can set only command, path, and args; other values come from the function defaults.
icon: material/gesture-tap-button
---

# :material-gesture-tap-button: Button

A tappable item in the Telegram menu. A button has a `name` and, when tapped,
runs a [function](function.md) on your server and sends the result back to the
chat.

Buttons and [categories](category.md) together form the [menu](menu.md) tree
under the top-level `menu` key in your [config file](config-file.md). A button
does the work; a category only opens a submenu.

## What a button looks like

Buttons appear on the keyboard under the message box, two per row by default.
The name you choose is the text on the key, so keep it short enough to read on
a phone. See [Menu → How the Telegram menu looks](menu.md#how-the-telegram-menu-looks).

## :material-format-list-checks: The parts of a button

!!! example "Every line controls one part of the button"

    ```yaml title="One button, fully labelled"
    - name: Restart nginx # (1)!
      type: button # (2)!
      icon: "🔄" # (3)!
      function: command # (4)!
      command: "systemctl restart nginx" # (5)!
      confirm: true # (6)!
    ```

    1.  The text on the key in Telegram. It must be unique among its neighbours in
        the same menu.
    2.  Always `button` here. Use `category` instead when you want a submenu.
    3.  Optional emoji shown before the name. It is decoration only and changes
        nothing about what runs.
    4.  Which [function](function.md) to use. `command` is the built-in one that
        runs a shell command.
    5.  What the `command` function runs. Anything you could type in a terminal
        works here.
    6.  Optional. Ask "Are you sure?" before running. Leave it out for anything
        that only reads.

## What happens when you tap it

1.  The bot posts a short **Running** line so you know it started.
2.  The command runs on the machine where the bot is running.
3.  The output comes back as a code block, with the exit code and how long it
    took. Long output arrives as several messages in a row.
4.  You stay on the same menu you were on, so **Back** still leaves that
    category.

## :material-code-braces: Common buttons

=== "Check something"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

=== "Manage a service"

    ```yaml title="Restart nginx button"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

=== "Read a log"

    ```yaml title="Nginx log button"
    - name: Nginx log
      type: button
      function: command
      command: "journalctl -u nginx -n 50 --no-pager"
    ```

=== "Run a script"

    ```yaml title="Nightly backup button"
    - name: Nightly backup
      type: button
      function: script
      path: "/usr/local/bin/backup.sh"
    ```

=== "Something destructive"

    ```yaml title="Stop nginx button"
    - name: Stop nginx
      type: button
      icon: "🛑"
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## :material-emoticon-outline: Icons

`icon` puts an emoji in front of the name. It is purely cosmetic, so you can
change or remove it at any time without touching what the button runs.

!!! example "An icon only changes the label"

    ```yaml title="The same button, with and without an icon"
    - name: Disk usage
      type: button
      function: command
      command: "df -h"

    - name: Disk usage
      type: button
      icon: "💾"
      function: command
      command: "df -h"
    ```

## :material-help-circle-outline: Ask before risky buttons

Add `confirm: true` and the bot asks Yes or Cancel first. Use it for anything
that stops a service, deletes data, or reboots the machine. The prompt expires
after a while (five minutes by default).

Read [Confirmation](confirmation.md) for when it is worth asking and how to
change the waiting time.

## Settings for one button only

Most global settings can be overridden on a single button, which is handy when
one job behaves differently from the rest:

!!! example "Override settings for one slow job"

    ```yaml title="A slow job that runs somewhere else"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"
      workdir: "/var/backups"
      env:
        BACKUP_MODE: "full"
    ```

`timeout` gives this one command longer to finish, `workdir` chooses the
directory it runs in, and `env` adds environment variables just for it.

## Configuration

For every field a button accepts, see
[Configuration → Menu](../configuration.md#menu).

## Related

<div class="grid cards cols-2" markdown>

-   :material-folder-outline:{ .middle } __Category__

    ---

    Opens a submenu instead of running something.

    [:octicons-arrow-right-24: Category](category.md)

-   :material-function:{ .middle } __Function__

    ---

    What actually runs when a button is tapped.

    [:octicons-arrow-right-24: Function](function.md)

-   :material-tune:{ .middle } __Parameter__

    ---

    Values a function needs from the button.

    [:octicons-arrow-right-24: Parameter](parameter.md)

-   :material-view-list:{ .middle } __Menu__

    ---

    Build and organise the whole tree.

    [:octicons-arrow-right-24: Menu](menu.md)

</div>
