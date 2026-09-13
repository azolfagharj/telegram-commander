---
icon: material/file-cog-outline
title: پیکربندی
description: همهٔ تنظیمات Telegram Commander همراه نوع، پیش‌فرض و معنی آن‌ها؛ telegram، menu، function_directory، timeout، محدودیت خروجی و logging.
---

# :material-file-cog-outline: پیکربندی

[فایل پیکربندی](concepts/config-file.md) کل ربات را توصیف می‌کند: اتصال
تلگرام، کاربران مجاز، منوی [دکمه‌ها](concepts/button.md) و گزارش‌گیری. آن را
با `--config` به فرمان‌های `run`، `validate`، `fmt` و `list-functions`
می‌دهید ([CLI](cli.md) را ببینید).

همهٔ کلیدها از `lower_snake_case` استفاده می‌کنند. **کلیدهای ناشناخته رد
می‌شوند**؛ بنابراین غلط تایپی خطایی است که هنگام
[اعتبارسنجی](cli.md#validate) فوراً می‌بینید.

**الزامی** یعنی اگر فیلد پس از اعمال پیش‌فرض‌ها وجود نداشته یا خالی باشد،
اعتبارسنجی ناموفق است.  
فیلدهای **اختیاری** را می‌توان حذف کرد؛ ستون پیش‌فرض مقدار استفاده‌شده را
نشان می‌دهد.

تازه شروع کرده‌اید؟ از [اجرا در CLI](installation/download-and-run.md)
آغاز کنید که ساخت نخستین پیکربندی را گام‌به‌گام توضیح می‌دهد. برای واژگان
زیر، [مفاهیم](concepts/config-file.md) را ببینید.

## :material-rocket-launch-outline: پیکربندی کمینه { #a-minimal-config }

فقط `telegram` (با توکن و یک [کاربر مجاز](concepts/allowed-users.md)) و
`menu` الزامی‌اند. بقیه مقدار پیش‌فرض دارند:

!!! example "شروع با یک کاربر مجاز و سه دکمهٔ وضعیت"

    ```yaml title="config.yaml (کمینه)"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: System Status
        type: category
        icon: "🖥️"
        items:
          - name: Uptime
            type: button
            icon: "🕒"
            function: command
            command: "uptime"
          - name: Memory
            type: button
            icon: "🧮"
            function: command
            command: "free -h"
          - name: Disk Space
            type: button
            icon: "💾"
            function: command
            command: "df -h /"
    ```

پوشهٔ `config-examples/` در نسخه، نمونهٔ کمینه و کامل را در خود دارد.

## :material-card-bulleted-outline: فیلدهای ریشه { #root-fields }

| فیلد | نوع | الزامی | پیش‌فرض | توضیح |
|-------|------|----------|---------|-------------|
| `telegram` | object | بله | — | تنظیمات تلگرام (پایین را ببینید) |
| `menu` | list | بله | — | درخت منو؛ دست‌کم یک گره |
| `function_directory` | string | خیر | تنظیم‌نشده | دایرکتوری YAML توابع سفارشی (قواعد پایین) |
| `shell` | string | خیر | `/bin/bash` | [Shell](concepts/shell.md) به‌شکل `shell -c "<command>"` |
| `timeout` | duration | خیر | `60s` | timeout پیش‌فرض فرمان |
| `max_output_bytes` | int | خیر | `524288` | بیشترین خروجی نگه‌داری‌شده برای هر فرمان ([مقدار خروجی فرمان](#how-much-command-output-you-see)) |
| `output` | `auto` \| `text` \| `file` | خیر | `auto` | نحوهٔ تحویل نتیجهٔ فرمان ([مقدار خروجی فرمان](#how-much-command-output-you-see)) |
| `max_output_messages` | int | خیر | `2` | فقط در ریشه. در حالت `auto`، اگر نتیجه به بیش از این تعداد پیام نیاز داشته باشد یک فایل `.txt` فرستاده می‌شود (۱ تا ۱۰). ننویسید تا همان `۲` بماند |
| `workdir` | string | خیر | cwd فرایند | دایرکتوری کاری پیش‌فرض فرمان‌ها |
| `env` | map | خیر | خالی | متغیرهای محیطی بیشتر برای فرمان‌ها |
| `menu_columns` | int | خیر | `2` | تعداد دکمه‌های گزینه در هر ردیف زیر کادر پیام |
| `page_size` | int | خیر | `8` | تعداد گزینه‌های هر صفحه پیش از صفحه‌بندی |
| `confirm_ttl` | duration | خیر | `5m` | مدت اعتبار پیام [تأیید](concepts/confirmation.md) |
| `enable_run_command` | bool | خیر | `false` | نمایش دکمهٔ **`$ >_ Run Command`** که پیام بعدی را به‌عنوان فرمان شِل اجرا می‌کند. پیش‌فرض خاموش است. هر کاربر ربات می‌تواند هر فرمانی روی میزبان اجرا کند؛ فقط اگر به همهٔ کاربران مجاز اعتماد دارید آن را روشن کنید. قرار دادن این کلید زیر `telegram` نامعتبر است. |
| `logging` | object | خیر | گزارش‌گیر پیش‌فرض توکار | گزارش‌گیرهای نام‌دار (پایین را ببینید) |

??? note "اگر `shell` را حذف کنم چه می‌شود؟"

    می‌توانید آن را حذف کنید؛ ربات از `/bin/bash` استفاده می‌کند. برای
    `timeout`، `page_size` و دیگر فیلدهای اختیاری ریشه نیز پیش‌فرض‌ها اعمال
    می‌شوند. فقط برای مقدار غیراستاندارد آن‌ها را تنظیم کنید، مانند
    `shell: /bin/sh`.

### مقدار خروجی فرمان که می‌بینید { #how-much-command-output-you-see }

دو محدودیت به‌ترتیب اعمال می‌شوند. `max_output_bytes` محدودیت **شما** است و
روی محدودیت تلگرام که عوض‌کردنی نیست می‌نشیند.

**۱. محدودیت شما: `max_output_bytes`** (پیش‌فرض `524288`، یعنی ۵۱۲ کیلوبایت)

هنگام اجرای فرمان، ربات حداکثر این مقدار از خروجی عادی و خطا را جداگانه نگه
می‌دارد. بیشتر از آن دور ریخته می‌شود، ولی خود فرمان تا پایان یا رسیدن به
`timeout` ادامه می‌یابد. در این حالت نتیجه با `(output truncated)` شروع
می‌شود.

**۲. نحوهٔ تحویل نتیجه: `output`** (پیش‌فرض `auto`)

تلگرام حداکثر ۴۰۹۶ بایت در هر پیام جای می‌دهد. ربات می‌تواند نتیجهٔ بلند را به
چند پیام ریپلای بشکند یا کل خروجی را به‌صورت فایل `.txt` بفرستد. حالت را با
`output` در ریشهٔ پیکربندی انتخاب کنید یا روی یک دکمه عوض کنید:

| مقدار | رفتار |
|-------|--------|
| `auto` | تا وقتی نتیجه در `max_output_messages` پیام جا شود (پیش‌فرض `۲`) متن می‌فرستد؛ اگر بیشتر لازم باشد یک فایل `.txt` می‌فرستد. |
| `text` | همیشه به پیام متنی می‌شکند. همچنان بعد از ۱۰ پیام می‌ایستد و بریده‌شدن را اعلام می‌کند. |
| `file` | همیشه یک فایل `.txt` با کپشن کوتاه (نام دکمه، کد خروج، مدت) می‌فرستد. |

در حالت‌های `auto` و `file` فایل شامل سربرگ و stdout و stderr است؛ پس بالا
بردن `max_output_bytes` واقعاً خروجی بیشتری به شما می‌رساند. اگر ارسال فایل
شکست بخورد، ربات به مسیر پیام متنی برمی‌گردد.

لازم نیست `output` یا `max_output_messages` را بنویسید. اگر ننویسید، ربات
`auto` و `۲` را به کار می‌برد. کانفیگی که الان کار می‌کند همان‌طور می‌ماند.
`max_output_messages` فقط در ریشه است. روی دکمه `output` را نگذارید تا مقدار
ریشه استفاده شود؛ فقط وقتی بنویسید که همان دکمه همیشه فایل یا همیشه متن باشد.

برای لاگ‌های بلند، `output: auto` (پیش‌فرض) یا `output: file` روی همان دکمه را
ترجیح دهید. اگر فقط انتهای لاگ لازم است، خود فرمان را هم می‌توانید کوتاه کنید.


### قواعد `function_directory` { #function_directory-rules }

| وضعیت | نتیجه |
|-----------|--------|
| کلید وجود ندارد | گزارش info؛ فقط توابع توکار |
| کلید وجود دارد اما خالی است (`""`) | گزارش info؛ فقط توابع توکار |
| کلید به مسیر ناموجود یا غیرقابل دسترسی اشاره می‌کند | خطای قطعی؛ فرایند متوقف می‌شود |
| مسیر وجود دارد اما دایرکتوری خالی است | مجاز |

!!! warning "مسیر نادرست ربات را متوقف می‌کند"

    اگر `function_directory` به پوشه‌ای ناموجود یا ناخوانا اشاره کند، برنامه
    به‌جای شروع بدون توابع سفارشی، با خطا متوقف می‌شود.

## :material-send-circle-outline: `telegram` { #telegram }

| فیلد | نوع | الزامی | پیش‌فرض | توضیح |
|-------|------|----------|---------|-------------|
| `bot_token` | string | بله | — | توکن ربات از BotFather |
| `allowed_users` | list of string | بله | — | [کاربران مجاز](concepts/allowed-users.md) |
| `api` | string | خیر | `https://api.telegram.org` | URL پایهٔ Bot API |
| `proxy.enabled` | bool | خیر | `false` | استفاده از پراکسی برای Telegram API |
| `proxy.url` | string | مشروط | — | هنگام `true` بودن `proxy.enabled` الزامی |
| `insecure` | bool | خیر | `false` | نادیده گرفتن بررسی TLS (توصیه نمی‌شود) |

کاربران غیرمجاز پیامی شامل `user_id` و `username` دریافت می‌کنند تا از مدیر
دسترسی بخواهند. بار نخست شناسهٔ خودتان را نیز همین‌گونه پیدا می‌کنید؛
[اجرا در CLI ← گام ۵](installation/download-and-run.md#step-5-find-your-user-id-if-needed)
را ببینید.

!!! example "اتصال از طریق پراکسی"

    ```yaml title="بخش telegram با پراکسی"
    telegram:
      bot_token: "123456789:AAExampleTokenValue"
      allowed_users:
        - "123456789"  # numeric user ID
        - "@alice"  # or username
      proxy:
        enabled: true
        url: "socks5://127.0.0.1:10808"
    ```

برای اینکه کاربران مجاز بتوانند فرمان شِل را در تلگرام بنویسند، این تنظیم
را در **ریشهٔ** فایل بگذارید (نه زیر `telegram`):

!!! tip "افزودن یک تنظیم ریشه"

    ```yaml title="فعال کردن دکمهٔ Run Command"
    enable_run_command: true
    ```

## :material-menu: منو { #menu }

این بخش مرجع فیلدهاست. برای توضیح هدایت‌شده همراه نمونه،
[منو](concepts/menu.md) را ببینید. هر گره [دکمه](concepts/button.md) یا
[دسته](concepts/category.md):

| فیلد | نوع | الزامی | توضیح |
|-------|------|----------|-------------|
| `name` | string | بله | نام نمایشی (میان هم‌سطح‌ها بدون توجه به بزرگی حروف یکتا) |
| `type` | `category` \| `button` | بله | نوع گره |
| `items` | list | برای `category` بله | فرزندان؛ دسته باید دست‌کم یکی داشته باشد |
| `function` | string | برای `button` بله | نام [تابع](concepts/function.md) |
| `command` | string | برای `function: command` بله | فرمان شِل برای `command` توکار |
| `path` | string | برای `function: script` بله | مسیر اسکریپت برای `script` توکار |
| `icon` | string | خیر | پیشوند ایموجی اختیاری |
| `id` | string | خیر | شناسهٔ اختیاری این گره؛ می‌توانید حذفش کنید |
| `confirm` | bool | خیر | درخواست [تأیید](concepts/confirmation.md) پیش از اجرا (پیش‌فرض `false`) |
| `timeout` | duration | خیر | بازنویسی timeout سراسری |
| `workdir` | string | خیر | بازنویسی دایرکتوری کاری |
| `env` | map | خیر | متغیر محیطی بیشتر برای این دکمه |
| `columns` | int | خیر | بازنویسی ستون‌های این دسته |
| `output` | `auto` \| `text` \| `file` | خیر | اختیاری. ننویسید تا `output` ریشه استفاده شود. فقط وقتی بگذارید که این دکمه حتماً `file` یا `text` باشد ([مقدار خروجی فرمان](#how-much-command-output-you-see)) |
| `args` | string | خیر | آرگومان‌های اختیاری `script` |
| هر نام پارامتر تعریف‌شده | scalar | طبق تعریف تابع | مقدار داده‌شده به تابع، مانند `url`، `host`، `unit` یا `lines` |

روی **دکمه**، هر کلید scalar دیگر پارامتر تابع است. نام آن باید با پارامتری
از تابع انتخاب‌شده مطابقت داشته باشد. نام ناشناخته
[`validate`](cli.md#validate) را ناموفق می‌کند. مقادیر `int` یا `bool` نیز
بررسی می‌شوند. رشته، عدد و بولی مستقیم به‌شکل مقدار YAML نوشته می‌شوند و
عدد به نقل‌قول نیاز ندارد.

روی **دسته**، هر کلید خارج از فیلدهای دستهٔ بالا خطاست. دسته تابع اجرا
نمی‌کند، پس کلید پارامتر نمی‌پذیرد.

`command`، `path` و `args` میانبر پارامترهایی با همان نام‌اند. نام دیگر
[پارامترها](concepts/parameter.md) مستقیم روی دکمه نوشته می‌شود. مقادیر را
در نگاشت تودرتوی `params:` نگذارید.
[توابع ← فرستادن مقادیر از دکمه](functions/index.md#passing-values-from-a-button)
را ببینید.

## :material-math-log: `logging` { #logging }

اختیاری است. اگر حذف شود، گزارش‌گیر پیش‌فرض کنسول روی `stderr` با سطح
`info` استفاده می‌شود.

گزارش‌گیرهای نام‌دار:

!!! example "نوشتن گزارش عادی و فایل حسابرسی"

    ```yaml title="بخش logging با فایل حسابرسی"
    logging:
      logs:
        default:
          level: info
          format: console  # or JSON
          output:
            - output: stderr
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

خروجی‌های پشتیبانی‌شده: `stdout`، `stderr`، `file` و `discard`.

گزارش‌گیر `audit` بالا هر فرمان را ثبت می‌کند: اجراکننده، دکمه، کد خروج و
مدت اجرا. [گزارش حسابرسی](concepts/audit-log.md) را ببینید.

## صفحات مرتبط

- [اجرا در CLI](installation/download-and-run.md) — ساخت و اجرای نخستین پیکربندی
- [منو](concepts/menu.md) — شرح کامل درخت منو
- [توابع](functions/index.md) — معنی `function`، `command`، `path` و `args`
- [CLI](cli.md) — اعتبارسنجی و اجرا با پیکربندی
