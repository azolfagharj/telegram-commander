# Long polling

How the bot receives messages from Telegram: it repeatedly asks Telegram for new
updates. There is nothing to configure and no inbound port to open.

If Telegram is blocked on your network, route it through a proxy under the
`telegram` section of your [config file](config-file.md).

## Configuration

For `api`, `proxy`, and `insecure` under Telegram settings, see
[Configuration → telegram](../configuration.md#telegram).

## Related

- [Allowed users](allowed-users.md) — who may use the bot
- [CLI → run](../cli.md#run) — start the bot in the foreground
- [Run as a service](../installation/run-as-a-service.md) — keep the bot running
