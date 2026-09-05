---
title: Catégorie
description: Regroupez les boutons dans des sous-menus. Une catégorie contient d’autres boutons ou catégories sous items pour garder un menu court et facile à utiliser.
icon: material/folder-outline
---

# :material-folder-outline: Catégorie

Un élément de menu qui ouvre un sous-menu au lieu d’exécuter une action. Il
possède des `items` (d’autres [boutons](button.md) ou catégories) plutôt qu’une
[fonction](function.md).

!!! example "Une catégorie contenant un bouton"

    ```yaml title="Une catégorie Système"
    - name: System
      type: category
      items:
        - name: Uptime
          type: button
          function: command
          command: "uptime"
    ```

Les boutons et les catégories forment ensemble une arborescence. Consultez
[Menu](menu.md).

## Configuration

Pour les champs d’une catégorie (`type`, `items`, `columns`, etc.), consultez
[Configuration → Menu](../configuration.md#menu).

## Pages associées

- [Bouton](button.md) — exécute une fonction lorsque vous appuyez dessus
- [Menu](menu.md) — imbrication, disposition et pagination
