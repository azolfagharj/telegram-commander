---
title: پینگ میزبان
description: تابع همراه بستهٔ ping-host چند بار یک میزبان را پینگ می‌کند. host و count اختیاری را مستقیم روی دکمه تنظیم کنید.
icon: material/access-point-network
---

# :material-access-point-network: پینگ میزبان

`ping-host` میزبانی را چند بار پینگ می‌کند و نتیجه را برمی‌گرداند. این تابع
یکی از [توابع همراه بسته](../index.md#custom-functions) است که می‌توانید
همان‌گونه که هست از یک دکمه استفاده کنید.

- اجرا: `ping -c {{.count}} {{.host}}`
- `host` (الزامی): نام میزبان یا IP
- `count` (اختیاری، پیش‌فرض `4`): تعداد بسته‌ها

!!! example "فایل تابع"

    ```yaml title="functions/ping-host.yaml"
    name: ping-host
    run: "ping -c {{.count}} {{.host}}"
    params:
      - name: host
        type: string
        required: true
        description: نام میزبان یا IP
      - name: count
        type: string
        required: false
        default: "4"
        description: تعداد بسته‌های ping
    ```

## افزودن دکمه

!!! example "سه بار پینگ کردن یک میزبان ثابت"

    ```yaml title="دکمهٔ پینگ دروازه"
    - name: Ping gateway
      type: button
      function: ping-host
      host: "192.168.1.1"
      count: 3
    ```

`count` عددی است، پس نیازی به نقل‌قول ندارد. همچنین می‌توانید آن را حذف
کنید تا مقدار پیش‌فرض `4` استفاده شود:

!!! example "استفاده از تعداد پیش‌فرض"

    ```yaml title="دکمهٔ پینگ DNS"
    - name: Ping DNS
      type: button
      function: ping-host
      host: "1.1.1.1"
    ```

## مرتبط

- [قواعد](../write-your-own/rules.md) — رفتار مقادیر پیش‌فرض و الزامی
- [توابع سفارشی](../index.md#custom-functions) — هر پنج نمونهٔ همراه بسته
