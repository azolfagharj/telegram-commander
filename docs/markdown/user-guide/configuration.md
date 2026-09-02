---
icon: material/file-cog-outline
title: Configuration
description: Every Telegram Commander setting with its type, default, and meaning — telegram, menu, function_directory, timeouts, output limits, and logging.
---

# :material-file-cog-outline: Configuration

The [config file](concepts/config-file.md) describes your whole bot: the Telegram
connection, who may use it, the [button](concepts/button.md) menu, and logging.
You pass it with `--config` to the commands that read it — `run`, `validate`,
`fmt`, and `list-functions` (see [CLI](cli.md)).

All keys use `lower_snake_case`. **Unknown keys are rejected**, so a typo is an
error you will see immediately when you [validate](cli.md#validate).

**Required** means validation fails if the field is missing or empty after
defaults are applied.  
**Optional** fields may be omitted; the Default column shows what is used then.

New to the project? Start with [Run in CLI](installation/download-and-run.md),
which walks through building a first config. See [Concepts](concepts/config-file.md)
for the vocabulary used below.

## :material-rocket-launch-outline: A minimal config

Only `telegram` (with a token and one [allowed user](concepts/allowed-users.md))
and `menu` are required. Everything else has a default:

!!! example "Start with one allowed user and one button"

    ```yaml title="config.yaml (minimal)"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

The `config-examples/` folder in the release includes both a minimal and a full
example.

## :material-card-bulleted-outline: Root fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `telegram` | object | yes | — | Telegram settings (see below) |
| `menu` | list | yes | — | Menu tree; at least one node |
| `function_directory` | string | no | unset | Custom function YAML directory (see rules below) |
| `shell` | string | no | `/bin/bash` | [Shell](concepts/shell.md) used as `shell -c "<command>"` |
| `timeout` | duration | no | `60s` | Default command timeout |
| `max_output_bytes` | int | no | `524288` | Max output kept per command (see [How much command output you see](#how-much-command-output-you-see)) |
| `workdir` | string | no | process cwd | Default working directory for commands |
| `env` | map | no | empty | Extra environment variables for commands |
| `menu_columns` | int | no | `2` | Item buttons per row under the message box |
| `page_size` | int | no | `8` | Items per page before pagination |
| `confirm_ttl` | duration | no | `5m` | How long a [confirmation](concepts/confirmation.md) prompt stays valid |
| `enable_run_command` | bool | no | `false` | Show a **`$ >_ Run Command`** button that runs the next message as a shell command. Off by default. Anyone who can use the bot can then run any command on the host, so only turn this on if you trust every allowed user. Putting this key under `telegram` is invalid. |
| `logging` | object | no | built-in default logger | Named loggers (see below) |

??? note "What if I omit `shell`?"

    You can omit it. The bot uses `/bin/bash`. Same for `timeout`, `page_size`,
    and other optional root fields: omit them and defaults apply. You only need
    to set them when you want a non-default value (for example
    `shell: /bin/sh`).

### How much command output you see

Two limits apply, one after the other. `max_output_bytes` is **your** limit and
comes on top of a Telegram limit you cannot change.

**1. Your limit: `max_output_bytes`** (default `524288`, so 512 KB)

While a command runs, the bot keeps at most this much of its output, counted
separately for normal output and error output. Anything past that is dropped,
but the command itself keeps running until it finishes or hits its `timeout`.
When this happens, the result starts with `(output truncated)`.

**2. Telegram's limit: one message holds at most 4096 bytes**

This one is fixed by Telegram. If the result is longer than a single message,
the bot splits it into several messages. Each part is sent as a reply to the
part before it, so they stay together and in order, and the menu buttons appear
on the last part. The split happens on line boundaries whenever possible, so
lines are not cut in half.

If the result is still very long after splitting, the bot stops after 10
messages and the last one ends with a note like
`(output too long; showing first N bytes)`, where `N` is how much of the output
you actually received.

So raising `max_output_bytes` lets the bot keep more output, but you still see
at most about ten messages of it. For output that long, it is usually better to
shorten the command itself (for example `journalctl -u nginx | tail -n 50`) or
write the full output to a file on the server.

### `function_directory` rules

| Situation | Result |
|-----------|--------|
| Key missing | Info log; built-in functions only |
| Key present but empty (`""`) | Info log; built-in functions only |
| Key set to a path that does not exist or is not accessible | Hard error; process stops |
| Path exists but directory is empty | OK |

!!! warning "A wrong path stops the bot"

    If `function_directory` points to a folder that does not exist or cannot be
    read, the program stops with an error instead of starting without your
    custom functions.

## :material-send-circle-outline: `telegram`

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `bot_token` | string | yes | — | Bot token from BotFather |
| `allowed_users` | list of string | yes | — | [Allowed users](concepts/allowed-users.md) |
| `api` | string | no | `https://api.telegram.org` | Bot API base URL |
| `proxy.enabled` | bool | no | `false` | Use proxy for Telegram API |
| `proxy.url` | string | conditional | — | Required when `proxy.enabled` is `true` |
| `insecure` | bool | no | `false` | Skip TLS verify (not recommended) |

Unauthorized users receive a message with their `user_id` and `username` so they can ask an admin for access. This is also how you find your own id the first time — see [Run in CLI → Step 5](installation/download-and-run.md#step-5-find-your-user-id-if-needed).

!!! example "Connect through a proxy"

    ```yaml title="telegram section with a proxy"
    telegram:
      bot_token: "123456789:AAExampleTokenValue"
      allowed_users:
        - "123456789"        # numeric user id
        - "@alice"           # or a username
      proxy:
        enabled: true
        url: "socks5://127.0.0.1:10808"
    ```

To let allowed users type a shell command from Telegram, set this at the **root** of the file (not under `telegram`):

!!! tip "Add one root setting"

    ```yaml title="Enable the Run Command button"
    enable_run_command: true
    ```

## :material-menu: Menu

This section is the field reference. For a guided explanation with examples, see
[Menu](concepts/menu.md). Each [button](concepts/button.md) or
[category](concepts/category.md) node:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | yes | Display name (unique among siblings, case-insensitive) |
| `type` | `category` \| `button` | yes | Node kind |
| `items` | list | yes if `category` | Children; category must have at least one |
| `function` | string | yes if `button` | [Function](concepts/function.md) name |
| `command` | string | yes if `function: command` | Shell command for built-in `command` |
| `path` | string | yes if `function: script` | Script path for built-in `script` |
| `icon` | string | no | Optional emoji prefix |
| `id` | string | no | Optional id for this node. You can omit it |
| `confirm` | bool | no | Ask for [confirmation](concepts/confirmation.md) before run (default `false`) |
| `timeout` | duration | no | Override global timeout |
| `workdir` | string | no | Override working directory |
| `env` | map | no | Extra env for this button |
| `columns` | int | no | Override columns for this category |
| `args` | string | no | Optional args for `script` |
| Any declared parameter name | scalar | as declared by the function | Value passed to the selected function, for example `url`, `host`, `unit`, or `lines` |

On a **button**, any other scalar key is treated as a function parameter.
Its name must match a parameter declared by the selected function. Unknown
parameter names fail [`validate`](cli.md#validate). Values declared as `int` or
`bool` are also checked. Strings, numbers, and booleans can be written directly
as YAML values; numbers do not need quotes.

On a **category**, any key outside the category fields above is an error.
Categories do not run functions, so they cannot have parameter keys.

`command`, `path`, and `args` are shortcut fields that fill parameters with the
same names. Other [parameter](concepts/parameter.md) names are written directly
on the button. Do not place button values inside a nested `params:` map. See
[Functions → Passing values from a button](functions/index.md#passing-values-from-a-button).

## :material-math-log: `logging`

Optional. If omitted, a default console logger on `stderr` at `info` is used.

Named loggers:

!!! example "Write normal logs and an audit file"

    ```yaml title="logging section with an audit file"
    logging:
      logs:
        default:
          level: info
          format: console   # or json
          output:
            - output: stderr
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

Supported outputs: `stdout`, `stderr`, `file`, `discard`.

The `audit` logger shown above records every command run (who, which button,
exit code, duration). See [Audit log](concepts/audit-log.md).

## Related pages

- [Run in CLI](installation/download-and-run.md) — build and run a first config
- [Menu](concepts/menu.md) — the menu tree in depth
- [Functions](functions/index.md) — what `function`, `command`, `path`, and `args` mean
- [CLI](cli.md) — validate and run with your config
