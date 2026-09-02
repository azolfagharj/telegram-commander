---
title: Echo Script
description: The bundled echo-script function runs a script file through Bash, so the script does not need the execute bit. It takes path and args from a button.
icon: material/bash
---

# :material-bash: Echo Script

`echo-script` runs a script through `bash`, so the script file itself does not
need the execute bit. It is one of the
[bundled functions](../index.md#custom-functions) you can use from a button as
it is.

- Runs: `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (required): path to the script file
- `args` (optional): extra arguments

!!! example "The function file"

    ```yaml title="functions/echo-script.yaml"
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

Write both parameter names directly on the button.

!!! example "Run a script through Bash"

    ```yaml title="Run cleanup button"
    - name: Run cleanup
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

That button runs `bash /opt/scripts/cleanup.sh --verbose`. Leave `args` out and
it runs `bash /opt/scripts/cleanup.sh`, because the `{{if .args}}` part is
skipped when the value is empty.

## Related

- [`script`](../built-in/script.md) — run an executable script directly
- [Placeholders](../write-your-own/placeholders.md) — how `{{if .args}}` works
- [Custom functions](../index.md#custom-functions) — all five bundled examples
