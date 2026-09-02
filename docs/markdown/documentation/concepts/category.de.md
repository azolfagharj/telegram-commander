---
title: Kategorie
description: Gruppieren Sie Schaltflächen in Untermenüs. Eine Kategorie enthält unter items weitere Schaltflächen oder Kategorien, damit ein langes Menü kurz und auf dem Smartphone leicht bedienbar bleibt.
icon: material/folder-outline
---

# :material-folder-outline: Kategorie

Ein Menüknoten, der ein Untermenü öffnet, statt etwas auszuführen. Er besitzt
`items` (weitere [Schaltflächen](button.md) oder Kategorien) statt einer
[Funktion](function.md).

!!! example "Eine Kategorie mit einer Schaltfläche"

    ```yaml title="Eine System-Kategorie"
    - name: System
      type: category
      items:
        - name: Laufzeit
          type: button
          function: command
          command: "uptime"
    ```

Schaltflächen und Kategorien bilden gemeinsam einen Baum. Siehe [Menü](menu.md).

## Konfiguration

Informationen zu Kategoriefeldern (`type`, `items`, `columns` und weiteren)
finden Sie unter [Konfiguration → Menü](../configuration.md#menu).

## Verwandte Themen

- [Schaltfläche](button.md) — führt beim Antippen eine Funktion aus
- [Menü](menu.md) — Verschachtelung, Layout und Seitennavigation
