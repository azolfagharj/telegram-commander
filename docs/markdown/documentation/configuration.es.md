---
icon: material/file-cog-outline
title: Configuración
description: >-
  Todos los ajustes de Telegram Commander con su tipo, valor predeterminado
  y significado: telegram, menu, function_directory, tiempos, salida y
  registro.
---

# :material-file-cog-outline: Configuración

El [archivo de configuración](concepts/config-file.md) describe todo el bot: la
conexión con Telegram, quién puede usarlo, el menú de
[botones](concepts/button.md) y el registro. Se pasa con `--config` a `run`,
`validate`, `fmt` y `list-functions` (consulte [CLI](cli.md)).

Todas las claves usan `lower_snake_case`. **Las claves desconocidas se
rechazan**, por lo que verá inmediatamente cualquier error al
[validar](cli.md#validate).

**Obligatorio** significa que la validación falla si el campo falta o queda
vacío tras aplicar los valores predeterminados.  
Los campos **opcionales** se pueden omitir; la columna Predeterminado indica el
valor usado.

Si es nuevo, empiece con
[Ejecutar en la CLI](installation/download-and-run.md), que explica cómo crear
una primera configuración. Consulte [Conceptos](concepts/config-file.md) para
conocer el vocabulario.

## :material-rocket-launch-outline: Una configuración mínima { #a-minimal-config }

Solo son obligatorios `telegram`, con un token y un
[usuario permitido](concepts/allowed-users.md), y `menu`. Todo lo demás tiene
un valor predeterminado:

!!! example "Empezar con un usuario permitido y un botón"

    ```yaml title="config.yaml (mínima)"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

La carpeta `config-examples/` de la versión contiene un ejemplo mínimo y otro completo.

## :material-card-bulleted-outline: Campos raíz { #root-fields }

| Campo | Tipo | Obligatorio | Predeterminado | Descripción |
|-------|------|-------------|----------------|-------------|
| `telegram` | objeto | sí | — | Ajustes de Telegram (véase más abajo) |
| `menu` | lista | sí | — | Árbol del menú; al menos un nodo |
| `function_directory` | cadena | no | sin definir | Directorio YAML de funciones personalizadas (consulte las reglas más abajo) |
| `shell` | cadena | no | `/bin/bash` | [Shell](concepts/shell.md) usado como `shell -c "<command>"` |
| `timeout` | duración | no | `60s` | Tiempo máximo predeterminado de un comando |
| `max_output_bytes` | entero | no | `524288` | Salida máxima guardada por comando (consulte [Cuánta salida de comando verá](#how-much-command-output-you-see)) |
| `workdir` | cadena | no | directorio de trabajo del proceso | Directorio de trabajo predeterminado para los comandos |
| `env` | mapa | no | vacío | Variables de entorno adicionales para los comandos |
| `menu_columns` | entero | no | `2` | Botones de elementos por fila bajo el cuadro de mensajes |
| `page_size` | entero | no | `8` | Elementos por página antes de paginar |
| `confirm_ttl` | duración | no | `5m` | Vigencia de una [confirmación](concepts/confirmation.md) |
| `enable_run_command` | booleano | no | `false` | Muestra un botón **`$ >_ Run Command`** que ejecuta el siguiente mensaje como comando de shell. Está desactivado de forma predeterminada. Cualquier persona que pueda usar el bot podrá ejecutar cualquier comando en el host, así que actívelo solo si confía en todos los usuarios permitidos. Esta clave no es válida bajo `telegram`. |
| `logging` | objeto | no | registrador integrado predeterminado | Registradores con nombre (consulte más abajo) |

??? note "¿Qué ocurre si omite `shell`?"

    Puede omitirlo. El bot usa `/bin/bash`. Lo mismo ocurre con `timeout`,
    `page_size` y los demás campos raíz opcionales: si los omite, se aplican sus
    valores predeterminados. Solo debe definirlos cuando quiera usar un valor
    distinto, por ejemplo `shell: /bin/sh`.

### Cuánta salida de comando verá { #how-much-command-output-you-see }

Se aplican dos límites consecutivos. `max_output_bytes` es **su** límite, además
de un límite de Telegram que no puede cambiar.

**1. Su límite: `max_output_bytes`** (predeterminado `524288`, es decir, 512 KB)

Mientras se ejecuta un comando, el bot conserva como máximo esa cantidad por
separado para la salida normal y la de error. Lo demás se descarta, pero el
comando continúa hasta terminar o alcanzar su `timeout`. El resultado comienza
entonces con `(output truncated)`.

**2. Límite de Telegram: un mensaje contiene como máximo 4096 bytes**

Si el resultado supera un mensaje, el bot lo divide en varios. Cada parte se
envía como respuesta a la anterior para mantenerlas juntas y ordenadas, y los
botones del menú aparecen en la última. Siempre que sea posible, la división se
realiza entre líneas para no cortarlas por la mitad.

Si sigue siendo muy largo, el bot se detiene después de 10 mensajes y añade una
nota final como `(output too long; showing first N bytes)`, donde `N` indica
cuánta salida ha recibido realmente.

Aumentar `max_output_bytes` permite conservar más salida, pero seguirá viendo
como máximo unos diez mensajes. Para una salida tan larga, suele ser mejor
acortar el comando, por ejemplo `journalctl -u nginx | tail -n 50`, o escribir
la salida completa en un archivo del servidor.

### Reglas de `function_directory` { #function_directory-rules }

| Situación | Resultado |
|-----------|-----------|
| Falta la clave | Registro informativo; solo funciones integradas |
| La clave está presente pero vacía (`""`) | Registro informativo; solo funciones integradas |
| La ruta no existe o no es accesible | Error grave; el proceso se detiene |
| La ruta existe pero el directorio está vacío | Correcto |

!!! warning "Una ruta incorrecta detiene el bot"

    Si `function_directory` apunta a una carpeta inexistente o ilegible, el
    programa se detiene en vez de iniciarse sin sus funciones personalizadas.

## :material-send-circle-outline: `telegram` { #telegram }

| Campo | Tipo | Obligatorio | Predeterminado | Descripción |
|-------|------|-------------|----------------|-------------|
| `bot_token` | cadena | sí | — | Token del bot proporcionado por BotFather |
| `allowed_users` | lista de cadenas | sí | — | [Usuarios permitidos](concepts/allowed-users.md) |
| `api` | cadena | no | `https://api.telegram.org` | URL base de la API del bot |
| `proxy.enabled` | booleano | no | `false` | Usar un proxy para la API de Telegram |
| `proxy.url` | cadena | condicional | — | Obligatorio si `proxy.enabled` es `true` |
| `insecure` | booleano | no | `false` | Omitir la verificación TLS (no recomendado) |

Los usuarios no autorizados reciben su `user_id` y `username` para solicitar
acceso. Así puede encontrar su propio id la primera vez; consulte
[Ejecutar en la CLI → Paso 5](installation/download-and-run.md#step-5-find-your-user-id-if-needed).

!!! example "Conectarse mediante un proxy"

    ```yaml title="Sección telegram con un proxy"
    telegram:
      bot_token: "123456789:AAExampleTokenValue"
      allowed_users:
        - "123456789"        # id numérico del usuario
        - "@alice"           # o un nombre de usuario
      proxy:
        enabled: true
        url: "socks5://127.0.0.1:10808"
    ```

Para permitir comandos escritos en Telegram, añada este ajuste a la **raíz**
del archivo, no bajo `telegram`:

!!! tip "Añadir un ajuste raíz"

    ```yaml title="Activar el botón Run Command"
    enable_run_command: true
    ```

## :material-menu: Menú { #menu }

Esta sección es la referencia de campos. Para una explicación guiada, consulte
[Menú](concepts/menu.md). Cada [botón](concepts/button.md) o
[categoría](concepts/category.md) acepta:

| Campo | Tipo | Obligatorio | Descripción |
|-------|------|-------------|-------------|
| `name` | cadena | sí | Nombre visible, único entre hermanos sin distinguir mayúsculas |
| `type` | `category` \| `button` | sí | Tipo de nodo |
| `items` | lista | si es `category` | Hijos; al menos uno |
| `function` | cadena | si es `button` | Nombre de la [función](concepts/function.md) |
| `command` | cadena | si `function: command` | Comando de shell integrado |
| `path` | cadena | si `function: script` | Ruta del script |
| `icon` | cadena | no | Emoji opcional |
| `id` | cadena | no | Identificador opcional para este nodo; puede omitirlo |
| `confirm` | booleano | no | Pedir [confirmación](concepts/confirmation.md), `false` de forma predeterminada |
| `timeout` | duración | no | Sustituir el tiempo máximo global |
| `workdir` | cadena | no | Sustituir el directorio de trabajo |
| `env` | mapa | no | Variables de entorno adicionales para este botón |
| `columns` | entero | no | Sustituir las columnas de la categoría |
| `args` | cadena | no | Argumentos opcionales de `script` |
| Cualquier parámetro declarado | escalar | según la función | Valor pasado, como `url`, `host`, `unit` o `lines` |

En un **botón**, cualquier otra clave escalar se trata como un parámetro de
función. Su nombre debe estar declarado por la función seleccionada.
[`validate`](cli.md#validate) rechaza los nombres desconocidos y también
comprueba los valores declarados como `int` o `bool`. Las cadenas, los números
y los booleanos se pueden escribir directamente como valores YAML; los números
no necesitan comillas.

En una **categoría**, cualquier clave fuera de los campos anteriores es un
error. Las categorías no ejecutan funciones, por lo que no pueden tener claves
de parámetros.

`command`, `path` y `args` completan parámetros del mismo nombre. Escriba los
demás [parámetros](concepts/parameter.md) directamente en el botón, sin un mapa
`params:`. Consulte
[Funciones → Pasar valores desde un botón](functions/index.md#passing-values-from-a-button).

## :material-math-log: `logging` { #logging }

Es opcional. Si se omite, se usa un registrador de consola en `stderr` con nivel `info`.

Registradores con nombre:

!!! example "Escribir registros normales y un archivo de auditoría"

    ```yaml title="Sección logging con un archivo de auditoría"
    logging:
      logs:
        default:
          level: info
          format: console   # o JSON
          output:
            - output: stderr
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

Salidas admitidas: `stdout`, `stderr`, `file`, `discard`.

El registrador `audit` guarda cada comando: quién lo ejecutó, qué botón, el
código de salida y la duración. Consulte
[Registro de auditoría](concepts/audit-log.md).

## Páginas relacionadas

- [Ejecutar en la CLI](installation/download-and-run.md) — cree y ejecute una primera configuración
- [Menú](concepts/menu.md) — el árbol del menú en detalle
- [Funciones](functions/index.md) — significado de `function`, `command`, `path` y `args`
- [CLI](cli.md) — valide y ejecute su configuración
