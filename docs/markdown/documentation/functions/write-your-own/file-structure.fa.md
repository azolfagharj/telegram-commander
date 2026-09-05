---
title: ساختار فایل
description: شرح همهٔ فیلدهای فایل تابع Telegram Commander، از name و run تا فهرست params شامل type، required، default و description.
icon: material/file-tree
---

# :material-file-tree: ساختار فایل

هر فایل یک تابع را توصیف می‌کند. فایل را هرجای
[`function_directory`](../../configuration.md#function_directory-rules)،
حتی در زیردایرکتوری، بگذارید و پسوند `.yaml` یا `.yml` به آن بدهید. نام
خود فایل اهمیتی ندارد؛ نام تابع از فیلد `name` درون آن می‌آید.

!!! example "تعریف یک تابع سفارشی در هر فایل"

    ```yaml title="فایل تابع سفارشی"
    name: my-function  # required; must be unique and not reserved
    run: "echo {{.args}}"  # required; command to run
    params:  # optional parameter list
      - name: args  # required for each parameter
        type: string  # optional: string (default), int, or bool
        required: true  # optional; default is false
        default: ""  # optional value used when missing
        description: یک متن  # optional note for yourself
    ```

## فیلدهای سطح بالا

| فیلد | الزامی | معنی |
|-------|----------|---------|
| `name` | بله | نامی که دکمه‌ها در فیلد `function` خود استفاده می‌کنند |
| `run` | بله | فرمان اجرا، همراه [جایگزین‌ها](placeholders.md) برای مقادیر |
| `params` | خیر | فهرست مقادیر نام‌داری که تابع می‌پذیرد |

تابعی که اصلاً `params` ندارد معتبر است؛ در این حالت `run` فرمانی ثابت است.

## فیلدهای پارامتر

| فیلد | الزامی | پیش‌فرض | معنی |
|-------|----------|---------|---------|
| `name` | بله | — | نام استفاده‌شده در `{{.name}}` داخل `run` |
| `type` | خیر | `string` | نوع مقدار: `string`، `int` یا `bool` |
| `required` | خیر | `false` | دکمه باید مقداری ارائه کند |
| `default` | خیر | خالی | وقتی مقداری ارائه نشود استفاده می‌شود |
| `description` | خیر | خالی | یادداشتی برای خودتان که در تلگرام نمایش داده نمی‌شود |

!!! info "`type` بررسی می‌شود"

    مقادیر و پیش‌فرض‌های تعریف‌شده با `int` باید عدد صحیح باشند. مقادیر و
    پیش‌فرض‌های تعریف‌شده با `bool` باید مقدار بولی معتبر باشند. مقادیر
    نامعتبر باعث شکست [`validate`](../../cli.md#validate) می‌شوند.

!!! info "کلیدهای دکمه با نام پارامترها یکسان‌اند"

    هر پارامتر را با همان نام مستقیماً روی دکمه بنویسید. نام‌هایی مانند
    `url`، `host` و `lines` در کنار `command`، `path` و `args` کار می‌کنند.
    [فرستادن مقادیر از دکمه](../index.md#passing-values-from-a-button) را
    ببینید.

## چیدمان پوشه

می‌توانید پوشه را هرطور می‌خواهید سازمان‌دهی کنید:

!!! example "زیردایرکتوری‌ها نیز خوانده می‌شوند"

    ```text title="functions/"
    functions/
      disk.yaml
      logs/
        nginx.yaml
        app.yml
    ```

هر سه فایل بارگذاری می‌شوند. فایل‌هایی با پسوندهای دیگر نادیده گرفته می‌شوند.

## مرتبط

- [قواعد](rules.md) — مواردی که بارگذار نمی‌پذیرد
- [جایگزین‌ها](placeholders.md) — نوشتن فرمان `run`
- [راهنمای گام‌به‌گام](step-by-step.md) — ساخت نخستین تابع
