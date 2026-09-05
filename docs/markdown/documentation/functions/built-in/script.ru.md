---
title: сценарий
description: Встроенная функция сценария запускает файл сценария на сервере с дополнительными аргументами, добавляемыми после пути. Скрипт должен быть исполняемым.
icon: material/script-text
---
# :material-script-text: `script`

`script` — это [встроенная функция ](../index.md#built-in-functions). Она
запускает файл сценария с необязательными аргументами.

| Параметр | Требуется | По умолчанию | Значение |
|-----------|----------|---------|---------|
| `path` | да | — | Путь к скрипту |
| `args` | нет | пустой | Аргументы, передаваемые после пути |

!!! example "Запустить скрипт с аргументами"

    ```yaml title="Nightly report button"
    - name: Nightly report
      type: button
      function: script
      path: "/usr/local/bin/report.sh"
      args: "--today"
    ```

Это запускает `/usr/local/bin/report.sh --today`. Если не указывать `args`,
запустится только `/usr/local/bin/report.sh`.

## Скрипт должен быть исполняемым

Путь запускается напрямую, поэтому файлу нужен бит выполнения:

!!! tip "Сделать скрипт исполняемым один раз"

    ```bash title="Allow the file to run"
    chmod +x /usr/local/bin/report.sh
    ```

Если вы не можете изменить файл, вместо этого запустите его через интерпретатор.
[Пример `echo-script`](../release-pack/echo-script.md) в выпуске
делает именно это, сначала вызывая `bash`.

## Похожие

- [`command`](command.md) — запустить встроенную команду вместо файла
- [Эхо-скрипт](../release-pack/echo-script.md) — запустить скрипт через Bash
- [Встроенные функции](../index.md#built-in-functions) — обе встроенные функции
- [Оболочка](../../concepts/shell.md) — как выполняются команды
