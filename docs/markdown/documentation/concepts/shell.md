---
title: Shell
description: Every command runs through a shell, /bin/bash by default, so pipes and redirects work. Change it once at the root of your config file.
icon: material/console-line
---

# :material-console-line: Shell

The program that runs your command, by default `/bin/bash`. Commands run as
`bash -c "your command"`, so pipes and redirects work.

You can change the shell globally with the `shell` key in your
[config file](config-file.md).

!!! example "Choosing the shell"

    ```yaml title="Root of config.yaml"
    shell: /bin/bash
    ```

Every [function](function.md) runs its command through this shell.

## Configuration

For the `shell` field and related root options, see
[Configuration → Root fields](../configuration.md#root-fields).

## Related

- [Function](function.md) — produces the command that the shell runs
- [Config file](config-file.md) — where `shell` is set
