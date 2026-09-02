---
title: Placeholders
description: How to write the run command of a function using placeholders, including optional parts that appear only when a button supplies a value.
icon: material/code-braces
---

# :material-code-braces: Placeholders

The `run` field is a small template. Placeholders written in double curly
braces are replaced with parameter values before the command is executed.

Two patterns cover almost everything:

- `{{.name}}` inserts the value of the parameter called `name`.
- `{{if .name}} ... {{end}}` includes the middle part only when `name` has a
  value.

## Inserting a value

!!! example "One value in the middle of a command"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
    ```

A button with `path: "/var/log/app.log"` runs
`tail -n 200 /var/log/app.log`.

## Optional parts

An empty value makes `{{if .name}}` false, so the whole block disappears. This
is how you add a flag only when it is wanted.

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

**App log** runs `tail -n 200 /var/log/app.log`, and **App errors** runs
`tail -n 200 /var/log/app.log | grep -- ERROR`.

You can also pick between two forms with `{{if .args}} ... {{else}} ... {{end}}`.

## Which names you can use

A placeholder may use any parameter you declared in `params`. Its value is
either what the button supplied or the parameter's `default`.

!!! warning "A name you did not declare fails at run time"

    Only the placeholder syntax is checked by
    [`validate`](../../cli.md#validate), not the names. If `run` mentions
    `{{.uri}}` but no parameter is called `uri`, the config validates fine and
    the button reports an error in the chat when you tap it.

!!! warning "Values are inserted as plain text"

    Nothing is quoted or escaped for you. A value with spaces or shell
    characters becomes part of the command as written, so prefer fixed values on
    buttons and add [`confirm: true`](../../concepts/confirmation.md) to
    anything destructive.

## Related

- [File structure](file-structure.md) — where `run` and `params` live
- [Rules](rules.md) — what the loader accepts
- [Step by step guide](step-by-step.md) — try one end to end
