---
title: script
description: The built-in script function runs a script file on the server, with optional arguments added after the path. The script must be executable.
icon: material/script-text
---

# :material-script-text: `script`

`script` is a [built-in function](../index.md#built-in-functions). It runs a script file, with
optional arguments.

| Parameter | Required | Default | Meaning |
|-----------|----------|---------|---------|
| `path` | yes | — | Path to the script |
| `args` | no | empty | Arguments passed after the path |

!!! example "Run a script with arguments"

    ```yaml title="Nightly report button"
    - name: Nightly report
      type: button
      function: script
      path: "/usr/local/bin/report.sh"
      args: "--today"
    ```

This runs `/usr/local/bin/report.sh --today`. If you leave `args` out, it runs
just `/usr/local/bin/report.sh`.

## The script must be executable

The path is run directly, so the file needs the execute bit:

!!! tip "Make the script executable once"

    ```bash title="Allow the file to run"
    chmod +x /usr/local/bin/report.sh
    ```

If you cannot change the file, run it through an interpreter instead. The
[`echo-script`](../release-pack/echo-script.md) example in the release
pack does exactly that by calling `bash` first.

## Related

- [`command`](command.md) — run an inline command instead of a file
- [Echo Script](../release-pack/echo-script.md) — run a script through Bash
- [Built-in functions](../index.md#built-in-functions) — both built-in functions
- [Shell](../../concepts/shell.md) — how commands are executed
