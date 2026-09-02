---
title: CLI-Referenz
description: Alle Befehlszeilenoptionen von Telegram Commander — vom Ausführen des Bots bis zum Prüfen und Formatieren Ihrer Konfigurationsdatei sowie das gemeinsam verwendete Flag --config.
icon: material/console-line
---

# :material-console-line: CLI-Referenz

!!! info "Befehle verwenden dieses Format"

    ```text title="Befehlssyntax"
    telegram-commander <command> [flags]
    ```

`--config` / `-c` ist für Befehle, die eine Konfiguration laden,
**erforderlich**. Der Pfad kann relativ zum aktuellen Arbeitsverzeichnis oder
absolut sein. Der Inhalt dieser Datei ist unter
[Konfiguration](configuration.md) dokumentiert.

Wenn Sie gerade erst beginnen, zeigt
[In der CLI ausführen](installation/download-and-run.md) diese Befehle in der
Reihenfolge, in der Sie sie verwenden werden.

## :material-format-list-bulleted-square: Befehle

### `run` { #run }

Den Bot im Vordergrund ausführen.

!!! info "Den Bot im Vordergrund starten"

    ```bash title="Den Bot ausführen"
    telegram-commander run --config /path/to/configfile.yaml
    ```

!!! note "Konfigurationsänderungen erfordern einen Neustart"

    Der Bot liest die Konfiguration einmalig beim Start. Starten Sie den Prozess
    nach einer Dateiänderung neu (zum Beispiel mit
    `systemctl restart telegram-commander`).

### `validate` { #validate }

Konfiguration, Funktionen und Schaltflächenverweise offline validieren.

!!! info "Die Konfiguration offline oder online prüfen"

    ```bash title="Die Konfiguration validieren"
    telegram-commander validate --config /path/to/configfile.yaml
    telegram-commander validate --config /path/to/configfile.yaml --online
    ```

!!! note "`--online` benötigt Internetzugriff"

    Mit `--online` fragt die Prüfung zusätzlich bei Telegram nach, ob das
    Bot-Token funktioniert. Der Rechner muss Telegram daher erreichen können.

### `version`

Die Programmversion ausgeben.

### `fmt`

Eine YAML-Konfigurationsdatei formatiert ausgeben.

!!! info "Formatiertes YAML ausgeben oder speichern"

    ```bash title="Eine Konfigurationsdatei formatieren"
    telegram-commander fmt --config /path/to/configfile.yaml
    telegram-commander fmt --config /path/to/configfile.yaml -w
    ```

### `environ`

Umgebungsvariablen des Prozesses ausgeben (hilfreich bei der Fehlersuche in
Dienst-Units).

### `list-functions` { #list-functions }

Integrierte und geladene eigene Funktionen auflisten. Damit können Sie prüfen,
ob Ihre eigenen Funktionsdateien gefunden wurden. Siehe
[Funktionen](functions/index.md).

!!! info "Alle verfügbaren Funktionen anzeigen"

    ```bash title="Verfügbare Funktionen auflisten"
    telegram-commander list-functions --config /path/to/configfile.yaml
    ```

### `completion`

Skripte zur Shell-Vervollständigung erzeugen:

!!! info "Die verwendete Shell auswählen"

    ```bash title="Ein Vervollständigungsskript erzeugen"
    telegram-commander completion bash
    telegram-commander completion zsh
    telegram-commander completion fish
    telegram-commander completion powershell
    ```

### `manpage`

Eine Handbuchseite auf stdout schreiben.

## Verwandte Seiten

- [Konfiguration](configuration.md) — die an `--config` übergebene Datei
- [Eigene Funktionen](functions/index.md#custom-functions) — was `list-functions` anzeigt
- [Als Dienst ausführen](installation/run-as-a-service.md) — `run` unter systemd ausführen
