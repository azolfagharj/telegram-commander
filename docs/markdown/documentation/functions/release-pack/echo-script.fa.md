---
title: اسکریپت اکو
description: تابع همراه بستهٔ echo-script فایل اسکریپت را از طریق Bash اجرا می‌کند؛ پس اسکریپت به مجوز اجرا نیاز ندارد. path و args را از دکمه می‌گیرد.
icon: material/bash
---

# :material-bash: اسکریپت اکو

`echo-script` اسکریپت را از طریق `bash` اجرا می‌کند؛ بنابراین خود فایل
اسکریپت به مجوز اجرا نیاز ندارد. این تابع یکی از
[توابع همراه بسته](../index.md#custom-functions) است که می‌توانید همان‌گونه
که هست از یک دکمه استفاده کنید.

- اجرا: `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (الزامی): مسیر فایل اسکریپت
- `args` (اختیاری): آرگومان‌های بیشتر

!!! example "فایل تابع"

    ```yaml title="functions/echo-script.yaml"
    name: echo-script
    run: "bash {{.path}}{{if .args}} {{.args}}{{end}}"
    params:
      - name: path
        type: string
        required: true
        description: مسیر فایل اسکریپت
      - name: args
        type: string
        required: false
        description: آرگومان‌های اختیاری اسکریپت
    ```

هر دو نام پارامتر را مستقیم روی دکمه بنویسید.

!!! example "اجرای اسکریپت از طریق Bash"

    ```yaml title="دکمهٔ اجرای پاک‌سازی"
    - name: Run cleanup
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

این دکمه `bash /opt/scripts/cleanup.sh --verbose` را اجرا می‌کند. اگر `args`
را حذف کنید، `bash /opt/scripts/cleanup.sh` اجرا می‌شود؛ زیرا بخش
`{{if .args}}` هنگام خالی بودن مقدار نادیده گرفته می‌شود.

## مرتبط

- [`script`](../built-in/script.md) — اجرای مستقیم اسکریپت اجرایی
- [جایگزین‌ها](../write-your-own/placeholders.md) — نحوهٔ کار `{{if .args}}`
- [توابع سفارشی](../index.md#custom-functions) — هر پنج نمونهٔ همراه بسته
