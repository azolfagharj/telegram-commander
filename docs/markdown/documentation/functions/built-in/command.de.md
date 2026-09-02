---
title: command
description: Die integrierte Funktion command führt einen einzelnen Shell-Befehl genau so aus, wie er auf der Schaltfläche steht, einschließlich Pipes, Umleitungen und verketteter Befehle.
icon: material/console
---

# :material-console: `command`

`command` ist eine [integrierte Funktion](../index.md#built-in-functions). Sie
führt einen einzelnen Shell-Befehl genau so aus, wie Sie ihn auf der
Schaltfläche angegeben haben.

| Parameter | Erforderlich | Standardwert | Bedeutung |
|-----------|----------|---------|---------|
| `command` | ja | — | Der auszuführende Shell-Befehl |

!!! example "Einen Befehl unverändert ausführen"

    ```yaml title="Schaltfläche zum Anzeigen des Kernels"
    - name: Kernel anzeigen
      type: button
      function: command
      command: "uname -a"
    ```

## Pipes und verkettete Befehle

Befehle werden über [`/bin/bash -c`](../../concepts/shell.md) ausgeführt.
Deshalb funktionieren Pipes, Umleitungen und `&&` innerhalb eines einzelnen
`command`-Werts.

!!! example "Pipes funktionieren innerhalb von Befehlen"

    ```yaml title="Schaltfläche für die wichtigsten Prozesse"
    - name: Wichtigste Prozesse
      type: button
      function: command
      command: "ps aux --sort=-%mem | head -n 10"
    ```

!!! tip "Verwenden Sie für wiederholte Befehle eine eigene Funktion"

    Wenn mehrere Schaltflächen dieselbe Befehlsstruktur wiederholen, erstellen
    Sie eine eigene Funktion und geben Sie auf jeder Schaltfläche nur die
    veränderlichen Werte an. Eigene Parameternamen wie `url` oder `host` können
    direkt auf diesen Schaltflächen stehen.

## Verwandte Themen

- [`script`](script.md) — eine Skriptdatei statt eines Inline-Befehls ausführen
- [Integrierte Funktionen](../index.md#built-in-functions) — beide integrierten Funktionen
- [Shell](../../concepts/shell.md) — wie Befehle ausgeführt werden
- [Bestätigung](../../concepts/confirmation.md) — vor einem riskanten Befehl nachfragen
