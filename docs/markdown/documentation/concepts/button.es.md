---
title: Botón
description: >-
  Un elemento del menú que ejecuta una función. Escriba cada valor
  directamente en el botón con el nombre de parámetro declarado por la
  función.
icon: material/gesture-tap-button
---

# :material-gesture-tap-button: Botón

Un elemento que se puede tocar en el menú de Telegram. Tiene un `name` y, al
tocarlo, ejecuta una [función](function.md) en el servidor y devuelve el
resultado al chat.

Los botones y las [categorías](category.md) forman el árbol del [menú](menu.md)
bajo la clave raíz `menu` del
[archivo de configuración](config-file.md). Un botón realiza una acción; una
categoría solo abre un submenú.

## Aspecto de un botón

Los botones aparecen en el teclado bajo el cuadro de mensajes, dos por fila de
forma predeterminada. El nombre que elija será el texto de la tecla, así que
procure que sea suficientemente breve para leerlo en un teléfono.
Consulte [Menú → Aspecto del menú](menu.md#how-the-telegram-menu-looks).

## :material-format-list-checks: Partes de un botón

!!! example "Cada línea controla una parte del botón"

    ```yaml title="Un botón con todas sus partes"
    - name: Restart nginx # (1)!
      type: button # (2)!
      icon: "🔄" # (3)!
      function: command # (4)!
      command: "systemctl restart nginx" # (5)!
      confirm: true # (6)!
    ```

    1.  El texto de la tecla en Telegram. Debe ser único entre los elementos
        vecinos del mismo menú.
    2.  Aquí siempre se usa `button`. Use `category` si quiere un submenú.
    3.  Emoji opcional que aparece antes del nombre. Es solo decorativo y no
        cambia lo que se ejecuta.
    4.  La [función](function.md) que se usará. `command` es la función
        integrada que ejecuta un comando de shell.
    5.  Lo que ejecuta la función `command`. Puede escribir aquí cualquier
        comando válido en una terminal.
    6.  Es opcional. Pregunta «¿Está seguro?» antes de ejecutar. Omítalo para
        acciones que solo consultan información.

## Qué ocurre al tocarlo

1. El bot publica una breve línea **Running** para indicar que ha empezado.
2. El comando se ejecuta en la máquina donde funciona el bot.
3. La salida vuelve como bloque de código, con el código de salida y la
   duración. Una salida larga llega en varios mensajes consecutivos.
4. Usted permanece en el mismo menú, por lo que **Atrás** sigue saliendo de esa
   categoría.

## :material-code-braces: Botones habituales

=== "Comprobar algo"

    ```yaml title="Botón Uptime"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

=== "Gestionar un servicio"

    ```yaml title="Botón Restart nginx"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

=== "Leer un registro"

    ```yaml title="Botón Nginx log"
    - name: Nginx log
      type: button
      function: command
      command: "journalctl -u nginx -n 50 --no-pager"
    ```

=== "Ejecutar un script"

    ```yaml title="Botón Nightly backup"
    - name: Nightly backup
      type: button
      function: script
      path: "/usr/local/bin/backup.sh"
    ```

=== "Algo destructivo"

    ```yaml title="Botón Stop nginx"
    - name: Stop nginx
      type: button
      icon: "🛑"
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## :material-emoticon-outline: Iconos

`icon` antepone un emoji al nombre. Es solo decorativo y se puede cambiar sin
alterar lo que ejecuta el botón.

!!! example "Un icono solo cambia la etiqueta"

    ```yaml title="El mismo botón, con y sin icono"
    - name: Disk usage
      type: button
      function: command
      command: "df -h"

    - name: Disk usage
      type: button
      icon: "💾"
      function: command
      command: "df -h"
    ```

## :material-help-circle-outline: Preguntar antes de acciones peligrosas

Añada `confirm: true` y el bot mostrará primero Sí o Cancelar. Úselo para
cualquier acción que detenga un servicio, elimine datos o reinicie la máquina.
La pregunta caduca tras un tiempo, cinco minutos de forma predeterminada.

Lea [Confirmación](confirmation.md) para saber cuándo usarla y cambiar la espera.

## Ajustes exclusivos de un botón

La mayoría de los ajustes globales se pueden sustituir en un solo botón, lo
que resulta útil cuando una tarea se comporta de forma distinta a las demás:

!!! example "Sustituir ajustes para una tarea lenta"

    ```yaml title="Una tarea lenta que se ejecuta en otro lugar"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"
      workdir: "/var/backups"
      env:
        BACKUP_MODE: "full"
    ```

`timeout` concede más tiempo a este comando, `workdir` elige el directorio
donde se ejecuta y `env` añade variables de entorno solo para él.

## Valores para la función

Escriba los valores de la función directamente en el botón. `command`, `path`
y `args` son atajos para parámetros del mismo nombre. Los nombres
personalizados, como `url`, `host`, `unit` y `lines`, funcionan igual.

!!! example "Pasar valores personalizados"

    ```yaml title="Registros recientes de Nginx"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

No coloque estos valores dentro de un mapa `params:`. Cada clave debe coincidir
con un parámetro declarado por la función seleccionada.
[`validate`](../cli.md#validate) detecta nombres desconocidos, valores
obligatorios ausentes y valores `int` o `bool` inválidos.

## Configuración

Para todos los campos admitidos, consulte
[Configuración → Menú](../configuration.md#menu).

## Relacionado

<div class="grid cards cols-2" markdown>

-   :material-folder-outline:{ .middle } __Categoría__

    ---

    Abre un submenú en vez de ejecutar algo.

    [:octicons-arrow-right-24: Categoría](category.md)

-   :material-function:{ .middle } __Función__

    ---

    Lo que se ejecuta al tocar un botón.

    [:octicons-arrow-right-24: Función](function.md)

-   :material-tune:{ .middle } __Parámetro__

    ---

    Valores que necesita una función y que proporciona el botón.

    [:octicons-arrow-right-24: Parámetro](parameter.md)

-   :material-view-list:{ .middle } __Menú__

    ---

    Cree y organice todo el árbol.

    [:octicons-arrow-right-24: Menú](menu.md)

</div>
