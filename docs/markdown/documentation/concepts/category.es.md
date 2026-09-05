---
title: Categoría
description: >-
  Agrupe botones en submenús. Una categoría contiene más botones o
  categorías bajo items para mantener un menú largo breve y fácil de usar.
icon: material/folder-outline
---

# :material-folder-outline: Categoría

Un nodo del menú que abre un submenú en lugar de ejecutar algo. Contiene `items`
(más [botones](button.md) o categorías) en vez de una [función](function.md).

!!! example "Una categoría que contiene un botón"

    ```yaml title="Una categoría System"
    - name: System
      type: category
      items:
        - name: Uptime
          type: button
          function: command
          command: "uptime"
    ```

Los botones y las categorías forman un árbol. Consulte [Menú](menu.md).

## Configuración

Para conocer los campos de una categoría (`type`, `items`, `columns` y otros),
consulte [Configuración → Menú](../configuration.md#menu).

## Relacionado

- [Botón](button.md) — ejecuta una función al tocarlo
- [Menú](menu.md) — anidamiento, disposición y paginación
