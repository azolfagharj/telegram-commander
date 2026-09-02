---
title: How the bot connects
description: The bot connects out to Telegram, so you never open a port on your server. Set a proxy if Telegram is blocked on your network.
icon: material/lan-disconnect
---

# :material-lan-disconnect: How the bot connects

The bot connects out to Telegram. You do not open a port on your server.

!!! tip "Nothing to expose to the internet"

    Because the connection only ever goes outward, there is no web panel to
    secure, no inbound firewall rule to add, and no public address for anyone
    to find. The machine can sit behind a home router or a strict firewall and
    the bot still works.

If Telegram is blocked on your network, set a proxy under `telegram` in your
[config file](config-file.md).

## Configuration

For `api`, `proxy`, and `insecure` under Telegram settings, see
[Configuration → telegram](../configuration.md#telegram).

## Related

- [Allowed users](allowed-users.md) — who may use the bot
- [CLI → run](../cli.md#run) — start the bot in the foreground
- [Run as a service](../installation/run-as-a-service.md) — keep the bot running
