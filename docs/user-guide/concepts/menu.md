# Menu

Your menu is a tree of nodes under the top-level `menu` key. There are two
kinds of node:

- A **[category](category.md)** opens a submenu. It has `items`.
- A **[button](button.md)** runs something. It has a [function](function.md).

If these words are new, read [Concepts](button.md) first. For the exact list
of every field, see [Configuration → Menu](../configuration.md#menu).

## A flat menu

The simplest menu is a list of buttons with no nesting:

```yaml
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

## Grouping with categories

As your menu grows, group related actions into categories. A category shows its
`items` when tapped. Home is always on the menu. Back appears only inside a
category.

```yaml
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

## Names must be unique among siblings

Two nodes under the same parent cannot share a name (comparison ignores case).
This is fine, because they are in different categories:

```yaml
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

## Icons

`icon` is an optional emoji shown before the name. It is purely cosmetic.

```yaml
- name: Reboot
  type: button
  icon: "🔁"
  function: command
  command: "reboot"
  confirm: true
```

Pick a simple, common emoji. A few emoji make some phones show the button
text cut off or overflowing the button. If a button looks cut off, try a
different emoji for it.

## How the Telegram menu looks

All buttons appear under the message box (the keyboard that shows and hides
with the small button at the right end of the message box). This keyboard
always spans the full width of the chat, so button text is never squeezed or
cut off. Menu titles (Home, a category, a page) replace the previous menu
title so the chat does not fill with empty screens. A **Running** line and
the command output stay in the chat, so you can still see what ran after you
open the menu again.

- **Home** is always the first button on every screen. Tap it to go back to
  the first screen.
- **Back** appears when you are inside a category.
- **$ >_ Run Command** appears when `enable_run_command` is on (see below).
- Items sit two per row by default. A category can change this with `columns`.
  If a screen has many items, **Prev** and **Next** let you page through them.
- Buttons with `confirm: true` ask Yes / Cancel before they run.
- Command output is shown as a code block. If it is longer than one Telegram
  message, it arrives as several messages, each a reply to the one before it.
  The last part keeps the same buttons as the page you were on, so Back still
  means leave that category. See
  [Configuration → How much command output you see](../configuration.md#how-much-command-output-you-see).

## Confirmation

Add `confirm: true` to any button to require a second tap ("Are you sure?")
before it runs. Use it for anything destructive. See
[Confirmation](confirmation.md) for the concept.

```yaml
- name: Stop nginx
  type: button
  function: command
  command: "systemctl stop nginx"
  confirm: true
```

The confirm prompt expires after a while (default 5 minutes). Change it with
`confirm_ttl`; see [Configuration → Root fields](../configuration.md#root-fields).

## Per-button overrides

Some global settings can be overridden on a single button:

```yaml
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

## Controlling layout

`menu_columns` sets how many **item** buttons appear per row (default 2).
A category can override it with `columns`. When a menu has more
than `page_size` items (default 8), it is split into pages and Prev/Next are
shown until you reach the ends. See
[Configuration → Root fields](../configuration.md#root-fields).

## Run Command

If you set `enable_run_command: true` at the root of your config, a
**$ >_ Run Command** button stays on the menu (after Back inside a category,
or after Home on the first screen). Tap it, then send the shell command you
want to run. The bot uses the same shell, timeout, working directory, and
output limits as your other buttons.

Home or Back cancels the prompt without running anything. This is off by
default. Anyone who can use the bot can then run any command on the host, so
only turn it on if you trust every allowed user. See
[Configuration → Root fields](../configuration.md#root-fields).

## What runs when a button is tapped

Every button points to a **function** through its `function` field. The button
in the examples above uses the built-in `command` function. To understand
functions, built-in versus custom, and how to add your own, read
[Functions](../functions.md).

## Related pages

- [Button](button.md) — what a button is
- [Category](category.md) — submenu nodes
- [Configuration → Menu](../configuration.md#menu) — every field
