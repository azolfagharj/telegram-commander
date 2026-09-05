---
title: اجرا در CLI
description: ربات را گام‌به‌گام از ترمینال راه‌اندازی کنید؛ از دریافت نسخه و نوشتن پیکربندی کوچک تا لمس نخستین دکمه.
icon: material/console
---

# :material-console: اجرا در CLI

این صفحه شما را از نقطهٔ شروع تا یک ربات در حال اجرا پیش می‌برد. نیازی به
تجربهٔ قبلی با پروژه ندارید. اگر واژه‌ای نامفهوم بود، صفحات
[مفاهیم](../concepts/config-file.md) را ببینید.

## :material-clipboard-check-outline: پیش از شروع

به دو مورد از تلگرام نیاز دارید:

1. **توکن ربات.** گفتگویی با [@BotFather](https://t.me/BotFather) باز کنید،
   `/newbot` را بفرستید، مراحل را دنبال کنید و توکنی را که می‌دهد کپی کنید.
   ظاهر آن شبیه `123456789:AAExampleTokenValue` است.
2. **شناسهٔ عددی کاربر شما.** این یک عدد است، نه `@username` شما. اگر آن را
   نمی‌دانید نگران نباشید؛ ربات نخستین بار که به آن پیام بدهید شناسه را
   اعلام می‌کند ([گام ۵](#step-5-find-your-user-id-if-needed) را ببینید).

## :material-download: گام ۱: دریافت

بایگانی نسخه را دریافت و استخراج کنید.

!!! example "دریافت و باز کردن پوشهٔ نسخه"

    ```bash title="دریافت و استخراج نسخه"
    wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
    tar -xzf telegram-commander.tar.gz
    cd telegram-commander
    ```

درون پوشه این موارد را می‌بینید:

- `telegram-commander-linux-amd64` و `telegram-commander-linux-arm64` — برنامه، یکی برای هر نوع CPU
- `config-examples/` — [فایل‌های پیکربندی](../concepts/config-file.md) آماده (به [پیکربندی](../configuration.md) مراجعه کنید)
- `functions/` — نمونهٔ [توابع](../concepts/function.md) سفارشی (به [توابع سفارشی](../functions/index.md#custom-functions) مراجعه کنید)

## :material-chip: گام ۲: انتخاب فایل اجرایی

!!! info "کدام فایل برای دستگاه شماست؟"

    بیشتر سرورها و رایانه‌ها `amd64` هستند (که `x86_64` نیز نامیده می‌شود).
    بردهای کوچک ARM و برخی ماشین‌های مجازی ابری `arm64` هستند.

    اگر مطمئن نیستید، `uname -m` را اجرا کنید: `x86_64` یعنی amd64 و
    `aarch64` یعنی arm64.

=== ":fontawesome-brands-linux: AMD64"

    ```bash title="نگه داشتن برنامهٔ amd64"
    mv telegram-commander-linux-amd64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-arm64
    ```

=== ":fontawesome-brands-linux: ARM64"

    ```bash title="نگه داشتن برنامهٔ arm64"
    mv telegram-commander-linux-arm64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-amd64
    ```

اکنون یک برنامه به نام `telegram-commander` دارید.

## :material-file-cog-outline: گام ۳: ساخت پیکربندی

نمونهٔ کمینه را در یک فایل کاری کپی کنید:

!!! example "ساخت پیکربندی قابل ویرایش"

    ```bash title="کپی کردن پیکربندی نمونه"
    cp config-examples/config.minimal.yaml ./config.yaml
    ```

`config.yaml` را باز کنید و دو جایگزین را عوض کنید:

- `YOUR_BOT_TOKEN` — توکن دریافتی از BotFather
- `YOUR_USER_ID` — شناسهٔ عددی شما (یا فعلاً آن را رها کنید و گام ۵ را ببینید)

برای آشنایی با همهٔ تنظیمات، [پیکربندی](../configuration.md) را بخوانید.

## :material-file-check-outline: گام ۴: اعتبارسنجی

همیشه پیش از اجرا پیکربندی را بررسی کنید. این کار غلط‌های تایپی و اشتباه‌ها
را بدون شروع ربات پیدا می‌کند.

!!! success "بررسی کار کردن پیکربندی"

    ```bash title="اعتبارسنجی پیکربندی"
    ./telegram-commander validate --config config.yaml
    ```

اگر `Valid configuration` چاپ شد، همه‌چیز آماده است. در غیر این صورت دقیقاً
مشکل و محل آن فهرست می‌شود. برای جزئیات
[صفحهٔ CLI](../cli.md#validate) را ببینید.

## :material-account-search: گام ۵: یافتن شناسهٔ کاربر (در صورت نیاز) { #step-5-find-your-user-id-if-needed }

اگر شناسهٔ کاربر خود را نمی‌دانید، فقط توکن را در `config.yaml` تنظیم کنید،
فعلاً هر عددی در `allowed_users` بگذارید و سپس ربات را اجرا کنید:

!!! info "یک‌بار اجرا برای دیدن شناسهٔ کاربر"

    ```bash title="اجرای ربات برای یافتن شناسه"
    ./telegram-commander run --config config.yaml
    ```

تلگرام را باز کنید، ربات خود را پیدا کنید و هر پیامی برایش بفرستید. چون هنوز
در [کاربران مجاز](../concepts/allowed-users.md) نیستید، ربات با `user_id` و
`username` شما پاسخ می‌دهد. آن شناسه را در `allowed_users` کپی کنید، ربات
را با `Ctrl+C` متوقف کنید و دوباره اجرا کنید.

این رفتار بخشی از کنترل دسترسی است؛
[پیکربندی ← telegram](../configuration.md#telegram) را ببینید.

## :material-play-circle-outline: گام ۶: اجرا

!!! example "شروع ربات در ترمینال"

    ```bash title="شروع ربات"
    ./telegram-commander run --config config.yaml
    ```

ربات خود را در تلگرام باز کنید و `/start` را بفرستید. باید منوی خود را
ببینید. برای اجرای فرمان، یک [دکمه](../concepts/button.md) را لمس کنید.

!!! success "ربات شما فعال است"

    منویی که در `config.yaml` توصیف کرده‌اید اکنون در گفتگوی شماست و هر
    لمس، فرمان همان دکمه را روی این دستگاه اجرا می‌کند.

برای روشن ماندن ربات پس از خروج از سرور، آن را به‌صورت سرویس تنظیم کنید.
[اجرا به‌صورت سرویس](run-as-a-service.md) را ببینید.

## :material-map-marker-path: گام بعدی

- افزودن [دکمه‌ها](../concepts/button.md) و [دسته‌ها](../concepts/category.md): [منو](../concepts/menu.md)
- درک چیزی که واقعاً اجرا می‌شود: [توابع](../functions/index.md)
- دیدن همهٔ گزینه‌های خط فرمان: [CLI](../cli.md)
