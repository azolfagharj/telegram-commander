---
title: Telegram Commander
description: Превратите файл YAML в бота Telegram, который запускает команды на вашем Linux-сервере одним нажатием за раз.
icon: material/cellphone-link
hide:
  - navigation
  - toc
---
# :material-cellphone-link: Управляйте своим Linux-сервером из Telegram

<div class="hero" markdown>
<div class="hero-art" markdown>
![Telegram Commander](/images/logo-large.png){ .off-glb width="230" }
</div>
<div class="hero-text" markdown>
**Одно касание запускает команду на вашем сервере и отправляет результат обратно
в ваш чат.**

Telegram Commander превращает простой YAML-[файл
конфигурации](documentation/concepts/config-file.md) в Telegram-бота с меню
[кнопок](documentation/concepts/button.md). Назначьте кнопке любую терминальную
команду и запускайте её с телефона. Вам не нужно писать код.

<p class="hero-lang">Читать эту страницу на другом языке:
<a href="/" hreflang="en" class="hero-lang__link">English</a>
<a href="/de/" hreflang="de" class="hero-lang__link">Deutsch</a>
<a href="/fr/" hreflang="fr" class="hero-lang__link">Français</a>
<a href="/es/" hreflang="es" class="hero-lang__link">Español</a>
<a href="/zh/" hreflang="zh" class="hero-lang__link">简体中文</a>
<a href="/fa/" hreflang="fa" class="hero-lang__link">فارسی</a>
</p>
</div>
</div>

<div style="text-align: center" markdown="span">
[Начало работы :material-arrow-right:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Пример конфигурации :material-file-code-outline:](documentation/configuration.md#a-minimal-config){ .md-button }
</div>

[Установка](documentation/installation/download-and-run.md) ·
[Концепции](documentation/concepts/config-file.md) ·
[Функции](documentation/functions/index.md) ·
[Конфигурация](documentation/configuration.md) ·
[CLI](documentation/cli.md)

## :material-image-multiple-outline:{ .shots } Скриншоты { .split }

Ваше меню, выполняемая команда, возвращаемый результат и ввод команды. вручную.
Нажмите на любую картинку, чтобы увидеть ее в полном размере.

<div style="text-align: center" markdown="span">
  ![Вывод команды при открытом системном меню](/images/01.jpeg){ width="140" loading=lazy }
  ![Кнопки ресурсов и процессов](/images/02.jpeg){ width="140" loading=lazy }
  ![Кнопки хранения и упаковки](/images/03.jpeg){ width="140" loading=lazy }
  ![Сетевые инструменты и ввод команды вручную](/images/04.jpeg){ width="140" loading=lazy }
  ![Вывод команды отправлен обратно в чат](/images/05.jpeg){ width="140" loading=lazy }
</div>

## :material-lightning-bolt:{ .bolt } Быстро и легко { .split }

<div class="grid cards cols-3 center-title step-cards" markdown>

- :material-file-document-outline:{ .middle } __Напишите свое меню__

    ---

    :material-numeric-1-circle:{ .step } Перечислите кнопки и их команды.

- :material-rocket-launch:{ .middle } __Запустить бота__

    ---

    :material-numeric-2-circle:{ .step } Запустите его сейчас или сохраните как
    сервис.

- :material-gesture-tap-button:{ .middle } __Нажмите и прочитайте__

    ---

    :material-numeric-3-circle:{ .step } Нажмите кнопку и прочитайте вывод в
    чат.

</div>

<div style="text-align: center" markdown="span">
[Начать сейчас :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
</div>

## :material-view-grid-outline:{ .grid-icon } Варианты использования { .split }

<div class="grid cards cols-4 icon-left" markdown>

- :material-restart:{ .lg } Перезапустить или остановить службу.
- :material-docker:{ .lg } Контейнеры запуска и остановки
- :material-package-down:{ .lg } Обновление системных пакетов.
- :material-text-box-search-outline:{ .lg } Чтение журналов
- :material-harddisk:{ .lg } Проверьте место на диске
- :material-memory:{ .lg } Смотреть процессор и память
- :material-access-point-network:{ .lg } Пропингуйте хосты и проверьте
  URL-адреса.
- :material-backup-restore:{ .lg } Создание и восстановление резервных копий
- :material-script-text:{ .lg } Запуск собственных скриптов
- :material-power:{ .lg } Перезагрузите или выключите хост.
- :material-console:{ .lg } Введите любую команду вручную.
- :material-all-inclusive:{ .lg } И почти все остальное

</div>

## :material-thumb-up-outline:{ .thumb } Зачем его использовать { .split }

<div class="grid cards cols-4 center-title" markdown>

- :material-clock-fast:{ .lg .middle } __Без кодирования__

    ---

    Опишите меню и команды в одном YAML-файле.

    [:octicons-arrow-right-24: Конфигурационный файл](documentation/concepts/config-file.md)

- :material-cellphone-link:{ .lg .middle } __Из любого места__

    ---

    Откройте Telegram на телефоне и управляйте сервером. Подключение к серверу
    по VPN не требуется.

    [:octicons-arrow-right-24: Как бот подключается](documentation/concepts/long-polling.md)

- :material-lan-disconnect:{ .lg .middle } __Нет открытых портов__

    ---

    Бот подключается к Telegram. Ничего не выставлено в Интернете.

    [:octicons-arrow-right-24: Как бот подключается](documentation/concepts/long-polling.md)

- :material-message-text-outline:{ .lg .middle } __Вывод в чат__

    ---

    Результат возвращается в виде сообщения. Вам не нужен сеанс SSH.

    [:octicons-arrow-right-24: Сколько вы видите ](documentation/configuration.md#how-much-command-output-you-see)

- :material-shield-lock:{ .lg .middle } __Контролируется и записывается__

    ---

    Выбирайте, кому доступно меню, подтверждайте рискованные действия и
    записывайте каждый запуск.

    [:octicons-arrow-right-24: Доступ и подтверждение](documentation/concepts/allowed-users.md)

- :material-folder-outline:{ .lg .middle } __Вложенные меню__

    ---

    Группируйте кнопки по категориям. Главная всегда находится сверху, а кнопка
    «Назад» возвращает на уровень выше.

    [:octicons-arrow-right-24: Меню](documentation/concepts/menu.md)

-   :material-function-variant:{ .lg .middle } __Многоразовые функции__

    ---

    Напишите команду один раз, затем введите разные значения для каждой кнопки.

    [:octicons-arrow-right-24: Функции](documentation/functions/index.md)

- :material-cog-play-outline:{ .lg .middle } __Продолжает работу__

    ---

    Установите его как сервис и бот запустится с хоста.

    [:octicons-arrow-right-24: Запуск как служба](documentation/installation/run-as-a-service.md)

</div>

## :material-file-code-outline:{ .code-icon } Крошечный пример { .split }

Этот конфиг делает бота с одной кнопкой под названием «Uptime». Нажатие на него
запускает Команда `uptime` на сервере.

!!! example "В этой полной конфигурации добавлена одна кнопка"

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

    1. Попросите у BotFather в Telegram токен при создании бота.
    2. Только указанные здесь учетные записи могут открыть меню. Вы можете
       использовать числовой идентификатор
        или `@username`.
    3. Сюда помещается все, что вы можете ввести в терминале.

Это полноценная рабочая конфигурация. Все остальное не является обязательным.

!!! tip "Создан для небольшой доверенной группы"

    Ничто не прослушивает входящие соединения, а только учетные записи в
    `allowed_users` получите меню. Каждый, кто может использовать бота, может
    запустить кнопки, которые вы определили, поэтому сделайте этот список
    коротким.

## :material-hand-pointing-right: Готовы попробовать? { .split }

<div style="text-align: center" markdown>
[Начать сейчас :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Концепции :material-book-open-variant:](documentation/concepts/config-file.md){ .md-button }
[Последний выпуск :material-download:](https://github.com/azolfagharj/telegram-commander/releases/latest){ .md-button .md-button--primary }

[Исходный код :fontawesome-brands-github:](https://github.com/azolfagharj/telegram-commander){ .md-button }
</div>

Telegram Commander бесплатен и имеет открытый исходный код. Если это сэкономит
вам время,
[рассмотрим возможность поддержки его разработки](https://azolfagharj.github.io/donate/) —
это помогает поддерживать развитие проекта.
