---
title: Menu
description: Build the menu your bot shows in Telegram. Mix buttons and categories into a tree, and learn how Home and Back move you through it.
icon: material/view-list
---

# :material-view-list: Menu

Your menu is a tree of nodes under the top-level `menu` key. There are two
kinds of node:

- A **[category](category.md)** opens a submenu. It has `items`.
- A **[button](button.md)** runs something. It has a [function](function.md).

If these words are new, read [Concepts](button.md) first. For the exact list
of every field, see [Configuration → Menu](../configuration.md#menu).

## :material-format-list-bulleted: A flat menu

The simplest menu is a list of buttons with no nesting:

!!! example "Make a menu with three buttons"

    ```yaml title="Three buttons, no categories"
    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"

      - name: Free memory
        type: button
        function: command
        command: "free -h"

      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

Send `/start` in Telegram. You will see your menu.

## :material-folder-outline: Grouping with categories

As your menu grows, group related actions into categories. A category shows its
`items` when tapped. Home is always on the menu. Back appears only inside a
category.

!!! example "Group buttons into categories"

    ```yaml title="System and Services categories"
    menu:
      - name: System
        type: category
        icon: "💻"
        items:
          - name: Uptime
            type: button
            function: command
            command: "uptime"

          - name: Free memory
            type: button
            function: command
            command: "free -h"

      - name: Services
        type: category
        icon: "🔧"
        items:
          - name: Restart nginx
            type: button
            function: command
            command: "systemctl restart nginx"
            confirm: true
    ```

Categories can be nested as deep as you like. A category must have at least one
item.

## :material-family-tree: Names must be unique among siblings

Two nodes under the same parent cannot share a name (comparison ignores case).
This is fine, because they are in different categories:

!!! example "Reuse a name under different parents"

    ```yaml title="The same name under two parents"
    menu:
      - name: Web
        type: category
        items:
          - name: Restart          # ok
            type: button
            function: command
            command: "systemctl restart nginx"
      - name: Database
        type: category
        items:
          - name: Restart          # ok, different parent
            type: button
            function: command
            command: "systemctl restart postgresql"
    ```

## :material-emoticon-outline: Icons

`icon` is an optional emoji shown before the name. It is purely cosmetic.

!!! warning "This button reboots the machine"

    ```yaml title="A button with an emoji icon"
    - name: Reboot
      type: button
      icon: "🔁"
      function: command
      command: "reboot"
      confirm: true
    ```

!!! tip "Pick a simple, common emoji"

    A few emoji make some phones show the button text cut off or overflowing
    the button. If a button looks cut off, try a different emoji for it.

## :material-cellphone-text: How the Telegram menu looks

All buttons appear on the keyboard under the message box (the keyboard that
shows and hides with the small button at the right end of the message box).
This keyboard always spans the full width of the chat, so button text is never
squeezed or cut off.

- **Home** is always the first button on every screen. Tap it to go back to
  the first screen.
- **Back** appears when you are inside a category.
- **$ >_ Run Command** appears when `enable_run_command` is on (see below).
- Items sit two per row by default. A category can change this with `columns`.
  If a screen has many items, **Prev** and **Next** let you page through them.
- Buttons with `confirm: true` ask Yes / Cancel before they run.

!!! info "Menu titles are reused, output stays"

    A new menu title (Home, a category, a page) replaces the previous one, so
    the chat does not fill with empty screens. The **Running** line and the
    command output stay in the chat, so you can still read what ran after you
    open the menu again.

!!! info "Long output arrives in several messages"

    Command output is shown as a code block. If it is longer than one Telegram
    message, it arrives as several messages, each a reply to the one before it.
    The last part keeps the same buttons as the page you were on, so **Back**
    still means leave that category. See
    [Configuration → How much command output you see](../configuration.md#how-much-command-output-you-see).

## :material-help-circle-outline: Confirmation

Add `confirm: true` to any button to require a second tap ("Are you sure?")
before it runs. Use it for anything destructive. See
[Confirmation](confirmation.md) for the concept.

!!! warning "This button stops a service"

    ```yaml title="A button that asks first"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

The confirm prompt expires after a while (default 5 minutes). Change it with
`confirm_ttl`; see [Configuration → Root fields](../configuration.md#root-fields).

## :material-tune-variant: Per-button overrides

Some global settings can be overridden on a single button:

!!! example "Give one button its own settings"

    ```yaml title="One button with its own timeout, folder, and variables"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"          # this one may take longer than the global timeout
      workdir: "/var/backups" # run it here
      env:
        BACKUP_MODE: "full"   # extra environment variable for this command
    ```

See the full field list in
[Configuration → Menu](../configuration.md#menu).

## :material-view-grid-outline: Controlling layout

`menu_columns` sets how many **item** buttons appear per row (default 2).
A category can override it with `columns`. When a menu has more
than `page_size` items (default 8), it is split into pages and Prev/Next are
shown until you reach the ends. See
[Configuration → Root fields](../configuration.md#root-fields).

## :material-console: Run Command

If you set `enable_run_command: true` at the root of your config, a
**$ >_ Run Command** button stays on the menu (after Back inside a category,
or after Home on the first screen). Tap it, then send the shell command you
want to run. The bot uses the same shell, timeout, working directory, and
output limits as your other buttons.

Home or Back cancels the prompt without running anything. This is off by
default. See [Configuration → Root fields](../configuration.md#root-fields).

!!! warning "This hands over the whole machine"

    With Run Command on, anyone allowed to use the bot can run any command on
    the host, not only the buttons you defined. Turn it on only if you trust
    every allowed user that far.

## :material-timer-sand: One command at a time

Your commands run one after another, not side by side.

!!! info "A second tap waits its turn"

    If you tap a second button while the first is still working, the second
    waits and then runs on its own. You will see its **Running** line once it
    starts. This keeps two of your own taps from fighting over the same
    service or file. Other people using the bot are not held up by your
    command; each person has their own turn.

## :material-function-variant: What runs when a button is tapped

Every button points to a **function** through its `function` field. The button
in the examples above uses the built-in `command` function. To understand
functions, built-in versus custom, and how to add your own, read
[Functions](../functions/index.md).

## :material-link-variant: Related pages

- [Button](button.md) — what a button is
- [Category](category.md) — submenu nodes
- [Configuration → Menu](../configuration.md#menu) — every field
