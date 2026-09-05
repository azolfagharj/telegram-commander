---
title: Ruta de disco
description: >-
  La función incluida disk-path muestra el uso del disco con df -h. Su único
  parámetro es path y usa el sistema de archivos raíz si se omite.
icon: material/harddisk
---

# :material-harddisk: Ruta de disco

`disk-path` muestra el uso del disco para una ruta con `df -h`. Es una de las
[funciones incluidas](../index.md#custom-functions) que puede usar sin cambios.

- Ejecuta: `df -h {{.path}}`
- `path` (opcional, valor predeterminado `/`): ruta que se comprobará

!!! example "El archivo de la función"

    ```yaml title="functions/disk-path.yaml"
    name: disk-path
    run: "df -h {{.path}}"
    params:
      - name: path
        type: string
        required: false
        default: "/"
        description: Filesystem path to check
    ```

Como `path` es opcional y tiene un valor predeterminado, un botón puede omitirlo
y seguir funcionando.

!!! example "Usar la ruta predeterminada o elegir una"

    ```yaml title="Botones de uso del disco"
    - name: Disk (root)
      type: button
      function: disk-path        # sin campo path: usa "/" de forma predeterminada
    - name: Disk (var)
      type: button
      function: disk-path
      path: "/var"
    ```

El primer botón ejecuta `df -h /` y el segundo ejecuta `df -h /var`.

## Relacionado

- [Reglas](../write-your-own/rules.md) — cómo funcionan los valores predeterminados y obligatorios
- [Funciones personalizadas](../index.md#custom-functions) — los cinco ejemplos incluidos
- [`command`](../built-in/command.md) — escriba un comando completo
