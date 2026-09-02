---
title: Category
description: Group buttons into submenus. A category holds more buttons or categories under items, so a long menu stays short and easy to tap on a phone.
icon: material/folder-outline
---

# :material-folder-outline: Category

A menu node that opens a submenu instead of running something. It has `items`
(more [buttons](button.md) or categories) instead of a [function](function.md).

!!! example "A category holding one button"

    ```yaml title="A System category"
    - name: System
      type: category
      items:
        - name: Uptime
          type: button
          function: command
          command: "uptime"
    ```

Buttons and categories together form a tree. See [Menu](menu.md).

## Configuration

For category fields (`type`, `items`, `columns`, and more), see
[Configuration → Menu](../configuration.md#menu).

## Related

- [Button](button.md) — runs a function when tapped
- [Menu](menu.md) — nesting, layout, and pagination
