---
title: Estructura de archivos
description: >-
  Explicación de todos los campos de un archivo de función, desde name y run
  hasta la lista params con type, required, default y description.
icon: material/file-tree
---

# :material-file-tree: Estructura de archivos

Cada archivo describe una función. Colóquelo en cualquier parte de su
[`function_directory`](../../configuration.md#function_directory-rules),
incluso en una subcarpeta, y use la extensión `.yaml` o `.yml`. El nombre del
archivo no importa: el nombre de la función procede del campo `name`.

!!! example "Definir una función personalizada por archivo"

    ```yaml title="Archivo de función personalizada"
    name: my-function          # obligatorio, debe ser único y no estar reservado
    run: "echo {{.args}}"      # obligatorio, el comando que se ejecutará
    params:                    # lista opcional de parámetros
      - name: args             # obligatorio para cada parámetro
        type: string           # opcional: string (predeterminado), int o bool
        required: true         # opcional, false de forma predeterminada
        default: ""            # valor opcional si no se proporciona
        description: Some text  # nota opcional para usted
    ```

## Campos de nivel superior

| Campo | Obligatorio | Significado |
|-------|-------------|-------------|
| `name` | sí | El nombre que usan los botones en su campo `function` |
| `run` | sí | El comando, con [variables de reemplazo](placeholders.md) para los valores |
| `params` | no | Lista de valores con nombre que acepta la función |

Una función sin ningún `params` es válida: `run` es entonces un comando fijo.

## Campos de los parámetros

| Campo | Obligatorio | Predeterminado | Significado |
|-------|-------------|----------------|-------------|
| `name` | sí | — | El nombre usado en `{{.name}}` dentro de `run` |
| `type` | no | `string` | Tipo del valor: `string`, `int` o `bool` |
| `required` | no | `false` | Un botón debe proporcionar un valor |
| `default` | no | vacío | Se usa si no se proporciona ningún valor |
| `description` | no | vacío | Nota para usted, no visible en Telegram |

!!! info "Se comprueba `type`"

    Los valores y predeterminados declarados como `int` deben contener un
    entero. Los declarados como `bool` deben contener un booleano válido. Los
    valores inválidos hacen que [`validate`](../../cli.md#validate) falle.

!!! info "Las claves del botón coinciden con los nombres de parámetros"

    Escriba cada parámetro directamente en un botón con el mismo nombre.
    `url`, `host` y `lines` funcionan junto a `command`, `path` y `args`.
    Consulte [Pasar valores desde un botón](../index.md#passing-values-from-a-button).

## Organización de la carpeta

Puede organizar la carpeta como prefiera:

!!! example "También se leen las subcarpetas"

    ```text title="functions/"
    functions/
      disk.yaml
      logs/
        nginx.yaml
        app.yml
    ```

Se cargan los tres archivos. Se omiten los archivos con otra extensión.

## Relacionado

- [Reglas](rules.md) — lo que rechaza el cargador
- [Variables de reemplazo](placeholders.md) — cómo escribir el comando `run`
- [Guía paso a paso](step-by-step.md) — cree su primera función
