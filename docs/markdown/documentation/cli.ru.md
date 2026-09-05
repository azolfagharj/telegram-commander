---
title: Справочник по интерфейсу командной строки
description: Все параметры командной строки Telegram Commander, от запуска бота до проверки и форматирования файла конфигурации, а также используемый ими флаг --config.
icon: material/console-line
---
# :material-console-line: Справочник по CLI

!!! info "Команды используют эту форму"

    ```text title="Command syntax"
    telegram-commander <command> [flags]
    ```

`--config`/`-c` **требуется** для команд, загружающих конфигурацию. Путь может
быть относительным к текущему рабочему каталогу или абсолютным. Что входит в это
файл описан в [Конфигурация](configuration.md).

Если вы только начинаете, [Запуск в CLI](installation/download-and-run.md)
показывает эти команды в том порядке, в котором вы будете их использовать.

## :material-format-list-bulleted-square: Команды

### `run` { #run }

Запустите бота на переднем плане.

!!! info "Запустить бота на переднем плане"

    ```bash title="Run the bot"
    telegram-commander run --config /path/to/configfile.yaml
    ```

!!! note "Изменения конфигурации требуют перезагрузки"

    Бот читает конфигурацию один раз при запуске. После редактирования файла
    перезапустите процесс (например, `systemctl restart telegram-commander`).

### `validate` { #validate }

Проверка конфигурации, функций и ссылок на кнопки в автономном режиме.

!!! info "Проверьте конфигурацию в автономном или онлайн-режиме"

    ```bash title="Validate the config"
    telegram-commander validate --config /path/to/configfile.yaml
    telegram-commander validate --config /path/to/configfile.yaml --online
    ```

!!! note "`--online` нужен Интернет"

    При использовании `--online` проверка также спрашивает Telegram, работает ли
    токен бота, поэтому машина должна иметь возможность связаться с Telegram.

### `version`

Выведите версию программы.

### `fmt`

Отформатируйте файл конфигурации YAML.

!!! info "Распечатать или сохранить форматированный YAML"

    ```bash title="Format a config file"
    telegram-commander fmt --config /path/to/configfile.yaml
    telegram-commander fmt --config /path/to/configfile.yaml -w
    ```

### `environ`

Вывести переменные среды процесса (полезно для отладки модулей службы).

### `list-functions` { #list-functions }

Список встроенных и загруженных пользовательских функций. Используйте его, чтобы
убедиться, что нужные файлы функций найдены файлы функций. См.
[Функции](functions/index.md).

!!! info "Показать все доступные функции"

    ```bash title="List available functions"
    telegram-commander list-functions --config /path/to/configfile.yaml
    ```

### `completion`

Создайте скрипты автодополнения для оболочки:

!!! info "Выберите оболочку, которую вы используете"

    ```bash title="Generate a completion script"
    telegram-commander completion bash
    telegram-commander completion zsh
    telegram-commander completion fish
    telegram-commander completion powershell
    ```

### `manpage`

Выведите страницу руководства в стандартный вывод.

## Похожие страницы

- [Конфигурация](configuration.md) — файл передан `--config`
- [Пользовательские функции](functions/index.md#custom-functions) — что
  показывает `list-functions`
- [Запуск как служба ](installation/run-as-a-service.md) — запустить `run` через
  systemd
