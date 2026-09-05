---
title: ¿Qué es una función?
description: >-
  Una función convierte los valores escritos en un botón en un comando de
  shell. Aprenda cómo proporciona el botón los parámetros declarados.
---

# :material-function-variant: ¿Qué es una función?

Una **función** es una receta que convierte valores con nombre, sus
[parámetros](../concepts/parameter.md), en un comando de shell. Cada
[botón](../concepts/button.md) nombra una función en su campo `function`.

Imagine una función como un comando con espacios por completar. Una función de
uso de disco puede tener un espacio, la ruta, que cada botón completa.

## Qué ocurre al tocar un botón

1. El bot busca la función indicada en el campo `function`.
2. Recoge los valores escritos en el botón.
3. Construye un comando de shell con esos valores.
4. Ejecuta el comando en el [shell](../concepts/shell.md) y devuelve la salida
   al chat como bloque de código.

Si la función no existe o falta un valor, el bot no se inicia:
[`validate`](../cli.md#validate) informa primero del problema.

## Un ejemplo completo

La función `command` está integrada y siempre disponible. Ejecuta lo escrito en
el campo `command` del botón.

!!! example "Usar la función integrada command"

    ```yaml title="Botón Uptime"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

Toque **Uptime** y el bot ejecutará `uptime` en el servidor y devolverá la
salida.

## Pasar valores desde un botón { #passing-values-from-a-button }

Escriba cada valor directamente en el botón. Hay dos formas:

1. Use los campos `command`, `path` y `args`, que completan parámetros iguales.
2. Para cualquier otro parámetro, use su nombre como clave del botón.

!!! example "Pasar una URL mediante su nombre de parámetro"

    ```yaml title="Botón Check API"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

Aquí `url` coincide con el parámetro declarado por `curl-url`. Lo mismo se
aplica a `host`, `unit` y `lines`.

!!! warning "No coloque valores dentro de `params:`"

    `params:` declara valores en un archivo de función. En un botón, escríbalos
    directamente:

    ```yaml title="Los valores se escriben directamente en el botón"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

    Los valores numéricos de YAML no necesitan comillas.

[`validate`](../cli.md#validate) comprueba cada clave con los parámetros
declarados por la función seleccionada. Un error ortográfico o un nombre no
declarado hacen fallar la validación. También comprueba los valores declarados
como `int` o `bool`. Los valores opcionales usan sus valores predeterminados
cuando el botón los omite.

## Dos tipos de función { #two-kinds-of-function }

Las funciones vienen con el programa o de un archivo YAML suyo. Ambos tipos se
comportan igual al usarlos desde un botón.

| | Integrada | Personalizada |
|---|---|---|
| Procedencia | Incluida en el programa | Un archivo YAML escrito por usted |
| ¿Debe crear un archivo? | No | Sí, uno por función |
| Nombres | Reservados (`command`, `script`) | Cualquier nombre no reservado |
| ¿Siempre disponible? | Sí | Solo con `function_directory` |
| ¿Editable? | No | Sí, son sus archivos |

Puede combinar ambos tipos. Empiece con botones `command` y cree funciones
personalizadas cuando repita un comando con pequeños cambios.

### Funciones integradas { #built-in-functions }

Siempre se cargan dos funciones, aunque no defina `function_directory`. Sus
campos `command`, `path` y `args` se escriben directamente en el botón.

| Función | Qué hace | Obligatorio | Opcional |
|---------|----------|-------------|----------|
| [`command`](built-in/command.md) | Ejecuta un comando tal como está escrito | `command` | — |
| [`script`](built-in/script.md) | Ejecuta un script con argumentos | `path` | `args` |

Ambos nombres están **reservados**. Una función personalizada no puede
reutilizarlos: el cargador se detiene con un error como
`function name "command" is reserved` y el bot no se inicia.

### Funciones personalizadas { #custom-functions }

Una función personalizada es un archivo YAML que describe un comando
reutilizable. Guárdelo en una carpeta y apunte
[`function_directory`](../configuration.md#function_directory-rules) a ella.

!!! example "Indicar al bot dónde están las funciones"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

El bot lee la carpeta y sus subcarpetas al iniciarse. Carga cada archivo `.yaml`
y `.yml` e ignora los demás.

La versión contiene una carpeta `functions/` con cinco ejemplos:

| Función | Qué hace | Valores del botón |
|---------|----------|--------------------|
| [Script de eco](release-pack/echo-script.md) | Ejecuta un script mediante Bash | `path`, `args` opcional |
| [Ruta de disco](release-pack/disk-path.md) | Muestra el uso del disco | `path` opcional |
| [URL Curl](release-pack/curl-url.md) | Obtiene una URL | `url` |
| [Ping a un host](release-pack/ping-host.md) | Hace ping a un host | `host`, `count` opcional |
| [Unidad de journal](release-pack/journal-unit.md) | Muestra registros recientes de un servicio | `unit`, `lines` opcional |

Para escribir una, empiece por
[Estructura de archivos](write-your-own/file-structure.md) o siga la
[guía paso a paso](write-your-own/step-by-step.md).

!!! tip "Comprobar qué se cargó"

    ```bash title="Enumerar todas las funciones visibles para el bot"
    ./telegram-commander list-functions --config config.yaml
    ```

    Las funciones integradas muestran `source=builtin`; las personalizadas
    muestran su archivo de origen.

## Notas de seguridad

!!! warning "Los botones se ejecutan con los privilegios del bot"

    Los comandos usan los privilegios de la cuenta que ejecuta el bot. Si es
    root, como en la configuración predeterminada del
    [servicio](../installation/run-as-a-service.md), los botones pueden hacer
    cualquier cosa en el host. Añada solo
    [usuarios permitidos](../configuration.md#telegram) en quienes confíe.

    Los valores de los parámetros se insertan en el comando como texto.
    Trátelos como entrada del shell: prefiera valores fijos en los botones y
    use [`confirm: true`](../concepts/confirmation.md) para acciones
    destructivas.

!!! info "La salida larga se recorta y divide"

    Los comandos se detienen al alcanzar su `timeout` y el bot conserva como
    máximo `max_output_bytes` de su salida. Cualquier salida que supere un
    mensaje de Telegram llega en varios mensajes. Consulte
    [Configuración → Cuánta salida verá](../configuration.md#how-much-command-output-you-see).

## Relacionado

- [`command`](built-in/command.md) — ejecute un comando de shell
- [`script`](built-in/script.md) — ejecute un archivo de script
- [Guía paso a paso](write-your-own/step-by-step.md) — cree su primera función
- [Menú](../concepts/menu.md) — cómo hacen referencia los botones a funciones
- [Parámetro](../concepts/parameter.md) — los valores que necesita una función
