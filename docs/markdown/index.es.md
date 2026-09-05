---
title: Telegram Commander
description: >-
  Convierta un archivo YAML en un bot de Telegram que ejecuta comandos en su
  servidor Linux con un solo toque.
icon: material/cellphone-link
hide:
  - navigation
  - toc
---

# :material-cellphone-link: Controle su servidor Linux desde Telegram

<div class="hero" markdown>
<div class="hero-art" markdown>
![Telegram Commander](/images/logo-large.png){ .off-glb width="230" }
</div>
<div class="hero-text" markdown>
**Un toque ejecuta un comando en su servidor y devuelve la salida al chat.**

Telegram Commander convierte un sencillo
[archivo de configuración](documentation/concepts/config-file.md) YAML en un bot
de Telegram con un menú de [botones](documentation/concepts/button.md). Coloque
un comando de terminal detrás de cada botón y ejecútelo desde su teléfono. No
necesita escribir código.

<p class="hero-lang">Lea esta página en otro idioma:
<a href="/" hreflang="en" class="hero-lang__link">English</a>
<a href="/de/" hreflang="de" class="hero-lang__link">Deutsch</a>
<a href="/fr/" hreflang="fr" class="hero-lang__link">Français</a>
<a href="/ru/" hreflang="ru" class="hero-lang__link">Русский</a>
<a href="/zh/" hreflang="zh" class="hero-lang__link">简体中文</a>
<a href="/fa/" hreflang="fa" class="hero-lang__link">فارسی</a>
</p>
</div>
</div>

<div style="text-align: center" markdown="span">
[Empezar :material-arrow-right:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Ver una configuración :material-file-code-outline:](documentation/configuration.md#a-minimal-config){ .md-button }
</div>

[Instalación](documentation/installation/download-and-run.md) ·
[Conceptos](documentation/concepts/config-file.md) ·
[Funciones](documentation/functions/index.md) ·
[Configuración](documentation/configuration.md) ·
[CLI](documentation/cli.md)

## :material-image-multiple-outline:{ .shots } Capturas de pantalla { .split }

Su menú, un comando en ejecución, la salida devuelta y la escritura manual de un
comando. Haga clic en cualquier imagen para verla a tamaño completo.

<div style="text-align: center" markdown="span">
  ![Salida de un comando con el menú System abierto](/images/01.jpeg){ width="140" loading=lazy }
  ![Botones de recursos y procesos](/images/02.jpeg){ width="140" loading=lazy }
  ![Botones de almacenamiento y paquetes](/images/03.jpeg){ width="140" loading=lazy }
  ![Herramientas de red y escritura manual de un comando](/images/04.jpeg){ width="140" loading=lazy }
  ![Salida de un comando devuelta al chat](/images/05.jpeg){ width="140" loading=lazy }
</div>

## :material-lightning-bolt:{ .bolt } Rápido y sencillo { .split }

<div class="grid cards cols-3 center-title step-cards" markdown>

-   :material-file-document-outline:{ .middle } __Escriba su menú__

    ---

    :material-numeric-1-circle:{ .step } Enumere los botones y sus comandos.

-   :material-rocket-launch:{ .middle } __Inicie el bot__

    ---

    :material-numeric-2-circle:{ .step } Ejecútelo ahora o manténgalo activo
    como servicio.

-   :material-gesture-tap-button:{ .middle } __Toque y lea__

    ---

    :material-numeric-3-circle:{ .step } Toque un botón y lea la salida en el
    chat.

</div>

<div style="text-align: center" markdown="span">
[Empezar ahora :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
</div>

## :material-view-grid-outline:{ .grid-icon } Casos de uso { .split }

<div class="grid cards cols-4 icon-left" markdown>

-   :material-restart:{ .lg } Reiniciar o detener un servicio
-   :material-docker:{ .lg } Iniciar y detener contenedores
-   :material-package-down:{ .lg } Actualizar paquetes del sistema
-   :material-text-box-search-outline:{ .lg } Leer registros y diarios
-   :material-harddisk:{ .lg } Comprobar el espacio en disco
-   :material-memory:{ .lg } Vigilar la CPU y la memoria
-   :material-access-point-network:{ .lg } Hacer ping y probar URL
-   :material-backup-restore:{ .lg } Crear y restaurar copias de seguridad
-   :material-script-text:{ .lg } Ejecutar sus propios scripts
-   :material-power:{ .lg } Reiniciar o apagar el host
-   :material-console:{ .lg } Escribir cualquier comando manualmente
-   :material-all-inclusive:{ .lg } Y casi cualquier otra cosa

</div>

## :material-thumb-up-outline:{ .thumb } Por qué usarlo { .split }

<div class="grid cards cols-4 center-title" markdown>

-   :material-clock-fast:{ .lg .middle } __Sin programar__

    ---

    Describa el menú y los comandos en un archivo YAML.

    [:octicons-arrow-right-24: Archivo de configuración](documentation/concepts/config-file.md)

-   :material-cellphone-link:{ .lg .middle } __Desde cualquier lugar__

    ---

    Abra Telegram en su teléfono y controle el servidor sin una VPN hacia él.

    [:octicons-arrow-right-24: Cómo se conecta el bot](documentation/concepts/long-polling.md)

-   :material-lan-disconnect:{ .lg .middle } __Sin puertos abiertos__

    ---

    El bot se conecta a Telegram. Nada queda expuesto a Internet.

    [:octicons-arrow-right-24: Cómo se conecta el bot](documentation/concepts/long-polling.md)

-   :material-message-text-outline:{ .lg .middle } __Salida en el chat__

    ---

    El resultado vuelve como mensaje. No necesita una sesión SSH.

    [:octicons-arrow-right-24: Cuánta salida verá](documentation/configuration.md#how-much-command-output-you-see)

-   :material-shield-lock:{ .lg .middle } __Controlado y registrado__

    ---

    Elija quién recibe el menú, confirme acciones peligrosas y registre cada ejecución.

    [:octicons-arrow-right-24: Acceso y confirmación](documentation/concepts/allowed-users.md)

-   :material-folder-outline:{ .lg .middle } __Menús anidados__

    ---

    Agrupe botones en categorías. Inicio permanece arriba y Atrás sube un nivel.

    [:octicons-arrow-right-24: Menú](documentation/concepts/menu.md)

-   :material-function-variant:{ .lg .middle } __Funciones reutilizables__

    ---

    Escriba un comando una vez y complete valores distintos en cada botón.

    [:octicons-arrow-right-24: Funciones](documentation/functions/index.md)

-   :material-cog-play-outline:{ .lg .middle } __Siempre activo__

    ---

    Instálelo como servicio para iniciar el bot con el host.

    [:octicons-arrow-right-24: Ejecutar como servicio](documentation/installation/run-as-a-service.md)

</div>

## :material-file-code-outline:{ .code-icon } Un ejemplo pequeño { .split }

Esta configuración crea un bot con un botón llamado «Uptime». Al tocarlo,
ejecuta el comando `uptime` en el servidor.

!!! example "Esta configuración completa añade un botón"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN" # (1)!
      allowed_users:
        - "YOUR_USER_ID" # (2)!

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime" # (3)!
    ```

    1.  Pida un token a BotFather en Telegram al crear el bot.
    2.  Solo las cuentas enumeradas pueden abrir el menú. Puede usar un id
        numérico o un `@username`.
    3.  Aquí puede escribir cualquier comando válido en una terminal.

Esta es una configuración completa y funcional. Todo lo demás es opcional.

!!! tip "Diseñado para un grupo pequeño y de confianza"

    Nada escucha conexiones entrantes y solo las cuentas de `allowed_users`
    reciben un menú. Quien pueda usar el bot puede ejecutar todos los botones
    definidos, así que mantenga breve esa lista.

## :material-hand-pointing-right: ¿Listo para probarlo? { .split }

<div style="text-align: center" markdown>
[Empezar ahora :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Conceptos :material-book-open-variant:](documentation/concepts/config-file.md){ .md-button }
[Última versión :material-download:](https://github.com/azolfagharj/telegram-commander/releases/latest){ .md-button .md-button--primary }

[Ver el código fuente :fontawesome-brands-github:](https://github.com/azolfagharj/telegram-commander){ .md-button }
</div>

Telegram Commander es gratuito y de código abierto. Si le ahorra tiempo,
[considere apoyar su desarrollo](https://azolfagharj.github.io/donate/); su
ayuda permite mantener vivo el proyecto.
