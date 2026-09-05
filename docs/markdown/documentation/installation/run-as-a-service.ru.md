---
title: Запуск как служба
description: Держите бот работающим в фоновом режиме с помощью systemd, чтобы он запускался при загрузке и возвращался после сбоя, как root или как обычный пользователь.
icon: material/server
---
# :material-server: Запуск как служба

На этой странице описывается запуск Telegram Commander в качестве службы
systemd. Для пошаговый первый запуск в терминале см. в разделе [Запуск в
CLI](download-and-run.md).

## :material-cog-outline: systemd (корневой)

Создайте файл модуля самостоятельно, например
`/etc/systemd/system/telegram-commander.service`:

!!! warning "В этом примере каждая кнопка запускается от имени пользователя root"

    ```ini title="/etc/systemd/system/telegram-commander.service"
    [Unit]
    Description=telegram-commander Telegram bot
    After=network-online.target
    Wants=network-online.target

    [Service]
    Type=simple
    ExecStart=/path/to/telegram-commander run --config /path/to/configfile.yaml
    Restart=on-failure
    RestartSec=5

    [Install]
    WantedBy=multi-user.target
    ```

    Без строки `User=` служба запускается от имени пользователя root, поэтому
    каждая команда в вашем меню имеет полные права на машине. Добавьте
    `User=someone` к `[Service]`. раздел, если вашим кнопкам не требуется
    столько прав.

Замените пути заполнителей, затем:

!!! example "Загрузить и запустить службу"

    ```bash title="Enable, start, and watch the service"
    sudo systemctl daemon-reload
    sudo systemctl enable --now telegram-commander
    sudo systemctl status telegram-commander
    sudo journalctl -u telegram-commander -f
    ```

!!! info "Изменения в конфигурации требуют перезагрузки"

    Служба поддерживает работу бота, но считывает ваши
    [файл конфигурации](../concepts/config-file.md) один раз при запуске. Беги
    `sudo systemctl restart telegram-commander` после его редактирования.

## Похожие страницы

- [Запуск в CLI](download-and-run.md) — объяснение первого запуска
- [Конфигурация](../configuration.md) — файл конфигурации
- [CLI](../cli.md) — `run`, `validate` и другие.
