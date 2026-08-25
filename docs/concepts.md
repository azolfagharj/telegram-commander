# Concepts

This page explains the words used across the documentation. Skim it once, then
refer back when a term is unclear.

## Config file

A single YAML file that describes everything: your bot token, who may use the
bot, and the menu of buttons. You pass it with `--config`. Full details are in
[Configuration](config-reference.md).

## Button

A tappable item in the Telegram menu. A button has a `name` and, when tapped,
runs a [function](#function). Example:

```yaml
- name: Uptime
  type: button
  function: command
  command: "uptime"
```

## Category

A button that opens a submenu instead of running something. It has `items`
(more buttons or categories) instead of a `function`. Categories let you group
related actions. Example:

```yaml
- name: System
  type: category
  items:
    - name: Uptime
      type: button
      function: command
      command: "uptime"
```

Buttons and categories together form a tree. See
[Buttons and menus](buttons.md).

## Function

The thing that runs when a button is tapped. A function takes some
**parameters** and turns them into a shell command. There are two kinds:

- **Built-in functions** ship with the program (`command`, `script`). You do not
  create them.
- **Custom functions** are extra functions you define as YAML files.

The full story, including how to write your own, is in
[Functions](functions-reference.md).

## Parameter

A named value a function needs. For example, the built-in `command` function
needs one parameter called `command`. You provide parameter values on the
button. Some parameters are required, some are optional with a default.

## Shell

The program that runs your command, by default `/bin/bash`. Commands run as
`bash -c "your command"`, so pipes and redirects work. You can change the shell
globally; see [Configuration](config-reference.md#root-fields).

## Allowed users

The list of Telegram accounts permitted to use the bot, by numeric `user_id` or
`@username`. Everyone else is refused. See
[Configuration → telegram](config-reference.md#telegram).

## Confirmation

An optional "Are you sure?" step before a button runs, enabled with
`confirm: true`. Useful for destructive actions. See
[Buttons and menus](buttons.md#confirmation).

## Long polling

How the bot receives messages from Telegram: it repeatedly asks Telegram for new
updates. There is nothing to configure and no inbound port to open. If Telegram
is blocked on your network, route it through a
[proxy](config-reference.md#telegram).

## Audit log

A separate log stream that records every command run: who ran it, which button,
the exit code, and how long it took. Configured under
[`logging`](config-reference.md#logging).
