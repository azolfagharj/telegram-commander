---
title: Variables de reemplazo
description: >-
  Cómo escribir el comando run de una función con variables de reemplazo,
  incluidas
  partes opcionales que solo aparecen cuando un botón proporciona un valor.
icon: material/code-braces
---

# :material-code-braces: Variables de reemplazo

El campo `run` es una plantilla pequeña. Las variables de reemplazo entre llaves
dobles se sustituyen por valores de parámetros antes de ejecutar el comando.

Dos patrones cubren casi todos los casos:

- `{{.name}}` inserta el valor del parámetro llamado `name`.
- `{{if .name}} ... {{end}}` incluye la parte central solo si `name` tiene valor.

## Insertar un valor

!!! example "Un valor en medio de un comando"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
    ```

Un botón con `path: "/var/log/app.log"` ejecuta
`tail -n 200 /var/log/app.log`.

## Partes opcionales

Un valor vacío hace que `{{if .name}}` sea falso, por lo que desaparece todo el
bloque. Así puede añadir una opción solo cuando se necesite.

!!! example "Crear una parte opcional con una variable de reemplazo"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}{{if .args}} | grep -- {{.args}}{{end}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        description: Optional text to filter for
    ```

Dos botones que la usan, uno con filtro y otro sin él:

!!! example "Proporcionar valores desde un botón"

    ```yaml title="Botones de registro de la aplicación"
    - name: App log
      type: button
      function: tail-log
      path: "/var/log/app.log"

    - name: App errors
      type: button
      function: tail-log
      path: "/var/log/app.log"
      args: "ERROR"
    ```

**App log** ejecuta `tail -n 200 /var/log/app.log` y **App errors** ejecuta
`tail -n 200 /var/log/app.log | grep -- ERROR`.

También puede elegir entre dos formas con
`{{if .args}} ... {{else}} ... {{end}}`.

## Qué nombres puede usar

Una variable de reemplazo puede usar cualquier parámetro declarado en `params`.
Su valor es el proporcionado por el botón o el `default` del parámetro.

!!! warning "Un nombre no declarado falla durante la ejecución"

    [`validate`](../../cli.md#validate) solo comprueba la sintaxis de las
    variables de reemplazo, no sus nombres. Si `run` menciona `{{.uri}}` pero no hay un
    parámetro `uri`, la configuración se valida y el botón muestra un error.

!!! warning "Los valores se insertan como texto sin formato"

    No se entrecomillan ni escapan automáticamente. Un valor con espacios o
    caracteres de shell pasa a formar parte del comando tal como está escrito;
    prefiera valores fijos y añada
    [`confirm: true`](../../concepts/confirmation.md) a las acciones destructivas.

## Relacionado

- [Estructura de archivos](file-structure.md) — dónde están `run` y `params`
- [Reglas](rules.md) — lo que acepta el cargador
- [Guía paso a paso](step-by-step.md) — pruebe el proceso completo
