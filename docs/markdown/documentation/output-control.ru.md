---
title: Управление выводом
description: Решите, как результат команды попадёт в ваш чат. Задайте output, max_output_messages и max_output_bytes в корне, переопределите output на отдельной кнопке и посмотрите все сочетания корня и кнопки с примерами.
icon: material/export-variant
---
# :material-export-variant: Управление выводом { #control-output }

Когда [кнопка](concepts/button.md) заканчивает работу, боту нужно положить
результат в ваш чат. Telegram никогда не принимает больше 4096 байт в одном
сообщении, поэтому длинный результат приходит либо **несколькими ответными
сообщениями**, либо **одним файлом `.txt`**, который можно открыть и скачать.

Что произойдёт, решают три необязательных поля. Можно не указывать ни одного —
бот всё равно поведёт себя разумно.

!!! info "Одно правило, которое стоит запомнить"

    Кнопка главнее. Если кнопка задаёт `output`, будет именно так. Если на
    кнопке ничего не сказано, применяется корневое значение. Если и в корне
    ничего нет, бот использует `auto`.

<div class="grid cards cols-3" markdown>

-   :material-scissors-cutting:{ .middle } __`max_output_bytes`__

    ---

    Сколько вывода команды сохраняется до того, как что-то будет отправлено.

    [:octicons-arrow-right-24: Подробнее](#max_output_bytes)

-   :material-tune:{ .middle } __`output`__

    ---

    Сообщения или файл. Единственное поле, которое можно задать и на кнопке.

    [:octicons-arrow-right-24: Подробнее](#output)

-   :material-counter:{ .middle } __`max_output_messages`__

    ---

    В режиме `auto` — сколько сообщений разрешено до отправки файла.

    [:octicons-arrow-right-24: Подробнее](#max_output_messages)

</div>

!!! tip "Каждая картинка на этой странице открывается"

    Снимки чата специально небольшие. Нажмите на любой, чтобы увидеть его
    в полный размер.

## :material-card-bulleted-outline: Три поля коротко { #the-three-fields-at-a-glance }

| Поле | Где задаётся | Тип | По умолчанию | Что решает |
|-------|------------------|------|---------|-----------------|
| `max_output_bytes` | только в корне | целое число | `524288` | Сколько байт вывода бот сохраняет для каждой команды |
| `output` | корень **и** кнопка | `auto` \| `text` \| `file` | `auto` | Придёт результат сообщениями в чат или одним файлом `.txt` |
| `max_output_messages` | только в корне | целое число, `1`–`10` | `2` | В режиме `auto` — сколько сообщений разрешено, прежде чем вместо них будет отправлен файл |

Все три живут в вашем [файле конфигурации](concepts/config-file.md), в
**корне**, рядом с `shell` и `timeout`, а не внутри `telegram`. Полный список
корневых ключей — на странице [Конфигурация](configuration.md#root-fields).

## :material-scissors-cutting: `max_output_bytes`

Это ваш собственный лимит, и он действует первым. Пока команда работает, бот
сохраняет не больше этого количества байт её вывода, отдельно для обычного
вывода и для вывода ошибок. Всё сверх этого отбрасывается, но сама команда
продолжает работать до завершения или до своего `timeout`.

!!! example "Сохранять до 2 МБ вывода на команду"

    ```yaml title="config.yaml"
    max_output_bytes: 2097152 # (1)!
    ```

    1.  2 МБ вместо 512 КБ по умолчанию. Считается отдельно для обычного
        вывода и для вывода ошибок.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Когда лимит достигнут, результат сообщает об этом отдельной строкой, прямо
    под сводкой. Всё выше `(output truncated)` — это часть, которую бот
    сохранил; остальное так и не покинуло сервер.

    </div>
    <div class="result-shot" markdown>

    ![Результат в Telegram со строкой об обрезанном выводе](/images/output-control/truncated.png){ .shot loading=lazy }

    </div>
    </div>

!!! tip "Увеличивать его есть смысл только когда приходит файл"

    В сообщение чата никогда не поместится больше 4096 байт, поэтому больший
    `max_output_bytes` доходит до вас целиком только тогда, когда результат
    приходит файлом `.txt`. Используйте его вместе с `output: auto` (по
    умолчанию) или `output: file`.

## :material-tune: `output`

`output` выбирает способ доставки. Он принимает три значения и это единственное
поле вывода, которое можно задать и на отдельной кнопке.

=== "auto"

    Текстовые сообщения, пока результат короткий, и один файл `.txt`, когда он
    длинный. Это значение по умолчанию — то, что действует, если вы ничего не
    написали.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Пусть бот сам решает для каждого результата"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: auto # (1)!
        max_output_messages: 2 # (2)!
        ```

        1.  Решать для каждого результата, а не навязывать один способ
            доставки.
        2.  Точка переключения. Два сообщения или меньше остаются в чате;
            всё, что длиннее, становится одним файлом.

    </div>
    <div class="result-shot" markdown>

    ![В режиме auto короткий вывод остаётся текстом, а длинный приходит файлом](/images/output-control/auto.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

=== "text"

    Всегда сообщения в чате, каким бы длинным ни был результат. Бот
    останавливается после десяти сообщений и добавляет пометку, что остальное
    обрезано.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Оставить всё в чате"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: text # (1)!
        ```

        1.  `max_output_messages` здесь не учитывается, потому что бот сам
            никогда не переключается на файл.

    </div>
    <div class="result-shot" markdown>

    ![Результат команды, доставленный сообщениями Telegram](/images/output-control/message.png){ .shot loading=lazy }

    </div>
    </div>

=== "file"

    Всегда один файл `.txt`, даже для результата в одну строку. Сводка идёт
    подписью к файлу.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Отправлять каждый результат вложением"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: file # (1)!
        ```

        1.  Теперь вложение отправляет каждая кнопка, включая самые быстрые.
            Если это слишком, указывайте это на отдельных кнопках.

    </div>
    <div class="result-shot" markdown>

    ![Результат команды, доставленный файлом Telegram](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

!!! note "Не указать его — то же самое, что написать `auto`"

    Писать `output` не обязательно. Конфигурация, которая работала до появления
    этого поля, ведёт себя точно так же.

## :material-counter: `max_output_messages`

Это поле имеет значение только в режиме `auto` и существует только в корне. Это
число сообщений в чате, которые бот готов отправить, прежде чем откажется от
сообщений и отправит один файл. Допустимые значения — от `1` до `10`; по
умолчанию `2`.

!!! example "Отправлять файл, как только результату нужно второе сообщение"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 1 # (1)!
    ```

    1.  Самая строгая настройка. В чате остаётся только результат, который
        умещается в одно сообщение.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Результат, который умещается в разрешённое число сообщений, остаётся
    сообщениями в чате. Как только ему понадобится ещё одно сообщение, весь
    результат придёт одним файлом `.txt`, и ни одного частичного сообщения
    отправлено не будет.

    </div>
    <div class="result-shot" markdown>

    ![Короткий результат сообщением рядом с более длинным результатом файлом](/images/output-control/allowance.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

!!! warning "Это поле есть только в корне"

    Кнопка не может задать `max_output_messages`. Укажите его в корне, а
    кнопкам, которым нужно другое поведение, задайте `output: text` или
    `output: file`.

## :material-table-arrow-right: Корень и кнопка вместе { #root-and-button-together }

У кнопки может быть свой `output`. Больше ничего про вывод на кнопке задать
нельзя. Действующий режим выбирается просто:

!!! abstract "Как выбирается действующий режим"

    ```text title="Order of precedence"
    button output  →  root output  →  auto
    ```

| Корневой `output` | `output` кнопки | Что происходит на самом деле |
|---------------|-----------------|-----------------------|
| не задан или `auto` | не задан | `auto` |
| не задан или `auto` | `auto` | `auto` |
| не задан или `auto` | `text` | `text` |
| не задан или `auto` | `file` | `file` |
| `text` | не задан | `text` |
| `text` | `auto` | `auto` |
| `text` | `text` | `text` |
| `text` | `file` | `file` |
| `file` | не задан | `file` |
| `file` | `auto` | `auto` |
| `file` | `text` | `text` |
| `file` | `file` | `file` |

Разделы ниже проходят все двенадцать сочетаний, по одному примеру на каждое.

## :material-numeric-1-box-outline: Когда в корне `auto` { #when-the-root-is-auto }

Сюда же относится случай, когда в корне вообще ничего не сказано: отсутствующий
`output` означает `auto`.

### Корень `auto`, на кнопке ничего { #the-root-is-auto-the-button-says-nothing }

Обычный случай. Короткие результаты остаются в чате; результат, которому нужно
больше сообщений, чем `max_output_messages`, приходит файлом.

!!! example "Простая кнопка при простом корне"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 2

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    `df -h` печатает несколько строк, поэтому результат умещается в одно
    сообщение и остаётся в чате блоком кода со сводкой сверху.

    </div>
    <div class="result-shot" markdown>

    ![Использование диска пришло одним сообщением Telegram](/images/output-control/disk-usage.png){ .shot loading=lazy }

    </div>
    </div>

### Корень `auto`, на кнопке `auto` { #the-root-is-auto-the-button-says-auto }

Написать `auto` на кнопке здесь ничего не меняет. Это стоит делать только
тогда, когда вы хотите, чтобы кнопка осталась в `auto`, чем бы позже ни стал
корень.

!!! example "Явно указываем значение по умолчанию на кнопке"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
        output: auto # (1)!
    ```

    1.  Сегодня поведение такое же, как если поле не указывать, но эта кнопка
        сохранит его, даже если корень позже сменится на `text` или `file`.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Ровно автоматическое поведение: текст, пока результат короткий, и один
    файл, когда он перерастает разрешённое число сообщений. На сервере с
    множеством смонтированных дисков та же самая кнопка может прислать вам
    файл.

    </div>
    <div class="result-shot" markdown>

    ![Одна и та же кнопка сообщением и файлом](/images/output-control/disk-usage-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### Корень `auto`, на кнопке `text` { #the-root-is-auto-the-button-says-text }

Используйте это для кнопки, вывод которой вы всегда хотите читать прямо в чате,
даже если он немного длинный — например список состояния, который вы
пролистываете, а не скачиваете.

!!! example "Оставить одну кнопку в чате"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Service status
        type: button
        function: command
        command: "systemctl status nginx --no-pager"
        output: text # (1)!
    ```

    1.  Никогда файл, что бы ни было в `max_output_messages`.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Результат разбивается на несколько сообщений, каждое из которых — ответ на
    предыдущее, поэтому порядок никогда не теряется. Кнопки остаются на
    последней части.

    </div>
    <div class="result-shot" markdown>

    ![Длинный статус, разбитый на три сообщения Telegram](/images/output-control/service-status-split.png){ .shot loading=lazy }

    </div>
    </div>

### Корень `auto`, на кнопке `file` { #the-root-is-auto-the-button-says-file }

Обычный выбор для журналов и резервных копий: читать их в чате почти никогда не
хочется, а файл проще сохранить.

!!! example "Эту всегда скачиваем"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Nginx logs
        type: button
        function: command
        command: "journalctl -u nginx -n 2000 --no-pager"
        output: file # (1)!
    ```

    1.  Остальное меню по-прежнему работает автоматически. Только эта кнопка
        закреплена на файле.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Одно вложение, а в подписи — имя кнопки, код выхода и длительность.
    Нажмите на него, чтобы прочитать весь журнал, или сохраните на потом.

    </div>
    <div class="result-shot" markdown>

    ![Журналы Nginx, доставленные файлом Telegram](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-2-box-outline: Когда в корне `output: text` { #when-the-root-is-output-text }

Каждая кнопка остаётся в чате, если не сказано иное. Выбирайте это, если не
любите скачивать файлы на телефон.

### Корень `text`, на кнопке ничего { #the-root-is-text-the-button-says-nothing }

!!! example "Сообщения в чате везде"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Одно короткое сообщение, потому что `uptime` печатает одну строку. Более
    длинные результаты просто разбиваются на большее число сообщений, до десяти.

    </div>
    <div class="result-shot" markdown>

    ![Время работы пришло одним сообщением Telegram](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### Корень `text`, на кнопке `auto` { #the-root-is-text-the-button-says-auto }

Кнопка снова включает автоматическое переключение, поэтому именно она может
прислать файл, хотя остальное меню этого никогда не делает. Где произойдёт
переключение, решает `max_output_messages` из корня.

!!! example "Одной кнопке разрешено отправить файл, когда вывод станет длинным"

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

    1.  Разрешённое число сообщений читает только эта кнопка, потому что она
        единственная в режиме `auto`.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Короткий список пакетов приходит до трёх сообщений. На заполненном сервере
    та же кнопка выходит за разрешённое число и присылает один файл.

    </div>
    <div class="result-shot" markdown>

    ![Список пакетов сообщениями и файлом](/images/output-control/package-list-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### Корень `text`, на кнопке `text` { #the-root-is-text-the-button-says-text }

То же самое, что ничего не писать, но намерение записано. Полезно в длинной
конфигурации, где вы не хотите, чтобы поздняя правка в корне изменила эту
кнопку.

!!! example "Закрепляем кнопку на сообщениях в чате"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Last logins
        type: button
        function: command
        command: "last -n 20"
        output: text
    ```

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Всегда сообщения в чате, до десяти штук, чем бы позже ни стал корень.

    </div>
    <div class="result-shot" markdown>

    ![Последние входы пришли сообщением Telegram](/images/output-control/last-logins.png){ .shot loading=lazy }

    </div>
    </div>

### Корень `text`, на кнопке `file` { #the-root-is-text-the-button-says-file }

!!! example "Одна скачиваемая кнопка в меню, где всё остальное — сообщения"

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

    1.  Отчёт может готовиться долго, поэтому эта кнопка получает больше
        времени, чем общий тайм-аут.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Единственное вложение в меню, где всё остальное остаётся сообщениями. В
    подписи видно, сколько занял отчёт, — это удобно, когда он идёт минутами.

    </div>
    <div class="result-shot" markdown>

    ![Полный отчёт о системе, доставленный файлом Telegram](/images/output-control/system-report-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-3-box-outline: Когда в корне `output: file` { #when-the-root-is-output-file }

Каждый результат приходит файлом, даже если он в одну строку. Это подходит
боту, которым вы пользуетесь в основном для отчётов и архивов.

### Корень `file`, на кнопке ничего { #the-root-is-file-the-button-says-nothing }

!!! example "Файлы везде"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Даже одна строка вывода приходит вложением. Сводка всё равно есть в
    подписи, поэтому код выхода можно прочитать, не открывая файл.

    </div>
    <div class="result-shot" markdown>

    ![Время работы, доставленное небольшим файлом Telegram](/images/output-control/uptime-file.png){ .shot loading=lazy }

    </div>
    </div>

!!! warning "Даже крошечные результаты нужно скачивать"

    Ответ в одну строку приходит вложением, которое придётся открыть. Если это
    раздражает, оставьте в корне `auto` и указывайте `output: file` только на
    тех кнопках, где это нужно.

### Корень `file`, на кнопке `auto` { #the-root-is-file-the-button-says-auto }

Кнопка выходит из правила «только файлы» и снова ведёт себя обычным образом.

!!! example "Пусть быстрая проверка останется читаемой"

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

    1.  Для этой кнопки — снова автоматическое переключение.

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Ответ в одну строку короткий, поэтому он остаётся в чате сообщением, а все
    остальные кнопки меню по-прежнему присылают файлы.

    </div>
    <div class="result-shot" markdown>

    ![Время работы снова в чате сообщением](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### Корень `file`, на кнопке `text` { #the-root-is-file-the-button-says-text }

!!! example "Вернуть одну кнопку в чат"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Who is logged in
        type: button
        function: command
        command: "who"
        output: text
    ```

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Сообщение, и никогда вложение, как бы ни вырос список вошедших
    пользователей. После десяти сообщений бот отмечает, что остальное обрезано.

    </div>
    <div class="result-shot" markdown>

    ![Вошедшие пользователи пришли сообщением Telegram](/images/output-control/who.png){ .shot loading=lazy }

    </div>
    </div>

### Корень `file`, на кнопке `file` { #the-root-is-file-the-button-says-file }

Повторение корневого значения. Вреда нет, и кнопка останется правильной, если
позже вы ослабите корень.

!!! example "Кнопка, которая всегда должна быть файлом"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Database dump log
        type: button
        function: command
        command: "cat /var/log/pg_dump.log"
        output: file
    ```

!!! success "Что вы увидите в чате"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Вложение — ровно то, что уже просил корень. Запись на кнопке сохраняет
    намерение, если корень изменится.

    </div>
    <div class="result-shot" markdown>

    ![Журнал дампа базы данных, доставленный файлом Telegram](/images/output-control/db-dump-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-file-document-outline: Как выглядит файл { #what-the-file-looks-like }

Когда результат приходит файлом, вы получаете одно вложение и короткую подпись.

- **Имя** — имя кнопки строчными буквами через дефисы, затем дата и время в
  UTC, например `nginx-logs-20260913-091204.txt`.
- **Подпись** — та же сводка, которую вы увидели бы в начале сообщения в чате:
  имя кнопки, код выхода, длительность, а также пометка о тайм-ауте или обрезке,
  если она есть.
- **Тело** — снова та же сводка, затем раздел `--- stdout ---`, а затем раздел
  `--- stderr ---`, если команда что-то в него написала.

!!! success "Внутри скачанного файла"

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

## :material-alert-outline: Ограничения, которые нельзя изменить { #limits-you-cannot-change }

!!! warning "Верхний предел ставит Telegram, а не бот"

    - Одно сообщение в чате вмещает не больше 4096 байт.
    - В режиме `text` бот отправляет не больше десяти сообщений и заканчивает
      строкой `(output too long; showing first N bytes)`.
    - Какое бы `max_output_messages` вы ни написали, десять сообщений остаются
      жёстким пределом.

!!! note "Если отправить файл не удалось, результат вы всё равно получите"

    Когда файл отправить невозможно — нет сети, Telegram его не принимает — бот
    возвращается к сообщениям в чате, чтобы результат не потерялся.

## :material-console: Выполнение команды берёт корневую настройку { #run-command-uses-the-root-setting }

Кнопка **`$ >_ Выполнить команду`** не входит в ваше меню, поэтому у неё нет
своего `output`. Она всегда следует корневым `output` и
`max_output_messages`. См. [Меню → Выполнение команды](concepts/menu.md#run-command).

## :material-link-variant: Похожие { #related }

<div class="grid cards cols-2" markdown>

-   :material-cog-outline:{ .middle } __Конфигурация__

    ---

    Каждый корневой ключ, включая три поля вывода в одной таблице.

    [:octicons-arrow-right-24: Конфигурация](configuration.md#root-fields)

-   :material-gesture-tap-button:{ .middle } __Кнопка__

    ---

    Что кнопка может переопределить, в том числе `output`.

    [:octicons-arrow-right-24: Кнопка](concepts/button.md)

-   :material-view-list:{ .middle } __Меню__

    ---

    Как результаты стоят рядом с меню и что делает «Выполнить команду».

    [:octicons-arrow-right-24: Меню](concepts/menu.md)

-   :material-timer-outline:{ .middle } __Shell__

    ---

    Оболочка, рабочий каталог и тайм-аут, с которыми выполняются ваши команды.

    [:octicons-arrow-right-24: Shell](concepts/shell.md)

</div>
