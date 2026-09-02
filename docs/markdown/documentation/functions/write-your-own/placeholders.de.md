---
title: Platzhalter
description: So schreiben Sie den run-Befehl einer Funktion mit Platzhaltern, einschließlich optionaler Teile, die nur erscheinen, wenn eine Schaltfläche einen Wert bereitstellt.
icon: material/code-braces
---

# :material-code-braces: Platzhalter

Das Feld `run` ist eine kleine Vorlage. Platzhalter in doppelten geschweiften
Klammern werden vor der Befehlsausführung durch Parameterwerte ersetzt.

Zwei Muster decken fast alle Fälle ab:

- `{{.name}}` setzt den Wert des Parameters namens `name` ein.
- `{{if .name}} ... {{end}}` fügt den mittleren Teil nur ein, wenn `name`
  einen Wert besitzt.

## Einen Wert einsetzen

!!! example "Ein Wert in der Mitte eines Befehls"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}"
    params:
      - name: path
        required: true
        description: Pfad zur Protokolldatei
    ```

Eine Schaltfläche mit `path: "/var/log/app.log"` führt
`tail -n 200 /var/log/app.log` aus.

## Optionale Teile

Bei einem leeren Wert ist `{{if .name}}` falsch, sodass der gesamte Block
entfällt. So fügen Sie ein Flag nur bei Bedarf hinzu.

!!! example "Einen optionalen Teil mit einem Platzhalter erstellen"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}{{if .args}} | grep -- {{.args}}{{end}}"
    params:
      - name: path
        required: true
        description: Pfad zur Protokolldatei
      - name: args
        description: Optionaler Text zum Filtern
    ```

Zwei Schaltflächen verwenden die Funktion, eine mit und eine ohne Filter:

!!! example "Werte von einer Schaltfläche bereitstellen"

    ```yaml title="Schaltflächen für das Anwendungsprotokoll"
    - name: Anwendungsprotokoll
      type: button
      function: tail-log
      path: "/var/log/app.log"

    - name: Anwendungsfehler
      type: button
      function: tail-log
      path: "/var/log/app.log"
      args: "ERROR"
    ```

**Anwendungsprotokoll** führt `tail -n 200 /var/log/app.log` aus,
**Anwendungsfehler** führt
`tail -n 200 /var/log/app.log | grep -- ERROR` aus.

Mit `{{if .args}} ... {{else}} ... {{end}}` können Sie auch zwischen zwei
Formen wählen.

## Verwendbare Namen

Ein Platzhalter kann jeden unter `params` deklarierten Parameter verwenden.
Sein Wert ist entweder der von der Schaltfläche bereitgestellte Wert oder der
`default` des Parameters.

!!! warning "Ein nicht deklarierter Name schlägt zur Laufzeit fehl"

    [`validate`](../../cli.md#validate) prüft nur die Platzhaltersyntax, nicht
    die Namen. Wenn `run` `{{.uri}}` enthält, aber kein Parameter `uri` heißt,
    wird die Konfiguration erfolgreich validiert. Beim Antippen meldet die
    Schaltfläche jedoch einen Fehler im Chat.

!!! warning "Werte werden als Klartext eingesetzt"

    Es werden keine Anführungszeichen oder Escape-Zeichen für Sie ergänzt. Ein
    Wert mit Leerzeichen oder Shell-Zeichen wird unverändert Teil des Befehls.
    Bevorzugen Sie daher feste Werte auf Schaltflächen und fügen Sie bei
    destruktiven Aktionen
    [`confirm: true`](../../concepts/confirmation.md) hinzu.

## Verwandte Themen

- [Dateistruktur](file-structure.md) — wo `run` und `params` stehen
- [Regeln](rules.md) — was der Loader akzeptiert
- [Schritt-für-Schritt-Anleitung](step-by-step.md) — ein vollständiges Beispiel ausprobieren
