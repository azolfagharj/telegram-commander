---
title: منو
description: منویی را بسازید که ربات در تلگرام نشان می‌دهد. دکمه‌ها و دسته‌ها را در یک درخت ترکیب کنید و حرکت با Home و Back را بیاموزید.
icon: material/view-list
---

# :material-view-list: منو

منوی شما درختی از گره‌ها زیر کلید سطح بالای `menu` است. دو نوع گره وجود دارد:

- **[دسته](category.md)** یک زیرمنو باز می‌کند و `items` دارد.
- **[دکمه](button.md)** کاری را اجرا می‌کند و یک [تابع](function.md) دارد.

اگر این واژه‌ها تازه‌اند، ابتدا [مفاهیم](button.md) را بخوانید. برای فهرست
دقیق همهٔ فیلدها، [پیکربندی ← منو](../configuration.md#menu) را ببینید.

## :material-format-list-bulleted: منوی تخت

ساده‌ترین منو فهرستی از دکمه‌ها بدون تودرتوسازی است:

!!! example "ساخت منویی با سه دکمه"

    ```yaml title="سه دکمه، بدون دسته"
    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"

      - name: Free memory
        type: button
        function: command
        command: "free -h"

      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

در تلگرام `/start` را بفرستید. منوی خود را خواهید دید.

## :material-folder-outline: گروه‌بندی با دسته‌ها

با بزرگ شدن منو، کارهای مرتبط را در دسته‌ها گروه‌بندی کنید. با لمس دسته،
`items` آن نمایش داده می‌شوند. Home همیشه در منوست و Back فقط داخل دسته
ظاهر می‌شود.

!!! example "گروه‌بندی دکمه‌ها در دسته‌ها"

    ```yaml title="دسته‌های System و Services"
    menu:
      - name: System
        type: category
        icon: "💻"
        items:
          - name: Uptime
            type: button
            function: command
            command: "uptime"

          - name: Free memory
            type: button
            function: command
            command: "free -h"

      - name: Services
        type: category
        icon: "🔧"
        items:
          - name: Restart nginx
            type: button
            function: command
            command: "systemctl restart nginx"
            confirm: true
    ```

دسته‌ها می‌توانند تا هر عمقی تودرتو شوند. هر دسته باید دست‌کم یک گزینه داشته
باشد.

## :material-family-tree: نام‌ها میان گزینه‌های هم‌سطح باید یکتا باشند

دو گره زیر یک والد نمی‌توانند نام یکسان داشته باشند (بزرگی حروف نادیده گرفته
می‌شود). نمونهٔ زیر درست است، چون گره‌ها در دسته‌های متفاوت‌اند:

!!! example "استفادهٔ دوباره از نام زیر والدهای متفاوت"

    ```yaml title="نام یکسان زیر دو والد"
    menu:
      - name: Web
        type: category
        items:
          - name: Restart  # allowed
            type: button
            function: command
            command: "systemctl restart nginx"
      - name: Database
        type: category
        items:
          - name: Restart  # allowed; different parent
            type: button
            function: command
            command: "systemctl restart postgresql"
    ```

## :material-emoticon-outline: آیکون‌ها

`icon` ایموجی اختیاری پیش از نام است و فقط ظاهر را تغییر می‌دهد.

!!! warning "این دکمه دستگاه را راه‌اندازی مجدد می‌کند"

    ```yaml title="دکمه‌ای با آیکون ایموجی"
    - name: Reboot
      type: button
      icon: "🔁"
      function: command
      command: "reboot"
      confirm: true
    ```

!!! tip "یک ایموجی ساده و رایج برگزینید"

    بعضی ایموجی‌ها باعث بریده شدن یا بیرون زدن متن دکمه در برخی گوشی‌ها
    می‌شوند. اگر دکمه بریده دیده شد، ایموجی دیگری امتحان کنید.

## :material-cellphone-text: ظاهر منوی تلگرام { #how-the-telegram-menu-looks }

همهٔ دکمه‌ها روی صفحه‌کلید زیر کادر پیام ظاهر می‌شوند (صفحه‌کلیدی که با دکمهٔ
کوچک سمت راست کادر پیام پنهان و آشکار می‌شود). این صفحه‌کلید همیشه تمام عرض
گفتگو را می‌گیرد؛ بنابراین متن دکمه فشرده یا بریده نمی‌شود.

- **Home** همیشه نخستین دکمهٔ هر صفحه است و به صفحهٔ نخست برمی‌گردد.
- **Back** هنگامی ظاهر می‌شود که داخل یک دسته باشید.
- **$ >_ Run Command** هنگام روشن بودن `enable_run_command` ظاهر می‌شود.
- گزینه‌ها به‌طور پیش‌فرض دو عدد در هر ردیف قرار می‌گیرند. دسته با `columns`
  می‌تواند آن را تغییر دهد. برای صفحه‌های شلوغ **Prev** و **Next** میان
  صفحه‌ها جابه‌جا می‌شوند.
- دکمه‌های دارای `confirm: true` پیش از اجرا Yes / Cancel می‌پرسند.

!!! info "عنوان منو دوباره استفاده می‌شود و خروجی می‌ماند"

    عنوان منوی جدید (Home، دسته یا صفحه) جای عنوان قبلی را می‌گیرد تا گفتگو
    از صفحه‌های خالی پر نشود. خط **Running** و خروجی فرمان در گفتگو می‌مانند
    تا پس از باز کردن دوبارهٔ منو نیز بدانید چه چیزی اجرا شده است.

!!! info "خروجی بلند در چند پیام می‌آید"

    خروجی فرمان به‌شکل بلوک کد نشان داده می‌شود. اگر از یک پیام تلگرام
    بلندتر باشد، در چند پیام می‌آید و هرکدام پاسخ پیام قبلی است. بخش آخر
    دکمه‌های همان صفحه را نگه می‌دارد؛ پس **Back** همچنان خروج از دسته است.
    [پیکربندی ← مقدار خروجی فرمان](../configuration.md#how-much-command-output-you-see)
    را ببینید.

## :material-help-circle-outline: تأیید { #confirmation }

برای نیاز به لمس دوم («مطمئن هستید؟») پیش از اجرا، `confirm: true` را به
دکمه بیفزایید. برای کارهای مخرب از آن استفاده کنید.
[تأیید](confirmation.md) را ببینید.

!!! warning "این دکمه سرویسی را متوقف می‌کند"

    ```yaml title="دکمه‌ای که ابتدا سؤال می‌کند"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

پیام تأیید پس از مدتی منقضی می‌شود (پیش‌فرض ۵ دقیقه). آن را با
`confirm_ttl` تغییر دهید؛
[پیکربندی ← فیلدهای ریشه](../configuration.md#root-fields) را ببینید.

## :material-tune-variant: بازنویسی تنظیمات هر دکمه

برخی تنظیمات سراسری را می‌توان روی یک دکمه بازنویسی کرد:

!!! example "دادن تنظیمات جداگانه به یک دکمه"

    ```yaml title="یک دکمه با timeout، پوشه و متغیرهای خودش"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"  # may take longer than the main timeout
      workdir: "/var/backups"  # run in this folder
      env:
        BACKUP_MODE: "full"  # extra environment variable for this command
    ```

فهرست کامل را در [پیکربندی ← منو](../configuration.md#menu) ببینید.

## :material-view-grid-outline: کنترل چیدمان

`menu_columns` تعداد دکمه‌های **گزینه** در هر ردیف را تعیین می‌کند
(پیش‌فرض ۲). دسته می‌تواند با `columns` آن را بازنویسی کند. وقتی منو بیش از
`page_size` گزینه داشته باشد (پیش‌فرض ۸)، صفحه‌بندی می‌شود و Prev/Next تا
رسیدن به ابتدا یا انتها نمایش داده می‌شوند.
[پیکربندی ← فیلدهای ریشه](../configuration.md#root-fields) را ببینید.

## :material-console: Run Command

اگر `enable_run_command: true` را در ریشهٔ پیکربندی بگذارید، دکمهٔ
**$ >_ Run Command** در منو می‌ماند (پس از Back داخل دسته یا پس از Home در
صفحهٔ نخست). آن را بزنید و سپس فرمان شِل را بفرستید. ربات از همان شِل،
timeout، دایرکتوری کاری و محدودیت خروجی دکمه‌های دیگر استفاده می‌کند.

Home یا Back بدون اجرای چیزی پیام را لغو می‌کند. این قابلیت پیش‌فرض خاموش
است. [پیکربندی ← فیلدهای ریشه](../configuration.md#root-fields) را ببینید.

!!! warning "این قابلیت کل دستگاه را واگذار می‌کند"

    با روشن بودن Run Command، همهٔ کاربران مجاز می‌توانند هر فرمانی را روی
    میزبان اجرا کنند، نه فقط دکمه‌های تعریف‌شده. تنها وقتی آن را روشن کنید
    که تا این حد به همهٔ کاربران مجاز اعتماد دارید.

## :material-timer-sand: هر بار یک فرمان

فرمان‌های شما یکی پس از دیگری اجرا می‌شوند، نه هم‌زمان.

!!! info "لمس دوم منتظر نوبت می‌ماند"

    اگر هنگام اجرای دکمهٔ نخست، دکمهٔ دیگری را بزنید، دومی منتظر می‌ماند و
    سپس جداگانه اجرا می‌شود. با شروع آن خط **Running** را می‌بینید. این کار
    از تداخل دو لمس خودتان بر سر یک سرویس یا فایل جلوگیری می‌کند. کاربران
    دیگر منتظر فرمان شما نمی‌مانند؛ هر شخص نوبت خودش را دارد.

## :material-function-variant: با لمس دکمه چه چیزی اجرا می‌شود

هر دکمه از طریق فیلد `function` به یک **تابع** اشاره می‌کند. دکمه‌های
نمونهٔ بالا از تابع توکار `command` استفاده می‌کنند. برای درک توابع، تفاوت
توکار و سفارشی و روش افزودن تابع خودتان، [توابع](../functions/index.md) را
بخوانید.

## :material-link-variant: صفحات مرتبط

- [دکمه](button.md) — دکمه چیست
- [دسته](category.md) — گره‌های زیرمنو
- [پیکربندی ← منو](../configuration.md#menu) — همهٔ فیلدها
