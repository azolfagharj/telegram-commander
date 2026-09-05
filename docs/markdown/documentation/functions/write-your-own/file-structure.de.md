---
title: Dateistruktur
description: Alle Felder einer Telegram-Commander-Funktionsdatei erklärt — von name und run bis zur params-Liste mit type, required, default und description.
icon: material/file-tree
---

# :material-file-tree: Dateistruktur

Eine Datei beschreibt eine Funktion. Legen Sie die Datei an einer beliebigen
Stelle innerhalb Ihres
[`function_directory`](../../configuration.md#function_directory-rules) ab,
auch in einem Unterordner, und verwenden Sie die Endung `.yaml` oder `.yml`.
Der Dateiname selbst ist unerheblich: Der Funktionsname stammt aus dem Feld
`name` in der Datei.

!!! example "Eine eigene Funktion pro Datei definieren"

    ```yaml title="Datei einer eigenen Funktion"
    name: my-function          # erforderlich, muss eindeutig und nicht reserviert sein
    run: "echo {{.args}}"      # erforderlich, der auszuführende Befehl
    params:                    # optionale Parameterliste
      - name: args             # für jeden Parameter erforderlich
        type: string           # optional: string (Standard), int oder bool
        required: true         # optional, Standardwert false
        default: ""            # optionaler Wert, falls nicht angegeben
        description: Ein Text  # optionale Notiz für Sie selbst
    ```

## Felder auf oberster Ebene

| Feld | Erforderlich | Bedeutung |
|-------|----------|---------|
| `name` | ja | Der Name, den Schaltflächen in ihrem Feld `function` verwenden |
| `run` | ja | Der auszuführende Befehl mit [Platzhaltern](placeholders.md) für Werte |
| `params` | nein | Liste der benannten Werte, die die Funktion akzeptiert |

Eine Funktion ganz ohne `params` ist gültig; `run` ist dann ein fester Befehl.

## Parameterfelder

| Feld | Erforderlich | Standardwert | Bedeutung |
|-------|----------|---------|---------|
| `name` | ja | — | Der in `{{.name}}` innerhalb von `run` verwendete Name |
| `type` | nein | `string` | Werttyp: `string`, `int` oder `bool` |
| `required` | nein | `false` | Eine Schaltfläche muss einen Wert bereitstellen |
| `default` | nein | leer | Wird verwendet, wenn kein Wert angegeben wurde |
| `description` | nein | leer | Eine Notiz für Sie selbst, die nicht in Telegram angezeigt wird |

!!! info "`type` wird geprüft"

    Als `int` deklarierte Werte und Standardwerte müssen eine Ganzzahl
    enthalten. Als `bool` deklarierte Werte und Standardwerte müssen einen
    gültigen booleschen Wert enthalten. Ungültige Werte lassen
    [`validate`](../../cli.md#validate) fehlschlagen.

!!! info "Schaltflächenschlüssel entsprechen den Parameternamen"

    Schreiben Sie jeden Parameter mit demselben Namen direkt auf eine
    Schaltfläche. Namen wie `url`, `host` und `lines` funktionieren neben
    `command`, `path` und `args`. Siehe
    [Werte von einer Schaltfläche übergeben](../index.md#passing-values-from-a-button).

## Ordnerstruktur

Sie können den Ordner beliebig organisieren:

!!! example "Unterordner werden ebenfalls gelesen"

    ```text title="functions/"
    functions/
      disk.yaml
      logs/
        nginx.yaml
        app.yml
    ```

Alle drei Dateien werden geladen. Dateien mit anderen Endungen werden
übersprungen.

## Verwandte Themen

- [Regeln](rules.md) — was der Loader ablehnt
- [Platzhalter](placeholders.md) — den Befehl `run` schreiben
- [Schritt-für-Schritt-Anleitung](step-by-step.md) — Ihre erste Funktion erstellen
