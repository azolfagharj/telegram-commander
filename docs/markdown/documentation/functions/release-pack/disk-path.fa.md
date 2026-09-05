---
title: مسیر دیسک
description: تابع همراه بستهٔ disk-path میزان استفاده از دیسک را با df -h نشان می‌دهد. تنها پارامتر آن path است و اگر حذف شود، از فایل‌سیستم ریشه استفاده می‌کند.
icon: material/harddisk
---

# :material-harddisk: مسیر دیسک

`disk-path` میزان استفاده از دیسک برای یک مسیر را با `df -h` نشان می‌دهد.
این تابع یکی از [توابع همراه بسته](../index.md#custom-functions) است که
می‌توانید همان‌گونه که هست از یک دکمه استفاده کنید.

- اجرا: `df -h {{.path}}`
- `path` (اختیاری، پیش‌فرض `/`): مسیر فایل‌سیستم برای بررسی

!!! example "فایل تابع"

    ```yaml title="functions/disk-path.yaml"
    name: disk-path
    run: "df -h {{.path}}"
    params:
      - name: path
        type: string
        required: false
        default: "/"
        description: مسیر فایل‌سیستم برای بررسی
    ```

چون `path` اختیاری و دارای مقدار پیش‌فرض است، دکمه می‌تواند آن را نداشته
باشد و همچنان کار کند.

!!! example "استفاده از مسیر پیش‌فرض یا انتخاب مسیر"

    ```yaml title="دکمه‌های استفاده از دیسک"
    - name: Disk (root)
      type: button
      function: disk-path  # no path: uses "/"
    - name: Disk (var)
      type: button
      function: disk-path
      path: "/var"
    ```

دکمهٔ نخست `df -h /` و دکمهٔ دوم `df -h /var` را اجرا می‌کند.

## مرتبط

- [قواعد](../write-your-own/rules.md) — رفتار مقادیر پیش‌فرض و الزامی
- [توابع سفارشی](../index.md#custom-functions) — هر پنج نمونهٔ همراه بسته
- [`command`](../built-in/command.md) — نوشتن یک فرمان کامل
