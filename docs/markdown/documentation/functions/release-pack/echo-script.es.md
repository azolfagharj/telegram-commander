---
title: Script de eco
description: >-
  La función incluida echo-script ejecuta un archivo mediante Bash, por lo
  que no necesita permiso de ejecución. Recibe path y args de un botón.
icon: material/bash
---

# :material-bash: Script de eco

`echo-script` ejecuta un script mediante `bash`, por lo que el archivo no
necesita permiso de ejecución. Es una de las
[funciones incluidas](../index.md#custom-functions) que puede usar sin cambios.

- Ejecuta: `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (obligatorio): ruta del archivo de script
- `args` (opcional): argumentos adicionales

!!! example "El archivo de la función"

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

Escriba ambos nombres de parámetro directamente en el botón.

!!! example "Ejecutar un script mediante Bash"

    ```yaml title="Botón Run cleanup"
    - name: Run cleanup
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

Ese botón ejecuta `bash /opt/scripts/cleanup.sh --verbose`. Si omite `args`,
ejecuta `bash /opt/scripts/cleanup.sh`, porque la parte `{{if .args}}` se omite
cuando el valor está vacío.

## Relacionado

- [`script`](../built-in/script.md) — ejecute directamente un script ejecutable
- [Variables de reemplazo](../write-your-own/placeholders.md) — cómo funciona `{{if .args}}`
- [Funciones personalizadas](../index.md#custom-functions) — los cinco ejemplos incluidos
