---
title: Menú
description: >-
  Cree el menú que muestra el bot en Telegram. Combine botones y categorías
  en un árbol y aprenda cómo Inicio y Atrás permiten recorrerlo.
icon: material/view-list
---

# :material-view-list: Menú

Su menú es un árbol de nodos bajo la clave raíz `menu`. Hay dos tipos de nodo:

- Una **[categoría](category.md)** abre un submenú y contiene `items`.
- Un **[botón](button.md)** ejecuta una [función](function.md).

Si estos términos son nuevos para usted, lea primero
[Conceptos](button.md). Para consultar la lista exacta de todos los campos,
consulte [Configuración → Menú](../configuration.md#menu).

## :material-format-list-bulleted: Un menú plano

El menú más sencillo es una lista de botones sin anidamiento:

!!! example "Crear un menú con tres botones"

    ```yaml title="Tres botones sin categorías"
    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"

      - name: Free memory
        type: button
        function: command
        command: "free -h"

      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

Envíe `/start` en Telegram para ver el menú.

## :material-folder-outline: Agrupar con categorías

A medida que crezca el menú, agrupe las acciones relacionadas en categorías.
Al tocar una categoría se muestran sus `items`. Inicio siempre está visible y
Atrás solo aparece dentro de una categoría.

!!! example "Agrupar botones en categorías"

    ```yaml title="Categorías System y Services"
    menu:
      - name: System
        type: category
        icon: "💻"
        items:
          - name: Uptime
            type: button
            function: command
            command: "uptime"

          - name: Free memory
            type: button
            function: command
            command: "free -h"

      - name: Services
        type: category
        icon: "🔧"
        items:
          - name: Restart nginx
            type: button
            function: command
            command: "systemctl restart nginx"
            confirm: true
    ```

Las categorías pueden anidarse sin límite y deben contener al menos un elemento.

## :material-family-tree: Los nombres deben ser únicos entre hermanos

Dos nodos del mismo padre no pueden compartir nombre, sin distinguir mayúsculas.
Sí pueden usar el mismo nombre en categorías distintas:

!!! example "Reutilizar un nombre bajo padres distintos"

    ```yaml title="El mismo nombre bajo dos elementos superiores"
    menu:
      - name: Web
        type: category
        items:
          - name: Restart          # válido
            type: button
            function: command
            command: "systemctl restart nginx"
      - name: Database
        type: category
        items:
          - name: Restart          # válido, padre diferente
            type: button
            function: command
            command: "systemctl restart postgresql"
    ```

## :material-emoticon-outline: Iconos

`icon` es un emoji opcional delante del nombre y solo tiene efecto visual.

!!! warning "Este botón reinicia la máquina"

    ```yaml title="Un botón con un icono emoji"
    - name: Reboot
      type: button
      icon: "🔁"
      function: command
      command: "reboot"
      confirm: true
    ```

!!! tip "Elija un emoji sencillo y común"

    Algunos emoji hacen que ciertos teléfonos corten el texto. Si ocurre, pruebe
    otro emoji.

## :material-cellphone-text: Aspecto del menú de Telegram { #how-the-telegram-menu-looks }

Todos los botones aparecen en el teclado bajo el cuadro de mensajes, el que se
muestra y se oculta con el pequeño botón situado en el extremo derecho del
cuadro. Este teclado ocupa siempre todo el ancho del chat, por lo que el texto
de los botones no queda apretado ni cortado.

- **Inicio** es siempre el primer botón de cada pantalla. Tóquelo para volver a
  la primera pantalla.
- **Atrás** aparece dentro de una categoría.
- **$ >_ Run Command** aparece si `enable_run_command` está activado (consulte
  más abajo).
- Los elementos aparecen dos por fila de forma predeterminada. Una categoría
  puede cambiarlo con `columns`. Si una pantalla tiene muchos elementos,
  **Prev** y **Next** permiten recorrer sus páginas.
- Los botones con `confirm: true` preguntan Sí / Cancelar antes de ejecutarse.

!!! info "Los títulos se reutilizan y la salida permanece"

    Cada nuevo título del menú —Inicio, una categoría o una página— sustituye al
    anterior para no llenar el chat de pantallas vacías. La línea **Running** y
    la salida del comando permanecen, por lo que podrá seguir viendo qué se
    ejecutó después de volver a abrir el menú.

!!! info "La salida larga llega en varios mensajes"

    La salida del comando se muestra como bloque de código. Si supera un
    mensaje de Telegram, llega en varios mensajes, cada uno como respuesta al
    anterior. La última parte conserva los botones de la página en la que se
    encontraba, por lo que **Atrás** sigue saliendo de esa categoría. Consulte
    [Configuración → Cuánta salida verá](../configuration.md#how-much-command-output-you-see).

## :material-help-circle-outline: Confirmación { #confirmation }

Añada `confirm: true` a cualquier botón para exigir un segundo toque
(«¿Está seguro?») antes de ejecutarlo. Úselo para cualquier acción destructiva.
Consulte [Confirmación](confirmation.md).

!!! warning "Este botón detiene un servicio"

    ```yaml title="Un botón que pregunta primero"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

La pregunta caduca tras 5 minutos de forma predeterminada. Cambie `confirm_ttl`
en [Configuración → Campos raíz](../configuration.md#root-fields).

## :material-tune-variant: Sustituciones por botón

Algunos ajustes globales se pueden sustituir en un botón:

!!! example "Dar ajustes propios a un botón"

    ```yaml title="Un botón con su propio tiempo, carpeta y variables"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"          # puede tardar más que el tiempo límite global
      workdir: "/var/backups" # ejecutar aquí
      env:
        BACKUP_MODE: "full"   # variable de entorno adicional para este comando
    ```

Consulte [Configuración → Menú](../configuration.md#menu).

## :material-view-grid-outline: Controlar la disposición

`menu_columns` define cuántos botones de **elementos** aparecen por fila
(2 de forma predeterminada). Una categoría puede sustituirlo con `columns`.
Si un menú tiene más elementos que `page_size` (8 de forma predeterminada), se
divide en páginas y se muestran Prev y Next hasta llegar a los extremos.
Consulte [Configuración → Campos raíz](../configuration.md#root-fields).

## :material-console: Ejecutar comando

Con `enable_run_command: true` en la raíz de la configuración, el botón
**$ >_ Run Command** permanece en el menú: aparece después de Atrás dentro de
una categoría o después de Inicio en la primera pantalla. Tóquelo y envíe el
comando de shell que quiera ejecutar. El bot usa el mismo shell, tiempo máximo,
directorio de trabajo y límites de salida que los demás botones.

Inicio o Atrás cancelan la petición sin ejecutar nada. Esta opción está
desactivada de forma predeterminada. Consulte
[Configuración → Campos raíz](../configuration.md#root-fields).

!!! warning "Esto entrega el control de toda la máquina"

    Cualquier usuario permitido podrá ejecutar cualquier comando, no solo los
    botones definidos. Actívelo únicamente si confía plenamente en todos.

## :material-timer-sand: Un comando cada vez

Sus comandos se ejecutan uno tras otro, no en paralelo.

!!! info "Un segundo toque espera su turno"

    Si toca otro botón mientras el primero sigue en ejecución, el segundo
    espera y después se ejecuta por separado. Verá su línea **Running** cuando
    se inicie. Así se evita que dos acciones suyas interfieran sobre el mismo
    servicio o archivo. Los comandos de otras personas no quedan bloqueados:
    cada una tiene su propio turno.

## :material-function-variant: Qué se ejecuta al tocar un botón

Cada botón señala una **función** mediante su campo `function`. Los ejemplos
anteriores usan la función integrada `command`. Para conocer las funciones
integradas y personalizadas y aprender a crear las suyas, consulte
[Funciones](../functions/index.md).

## :material-link-variant: Páginas relacionadas

- [Botón](button.md) — qué es un botón
- [Categoría](category.md) — nodos de submenú
- [Configuración → Menú](../configuration.md#menu) — todos los campos
