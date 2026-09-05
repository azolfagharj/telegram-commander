---
title: command
description: >-
  La función integrada command ejecuta un único comando de shell exactamente
  como está escrito en el botón, incluidas tuberías, redirecciones y
  comandos encadenados.
icon: material/console
---

# :material-console: `command`

`command` es una [función integrada](../index.md#built-in-functions). Ejecuta un
único comando de shell exactamente como usted lo escribió en el botón.

| Parámetro | Obligatorio | Predeterminado | Significado |
|-----------|-------------|----------------|-------------|
| `command` | sí | — | El comando de shell que se ejecutará |

!!! example "Ejecutar un comando tal como está escrito"

    ```yaml title="Botón Show kernel"
    - name: Show kernel
      type: button
      function: command
      command: "uname -a"
    ```

## Tuberías y comandos encadenados

Los comandos se ejecutan mediante
[`/bin/bash -c`](../../concepts/shell.md), por lo que las tuberías,
redirecciones y `&&` funcionan dentro de un único valor `command`.

!!! example "Las tuberías funcionan dentro de los comandos"

    ```yaml title="Botón Top processes"
    - name: Top processes
      type: button
      function: command
      command: "ps aux --sort=-%mem | head -n 10"
    ```

!!! tip "Use una función personalizada para comandos repetidos"

    Si varios botones repiten la misma forma de comando, cree una función
    personalizada y asigne a cada botón solo los valores que cambian. Los
    nombres de parámetros personalizados, como `url` o `host`, se pueden
    escribir directamente en esos botones.

## Relacionado

- [`script`](script.md) — ejecute un archivo de script en vez de un comando en línea
- [Funciones integradas](../index.md#built-in-functions) — las dos funciones integradas
- [Shell](../../concepts/shell.md) — cómo se ejecutan los comandos
- [Confirmación](../../concepts/confirmation.md) — pregunte antes de un comando peligroso
