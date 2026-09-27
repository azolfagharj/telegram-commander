---
title: Reintentos de entrega
description: Una breve caída de red ya no le hace perder el resultado. El bot reintenta el envío con una espera creciente, respeta la espera que pida Telegram por límite de ritmo y mantiene el orden de los mensajes por chat.
icon: material/refresh
---

# :material-refresh: Reintentos de entrega

Si la conexión con Telegram se corta un momento mientras el bot envía un
resultado, el mensaje no se pierde. El bot lo intenta de nuevo por su cuenta,
esperando un poco más cada vez, hasta que el envío se completa.

Cuando el propio Telegram le pide al bot que vaya más despacio, el bot espera
exactamente el tiempo que Telegram indica antes de reintentar, en lugar de
adivinar su propia espera.

Los mensajes a un mismo chat siguen llegando en el orden en que los disparó,
incluso si toca varios botones mientras la conexión tiene problemas.

Si la caída dura demasiado, el bot finalmente deja de intentarlo para ese
resultado. No aparece nada en el chat y no queda nada a medio enviar — el
intento solo se anota en el registro del servidor, nunca se le muestra.

!!! tip "Nada que configurar para una instalación normal"

    Los valores predeterminados absorben por sí solos una breve interrupción
    de red. Cambie cuánto tiempo sigue intentando el bot con
    `delivery_retry_backoff`, `delivery_retry_backoff_max` y
    `delivery_retry_ttl` en su [archivo de configuración](config-file.md).

## Configuración

Para `delivery_retry_backoff`, `delivery_retry_backoff_max` y
`delivery_retry_ttl`, vea [Configuración](../configuration.md).

## Relacionado

- [Confirmación](confirmation.md) — otro momento en el que el bot espera antes de actuar
- [Cómo se conecta el bot](long-polling.md) — la conexión saliente que esto protege
- [Controlar la salida](../output-control.md) — cómo llega un resultado a su chat
