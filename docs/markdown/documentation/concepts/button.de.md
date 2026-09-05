---
title: Schaltfläche
description: Ein antippbarer Menüeintrag, der eine Funktion ausführt. Schreiben Sie jeden Funktionswert direkt auf die Schaltfläche und verwenden Sie den von der Funktion deklarierten Parameternamen.
icon: material/gesture-tap-button
---

# :material-gesture-tap-button: Schaltfläche

Ein antippbarer Eintrag im Telegram-Menü. Eine Schaltfläche besitzt einen
`name` und führt beim Antippen eine [Funktion](function.md) auf Ihrem Server
aus. Das Ergebnis wird an den Chat zurückgesendet.

Schaltflächen und [Kategorien](category.md) bilden gemeinsam den
[Menü](menu.md)-Baum unter dem obersten Schlüssel `menu` in Ihrer
[Konfigurationsdatei](config-file.md). Eine Schaltfläche erledigt die Arbeit;
eine Kategorie öffnet nur ein Untermenü.

## Darstellung einer Schaltfläche

Schaltflächen erscheinen auf der Tastatur unter dem Nachrichtenfeld,
standardmäßig zwei pro Zeile. Der von Ihnen gewählte Name ist der Text auf der
Taste. Halten Sie ihn kurz genug, damit er auf einem Smartphone lesbar bleibt.
Siehe [Menü → Darstellung des Telegram-Menüs](menu.md#how-the-telegram-menu-looks).

## :material-format-list-checks: Bestandteile einer Schaltfläche

!!! example "Jede Zeile steuert einen Bestandteil der Schaltfläche"

    ```yaml title="Eine vollständig beschriftete Schaltfläche"
    - name: Restart nginx # (1)!
      type: button # (2)!
      icon: "🔄" # (3)!
      function: command # (4)!
      command: "systemctl restart nginx" # (5)!
      confirm: true # (6)!
    ```

    1.  Der Text auf der Taste in Telegram. Er muss unter den benachbarten
        Einträgen desselben Menüs eindeutig sein.
    2.  Hier steht immer `button`. Verwenden Sie `category`, wenn Sie ein
        Untermenü wünschen.
    3.  Optionales Emoji vor dem Namen. Es dient nur der Darstellung und ändert
        nicht, was ausgeführt wird.
    4.  Die zu verwendende [Funktion](function.md). `command` ist die
        integrierte Funktion zum Ausführen eines Shell-Befehls.
    5.  Was die Funktion `command` ausführt. Hier funktioniert alles, was Sie
        in ein Terminal eingeben könnten.
    6.  Optional. Vor der Ausführung wird „Are you sure?“ gefragt. Lassen Sie
        dies bei reinen Leseaktionen weg.

## Was beim Antippen geschieht

1.  Der Bot sendet eine kurze Zeile **Running**, damit Sie wissen, dass die
    Ausführung begonnen hat.
2.  Der Befehl wird auf dem Rechner ausgeführt, auf dem der Bot läuft.
3.  Die Ausgabe kommt als Codeblock mit Exit-Code und Ausführungsdauer zurück.
    Lange Ausgaben werden in mehreren aufeinanderfolgenden Nachrichten gesendet.
4.  Sie bleiben im aktuellen Menü, sodass **Back** diese Kategorie weiterhin
    verlässt.

## :material-code-braces: Häufige Schaltflächen

=== "Etwas prüfen"

    ```yaml title="Laufzeit-Schaltfläche"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

=== "Einen Dienst verwalten"

    ```yaml title="Schaltfläche zum Neustarten von nginx"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

=== "Ein Protokoll lesen"

    ```yaml title="Schaltfläche für das nginx-Protokoll"
    - name: Nginx log
      type: button
      function: command
      command: "journalctl -u nginx -n 50 --no-pager"
    ```

=== "Ein Skript ausführen"

    ```yaml title="Schaltfläche für die nächtliche Sicherung"
    - name: Nightly backup
      type: button
      function: script
      path: "/usr/local/bin/backup.sh"
    ```

=== "Eine destruktive Aktion"

    ```yaml title="Schaltfläche zum Stoppen von nginx"
    - name: Stop nginx
      type: button
      icon: "🛑"
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## :material-emoticon-outline: Symbole

`icon` setzt ein Emoji vor den Namen. Es ist rein dekorativ und kann jederzeit
geändert oder entfernt werden, ohne den ausgeführten Befehl zu verändern.

!!! example "Ein Symbol ändert nur die Beschriftung"

    ```yaml title="Dieselbe Schaltfläche mit und ohne Symbol"
    - name: Disk usage
      type: button
      function: command
      command: "df -h"

    - name: Disk usage
      type: button
      icon: "💾"
      function: command
      command: "df -h"
    ```

## :material-help-circle-outline: Vor riskanten Aktionen nachfragen

Fügen Sie `confirm: true` hinzu. Der Bot fragt dann zunächst mit Yes oder
Cancel nach. Verwenden Sie dies für alles, was einen Dienst stoppt, Daten
löscht oder den Rechner neu startet. Die Aufforderung läuft nach einer Weile
ab (standardmäßig nach fünf Minuten).

Unter [Bestätigung](confirmation.md) erfahren Sie, wann eine Rückfrage sinnvoll
ist und wie Sie die Wartezeit ändern.

## Einstellungen nur für eine Schaltfläche

Die meisten globalen Einstellungen lassen sich für eine einzelne Schaltfläche
überschreiben. Das ist praktisch, wenn sich ein Auftrag anders verhält:

!!! example "Einstellungen für einen langsamen Auftrag überschreiben"

    ```yaml title="Ein langsamer Auftrag in einem anderen Verzeichnis"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"
      workdir: "/var/backups"
      env:
        BACKUP_MODE: "full"
    ```

`timeout` gibt diesem Befehl mehr Zeit, `workdir` wählt sein
Arbeitsverzeichnis und `env` ergänzt Umgebungsvariablen nur für ihn.

## Werte für die Funktion

Schreiben Sie Funktionswerte direkt auf die Schaltfläche. `command`, `path` und
`args` sind Kurzfelder für gleichnamige Parameter. Eigene Namen wie `url`,
`host`, `unit` und `lines` funktionieren genauso.

!!! example "Eigene Werte übergeben"

    ```yaml title="Aktuelle nginx-Protokolle"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

Verschachteln Sie diese Werte nicht unter `params:`. Jeder Schlüssel muss einem
von der ausgewählten Funktion deklarierten Parameter entsprechen.
[`validate`](../cli.md#validate) meldet unbekannte Namen, fehlende erforderliche
Werte sowie ungültige `int`- oder `bool`-Werte.

## Konfiguration

Alle von einer Schaltfläche akzeptierten Felder finden Sie unter
[Konfiguration → Menü](../configuration.md#menu).

## Verwandte Themen

<div class="grid cards cols-2" markdown>

-   :material-folder-outline:{ .middle } __Kategorie__

    ---

    Öffnet ein Untermenü, statt etwas auszuführen.

    [:octicons-arrow-right-24: Kategorie](category.md)

-   :material-function:{ .middle } __Funktion__

    ---

    Was beim Antippen einer Schaltfläche tatsächlich ausgeführt wird.

    [:octicons-arrow-right-24: Funktion](function.md)

-   :material-tune:{ .middle } __Parameter__

    ---

    Werte, die eine Funktion von der Schaltfläche benötigt.

    [:octicons-arrow-right-24: Parameter](parameter.md)

-   :material-view-list:{ .middle } __Menü__

    ---

    Den gesamten Baum erstellen und organisieren.

    [:octicons-arrow-right-24: Menü](menu.md)

</div>
