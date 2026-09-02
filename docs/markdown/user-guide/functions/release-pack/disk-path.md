---
title: Disk path
description: The bundled disk-path function shows disk usage with df -h. Its only parameter is path, and it falls back to the root filesystem when you omit it.
icon: material/harddisk
---

# :material-harddisk: Disk path

`disk-path` shows disk usage for a path with `df -h`. It is one of the
[bundled functions](../index.md#custom-functions) you can use from a button as
it is.

- Runs: `df -h {{.path}}`
- `path` (optional, default `/`): filesystem path to check

!!! example "The function file"

    ```yaml title="functions/disk-path.yaml"
    name: disk-path
    run: "df -h {{.path}}"
    params:
      - name: path
        type: string
        required: false
        default: "/"
        description: Filesystem path to check
    ```

Because `path` is optional and has a default, a button can leave it out and
still work.

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

The first button runs `df -h /` and the second runs `df -h /var`.

## Related

- [Rules](../write-your-own/rules.md) — how defaults and required values behave
- [Custom functions](../index.md#custom-functions) — all five bundled examples
- [`command`](../built-in/command.md) — write a full command instead
