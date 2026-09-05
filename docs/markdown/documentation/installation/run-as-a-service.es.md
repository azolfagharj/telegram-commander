---
title: Ejecutar como servicio
description: >-
  Mantenga el bot en segundo plano con systemd para que se inicie al
  arrancar y vuelva tras un fallo, como root o como usuario normal.
icon: material/server
---

# :material-server: Ejecutar como servicio

Esta página explica cómo ejecutar Telegram Commander como servicio systemd. Para
un primer inicio paso a paso en la terminal, consulte
[Ejecutar en la CLI](download-and-run.md).

## :material-cog-outline: systemd (root)

Cree usted mismo el archivo de unidad, por ejemplo
`/etc/systemd/system/telegram-commander.service`:

!!! warning "Este ejemplo ejecuta todos los botones como root"

    ```ini title="/etc/systemd/system/telegram-commander.service"
    [Unit]
    Description=telegram-commander Telegram bot
    After=network-online.target
    Wants=network-online.target

    [Service]
    Type=simple
    ExecStart=/path/to/telegram-commander run --config /path/to/configfile.yaml
    Restart=on-failure
    RestartSec=5

    [Install]
    WantedBy=multi-user.target
    ```

    Sin una línea `User=`, el servicio se inicia como root, por lo que todos los
    comandos del menú tienen derechos completos en la máquina. Añada
    `User=someone` a la sección `[Service]` si sus botones no necesitan tantos
    privilegios.

Sustituya las rutas de ejemplo y después:

!!! example "Cargar e iniciar el servicio"

    ```bash title="Activar, iniciar y observar el servicio"
    sudo systemctl daemon-reload
    sudo systemctl enable --now telegram-commander
    sudo systemctl status telegram-commander
    sudo journalctl -u telegram-commander -f
    ```

!!! info "Los cambios de configuración requieren un reinicio"

    El servicio mantiene el bot en ejecución, pero lee el
    [archivo de configuración](../concepts/config-file.md) una vez al iniciarse.
    Ejecute `sudo systemctl restart telegram-commander` después de modificarlo.

## Páginas relacionadas

- [Ejecutar en la CLI](download-and-run.md) — primer inicio explicado
- [Configuración](../configuration.md) — el archivo de configuración
- [CLI](../cli.md) — `run`, `validate` y otros comandos
