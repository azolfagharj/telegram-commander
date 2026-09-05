---
title: Referencia de la CLI
description: >-
  Todas las opciones de línea de comandos de Telegram Commander, desde
  ejecutar el bot hasta comprobar y dar formato a la configuración, y el
  indicador --config.
icon: material/console-line
---

# :material-console-line: Referencia de la CLI

!!! info "Los comandos usan esta forma"

    ```text title="Sintaxis del comando"
    telegram-commander <command> [flags]
    ```

`--config` / `-c` es **obligatorio** para los comandos que cargan la
configuración. La ruta puede ser relativa al directorio de trabajo actual o
absoluta. El contenido del archivo se documenta en
[Configuración](configuration.md).

Si acaba de empezar, [Ejecutar en la CLI](installation/download-and-run.md)
muestra estos comandos en el orden en que los usará.

## :material-format-list-bulleted-square: Comandos

### `run` { #run }

Ejecuta el bot en primer plano.

!!! info "Iniciar el bot en primer plano"

    ```bash title="Ejecutar el bot"
    telegram-commander run --config /path/to/configfile.yaml
    ```

!!! note "Los cambios de configuración requieren un reinicio"

    El bot lee la configuración una vez al iniciarse. Después de modificar el
    archivo, reinicie el proceso, por ejemplo con
    `systemctl restart telegram-commander`.

### `validate` { #validate }

Valida sin conexión la configuración, las funciones y las referencias de botones.

!!! info "Comprobar la configuración sin conexión o en línea"

    ```bash title="Validar la configuración"
    telegram-commander validate --config /path/to/configfile.yaml
    telegram-commander validate --config /path/to/configfile.yaml --online
    ```

!!! note "`--online` necesita Internet"

    Con `--online`, la comprobación también pregunta a Telegram si funciona el
    token del bot, por lo que la máquina debe poder conectarse a Telegram.

### `version`

Muestra la versión del programa.

### `fmt`

Muestra un archivo de configuración YAML con un formato uniforme.

!!! info "Mostrar o guardar YAML con formato"

    ```bash title="Dar formato a un archivo de configuración"
    telegram-commander fmt --config /path/to/configfile.yaml
    telegram-commander fmt --config /path/to/configfile.yaml -w
    ```

### `environ`

Muestra las variables de entorno del proceso, útiles para depurar unidades de servicio.

### `list-functions` { #list-functions }

Enumera las funciones integradas y personalizadas cargadas. Úselo para confirmar
que se encontraron sus archivos. Consulte [Funciones](functions/index.md).

!!! info "Mostrar todas las funciones disponibles"

    ```bash title="Enumerar las funciones disponibles"
    telegram-commander list-functions --config /path/to/configfile.yaml
    ```

### `completion`

Genera scripts de completado del shell:

!!! info "Elegir el shell que usa"

    ```bash title="Generar un script de completado"
    telegram-commander completion bash
    telegram-commander completion zsh
    telegram-commander completion fish
    telegram-commander completion powershell
    ```

### `manpage`

Escribe una página de manual en stdout.

## Páginas relacionadas

- [Configuración](configuration.md) — el archivo que se pasa con `--config`
- [Funciones personalizadas](functions/index.md#custom-functions) — lo que muestra `list-functions`
- [Ejecutar como servicio](installation/run-as-a-service.md) — ejecute `run` con systemd
