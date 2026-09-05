---
title: Reglas
description: >-
  Comprobaciones que debe superar un archivo de función: nombres admitidos,
  palabras reservadas, nombres únicos, rechazo de errores y valores
  predeterminados.
icon: material/format-list-checks
---

# :material-format-list-checks: Reglas

Cada archivo de función se comprueba al iniciar el bot y al ejecutar
[`validate`](../../cli.md#validate). Si un archivo es incorrecto, nada se
ejecuta, por lo que los errores se detectan antes de publicar el menú.

## Qué archivos se leen

- Solo se leen los archivos terminados en `.yaml` o `.yml`.
- También se leen las subcarpetas de `function_directory`.
- Cada archivo debe describir exactamente una función.

## Nombres

- **`name` es obligatorio** y debe coincidir con `^[A-Za-z0-9._-]+$`: letras,
  dígitos, `.`, `-` y `_`, sin espacios.
- **Los nombres son únicos, sin distinguir mayúsculas.** `Deploy` y `deploy`
  entran en conflicto aunque estén en archivos distintos.
- **Los nombres reservados están prohibidos.** Una función no puede llamarse
  `command` ni `script`; pertenecen a las
  [funciones integradas](../index.md#built-in-functions).

## El comando

- **`run` es obligatorio.** Escriba el comando con
  [variables de reemplazo](placeholders.md) donde deban ir los valores.
- La sintaxis de las variables de reemplazo se comprueba al validar la
  configuración.

## Parámetros

- **Cada parámetro necesita un `name`** que siga la misma regla de caracteres.
- **`type` debe ser** `string`, `int` o `bool`. Si se omite, es `string`. Se
  comprueban los valores y predeterminados declarados como `int` o `bool`.
- Un parámetro no puede llamarse como un ajuste de botón: `name`, `type`,
  `icon`, `id`, `function`, `confirm`, `timeout`, `workdir`, `env`, `columns`
  o `items`. Se permiten `command`, `path` y `args`.
- Cada valor de un botón debe coincidir con un parámetro declarado por su función.

El campo `run` solo puede usar variables de reemplazo de parámetros declarados.
Una variable no declarada es un error.

## Los errores tipográficos son errores

**Las claves desconocidas se rechazan.** Un error como `requird:` en vez de
`required:` detiene la validación en lugar de ignorarse.

!!! warning "Este archivo no se carga"

    ```yaml title="functions/broken.yaml"
    name: broken
    run: "echo {{.args}}"
    params:
      - name: args
        requird: true      # rechazado: clave desconocida
    ```

## Valores obligatorios y predeterminados

- Un **parámetro obligatorio sin valor** hace fallar la validación con un
  mensaje como `required parameter "args" for function "greet" is missing`.
- Un valor presente pero vacío se considera ausente.
- Un **parámetro opcional** usa su `default`. Sin `default`, queda vacío, lo que
  permite omitir partes opcionales del comando.

!!! example "Un parámetro obligatorio y otro opcional"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n {{.args}} {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        default: "200"
        description: Number of lines
    ```

    Un botón que solo define `path` ejecuta
    `tail -n 200 /var/log/app.log`.

## Relacionado

- [Estructura de archivos](file-structure.md) — explicación de todos los campos
- [Variables de reemplazo](placeholders.md) — cómo escribir el comando `run`
- [CLI → validate](../../cli.md#validate) — ejecute usted mismo las comprobaciones
