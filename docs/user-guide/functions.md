# Functions

This page explains what a function is, the two kinds of functions (built-in and
custom), the ready-made custom functions that ship in the release, and how to
write your own. It is meant to be read top to bottom the first time.

If you have not yet seen how buttons reference functions, read
[Buttons and menus](buttons.md) first, and check [Concepts](concepts/function.md) for the
vocabulary.

## What is a function?

A **[function](concepts/function.md)** is a recipe that turns some named values (its **[parameters](concepts/parameter.md)**)
into a shell command. When you tap a [button](concepts/button.md), the bot:

1. Looks up the function named in the button's `function` field.
2. Collects the button's parameter values.
3. Builds the shell command from those values.
4. Runs the command in the [shell](concepts/shell.md) and sends the output back.

Think of a function as a fill-in-the-blanks command. For example, a "ping a
host" function has one blank — the host — and you fill it in on each button.

### A worked example

This button uses the built-in `command` function:

```yaml
- name: Uptime
  type: button
  function: command
  command: "uptime"
```

The bot runs `uptime` on the server and sends the output back. To make reusable
fill-in-the-blanks commands, see [Writing custom functions](#writing-custom-functions).

## The two kinds of functions

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

## Built-in functions

Two functions are built in. They need no files and are always available. Their
names are **reserved**: a custom function file may not use the name `command` or
`script` (validation will reject it).

### `command`

Runs a single shell command exactly as written.

| Parameter | Required | Default | Meaning |
|-----------|----------|---------|---------|
| `command` | yes | — | The shell command to run |

```yaml
- name: Show kernel
  type: button
  function: command
  command: "uname -a"
```

Because commands run through [`/bin/bash -c`](concepts/shell.md), you can use
pipes, redirects, and `&&`:

```yaml
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

```yaml
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

## Custom functions in the release pack

The release archive contains a `functions/` folder with ready-to-use custom
functions. To use them, point `function_directory` at that folder (see
[Configuration](configuration.md#function_directory-rules)) and reference each
by name from a button.

Below is exactly what each one does.

### `echo-script`

Runs a script through `bash`, so the script itself does not need the execute
bit.

- Runs: `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (required): path to the script file
- `args` (optional): extra arguments

```yaml
- name: Run cleanup
  type: button
  function: echo-script
  path: "/opt/scripts/cleanup.sh"
  args: "--verbose"
```

### `curl-url`

Fetches a URL with `curl`. Handy for health checks. It fails on HTTP errors
(`-f`) and times out after 30 seconds.

- Runs: `curl -fsSL --max-time 30 {{.url}}`
- `url` (required): the URL to request

```yaml
- name: Check API
  type: button
  function: curl-url
  params:
    url: "https://example.com/health"
```

Note: `url` is not one of the built-in button fields (`command`, `path`,
`args`), so it goes under `params`. See
[Passing parameters from a button](#passing-parameters-from-a-button).

### `journal-unit`

Shows the most recent `journalctl` logs for a systemd unit.

- Runs: `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (required): unit name, for example `nginx.service`
- `lines` (optional, default `50`): how many lines to show

```yaml
- name: Nginx logs
  type: button
  function: journal-unit
  params:
    unit: "nginx.service"
    lines: "100"
```

### `disk-path`

Shows disk usage for a path with `df -h`.

- Runs: `df -h {{.path}}`
- `path` (optional, default `/`): filesystem path to check

```yaml
- name: Disk (root)
  type: button
  function: disk-path        # uses the default path "/"
- name: Disk (var)
  type: button
  function: disk-path
  params:
    path: "/var"
```

### `ping-host`

Pings a host a few times.

- Runs: `ping -c {{.count}} {{.host}}`
- `host` (required): hostname or IP
- `count` (optional, default `4`): number of packets

```yaml
- name: Ping gateway
  type: button
  function: ping-host
  params:
    host: "192.168.1.1"
    count: "3"
```

## Passing parameters from a button

A button supplies parameter values in one of two ways:

1. **Shortcut fields** for the built-in functions: `command`, `path`, and
   `args` can be written directly on the button.
2. **The `params` map** for everything else. Any parameter name works here.

These two buttons are equivalent:

```yaml
# Using the shortcut field
- name: A
  type: button
  function: command
  command: "uptime"

# Using params
- name: B
  type: button
  function: command
  params:
    command: "uptime"
```

For custom functions with their own parameter names (like `url` or `unit`),
always use `params`.

## Writing custom functions

A custom function is a single YAML file placed in your
[`function_directory`](configuration.md#function_directory-rules). The file
name does not matter (use `.yaml` or `.yml`); the function's name comes from
inside the file.

### File structure

```yaml
name: my-function          # required, must be unique and not reserved
run: "echo {{.text}}"      # required, the command to run
params:                    # optional list of parameters
  - name: text             # required for each parameter
    type: string           # optional: string (default), int, or bool
    required: true         # optional, default false
    default: ""            # optional value used when not provided
    description: Some text  # optional, shown by list-functions
```

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
- **Each parameter needs a `name`** matching the same character rule.
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

Example that adds an optional flag only when provided:

```yaml
name: tail-log
run: "tail -n {{.lines}}{{if .follow}} -f{{end}} {{.file}}"
params:
  - name: file
    required: true
    description: Log file path
  - name: lines
    default: "100"
    description: Lines to show
  - name: follow
    description: Set to any value to stream new lines
```

A button using it:

```yaml
- name: App log
  type: button
  function: tail-log
  params:
    file: "/var/log/app.log"
    lines: "200"
```

### Step by step: create one now

1. Make sure your config has a `function_directory`, for example:

   ```yaml
   function_directory: "./functions"
   ```

2. Create `functions/greet.yaml`:

   ```yaml
   name: greet
   run: "echo Hello {{.who}}"
   params:
     - name: who
       required: true
       description: Name to greet
   ```

3. Add a button that uses it:

   ```yaml
   - name: Say hello
     type: button
     function: greet
     params:
       who: "world"
   ```

4. Validate, then confirm it is loaded:

   ```bash
   ./telegram-commander validate --config config.yaml
   ./telegram-commander list-functions --config config.yaml
   ```

   You should see `greet` in the list. See
   [CLI → list-functions](cli.md#list-functions).

5. Run the bot and tap **Say hello**. It runs `echo Hello world`.

## Safety notes

- Commands run with the privileges of the account running the bot. If that is
  root (the default [service](installation/run-as-a-service.md) setup), buttons can do anything on the
  host. Only add [allowed users](configuration.md#telegram) you trust.
- Parameter values are inserted into the command as text. Treat them like shell
  input: prefer fixed values on buttons, and add
  [`confirm: true`](concepts/confirmation.md) to anything destructive.
- Output is truncated at `max_output_bytes` and commands stop at their
  `timeout`. See [Configuration → Root fields](configuration.md#root-fields).

## Related pages

- [Buttons and menus](buttons.md) — how buttons reference functions
- [Function](concepts/function.md) — what a function is
- [Parameter](concepts/parameter.md) — values a function needs
- [Configuration](configuration.md) — `function_directory` and button fields
- [CLI](cli.md) — `validate` and `list-functions`
