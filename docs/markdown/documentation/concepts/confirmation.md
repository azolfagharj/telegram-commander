---
title: Confirmation
description: Make the bot ask before a button runs. Turn on confirm for anything you cannot undo, and set how long the Yes or Cancel prompt stays valid.
icon: material/help-circle-outline
---

# :material-help-circle-outline: Confirmation

An optional "Are you sure?" step before a [button](button.md) runs, enabled
with `confirm: true`. Useful for destructive actions.

The bot asks (`Confirm: … ?`) with Yes / Cancel, plus Home (and Back when you
are inside a category). When Run Command is enabled, that button stays on the
confirm screen too. Nothing runs until you tap Yes; Cancel, Home, and Back all
leave the command alone.

The prompt expires after a while (default 5 minutes). Change it with
`confirm_ttl` in your [config file](config-file.md). If it expires, tap the
button again to get a fresh prompt.

!!! warning "Use it for anything you cannot undo"

    A menu key is one tap away on a phone, and everyone allowed to use the bot
    sees the same menu. Put `confirm: true` on buttons that stop a service,
    delete data, or reboot the machine.

!!! example "A button that asks before it runs"

    ```yaml title="Stop nginx button"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## Configuration

For `confirm` on buttons and global `confirm_ttl`, see
[Configuration](../configuration.md).

## Related

- [Button](button.md) — the node that can require confirmation
- [Menu → Confirmation](menu.md#confirmation) — examples and behavior
