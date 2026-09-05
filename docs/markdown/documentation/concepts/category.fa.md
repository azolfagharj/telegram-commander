---
title: دسته
description: دکمه‌ها را در زیرمنوها گروه‌بندی کنید. دسته، دکمه‌ها یا دسته‌های بیشتری را زیر items نگه می‌دارد تا منوی بلند روی گوشی کوتاه و کاربردی بماند.
icon: material/folder-outline
---

# :material-folder-outline: دسته

گرهی از منو که به‌جای اجرای کاری، یک زیرمنو را باز می‌کند. دسته به‌جای
[تابع](function.md)، دارای `items` است (یعنی [دکمه‌ها](button.md) یا
دسته‌های بیشتر).

!!! example "دسته‌ای شامل یک دکمه"

    ```yaml title="دستهٔ سیستم"
    - name: System
      type: category
      items:
        - name: Uptime
          type: button
          function: command
          command: "uptime"
    ```

دکمه‌ها و دسته‌ها در کنار هم یک درخت می‌سازند. [منو](menu.md) را ببینید.

## پیکربندی

برای فیلدهای دسته (`type`، `items`، `columns` و موارد دیگر)،
[پیکربندی ← منو](../configuration.md#menu) را ببینید.

## مرتبط

- [دکمه](button.md) — با لمس شدن یک تابع را اجرا می‌کند
- [منو](menu.md) — تودرتوسازی، چیدمان و صفحه‌بندی
