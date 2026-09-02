---
title: What is a function?
description: A function turns values written on a button into one shell command. Learn how a button supplies the parameters that each function declares.
---

# :material-function-variant: What is a function?

A **function** is a recipe that turns some named values (its
[parameters](../concepts/parameter.md)) into a shell command. Every
[button](../concepts/button.md) in your menu names exactly one function in its
`function` field.

Think of a function as a fill-in-the-blanks command. A "disk usage" function has
one blank — the path — and you fill it in on each button.

## What happens when you tap a button

1. The bot looks up the function named in the button's `function` field.
2. It collects the values written on that button.
3. It builds one shell command from those values.
4. It runs the command in the [shell](../concepts/shell.md) and sends the output
   back to the chat as a code block.

If the function does not exist, or a value it needs is missing, the bot never
starts: [`validate`](../cli.md#validate) reports the problem first.

## A worked example

The `command` function is built in and always available. It runs whatever you
write in the button's `command` field.

!!! example "Use the built-in command function"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

Tap **Uptime** and the bot runs `uptime` on the server and sends the output back.

## Passing values from a button

Write each value directly on the button. There are two ways to do this:

1. Use the shortcut fields `command`, `path`, and `args`. Each fills the
   parameter with the same name.
2. For any other parameter, use its name as a key on the button.

!!! example "Pass a URL by its parameter name"

    ```yaml title="Check API button"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

Here `url` matches the `url` parameter declared by `curl-url`. The same rule
works for names such as `host`, `unit`, and `lines`.

!!! warning "Do not put values inside `params:`"

    `params:` belongs in a custom function file, where it declares the values
    that function accepts. On a button, write each value directly:

    ```yaml title="Values belong directly on the button"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

    Numeric YAML values do not need quotes.

[`validate`](../cli.md#validate) checks every key against the parameters
declared by the selected function. A misspelling or undeclared name fails
validation. It also checks values declared as `int` or `bool`. Optional values
use their defaults when the button leaves them out.

## Two kinds of function

Functions either ship inside the program or come from a YAML file you keep on
the server. They behave the same way once a button uses them.

| | Built-in | Custom |
|---|---|---|
| Where it comes from | Ships inside the program | A YAML file you write |
| Do you create a file? | No | Yes, one file per function |
| Names | Reserved (`command`, `script`) | Any name that is not reserved |
| Always available? | Yes | Only if you set `function_directory` |
| Editable? | No | Yes, they are your files |

You can mix both kinds freely in the same menu. Most menus start with `command`
buttons only; move to custom functions when you notice yourself repeating the
same command with a small change.

### Built-in functions

Two functions are always loaded, even when you do not set a
`function_directory`. Their `command`, `path`, and `args` fields go straight on
a button.

| Function | What it does | Required | Optional |
|----------|--------------|----------|----------|
| [`command`](built-in/command.md) | Runs one shell command as written | `command` | — |
| [`script`](built-in/script.md) | Runs a script file with arguments | `path` | `args` |

Both names are **reserved**. A custom function file may not reuse them: the
loader stops with an error such as `function name "command" is reserved`, and
the bot does not start.

### Custom functions

A custom function is a single YAML file that describes one reusable command.
Keep those files in a folder of your own and point
[`function_directory`](../configuration.md#function_directory-rules) at it.

!!! example "Tell the bot where your function files live"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

The bot reads that folder when it starts, including sub-folders, and picks up
every `.yaml` and `.yml` file in it. Other files are ignored.

The release archive already contains a `functions/` folder with five examples
you can use as they are:

| Function | What it does | Button values |
|----------|--------------|---------------|
| [Echo Script](release-pack/echo-script.md) | Runs a script through Bash | `path`, optional `args` |
| [Disk path](release-pack/disk-path.md) | Shows disk usage | optional `path` |
| [Curl URL](release-pack/curl-url.md) | Fetches a URL | `url` |
| [Ping Host](release-pack/ping-host.md) | Pings a host | `host`, optional `count` |
| [Journal Unit](release-pack/journal-unit.md) | Shows recent service logs | `unit`, optional `lines` |

To write your own, start with
[File structure](write-your-own/file-structure.md) or follow the
[step by step guide](write-your-own/step-by-step.md).

!!! tip "Check what got loaded"

    ```bash title="List every function the bot can see"
    ./telegram-commander list-functions --config config.yaml
    ```

    Built-in functions show `source=builtin`; custom ones show the file they
    came from.

## Safety notes

!!! warning "Buttons run with the bot's privileges"

    Commands run with the privileges of the account running the bot. If that is
    root (the default [service](../installation/run-as-a-service.md) setup),
    buttons can do anything on the host. Only add
    [allowed users](../configuration.md#telegram) you trust.

    Parameter values are inserted into the command as text. Treat them like
    shell input: prefer fixed values on buttons, and add
    [`confirm: true`](../concepts/confirmation.md) to anything destructive.

!!! info "Long output is cut and split"

    Commands stop at their `timeout`, and the bot keeps at most
    `max_output_bytes` of their output. Anything longer than one Telegram
    message arrives as several messages. See
    [Configuration → How much command output you see](../configuration.md#how-much-command-output-you-see).

## Related

- [`command`](built-in/command.md) — run one shell command
- [`script`](built-in/script.md) — run a script file
- [Step by step guide](write-your-own/step-by-step.md) — build your first custom function
- [Menu](../concepts/menu.md) — how buttons reference functions
- [Parameter](../concepts/parameter.md) — the named values a function needs
