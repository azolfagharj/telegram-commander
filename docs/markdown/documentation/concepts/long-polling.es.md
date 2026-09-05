---
title: Cómo se conecta el bot
description: >-
  El bot se conecta a Telegram, por lo que nunca abre un puerto en su
  servidor. Configure un proxy si Telegram está bloqueado en su red.
icon: material/lan-disconnect
---

# :material-lan-disconnect: Cómo se conecta el bot

El bot se conecta a Telegram. Usted no abre ningún puerto en el servidor.

!!! tip "Nada que exponer a Internet"

    Como la conexión siempre es saliente, no hay panel web que proteger, regla
    de firewall entrante que añadir ni dirección pública que alguien pueda
    encontrar. La máquina puede estar detrás de un router doméstico o un
    firewall estricto y el bot seguirá funcionando.

Si Telegram está bloqueado en su red, configure un proxy bajo `telegram` en el
[archivo de configuración](config-file.md).

## Configuración

Para `api`, `proxy` e `insecure` en los ajustes de Telegram, consulte
[Configuración → telegram](../configuration.md#telegram).

## Relacionado

- [Usuarios permitidos](allowed-users.md) — quién puede usar el bot
- [CLI → run](../cli.md#run) — inicie el bot en primer plano
- [Ejecutar como servicio](../installation/run-as-a-service.md) — mantenga el bot en ejecución
