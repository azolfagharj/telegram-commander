---
title: Regeln
description: Die Prüfungen, die eine Telegram-Commander-Funktionsdatei bestehen muss — zulässige Namen, reservierte Wörter, eindeutige Namen, Ablehnung von Tippfehlern und Verwendung von Standardwerten.
icon: material/format-list-checks
---

# :material-format-list-checks: Regeln

Jede Funktionsdatei wird beim Start des Bots und bei der Ausführung von
[`validate`](../../cli.md#validate) geprüft. Ist eine Datei fehlerhaft, wird
nichts ausgeführt. Fehler werden somit erkannt, bevor Ihr Menü aktiv wird.

## Gelesene Dateien

- Nur Dateien mit der Endung `.yaml` oder `.yml` werden gelesen. Alle anderen werden ignoriert.
- Unterordner Ihres `function_directory` werden ebenfalls gelesen.
- Eine Datei muss genau eine Funktion beschreiben.

## Namensgebung

- **`name` ist erforderlich** und muss `^[A-Za-z0-9._-]+$` entsprechen —
  Buchstaben, Ziffern, `.`, `-` und `_`. Leerzeichen sind nicht zulässig.
- **Bei Namen wird die Groß-/Kleinschreibung ignoriert und sie müssen eindeutig
  sein.** `Deploy` und `deploy` kollidieren, selbst wenn sie in verschiedenen
  Dateien stehen.
- **Reservierte Namen sind verboten.** Eine Funktion darf nicht `command` oder
  `script` heißen; diese gehören zu den
  [integrierten Funktionen](../index.md#built-in-functions).

## Der Befehl

- **`run` ist erforderlich.** Schreiben Sie den Befehl mit
  [Platzhaltern](placeholders.md) an den Positionen der Werte.
- Die Platzhaltersyntax wird beim Validieren der Konfiguration geprüft, sodass
  eine fehlerhafte öffnende Klammer `{{` frühzeitig gemeldet wird.

## Parameter

- **Jeder Parameter benötigt einen `name`**, der derselben Zeichenregel wie der
  Funktionsname entspricht.
- **`type` muss** `string`, `int` oder `bool` sein. Ohne Angabe wird `string`
  verwendet. Als `int` oder `bool` deklarierte Werte und Standardwerte werden
  geprüft.
- Ein Parametername darf keiner Schaltflächeneinstellung entsprechen: `name`,
  `type`, `icon`, `id`, `function`, `confirm`, `timeout`, `workdir`, `env`,
  `columns` oder `items`. Die Namen `command`, `path` und `args` sind zulässig.
- Jeder auf einer Schaltfläche angegebene Wert muss einem von ihrer Funktion
  deklarierten Parameter entsprechen.

Das Feld `run` der Funktion darf nur Platzhalter für deklarierte Parameter
verwenden. Ein nicht deklarierter Platzhalter ist ein Fehler.

## Tippfehler sind Fehler

**Unbekannte Schlüssel werden abgelehnt.** Ein Tippfehler wie `requird:` statt
`required:` stoppt die Validierung, anstatt stillschweigend ignoriert zu werden.

!!! warning "Diese Datei wird nicht geladen"

    ```yaml title="functions/broken.yaml"
    name: broken
    run: "echo {{.args}}"
    params:
      - name: args
        requird: true      # abgelehnt: unbekannter Schlüssel
    ```

## Erforderliche Werte und Standardwerte

- Ein **erforderlicher Parameter ohne Wert** lässt die Validierung mit einer
  Meldung wie `required parameter "args" for function "greet" is missing`
  fehlschlagen.
- Ein vorhandener, aber leerer Wert gilt als fehlend.
- Ein **optionaler Parameter** verwendet seinen `default`. Ohne `default` wird
  er leer. So werden optionale Befehlsteile ausgelassen.

!!! example "Erforderlicher und optionaler Wert in einer Datei"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n {{.args}} {{.path}}"
    params:
      - name: path
        required: true
        description: Pfad zur Protokolldatei
      - name: args
        default: "200"
        description: Anzahl der Zeilen
    ```

    Eine Schaltfläche, die nur `path` festlegt, führt
    `tail -n 200 /var/log/app.log` aus.

## Verwandte Themen

- [Dateistruktur](file-structure.md) — Erklärung aller Felder
- [Platzhalter](placeholders.md) — den Befehl `run` schreiben
- [CLI → validate](../../cli.md#validate) — die Prüfungen selbst ausführen
