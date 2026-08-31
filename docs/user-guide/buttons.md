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

Send `/start` in Telegram. The bot keeps **one menu message** and updates it
when you tap. Action buttons sit on that message (inline keyboard). A small
reply keyboard at the bottom repeats Home, and Back / Prev / Next only when
those actions work.

## Grouping with categories

As your menu grows, group related actions into categories. A category shows its
`items` when tapped. Home is always on the menu. Back appears only inside a
category.

```yaml
buttons:
  - name: System
    type: category
    icon: "🖥️"
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
  icon: "♻️"
  function: command
  command: "reboot"
  confirm: true
```

## How the Telegram menu looks

The last menu message is reused: the bot edits it instead of sending a new
menu each time.

- **On the message (inline):** `Home` is always there. Category and button
  items use two columns by default (`buttons_columns`), at most `page_size`
  items per page (default 8). `Back` appears only inside a category. `Prev` /
  `Next` appear only when there is another page.
- **Bottom reply keyboard:** the same nav buttons, with the same rules. Menu
  items are not duplicated there.

Tapping Home, Back, Prev, Next, or an item updates that same message.

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

`buttons_columns` sets how many **item** buttons appear per row on the message
(default 2). A category can override it with `columns`. When a menu has more
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
