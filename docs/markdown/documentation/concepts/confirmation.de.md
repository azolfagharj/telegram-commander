---
title: Bestätigung
description: Lassen Sie den Bot vor dem Ausführen einer Schaltfläche nachfragen. Aktivieren Sie confirm für nicht rückgängig zu machende Aktionen und legen Sie fest, wie lange die Aufforderung mit Yes oder Cancel gültig bleibt.
icon: material/help-circle-outline
---

# :material-help-circle-outline: Bestätigung

Ein optionaler Schritt „Sind Sie sicher?“, bevor eine
[Schaltfläche](button.md) ausgeführt wird. Er wird mit `confirm: true`
aktiviert und eignet sich für destruktive Aktionen.

Der Bot fragt (`Confirm: … ?`) mit Yes / Cancel sowie Home (und Back, wenn Sie
sich in einer Kategorie befinden). Wenn Run Command aktiviert ist, bleibt diese
Schaltfläche auch auf dem Bestätigungsbildschirm. Bis Sie Yes antippen, wird
nichts ausgeführt; Cancel, Home und Back lassen den Befehl unangetastet.

Die Aufforderung läuft nach einer Weile ab (standardmäßig nach 5 Minuten).
Ändern Sie dies mit `confirm_ttl` in Ihrer
[Konfigurationsdatei](config-file.md). Ist sie abgelaufen, tippen Sie die
Schaltfläche erneut an, um eine neue Aufforderung zu erhalten.

!!! warning "Verwenden Sie dies für alles, was sich nicht rückgängig machen lässt"

    Eine Menüschaltfläche ist auf dem Smartphone nur ein Antippen entfernt, und
    alle zugelassenen Benutzer sehen dasselbe Menü. Setzen Sie `confirm: true`
    für Schaltflächen, die einen Dienst stoppen, Daten löschen oder den Rechner
    neu starten.

!!! example "Eine Schaltfläche, die vor der Ausführung nachfragt"

    ```yaml title="Schaltfläche zum Stoppen von nginx"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## Konfiguration

Informationen zu `confirm` für Schaltflächen und zum globalen `confirm_ttl`
finden Sie unter [Konfiguration](../configuration.md).

## Verwandte Themen

- [Schaltfläche](button.md) — der Knoten, der eine Bestätigung verlangen kann
- [Menü → Bestätigung](menu.md#confirmation) — Beispiele und Verhalten
