---
title: Menü
description: Erstellen Sie das Menü, das Ihr Bot in Telegram anzeigt. Kombinieren Sie Schaltflächen und Kategorien zu einem Baum und erfahren Sie, wie Home und Back die Navigation ermöglichen.
icon: material/view-list
---

# :material-view-list: Menü

Ihr Menü ist ein Knotenbaum unter dem obersten Schlüssel `menu`. Es gibt zwei
Arten von Knoten:

- Eine **[Kategorie](category.md)** öffnet ein Untermenü. Sie besitzt `items`.
- Eine **[Schaltfläche](button.md)** führt etwas aus. Sie besitzt eine [Funktion](function.md).

Wenn diese Begriffe neu für Sie sind, lesen Sie zuerst die
[Grundlagen](button.md). Die vollständige Liste aller Felder finden Sie unter
[Konfiguration → Menü](../configuration.md#menu).

## :material-format-list-bulleted: Ein flaches Menü

Das einfachste Menü ist eine Liste von Schaltflächen ohne Verschachtelung:

!!! example "Ein Menü mit drei Schaltflächen erstellen"

    ```yaml title="Drei Schaltflächen ohne Kategorien"
    menu:
      - name: Laufzeit
        type: button
        function: command
        command: "uptime"

      - name: Freier Speicher
        type: button
        function: command
        command: "free -h"

      - name: Datenträgerbelegung
        type: button
        function: command
        command: "df -h"
    ```

Senden Sie `/start` in Telegram. Ihr Menü wird angezeigt.

## :material-folder-outline: Gruppierung mit Kategorien

Wenn Ihr Menü wächst, gruppieren Sie zusammengehörige Aktionen in Kategorien.
Beim Antippen zeigt eine Kategorie ihre `items`. Home ist immer im Menü
vorhanden. Back erscheint nur innerhalb einer Kategorie.

!!! example "Schaltflächen in Kategorien gruppieren"

    ```yaml title="Kategorien System und Dienste"
    menu:
      - name: System
        type: category
        icon: "💻"
        items:
          - name: Laufzeit
            type: button
            function: command
            command: "uptime"

          - name: Freier Speicher
            type: button
            function: command
            command: "free -h"

      - name: Dienste
        type: category
        icon: "🔧"
        items:
          - name: nginx neu starten
            type: button
            function: command
            command: "systemctl restart nginx"
            confirm: true
    ```

Kategorien können beliebig tief verschachtelt werden. Eine Kategorie muss
mindestens einen Eintrag enthalten.

## :material-family-tree: Namen müssen unter Geschwistern eindeutig sein

Zwei Knoten unter demselben übergeordneten Knoten dürfen nicht denselben Namen
besitzen (Groß-/Kleinschreibung wird ignoriert). Folgendes ist zulässig, da die
Knoten in verschiedenen Kategorien liegen:

!!! example "Einen Namen unter verschiedenen übergeordneten Knoten wiederverwenden"

    ```yaml title="Derselbe Name unter zwei übergeordneten Knoten"
    menu:
      - name: Web
        type: category
        items:
          - name: Neustart          # ok
            type: button
            function: command
            command: "systemctl restart nginx"
      - name: Datenbank
        type: category
        items:
          - name: Neustart          # ok, different parent
            type: button
            function: command
            command: "systemctl restart postgresql"
    ```

## :material-emoticon-outline: Symbole

`icon` ist ein optionales Emoji vor dem Namen. Es dient nur der Darstellung.

!!! warning "Diese Schaltfläche startet den Rechner neu"

    ```yaml title="Eine Schaltfläche mit Emoji-Symbol"
    - name: Neustart
      type: button
      icon: "🔁"
      function: command
      command: "reboot"
      confirm: true
    ```

!!! tip "Wählen Sie ein einfaches, gebräuchliches Emoji"

    Einige Emojis führen auf manchen Smartphones dazu, dass der
    Schaltflächentext abgeschnitten wird oder über die Schaltfläche hinausragt.
    Probieren Sie in diesem Fall ein anderes Emoji aus.

## :material-cellphone-text: Darstellung des Telegram-Menüs { #how-the-telegram-menu-looks }

Alle Schaltflächen erscheinen auf der Tastatur unter dem Nachrichtenfeld (der
Tastatur, die sich über die kleine Schaltfläche am rechten Ende des
Nachrichtenfelds ein- und ausblenden lässt). Diese Tastatur nimmt immer die
volle Breite des Chats ein, sodass Schaltflächentext nicht zusammengedrückt
oder abgeschnitten wird.

- **Home** ist auf jedem Bildschirm immer die erste Schaltfläche. Tippen Sie sie an, um zum ersten Bildschirm zurückzukehren.
- **Back** erscheint innerhalb einer Kategorie.
- **$ >_ Run Command** erscheint, wenn `enable_run_command` aktiviert ist (siehe unten).
- Einträge stehen standardmäßig zu zweit in einer Zeile. Eine Kategorie kann dies mit `columns` ändern. Bei vielen Einträgen ermöglichen **Prev** und **Next** die Seitennavigation.
- Schaltflächen mit `confirm: true` fragen vor der Ausführung mit Yes / Cancel nach.

!!! info "Menütitel werden wiederverwendet, Ausgaben bleiben erhalten"

    Ein neuer Menütitel (Home, eine Kategorie oder eine Seite) ersetzt den
    vorherigen, damit sich der Chat nicht mit leeren Bildschirmen füllt. Die
    Zeile **Running** und die Befehlsausgabe bleiben im Chat, sodass Sie auch
    nach erneutem Öffnen des Menüs sehen können, was ausgeführt wurde.

!!! info "Lange Ausgaben werden in mehreren Nachrichten gesendet"

    Befehlsausgaben werden als Codeblock dargestellt. Ist die Ausgabe länger
    als eine Telegram-Nachricht, wird sie in mehreren Nachrichten gesendet, von
    denen jede auf die vorherige antwortet. Der letzte Teil behält die
    Schaltflächen der aktuellen Seite, sodass **Back** weiterhin die Kategorie
    verlässt. Siehe
    [Konfiguration → Umfang der angezeigten Befehlsausgabe](../configuration.md#how-much-command-output-you-see).

## :material-help-circle-outline: Bestätigung { #confirmation }

Fügen Sie einer Schaltfläche `confirm: true` hinzu, um vor der Ausführung ein
zweites Antippen („Are you sure?“) zu verlangen. Verwenden Sie dies für
destruktive Aktionen. Siehe [Bestätigung](confirmation.md).

!!! warning "Diese Schaltfläche stoppt einen Dienst"

    ```yaml title="Eine Schaltfläche, die zuerst nachfragt"
    - name: nginx stoppen
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

Die Bestätigungsaufforderung läuft nach einer Weile ab (standardmäßig nach
5 Minuten). Ändern Sie dies mit `confirm_ttl`; siehe
[Konfiguration → Felder auf oberster Ebene](../configuration.md#root-fields).

## :material-tune-variant: Schaltflächenspezifische Einstellungen

Einige globale Einstellungen lassen sich für eine einzelne Schaltfläche
überschreiben:

!!! example "Einer Schaltfläche eigene Einstellungen geben"

    ```yaml title="Eine Schaltfläche mit eigener Zeitüberschreitung, eigenem Ordner und eigenen Variablen"
    - name: Lange Sicherung
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"          # this one may take longer than the global timeout
      workdir: "/var/backups" # run it here
      env:
        BACKUP_MODE: "full"   # extra environment variable for this command
    ```

Die vollständige Feldliste finden Sie unter
[Konfiguration → Menü](../configuration.md#menu).

## :material-view-grid-outline: Layout steuern

`menu_columns` legt die Anzahl der **Eintrags**-Schaltflächen pro Zeile fest
(Standardwert 2). Eine Kategorie kann dies mit `columns` überschreiben. Besitzt
ein Menü mehr als `page_size` Einträge (Standardwert 8), wird es in Seiten
aufgeteilt. Prev/Next werden angezeigt, bis Sie das jeweilige Ende erreichen.
Siehe [Konfiguration → Felder auf oberster Ebene](../configuration.md#root-fields).

## :material-console: Run Command

Wenn Sie auf der obersten Ebene Ihrer Konfiguration
`enable_run_command: true` festlegen, bleibt eine Schaltfläche
**$ >_ Run Command** im Menü (innerhalb einer Kategorie hinter Back, auf dem
ersten Bildschirm hinter Home). Tippen Sie sie an und senden Sie anschließend
den auszuführenden Shell-Befehl. Der Bot verwendet dieselbe Shell, dieselbe
Zeitüberschreitung, dasselbe Arbeitsverzeichnis und dieselben Ausgabelimits wie
Ihre anderen Schaltflächen.

Home oder Back bricht die Aufforderung ab, ohne etwas auszuführen. Diese
Funktion ist standardmäßig deaktiviert. Siehe
[Konfiguration → Felder auf oberster Ebene](../configuration.md#root-fields).

!!! warning "Dies gewährt Zugriff auf den gesamten Rechner"

    Wenn Run Command aktiviert ist, kann jeder zugelassene Benutzer jeden
    Befehl auf dem Host ausführen, nicht nur Ihre definierten Schaltflächen.
    Aktivieren Sie die Funktion nur, wenn Sie allen zugelassenen Benutzern
    entsprechend vertrauen.

## :material-timer-sand: Ein Befehl nach dem anderen

Ihre Befehle werden nacheinander statt gleichzeitig ausgeführt.

!!! info "Ein zweites Antippen wartet, bis es an der Reihe ist"

    Tippen Sie eine zweite Schaltfläche an, während die erste noch arbeitet,
    wartet die zweite und wird danach separat ausgeführt. Ihre Zeile
    **Running** erscheint, sobald sie startet. So geraten zwei Ihrer Aktionen
    nicht beim Zugriff auf denselben Dienst oder dieselbe Datei in Konflikt.
    Andere Personen werden durch Ihren Befehl nicht aufgehalten; jede Person
    besitzt ihre eigene Warteschlange.

## :material-function-variant: Was beim Antippen ausgeführt wird

Jede Schaltfläche verweist über ihr Feld `function` auf eine **Funktion**. Die
Schaltflächen in den obigen Beispielen verwenden die integrierte Funktion
`command`. Informationen zu integrierten und eigenen Funktionen sowie zum
Hinzufügen eigener Funktionen finden Sie unter
[Funktionen](../functions/index.md).

## :material-link-variant: Verwandte Seiten

- [Schaltfläche](button.md) — was eine Schaltfläche ist
- [Kategorie](category.md) — Knoten für Untermenüs
- [Konfiguration → Menü](../configuration.md#menu) — alle Felder
