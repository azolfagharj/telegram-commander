---
title: Unidad de journal
description: >-
  La función incluida journal-unit muestra registros recientes de journalctl
  para una unidad systemd. Defina unit y el número opcional de líneas en el
  botón.
icon: material/text-box-search-outline
---

# :material-text-box-search-outline: Unidad de journal

`journal-unit` muestra los registros más recientes de `journalctl` para una
unidad systemd. Es una de las
[funciones incluidas](../index.md#custom-functions) que puede usar sin cambios.

- Ejecuta: `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (obligatorio): nombre de la unidad, por ejemplo `nginx.service`
- `lines` (opcional, valor predeterminado `50`): líneas que se mostrarán

!!! example "El archivo de la función"

    ```yaml title="functions/journal-unit.yaml"
    name: journal-unit
    run: "journalctl -u {{.unit}} -n {{.lines}} --no-pager"
    params:
      - name: unit
        type: string
        required: true
        description: Systemd unit name (for example nginx.service)
      - name: lines
        type: string
        required: false
        default: "50"
        description: Number of log lines
    ```

## Añadir un botón

!!! example "Leer registros recientes de un servicio"

    ```yaml title="Botón Nginx logs"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`lines` es numérico, por lo que no necesita comillas. Omítalo para usar el valor
predeterminado `50`.

!!! example "Usar el número predeterminado de líneas"

    ```yaml title="Botón SSH logs"
    - name: SSH logs
      type: button
      function: journal-unit
      unit: "ssh.service"
    ```

## Relacionado

- [Variables de reemplazo](../write-your-own/placeholders.md) — cómo se completa `{{.unit}}`
- [Funciones personalizadas](../index.md#custom-functions) — los cinco ejemplos incluidos
