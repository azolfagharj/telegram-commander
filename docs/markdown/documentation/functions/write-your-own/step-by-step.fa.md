---
title: راهنمای گام‌به‌گام
description: نخستین تابع Telegram Commander را از ابتدا بسازید، دکمه‌ای برایش بیفزایید، پیکربندی را اعتبارسنجی کنید و ربات را اجرا کنید تا نتیجه را در گفتگو ببینید.
icon: material/format-list-numbered
---

# :material-format-list-numbered: راهنمای گام‌به‌گام

این راهنما یک تابع کوچک را از پوشه‌ای خالی تا دکمه‌ای عملی پیش می‌برد.
چند دقیقه زمان می‌برد و جز فایل پیکربندی به چیز دیگری نیاز ندارد.

## ۱. معرفی پوشه در پیکربندی

!!! example "افزودن پوشه به پیکربندی"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

مسیر نسبت به فایل پیکربندی است.
[پیکربندی ← قواعد `function_directory`](../../configuration.md#function_directory-rules)
را ببینید.

## ۲. نوشتن فایل تابع

فایل `functions/greet.yaml` را بسازید. پارامتر `args` نام دارد تا دکمه
بتواند آن را پر کند:

!!! example "نخستین تابع شما"

    ```yaml title="functions/greet.yaml"
    name: greet
    run: "echo Hello {{.args}}"
    params:
      - name: args
        required: true
        description: نام برای سلام کردن
    ```

## ۳. افزودن دکمه‌ای که از آن استفاده کند

!!! example "دکمه‌ای برای تابع جدید"

    ```yaml title="دکمهٔ سلام کردن"
    - name: Say hello
      type: button
      function: greet
      args: "world"
    ```

## ۴. اعتبارسنجی و فهرست کردن توابع

!!! example "بررسی پیکربندی و سپس دیدن موارد بارگذاری‌شده"

    ```bash title="اعتبارسنجی و سپس فهرست توابع"
    ./telegram-commander validate --config config.yaml
    ./telegram-commander list-functions --config config.yaml
    ```

باید خطی برای `greet` ببینید. هر خط نام، منشأ تابع و تعداد پارامترهای آن را
نشان می‌دهد. [CLI ← list-functions](../../cli.md#list-functions) را ببینید.

اگر `validate` ایرادی اعلام کرد، پیام آن نام دکمه و مقدار گم‌شده را مشخص
می‌کند؛ [قواعد](rules.md) را بررسی کنید.

## ۵. اجرای ربات و لمس دکمه

!!! example "اجرای ربات در پیش‌زمینه"

    ```bash title="اجرا و مشاهدهٔ خروجی"
    ./telegram-commander run --config config.yaml
    ```

در تلگرام `/start` را بفرستید، **سلام کردن** را بزنید؛ ربات
`echo Hello world` را اجرا می‌کند و `Hello world` را در یک بلوک کد
برمی‌گرداند.

## ۶. استفادهٔ دوباره از تابع

مقدار `args` دکمه را تغییر دهید یا دکمهٔ دومی با مقداری متفاوت بیفزایید.
خود تابع بدون تغییر می‌ماند:

!!! example "دو دکمه، یک تابع"

    ```yaml title="دکمه‌های خوشامدگویی"
    - name: Greet world
      type: button
      function: greet
      args: "world"
    - name: Greet team
      type: button
      function: greet
      args: "team"
    ```

## مرتبط

- [جایگزین‌ها](placeholders.md) — افزودن بخش‌های اختیاری به فرمان
- [قواعد](rules.md) — موارد پذیرفته‌شده توسط بارگذار
- [ساختار فایل](file-structure.md) — شرح همهٔ فیلدها
- [منو](../../concepts/menu.md) — محل قرار دادن دکمه در منو
