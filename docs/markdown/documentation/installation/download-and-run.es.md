---
title: Ejecutar en la CLI
description: >-
  Ponga en marcha su bot desde la terminal, paso a paso, desde descargar la
  versión y escribir una configuración pequeña hasta tocar el primer botón.
icon: material/console
---

# :material-console: Ejecutar en la CLI

Esta página le guía desde cero hasta un bot en ejecución. No necesita experiencia
previa con el proyecto. Si un término no está claro, consulte las páginas de
[Conceptos](../concepts/config-file.md).

## :material-clipboard-check-outline: Antes de empezar

Necesita dos datos de Telegram:

1. **Un token de bot.** Abra un chat con
   [@BotFather](https://t.me/BotFather), envíe `/newbot`, siga las indicaciones
   y copie el token recibido. Se parece a `123456789:AAExampleTokenValue`.
2. **Su id numérico de usuario.** Es un número, no su `@username`. Si no lo
   conoce, el bot se lo indicará la primera vez que le escriba; consulte el
   [paso 5](#step-5-find-your-user-id-if-needed).

## :material-download: Paso 1: descargar

Descargue el archivo de la versión y extráigalo.

!!! example "Descargar y abrir la carpeta de la versión"

    ```bash title="Descargar y extraer la versión"
    wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
    tar -xzf telegram-commander.tar.gz
    cd telegram-commander
    ```

Dentro de la carpeta encontrará:

- `telegram-commander-linux-amd64` y `telegram-commander-linux-arm64` — el programa, uno para cada tipo de CPU
- `config-examples/` — [configuraciones](../concepts/config-file.md) listas
para usar (consulte [Configuración](../configuration.md))
- `functions/` — [funciones](../concepts/function.md) personalizadas de
ejemplo (consulte [Funciones
personalizadas](../functions/index.md#custom-functions))

## :material-chip: Paso 2: elegir el binario

!!! info "¿Qué archivo corresponde a su máquina?"

    La mayoría de los servidores y PC son `amd64`, también llamado `x86_64`.
    Las placas ARM pequeñas y algunas máquinas virtuales son `arm64`.

    Si no está seguro, ejecute `uname -m`: `x86_64` significa amd64 y
    `aarch64` significa arm64.

=== ":fontawesome-brands-linux: AMD64"

    ```bash title="Conservar el programa amd64"
    mv telegram-commander-linux-amd64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-arm64
    ```

=== ":fontawesome-brands-linux: ARM64"

    ```bash title="Conservar el programa arm64"
    mv telegram-commander-linux-arm64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-amd64
    ```

Ahora tiene un único programa llamado `telegram-commander`.

## :material-file-cog-outline: Paso 3: crear la configuración

Copie el ejemplo mínimo a un archivo de trabajo:

!!! example "Crear una configuración editable"

    ```bash title="Copiar la configuración de ejemplo"
    cp config-examples/config.minimal.yaml ./config.yaml
    ```

Abra `config.yaml` y sustituya dos marcadores:

- `YOUR_BOT_TOKEN` — el token de BotFather
- `YOUR_USER_ID` — su id numérico, o déjelo por ahora y consulte el paso 5

Para conocer todos los ajustes, lea [Configuración](../configuration.md).

## :material-file-check-outline: Paso 4: validar

Compruebe siempre la configuración antes de ejecutar el bot. Así detectará
errores sin iniciarlo.

!!! success "Comprobar que la configuración funciona"

    ```bash title="Validar la configuración"
    ./telegram-commander validate --config config.yaml
    ```

Si muestra `Valid configuration`, todo está listo. En caso contrario, indica
exactamente el problema y su ubicación. Consulte la
[página de la CLI](../cli.md#validate).

## :material-account-search: Paso 5: encontrar su id de usuario (si es necesario) { #step-5-find-your-user-id-if-needed }

Si no conoce su id, escriba solo el token en `config.yaml`, ponga por ahora
cualquier número en `allowed_users` y ejecute el bot:

!!! info "Iniciar una vez para ver su id de usuario"

    ```bash title="Ejecutar el bot para conocer su id"
    ./telegram-commander run --config config.yaml
    ```

Abra Telegram, busque su bot y envíele cualquier mensaje. Como aún no está entre
los [usuarios permitidos](../concepts/allowed-users.md), el bot responde con su
`user_id` y `username`. Copie el id a `allowed_users`, detenga el bot con
`Ctrl+C` y vuelva a ejecutarlo.

Este comportamiento forma parte del control de acceso; consulte
[Configuración → telegram](../configuration.md#telegram).

## :material-play-circle-outline: Paso 6: ejecutar

!!! example "Iniciar el bot en la terminal"

    ```bash title="Iniciar el bot"
    ./telegram-commander run --config config.yaml
    ```

Abra el bot en Telegram y envíe `/start`. Debería ver el menú. Toque un
[botón](../concepts/button.md) para ejecutar su comando.

!!! success "Su bot está activo"

    El menú descrito en `config.yaml` ya está en su chat y cada toque ejecuta
    su comando en esta máquina.

Para mantener el bot en ejecución después de cerrar la sesión del servidor,
configúrelo como servicio. Consulte
[Ejecutar como servicio](run-as-a-service.md).

## :material-map-marker-path: Siguientes pasos

- Añada más [botones](../concepts/button.md) y [categorías](../concepts/category.md): [Menú](../concepts/menu.md)
- Comprenda qué se ejecuta realmente: [Funciones](../functions/index.md)
- Consulte todas las opciones de línea de comandos: [CLI](../cli.md)
