---
title: مرجع CLI
description: همهٔ گزینه‌های خط فرمان Telegram Commander، از اجرای ربات تا بررسی و قالب‌بندی فایل پیکربندی و flag مشترک --config.
icon: material/console-line
---

# :material-console-line: مرجع CLI

!!! info "فرمان‌ها از این قالب استفاده می‌کنند"

    ```text title="نحو فرمان"
    telegram-commander <command> [flags]
    ```

`--config` / `-c` برای فرمان‌هایی که پیکربندی را بارگذاری می‌کنند
**الزامی** است. مسیر می‌تواند نسبت به دایرکتوری کاری جاری یا مطلق باشد.
محتوای آن فایل در [پیکربندی](configuration.md) مستند شده است.

اگر تازه شروع کرده‌اید، [اجرا در CLI](installation/download-and-run.md)
این فرمان‌ها را به‌ترتیب استفاده نشان می‌دهد.

## :material-format-list-bulleted-square: فرمان‌ها

### `run` { #run }

ربات را در پیش‌زمینه اجرا می‌کند.

!!! info "اجرای ربات در پیش‌زمینه"

    ```bash title="اجرای ربات"
    telegram-commander run --config /path/to/configfile.yaml
    ```

!!! note "تغییر پیکربندی به راه‌اندازی مجدد نیاز دارد"

    ربات پیکربندی را یک‌بار هنگام شروع می‌خواند. پس از ویرایش فایل، فرایند
    را دوباره راه‌اندازی کنید (برای نمونه
    `systemctl restart telegram-commander`).

### `validate` { #validate }

پیکربندی، توابع و ارجاع‌های دکمه را بدون اجرای ربات اعتبارسنجی می‌کند.

!!! info "بررسی پیکربندی به‌صورت آفلاین یا آنلاین"

    ```bash title="اعتبارسنجی پیکربندی"
    telegram-commander validate --config /path/to/configfile.yaml
    telegram-commander validate --config /path/to/configfile.yaml --online
    ```

!!! note "`--online` به اینترنت نیاز دارد"

    با `--online` بررسی از تلگرام نیز می‌پرسد که توکن ربات کار می‌کند یا
    خیر؛ پس دستگاه باید به تلگرام دسترسی داشته باشد.

### `version`

نسخهٔ برنامه را چاپ می‌کند.

### `fmt`

یک فایل پیکربندی YAML را با قالبی خوانا چاپ می‌کند.

!!! info "چاپ یا ذخیرهٔ YAML قالب‌بندی‌شده"

    ```bash title="قالب‌بندی فایل پیکربندی"
    telegram-commander fmt --config /path/to/configfile.yaml
    telegram-commander fmt --config /path/to/configfile.yaml -w
    ```

### `environ`

متغیرهای محیطی فرایند را چاپ می‌کند (برای عیب‌یابی واحدهای سرویس مفید است).

### `list-functions` { #list-functions }

توابع توکار و سفارشی بارگذاری‌شده را فهرست می‌کند. برای اطمینان از پیدا شدن
فایل‌های تابع سفارشی خود از آن استفاده کنید. [توابع](functions/index.md) را
ببینید.

!!! info "نمایش همهٔ توابع موجود"

    ```bash title="فهرست توابع موجود"
    telegram-commander list-functions --config /path/to/configfile.yaml
    ```

### `completion`

اسکریپت‌های تکمیل خودکار شِل را تولید می‌کند:

!!! info "شِل مورد استفادهٔ خود را انتخاب کنید"

    ```bash title="تولید اسکریپت تکمیل خودکار"
    telegram-commander completion bash
    telegram-commander completion zsh
    telegram-commander completion fish
    telegram-commander completion powershell
    ```

### `manpage`

صفحهٔ راهنما را در خروجی استاندارد می‌نویسد.

## صفحات مرتبط

- [پیکربندی](configuration.md) — فایلی که به `--config` داده می‌شود
- [توابع سفارشی](functions/index.md#custom-functions) — خروجی `list-functions`
- [اجرا به‌صورت سرویس](installation/run-as-a-service.md) — اجرای `run` زیر systemd
