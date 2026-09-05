---
title: Telegram Commander
description: یک فایل YAML را به ربات تلگرامی تبدیل کنید که فرمان‌ها را روی سرور لینوکس شما اجرا می‌کند، با هر بار زدن دکمه.
icon: material/cellphone-link
hide:
  - navigation
  - toc
---

# :material-cellphone-link: سرور لینوکس را از تلگرام کنترل کنید

<div class="hero" markdown>
<div class="hero-art" markdown>
![Telegram Commander](/images/logo-large.png){ .off-glb width="230" }
</div>
<div class="hero-text" markdown>
**یک ضربه فرمان را روی سرور اجرا می‌کند و خروجی را به گفتگو برمی‌گرداند.**

Telegram Commander یک [فایل پیکربندی](documentation/concepts/config-file.md)
سادهٔ YAML را به ربات تلگرام با منوی [دکمه](documentation/concepts/button.md)
تبدیل می‌کند. هر فرمان ترمینال را پشت یک دکمه بگذارید و از گوشی اجرا کنید.
لازم نیست کدی بنویسید.

<p class="hero-lang">این صفحه را به زبان دیگری بخوانید:
<a href="/" hreflang="en" class="hero-lang__link">English</a>
<a href="/de/" hreflang="de" class="hero-lang__link">Deutsch</a>
<a href="/fr/" hreflang="fr" class="hero-lang__link">Français</a>
<a href="/es/" hreflang="es" class="hero-lang__link">Español</a>
<a href="/ru/" hreflang="ru" class="hero-lang__link">Русский</a>
<a href="/zh/" hreflang="zh" class="hero-lang__link">简体中文</a>
</p>
</div>
</div>

<div style="text-align: center" markdown="span">
[شروع کنید :material-arrow-left:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[دیدن پیکربندی :material-file-code-outline:](documentation/configuration.md#a-minimal-config){ .md-button }
</div>

[نصب](documentation/installation/download-and-run.md) ·
[مفاهیم](documentation/concepts/config-file.md) ·
[توابع](documentation/functions/index.md) ·
[پیکربندی](documentation/configuration.md) ·
[CLI](documentation/cli.md)

## :material-image-multiple-outline:{ .shots } نماگرفت‌ها { .split }

منوی شما، فرمان در حال اجرا، خروجی برگشتی، و تایپ دستی یک فرمان. روی هر
تصویر بزنید تا تمام‌صفحه دیده شود.

<div style="text-align: center" markdown="span">
  ![خروجی فرمان با منوی System باز](/images/01.jpeg){ width="140" loading=lazy }
  ![دکمه‌های منابع و فرایندها](/images/02.jpeg){ width="140" loading=lazy }
  ![دکمه‌های فضای ذخیره‌سازی و بسته‌ها](/images/03.jpeg){ width="140" loading=lazy }
  ![ابزار شبکه و تایپ دستی فرمان](/images/04.jpeg){ width="140" loading=lazy }
  ![خروجی فرمان برگشته به گفتگو](/images/05.jpeg){ width="140" loading=lazy }
</div>

## :material-lightning-bolt:{ .bolt } سریع و ساده { .split }

<div class="grid cards cols-3 center-title step-cards" markdown>

-   :material-file-document-outline:{ .middle } __منو را بنویسید__

    ---

    :material-numeric-1-circle:{ .step } دکمه‌ها و فرمان‌هایشان را فهرست کنید.

-   :material-rocket-launch:{ .middle } __ربات را شروع کنید__

    ---

    :material-numeric-2-circle:{ .step } همین حالا اجرا کنید، یا به‌صورت
    سرویس نگه دارید.

-   :material-gesture-tap-button:{ .middle } __بزنید و بخوانید__

    ---

    :material-numeric-3-circle:{ .step } دکمه را بزنید و خروجی را در گفتگو
    بخوانید.

</div>

<div style="text-align: center" markdown="span">
[الان شروع کنید :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
</div>

## :material-view-grid-outline:{ .grid-icon } کاربردها { .split }

<div class="grid cards cols-4 icon-left" markdown>

-   :material-restart:{ .lg } راه‌اندازی مجدد یا توقف سرویس
-   :material-docker:{ .lg } شروع و توقف کانتینر
-   :material-package-down:{ .lg } به‌روزرسانی بسته‌های سیستم
-   :material-text-box-search-outline:{ .lg } خواندن لاگ و journal
-   :material-harddisk:{ .lg } بررسی فضای دیسک
-   :material-memory:{ .lg } نظارت بر CPU و حافظه
-   :material-access-point-network:{ .lg } پینگ میزبان و آزمایش URL
-   :material-backup-restore:{ .lg } تهیه و بازگردانی پشتیبان
-   :material-script-text:{ .lg } اجرای اسکریپت‌های خودتان
-   :material-power:{ .lg } راه‌اندازی مجدد یا خاموشی میزبان
-   :material-console:{ .lg } تایپ دستی هر فرمان
-   :material-all-inclusive:{ .lg } و تقریباً هر کار دیگر

</div>

## :material-thumb-up-outline:{ .thumb } چرا از آن استفاده کنید { .split }

<div class="grid cards cols-4 center-title" markdown>

-   :material-clock-fast:{ .lg .middle } __بدون کدنویسی__

    ---

    منو و فرمان‌ها را در یک فایل YAML توصیف کنید.

    [:octicons-arrow-left-24: فایل پیکربندی](documentation/concepts/config-file.md)

-   :material-cellphone-link:{ .lg .middle } __از هر جا__

    ---

    تلگرام را روی گوشی باز کنید و سرور را اجرا کنید. VPN به میزبان لازم نیست.

    [:octicons-arrow-left-24: نحوهٔ اتصال ربات](documentation/concepts/long-polling.md)

-   :material-lan-disconnect:{ .lg .middle } __بدون پورت باز__

    ---

    ربات به بیرون به تلگرام وصل می‌شود. چیزی به اینترنت نمایان نیست.

    [:octicons-arrow-left-24: نحوهٔ اتصال ربات](documentation/concepts/long-polling.md)

-   :material-message-text-outline:{ .lg .middle } __خروجی در گفتگو__

    ---

    نتیجه به‌صورت پیام برمی‌گردد. به نشست SSH نیاز ندارید.

    [:octicons-arrow-left-24: چقدر از خروجی را می‌بینید](documentation/configuration.md#how-much-command-output-you-see)

-   :material-shield-lock:{ .lg .middle } __کنترل‌شده و ثبت‌شده__

    ---

    مشخص کنید چه کسی منو را ببیند، کارهای پرخطر را تأیید کنید و هر اجرا را ثبت کنید.

    [:octicons-arrow-left-24: دسترسی و تأیید](documentation/concepts/allowed-users.md)

-   :material-folder-outline:{ .lg .middle } __منوهای تو در تو__

    ---

    دکمه‌ها را در دسته گروه‌بندی کنید. خانه بالا می‌ماند؛ بازگشت یک سطح بالا می‌رود.

    [:octicons-arrow-left-24: منو](documentation/concepts/menu.md)

-   :material-function-variant:{ .lg .middle } __توابع قابل استفاده مجدد__

    ---

    فرمان را یک‌بار بنویسید، بعد روی هر دکمه مقدارهای مختلف بگذارید.

    [:octicons-arrow-left-24: توابع](documentation/functions/index.md)

-   :material-cog-play-outline:{ .lg .middle } __در حال اجرا می‌ماند__

    ---

    به‌صورت سرویس نصب کنید تا ربات با میزبان شروع شود.

    [:octicons-arrow-left-24: اجرا به‌صورت سرویس](documentation/installation/run-as-a-service.md)

</div>

## :material-file-code-outline:{ .code-icon } یک نمونهٔ کوچک { .split }

این پیکربندی رباتی با یک دکمه به نام «Uptime» می‌سازد. زدنش فرمان `uptime`
را روی سرور اجرا می‌کند.

!!! example "این پیکربندی کامل یک دکمه اضافه می‌کند"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN" # (1)!
      allowed_users:
        - "YOUR_USER_ID" # (2)!

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime" # (3)!
    ```

    1.  هنگام ساخت ربات، توکن را از BotFather در تلگرام بگیرید.
    2.  فقط حساب‌های این فهرست می‌توانند منو را باز کنند. می‌توانید شناسهٔ
        عددی یا `@username` بگذارید.
    3.  هر چیزی که در ترمینال می‌توانید تایپ کنید اینجا می‌آید.

این یک پیکربندی کامل و کارکننده است. بقیه اختیاری‌اند.

!!! tip "برای یک گروه کوچک و مورد اعتماد"

    چیزی به اتصال ورودی گوش نمی‌دهد و فقط حساب‌های `allowed_users` منو
    می‌گیرند. هرکس بتواند از ربات استفاده کند، دکمه‌هایی را که تعریف کرده‌اید
    اجرا می‌کند؛ پس این فهرست را کوتاه نگه دارید.

## :material-hand-pointing-left: آمادهٔ امتحانید؟ { .split }

<div style="text-align: center" markdown>
[الان شروع کنید :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[مفاهیم :material-book-open-variant:](documentation/concepts/config-file.md){ .md-button }
[آخرین انتشار :material-download:](https://github.com/azolfagharj/telegram-commander/releases/latest){ .md-button .md-button--primary }

[دیدن منبع :fontawesome-brands-github:](https://github.com/azolfagharj/telegram-commander){ .md-button }
</div>

Telegram Commander رایگان و متن‌باز است. اگر وقتتان را ذخیره می‌کند،
[از توسعهٔ آن حمایت کنید](https://azolfagharj.github.io/donate/) —
کمک می‌کند پروژه زنده و نگهداری‌شده بماند.
