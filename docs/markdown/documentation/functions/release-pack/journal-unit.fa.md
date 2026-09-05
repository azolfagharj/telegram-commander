---
title: واحد journal
description: تابع همراه بستهٔ journal-unit گزارش‌های اخیر journalctl را برای یک واحد systemd نشان می‌دهد. واحد و تعداد خط اختیاری را مستقیم روی دکمه تنظیم کنید.
icon: material/text-box-search-outline
---

# :material-text-box-search-outline: واحد journal

`journal-unit` تازه‌ترین گزارش‌های `journalctl` را برای یک واحد systemd
نشان می‌دهد. این تابع یکی از
[توابع همراه بسته](../index.md#custom-functions) است که می‌توانید همان‌گونه
که هست از یک دکمه استفاده کنید.

- اجرا: `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (الزامی): نام واحد، برای نمونه `nginx.service`
- `lines` (اختیاری، پیش‌فرض `50`): تعداد خط‌های نمایش‌داده‌شده

!!! example "فایل تابع"

    ```yaml title="functions/journal-unit.yaml"
    name: journal-unit
    run: "journalctl -u {{.unit}} -n {{.lines}} --no-pager"
    params:
      - name: unit
        type: string
        required: true
        description: نام واحد systemd (برای نمونه nginx.service)
      - name: lines
        type: string
        required: false
        default: "50"
        description: تعداد خط‌های گزارش
    ```

## افزودن دکمه

!!! example "خواندن گزارش‌های اخیر یک سرویس"

    ```yaml title="دکمهٔ گزارش‌های Nginx"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`lines` عددی است، پس نیازی به نقل‌قول ندارد. برای استفاده از مقدار پیش‌فرض
`50` آن را حذف کنید.

!!! example "استفاده از تعداد خط پیش‌فرض"

    ```yaml title="دکمهٔ گزارش‌های SSH"
    - name: SSH logs
      type: button
      function: journal-unit
      unit: "ssh.service"
    ```

## مرتبط

- [جایگزین‌ها](../write-your-own/placeholders.md) — نحوهٔ پر شدن `{{.unit}}`
- [توابع سفارشی](../index.md#custom-functions) — هر پنج نمونهٔ همراه بسته
