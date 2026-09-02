---
title: Echo Script
description: Die mitgelieferte Funktion echo-script führt eine Skriptdatei über Bash aus, sodass die Datei nicht ausführbar sein muss. Sie übernimmt path und args von einer Schaltfläche.
icon: material/bash
---

# :material-bash: Echo Script

`echo-script` führt ein Skript über `bash` aus, sodass die Skriptdatei selbst
kein Ausführungsbit benötigt. Die Funktion gehört zu den
[mitgelieferten Funktionen](../index.md#custom-functions), die Sie unverändert
von einer Schaltfläche aus verwenden können.

- Führt aus: `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (erforderlich): Pfad zur Skriptdatei
- `args` (optional): zusätzliche Argumente

!!! example "Die Funktionsdatei"

    ```yaml title="functions/echo-script.yaml"
    name: echo-script
    run: "bash {{.path}}{{if .args}} {{.args}}{{end}}"
    params:
      - name: path
        type: string
        required: true
        description: Pfad zur Skriptdatei
      - name: args
        type: string
        required: false
        description: Optionale Skriptargumente
    ```

Schreiben Sie beide Parameternamen direkt auf die Schaltfläche.

!!! example "Ein Skript über Bash ausführen"

    ```yaml title="Schaltfläche zum Ausführen der Bereinigung"
    - name: Bereinigung ausführen
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

Diese Schaltfläche führt `bash /opt/scripts/cleanup.sh --verbose` aus. Lassen
Sie `args` weg, wird `bash /opt/scripts/cleanup.sh` ausgeführt, da der Abschnitt
`{{if .args}}` bei einem leeren Wert übersprungen wird.

## Verwandte Themen

- [`script`](../built-in/script.md) — ein ausführbares Skript direkt ausführen
- [Platzhalter](../write-your-own/placeholders.md) — Funktionsweise von `{{if .args}}`
- [Eigene Funktionen](../index.md#custom-functions) — alle fünf mitgelieferten Beispiele
