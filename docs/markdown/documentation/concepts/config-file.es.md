---
icon: material/file-document-outline
title: Archivo de configuración
description: >-
  Un archivo YAML contiene el token del bot, los usuarios permitidos y el
  menú de botones. Páselo con --config a run, validate, fmt y
  list-functions.
---

# :material-file-document-outline: Archivo de configuración

Un único archivo YAML describe todo: el token del bot, quién puede usarlo y el
menú de [botones](button.md). Se pasa con `--config` a los comandos que lo leen:
`run`, `validate`, `fmt` y `list-functions`. Otros comandos, como `version` y
`completion`, no reciben un archivo de configuración.

!!! example "Una configuración funcional con un botón"

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

## Relacionado

- [Configuración](../configuration.md) — todos los campos, valores predeterminados y reglas de validación
- [Menú](menu.md) — el árbol de botones y categorías
- [CLI](../cli.md) — pase el archivo con `--config`
- [Ejecutar en la CLI](../installation/download-and-run.md) — cree su primera configuración
