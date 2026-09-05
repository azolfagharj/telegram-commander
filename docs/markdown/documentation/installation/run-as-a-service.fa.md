---
title: اجرا به‌صورت سرویس
description: ربات را با systemd در پس‌زمینه روشن نگه دارید تا هنگام راه‌اندازی سیستم اجرا شود و پس از خرابی بازگردد؛ با کاربر root یا کاربری عادی.
icon: material/server
---

# :material-server: اجرا به‌صورت سرویس

این صفحه اجرای Telegram Commander به‌صورت سرویس systemd را توضیح می‌دهد.
برای نخستین اجرای گام‌به‌گام در ترمینال،
[اجرا در CLI](download-and-run.md) را ببینید.

## :material-cog-outline: systemd (root)

فایل واحد را خودتان بسازید، برای نمونه
`/etc/systemd/system/telegram-commander.service`:

!!! warning "این نمونه همهٔ دکمه‌ها را با دسترسی root اجرا می‌کند"

    ```ini title="/etc/systemd/system/telegram-commander.service"
    [Unit]
    Description=telegram-commander Telegram bot
    After=network-online.target
    Wants=network-online.target

    [Service]
    Type=simple
    ExecStart=/path/to/telegram-commander run --config /path/to/configfile.yaml
    Restart=on-failure
    RestartSec=5

    [Install]
    WantedBy=multi-user.target
    ```

    چون خط `User=` وجود ندارد، سرویس با root شروع می‌شود و همهٔ فرمان‌های
    منو روی دستگاه دسترسی کامل دارند. اگر دکمه‌ها به این سطح از دسترسی نیاز
    ندارند، `User=someone` را به بخش `[Service]` بیفزایید.

مسیرهای جایگزین را عوض کنید، سپس:

!!! example "بارگذاری و شروع سرویس"

    ```bash title="فعال‌سازی، شروع و مشاهدهٔ سرویس"
    sudo systemctl daemon-reload
    sudo systemctl enable --now telegram-commander
    sudo systemctl status telegram-commander
    sudo journalctl -u telegram-commander -f
    ```

!!! info "تغییر پیکربندی به راه‌اندازی مجدد نیاز دارد"

    سرویس ربات را روشن نگه می‌دارد، اما
    [فایل پیکربندی](../concepts/config-file.md) را یک‌بار هنگام شروع
    می‌خواند. پس از ویرایش آن،
    `sudo systemctl restart telegram-commander` را اجرا کنید.

## صفحات مرتبط

- [اجرا در CLI](download-and-run.md) — توضیح نخستین اجرا
- [پیکربندی](../configuration.md) — فایل پیکربندی
- [CLI](../cli.md) — `run`، `validate` و موارد دیگر
