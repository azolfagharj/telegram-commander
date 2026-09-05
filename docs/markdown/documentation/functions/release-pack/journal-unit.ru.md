---
title: Журналы службы
description: Функция `journal-unit` из комплекта показывает последние записи `journalctl` для службы systemd. Установите имя службы и дополнительное количество строк непосредственно на кнопке.
icon: material/text-box-search-outline
---
# :material-text-box-search-outline: Журналы службы

`journal-unit` показывает самые последние журналы `journalctl` для модуля
systemd. Это одна из [входящих в комплект
функций](../index.md#custom-functions), которую вы можете использовать из кнопка
как есть.

- Запускается: `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (обязательно): название устройства, например `nginx.service`
- `lines` (необязательно, по умолчанию `50`): сколько строк показывать.

!!! example "Файл функции"

    ```yaml title="functions/journal-unit.yaml"
    name: journal-unit
    run: "journalctl -u {{.unit}} -n {{.lines}} --no-pager"
    params:
      - name: unit
        type: string
        required: true
        description: Systemd unit name (for example nginx.service)
      - name: lines
        type: string
        required: false
        default: "50"
        description: Number of log lines
    ```

## Добавляем кнопку

!!! example "Читать последние журналы одной службы"

    ```yaml title="Nginx logs button"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`lines` — числовое значение, поэтому кавычки ему не нужны. Не указывайте его,
чтобы использовать значение по умолчанию `50`.

!!! example "Использовать количество строк по умолчанию"

    ```yaml title="SSH logs button"
    - name: SSH logs
      type: button
      function: journal-unit
      unit: "ssh.service"
    ```

## Похожие

- [Заполнители](../write-your-own/placeholders.md) — как заполняется `{{.unit}}`
- [Пользовательские функции](../index.md#custom-functions) — все пять примеров в
  комплекте
