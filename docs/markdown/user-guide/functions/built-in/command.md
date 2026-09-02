---
title: command
description: The built-in command function runs a single shell command exactly as written on the button, including pipes, redirects and chained commands.
icon: material/console
---

# :material-console: `command`

`command` is a [built-in function](../index.md#built-in-functions). It runs a single shell command
exactly as you wrote it on the button.

| Parameter | Required | Default | Meaning |
|-----------|----------|---------|---------|
| `command` | yes | — | The shell command to run |

!!! example "Run one command as written"

    ```yaml title="Show kernel button"
    - name: Show kernel
      type: button
      function: command
      command: "uname -a"
    ```

## Pipes and chained commands

Commands run through [`/bin/bash -c`](../../concepts/shell.md), so pipes,
redirects, and `&&` all work inside one `command` value.

!!! example "Pipes work inside commands"

    ```yaml title="Top processes button"
    - name: Top processes
      type: button
      function: command
      command: "ps aux --sort=-%mem | head -n 10"
    ```

!!! tip "Use a custom function for repeated commands"

    If several buttons repeat the same command shape, create a custom function
    and give each button only the values that change. Custom parameter names
    such as `url` or `host` can be written directly on those buttons.

## Related

- [`script`](script.md) — run a script file instead of an inline command
- [Built-in functions](../index.md#built-in-functions) — both built-in functions
- [Shell](../../concepts/shell.md) — how commands are executed
- [Confirmation](../../concepts/confirmation.md) — ask before a risky command
