---
icon: material/tune
title: Parameter
description: A parameter is a named value a function needs. A button can fill only command, path, and args; other parameters must be optional with a default.
---

# :material-tune: Parameter

A named value a [function](function.md) needs. For example, the built-in
`command` function needs one parameter called `command`. You provide parameter
values on the [button](button.md).

Some parameters are required; others are optional with a default defined in the
function file.

## Only three names come from a button

A button can pass exactly three values, and each is a field on the button
itself:

| Button field | Fills the parameter named |
|--------------|---------------------------|
| `command` | `command` |
| `path` | `path` |
| `args` | `args` |

!!! example "The value is a field on the button"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

A button has no other place to put a value. Writing `params:` on a button is an
error and the config will not load.

!!! warning "`params:` belongs in a function file, not on a button"

    ```yaml title="This config does not load"
    - name: Ping gateway
      type: button
      function: ping-host
      params:                  # refused: not a button field
        host: "192.168.1.1"
    ```

So a function works from a button when every **required** parameter it has is
named `command`, `path`, or `args`. Any other parameter must be optional, and
it always uses the default from the function file. If a required parameter has
another name, [`validate`](../cli.md#validate) fails and tells you which
parameter is missing.

See [Functions → Passing values from a button](../functions.md#passing-values-from-a-button)
for the full explanation and examples.

## Configuration

For the button fields `command`, `path`, and `args`, see
[Configuration → Menu](../configuration.md#menu).

## Related

- [Function](function.md) — what uses parameters
- [Functions](../functions.md) — parameter rules for custom functions
