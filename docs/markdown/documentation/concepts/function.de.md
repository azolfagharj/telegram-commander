---
title: Funktion
description: Eine Funktion wandelt die Werte einer Schaltfläche in einen Shell-Befehl um. Verwenden Sie die integrierten Funktionen command und script oder fügen Sie eigene Funktionen als YAML-Dateien hinzu.
icon: material/function
---

# :material-function: Funktion

Das, was ausgeführt wird, wenn eine [Schaltfläche](button.md) angetippt wird.
Eine [Funktion](function.md) nimmt einige [Parameter](parameter.md) entgegen und
wandelt sie in einen Shell-Befehl um.

Es gibt zwei Arten:

- **Integrierte Funktionen** werden mit dem Programm ausgeliefert (`command`,
  `script`). Sie müssen sie nicht erstellen.
- **Eigene Funktionen** sind zusätzliche Funktionen, die Sie als YAML-Dateien
  in `function_directory` definieren.

Wenn Sie eine Schaltfläche antippen, sucht der Bot die Funktion, setzt die
Werte der Schaltfläche ein und führt das Ergebnis in der [Shell](shell.md) aus.

!!! example "Eine Schaltfläche mit der integrierten Funktion `command`"

    ```yaml title="Laufzeit-Schaltfläche"
    - name: Laufzeit
      type: button
      function: command
      command: "uptime"
    ```

## Konfiguration

Informationen zu `function_directory` und den Schaltflächenfeldern, die auf
Funktionen verweisen, finden Sie unter [Konfiguration](../configuration.md).

## Verwandte Themen

- [Parameter](parameter.md) — benannte Werte, die eine Funktion benötigt
- [Funktionen](../functions/index.md#two-kinds-of-function) — integrierte und eigene Funktionen im Detail
- [Menü](menu.md) — wie Schaltflächen auf Funktionen verweisen
