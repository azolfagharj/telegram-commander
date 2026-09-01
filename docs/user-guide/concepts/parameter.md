# Parameter

A named value a [function](function.md) needs. For example, the built-in
`command` function needs one parameter called `command`. You provide parameter
values on the [button](button.md).

Some parameters are required; others are optional with a default defined in the
function file.

## Example

Using the shortcut field on a button:

```yaml
- name: Uptime
  type: button
  function: command
  command: "uptime"
```

The same value through `params`:

```yaml
- name: Uptime
  type: button
  function: command
  params:
    command: "uptime"
```

For custom functions, use `params` for any parameter name. See
[Functions → Passing parameters from a button](../functions.md#passing-parameters-from-a-button).

## Configuration

For button fields `command`, `path`, `args`, and `params`, see
[Configuration → Menu](../configuration.md#menu).

## Related

- [Function](function.md) — what uses parameters
- [Functions](../functions.md) — parameter rules for custom functions
