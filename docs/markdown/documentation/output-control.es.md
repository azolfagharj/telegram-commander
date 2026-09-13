---
title: Controlar la salida
description: >-
  Decida cómo llegan al chat los resultados de los comandos. Defina output,
  max_output_messages y max_output_bytes en la raíz, sustituya output en un
  botón concreto y vea todas las combinaciones de raíz y botón con ejemplos.
icon: material/export-variant
---

# :material-export-variant: Controlar la salida

Cuando un [botón](concepts/button.md) termina, el bot tiene que llevar el
resultado a su chat. Telegram nunca acepta más de 4096 bytes en un mensaje, así
que un resultado largo llega como **varios mensajes de respuesta** o como **un
archivo `.txt`** que puede abrir y descargar.

Tres campos opcionales deciden qué ocurre. Puede omitirlos todos y el bot
seguirá comportándose de forma razonable.

!!! info "La única regla que debe recordar"

    Manda el botón. Si un botón define `output`, eso es lo que ocurre. Si el
    botón no dice nada, se aplica el valor de la raíz. Y si la raíz tampoco
    dice nada, el bot usa `auto`.

<div class="grid cards cols-3" markdown>

-   :material-scissors-cutting:{ .middle } __`max_output_bytes`__

    ---

    Cuánta salida del comando se guarda antes de enviar nada.

    [:octicons-arrow-right-24: Leer más](#max_output_bytes)

-   :material-tune:{ .middle } __`output`__

    ---

    Mensajes o archivo. El único campo que también puede definir en un botón.

    [:octicons-arrow-right-24: Leer más](#output)

-   :material-counter:{ .middle } __`max_output_messages`__

    ---

    En modo `auto`, cuántos mensajes se permiten antes de enviar un archivo.

    [:octicons-arrow-right-24: Leer más](#max_output_messages)

</div>

!!! tip "Todas las imágenes de esta página se amplían"

    Las capturas del chat son pequeñas a propósito. Pulse una para verla a
    tamaño completo.

## :material-card-bulleted-outline: Los tres campos de un vistazo { #the-three-fields-at-a-glance }

| Campo | Dónde se define | Tipo | Predeterminado | Qué decide |
|-------|-----------------|------|----------------|------------|
| `max_output_bytes` | solo en la raíz | entero | `524288` | Cuántos bytes de salida guarda el bot por comando |
| `output` | raíz **y** botón | `auto` \| `text` \| `file` | `auto` | Si el resultado llega como mensajes de chat o como un archivo `.txt` |
| `max_output_messages` | solo en la raíz | entero, `1`–`10` | `2` | En modo `auto`, cuántos mensajes se permiten antes de enviar un archivo |

Los tres están en su [archivo de configuración](concepts/config-file.md), en la
**raíz**, junto a `shell` y `timeout`, no bajo `telegram`. La lista completa de
claves raíz está en la página [Configuración](configuration.md#root-fields).

## :material-scissors-cutting: `max_output_bytes`

Este es su propio límite y se aplica primero. Mientras corre un comando, el bot
guarda como máximo esta cantidad de bytes de su salida, contada por separado
para la salida normal y la de error. Lo que pase de ahí se descarta, pero el
comando sigue en marcha hasta terminar o alcanzar su `timeout`.

!!! example "Guardar hasta 2 MB de salida por comando"

    ```yaml title="config.yaml"
    max_output_bytes: 2097152 # (1)!
    ```

    1.  2 MB en vez de los 512 KB predeterminados. Se cuenta por separado para
        la salida normal y para la de error.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Cuando se alcanza el límite, el resultado lo indica en su propia línea,
    justo debajo del resumen. Todo lo que está encima de `(output truncated)`
    es la parte que guardó el bot; el resto del registro nunca salió del
    servidor.

    </div>
    <div class="result-shot" markdown>

    ![Un resultado de Telegram que indica que la salida se recortó](/images/output-control/truncated.png){ .shot loading=lazy }

    </div>
    </div>

!!! tip "Subirlo solo sirve si se envía un archivo"

    Un mensaje de chat nunca puede llevar más de 4096 bytes, así que un
    `max_output_bytes` mayor solo le llega completo cuando el resultado viene
    como archivo `.txt`. Combínelo con `output: auto` (el predeterminado) o con
    `output: file`.

## :material-tune: `output`

`output` elige la entrega. Acepta tres valores y es el único campo de salida
que también puede poner en un botón concreto.

=== "auto"

    Mensajes de texto mientras el resultado es corto y un archivo `.txt` cuando
    no lo es. Es el valor predeterminado y el que conserva si no escribe nada.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Dejar que el bot decida en cada resultado"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: auto # (1)!
        max_output_messages: 2 # (2)!
        ```

        1.  Decidir según el resultado en vez de forzar una única entrega.
        2.  El punto de cambio. Dos mensajes o menos se quedan en el chat;
            cualquier cosa más larga se convierte en un archivo.

    </div>
    <div class="result-shot" markdown>

    ![El modo auto deja la salida corta como texto y envía la más larga como archivo](/images/output-control/auto.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

=== "text"

    Siempre mensajes de chat, por largo que sea el resultado. El bot se detiene
    tras diez mensajes y añade una nota de que se cortó el resto.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Mantener todo en el chat"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: text # (1)!
        ```

        1.  Aquí se ignora `max_output_messages`, porque el bot nunca cambia a
            un archivo por su cuenta.

    </div>
    <div class="result-shot" markdown>

    ![El resultado de un comando entregado como mensajes de Telegram](/images/output-control/message.png){ .shot loading=lazy }

    </div>
    </div>

=== "file"

    Siempre un archivo `.txt`, incluso para un resultado de una línea. El
    resumen llega como pie del archivo.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Enviar cada resultado como adjunto"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: file # (1)!
        ```

        1.  Ahora todos los botones envían un adjunto, incluso los más
            rápidos. Si es demasiado, póngalo en botones concretos.

    </div>
    <div class="result-shot" markdown>

    ![El resultado de un comando entregado como archivo de Telegram](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

!!! note "Omitirlo es lo mismo que escribir `auto`"

    Nunca hace falta escribir `output`. Una configuración que funcionaba antes
    de que existiera este campo sigue comportándose igual.

## :material-counter: `max_output_messages`

Este campo solo importa en modo `auto` y solo existe en la raíz. Es el número
de mensajes de chat que el bot está dispuesto a enviar antes de renunciar a los
mensajes y enviar un solo archivo. Los valores permitidos van de `1` a `10`; el
predeterminado es `2`.

!!! example "Enviar un archivo en cuanto el resultado necesite un segundo mensaje"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 1 # (1)!
    ```

    1.  El ajuste más estricto. Solo un resultado que quepa en un mensaje se
        queda en el chat.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un resultado que cabe en el margen permitido se queda como mensajes de
    chat. En cuanto necesitaría un mensaje más, todo el resultado llega como un
    único archivo `.txt` y no se envía ningún mensaje parcial.

    </div>
    <div class="result-shot" markdown>

    ![Un resultado corto como mensaje junto a otro más largo como archivo](/images/output-control/allowance.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

!!! warning "Es un campo solo de la raíz"

    Un botón no puede definir `max_output_messages`. Póngalo en la raíz y deje
    que los botones que necesiten algo distinto usen `output: text` u
    `output: file`.

## :material-table-arrow-right: La raíz y el botón juntos { #root-and-button-together }

Un botón puede llevar su propio `output`. No hay nada más de la salida que se
pueda definir por botón. El modo efectivo se elige así:

!!! abstract "Cómo se elige el modo efectivo"

    ```text title="Orden de precedencia"
    button output  →  root output  →  auto
    ```

| `output` de la raíz | `output` del botón | Qué ocurre en realidad |
|---------------------|--------------------|------------------------|
| sin definir o `auto` | sin definir | `auto` |
| sin definir o `auto` | `auto` | `auto` |
| sin definir o `auto` | `text` | `text` |
| sin definir o `auto` | `file` | `file` |
| `text` | sin definir | `text` |
| `text` | `auto` | `auto` |
| `text` | `text` | `text` |
| `text` | `file` | `file` |
| `file` | sin definir | `file` |
| `file` | `auto` | `auto` |
| `file` | `text` | `text` |
| `file` | `file` | `file` |

Las secciones siguientes recorren las doce combinaciones, con un ejemplo cada
una.

## :material-numeric-1-box-outline: Cuando la raíz es `auto` { #when-the-root-is-auto }

Aquí se incluye el caso en que la raíz no dice nada, porque un `output` ausente
significa `auto`.

### La raíz es `auto` y el botón no dice nada { #the-root-is-auto-the-button-says-nothing }

El caso de cada día. Los resultados cortos se quedan en el chat; un resultado
que necesitaría más de `max_output_messages` mensajes llega como archivo.

!!! example "Un botón sencillo con una raíz sencilla"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 2

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    `df -h` imprime unas pocas líneas, así que cabe en un mensaje y se queda en
    el chat como bloque de código con el resumen encima.

    </div>
    <div class="result-shot" markdown>

    ![El uso de disco devuelto en un solo mensaje de Telegram](/images/output-control/disk-usage.png){ .shot loading=lazy }

    </div>
    </div>

### La raíz es `auto` y el botón dice `auto` { #the-root-is-auto-the-button-says-auto }

Escribir `auto` en el botón no cambia nada aquí. Solo vale la pena escribirlo
si quiere que el botón siga en `auto` pase lo que pase con la raíz más
adelante.

!!! example "Escribir el valor predeterminado en el botón"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
        output: auto # (1)!
    ```

    1.  Hoy se comporta igual que si lo omitiera, pero este botón lo conserva
        aunque la raíz cambie a `text` o `file` más adelante.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Exactamente el comportamiento automático: texto mientras el resultado es
    corto y un archivo en cuanto supera el margen permitido. En un servidor con
    muchos discos montados, ese mismo botón puede enviarle un archivo.

    </div>
    <div class="result-shot" markdown>

    ![El mismo botón como mensaje y como archivo](/images/output-control/disk-usage-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### La raíz es `auto` y el botón dice `text` { #the-root-is-auto-the-button-says-text }

Úselo para un botón cuya salida quiera leer siempre en el chat, aunque se
alargue un poco: una lista de estado que se recorre en vez de descargarla.

!!! example "Mantener un botón en el chat"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Service status
        type: button
        function: command
        command: "systemctl status nginx --no-pager"
        output: text # (1)!
    ```

    1.  Nunca un archivo, diga lo que diga `max_output_messages`.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    El resultado se parte en varios mensajes, cada uno como respuesta al
    anterior, así que nunca se pierde el orden. Los botones se quedan en la
    última parte.

    </div>
    <div class="result-shot" markdown>

    ![Un estado largo repartido en tres mensajes de Telegram](/images/output-control/service-status-split.png){ .shot loading=lazy }

    </div>
    </div>

### La raíz es `auto` y el botón dice `file` { #the-root-is-auto-the-button-says-file }

La opción típica para registros y copias de seguridad: casi nunca quiere
leerlos en el chat y un archivo es más fácil de guardar.

!!! example "Descargar siempre este"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Nginx logs
        type: button
        function: command
        command: "journalctl -u nginx -n 2000 --no-pager"
        output: file # (1)!
    ```

    1.  El resto del menú sigue comportándose de forma automática. Solo este
        botón queda fijado a un archivo.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un adjunto con el nombre del botón, el código de salida y la duración como
    pie. Tóquelo para leer todo el registro o guárdelo para más adelante.

    </div>
    <div class="result-shot" markdown>

    ![Los registros de Nginx entregados como archivo de Telegram](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-2-box-outline: Cuando la raíz es `output: text` { #when-the-root-is-output-text }

Todos los botones se quedan en el chat salvo que digan otra cosa. Elija esto si
no le gusta descargar archivos en el teléfono.

### La raíz es `text` y el botón no dice nada { #the-root-is-text-the-button-says-nothing }

!!! example "Mensajes de chat en todas partes"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un mensaje corto, porque `uptime` imprime una sola línea. Los resultados
    más largos simplemente se reparten en más mensajes, hasta diez.

    </div>
    <div class="result-shot" markdown>

    ![El tiempo de actividad devuelto en un solo mensaje de Telegram](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### La raíz es `text` y el botón dice `auto` { #the-root-is-text-the-button-says-auto }

El botón vuelve al cambio automático, así que este botón puede enviar un
archivo aunque el resto del menú no lo haga nunca. El `max_output_messages` de
la raíz decide dónde ocurre el cambio.

!!! example "Un botón autorizado a enviar un archivo cuando se alarga"

    ```yaml title="config.yaml"
    output: text
    max_output_messages: 3 # (1)!

    menu:
      - name: Package list
        type: button
        function: command
        command: "dpkg -l"
        output: auto
    ```

    1.  Solo este botón lee el margen permitido, porque es el único en modo
        `auto`.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Una lista de paquetes corta llega en hasta tres mensajes. En un servidor
    lleno, el mismo botón supera el margen y envía un solo archivo.

    </div>
    <div class="result-shot" markdown>

    ![Una lista de paquetes como mensajes y como archivo](/images/output-control/package-list-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### La raíz es `text` y el botón dice `text` { #the-root-is-text-the-button-says-text }

Lo mismo que no decir nada, pero con la intención escrita. Es útil en una
configuración larga donde no quiere que un cambio posterior en la raíz afecte a
este botón.

!!! example "Fijar un botón a los mensajes de chat"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Last logins
        type: button
        function: command
        command: "last -n 20"
        output: text
    ```

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Siempre mensajes de chat, hasta diez, sea lo que sea que cambie en la raíz
    más adelante.

    </div>
    <div class="result-shot" markdown>

    ![Los últimos inicios de sesión devueltos en un mensaje de Telegram](/images/output-control/last-logins.png){ .shot loading=lazy }

    </div>
    </div>

### La raíz es `text` y el botón dice `file` { #the-root-is-text-the-button-says-file }

!!! example "Un botón descargable en un menú de solo chat"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Full system report
        type: button
        function: command
        command: "/usr/local/bin/report.sh"
        output: file
        timeout: "5m" # (1)!
    ```

    1.  Un informe puede tardar, así que este botón recibe más tiempo que el
        límite global.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    El único adjunto en un menú que por lo demás es solo de chat. El pie
    muestra cuánto tardó el informe, lo que resulta útil cuando corre durante
    minutos.

    </div>
    <div class="result-shot" markdown>

    ![Un informe completo del sistema entregado como archivo de Telegram](/images/output-control/system-report-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-3-box-outline: Cuando la raíz es `output: file` { #when-the-root-is-output-file }

Cada resultado es un archivo, incluso el de una sola línea. Encaja con un bot
usado sobre todo para informes y copias guardadas.

### La raíz es `file` y el botón no dice nada { #the-root-is-file-the-button-says-nothing }

!!! example "Archivos en todas partes"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Incluso una sola línea de salida llega como adjunto. El pie sigue llevando
    el resumen, así que puede leer el código de salida sin abrirlo.

    </div>
    <div class="result-shot" markdown>

    ![El tiempo de actividad entregado como archivo pequeño de Telegram](/images/output-control/uptime-file.png){ .shot loading=lazy }

    </div>
    </div>

!!! warning "Hasta los resultados mínimos se convierten en descargas"

    Una respuesta de una línea llega como adjunto que hay que abrir. Si eso le
    molesta, deje la raíz en `auto` y ponga `output: file` solo en los botones
    que lo necesiten.

### La raíz es `file` y el botón dice `auto` { #the-root-is-file-the-button-says-auto }

El botón escapa de la regla de solo archivos y vuelve a comportarse con
normalidad.

!!! example "Dejar legible una comprobación rápida"

    ```yaml title="config.yaml"
    output: file
    max_output_messages: 2

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
        output: auto # (1)!
    ```

    1.  Vuelta al cambio automático solo para este botón.

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    La respuesta de una línea es corta, así que se queda en el chat como
    mensaje mientras el resto de los botones del menú siguen enviando archivos.

    </div>
    <div class="result-shot" markdown>

    ![El tiempo de actividad de vuelta en el chat como mensaje](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### La raíz es `file` y el botón dice `text` { #the-root-is-file-the-button-says-text }

!!! example "Forzar un botón de vuelta al chat"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Who is logged in
        type: button
        function: command
        command: "who"
        output: text
    ```

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un mensaje, nunca un adjunto, por mucho que crezca la lista de usuarios
    conectados. Por encima de diez mensajes el bot indica que se cortó el
    resto.

    </div>
    <div class="result-shot" markdown>

    ![Los usuarios conectados devueltos en un mensaje de Telegram](/images/output-control/who.png){ .shot loading=lazy }

    </div>
    </div>

### La raíz es `file` y el botón dice `file` { #the-root-is-file-the-button-says-file }

Repite el valor de la raíz. Es inofensivo y mantiene el botón correcto si más
adelante relaja la raíz.

!!! example "Un botón que siempre debe ser un archivo"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Database dump log
        type: button
        function: command
        command: "cat /var/log/pg_dump.log"
        output: file
    ```

!!! success "Qué verá en el chat"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un adjunto, exactamente lo que ya pedía la raíz. Escribirlo en el botón
    hace que la intención sobreviva a un cambio en la raíz.

    </div>
    <div class="result-shot" markdown>

    ![El registro de un volcado de base de datos entregado como archivo de Telegram](/images/output-control/db-dump-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-file-document-outline: Cómo es el archivo { #what-the-file-looks-like }

Cuando un resultado se entrega como archivo, recibe un adjunto y un pie breve.

- **Nombre** — el nombre del botón en minúsculas y con guiones, después la
  fecha y la hora en UTC, por ejemplo `nginx-logs-20260913-091204.txt`.
- **Pie** — el mismo resumen que vería al principio de un mensaje de chat:
  nombre del botón, código de salida, duración y la nota de tiempo agotado o de
  recorte cuando la haya.
- **Cuerpo** — otra vez ese resumen, después una sección `--- stdout ---` y una
  sección `--- stderr ---` cuando el comando haya escrito algo en ella.

!!! success "Dentro del archivo descargado"

    ```text title="nginx-logs-20260913-091204.txt"
    Button: Nginx logs
    Exit: 0
    Duration: 316ms

    --- stdout ---
    2026-09-13T09:11:58 nginx: worker process started
    ...

    --- stderr ---
    2026-09-13T09:12:01 nginx: warning: duplicate server name
    ```

## :material-alert-outline: Límites que no puede cambiar { #limits-you-cannot-change }

!!! warning "El techo lo pone Telegram, no el bot"

    - Un mensaje de chat admite como máximo 4096 bytes.
    - En modo `text` el bot envía como máximo diez mensajes y termina con
      `(output too long; showing first N bytes)`.
    - Escriba el `max_output_messages` que quiera: diez mensajes sigue siendo
      el techo absoluto.

!!! note "Si falla el envío del archivo, seguirá recibiendo el resultado"

    Cuando el archivo no se puede enviar, porque no hay red o porque Telegram
    lo rechaza, el bot vuelve a los mensajes de chat para que el resultado no
    se pierda.

## :material-console: Ejecutar comando usa el ajuste de la raíz { #run-command-uses-the-root-setting }

El botón **`$ >_ Run Command`** no forma parte de su menú, así que no tiene un
`output` propio. Siempre sigue el `output` y el `max_output_messages` de la
raíz. Consulte [Menú → Ejecutar comando](concepts/menu.md#run-command).

## :material-link-variant: Relacionado { #related }

<div class="grid cards cols-2" markdown>

-   :material-cog-outline:{ .middle } __Configuración__

    ---

    Todas las claves raíz, con los tres campos de salida en una tabla.

    [:octicons-arrow-right-24: Configuración](configuration.md#root-fields)

-   :material-gesture-tap-button:{ .middle } __Botón__

    ---

    Lo que un botón puede sustituir, `output` incluido.

    [:octicons-arrow-right-24: Botón](concepts/button.md)

-   :material-view-list:{ .middle } __Menú__

    ---

    Cómo conviven los resultados con el menú y qué hace Ejecutar comando.

    [:octicons-arrow-right-24: Menú](concepts/menu.md)

-   :material-timer-outline:{ .middle } __Shell__

    ---

    El shell, el directorio de trabajo y el tiempo máximo de sus comandos.

    [:octicons-arrow-right-24: Shell](concepts/shell.md)

</div>
