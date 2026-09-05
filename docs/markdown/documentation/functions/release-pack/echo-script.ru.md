---
title: Эхо-скрипт
description: Функция `echo-script` из комплекта запускает файл сценария через Bash, поэтому для сценария не требуется бит выполнения. Он принимает путь и аргументы от кнопки.
icon: material/bash
---
# :material-bash: Эхо-скрипт

`echo-script` запускает сценарий через `bash`, поэтому сам файл сценария не
нужен бит выполнения. Это один из
[функций из комплекта](../index.md#custom-functions), которые вы можете использовать с кнопки как
есть.

- Запускается: `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (обязательно): путь к файлу скрипта.
- `args` (необязательно): дополнительные аргументы.

!!! example "Файл функции"

    ```yaml title="functions/echo-script.yaml"
    name: echo-script
    run: "bash {{.path}}{{if .args}} {{.args}}{{end}}"
    params:
      - name: path
        type: string
        required: true
        description: Path to the script file
      - name: args
        type: string
        required: false
        description: Optional script arguments
    ```

Напишите названия обоих параметров прямо на кнопке.

!!! example "Запуск сценария через Bash"

    ```yaml title="Run cleanup button"
    - name: Run cleanup
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

Эта кнопка запускает `bash /opt/scripts/cleanup.sh --verbose`. Если не указывать
`args`, кнопка запускает `bash /opt/scripts/cleanup.sh`, потому что часть `{{if
.args}}` пропускается, если значение пустое.

## Похожие

- [`script`](../built-in/script.md) — запустить исполняемый скрипт напрямую
- [Заполнители](../write-your-own/placeholders.md) — как работает `{{if .args}}`
- [Пользовательские функции](../index.md#custom-functions) — все пять примеров в
  комплекте
