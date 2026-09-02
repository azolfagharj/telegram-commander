---
title: Journal Unit
description: The bundled journal-unit function shows recent journalctl logs for a systemd unit. Set the unit and optional line count directly on the button.
icon: material/text-box-search-outline
---

# :material-text-box-search-outline: Journal Unit

`journal-unit` shows the most recent `journalctl` logs for a systemd unit. It is
one of the [bundled functions](../index.md#custom-functions) you can use from a
button as it is.

- Runs: `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (required): unit name, for example `nginx.service`
- `lines` (optional, default `50`): how many lines to show

!!! example "The function file"

    ```yaml title="functions/journal-unit.yaml"
    name: journal-unit
    run: "journalctl -u {{.unit}} -n {{.lines}} --no-pager"
    params:
      - name: unit
        type: string
        required: true
        description: Systemd unit name (for example nginx.service)
      - name: lines
        type: string
        required: false
        default: "50"
        description: Number of log lines
    ```

## Add a button

!!! example "Read recent logs of one service"

    ```yaml title="Nginx logs button"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`lines` is numeric, so it does not need quotes. Leave it out to use the default
value `50`.

!!! example "Use the default line count"

    ```yaml title="SSH logs button"
    - name: SSH logs
      type: button
      function: journal-unit
      unit: "ssh.service"
    ```

## Related

- [Placeholders](../write-your-own/placeholders.md) — how `{{.unit}}` is filled in
- [Custom functions](../index.md#custom-functions) — all five bundled examples
