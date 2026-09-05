---
title: Función
description: >-
  Una función convierte los valores de un botón en un comando de shell. Use
  command y script integradas o añada las suyas como archivos YAML.
icon: material/function
---

# :material-function: Función

Lo que se ejecuta al tocar un [botón](button.md). Una [función](function.md)
recibe algunos [parámetros](parameter.md) y los convierte en un comando de shell.

Hay dos tipos:

- Las **funciones integradas** vienen con el programa (`command`, `script`).
  Usted no las crea.
- Las **funciones personalizadas** son funciones adicionales definidas como
  archivos YAML en `function_directory`.

Al tocar un botón, el bot busca la función, completa los valores del botón y
ejecuta el resultado en el [shell](shell.md).

!!! example "Un botón que usa la función integrada `command`"

    ```yaml title="Botón Uptime"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

## Configuración

Para `function_directory` y los campos de botón que hacen referencia a
funciones, consulte [Configuración](../configuration.md).

## Relacionado

- [Parámetro](parameter.md) — valores con nombre que necesita una función
- [Funciones](../functions/index.md#two-kinds-of-function) — funciones integradas y personalizadas en detalle
- [Menú](menu.md) — cómo hacen referencia los botones a las funciones
