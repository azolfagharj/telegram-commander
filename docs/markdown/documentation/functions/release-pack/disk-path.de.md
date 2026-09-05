---
title: Datenträgerpfad
description: Die mitgelieferte Funktion disk-path zeigt mit df -h die Datenträgerbelegung. Ihr einziger Parameter ist path; wenn Sie ihn auslassen, wird das Wurzeldateisystem verwendet.
icon: material/harddisk
---

# :material-harddisk: Datenträgerpfad

`disk-path` zeigt mit `df -h` die Datenträgerbelegung für einen Pfad. Die
Funktion gehört zu den
[mitgelieferten Funktionen](../index.md#custom-functions), die Sie unverändert
von einer Schaltfläche aus verwenden können.

- Führt aus: `df -h {{.path}}`
- `path` (optional, Standardwert `/`): zu prüfender Dateisystempfad

!!! example "Die Funktionsdatei"

    ```yaml title="functions/disk-path.yaml"
    name: disk-path
    run: "df -h {{.path}}"
    params:
      - name: path
        type: string
        required: false
        default: "/"
        description: Zu prüfender Dateisystempfad
    ```

Da `path` optional ist und einen Standardwert besitzt, kann eine Schaltfläche
das Feld auslassen und funktioniert dennoch.

!!! example "Den Standardpfad verwenden oder einen Pfad auswählen"

    ```yaml title="Schaltflächen für die Datenträgerbelegung"
    - name: Disk (root)
      type: button
      function: disk-path        # kein Feld path: verwendet den Standardwert "/"
    - name: Disk (var)
      type: button
      function: disk-path
      path: "/var"
    ```

Die erste Schaltfläche führt `df -h /`, die zweite `df -h /var` aus.

## Verwandte Themen

- [Regeln](../write-your-own/rules.md) — Verhalten von Standardwerten und erforderlichen Werten
- [Eigene Funktionen](../index.md#custom-functions) — alle fünf mitgelieferten Beispiele
- [`command`](../built-in/command.md) — stattdessen einen vollständigen Befehl schreiben
