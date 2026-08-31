# Buttons and menus

Your menu is a tree of nodes under the top-level `buttons` key. There are two
kinds of node:

- A **[category](concepts/category.md)** opens a submenu. It has `items`.
- A **[button](concepts/button.md)** runs something. It has a [function](concepts/function.md).

If these words are new, read [Concepts](concepts/button.md) first. For the exact list
of every field, see [Configuration → Buttons](configuration.md#buttons).

## A flat menu

The simplest menu is a list of buttons with no nesting:

```yaml
buttons:
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
buttons:
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
buttons:
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
cut off. Every screen (the menu, a confirmation, a command result) is its own
message, and the bot deletes the previous screen right after sending the new
one, so the chat stays tidy and the current screen is always the last message.

- **Home** is always the first button on every screen. Tap it to go back to
  the first screen.
- **Back** appears when you are inside a category.
- Items sit two per row by default. A category can change this with `columns`.
  If a screen has many items, **Prev** and **Next** let you page through them.
- Buttons with `confirm: true` ask Yes / Cancel before they run.

## Confirmation

Add `confirm: true` to any button to require a second tap ("Are you sure?")
before it runs. Use it for anything destructive. See
[Confirmation](concepts/confirmation.md) for the concept.

```yaml
- name: Stop nginx
  type: button
  function: command
  command: "systemctl stop nginx"
  confirm: true
```

The confirm prompt expires after a while (default 5 minutes). Change it with
`confirm_ttl`; see [Configuration → Root fields](configuration.md#root-fields).

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
[Configuration → Buttons](configuration.md#buttons).

## Controlling layout

`buttons_columns` sets how many **item** buttons appear per row (default 2).
A category can override it with `columns`. When a menu has more
than `page_size` items (default 8), it is split into pages and Prev/Next are
shown until you reach the ends. See
[Configuration → Root fields](configuration.md#root-fields).

## Running a button by name

If you enable `enable_run_command` under `telegram`, you can run a button
without opening the menu:

```
/run Restart nginx
```

This is off by default. See
[Configuration → telegram](configuration.md#telegram).

## What runs when a button is tapped

Every button points to a **function** through its `function` field. The button
in the examples above uses the built-in `command` function. To understand
functions, built-in versus custom, and how to add your own, read
[Functions](functions.md).

## Related pages

- [Button](concepts/button.md) — what a button is
- [Category](concepts/category.md) — submenu nodes
- [Configuration → Buttons](configuration.md#buttons) — every field
