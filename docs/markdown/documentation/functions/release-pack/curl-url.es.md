---
title: URL Curl
description: >-
  La función incluida curl-url obtiene una URL con curl. Añada la URL
  directamente a un botón con la clave url.
icon: material/web
---

# :material-web: URL Curl

`curl-url` obtiene una URL con `curl`. Falla ante errores HTTP (`-f`) y abandona
después de 30 segundos. Es una de las
[funciones incluidas](../index.md#custom-functions) que puede usar sin cambios.

- Ejecuta: `curl -fsSL --max-time 30 {{.url}}`
- `url` (obligatorio): la URL que se solicitará

!!! example "El archivo de la función"

    ```yaml title="functions/curl-url.yaml"
    name: curl-url
    run: "curl -fsSL --max-time 30 {{.url}}"
    params:
      - name: url
        type: string
        required: true
        description: URL to request
    ```

## Añadir un botón

Escriba `url` directamente en el botón. Su nombre coincide con el parámetro
declarado por la función.

!!! example "Comprobar un punto de conexión"

    ```yaml title="Botón Check API"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

El botón ejecuta `curl -fsSL --max-time 30 https://example.com/health`.
[`validate`](../../cli.md#validate) informa de un error si falta `url`.

## Relacionado

- [Parámetros](../../concepts/parameter.md) — cómo se comprueban los valores del botón
- [Funciones personalizadas](../index.md#custom-functions) — los cinco ejemplos incluidos
