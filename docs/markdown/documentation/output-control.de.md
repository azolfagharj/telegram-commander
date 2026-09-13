---
title: Ausgabe steuern
description: Entscheiden Sie, wie Befehlsergebnisse in Ihren Chat kommen. Setzen Sie output, max_output_messages und max_output_bytes auf Root-Ebene, überschreiben Sie output an einer einzelnen Schaltfläche und sehen Sie jede Kombination aus Root und Schaltfläche mit Beispielen.
icon: material/export-variant
---

# :material-export-variant: Ausgabe steuern

Wenn eine [Schaltfläche](concepts/button.md) fertig ist, muss der Bot das
Ergebnis in Ihren Chat bringen. Telegram nimmt niemals mehr als 4096 Bytes in
einer Nachricht an. Ein langes Ergebnis kommt deshalb entweder als **mehrere
Antwortnachrichten** oder als **eine `.txt`-Datei**, die Sie öffnen und
herunterladen können.

Drei optionale Felder entscheiden, was geschieht. Sie können alle drei
weglassen, und der Bot verhält sich trotzdem sinnvoll.

!!! info "Die eine Regel, die Sie sich merken sollten"

    Die Schaltfläche gewinnt. Setzt eine Schaltfläche `output`, dann gilt das.
    Sagt die Schaltfläche nichts, gilt der Root-Wert. Sagt auch Root nichts,
    verwendet der Bot `auto`.

<div class="grid cards cols-3" markdown>

-   :material-scissors-cutting:{ .middle } __`max_output_bytes`__

    ---

    Wie viel Ausgabe vom Befehl behalten wird, bevor überhaupt etwas gesendet wird.

    [:octicons-arrow-right-24: Mehr lesen](#max_output_bytes)

-   :material-tune:{ .middle } __`output`__

    ---

    Nachrichten oder Datei. Das einzige Feld, das Sie auch an einer Schaltfläche setzen können.

    [:octicons-arrow-right-24: Mehr lesen](#output)

-   :material-counter:{ .middle } __`max_output_messages`__

    ---

    Im Modus `auto`: wie viele Nachrichten erlaubt sind, bevor eine Datei gesendet wird.

    [:octicons-arrow-right-24: Mehr lesen](#max_output_messages)

</div>

!!! tip "Jedes Bild auf dieser Seite lässt sich öffnen"

    Die Chatbilder sind absichtlich klein. Klicken Sie eines an, um es in
    voller Größe zu sehen.

## :material-card-bulleted-outline: Die drei Felder im Überblick

| Feld | Wo Sie es setzen | Typ | Standard | Was es entscheidet |
|------|------------------|-----|----------|--------------------|
| `max_output_bytes` | nur Root | Ganzzahl | `524288` | Wie viele Bytes Ausgabe der Bot pro Befehl behält |
| `output` | Root **und** Schaltfläche | `auto` \| `text` \| `file` | `auto` | Ob das Ergebnis als Chatnachrichten oder als eine `.txt`-Datei ankommt |
| `max_output_messages` | nur Root | Ganzzahl, `1`–`10` | `2` | Im Modus `auto`: die Anzahl der erlaubten Nachrichten, bevor stattdessen eine Datei gesendet wird |

Alle drei stehen in Ihrer [Konfigurationsdatei](concepts/config-file.md) auf
der **obersten Ebene**, neben `shell` und `timeout`, nicht unter `telegram`.
Die vollständige Liste der Schlüssel auf oberster Ebene finden Sie auf der
Seite [Konfiguration](configuration.md#root-fields).

## :material-scissors-cutting: `max_output_bytes`

Dies ist Ihr eigenes Limit, und es kommt zuerst. Während ein Befehl läuft,
behält der Bot höchstens so viele Bytes seiner Ausgabe, getrennt gezählt für
normale Ausgabe und Fehlerausgabe. Alles darüber wird verworfen, der Befehl
selbst läuft aber weiter, bis er endet oder sein `timeout` erreicht.

!!! example "Bis zu 2 MB Ausgabe pro Befehl behalten"

    ```yaml title="config.yaml"
    max_output_bytes: 2097152 # (1)!
    ```

    1.  2 MB statt der standardmäßigen 512 KB. Wird für normale Ausgabe und
        für Fehlerausgabe getrennt gezählt.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Wurde das Limit erreicht, steht das direkt unter der Zusammenfassung in
    einer eigenen Zeile. Alles über `(output truncated)` ist der Teil, den der
    Bot behalten hat; der Rest des Protokolls hat den Server nie verlassen.

    </div>
    <div class="result-shot" markdown>

    ![Ein Telegram-Ergebnis mit dem Hinweis output truncated](/images/output-control/truncated.png){ .shot loading=lazy }

    </div>
    </div>

!!! tip "Ein höherer Wert hilft nur, wenn eine Datei gesendet wird"

    Eine Chatnachricht kann niemals mehr als 4096 Bytes tragen. Ein größeres
    `max_output_bytes` erreicht Sie deshalb nur vollständig, wenn das Ergebnis
    als `.txt`-Datei ankommt. Kombinieren Sie es mit `output: auto` (dem
    Standard) oder mit `output: file`.

## :material-tune: `output`

`output` wählt die Zustellung. Es nimmt drei Werte an und ist das einzige
Ausgabefeld, das Sie auch an eine einzelne Schaltfläche schreiben können.

=== "auto"

    Textnachrichten, solange das Ergebnis kurz ist, und eine `.txt`-Datei,
    wenn nicht. Das ist der Standard und der Wert, der gilt, wenn Sie nichts
    schreiben.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Den Bot für jedes Ergebnis entscheiden lassen"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: auto # (1)!
        max_output_messages: 2 # (2)!
        ```

        1.  Pro Ergebnis entscheiden, statt eine Zustellung zu erzwingen.
        2.  Der Umschaltpunkt. Zwei Nachrichten oder weniger bleiben im Chat;
            alles Längere wird eine Datei.

    </div>
    <div class="result-shot" markdown>

    ![Im Modus auto bleibt kurze Ausgabe Text und längere Ausgabe wird als Datei gesendet](/images/output-control/auto.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

=== "text"

    Immer Chatnachrichten, wie lang das Ergebnis auch ist. Der Bot stoppt nach
    zehn Nachrichten und ergänzt einen Hinweis, dass der Rest abgeschnitten
    wurde.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Alles im Chat behalten"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: text # (1)!
        ```

        1.  `max_output_messages` wird hier ignoriert, weil der Bot nie von
            selbst zu einer Datei wechselt.

    </div>
    <div class="result-shot" markdown>

    ![Ein Befehlsergebnis, das als Telegram-Nachrichten zugestellt wird](/images/output-control/message.png){ .shot loading=lazy }

    </div>
    </div>

=== "file"

    Immer eine `.txt`-Datei, selbst bei einem einzeiligen Ergebnis. Die
    Zusammenfassung kommt als Beschriftung der Datei mit.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Jedes Ergebnis als Anhang senden"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: file # (1)!
        ```

        1.  Jede Schaltfläche sendet nun einen Anhang, auch die schnellen.
            Setzen Sie dies stattdessen an einzelne Schaltflächen, wenn das zu
            viel ist.

    </div>
    <div class="result-shot" markdown>

    ![Ein Befehlsergebnis, das als Telegram-Datei zugestellt wird](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

!!! note "Weglassen ist dasselbe wie `auto` zu schreiben"

    Sie müssen `output` nie schreiben. Eine Konfiguration, die schon vor
    diesem Feld funktioniert hat, verhält sich genau wie vorher.

## :material-counter: `max_output_messages`

Dieses Feld ist nur im Modus `auto` relevant, und es gibt es nur auf oberster
Ebene. Es ist die Anzahl der Chatnachrichten, die der Bot senden will, bevor
er Nachrichten aufgibt und stattdessen eine Datei sendet. Erlaubt sind `1` bis
`10`; der Standardwert ist `2`.

!!! example "Eine Datei senden, sobald das Ergebnis eine zweite Nachricht braucht"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 1 # (1)!
    ```

    1.  Die strengste Einstellung. Nur ein Ergebnis, das in eine Nachricht
        passt, bleibt im Chat.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Ein Ergebnis, das in das erlaubte Kontingent passt, bleibt als
    Chatnachrichten. Sobald es eine Nachricht mehr bräuchte, kommt das ganze
    Ergebnis stattdessen als eine einzige `.txt`-Datei, und es werden keine
    Teilnachrichten gesendet.

    </div>
    <div class="result-shot" markdown>

    ![Ein kurzes Ergebnis als Nachricht neben einem längeren als Datei](/images/output-control/allowance.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

!!! warning "Das Feld gibt es nur auf oberster Ebene"

    Eine Schaltfläche kann `max_output_messages` nicht setzen. Setzen Sie es
    auf oberster Ebene und lassen Sie Schaltflächen, die etwas anderes
    brauchen, `output: text` oder `output: file` setzen.

## :material-table-arrow-right: Root und Schaltfläche zusammen

Eine Schaltfläche darf ihr eigenes `output` tragen. Sonst lässt sich nichts
zur Ausgabe pro Schaltfläche einstellen. Der wirksame Modus ergibt sich
einfach so:

!!! abstract "Wie der wirksame Modus gewählt wird"

    ```text title="Reihenfolge des Vorrangs"
    button output  →  root output  →  auto
    ```

| Root `output` | `output` der Schaltfläche | Was tatsächlich passiert |
|---------------|---------------------------|--------------------------|
| nicht gesetzt oder `auto` | nicht gesetzt | `auto` |
| nicht gesetzt oder `auto` | `auto` | `auto` |
| nicht gesetzt oder `auto` | `text` | `text` |
| nicht gesetzt oder `auto` | `file` | `file` |
| `text` | nicht gesetzt | `text` |
| `text` | `auto` | `auto` |
| `text` | `text` | `text` |
| `text` | `file` | `file` |
| `file` | nicht gesetzt | `file` |
| `file` | `auto` | `auto` |
| `file` | `text` | `text` |
| `file` | `file` | `file` |

Die folgenden Abschnitte gehen alle zwölf Kombinationen durch, mit je einem
Beispiel.

## :material-numeric-1-box-outline: Wenn Root `auto` ist

Dazu gehört auch der Fall, dass Root gar nichts sagt, denn ein fehlendes
`output` bedeutet `auto`.

### Root ist `auto`, die Schaltfläche sagt nichts

Der Alltagsfall. Kurze Ergebnisse bleiben im Chat; ein Ergebnis, das mehr als
`max_output_messages` Nachrichten bräuchte, kommt als Datei.

!!! example "Eine einfache Schaltfläche unter einem einfachen Root"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 2

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    `df -h` gibt einige Zeilen aus, passt also in eine Nachricht und bleibt
    als Codeblock im Chat, mit der Zusammenfassung darüber.

    </div>
    <div class="result-shot" markdown>

    ![Datenträgerbelegung als eine Telegram-Nachricht](/images/output-control/disk-usage.png){ .shot loading=lazy }

    </div>
    </div>

### Root ist `auto`, die Schaltfläche sagt `auto`

`auto` an der Schaltfläche zu schreiben ändert hier nichts. Es lohnt sich nur,
wenn die Schaltfläche in `auto` bleiben soll, egal was aus Root später wird.

!!! example "Den Standard an der Schaltfläche ausschreiben"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
        output: auto # (1)!
    ```

    1.  Heute dasselbe Verhalten wie ohne die Zeile, aber diese Schaltfläche
        behält es auch, wenn Root später auf `text` oder `file` wechselt.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Genau das automatische Verhalten: Text, solange das Ergebnis kurz ist, und
    eine Datei, sobald es über das Kontingent wächst. Auf einem Server mit
    vielen eingebundenen Datenträgern kann genau dieselbe Schaltfläche Ihnen
    stattdessen eine Datei senden.

    </div>
    <div class="result-shot" markdown>

    ![Dieselbe Schaltfläche als Nachricht und als Datei](/images/output-control/disk-usage-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### Root ist `auto`, die Schaltfläche sagt `text`

Verwenden Sie das für eine Schaltfläche, deren Ausgabe Sie immer direkt im Chat
lesen wollen, auch wenn sie etwas länger wird — eine Statusliste, durch die Sie
scrollen, statt sie herunterzuladen.

!!! example "Eine Schaltfläche im Chat behalten"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Service status
        type: button
        function: command
        command: "systemctl status nginx --no-pager"
        output: text # (1)!
    ```

    1.  Nie eine Datei, egal was `max_output_messages` sagt.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Das Ergebnis wird in mehrere Nachrichten geteilt, von denen jede auf die
    vorherige antwortet, sodass die Reihenfolge nie verloren geht. Die
    Schaltflächen bleiben am letzten Teil.

    </div>
    <div class="result-shot" markdown>

    ![Ein langer Status, auf drei Telegram-Nachrichten geteilt](/images/output-control/service-status-split.png){ .shot loading=lazy }

    </div>
    </div>

### Root ist `auto`, die Schaltfläche sagt `file`

Die typische Wahl für Protokolle und Sicherungen: Sie wollen sie fast nie im
Chat lesen, und eine Datei lässt sich leichter aufbewahren.

!!! example "Dieses eine Ergebnis immer herunterladen"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Nginx logs
        type: button
        function: command
        command: "journalctl -u nginx -n 2000 --no-pager"
        output: file # (1)!
    ```

    1.  Der Rest des Menüs verhält sich weiterhin automatisch. Nur diese
        Schaltfläche ist auf eine Datei festgelegt.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Ein Anhang mit dem Schaltflächennamen, dem Exit-Code und der Dauer als
    Beschriftung. Tippen Sie ihn an, um das ganze Protokoll zu lesen, oder
    behalten Sie ihn für später.

    </div>
    <div class="result-shot" markdown>

    ![Nginx-Protokolle, die als Telegram-Datei zugestellt werden](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-2-box-outline: Wenn Root `output: text` ist

Jede Schaltfläche bleibt im Chat, sofern sie nichts anderes sagt. Wählen Sie
das, wenn Sie keine Dateien auf Ihr Telefon herunterladen mögen.

### Root ist `text`, die Schaltfläche sagt nichts

!!! example "Überall Chatnachrichten"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Eine kurze Nachricht, weil `uptime` nur eine Zeile ausgibt. Längere
    Ergebnisse werden einfach auf mehr Nachrichten geteilt, bis zu zehn.

    </div>
    <div class="result-shot" markdown>

    ![Laufzeit als eine Telegram-Nachricht](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### Root ist `text`, die Schaltfläche sagt `auto`

Die Schaltfläche schaltet den automatischen Wechsel wieder ein, sodass genau
diese eine Schaltfläche eine Datei senden darf, obwohl das restliche Menü das
nie tut. `max_output_messages` von Root entscheidet, wo der Wechsel passiert.

!!! example "Eine Schaltfläche darf bei langer Ausgabe eine Datei senden"

    ```yaml title="config.yaml"
    output: text
    max_output_messages: 3 # (1)!

    menu:
      - name: Package list
        type: button
        function: command
        command: "dpkg -l"
        output: auto
    ```

    1.  Nur diese Schaltfläche liest das Kontingent, weil sie die einzige im
        Modus `auto` ist.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Eine kurze Paketliste kommt als bis zu drei Nachrichten. Auf einem voll
    belegten Server überschreitet dieselbe Schaltfläche das Kontingent und
    sendet stattdessen eine Datei.

    </div>
    <div class="result-shot" markdown>

    ![Eine Paketliste als Nachrichten und als Datei](/images/output-control/package-list-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### Root ist `text`, die Schaltfläche sagt `text`

Dasselbe wie nichts zu sagen, aber die Absicht steht geschrieben. Nützlich in
einer langen Konfiguration, in der eine spätere Änderung auf oberster Ebene
diese Schaltfläche nicht verändern soll.

!!! example "Eine Schaltfläche auf Chatnachrichten festlegen"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Last logins
        type: button
        function: command
        command: "last -n 20"
        output: text
    ```

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Immer Chatnachrichten, bis zu zehn davon, egal worauf Root später
    geändert wird.

    </div>
    <div class="result-shot" markdown>

    ![Letzte Anmeldungen als eine Telegram-Nachricht](/images/output-control/last-logins.png){ .shot loading=lazy }

    </div>
    </div>

### Root ist `text`, die Schaltfläche sagt `file`

!!! example "Eine herunterladbare Schaltfläche in einem reinen Chat-Menü"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Full system report
        type: button
        function: command
        command: "/usr/local/bin/report.sh"
        output: file
        timeout: "5m" # (1)!
    ```

    1.  Ein Bericht kann eine Weile dauern, deshalb erhält diese Schaltfläche
        mehr Zeit als das globale Zeitlimit.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Der einzige Anhang in einem ansonsten reinen Chat-Menü. Die Beschriftung
    zeigt, wie lange der Bericht gedauert hat. Das ist nützlich, wenn er
    minutenlang läuft.

    </div>
    <div class="result-shot" markdown>

    ![Ein vollständiger Systembericht, der als Telegram-Datei zugestellt wird](/images/output-control/system-report-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-3-box-outline: Wenn Root `output: file` ist

Jedes Ergebnis ist eine Datei, auch ein einzeiliges. Das passt zu einem Bot,
der vor allem für Berichte und Archive genutzt wird.

### Root ist `file`, die Schaltfläche sagt nichts

!!! example "Überall Dateien"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Selbst eine einzige Ausgabezeile kommt als Anhang. Die Beschriftung trägt
    weiterhin die Zusammenfassung, sodass Sie den Exit-Code lesen können, ohne
    die Datei zu öffnen.

    </div>
    <div class="result-shot" markdown>

    ![Laufzeit als kleine Telegram-Datei](/images/output-control/uptime-file.png){ .shot loading=lazy }

    </div>
    </div>

!!! warning "Auch winzige Ergebnisse werden Downloads"

    Eine einzeilige Antwort kommt als Anhang, den Sie öffnen müssen. Wenn Sie
    das stört, lassen Sie Root auf `auto` und setzen `output: file` nur an
    den Schaltflächen, die es brauchen.

### Root ist `file`, die Schaltfläche sagt `auto`

Die Schaltfläche entkommt der Nur-Datei-Regel und verhält sich wieder normal.

!!! example "Eine schnelle Prüfung lesbar lassen"

    ```yaml title="config.yaml"
    output: file
    max_output_messages: 2

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
        output: auto # (1)!
    ```

    1.  Zurück zum automatischen Wechsel, nur für diese Schaltfläche.

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Die einzeilige Antwort ist kurz, sie bleibt also als Nachricht im Chat,
    während jede andere Schaltfläche im Menü weiterhin Dateien sendet.

    </div>
    <div class="result-shot" markdown>

    ![Laufzeit wieder als Nachricht im Chat](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### Root ist `file`, die Schaltfläche sagt `text`

!!! example "Eine Schaltfläche zurück in den Chat zwingen"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Who is logged in
        type: button
        function: command
        command: "who"
        output: text
    ```

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Eine Nachricht, nie ein Anhang, wie lang die Liste der angemeldeten
    Benutzer auch wird. Über zehn Nachrichten hinaus vermerkt der Bot, dass
    der Rest abgeschnitten wurde.

    </div>
    <div class="result-shot" markdown>

    ![Angemeldete Benutzer als eine Telegram-Nachricht](/images/output-control/who.png){ .shot loading=lazy }

    </div>
    </div>

### Root ist `file`, die Schaltfläche sagt `file`

Der Root-Wert wird wiederholt. Das ist harmlos und hält die Schaltfläche
korrekt, wenn Sie Root später lockern.

!!! example "Eine Schaltfläche, die immer eine Datei sein muss"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Database dump log
        type: button
        function: command
        command: "cat /var/log/pg_dump.log"
        output: file
    ```

!!! success "Was Sie im Chat erhalten"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Ein Anhang, genau wie Root es ohnehin verlangt. An der Schaltfläche
    geschrieben, übersteht die Absicht eine Änderung auf oberster Ebene.

    </div>
    <div class="result-shot" markdown>

    ![Ein Datenbank-Dump-Protokoll, das als Telegram-Datei zugestellt wird](/images/output-control/db-dump-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-file-document-outline: Wie die Datei aussieht

Wird ein Ergebnis als Datei zugestellt, erhalten Sie einen Anhang und eine
kurze Beschriftung.

- **Name** — der Schaltflächenname in Kleinbuchstaben mit Bindestrichen, dann
  Datum und Uhrzeit in UTC, zum Beispiel `nginx-logs-20260913-091204.txt`.
- **Beschriftung** — dieselbe Zusammenfassung, die Sie oben in einer
  Chatnachricht sehen würden: Schaltflächenname, Exit-Code, Dauer sowie der
  Hinweis auf Zeitüberschreitung oder Kürzung, wenn es einen gibt.
- **Inhalt** — nochmals diese Zusammenfassung, dann ein Abschnitt
  `--- stdout ---` und ein Abschnitt `--- stderr ---`, wenn der Befehl dorthin
  etwas geschrieben hat.

!!! success "Inhalt der heruntergeladenen Datei"

    ```text title="nginx-logs-20260913-091204.txt"
    Button: Nginx logs
    Exit: 0
    Duration: 316ms

    --- stdout ---
    2026-09-13T09:11:58 nginx: worker process started
    ...

    --- stderr ---
    2026-09-13T09:12:01 nginx: warning: duplicate server name
    ```

## :material-alert-outline: Grenzen, die Sie nicht ändern können

!!! warning "Telegram setzt die Obergrenze, nicht der Bot"

    - Eine Chatnachricht trägt höchstens 4096 Bytes.
    - Im Modus `text` sendet der Bot höchstens zehn Nachrichten und endet mit
      `(output too long; showing first N bytes)`.
    - Welches `max_output_messages` Sie auch schreiben: zehn Nachrichten
      bleiben die harte Obergrenze.

!!! note "Scheitert der Upload, erhalten Sie das Ergebnis trotzdem"

    Kann die Datei nicht gesendet werden — kein Netz, Telegram lehnt sie ab —,
    fällt der Bot auf Chatnachrichten zurück, damit das Ergebnis nicht
    verloren geht.

## :material-console: Run Command nutzt die Root-Einstellung

Die Schaltfläche **`$ >_ Run Command`** ist nicht Teil Ihres Menüs und hat
deshalb kein eigenes `output`. Sie folgt immer `output` und
`max_output_messages` auf oberster Ebene. Siehe
[Menü → Run Command](concepts/menu.md#run-command).

## :material-link-variant: Verwandte Themen

<div class="grid cards cols-2" markdown>

-   :material-cog-outline:{ .middle } __Konfiguration__

    ---

    Alle Schlüssel auf oberster Ebene, mit den drei Ausgabefeldern in einer Tabelle.

    [:octicons-arrow-right-24: Konfiguration](configuration.md#root-fields)

-   :material-gesture-tap-button:{ .middle } __Schaltfläche__

    ---

    Was eine Schaltfläche überschreiben kann, `output` eingeschlossen.

    [:octicons-arrow-right-24: Schaltfläche](concepts/button.md)

-   :material-view-list:{ .middle } __Menü__

    ---

    Wie Ergebnisse neben dem Menü liegen und was Run Command tut.

    [:octicons-arrow-right-24: Menü](concepts/menu.md)

-   :material-timer-outline:{ .middle } __Shell__

    ---

    Shell, Arbeitsverzeichnis und Zeitlimit, mit denen Ihre Befehle laufen.

    [:octicons-arrow-right-24: Shell](concepts/shell.md)

</div>
