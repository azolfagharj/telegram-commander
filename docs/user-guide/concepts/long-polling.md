# How the bot connects

The bot connects out to Telegram. You do not open a port on your server.

If Telegram is blocked on your network, set a proxy under `telegram` in your
[config file](config-file.md).

## Configuration

For `api`, `proxy`, and `insecure` under Telegram settings, see
[Configuration → telegram](../configuration.md#telegram).

## Related

- [Allowed users](allowed-users.md) — who may use the bot
- [CLI → run](../cli.md#run) — start the bot in the foreground
- [Run as a service](../installation/run-as-a-service.md) — keep the bot running
