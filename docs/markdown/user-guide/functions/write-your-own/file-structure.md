---
title: File structure
description: Every field of a Telegram Commander function file explained, from name and run to the params list with type, required, default and description.
icon: material/file-tree
---

# :material-file-tree: File structure

One file describes one function. Put the file anywhere inside your
[`function_directory`](../../configuration.md#function_directory-rules),
including a sub-folder, and give it a `.yaml` or `.yml` name. The file name
itself does not matter: the function's name comes from the `name` field inside.

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

## Top-level fields

| Field | Required | Meaning |
|-------|----------|---------|
| `name` | yes | The name buttons use in their `function` field |
| `run` | yes | The command to run, with [placeholders](placeholders.md) for values |
| `params` | no | List of named values the function accepts |

A function with no `params` at all is valid: `run` is then a fixed command.

## Parameter fields

| Field | Required | Default | Meaning |
|-------|----------|---------|---------|
| `name` | yes | — | The name used in `{{.name}}` inside `run` |
| `type` | no | `string` | Value type: `string`, `int`, or `bool` |
| `required` | no | `false` | A button must supply a value |
| `default` | no | empty | Used when no value is supplied |
| `description` | no | empty | A note for yourself, not shown in Telegram |

!!! info "`type` is checked"

    Values and defaults declared as `int` must contain an integer. Values and
    defaults declared as `bool` must contain a valid boolean. Invalid values
    make [`validate`](../../cli.md#validate) fail.

!!! info "Button keys match parameter names"

    Write any parameter directly on a button with the same name. Names such as
    `url`, `host`, and `lines` work alongside `command`, `path`, and `args`. See
    [Passing values from a button](../index.md#passing-values-from-a-button).

## Folder layout

You are free to organise the folder however you like:

!!! example "Sub-folders are read too"

    ```text title="functions/"
    functions/
      disk.yaml
      logs/
        nginx.yaml
        app.yml
    ```

All three files are loaded. Files with any other extension are skipped.

## Related

- [Rules](rules.md) — what the loader refuses
- [Placeholders](placeholders.md) — writing the `run` command
- [Step by step guide](step-by-step.md) — build your first function
