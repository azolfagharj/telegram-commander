---
icon: material/file-document-outline
title: Config file
description: One YAML file holds your bot token, the allowed users, and the button menu. Pass it with --config to run, validate, fmt, and list-functions.
---

# :material-file-document-outline: Config file

A single YAML file that describes everything: your bot token, who may use the
bot, and the menu of [buttons](button.md). You pass it with `--config` to the
commands that read it: `run`, `validate`, `fmt`, and `list-functions`. Other
commands, such as `version` and `completion`, do not take a config file.

!!! example "A working config with one button"

    ```yaml title="config.yaml"
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

## Related

- [Configuration](../configuration.md) — every field, default, and validation rule
- [Menu](menu.md) — the button and category tree
- [CLI](../cli.md) — pass the file with `--config`
- [Run in CLI](../installation/download-and-run.md) — create your first config
