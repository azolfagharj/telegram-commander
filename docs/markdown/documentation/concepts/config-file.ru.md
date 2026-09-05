---
icon: material/file-document-outline
title: Конфигурационный файл
description: Один файл YAML содержит токен вашего бота, разрешенных пользователей и меню кнопок. Передайте его с --config для запуска, проверки, fmt и функций списка.
---
# :material-file-document-outline: Конфигурационный файл

Один файл YAML, описывающий все: токен вашего бота, который может использовать
бота и меню [кнопки](button.md). Вы передаете его с помощью `--config` в
команды, которые его читают: `run`, `validate`, `fmt` и `list-functions`. Другое
команды, такие как `version` и `completion`, не принимают файл конфигурации.

!!! example "Рабочий конфиг одной кнопкой"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

## Похожие

- [Конфигурация](../configuration.md) — каждое поле, значение по умолчанию и
  правило проверки.
- [Меню](menu.md) — кнопка и дерево категорий
- [CLI](../cli.md) — передать файл с `--config`
- [Запуск в CLI](../installation/download-and-run.md) — создайте свою первую
  конфигурацию
