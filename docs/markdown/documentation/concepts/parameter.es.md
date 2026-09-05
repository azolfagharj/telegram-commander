---
icon: material/tune
title: Parámetro
description: >-
  Un parámetro es un valor con nombre que necesita una función. Escriba cada
  valor directamente en el botón con el nombre declarado por la función.
---

# :material-tune: Parámetro

Un valor con nombre que necesita una [función](function.md). Por ejemplo, la
función integrada `command` necesita un parámetro llamado `command`. Usted
proporciona los valores de los parámetros en el [botón](button.md).

Algunos parámetros son obligatorios; otros son opcionales y tienen un valor
predeterminado definido en el archivo de la función.

## Pasar un valor desde un botón

Escriba el nombre del parámetro como clave directamente en el botón. Los campos
integrados `command`, `path` y `args` siguen esta regla, al igual que los nombres
de parámetros personalizados.

!!! example "Proporcionar nombres de parámetros personalizados"

    ```yaml title="Registros recientes del servicio"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`unit` y `lines` deben estar declarados por `journal-unit`. El valor numérico de
`lines` no necesita comillas.

!!! warning "No use un mapa `params:` anidado"

    Un archivo de función personalizada usa `params:` para declarar sus
    parámetros. Un botón no. Escriba `host:`, `unit:` o cualquier otro valor
    directamente en el botón.

## Validación

[`validate`](../cli.md#validate) comprueba que:

- cada parámetro obligatorio tenga un valor;
- cada valor del botón esté declarado por su función;
- los valores declarados como `int` contengan un entero;
- los valores declarados como `bool` contengan un booleano válido;
- los valores predeterminados coincidan con el tipo declarado.

Los parámetros opcionales usan su valor predeterminado cuando el botón los
omite.

Los nombres de parámetros no pueden coincidir con ajustes del botón: `name`,
`type`, `icon`, `id`, `function`, `confirm`, `timeout`, `workdir`, `env`,
`columns` o `items`. Los nombres `command`, `path` y `args` están permitidos.

Consulte [Funciones → Pasar valores desde un botón](../functions/index.md#passing-values-from-a-button)
para ver la explicación completa y ejemplos.

## Configuración

Para los ajustes de botón y las claves de parámetros, consulte
[Configuración → Menú](../configuration.md#menu).

## Relacionado

- [Función](function.md) — lo que usa los parámetros
- [Reglas](../functions/write-your-own/rules.md) — reglas de parámetros para funciones personalizadas
