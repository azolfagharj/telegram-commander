---
title: Shell
description: >-
  Cada comando se ejecuta mediante un shell, /bin/bash de forma
  predeterminada, por lo que funcionan las tuberías y redirecciones.
  Cámbielo en la raíz de su configuración.
icon: material/console-line
---

# :material-console-line: Shell

El programa que ejecuta su comando, `/bin/bash` de forma predeterminada. Los
comandos se ejecutan como `bash -c "your command"`, por lo que funcionan las
tuberías y redirecciones.

Puede cambiar el shell globalmente con la clave `shell` del
[archivo de configuración](config-file.md).

!!! example "Elegir el shell"

    ```yaml title="Raíz de config.yaml"
    shell: /bin/bash
    ```

Cada [función](function.md) ejecuta su comando mediante este shell.

## Configuración

Para el campo `shell` y las opciones raíz relacionadas, consulte
[Configuración → Campos raíz](../configuration.md#root-fields).

## Relacionado

- [Función](function.md) — produce el comando que ejecuta el shell
- [Archivo de configuración](config-file.md) — donde se define `shell`
