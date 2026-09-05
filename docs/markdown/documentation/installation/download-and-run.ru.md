---
title: Запустить в CLI
description: Запустите своего бота с терминала, шаг за шагом, от загрузки релиза и написания небольшого файла конфигурации до нажатия первой кнопки.
icon: material/console
---
# :material-console: Запуск в CLI

Эта страница приведет вас из ничего к работающему боту. Никакого
предварительного опыта работы с проект нужен. Если слово непонятно, проверьте
[Основные понятия](../concepts/config-file.md) страницы.

## :material-clipboard-check-outline: Прежде чем начать

Вам нужны две вещи от Telegram:

1. **Токен бота.** Откройте чат с [@BotFather](https://t.me/BotFather),
   отправьте `/newbot`, следуйте инструкциям и скопируйте токен, который он вам
   предоставит. Это выглядит например `123456789:AAExampleTokenValue`.
2. **Ваш числовой идентификатор пользователя.** Это номер, а не ваш `@username`.
   Если вы его не знаете, не волнуйтесь — бот сообщит вам ваш идентификатор при
   первом посещении сообщите об этом (см. [шаг
   5](#step-5-find-your-user-id-if-needed)).

## :material-download: Шаг 1: загрузка

Загрузите архив релиза и распакуйте его.

!!! example "Загрузите и откройте папку выпуска"

    ```bash title="Download and extract the release"
    wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
    tar -xzf telegram-commander.tar.gz
    cd telegram-commander
    ```

Внутри папки вы найдете:

- `telegram-commander-linux-amd64` и `telegram-commander-linux-arm64` —
  программы, по одной на каждый тип процессора.
- `config-examples/` — готовые [конфигурационные файлы
  ](../concepts/config-file.md) (см. [Конфигурация](../configuration.md))
- `functions/` — пример пользовательских [функций ](../concepts/function.md)
  (см. [Пользовательские функции](../functions/index.md#custom-functions))

## :material-chip: Шаг 2: выберите бинарный файл

!!! info "Какой файл предназначен для вашей машины?"

    Большинство серверов и компьютеров имеют код `amd64` (также называемый
    `x86_64`). Маленькие платы ARM а некоторые облачные виртуальные машины —
    `arm64`.

    Если вы не уверены, запустите `uname -m`: `x86_64` означает amd64, `aarch64`
    означает arm64.

=== ":fontawesome-brands-linux: AMD64"

    ```bash title="Keep the amd64 program"
    mv telegram-commander-linux-amd64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-arm64
    ```

=== ":fontawesome-brands-linux: ARM64"

    ```bash title="Keep the arm64 program"
    mv telegram-commander-linux-arm64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-amd64
    ```

Теперь у вас есть одна программа с именем `telegram-commander`.

## :material-file-cog-outline: Шаг 3: создайте свою конфигурацию

Скопируйте минимальный пример в рабочий файл:

!!! example "Создать редактируемый конфиг"

    ```bash title="Copy the example config"
    cp config-examples/config.minimal.yaml ./config.yaml
    ```

Откройте `config.yaml` и замените два заполнителя:

- `YOUR_BOT_TOKEN` — токен от BotFather
- `YOUR_USER_ID` — ваш числовой идентификатор (или пока оставьте его и смотрите
  шаг 5)

Чтобы узнать, что означает каждый параметр, прочтите
[Конфигурация](../configuration.md).

## :material-file-check-outline: Шаг 4: проверка

Всегда проверяйте конфигурацию перед запуском. Это выявляет опечатки и ошибки
без запуск бота.

!!! success "Проверьте работоспособность конфигурации"

    ```bash title="Validate the config"
    ./telegram-commander validate --config config.yaml
    ```

Если он напечатает `Valid configuration`, все в порядке. Если нет, то там
указано, что именно ошибается и где. Подробности см. на странице
[CLI](../cli.md#validate).

## :material-account-search: Шаг 5: найдите свой идентификатор пользователя (при необходимости) { #step-5-find-your-user-id-if-needed }

Если вы не знали свой идентификатор пользователя, задайте в `config.yaml` только
токен, поставьте любой номер в `allowed_users` на данный момент, затем запустите
бота:

!!! info "Начните один раз, чтобы увидеть свой идентификатор пользователя"

    ```bash title="Run the bot to learn your id"
    ./telegram-commander run --config config.yaml
    ```

Откройте Telegram, найдите своего бота и отправьте ему любое сообщение. Потому
что вы не в
[разрешённых пользователей](../concepts/allowed-users.md), бот отвечает вашим
`user_id` и `username`. Скопируйте этот идентификатор в `allowed_users`,
остановите бота с помощью `Ctrl+C` и запустите его еще раз.

Такое поведение является частью того, как работает контроль доступа; увидеть
[Конфигурация → `telegram`](../configuration.md#telegram).

## :material-play-circle-outline: Шаг 6: запустить

!!! example "Запустить бота в терминале"

    ```bash title="Start the bot"
    ./telegram-commander run --config config.yaml
    ```

Откройте своего бота в Telegram и отправьте `/start`. Вы должны увидеть свое
меню. Нажмите
[кнопка](../concepts/button.md), чтобы запустить команду.

!!! success "Ваш бот активен"

    Меню, которое вы описали в `config.yaml`, теперь есть в вашем чате, и каждое
    нажатие запускает свою команду на этой машине.

Чтобы бот продолжал работать после выхода из сервера, настройте его как службу.
См. [Запуск как служба ](run-as-a-service.md).

## :material-map-marker-path: Что дальше

- Добавьте больше [кнопки](../concepts/button.md) и
  [категории](../concepts/category.md): [Меню](../concepts/menu.md)
- Поймите, что на самом деле работает: [Функции](../functions/index.md)
- Просмотреть все параметры командной строки: [CLI](../cli.md)
