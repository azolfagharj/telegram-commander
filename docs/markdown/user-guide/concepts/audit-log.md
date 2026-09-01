---
title: Audit log
description: Keep a record of every command the bot runs, including who tapped the button, which button it was, the exit code, and how long it took.
icon: material/history
---

# :material-history: Audit log

A separate log stream that records every command run: who ran it, which
[button](button.md), the exit code, and how long it took. Configured under
`logging` in your [config file](config-file.md).

This is the record you check later when you want to know who restarted a
service at three in the morning, or whether last night's backup button really
did finish. Because it is its own stream, you can keep it after the ordinary
application log has rotated away.

Each run adds one line, and between them they answer four questions:

| Question | What the line tells you |
|---|---|
| Who? | The Telegram account that tapped the button |
| What? | The button they used |
| Did it work? | The exit code the command returned |
| How long? | The time the command took to finish |

!!! example "Writing the audit log to its own file"

    ```yaml title="logging section of config.yaml"
    logging:
      logs:
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

## Configuration

For the full `logging` schema and supported outputs, see
[Configuration → logging](../configuration.md#logging).

## Related

- [Config file](config-file.md) — where logging is defined
- [Function](function.md) — what runs when a button is tapped
