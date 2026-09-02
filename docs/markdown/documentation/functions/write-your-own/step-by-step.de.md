---
title: Schritt-für-Schritt-Anleitung
description: Erstellen Sie Ihre erste Telegram-Commander-Funktion von Grund auf, fügen Sie eine Schaltfläche hinzu, validieren Sie die Konfiguration und starten Sie den Bot, um das Ergebnis im Chat zu sehen.
icon: material/format-list-numbered
---

# :material-format-list-numbered: Schritt-für-Schritt-Anleitung

Diese Anleitung führt Sie von einem leeren Ordner zu einer kleinen,
funktionierenden Schaltfläche. Sie dauert nur wenige Minuten und benötigt
lediglich Ihre Konfigurationsdatei.

## 1. Einen Ordner in der Konfiguration angeben

!!! example "Den Ordner zur Konfiguration hinzufügen"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

Der Pfad ist relativ zur Konfigurationsdatei. Siehe
[Konfiguration → Regeln für `function_directory`](../../configuration.md#function_directory-rules).

## 2. Die Funktionsdatei schreiben

Erstellen Sie `functions/greet.yaml`. Der Parameter heißt `args`, damit eine
Schaltfläche ihn ausfüllen kann:

!!! example "Ihre erste Funktion"

    ```yaml title="functions/greet.yaml"
    name: greet
    run: "echo Hello {{.args}}"
    params:
      - name: args
        required: true
        description: Zu begrüßender Name
    ```

## 3. Eine Schaltfläche für die Funktion hinzufügen

!!! example "Eine Schaltfläche für die neue Funktion"

    ```yaml title="Schaltfläche zum Begrüßen"
    - name: Hallo sagen
      type: button
      function: greet
      args: "world"
    ```

## 4. Validieren und Funktionen auflisten

!!! example "Die Konfiguration prüfen und anschließend die geladenen Funktionen anzeigen"

    ```bash title="Validieren und anschließend die Funktionen auflisten"
    ./telegram-commander validate --config config.yaml
    ./telegram-commander list-functions --config config.yaml
    ```

Es sollte eine Zeile für `greet` erscheinen. Jede Zeile zeigt den Namen, die
Herkunft der Funktion und die Anzahl ihrer Parameter. Siehe
[CLI → list-functions](../../cli.md#list-functions).

Wenn `validate` einen Fehler meldet, nennt die Meldung die Schaltfläche und den
fehlenden Wert — prüfen Sie die [Regeln](rules.md).

## 5. Den Bot starten und die Schaltfläche antippen

!!! example "Den Bot im Vordergrund starten"

    ```bash title="Starten und die Ausgabe beobachten"
    ./telegram-commander run --config config.yaml
    ```

Senden Sie `/start` in Telegram und tippen Sie **Hallo sagen** an. Der Bot
führt `echo Hello world` aus und sendet `Hello world` als Codeblock zurück.

## 6. Wiederverwendbar machen

Ändern Sie den `args`-Wert der Schaltfläche oder fügen Sie eine zweite
Schaltfläche mit einem anderen Wert hinzu. Die Funktion bleibt unverändert:

!!! example "Zwei Schaltflächen, eine Funktion"

    ```yaml title="Begrüßungsschaltflächen"
    - name: Welt begrüßen
      type: button
      function: greet
      args: "world"
    - name: Team begrüßen
      type: button
      function: greet
      args: "team"
    ```

## Verwandte Themen

- [Platzhalter](placeholders.md) — optionale Teile zum Befehl hinzufügen
- [Regeln](rules.md) — was der Loader akzeptiert
- [Dateistruktur](file-structure.md) — Erklärung aller Felder
- [Menü](../../concepts/menu.md) — wo die Schaltfläche im Menü eingefügt wird
