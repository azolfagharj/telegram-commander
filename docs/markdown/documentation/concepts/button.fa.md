---
title: دکمه
description: گزینه‌ای قابل لمس در منو که تابعی را اجرا می‌کند. هر مقدار تابع را با نام پارامتری که تابع تعریف کرده، مستقیم روی دکمه بنویسید.
icon: material/gesture-tap-button
---

# :material-gesture-tap-button: دکمه

گزینه‌ای قابل لمس در منوی تلگرام. دکمه یک `name` دارد و با لمس شدن،
[تابعی](function.md) را روی سرور اجرا می‌کند و نتیجه را به گفتگو می‌فرستد.

دکمه‌ها و [دسته‌ها](category.md) با هم درخت [منو](menu.md) را زیر کلید
سطح بالای `menu` در [فایل پیکربندی](config-file.md) می‌سازند. دکمه کار را
انجام می‌دهد؛ دسته فقط یک زیرمنو باز می‌کند.

## ظاهر دکمه

دکمه‌ها روی صفحه‌کلید زیر کادر پیام ظاهر می‌شوند؛ به‌طور پیش‌فرض دو دکمه در
هر ردیف. نامی که انتخاب می‌کنید متن روی کلید است، پس آن را برای خواندن روی
گوشی به‌اندازهٔ کافی کوتاه نگه دارید.
[منو ← ظاهر منوی تلگرام](menu.md#how-the-telegram-menu-looks) را ببینید.

## :material-format-list-checks: اجزای دکمه

!!! example "هر خط یک بخش از دکمه را کنترل می‌کند"

    ```yaml title="یک دکمه با همهٔ برچسب‌ها"
    - name: Restart nginx # (1)!
      type: button # (2)!
      icon: "🔄" # (3)!
      function: command # (4)!
      command: "systemctl restart nginx" # (5)!
      confirm: true # (6)!
    ```

    1.  متن روی کلید در تلگرام. باید میان گزینه‌های هم‌سطح خود در همان منو
        یکتا باشد.
    2.  اینجا همیشه `button` است. برای زیرمنو از `category` استفاده کنید.
    3.  ایموجی اختیاری پیش از نام. فقط تزئینی است و چیزی را در اجرا تغییر
        نمی‌دهد.
    4.  [تابع](function.md) مورد استفاده. `command` تابع توکاری است که فرمان
        شِل اجرا می‌کند.
    5.  چیزی که تابع `command` اجرا می‌کند. هرچه بتوانید در ترمینال بنویسید،
        اینجا کار می‌کند.
    6.  اختیاری. پیش از اجرا می‌پرسد «مطمئن هستید؟». برای کارهایی که فقط
        اطلاعات می‌خوانند آن را حذف کنید.

## با لمس دکمه چه رخ می‌دهد

1. ربات خط کوتاه **Running** را می‌فرستد تا بدانید کار شروع شده است.
2. فرمان روی دستگاهی اجرا می‌شود که ربات در آن در حال اجراست.
3. خروجی به‌شکل بلوک کد، همراه کد خروج و مدت اجرا برمی‌گردد. خروجی بلند در
   چند پیام پیاپی می‌آید.
4. در همان منو می‌مانید؛ بنابراین **Back** همچنان شما را از آن دسته بیرون
   می‌برد.

## :material-code-braces: دکمه‌های رایج

=== "بررسی یک مورد"

    ```yaml title="دکمهٔ Uptime"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

=== "مدیریت سرویس"

    ```yaml title="دکمهٔ راه‌اندازی مجدد nginx"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

=== "خواندن گزارش"

    ```yaml title="دکمهٔ گزارش Nginx"
    - name: Nginx log
      type: button
      function: command
      command: "journalctl -u nginx -n 50 --no-pager"
    ```

=== "اجرای اسکریپت"

    ```yaml title="دکمهٔ پشتیبان‌گیری شبانه"
    - name: Nightly backup
      type: button
      function: script
      path: "/usr/local/bin/backup.sh"
    ```

=== "کاری مخرب"

    ```yaml title="دکمهٔ توقف nginx"
    - name: Stop nginx
      type: button
      icon: "🛑"
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## :material-emoticon-outline: آیکون‌ها

`icon` یک ایموجی پیش از نام می‌گذارد. کاملاً ظاهری است؛ پس هر زمان می‌توانید
آن را بدون دست زدن به کاری که دکمه اجرا می‌کند تغییر دهید یا حذف کنید.

!!! example "آیکون فقط برچسب را تغییر می‌دهد"

    ```yaml title="یک دکمه با آیکون و بدون آن"
    - name: Disk usage
      type: button
      function: command
      command: "df -h"

    - name: Disk usage
      type: button
      icon: "💾"
      function: command
      command: "df -h"
    ```

## :material-help-circle-outline: پرسش پیش از دکمه‌های پرخطر

`confirm: true` را بیفزایید تا ربات ابتدا Yes یا Cancel بپرسد. برای هر کاری
که سرویسی را متوقف، داده‌ای را حذف یا دستگاه را راه‌اندازی مجدد می‌کند از
آن استفاده کنید. پیام پس از مدتی منقضی می‌شود (پیش‌فرض پنج دقیقه).

برای موارد مناسب و تغییر زمان انتظار، [تأیید](confirmation.md) را بخوانید.

## تنظیمات تنها برای یک دکمه

بیشتر تنظیمات سراسری را می‌توان روی یک دکمه بازنویسی کرد؛ وقتی یک کار با
بقیه متفاوت است، این قابلیت مفید خواهد بود:

!!! example "بازنویسی تنظیمات یک کار کند"

    ```yaml title="کاری کند که جای دیگری اجرا می‌شود"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"
      workdir: "/var/backups"
      env:
        BACKUP_MODE: "full"
    ```

`timeout` زمان بیشتری برای پایان این فرمان می‌دهد، `workdir` دایرکتوری اجرا
را تعیین می‌کند و `env` فقط برای همین فرمان متغیر محیطی می‌افزاید.

## مقادیر تابع

مقادیر تابع را مستقیم روی دکمه بنویسید. `command`، `path` و `args` فیلدهای
میانبری برای پارامترهایی با همین نام هستند. نام‌های سفارشی مانند `url`،
`host`، `unit` و `lines` نیز به همین شکل کار می‌کنند.

!!! example "فرستادن مقادیر سفارشی"

    ```yaml title="گزارش‌های اخیر Nginx"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

این مقادیر را درون `params:` نگذارید. هر کلید باید با پارامتری که تابع
انتخاب‌شده تعریف کرده مطابقت داشته باشد. [`validate`](../cli.md#validate)
نام‌های ناشناخته، مقادیر الزامی گم‌شده و مقادیر نامعتبر `int` یا `bool` را
گزارش می‌کند.

## پیکربندی

برای همهٔ فیلدهایی که دکمه می‌پذیرد،
[پیکربندی ← منو](../configuration.md#menu) را ببینید.

## مرتبط

<div class="grid cards cols-2" markdown>

-   :material-folder-outline:{ .middle } __دسته__

    ---

    به‌جای اجرای کاری، زیرمنو باز می‌کند.

    [:octicons-arrow-left-24: دسته](category.md)

-   :material-function:{ .middle } __تابع__

    ---

    چیزی که با لمس دکمه واقعاً اجرا می‌شود.

    [:octicons-arrow-left-24: تابع](function.md)

-   :material-tune:{ .middle } __پارامتر__

    ---

    مقادیری که تابع از دکمه نیاز دارد.

    [:octicons-arrow-left-24: پارامتر](parameter.md)

-   :material-view-list:{ .middle } __منو__

    ---

    ساخت و سازمان‌دهی کل درخت.

    [:octicons-arrow-left-24: منو](menu.md)

</div>
