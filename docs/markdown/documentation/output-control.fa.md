---
title: کنترل خروجی
description: تعیین کنید نتیجهٔ فرمان چگونه به گفتگو برسد. output و max_output_messages و max_output_bytes را در ریشه تنظیم کنید، output را روی یک دکمه بازنویسی کنید و همهٔ ترکیب‌های ریشه و دکمه را با نمونه ببینید.
icon: material/export-variant
---

# :material-export-variant: کنترل خروجی

وقتی کار یک [دکمه](concepts/button.md) تمام می‌شود، ربات باید نتیجه را در
گفتگوی شما بگذارد. تلگرام هرگز بیش از ۴۰۹۶ بایت را در یک پیام نمی‌پذیرد؛ پس
نتیجهٔ بلند یا در **چند پیام پاسخ‌به‌پاسخ** می‌آید یا به‌صورت **یک فایل `.txt`**
که می‌توانید آن را باز کنید و بگیرید.

سه فیلد اختیاری تعیین می‌کنند چه رخ دهد. می‌توانید هر سه را ننویسید و ربات
باز هم رفتار معقولی دارد.

!!! info "تنها قاعده‌ای که باید به یاد بسپارید"

    دکمه برنده است. اگر دکمه `output` را تنظیم کند، همان اتفاق می‌افتد. اگر
    دکمه چیزی نگوید، مقدار ریشه اعمال می‌شود. اگر ریشه هم چیزی نگوید، ربات
    از `auto` استفاده می‌کند.

<div class="grid cards cols-3" markdown>

-   :material-scissors-cutting:{ .middle } __`max_output_bytes`__

    ---

    پیش از فرستادن هر چیزی، چه مقدار از خروجی فرمان نگه داشته می‌شود.

    [:octicons-arrow-left-24: بیشتر بخوانید](#max_output_bytes)

-   :material-tune:{ .middle } __`output`__

    ---

    پیام یا فایل. تنها فیلدی که روی دکمه هم می‌توانید تنظیم کنید.

    [:octicons-arrow-left-24: بیشتر بخوانید](#output)

-   :material-counter:{ .middle } __`max_output_messages`__

    ---

    در حالت `auto`، چند پیام اجازه دارد پیش از فرستادن فایل بیاید.

    [:octicons-arrow-left-24: بیشتر بخوانید](#max_output_messages)

</div>

!!! tip "هر تصویر این صفحه باز می‌شود"

    تصویرهای گفتگو از قصد کوچک هستند. روی هرکدام بزنید تا در اندازهٔ کامل
    ببینید.

## :material-card-bulleted-outline: سه فیلد در یک نگاه { #the-three-fields-at-a-glance }

| فیلد | جای تنظیم | نوع | پیش‌فرض | چه چیزی را تعیین می‌کند |
|------|-----------|-----|---------|--------------------------|
| `max_output_bytes` | فقط ریشه | int | `524288` | ربات برای هر فرمان چند بایت از خروجی را نگه می‌دارد |
| `output` | ریشه **و** دکمه | `auto` \| `text` \| `file` | `auto` | نتیجه به‌صورت پیام‌های گفتگو بیاید یا یک فایل `.txt` |
| `max_output_messages` | فقط ریشه | int، `1` تا `10` | `2` | در حالت `auto`، چند پیام اجازه دارد بیاید پیش از آنکه به‌جای آن فایل فرستاده شود |

هر سه در [فایل پیکربندی](concepts/config-file.md) شما و در **ریشه**، کنار
`shell` و `timeout` قرار می‌گیرند، نه زیر `telegram`. فهرست کامل کلیدهای ریشه
در صفحهٔ [پیکربندی](configuration.md#root-fields) است.

## :material-scissors-cutting: `max_output_bytes`

این محدودیت خودتان است و اول از همه اعمال می‌شود. هنگام اجرای فرمان، ربات
حداکثر همین تعداد بایت از خروجی آن را نگه می‌دارد و خروجی عادی و خروجی خطا را
جداگانه می‌شمارد. بیشتر از آن دور ریخته می‌شود، ولی خود فرمان تا پایان کار یا
رسیدن به `timeout` خودش ادامه می‌دهد.

!!! example "نگه داشتن تا ۲ مگابایت خروجی برای هر فرمان"

    ```yaml title="config.yaml"
    max_output_bytes: 2097152 # (1)!
    ```

    1.  ۲ مگابایت به‌جای پیش‌فرض ۵۱۲ کیلوبایت. برای خروجی عادی و خروجی خطا
        جداگانه شمرده می‌شود.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    وقتی محدودیت پر شود، نتیجه همین را در خط جداگانه‌ای درست زیر خلاصه
    می‌گوید. هر چیزی بالای `(output truncated)` بخشی است که ربات نگه داشته؛
    بقیهٔ گزارش هرگز از سرور بیرون نیامده است.

    </div>
    <div class="result-shot" markdown>

    ![نتیجه‌ای در تلگرام که بریده‌شدن خروجی را اعلام می‌کند](/images/output-control/truncated.png){ .shot loading=lazy }

    </div>
    </div>

!!! tip "بالا بردن این مقدار فقط وقتی فایل فرستاده شود به کار می‌آید"

    یک پیام گفتگو هرگز بیش از ۴۰۹۶ بایت را نمی‌برد؛ پس `max_output_bytes`
    بزرگ‌تر تنها وقتی کامل به دست شما می‌رسد که نتیجه به‌صورت فایل `.txt`
    بیاید. آن را با `output: auto` (پیش‌فرض) یا `output: file` همراه کنید.

## :material-tune: `output`

`output` شیوهٔ تحویل را انتخاب می‌کند. سه مقدار می‌پذیرد و تنها فیلد خروجی است
که روی یک دکمه هم می‌توانید بگذارید.

=== "auto"

    تا وقتی نتیجه کوتاه است پیام متنی، و وقتی کوتاه نیست یک فایل `.txt`.
    این حالت پیش‌فرض است و اگر چیزی ننویسید همین می‌ماند.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "بگذارید ربات برای هر نتیجه خودش انتخاب کند"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: auto # (1)!
        max_output_messages: 2 # (2)!
        ```

        1.  به‌جای تحمیل یک شیوهٔ تحویل، برای هر نتیجه جداگانه تصمیم بگیرید.
        2.  نقطهٔ تغییر. دو پیام یا کمتر در گفتگو می‌ماند؛ بلندتر از آن یک
            فایل می‌شود.

    </div>
    <div class="result-shot" markdown>

    ![حالت auto خروجی کوتاه را متن نگه می‌دارد و خروجی بلندتر را فایل می‌فرستد](/images/output-control/auto.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

=== "text"

    همیشه پیام‌های گفتگو، هر اندازه که نتیجه بلند باشد. ربات پس از ده پیام
    می‌ایستد و یادداشتی می‌افزاید که بقیه بریده شد.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "نگه داشتن همه چیز در گفتگو"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: text # (1)!
        ```

        1.  `max_output_messages` اینجا نادیده گرفته می‌شود، چون ربات هرگز
            خودش به فایل تغییر نمی‌دهد.

    </div>
    <div class="result-shot" markdown>

    ![نتیجهٔ یک فرمان که به‌شکل پیام‌های تلگرام آمده است](/images/output-control/message.png){ .shot loading=lazy }

    </div>
    </div>

=== "file"

    همیشه یک فایل `.txt`، حتی برای نتیجهٔ یک‌خطی. خلاصه هم به‌عنوان کپشن فایل
    همراهش می‌آید.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "فرستادن هر نتیجه به‌صورت پیوست"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: file # (1)!
        ```

        1.  از این پس هر دکمه پیوست می‌فرستد، حتی دکمه‌های سریع. اگر این
            زیاده‌روی است، به‌جایش آن را روی دکمه‌های تک بگذارید.

    </div>
    <div class="result-shot" markdown>

    ![نتیجهٔ یک فرمان که به‌شکل فایل تلگرام آمده است](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

!!! note "ننوشتن آن مثل نوشتن `auto` است"

    هرگز لازم نیست `output` را بنویسید. پیکربندی‌ای که پیش از وجود این فیلد
    کار می‌کرد، دقیقاً همان رفتار را نگه می‌دارد.

## :material-counter: `max_output_messages`

این فیلد فقط در حالت `auto` اهمیت دارد و فقط در ریشه وجود دارد. تعداد
پیام‌های گفتگویی است که ربات حاضر است بفرستد، پیش از آنکه از پیام دست بکشد و
به‌جای آن یک فایل بفرستد. مقدارهای مجاز `1` تا `10` هستند و پیش‌فرض `2` است.

!!! example "فرستادن فایل به‌محض آنکه نتیجه به پیام دوم نیاز داشت"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 1 # (1)!
    ```

    1.  سخت‌گیرانه‌ترین تنظیم. فقط نتیجه‌ای که در یک پیام جا شود در گفتگو
        می‌ماند.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    نتیجه‌ای که در سهمیهٔ تعیین‌شده جا شود به‌شکل پیام‌های گفتگو می‌ماند. همان
    لحظه که به یک پیام بیشتر نیاز پیدا کند، کل نتیجه به‌جای آن در یک فایل
    `.txt` می‌آید و هیچ پیام نیمه‌کاره‌ای فرستاده نمی‌شود.

    </div>
    <div class="result-shot" markdown>

    ![یک نتیجهٔ کوتاه به‌شکل پیام در کنار نتیجه‌ای بلندتر به‌شکل فایل](/images/output-control/allowance.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

!!! warning "این فیلد فقط در ریشه است"

    دکمه نمی‌تواند `max_output_messages` را تنظیم کند. آن را در ریشه بگذارید
    و بگذارید دکمه‌هایی که چیز دیگری می‌خواهند `output: text` یا
    `output: file` را تنظیم کنند.

## :material-table-arrow-left: ریشه و دکمه با هم { #root-and-button-together }

دکمه می‌تواند `output` خودش را داشته باشد. هیچ چیز دیگری از خروجی را نمی‌توان
برای هر دکمه جداگانه تنظیم کرد. حالت مؤثر به‌سادگی این است:

!!! abstract "حالت مؤثر چگونه انتخاب می‌شود"

    ```text title="ترتیب اولویت"
    button output  →  root output  →  auto
    ```

| `output` ریشه | `output` دکمه | واقعاً چه رخ می‌دهد |
|---------------|---------------|----------------------|
| تنظیم‌نشده یا `auto` | تنظیم‌نشده | `auto` |
| تنظیم‌نشده یا `auto` | `auto` | `auto` |
| تنظیم‌نشده یا `auto` | `text` | `text` |
| تنظیم‌نشده یا `auto` | `file` | `file` |
| `text` | تنظیم‌نشده | `text` |
| `text` | `auto` | `auto` |
| `text` | `text` | `text` |
| `text` | `file` | `file` |
| `file` | تنظیم‌نشده | `file` |
| `file` | `auto` | `auto` |
| `file` | `text` | `text` |
| `file` | `file` | `file` |

بخش‌های زیر همهٔ دوازده ترکیب را با یک نمونه برای هرکدام مرور می‌کنند.

## :material-numeric-1-box-outline: وقتی ریشه `auto` است { #when-the-root-is-auto }

این حالت شامل وقتی هم می‌شود که ریشه هیچ چیزی نگفته باشد، چون نبودن `output`
یعنی `auto`.

### ریشه `auto` است و دکمه چیزی نمی‌گوید { #the-root-is-auto-the-button-says-nothing }

حالت هر روزه. نتیجهٔ کوتاه در گفتگو می‌ماند؛ نتیجه‌ای که به بیش از
`max_output_messages` پیام نیاز داشته باشد به‌صورت فایل می‌آید.

!!! example "یک دکمهٔ ساده زیر یک ریشهٔ ساده"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 2

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    `df -h` چند خط چاپ می‌کند؛ پس در یک پیام جا می‌شود و به‌شکل بلوک کد با
    خلاصه‌ای بالای آن در گفتگو می‌ماند.

    </div>
    <div class="result-shot" markdown>

    ![استفادهٔ دیسک که در یک پیام تلگرام برگشته است](/images/output-control/disk-usage.png){ .shot loading=lazy }

    </div>
    </div>

### ریشه `auto` است و دکمه `auto` می‌گوید { #the-root-is-auto-the-button-says-auto }

نوشتن `auto` روی دکمه اینجا چیزی را تغییر نمی‌دهد. تنها وقتی ارزش نوشتن دارد
که بخواهید این دکمه هر تغییری هم که بعداً در ریشه بدهید در حالت `auto` بماند.

!!! example "نوشتن صریح مقدار پیش‌فرض روی دکمه"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
        output: auto # (1)!
    ```

    1.  امروز رفتارش با ننوشتن آن یکی است، ولی این دکمه همین رفتار را نگه
        می‌دارد حتی اگر ریشه بعداً به `text` یا `file` تغییر کند.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    دقیقاً همان رفتار خودکار: تا وقتی نتیجه کوتاه است متن، و به‌محض گذشتن از
    سهمیه یک فایل. روی سروری با دیسک‌های زیاد همین دکمه می‌تواند به‌جای متن
    فایل بفرستد.

    </div>
    <div class="result-shot" markdown>

    ![همان دکمه یک بار به‌شکل پیام و یک بار به‌شکل فایل](/images/output-control/disk-usage-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### ریشه `auto` است و دکمه `text` می‌گوید { #the-root-is-auto-the-button-says-text }

این را برای دکمه‌ای به کار ببرید که همیشه می‌خواهید خروجی‌اش را همان‌جا در
گفتگو بخوانید، حتی وقتی کمی بلند می‌شود؛ مثل فهرست وضعیتی که آن را پایین
می‌کشید، نه فایلی که می‌گیرید.

!!! example "نگه داشتن یک دکمه در گفتگو"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Service status
        type: button
        function: command
        command: "systemctl status nginx --no-pager"
        output: text # (1)!
    ```

    1.  هرگز فایل نمی‌شود، `max_output_messages` هر چه بگوید.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    نتیجه به چند پیام تقسیم می‌شود و هر پیام پاسخ پیام پیش از خودش است؛ پس
    ترتیب هرگز گم نمی‌شود. دکمه‌ها روی آخرین بخش می‌مانند.

    </div>
    <div class="result-shot" markdown>

    ![یک وضعیت بلند که به سه پیام تلگرام تقسیم شده است](/images/output-control/service-status-split.png){ .shot loading=lazy }

    </div>
    </div>

### ریشه `auto` است و دکمه `file` می‌گوید { #the-root-is-auto-the-button-says-file }

انتخاب معمول برای گزارش‌ها و پشتیبان‌ها: تقریباً هرگز نمی‌خواهید آن‌ها را در
گفتگو بخوانید و فایل را راحت‌تر نگه می‌دارید.

!!! example "این یکی همیشه گرفته شود"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Nginx logs
        type: button
        function: command
        command: "journalctl -u nginx -n 2000 --no-pager"
        output: file # (1)!
    ```

    1.  بقیهٔ منو همچنان خودکار رفتار می‌کند. فقط همین دکمه به فایل بسته شده
        است.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    یک پیوست با نام دکمه، کد خروج و مدت اجرا به‌عنوان کپشن. روی آن بزنید تا
    کل گزارش را بخوانید، یا آن را برای بعد نگه دارید.

    </div>
    <div class="result-shot" markdown>

    ![گزارش‌های nginx که به‌شکل فایل تلگرام آمده‌اند](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-2-box-outline: وقتی ریشه `output: text` است { #when-the-root-is-output-text }

هر دکمه در گفتگو می‌ماند، مگر خودش چیز دیگری بگوید. وقتی گرفتن فایل روی گوشی
را دوست ندارید این را انتخاب کنید.

### ریشه `text` است و دکمه چیزی نمی‌گوید { #the-root-is-text-the-button-says-nothing }

!!! example "پیام‌های گفتگو در همه جا"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    یک پیام کوتاه، چون `uptime` تنها یک خط چاپ می‌کند. نتیجه‌های بلندتر
    به‌سادگی به پیام‌های بیشتر، تا ده پیام، تقسیم می‌شوند.

    </div>
    <div class="result-shot" markdown>

    ![uptime که در یک پیام تلگرام برگشته است](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### ریشه `text` است و دکمه `auto` می‌گوید { #the-root-is-text-the-button-says-auto }

دکمه دوباره به تغییر خودکار می‌پیوندد؛ پس همین یک دکمه می‌تواند فایل بفرستد،
هرچند بقیهٔ منو هرگز این کار را نمی‌کند. `max_output_messages` ریشه تعیین
می‌کند این تغییر کجا رخ دهد.

!!! example "یک دکمه که وقتی خروجی‌اش بلند شد اجازهٔ فرستادن فایل دارد"

    ```yaml title="config.yaml"
    output: text
    max_output_messages: 3 # (1)!

    menu:
      - name: Package list
        type: button
        function: command
        command: "dpkg -l"
        output: auto
    ```

    1.  فقط همین دکمه سهمیه را می‌خواند، چون تنها دکمهٔ حالت `auto` است.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    فهرست کوتاه بسته‌ها تا سه پیام می‌آید. روی سروری پر از بسته، همین دکمه از
    سهمیه می‌گذرد و به‌جای آن یک فایل می‌فرستد.

    </div>
    <div class="result-shot" markdown>

    ![فهرست بسته‌ها یک بار به‌شکل پیام و یک بار به‌شکل فایل](/images/output-control/package-list-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### ریشه `text` است و دکمه `text` می‌گوید { #the-root-is-text-the-button-says-text }

مثل چیزی نگفتن است، ولی خواستهٔ شما نوشته می‌شود. در پیکربندی بلندی که
نمی‌خواهید ویرایش بعدی ریشه این دکمه را تغییر دهد، به کار می‌آید.

!!! example "بستن یک دکمه به پیام‌های گفتگو"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Last logins
        type: button
        function: command
        command: "last -n 20"
        output: text
    ```

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    همیشه پیام‌های گفتگو، تا ده پیام، هر چیزی هم که بعداً در ریشه بنویسید.

    </div>
    <div class="result-shot" markdown>

    ![ورودهای اخیر که در یک پیام تلگرام برگشته‌اند](/images/output-control/last-logins.png){ .shot loading=lazy }

    </div>
    </div>

### ریشه `text` است و دکمه `file` می‌گوید { #the-root-is-text-the-button-says-file }

!!! example "یک دکمهٔ قابل گرفتن در منویی که همه‌اش گفتگوست"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Full system report
        type: button
        function: command
        command: "/usr/local/bin/report.sh"
        output: file
        timeout: "5m" # (1)!
    ```

    1.  گزارش ممکن است طول بکشد؛ پس این دکمه زمانی بیشتر از timeout سراسری
        می‌گیرد.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    تنها پیوست در منویی که جز این همه‌اش گفتگوست. کپشن نشان می‌دهد گزارش چقدر
    طول کشیده، که وقتی چند دقیقه اجرا می‌شود به کار می‌آید.

    </div>
    <div class="result-shot" markdown>

    ![گزارش کامل سیستم که به‌شکل فایل تلگرام آمده است](/images/output-control/system-report-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-3-box-outline: وقتی ریشه `output: file` است { #when-the-root-is-output-file }

هر نتیجه یک فایل است، حتی نتیجهٔ یک‌خطی. این برای رباتی مناسب است که بیشتر
برای گزارش‌ها و بایگانی به کار می‌رود.

### ریشه `file` است و دکمه چیزی نمی‌گوید { #the-root-is-file-the-button-says-nothing }

!!! example "فایل در همه جا"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    حتی یک خط خروجی هم به‌شکل پیوست می‌آید. کپشن باز هم خلاصه را همراه دارد؛
    پس کد خروج را بی‌آنکه فایل را باز کنید می‌خوانید.

    </div>
    <div class="result-shot" markdown>

    ![uptime که به‌شکل یک فایل کوچک تلگرام آمده است](/images/output-control/uptime-file.png){ .shot loading=lazy }

    </div>
    </div>

!!! warning "حتی نتیجه‌های بسیار کوچک هم فایل می‌شوند"

    پاسخ یک‌خطی هم به‌شکل پیوستی می‌آید که باید بازش کنید. اگر این آزارتان
    می‌دهد، ریشه را روی `auto` بگذارید و `output: file` را فقط روی دکمه‌هایی
    بنویسید که به آن نیاز دارند.

### ریشه `file` است و دکمه `auto` می‌گوید { #the-root-is-file-the-button-says-auto }

دکمه از قاعدهٔ فقط‌فایل بیرون می‌آید و دوباره عادی رفتار می‌کند.

!!! example "بگذارید یک بررسی سریع خواندنی بماند"

    ```yaml title="config.yaml"
    output: file
    max_output_messages: 2

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
        output: auto # (1)!
    ```

    1.  بازگشت به تغییر خودکار، فقط برای همین دکمه.

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    پاسخ یک‌خطی کوتاه است؛ پس به‌شکل پیام در گفتگو می‌ماند، در حالی که هر
    دکمهٔ دیگر منو همچنان فایل می‌فرستد.

    </div>
    <div class="result-shot" markdown>

    ![uptime که دوباره به‌شکل پیام در گفتگو آمده است](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### ریشه `file` است و دکمه `text` می‌گوید { #the-root-is-file-the-button-says-text }

!!! example "برگرداندن یک دکمه به گفتگو"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Who is logged in
        type: button
        function: command
        command: "who"
        output: text
    ```

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    یک پیام، هرگز پیوست، فهرست کاربران واردشده هر چه هم بلند شود. بالای ده
    پیام، ربات یادداشت می‌کند که بقیه بریده شد.

    </div>
    <div class="result-shot" markdown>

    ![کاربران واردشده که در یک پیام تلگرام برگشته‌اند](/images/output-control/who.png){ .shot loading=lazy }

    </div>
    </div>

### ریشه `file` است و دکمه `file` می‌گوید { #the-root-is-file-the-button-says-file }

تکرار مقدار ریشه. بی‌آزار است و اگر بعداً ریشه را آزادتر کنید، دکمه را درست
نگه می‌دارد.

!!! example "دکمه‌ای که همیشه باید فایل باشد"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Database dump log
        type: button
        function: command
        command: "cat /var/log/pg_dump.log"
        output: file
    ```

!!! success "در گفتگو چه می‌بینید"

    <div class="result" markdown>
    <div class="result-text" markdown>

    یک پیوست، دقیقاً همان‌طور که ریشه از قبل خواسته بود. نوشتن آن روی دکمه
    باعث می‌شود این خواسته از تغییر ریشه جان سالم ببرد.

    </div>
    <div class="result-shot" markdown>

    ![گزارش دامپ پایگاه داده که به‌شکل فایل تلگرام آمده است](/images/output-control/db-dump-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-file-document-outline: فایل چه شکلی است { #what-the-file-looks-like }

وقتی نتیجه به‌صورت فایل تحویل می‌شود، یک پیوست به‌همراه کپشنی کوتاه می‌گیرید.

- **نام** — نام دکمه با حروف کوچک و خط تیره، سپس تاریخ و ساعت به وقت UTC،
  مثلاً `nginx-logs-20260913-091204.txt`.
- **کپشن** — همان خلاصه‌ای که بالای پیام گفتگو می‌دیدید: نام دکمه، کد خروج،
  مدت اجرا و یادداشت timeout یا بریده‌شدن، وقتی چنین چیزی باشد.
- **بدنه** — همان خلاصه، سپس بخش `--- stdout ---` و سپس بخش
  `--- stderr ---` وقتی فرمان چیزی در آن نوشته باشد.

!!! success "درون فایل گرفته‌شده"

    ```text title="nginx-logs-20260913-091204.txt"
    Button: Nginx logs
    Exit: 0
    Duration: 316ms

    --- stdout ---
    2026-09-13T09:11:58 nginx: worker process started
    ...

    --- stderr ---
    2026-09-13T09:12:01 nginx: warning: duplicate server name
    ```

## :material-alert-outline: محدودیت‌هایی که تغییرشان نمی‌دهید { #limits-you-cannot-change }

!!! warning "سقف را تلگرام تعیین می‌کند، نه ربات"

    - یک پیام گفتگو حداکثر ۴۰۹۶ بایت جا می‌دهد.
    - در حالت `text` ربات حداکثر ده پیام می‌فرستد و با
      `(output too long; showing first N bytes)` تمام می‌کند.
    - هر `max_output_messages` هم که بنویسید، ده پیام باز هم سقف قطعی است.

!!! note "اگر فرستادن فایل شکست بخورد، نتیجه را از دست نمی‌دهید"

    وقتی فایل فرستادنی نباشد — نبود شبکه، نپذیرفتن تلگرام — ربات به
    پیام‌های گفتگو برمی‌گردد تا نتیجه گم نشود.

## :material-console: Run Command از تنظیم ریشه پیروی می‌کند { #run-command-uses-the-root-setting }

دکمهٔ **`$ >_ Run Command`** بخشی از منوی شما نیست؛ پس `output` خودش را ندارد.
همیشه از `output` و `max_output_messages` ریشه پیروی می‌کند.
[منو ← Run Command](concepts/menu.md#run-command) را ببینید.

## :material-link-variant: مرتبط { #related }

<div class="grid cards cols-2" markdown>

-   :material-cog-outline:{ .middle } __پیکربندی__

    ---

    همهٔ کلیدهای ریشه، از جمله سه فیلد خروجی در یک جدول.

    [:octicons-arrow-left-24: پیکربندی](configuration.md#root-fields)

-   :material-gesture-tap-button:{ .middle } __دکمه__

    ---

    چیزهایی که دکمه می‌تواند بازنویسی کند، `output` یکی از آن‌هاست.

    [:octicons-arrow-left-24: دکمه](concepts/button.md)

-   :material-view-list:{ .middle } __منو__

    ---

    نتیجه‌ها چگونه کنار منو می‌نشینند و Run Command چه می‌کند.

    [:octicons-arrow-left-24: منو](concepts/menu.md)

-   :material-timer-outline:{ .middle } __Shell__

    ---

    شِل، دایرکتوری کاری و timeout که فرمان‌های شما با آن اجرا می‌شوند.

    [:octicons-arrow-left-24: Shell](concepts/shell.md)

</div>
