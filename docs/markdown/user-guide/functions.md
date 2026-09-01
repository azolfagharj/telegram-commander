---
icon: material/function-variant
title: Functions
description: Built-in command and script functions, the ready-made functions in the release pack, and how to write your own function files for Telegram Commander buttons.
---

# :material-function-variant: Functions

This page explains what a function is, the two kinds of functions (built-in and
custom), the ready-made custom functions that ship in the release, and how to
write your own. It is meant to be read top to bottom the first time.

If you have not yet seen how buttons reference functions, read
[Menu](concepts/menu.md) first, and check [Concepts](concepts/function.md) for the
vocabulary.

## :material-help-circle-outline: What is a function?

A **[function](concepts/function.md)** is a recipe that turns some named values (its **[parameters](concepts/parameter.md)**)
into a shell command. When you tap a [button](concepts/button.md), the bot:

1. Looks up the function named in the button's `function` field.
2. Collects the button's parameter values.
3. Builds the shell command from those values.
4. Runs the command in the [shell](concepts/shell.md) and sends the output back as a code block.

Think of a function as a fill-in-the-blanks command. For example, a "disk
usage" function has one blank — the path — and you fill it in on each button.
A button can fill three blanks by name: `command`, `path`, and `args`. See
[Passing values from a button](#passing-values-from-a-button).

### A worked example

This button uses the built-in `command` function:

!!! example "Use the built-in command function"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

The bot runs `uptime` on the server and sends the output back as a code block. To make reusable
fill-in-the-blanks commands, see [Writing custom functions](#writing-custom-functions).

## :material-compare-horizontal: The two kinds of functions

There are **built-in** functions and **custom** functions. They behave the same
way when a button uses them; the difference is where they come from.

| | Built-in | Custom |
|---|---|---|
| Where it comes from | Ships inside the program | A YAML file you write |
| Do you create a file? | No | Yes, one file per function |
| Names | Reserved (`command`, `script`) — you cannot reuse them | Any valid name that is not reserved |
| Always available? | Yes | Only if you set `function_directory` |
| Editable? | No | Yes, they are your files |

You can mix both freely in the same menu.

## :material-package-variant-closed: Built-in functions

Two functions are built in. They need no files and are always available. Their
names are **reserved**: a custom function file may not use the name `command` or
`script` (validation will reject it).

### `command`

Runs a single shell command exactly as written.

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

Because commands run through [`/bin/bash -c`](concepts/shell.md), you can use
pipes, redirects, and `&&`:

!!! example "Pipes work inside commands"

    ```yaml title="Top processes button"
    - name: Top processes
      type: button
      function: command
      command: "ps aux --sort=-%mem | head -n 10"
    ```

### `script`

Runs a script file, with optional arguments.

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

This produces the command `/usr/local/bin/report.sh --today`. If you omit
`args`, it is just `/usr/local/bin/report.sh`. The script must be executable, or
you should invoke it through an interpreter (see the `echo-script` example
below, which calls `bash`).

## :material-variable: Passing values from a button

This is the most important rule on this page. A button can pass **three
values** to a function, and each one is a field written directly on the button:

| Button field | Fills the parameter named |
|--------------|---------------------------|
| `command` | `command` |
| `path` | `path` |
| `args` | `args` |

There is no other way. A button has a fixed set of allowed fields (see
[Configuration → Menu](configuration.md#menu)), and any other key is refused
when the config is read.

!!! warning "`params:` on a button is not valid"

    ```yaml title="This config does not load"
    - name: Check API
      type: button
      function: curl-url
      params:                  # refused: not a button field
        url: "https://example.com/health"
    ```

    [`validate`](cli.md#validate) stops with an unknown-field error, so the bot
    never starts with a config like this.

A function is therefore usable from a button when **either**:

- every required parameter it has is named `command`, `path`, or `args`, **or**
- its other parameters are optional and their defaults are what you want.

If a required parameter has some other name, the button cannot fill it and
`validate` reports something like
`required parameter "url" for function "curl-url" is missing`.

When you [write your own function](#writing-custom-functions), name the value
you want to pass from the button `command`, `path`, or `args`.

## :material-folder-multiple-outline: Custom functions in the release pack

The release archive contains a `functions/` folder with example custom
functions. To load them, point `function_directory` at that folder (see
[Configuration](configuration.md#function_directory-rules)).

Because of the rule above, only some of them can be driven from a button
today:

| Function | Usable from a button? |
|----------|-----------------------|
| `echo-script` | Yes — its parameters are `path` and `args` |
| `disk-path` | Yes — its only parameter is `path` (default `/`) |
| `curl-url` | No — needs `url` |
| `journal-unit` | No — needs `unit` |
| `ping-host` | No — needs `host` |

The last three are useful as reading material: they show how a function file is
written. Treat them as templates to copy, not as buttons you can add as they
are. To get the same result today, either rename their parameter to `path` or
`args`, or use the built-in `command` function with the full command written
out.

### `echo-script`

Runs a script through `bash`, so the script itself does not need the execute
bit.

- Runs: `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (required): path to the script file
- `args` (optional): extra arguments

!!! example "Run a script through Bash"

    ```yaml title="Run cleanup button"
    - name: Run cleanup
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

### `disk-path`

Shows disk usage for a path with `df -h`.

- Runs: `df -h {{.path}}`
- `path` (optional, default `/`): filesystem path to check

!!! example "Use the default path or choose one"

    ```yaml title="Disk usage buttons"
    - name: Disk (root)
      type: button
      function: disk-path        # no path field: uses the default "/"
    - name: Disk (var)
      type: button
      function: disk-path
      path: "/var"
    ```

### `curl-url`

Fetches a URL with `curl`. It fails on HTTP errors (`-f`) and times out after
30 seconds.

- Runs: `curl -fsSL --max-time 30 {{.url}}`
- `url` (required): the URL to request

A button cannot fill `url`, so use the built-in `command` function for a health
check, or copy the file and rename the parameter to `path`:

!!! example "A health check you can add today"

    ```yaml title="Check API button"
    - name: Check API
      type: button
      function: command
      command: "curl -fsSL --max-time 30 https://example.com/health"
    ```

### `journal-unit`

Shows the most recent `journalctl` logs for a systemd unit.

- Runs: `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (required): unit name, for example `nginx.service`
- `lines` (optional, default `50`): how many lines to show

A button cannot fill `unit`. For one fixed service, write the command out:

!!! example "Read recent service logs"

    ```yaml title="Nginx logs button"
    - name: Nginx logs
      type: button
      function: command
      command: "journalctl -u nginx.service -n 100 --no-pager"
    ```

### `ping-host`

Pings a host a few times.

- Runs: `ping -c {{.count}} {{.host}}`
- `host` (required): hostname or IP
- `count` (optional, default `4`): number of packets

A button cannot fill `host`. For one fixed host:

!!! example "Ping a host three times"

    ```yaml title="Ping gateway button"
    - name: Ping gateway
      type: button
      function: command
      command: "ping -c 3 192.168.1.1"
    ```

## :material-pencil-plus-outline: Writing custom functions

A custom function is a single YAML file placed in your
[`function_directory`](configuration.md#function_directory-rules). The file
name does not matter (use `.yaml` or `.yml`); the function's name comes from
inside the file.

### File structure

!!! example "Define one custom function per file"

    ```yaml title="Custom function file"
    name: my-function          # required, must be unique and not reserved
    run: "echo {{.args}}"      # required, the command to run
    params:                    # optional list of parameters
      - name: args             # required for each parameter
        type: string           # optional: string (default), int, or bool
        required: true         # optional, default false
        default: ""            # optional value used when not provided
        description: Some text  # optional note for yourself
    ```

!!! tip "Name your parameters `command`, `path`, or `args`"

    Those are the only three names a button can fill (see
    [Passing values from a button](#passing-values-from-a-button)). A parameter
    with any other name only works if it is optional and its `default` is
    what you want on every button.

### The rules

The loader validates every file. Keep these rules in mind:

- **`name` is required** and must match `^[A-Za-z0-9._-]+$` (letters, digits,
  `.`, `-`, `_`). No spaces.
- **Names are unique, case-insensitively.** `Deploy` and `deploy` collide. This
  also applies across built-ins and other custom files.
- **Reserved names are forbidden:** you cannot name a custom function `command`
  or `script`.
- **`run` is required.** Write the command with placeholders: every `{{.name}}`
  is replaced by the matching parameter value.
- **Each parameter needs a `name`** matching the same character rule. Only
  `command`, `path`, and `args` can be filled from a button.
- **`type` must be** `string`, `int`, or `bool` (empty means `string`).
- **Unknown YAML keys are rejected**, so a typo like `requird:` is an error, not
  a silent no-op.
- A **required parameter with no value** makes the button fail validation, so
  problems are caught by [`validate`](cli.md#validate) before the bot
  runs.

### Placeholders in `run`

Two patterns cover almost everything:

- `{{.name}}` inserts the value of parameter `name`.
- `{{if .name}} ... {{end}}` includes the middle part only when `name` has a
  value. This is how optional parameters are added cleanly.

Example that adds an optional filter only when provided:

!!! example "Make an optional part with a placeholder"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}{{if .args}} | grep -- {{.args}}{{end}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        description: Optional text to filter for
    ```

Two buttons using it, one with the filter and one without:

!!! example "Supply values from a button"

    ```yaml title="App log buttons"
    - name: App log
      type: button
      function: tail-log
      path: "/var/log/app.log"

    - name: App errors
      type: button
      function: tail-log
      path: "/var/log/app.log"
      args: "ERROR"
    ```

### Step by step: create one now

1. Make sure your config has a `function_directory`, for example:

   ```yaml title="config.yaml"
   function_directory: "./functions"
   ```

2. Create `functions/greet.yaml`. The parameter is called `args` so a button
   can fill it:

   ```yaml title="functions/greet.yaml"
   name: greet
   run: "echo Hello {{.args}}"
   params:
     - name: args
       required: true
       description: Name to greet
   ```

3. Add a button that uses it:

   ```yaml title="Say hello button"
   - name: Say hello
     type: button
     function: greet
     args: "world"
   ```

4. Validate, then confirm it is loaded:

   ```bash title="Validate, then list the functions"
   ./telegram-commander validate --config config.yaml
   ./telegram-commander list-functions --config config.yaml
   ```

   You should see a `greet` line. Each line shows the name, where the function
   came from, and how many parameters it has. See
   [CLI → list-functions](cli.md#list-functions).

5. Run the bot and tap **Say hello**. It runs `echo Hello world`.

## :material-shield-alert-outline: Safety notes

!!! warning "Buttons run with the bot's privileges"

    Commands run with the privileges of the account running the bot. If that is
    root (the default [service](installation/run-as-a-service.md) setup),
    buttons can do anything on the host. Only add
    [allowed users](configuration.md#telegram) you trust.

    Parameter values are inserted into the command as text. Treat them like
    shell input: prefer fixed values on buttons, and add
    [`confirm: true`](concepts/confirmation.md) to anything destructive.

!!! info "Long output is cut and split"

    Commands stop at their `timeout`, and the bot keeps at most
    `max_output_bytes` of their output. Anything longer than one Telegram
    message arrives as several messages. See
    [Configuration → How much command output you see](configuration.md#how-much-command-output-you-see).

## Related pages

- [Menu](concepts/menu.md) — how buttons reference functions
- [Function](concepts/function.md) — what a function is
- [Parameter](concepts/parameter.md) — values a function needs
- [Configuration](configuration.md) — `function_directory` and button fields
- [CLI](cli.md) — `validate` and `list-functions`
