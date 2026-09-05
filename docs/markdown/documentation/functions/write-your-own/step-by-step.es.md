---
title: Guía paso a paso
description: >-
  Cree desde cero su primera función de Telegram Commander, añada un botón,
  valide la configuración y ejecute el bot para ver el resultado.
icon: material/format-list-numbered
---

# :material-format-list-numbered: Guía paso a paso

Este recorrido lleva una función pequeña desde una carpeta vacía hasta un botón
funcional. Tarda pocos minutos y solo necesita su archivo de configuración.

## 1. Apunte la configuración a una carpeta

!!! example "Añadir la carpeta a la configuración"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

La ruta es relativa al archivo de configuración. Consulte
[Configuración → Reglas de `function_directory`](../../configuration.md#function_directory-rules).

## 2. Escriba el archivo de la función

Cree `functions/greet.yaml`. El parámetro se llama `args` para que un botón
pueda completarlo:

!!! example "Su primera función"

    ```yaml title="functions/greet.yaml"
    name: greet
    run: "echo Hello {{.args}}"
    params:
      - name: args
        required: true
        description: Name to greet
    ```

## 3. Añada un botón que la use

!!! example "Un botón para la función nueva"

    ```yaml title="Botón Say hello"
    - name: Say hello
      type: button
      function: greet
      args: "world"
    ```

## 4. Valide y enumere las funciones

!!! example "Comprobar la configuración y ver qué se cargó"

    ```bash title="Validar y después enumerar las funciones"
    ./telegram-commander validate --config config.yaml
    ./telegram-commander list-functions --config config.yaml
    ```

Debería ver una línea `greet`. Cada línea muestra el nombre, la procedencia de
la función y cuántos parámetros tiene. Consulte
[CLI → list-functions](../../cli.md#list-functions).

Si `validate` muestra una queja, el mensaje indica el botón y el valor que
falta; consulte [Reglas](rules.md).

## 5. Ejecute el bot y toque el botón

!!! example "Iniciar el bot en primer plano"

    ```bash title="Ejecutar y observar la salida"
    ./telegram-commander run --config config.yaml
    ```

Envíe `/start` en Telegram, toque **Say hello** y el bot ejecutará
`echo Hello world` y devolverá `Hello world` como bloque de código.

## 6. Hágala reutilizable

Cambie el valor `args` del botón o añada otro botón con un valor distinto. La
función no cambia:

!!! example "Dos botones, una función"

    ```yaml title="Botones de saludo"
    - name: Greet world
      type: button
      function: greet
      args: "world"
    - name: Greet team
      type: button
      function: greet
      args: "team"
    ```

## Relacionado

- [Variables de reemplazo](placeholders.md) — añada partes opcionales al comando
- [Reglas](rules.md) — lo que acepta el cargador
- [Estructura de archivos](file-structure.md) — explicación de todos los campos
- [Menú](../../concepts/menu.md) — dónde colocar el botón en su menú
