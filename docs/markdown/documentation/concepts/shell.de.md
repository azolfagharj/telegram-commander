---
title: Shell
description: Jeder Befehl wird über eine Shell ausgeführt, standardmäßig /bin/bash, sodass Pipes und Umleitungen funktionieren. Ändern Sie sie einmalig auf der obersten Ebene Ihrer Konfigurationsdatei.
icon: material/console-line
---

# :material-console-line: Shell

Das Programm, das Ihren Befehl ausführt, standardmäßig `/bin/bash`. Befehle
werden als `bash -c "your command"` ausgeführt, sodass Pipes und Umleitungen
funktionieren.

Sie können die Shell global mit dem Schlüssel `shell` in Ihrer
[Konfigurationsdatei](config-file.md) ändern.

!!! example "Die Shell auswählen"

    ```yaml title="Oberste Ebene von config.yaml"
    shell: /bin/bash
    ```

Jede [Funktion](function.md) führt ihren Befehl über diese Shell aus.

## Konfiguration

Informationen zum Feld `shell` und verwandten Optionen auf oberster Ebene
finden Sie unter [Konfiguration → Felder auf oberster Ebene](../configuration.md#root-fields).

## Verwandte Themen

- [Funktion](function.md) — erzeugt den Befehl, den die Shell ausführt
- [Konfigurationsdatei](config-file.md) — hier wird `shell` festgelegt
