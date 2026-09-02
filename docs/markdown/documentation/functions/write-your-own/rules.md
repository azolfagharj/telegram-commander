---
title: Rules
description: The checks a Telegram Commander function file must pass, covering allowed names, reserved words, unique names, typo rejection and how defaults are used.
icon: material/format-list-checks
---

# :material-format-list-checks: Rules

Every function file is checked when the bot starts and when you run
[`validate`](../../cli.md#validate). If one file is wrong, nothing runs, so
mistakes are caught before your menu goes live.

## Which files are read

- Only files ending in `.yaml` or `.yml` are read. Anything else is ignored.
- Sub-folders of your `function_directory` are read too.
- One file must describe exactly one function.

## Naming

- **`name` is required** and must match `^[A-Za-z0-9._-]+$` — letters, digits,
  `.`, `-`, and `_`. No spaces.
- **Names are unique, ignoring case.** `Deploy` and `deploy` collide, even if
  they live in different files.
- **Reserved names are forbidden.** You cannot name a function `command` or
  `script`; those belong to the
  [built-in functions](../index.md#built-in-functions).

## The command

- **`run` is required.** Write the command with
  [placeholders](placeholders.md) where values go.
- The placeholder syntax is checked when the config is validated, so a broken
  `{{` is reported early.

## Parameters

- **Each parameter needs a `name`**, matching the same character rule as the
  function name.
- **`type` must be** `string`, `int`, or `bool`. Leaving it out means `string`.
  Values and defaults declared as `int` or `bool` are checked.
- A parameter name cannot be a button setting: `name`, `type`, `icon`, `id`,
  `function`, `confirm`, `timeout`, `workdir`, `env`, `columns`, or `items`.
  The names `command`, `path`, and `args` are allowed.
- Every value written on a button must match a parameter declared by that
  button's function.

The function's `run` field can only use placeholders for declared parameters.
An undeclared placeholder is an error.

## Typos are errors

**Unknown keys are rejected.** A typo like `requird:` instead of `required:`
stops validation instead of being silently ignored.

!!! warning "This file does not load"

    ```yaml title="functions/broken.yaml"
    name: broken
    run: "echo {{.args}}"
    params:
      - name: args
        requird: true      # refused: unknown key
    ```

## Required values and defaults

- A **required parameter with no value** makes validation fail with a message
  like `required parameter "args" for function "greet" is missing`.
- A value that is present but empty counts as missing.
- An **optional parameter** falls back to its `default`. With no `default`, it
  becomes empty, which is how optional parts of a command are left out.

!!! example "Required and optional in one file"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n {{.args}} {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        default: "200"
        description: Number of lines
    ```

    A button that sets only `path` runs `tail -n 200 /var/log/app.log`.

## Related

- [File structure](file-structure.md) — every field explained
- [Placeholders](placeholders.md) — writing the `run` command
- [CLI → validate](../../cli.md#validate) — run the checks yourself
