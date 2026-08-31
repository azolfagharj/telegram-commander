# Allowed users

The list of Telegram accounts permitted to use the bot, by numeric `user_id` or
`@username`. Everyone else is refused.

Unauthorized users receive a message with their `user_id` and `username` so
they can ask an admin for access. This is also how you find your own id the
first time — see [Run in CLI → Step 5](../installation/download-and-run.md#step-5-find-your-user-id-if-needed).

## Example

```yaml
telegram:
  allowed_users:
    - "123456789"
    - "@alice"
```

## Configuration

For `allowed_users` and other Telegram settings, see
[Configuration → telegram](../configuration.md#telegram).

## Related

- [Config file](config-file.md) — holds the `telegram` section
- [Long polling](long-polling.md) — how the bot receives messages
