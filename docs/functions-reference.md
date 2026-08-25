# Functions reference

Functions define how a button turns into a runnable shell command.

## Built-in functions

### `command`

Runs an arbitrary shell command.

Parameters:

| Name | Required | Description |
|------|----------|-------------|
| `command` | yes | Command string passed to the configured shell |

Button example:

```yaml
- name: Restart nginx
  type: button
  function: command
  command: "systemctl restart nginx"
  confirm: true
```

### `script`

Runs a script path with optional arguments.

Parameters:

| Name | Required | Description |
|------|----------|-------------|
| `path` | yes | Script path |
| `args` | no | Arguments appended after the path |

## Custom functions

Place one YAML file per function under `function_directory` (recursive, `.yaml` / `.yml`).

File name does not matter. The `name` field inside the file is the function name.

Rules:

- Name pattern: letters, digits, `.`, `-`, `_`
- Names are unique case-insensitively (`MyFunc` and `myfunc` conflict)
- Reserved names cannot be overridden: `command`, `script`
- `run` is a Go `text/template` rendered with parameter values

Example (`echo-script.yaml`):

```yaml
name: echo-script
run: "bash {{.path}}{{if .args}} {{.args}}{{end}}"
params:
  - name: path
    type: string
    required: true
    description: Path to the script file
  - name: args
    type: string
    required: false
    description: Optional script arguments
```

Then reference it from a button:

```yaml
- name: Run helper
  type: button
  function: echo-script
  path: /usr/local/bin/helper.sh
  args: "--verbose"
```

## Sample functions in the release pack

The release archive includes a `functions/` folder with examples:

| Name | Purpose |
|------|---------|
| `echo-script` | Run a bash script with optional args |
| `curl-url` | HTTP GET a URL with curl |
| `journal-unit` | Show recent `journalctl` lines for a unit |
| `disk-path` | Show `df -h` for a path |
| `ping-host` | Ping a host |

The minimal config does not enable `function_directory`. The full example config points at `./functions` after extract (or `examples/functions` in this repository).

## Validation

`telegram-commander validate --config ...` checks:

- Config schema and required fields
- Function name uniqueness and reserved names
- Button → function references
- Required parameters present on buttons
- Run template syntax
