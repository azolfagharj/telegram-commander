---
title: Usuarios permitidos
description: >-
  Elija qué cuentas de Telegram pueden abrir el menú del bot. Añada personas
  por id numérico o nombre de usuario y vea qué reciben las demás.
icon: material/shield-lock
---

# :material-shield-lock: Usuarios permitidos

La lista de cuentas de Telegram autorizadas para usar el bot, mediante
`user_id` numérico o `@username`. El acceso se rechaza a todas las demás.

Los usuarios no autorizados reciben un mensaje con su `user_id` y `username`
para que puedan solicitar acceso a un administrador. Así puede encontrar
también su propio id la primera vez; consulte
[Ejecutar en la CLI → Paso 5](../installation/download-and-run.md#step-5-find-your-user-id-if-needed).

Ese mensaje también sirve para añadir a alguien. Pida a la persona que envíe
cualquier mensaje al bot, que copie el id de la respuesta, añádalo a
`allowed_users` y reinicie el bot. La próxima vez recibirá el menú en vez del
rechazo.

!!! warning "Todas las personas de la lista comparten el mismo menú"

    No hay permisos por usuario. Cualquier persona añadida puede tocar todos los
    botones definidos, así que limite la lista a personas a quienes confiaría
    la propia máquina.

!!! example "Dos cuentas autorizadas para usar el bot"

    ```yaml title="Sección telegram de config.yaml"
    telegram:
      allowed_users:
        - "123456789"
        - "@alice"
    ```

## Configuración

Para `allowed_users` y otros ajustes de Telegram, consulte
[Configuración → telegram](../configuration.md#telegram).

## Relacionado

- [Archivo de configuración](config-file.md) — contiene la sección `telegram`
- [Cómo se conecta el bot](long-polling.md) — no hay que abrir un puerto entrante
