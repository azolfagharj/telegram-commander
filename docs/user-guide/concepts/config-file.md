# Config file

A single YAML file that describes everything: your bot token, who may use the
bot, and the menu of [buttons](button.md). You pass it with `--config` on every
CLI command.

## Example

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

## Related

- [Configuration](../configuration.md) — every field, default, and validation rule
- [CLI](../cli.md) — pass the file with `--config`
- [Run in CLI](../installation/download-and-run.md) — create your first config
