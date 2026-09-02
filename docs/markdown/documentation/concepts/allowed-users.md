---
title: Allowed users
description: Choose which Telegram accounts can open the bot menu. Add people by numeric user id or username, and see what everyone else gets instead.
icon: material/shield-lock
---

# :material-shield-lock: Allowed users

The list of Telegram accounts permitted to use the bot, by numeric `user_id` or
`@username`. Everyone else is refused.

Unauthorized users receive a message with their `user_id` and `username` so
they can ask an admin for access. This is also how you find your own id the
first time — see [Run in CLI → Step 5](../installation/download-and-run.md#step-5-find-your-user-id-if-needed).

That message doubles as the way you add someone. Ask the new person to send
the bot any message, have them copy the id it replies with, add that id to
`allowed_users`, and restart the bot. Next time they write, they get the menu
instead of the refusal.

!!! warning "Everyone on the list shares the same menu"

    There are no per-user permissions. Anyone you add can tap every button you
    defined, so keep the list to people you trust with the machine itself.

!!! example "Two accounts allowed to use the bot"

    ```yaml title="telegram section of config.yaml"
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
- [How the bot connects](long-polling.md) — no inbound port to open
