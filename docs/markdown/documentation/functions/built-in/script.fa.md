---
title: script
description: تابع توکار script یک فایل اسکریپت را روی سرور اجرا می‌کند و آرگومان‌های اختیاری را پس از مسیر می‌افزاید. اسکریپت باید اجرایی باشد.
icon: material/script-text
---

# :material-script-text: `script`

`script` یک [تابع توکار](../index.md#built-in-functions) است. این تابع یک
فایل اسکریپت را با آرگومان‌های اختیاری اجرا می‌کند.

| پارامتر | الزامی | پیش‌فرض | معنی |
|-----------|----------|---------|---------|
| `path` | بله | — | مسیر اسکریپت |
| `args` | خیر | خالی | آرگومان‌هایی که پس از مسیر داده می‌شوند |

!!! example "اجرای اسکریپت با آرگومان"

    ```yaml title="دکمهٔ گزارش شبانه"
    - name: Nightly report
      type: button
      function: script
      path: "/usr/local/bin/report.sh"
      args: "--today"
    ```

این دکمه `/usr/local/bin/report.sh --today` را اجرا می‌کند. اگر `args` را
حذف کنید، فقط `/usr/local/bin/report.sh` اجرا می‌شود.

## اسکریپت باید اجرایی باشد

مسیر مستقیماً اجرا می‌شود؛ بنابراین فایل باید مجوز اجرا داشته باشد:

!!! tip "یک‌بار اسکریپت را اجرایی کنید"

    ```bash title="دادن مجوز اجرای فایل"
    chmod +x /usr/local/bin/report.sh
    ```

اگر نمی‌توانید فایل را تغییر دهید، آن را از طریق یک مفسر اجرا کنید. نمونهٔ
[`echo-script`](../release-pack/echo-script.md) در بستهٔ انتشار دقیقاً با
فراخوانی نخست `bash` همین کار را انجام می‌دهد.

## مرتبط

- [`command`](command.md) — اجرای فرمان درون‌خطی به‌جای فایل
- [اسکریپت اکو](../release-pack/echo-script.md) — اجرای اسکریپت از طریق Bash
- [توابع توکار](../index.md#built-in-functions) — هر دو تابع توکار
- [Shell](../../concepts/shell.md) — نحوهٔ اجرای فرمان‌ها
