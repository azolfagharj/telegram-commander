---
title: URL با curl
description: تابع همراه بستهٔ curl-url یک URL را با curl دریافت می‌کند. URL را با کلید url مستقیماً به دکمه بیفزایید.
icon: material/web
---

# :material-web: URL با curl

`curl-url` یک URL را با `curl` دریافت می‌کند. در خطاهای HTTP ناموفق می‌شود
(`-f`) و پس از ۳۰ ثانیه دست می‌کشد. این تابع یکی از
[توابع همراه بسته](../index.md#custom-functions) است که می‌توانید همان‌گونه
که هست از یک دکمه استفاده کنید.

- اجرا: `curl -fsSL --max-time 30 {{.url}}`
- `url` (الزامی): URL مورد درخواست

!!! example "فایل تابع"

    ```yaml title="functions/curl-url.yaml"
    name: curl-url
    run: "curl -fsSL --max-time 30 {{.url}}"
    params:
      - name: url
        type: string
        required: true
        description: URL مورد درخواست
    ```

## افزودن دکمه

`url` را مستقیماً روی دکمه بنویسید. نام آن با پارامتری که تابع تعریف کرده
یکسان است.

!!! example "بررسی یک endpoint"

    ```yaml title="دکمهٔ بررسی API"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

دکمه `curl -fsSL --max-time 30 https://example.com/health` را اجرا می‌کند.
اگر `url` وجود نداشته باشد، [`validate`](../../cli.md#validate) خطا می‌دهد.

## مرتبط

- [پارامترها](../../concepts/parameter.md) — نحوهٔ بررسی مقادیر دکمه
- [توابع سفارشی](../index.md#custom-functions) — هر پنج نمونهٔ همراه بسته
