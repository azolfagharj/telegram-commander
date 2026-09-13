---
icon: material/file-cog-outline
title: Konfiguration
description: Alle Einstellungen von Telegram Commander mit Typ, Standardwert und Bedeutung — telegram, menu, function_directory, Zeitüberschreitungen, Ausgabelimits und Protokollierung.
---

# :material-file-cog-outline: Konfiguration

Die [Konfigurationsdatei](concepts/config-file.md) beschreibt Ihren gesamten
Bot: die Telegram-Verbindung, zugelassene Benutzer, das
[Schaltflächen](concepts/button.md)-Menü und die Protokollierung. Sie übergeben
sie mit `--config` an `run`, `validate`, `fmt` und `list-functions` (siehe
[CLI](cli.md)).

Alle Schlüssel verwenden `lower_snake_case`. **Unbekannte Schlüssel werden
abgelehnt**, sodass ein Tippfehler bei der [Validierung](cli.md#validate) sofort
sichtbar wird.

**Erforderlich** bedeutet, dass die Validierung fehlschlägt, wenn das Feld nach
Anwendung der Standardwerte fehlt oder leer ist.  
**Optionale** Felder können fehlen; die Spalte Standardwert zeigt den dann
verwendeten Wert.

Wenn Sie neu im Projekt sind, beginnen Sie mit
[In der CLI ausführen](installation/download-and-run.md), wo die erste
Konfiguration schrittweise erstellt wird. Die verwendeten Begriffe werden
unter [Grundlagen](concepts/config-file.md) erläutert.

## :material-rocket-launch-outline: Eine minimale Konfiguration { #a-minimal-config }

Nur `telegram` (mit Token und einem
[zugelassenen Benutzer](concepts/allowed-users.md)) sowie `menu` sind
erforderlich. Alles andere besitzt einen Standardwert:

!!! example "Mit einem zugelassenen Benutzer und drei Status-Schaltflächen beginnen"

    ```yaml title="config.yaml (minimal)"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: System Status
        type: category
        icon: "🖥️"
        items:
          - name: Uptime
            type: button
            icon: "🕒"
            function: command
            command: "uptime"
          - name: Memory
            type: button
            icon: "🧮"
            function: command
            command: "free -h"
          - name: Disk Space
            type: button
            icon: "💾"
            function: command
            command: "df -h /"
    ```

Der Ordner `config-examples/` im Release enthält ein minimales und ein
vollständiges Beispiel.

## :material-card-bulleted-outline: Felder auf oberster Ebene { #root-fields }

| Feld | Typ | Erforderlich | Standardwert | Beschreibung |
|-------|------|----------|---------|-------------|
| `telegram` | Objekt | ja | — | Telegram-Einstellungen (siehe unten) |
| `menu` | Liste | ja | — | Menübaum; mindestens ein Knoten |
| `function_directory` | Zeichenfolge | nein | nicht gesetzt | YAML-Verzeichnis eigener Funktionen (siehe Regeln unten) |
| `shell` | Zeichenfolge | nein | `/bin/bash` | Als `shell -c "<command>"` verwendete [Shell](concepts/shell.md) |
| `timeout` | Dauer | nein | `60s` | Standardmäßige Befehlszeitüberschreitung |
| `max_output_bytes` | Ganzzahl | nein | `524288` | Maximal aufbewahrte Ausgabe pro Befehl (siehe [Umfang der angezeigten Befehlsausgabe](#how-much-command-output-you-see) und [Ausgabe steuern](output-control.md)) |
| `output` | `auto` \| `text` \| `file` | nein | `auto` | Wie Befehlsergebnisse zugestellt werden (siehe [Umfang der angezeigten Befehlsausgabe](#how-much-command-output-you-see) und [Ausgabe steuern](output-control.md)) |
| `max_output_messages` | Ganzzahl | nein | `2` | Nur auf Root-Ebene. Im Modus `auto`: eine `.txt`-Datei senden, wenn das Ergebnis mehr als so viele Nachrichten bräuchte (1–10). Auslassen, um `2` zu behalten. Siehe [Ausgabe steuern](output-control.md) |
| `workdir` | Zeichenfolge | nein | Arbeitsverzeichnis des Prozesses | Standardarbeitsverzeichnis für Befehle |
| `env` | Zuordnung | nein | leer | Zusätzliche Umgebungsvariablen für Befehle |
| `menu_columns` | Ganzzahl | nein | `2` | Menüeinträge pro Zeile unter dem Nachrichtenfeld |
| `page_size` | Ganzzahl | nein | `8` | Einträge pro Seite vor der Seitennavigation |
| `confirm_ttl` | Dauer | nein | `5m` | Gültigkeitsdauer einer Aufforderung zur [Bestätigung](concepts/confirmation.md) |
| `enable_run_command` | bool | nein | `false` | Zeigt eine Schaltfläche **`$ >_ Run Command`**, die die nächste Nachricht als Shell-Befehl ausführt. Standardmäßig aus. Jeder Bot-Benutzer kann damit jeden Befehl auf dem Host ausführen. Aktivieren Sie dies nur, wenn Sie allen zugelassenen Benutzern vertrauen. Unter `telegram` ist dieser Schlüssel ungültig. |
| `logging` | Objekt | nein | integrierter Standard-Logger | Benannte Logger (siehe unten) |

??? note "Was geschieht, wenn ich `shell` auslasse?"

    Sie können das Feld auslassen. Der Bot verwendet `/bin/bash`. Dasselbe gilt
    für `timeout`, `page_size` und andere optionale Felder: Ohne Angabe gelten
    die Standardwerte. Legen Sie sie nur für einen abweichenden Wert fest
    (beispielsweise `shell: /bin/sh`).

### Umfang der angezeigten Befehlsausgabe { #how-much-command-output-you-see }

Zwei Limits werden nacheinander angewendet. `max_output_bytes` ist **Ihr** Limit
und kommt zusätzlich zu einem Telegram-Limit, das Sie nicht ändern können.

**1. Ihr Limit: `max_output_bytes`** (Standardwert `524288`, also 512 KB)

Während ein Befehl läuft, behält der Bot höchstens so viel Ausgabe, getrennt
für normale Ausgabe und Fehlerausgabe. Alles darüber wird verworfen, der Befehl
läuft aber weiter, bis er endet oder sein `timeout` erreicht. In diesem Fall
beginnt das Ergebnis mit `(output truncated)`.

**2. Zustellung des Ergebnisses: `output`** (Standardwert `auto`)

Telegram erlaubt höchstens 4096 Bytes pro Nachricht. Der Bot kann ein langes
Ergebnis in mehrere Antwortnachrichten teilen oder die vollständige Ausgabe als
`.txt`-Datei senden. Wählen Sie den Modus mit `output` in der Root-Konfiguration
oder überschreiben Sie ihn an einem einzelnen Button:

| Wert | Verhalten |
|------|-----------|
| `auto` | Textnachrichten, solange das Ergebnis in `max_output_messages` Nachrichten passt (Standard `2`). Braucht es mehr, wird eine `.txt`-Datei gesendet. |
| `text` | Immer in Textnachrichten teilen. Stoppt weiterhin nach 10 Nachrichten und vermerkt den Abbruch. |
| `file` | Immer eine `.txt`-Datei mit kurzer Beschriftung (Buttonname, Exit-Code, Dauer) senden. |

In den Modi `auto` und `file` enthält die Datei Kopfzeile sowie stdout und
stderr — ein höheres `max_output_bytes` liefert Ihnen also wirklich mehr Ausgabe.
Schlägt das Senden der Datei fehl, fällt der Bot auf den Textnachrichten-Weg zurück.

`output` und `max_output_messages` müssen Sie nicht schreiben. Lassen Sie sie
weg, verwendet der Bot `auto` und `2`. Eine bereits funktionierende Datei
bleibt gleich. `max_output_messages` gibt es nur auf Root-Ebene. An einem
Button lassen Sie `output` weg, um den Root-Wert zu nutzen; setzen Sie ihn
nur, wenn genau dieser Button immer eine Datei oder immer Text senden soll.

Für lange Logs eignet sich `output: auto` (Standard) oder `output: file` an
diesem Button. Den Befehl selbst können Sie trotzdem kürzen, wenn nur ein
Ausschnitt nötig ist.

!!! abstract ":material-export-variant: Mehr dazu: Ausgabe steuern"

    Diesen drei Feldern ist eine eigene Seite gewidmet. Sie geht jede
    Kombination aus Root und Button einzeln durch und zeigt zu jeder ein Bild
    des Telegram-Chats.

    [:octicons-arrow-right-24: Ausgabe steuern lesen](output-control.md)

### Regeln für `function_directory` { #function_directory-rules }

| Situation | Ergebnis |
|-----------|--------|
| Schlüssel fehlt | Info-Protokoll; nur integrierte Funktionen |
| Schlüssel vorhanden, aber leer (`""`) | Info-Protokoll; nur integrierte Funktionen |
| Pfad existiert nicht oder ist nicht zugänglich | Schwerer Fehler; Prozess stoppt |
| Pfad existiert, Verzeichnis ist aber leer | OK |

!!! warning "Ein falscher Pfad stoppt den Bot"

    Verweist `function_directory` auf einen nicht vorhandenen oder nicht
    lesbaren Ordner, stoppt das Programm mit einem Fehler, statt ohne Ihre
    eigenen Funktionen zu starten.

## :material-send-circle-outline: `telegram` { #telegram }

| Feld | Typ | Erforderlich | Standardwert | Beschreibung |
|-------|------|----------|---------|-------------|
| `bot_token` | Zeichenfolge | ja | — | Bot-Token von BotFather |
| `allowed_users` | Liste von Zeichenfolgen | ja | — | [Zugelassene Benutzer](concepts/allowed-users.md) |
| `api` | Zeichenfolge | nein | `https://api.telegram.org` | Basis-URL der Bot-API |
| `proxy.enabled` | bool | nein | `false` | Proxy für die Telegram-API verwenden |
| `proxy.url` | Zeichenfolge | bedingt | — | Erforderlich, wenn `proxy.enabled` auf `true` steht |
| `insecure` | bool | nein | `false` | TLS-Prüfung überspringen (nicht empfohlen) |

Nicht autorisierte Benutzer erhalten ihre `user_id` und ihren `username`, damit
sie einen Administrator um Zugriff bitten können. So finden Sie beim ersten
Mal auch Ihre eigene ID — siehe
[In der CLI ausführen → Schritt 5](installation/download-and-run.md#step-5-find-your-user-id-if-needed).

!!! example "Verbindung über einen Proxy herstellen"

    ```yaml title="telegram-Abschnitt mit Proxy"
    telegram:
      bot_token: "123456789:AAExampleTokenValue"
      allowed_users:
        - "123456789"        # numerische Benutzer-ID
        - "@alice"           # oder ein Benutzername
      proxy:
        enabled: true
        url: "socks5://127.0.0.1:10808"
    ```

Damit zugelassene Benutzer einen Shell-Befehl in Telegram eingeben können,
legen Sie Folgendes auf der **obersten Ebene** fest (nicht unter `telegram`):

!!! tip "Eine Einstellung auf oberster Ebene hinzufügen"

    ```yaml title="Die Schaltfläche Run Command aktivieren"
    enable_run_command: true
    ```

## :material-menu: Menü { #menu }

Dieser Abschnitt ist die Feldreferenz. Eine geführte Erläuterung mit Beispielen
finden Sie unter [Menü](concepts/menu.md). Jeder Knoten vom Typ
[Schaltfläche](concepts/button.md) oder [Kategorie](concepts/category.md):

| Feld | Typ | Erforderlich | Beschreibung |
|-------|------|----------|-------------|
| `name` | Zeichenfolge | ja | Anzeigename (unter Geschwistern eindeutig, Groß-/Kleinschreibung ignoriert) |
| `type` | `category` \| `button` | ja | Knotenart |
| `items` | Liste | bei `category` | Unterknoten; Kategorie benötigt mindestens einen |
| `function` | Zeichenfolge | bei `button` | Name der [Funktion](concepts/function.md) |
| `command` | Zeichenfolge | bei `function: command` | Shell-Befehl für das integrierte `command` |
| `path` | Zeichenfolge | bei `function: script` | Skriptpfad für das integrierte `script` |
| `icon` | Zeichenfolge | nein | Optionales Emoji-Präfix |
| `id` | Zeichenfolge | nein | Optionale ID dieses Knotens. Sie können sie weglassen. |
| `confirm` | bool | nein | Vor der Ausführung [Bestätigung](concepts/confirmation.md) verlangen (Standard `false`) |
| `timeout` | Dauer | nein | Globale Zeitüberschreitung überschreiben |
| `workdir` | Zeichenfolge | nein | Arbeitsverzeichnis überschreiben |
| `env` | Zuordnung | nein | Zusätzliche Umgebungsvariablen für diese Schaltfläche |
| `columns` | Ganzzahl | nein | Spalten für diese Kategorie überschreiben |
| `output` | `auto` \| `text` \| `file` | nein | Optional. Weglassen, um den Root-Wert von `output` zu nutzen. Nur setzen, um diesen Button auf `file` oder `text` zu zwingen (siehe [Umfang der angezeigten Befehlsausgabe](#how-much-command-output-you-see) und [Ausgabe steuern](output-control.md)) |
| `args` | Zeichenfolge | nein | Optionale Argumente für `script` |
| Jeder deklarierte Parametername | Skalar | wie von der Funktion deklariert | An die Funktion übergebener Wert, z. B. `url`, `host`, `unit` oder `lines` |

Auf einer **Schaltfläche** gilt jeder weitere skalare Schlüssel als
Funktionsparameter. Sein Name muss einem Parameter der ausgewählten Funktion
entsprechen. Unbekannte Namen lassen [`validate`](cli.md#validate) fehlschlagen.
Als `int` oder `bool` deklarierte Werte werden ebenfalls geprüft. Zeichenfolgen,
Zahlen und boolesche Werte können direkt als YAML-Werte geschrieben werden;
Zahlen benötigen keine Anführungszeichen.

Bei einer **Kategorie** ist jeder nicht oben aufgeführte Schlüssel ein Fehler.
Kategorien führen keine Funktionen aus und können keine Parameterschlüssel besitzen.

`command`, `path` und `args` sind Kurzfelder für gleichnamige Parameter. Andere
[Parameternamen](concepts/parameter.md) stehen direkt auf der Schaltfläche,
nicht in einer verschachtelten `params:`-Zuordnung. Siehe
[Funktionen → Werte von einer Schaltfläche übergeben](functions/index.md#passing-values-from-a-button).

## :material-math-log: `logging` { #logging }

Optional. Ohne Angabe wird ein standardmäßiger Konsolen-Logger auf `stderr` mit
Stufe `info` verwendet.

Benannte Logger:

!!! example "Normale Protokolle und eine Audit-Datei schreiben"

    ```yaml title="logging-Abschnitt mit Audit-Datei"
    logging:
      logs:
        default:
          level: info
          format: console   # oder JSON
          output:
            - output: stderr
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

Unterstützte Ausgaben: `stdout`, `stderr`, `file`, `discard`.

Der gezeigte Logger `audit` erfasst jede Befehlsausführung (Person,
Schaltfläche, Exit-Code und Dauer). Siehe
[Audit-Protokoll](concepts/audit-log.md).

## Verwandte Seiten

- [In der CLI ausführen](installation/download-and-run.md) — eine erste Konfiguration erstellen und ausführen
- [Menü](concepts/menu.md) — der Menübaum im Detail
- [Ausgabe steuern](output-control.md) — Nachrichten oder eine `.txt`-Datei, mit jeder Kombination aus Root und Button
- [Funktionen](functions/index.md) — Bedeutung von `function`, `command`, `path` und `args`
- [CLI](cli.md) — Ihre Konfiguration validieren und ausführen
