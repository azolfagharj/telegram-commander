---
title: Delivery retries
description: A short network drop no longer loses the result. The bot retries the send with a growing wait, honors Telegram's own rate-limit wait, and keeps per-chat order.
icon: material/refresh
---

# :material-refresh: Delivery retries

If the connection to Telegram drops for a moment while the bot is sending a
result, the message is not lost. The bot tries again on its own, waiting a
little longer each time, until the send goes through.

When Telegram itself asks the bot to slow down, the bot waits exactly as long
as Telegram asks before trying again, instead of guessing its own wait time.

Messages to the same chat still arrive in the order you triggered them, even
if you tap several buttons while the connection is having trouble.

If the drop lasts too long, the bot eventually stops trying for that one
result. Nothing appears in the chat and nothing is left half-sent — the
attempt is only noted in the server log, never shown to you.

!!! tip "Nothing to set for a normal setup"

    The defaults handle a short blip in connectivity on their own. Change how
    long the bot keeps trying with `delivery_retry_backoff`,
    `delivery_retry_backoff_max`, and `delivery_retry_ttl` in your
    [config file](config-file.md).

## Configuration

For `delivery_retry_backoff`, `delivery_retry_backoff_max`, and
`delivery_retry_ttl`, see [Configuration](../configuration.md).

## Related

- [Confirmation](confirmation.md) — another moment where the bot waits before acting
- [How the bot connects](long-polling.md) — the outbound connection this protects
- [Control output](../output-control.md) — how a result reaches your chat
