---
icon: material/file-document-outline
title: فایل پیکربندی
description: یک فایل YAML توکن ربات، کاربران مجاز و منوی دکمه‌ها را نگه می‌دارد. آن را با --config به run، validate، fmt و list-functions بدهید.
---

# :material-file-document-outline: فایل پیکربندی

یک فایل YAML همه‌چیز را توصیف می‌کند: توکن ربات، افرادی که می‌توانند از
ربات استفاده کنند و منوی [دکمه‌ها](button.md). این فایل را با `--config` به
فرمان‌هایی می‌دهید که آن را می‌خوانند: `run`، `validate`، `fmt` و
`list-functions`. فرمان‌های دیگر، مانند `version` و `completion`، فایل
پیکربندی نمی‌گیرند.

!!! example "یک پیکربندی عملی با یک دکمه"

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

## مرتبط

- [پیکربندی](../configuration.md) — همهٔ فیلدها، مقادیر پیش‌فرض و قواعد اعتبارسنجی
- [منو](menu.md) — درخت دکمه‌ها و دسته‌ها
- [CLI](../cli.md) — دادن فایل با `--config`
- [اجرا در CLI](../installation/download-and-run.md) — ساخت نخستین پیکربندی
