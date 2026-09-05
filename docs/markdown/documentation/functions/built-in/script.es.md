---
title: script
description: >-
  La función integrada script ejecuta un archivo de script del servidor y
  añade argumentos opcionales tras la ruta. El script debe ser ejecutable.
icon: material/script-text
---

# :material-script-text: `script`

`script` es una [función integrada](../index.md#built-in-functions). Ejecuta un
archivo de script con argumentos opcionales.

| Parámetro | Obligatorio | Predeterminado | Significado |
|-----------|-------------|----------------|-------------|
| `path` | sí | — | Ruta del script |
| `args` | no | vacío | Argumentos añadidos después de la ruta |

!!! example "Ejecutar un script con argumentos"

    ```yaml title="Botón Nightly report"
    - name: Nightly report
      type: button
      function: script
      path: "/usr/local/bin/report.sh"
      args: "--today"
    ```

Esto ejecuta `/usr/local/bin/report.sh --today`. Si omite `args`, solo ejecuta
`/usr/local/bin/report.sh`.

## El script debe ser ejecutable

La ruta se ejecuta directamente, por lo que el archivo necesita permiso de
ejecución:

!!! tip "Hacer ejecutable el script una vez"

    ```bash title="Permitir que se ejecute el archivo"
    chmod +x /usr/local/bin/report.sh
    ```

Si no puede cambiar el archivo, ejecútelo mediante un intérprete. El ejemplo
[`echo-script`](../release-pack/echo-script.md) del paquete de la versión hace
exactamente eso al llamar primero a `bash`.

## Relacionado

- [`command`](command.md) — ejecute un comando en línea en vez de un archivo
- [Script de eco](../release-pack/echo-script.md) — ejecute un script mediante Bash
- [Funciones integradas](../index.md#built-in-functions) — las dos funciones integradas
- [Shell](../../concepts/shell.md) — cómo se ejecutan los comandos
