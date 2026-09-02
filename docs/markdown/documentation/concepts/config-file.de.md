---
icon: material/file-document-outline
title: Konfigurationsdatei
description: Eine YAML-Datei enthält Ihr Bot-Token, die zugelassenen Benutzer und das Schaltflächenmenü. Übergeben Sie sie mit --config an run, validate, fmt und list-functions.
---

# :material-file-document-outline: Konfigurationsdatei

Eine einzelne YAML-Datei beschreibt alles: Ihr Bot-Token, wer den Bot verwenden
darf und das Menü aus [Schaltflächen](button.md). Sie übergeben sie mit
`--config` an die Befehle, die sie einlesen: `run`, `validate`, `fmt` und
`list-functions`. Andere Befehle wie `version` und `completion` verwenden keine
Konfigurationsdatei.

!!! example "Eine funktionierende Konfiguration mit einer Schaltfläche"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: Laufzeit
        type: button
        function: command
        command: "uptime"
    ```

## Verwandte Themen

- [Konfiguration](../configuration.md) — alle Felder, Standardwerte und Validierungsregeln
- [Menü](menu.md) — der Baum aus Schaltflächen und Kategorien
- [CLI](../cli.md) — die Datei mit `--config` übergeben
- [In der CLI ausführen](../installation/download-and-run.md) — Ihre erste Konfiguration erstellen
