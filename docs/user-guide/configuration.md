# Configuration

The [config file](concepts/config-file.md) describes your whole bot: the Telegram
connection, who may use it, the [button](concepts/button.md) menu, and logging.
You pass it to every command with `--config` (see [CLI](cli.md)).

All keys use `lower_snake_case`. **Unknown keys are rejected**, so a typo is an
error you will see immediately when you [validate](cli.md#validate).

**Required** means validation fails if the field is missing or empty after
defaults are applied.  
**Optional** fields may be omitted; the Default column shows what is used then.

New to the project? Start with [Run in CLI](installation/download-and-run.md),
which walks through building a first config. See [Concepts](concepts/config-file.md)
for the vocabulary used below.

## A minimal config

Only `telegram` (with a token and one [allowed user](concepts/allowed-users.md))
and `buttons` are required. Everything else has a default:

```yaml
telegram:
  bot_token: "YOUR_BOT_TOKEN"
  allowed_users:
    - "123456789"

buttons:
  - name: Uptime
    type: button
    function: command
    command: "uptime"
```

The `config-examples/` folder in the release includes both a minimal and a full
example.

## Root fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `telegram` | object | yes | — | Telegram settings (see below) |
| `buttons` | list | yes | — | Menu tree; at least one node |
| `function_directory` | string | no | unset | Custom function YAML directory (see rules below) |
| `shell` | string | no | `/bin/bash` | [Shell](concepts/shell.md) used as `shell -c "<command>"` |
| `timeout` | duration | no | `60s` | Default command timeout |
| `max_output_bytes` | int | no | `524288` | Max captured stdout/stderr bytes |
| `workdir` | string | no | process cwd | Default working directory for commands |
| `env` | map | no | empty | Extra environment variables for commands |
| `buttons_columns` | int | no | `2` | Item buttons per row on the menu message |
| `page_size` | int | no | `8` | Items per page before pagination |
| `confirm_ttl` | duration | no | `5m` | How long a [confirmation](concepts/confirmation.md) prompt stays valid |
| `logging` | object | no | built-in default logger | Named loggers (see below) |

### What if I omit `shell`?

You can omit it. The bot uses `/bin/bash`. Same for `timeout`, `page_size`, and other optional root fields: omit them and defaults apply. You only need to set them when you want a non-default value (for example `shell: /bin/sh`).

### `function_directory` rules

| Situation | Result |
|-----------|--------|
| Key missing | Info log; built-in functions only |
| Key present but empty (`""`) | Info log; built-in functions only |
| Key set to a path that does not exist or is not accessible | Hard error; process stops |
| Path exists but directory is empty | OK |

## `telegram`

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `bot_token` | string | yes | — | Bot token from BotFather |
| `allowed_users` | list of string | yes | — | [Allowed users](concepts/allowed-users.md) |
| `api` | string | no | `https://api.telegram.org` | Bot API base URL |
| `proxy.enabled` | bool | no | `false` | Use proxy for Telegram API |
| `proxy.url` | string | conditional | — | Required when `proxy.enabled` is `true` |
| `insecure` | bool | no | `false` | Skip TLS verify (not recommended) |
| `enable_run_command` | bool | no | `false` | Enable `/run <button name>` |

Unauthorized users receive a message with their `user_id` and `username` so they can ask an admin for access. This is also how you find your own id the first time — see [Run in CLI → Step 5](installation/download-and-run.md#step-5-find-your-user-id-if-needed).

Example:

```yaml
telegram:
  bot_token: "123456789:AAExampleTokenValue"
  allowed_users:
    - "123456789"        # numeric user id
    - "@alice"           # or a username
  proxy:
    enabled: true
    url: "socks5://127.0.0.1:10808"
  enable_run_command: true
```

## Buttons

This section is the field reference. For a guided explanation with examples, see
[Buttons and menus](buttons.md). Each [button](concepts/button.md) or
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
| `id` | string | no | Stable callback id; auto-hashed if omitted |
| `confirm` | bool | no | Ask for [confirmation](concepts/confirmation.md) before run (default `false`) |
| `timeout` | duration | no | Override global timeout |
| `workdir` | string | no | Override working directory |
| `env` | map | no | Extra env for this button |
| `columns` | int | no | Override columns for this category |
| `args` | string | no | Optional args for `script` |
| `params` | map | no | [Parameter](concepts/parameter.md) values for the button's function |

`command`, `path`, and `args` are shortcuts for the matching built-in function
parameters. For any other function, pass values under `params`. See
[Functions → Passing parameters from a button](functions.md#passing-parameters-from-a-button).

## `logging`

Optional. If omitted, a default console logger on `stderr` at `info` is used.

Caddy-style named loggers:

```yaml
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
- [Buttons and menus](buttons.md) — the menu tree in depth
- [Functions](functions.md) — what `function`, `command`, and `params` mean
- [CLI](cli.md) — validate and run with your config
