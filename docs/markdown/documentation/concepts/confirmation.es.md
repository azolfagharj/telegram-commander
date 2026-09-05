---
title: Confirmación
description: >-
  Haga que el bot pregunte antes de ejecutar un botón. Active confirm para
  acciones irreversibles y defina cuánto dura la pregunta Sí o Cancelar.
icon: material/help-circle-outline
---

# :material-help-circle-outline: Confirmación

Un paso opcional «¿Está seguro?» antes de ejecutar un [botón](button.md), que se
activa con `confirm: true`. Es útil para acciones destructivas.

El bot pregunta (`Confirm: … ?`) y muestra Sí / Cancelar, además de Inicio (y
Atrás dentro de una categoría). Si Ejecutar comando está activado, ese botón
también permanece en la pantalla de confirmación. Nada se ejecuta hasta que
toque Sí; Cancelar, Inicio y Atrás dejan el comando sin ejecutar.

La pregunta caduca tras un tiempo (5 minutos de forma predeterminada). Cámbielo
con `confirm_ttl` en el [archivo de configuración](config-file.md). Si caduca,
toque de nuevo el botón para obtener otra.

!!! warning "Úselo para cualquier acción irreversible"

    Una tecla del menú está a un solo toque en el teléfono y todas las personas
    autorizadas ven el mismo menú. Añada `confirm: true` a los botones que
    detienen un servicio, eliminan datos o reinician la máquina.

!!! example "Un botón que pregunta antes de ejecutarse"

    ```yaml title="Botón Stop nginx"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## Configuración

Para `confirm` en los botones y el valor global `confirm_ttl`, consulte
[Configuración](../configuration.md).

## Relacionado

- [Botón](button.md) — el nodo que puede requerir confirmación
- [Menú → Confirmación](menu.md#confirmation) — ejemplos y comportamiento
