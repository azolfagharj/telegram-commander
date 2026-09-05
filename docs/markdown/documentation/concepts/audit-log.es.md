---
title: Registro de auditoría
description: >-
  Conserve un registro de cada comando que ejecuta el bot, quién tocó el
  botón, qué botón fue, el código de salida y cuánto tardó.
icon: material/history
---

# :material-history: Registro de auditoría

Un flujo de registro independiente que guarda cada comando ejecutado: quién lo
ejecutó, qué [botón](button.md), el código de salida y cuánto tardó. Se configura
bajo `logging` en su [archivo de configuración](config-file.md).

Este es el registro que puede consultar después para saber quién reinició un
servicio a las tres de la mañana o si el botón de copia de seguridad de anoche
terminó realmente. Como es un flujo propio, puede conservarlo aunque el registro
normal de la aplicación ya haya rotado.

Cada ejecución añade una línea y, entre todas, responden a cuatro preguntas:

| Pregunta | Lo que indica la línea |
|---|---|
| ¿Quién? | La cuenta de Telegram que tocó el botón |
| ¿Qué? | El botón que usó |
| ¿Funcionó? | El código de salida devuelto por el comando |
| ¿Cuánto tardó? | El tiempo que tardó el comando en terminar |

!!! example "Escribir el registro de auditoría en su propio archivo"

    ```yaml title="Sección logging de config.yaml"
    logging:
      logs:
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

## Configuración

Para el esquema completo de `logging` y las salidas admitidas, consulte
[Configuración → logging](../configuration.md#logging).

## Relacionado

- [Archivo de configuración](config-file.md) — donde se define el registro
- [Función](function.md) — lo que se ejecuta al tocar un botón
