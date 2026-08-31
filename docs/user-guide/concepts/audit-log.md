# Audit log

A separate log stream that records every command run: who ran it, which
[button](button.md), the exit code, and how long it took. Configured under
`logging` in your [config file](config-file.md).

## Example

```yaml
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
