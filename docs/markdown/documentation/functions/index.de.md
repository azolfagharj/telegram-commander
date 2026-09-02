---
title: Was ist eine Funktion?
description: Eine Funktion wandelt die auf einer Schaltfläche angegebenen Werte in einen Shell-Befehl um. Erfahren Sie, wie eine Schaltfläche die von einer Funktion deklarierten Parameter bereitstellt.
---

# :material-function-variant: Was ist eine Funktion?

Eine **Funktion** ist ein Rezept, das benannte Werte (ihre
[Parameter](../concepts/parameter.md)) in einen Shell-Befehl umwandelt. Jede
[Schaltfläche](../concepts/button.md) in Ihrem Menü nennt in ihrem Feld
`function` genau eine Funktion.

Stellen Sie sich eine Funktion als Lückentext für einen Befehl vor. Eine
Funktion zur Anzeige der Datenträgerbelegung besitzt eine Lücke — den Pfad —
die Sie auf jeder Schaltfläche ausfüllen.

## Was beim Antippen einer Schaltfläche geschieht

1. Der Bot sucht die Funktion, die im Feld `function` der Schaltfläche genannt ist.
2. Er sammelt die auf dieser Schaltfläche angegebenen Werte.
3. Er erstellt aus diesen Werten einen Shell-Befehl.
4. Er führt den Befehl in der [Shell](../concepts/shell.md) aus und sendet die
   Ausgabe als Codeblock an den Chat zurück.

Wenn die Funktion nicht existiert oder ein benötigter Wert fehlt, startet der
Bot gar nicht erst: [`validate`](../cli.md#validate) meldet das Problem zuvor.

## Ein vollständiges Beispiel

Die Funktion `command` ist integriert und immer verfügbar. Sie führt den Inhalt
des Felds `command` einer Schaltfläche aus.

!!! example "Die integrierte Funktion command verwenden"

    ```yaml title="Laufzeit-Schaltfläche"
    - name: Laufzeit
      type: button
      function: command
      command: "uptime"
    ```

Tippen Sie **Laufzeit** an. Der Bot führt `uptime` auf dem Server aus und sendet
die Ausgabe zurück.

## Werte von einer Schaltfläche übergeben { #passing-values-from-a-button }

Schreiben Sie jeden Wert direkt auf die Schaltfläche. Dafür gibt es zwei
Möglichkeiten:

1. Verwenden Sie die Kurzfelder `command`, `path` und `args`. Jedes davon füllt
   den gleichnamigen Parameter.
2. Verwenden Sie für jeden anderen Parameter dessen Namen als Schlüssel auf der
   Schaltfläche.

!!! example "Eine URL über ihren Parameternamen übergeben"

    ```yaml title="Schaltfläche zum Prüfen der API"
    - name: API prüfen
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

Hier entspricht `url` dem von `curl-url` deklarierten Parameter `url`. Dieselbe
Regel gilt für Namen wie `host`, `unit` und `lines`.

!!! warning "Werte nicht in `params:` verschachteln"

    `params:` gehört in eine eigene Funktionsdatei und deklariert dort die von
    der Funktion akzeptierten Werte. Schreiben Sie auf einer Schaltfläche jeden
    Wert direkt:

    ```yaml title="Werte gehören direkt auf die Schaltfläche"
    - name: nginx-Protokolle
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

    Numerische YAML-Werte benötigen keine Anführungszeichen.

[`validate`](../cli.md#validate) prüft jeden Schlüssel gegen die Parameter der
ausgewählten Funktion. Ein Tippfehler oder nicht deklarierter Name lässt die
Validierung fehlschlagen. Auch als `int` oder `bool` deklarierte Werte werden
geprüft. Für ausgelassene optionale Werte werden die Standardwerte verwendet.

## Zwei Arten von Funktionen { #two-kinds-of-function }

Funktionen sind entweder im Programm enthalten oder stammen aus einer
YAML-Datei auf Ihrem Server. Sobald eine Schaltfläche sie verwendet, verhalten
sich beide Arten gleich.

| | Integriert | Eigen |
|---|---|---|
| Herkunft | Im Programm enthalten | Eine von Ihnen geschriebene YAML-Datei |
| Müssen Sie eine Datei erstellen? | Nein | Ja, eine Datei pro Funktion |
| Namen | Reserviert (`command`, `script`) | Jeder nicht reservierte Name |
| Immer verfügbar? | Ja | Nur, wenn Sie `function_directory` festlegen |
| Bearbeitbar? | Nein | Ja, es sind Ihre Dateien |

Sie können beide Arten im selben Menü frei kombinieren. Die meisten Menüs
beginnen ausschließlich mit `command`-Schaltflächen. Wechseln Sie zu eigenen
Funktionen, wenn Sie denselben Befehl mit kleinen Abweichungen wiederholen.

### Integrierte Funktionen { #built-in-functions }

Zwei Funktionen werden immer geladen, auch wenn Sie kein `function_directory`
festlegen. Ihre Felder `command`, `path` und `args` stehen direkt auf einer
Schaltfläche.

| Funktion | Aufgabe | Erforderlich | Optional |
|----------|--------------|----------|----------|
| [`command`](built-in/command.md) | Führt einen Shell-Befehl unverändert aus | `command` | — |
| [`script`](built-in/script.md) | Führt eine Skriptdatei mit Argumenten aus | `path` | `args` |

Beide Namen sind **reserviert**. Eine eigene Funktionsdatei darf sie nicht
wiederverwenden. Der Loader stoppt mit einem Fehler wie
`function name "command" is reserved`, und der Bot startet nicht.

### Eigene Funktionen { #custom-functions }

Eine eigene Funktion ist eine einzelne YAML-Datei, die einen
wiederverwendbaren Befehl beschreibt. Bewahren Sie diese Dateien in einem
eigenen Ordner auf und lassen Sie
[`function_directory`](../configuration.md#function_directory-rules) darauf
verweisen.

!!! example "Dem Bot den Speicherort Ihrer Funktionsdateien mitteilen"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

Der Bot liest diesen Ordner einschließlich seiner Unterordner beim Start und
lädt jede `.yaml`- und `.yml`-Datei. Andere Dateien werden ignoriert.

Das Release-Archiv enthält bereits einen Ordner `functions/` mit fünf
Beispielen, die Sie unverändert verwenden können:

| Funktion | Aufgabe | Schaltflächenwerte |
|----------|--------------|---------------|
| [Echo Script](release-pack/echo-script.md) | Führt ein Skript über Bash aus | `path`, optional `args` |
| [Disk path](release-pack/disk-path.md) | Zeigt die Datenträgerbelegung | optional `path` |
| [Curl URL](release-pack/curl-url.md) | Ruft eine URL ab | `url` |
| [Ping Host](release-pack/ping-host.md) | Pingt einen Host | `host`, optional `count` |
| [Journal Unit](release-pack/journal-unit.md) | Zeigt aktuelle Dienstprotokolle | `unit`, optional `lines` |

Beginnen Sie zum Schreiben einer eigenen Funktion mit der
[Dateistruktur](write-your-own/file-structure.md) oder folgen Sie der
[Schritt-für-Schritt-Anleitung](write-your-own/step-by-step.md).

!!! tip "Geladene Funktionen prüfen"

    ```bash title="Alle für den Bot sichtbaren Funktionen auflisten"
    ./telegram-commander list-functions --config config.yaml
    ```

    Integrierte Funktionen zeigen `source=builtin`; eigene Funktionen zeigen
    die Datei, aus der sie stammen.

## Sicherheitshinweise

!!! warning "Schaltflächen werden mit den Rechten des Bots ausgeführt"

    Befehle werden mit den Rechten des Kontos ausgeführt, unter dem der Bot
    läuft. Ist dies root (wie bei der standardmäßigen
    [Dienst](../installation/run-as-a-service.md)-Einrichtung), können
    Schaltflächen alles auf dem Host ausführen. Fügen Sie nur
    [zugelassene Benutzer](../configuration.md#telegram) hinzu, denen Sie
    vertrauen.

    Parameterwerte werden als Text in den Befehl eingesetzt. Behandeln Sie sie
    wie Shell-Eingaben: Verwenden Sie vorzugsweise feste Schaltflächenwerte und
    ergänzen Sie [`confirm: true`](../concepts/confirmation.md) für alle
    destruktiven Aktionen.

!!! info "Lange Ausgaben werden gekürzt und aufgeteilt"

    Befehle enden nach ihrem `timeout`, und der Bot behält höchstens
    `max_output_bytes` ihrer Ausgabe. Alles, was länger als eine
    Telegram-Nachricht ist, wird als mehrere Nachrichten gesendet. Siehe
    [Konfiguration → Umfang der angezeigten Befehlsausgabe](../configuration.md#how-much-command-output-you-see).

## Verwandte Themen

- [`command`](built-in/command.md) — einen Shell-Befehl ausführen
- [`script`](built-in/script.md) — eine Skriptdatei ausführen
- [Schritt-für-Schritt-Anleitung](write-your-own/step-by-step.md) — Ihre erste eigene Funktion erstellen
- [Menü](../concepts/menu.md) — wie Schaltflächen auf Funktionen verweisen
- [Parameter](../concepts/parameter.md) — benannte Werte, die eine Funktion benötigt
