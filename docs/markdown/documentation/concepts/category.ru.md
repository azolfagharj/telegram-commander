---
title: Категория
description: Группируйте кнопки в подменю. Категория содержит больше кнопок или категорий под элементами, поэтому длинное меню остается коротким и его легко нажимать на телефоне.
icon: material/folder-outline
---
# :material-folder-outline: Категория

Узел меню, который открывает подменю вместо запуска чего-либо. Имеет `items`
(больше [кнопки](button.md) или категорий) вместо [функция](function.md).

!!! example "Категория, удерживающая одну кнопку"

    ```yaml title="A System category"
    - name: System
      type: category
      items:
        - name: Uptime
          type: button
          function: command
          command: "uptime"
    ```

Кнопки и категории вместе образуют дерево. См. [Меню](menu.md).

## Конфигурация

Поля категорий (`type`, `items`, `columns` и т. д.) см.
[Конфигурация → Меню](../configuration.md#menu).

## Похожие

- [Кнопка](button.md) — запускает функцию при нажатии
- [Меню](menu.md) — вложение, макет и нумерация страниц.
