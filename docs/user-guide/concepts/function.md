# Function

The thing that runs when a [button](button.md) is tapped. A [function](function.md)
takes some [parameters](parameter.md) and turns them into a shell command.

There are two kinds:

- **Built-in functions** ship with the program (`command`, `script`). You do not
  create them.
- **Custom functions** are extra functions you define as YAML files in
  `function_directory`.

When you tap a button, the bot looks up the function, fills in the values from
the button, and runs the result in the [shell](shell.md).

## Example

This button uses the built-in `command` function:

```yaml
- name: Uptime
  type: button
  function: command
  command: "uptime"
```

## Configuration

For `function_directory` and button fields that reference functions, see
[Configuration](../configuration.md).

## Related

- [Parameter](parameter.md) — named values a function needs
- [Functions](../functions.md) — built-in and custom functions in depth
- [Menu](menu.md) — how buttons reference functions
