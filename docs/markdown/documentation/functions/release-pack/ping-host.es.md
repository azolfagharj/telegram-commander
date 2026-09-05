---
title: Ping a un host
description: >-
  La función incluida ping-host hace ping a un host varias veces. Defina
  host y el número opcional count directamente en el botón.
icon: material/access-point-network
---

# :material-access-point-network: Ping a un host

`ping-host` hace ping a un host varias veces y devuelve el resultado. Es una de
las [funciones incluidas](../index.md#custom-functions) que puede usar sin
cambios.

- Ejecuta: `ping -c {{.count}} {{.host}}`
- `host` (obligatorio): nombre de host o IP
- `count` (opcional, valor predeterminado `4`): número de paquetes

!!! example "El archivo de la función"

    ```yaml title="functions/ping-host.yaml"
    name: ping-host
    run: "ping -c {{.count}} {{.host}}"
    params:
      - name: host
        type: string
        required: true
        description: Hostname or IP
      - name: count
        type: string
        required: false
        default: "4"
        description: Number of ping packets
    ```

## Añadir un botón

!!! example "Hacer ping tres veces a un host fijo"

    ```yaml title="Botón Ping gateway"
    - name: Ping gateway
      type: button
      function: ping-host
      host: "192.168.1.1"
      count: 3
    ```

`count` es numérico, por lo que no necesita comillas. También puede omitirlo
para usar el valor predeterminado `4`:

!!! example "Usar el número predeterminado"

    ```yaml title="Botón Ping DNS"
    - name: Ping DNS
      type: button
      function: ping-host
      host: "1.1.1.1"
    ```

## Relacionado

- [Reglas](../write-your-own/rules.md) — cómo funcionan los valores predeterminados y obligatorios
- [Funciones personalizadas](../index.md#custom-functions) — los cinco ejemplos incluidos
