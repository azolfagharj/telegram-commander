---
title: script
description: Die integrierte Funktion script führt eine Skriptdatei auf dem Server aus und fügt optionale Argumente hinter dem Pfad hinzu. Das Skript muss ausführbar sein.
icon: material/script-text
---

# :material-script-text: `script`

`script` ist eine [integrierte Funktion](../index.md#built-in-functions). Sie
führt eine Skriptdatei mit optionalen Argumenten aus.

| Parameter | Erforderlich | Standardwert | Bedeutung |
|-----------|----------|---------|---------|
| `path` | ja | — | Pfad zum Skript |
| `args` | nein | leer | Argumente, die hinter dem Pfad übergeben werden |

!!! example "Ein Skript mit Argumenten ausführen"

    ```yaml title="Schaltfläche für den nächtlichen Bericht"
    - name: Nightly report
      type: button
      function: script
      path: "/usr/local/bin/report.sh"
      args: "--today"
    ```

Dies führt `/usr/local/bin/report.sh --today` aus. Wenn Sie `args` weglassen,
wird nur `/usr/local/bin/report.sh` ausgeführt.

## Das Skript muss ausführbar sein

Der Pfad wird direkt ausgeführt. Daher benötigt die Datei das Ausführungsbit:

!!! tip "Das Skript einmalig ausführbar machen"

    ```bash title="Die Ausführung der Datei erlauben"
    chmod +x /usr/local/bin/report.sh
    ```

Wenn Sie die Datei nicht ändern können, führen Sie sie stattdessen über einen
Interpreter aus. Das Beispiel [`echo-script`](../release-pack/echo-script.md)
der mitgelieferten Funktionen tut genau dies, indem es zuerst `bash` aufruft.

## Verwandte Themen

- [`command`](command.md) — einen Inline-Befehl statt einer Datei ausführen
- [Echo-Skript](../release-pack/echo-script.md) — ein Skript über Bash ausführen
- [Integrierte Funktionen](../index.md#built-in-functions) — beide integrierten Funktionen
- [Shell](../../concepts/shell.md) — wie Befehle ausgeführt werden
