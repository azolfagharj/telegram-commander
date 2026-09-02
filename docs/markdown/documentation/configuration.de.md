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

!!! example "Mit einem zugelassenen Benutzer und einer Schaltfläche beginnen"

    ```yaml title="config.yaml (minimal)"
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
| `max_output_bytes` | Ganzzahl | nein | `524288` | Maximal aufbewahrte Ausgabe pro Befehl (siehe [Umfang der angezeigten Befehlsausgabe](#how-much-command-output-you-see)) |
| `workdir` | Zeichenfolge | nein | Prozess-cwd | Standardarbeitsverzeichnis für Befehle |
| `env` | Zuordnung | nein | leer | Zusätzliche Umgebungsvariablen für Befehle |
| `menu_columns` | Ganzzahl | nein | `2` | Eintragsschaltflächen pro Zeile unter dem Nachrichtenfeld |
| `page_size` | Ganzzahl | nein | `8` | Einträge pro Seite vor der Seitennavigation |
| `confirm_ttl` | Dauer | nein | `5m` | Gültigkeitsdauer einer [Bestätigungs](concepts/confirmation.md)-aufforderung |
| `enable_run_command` | bool | nein | `false` | Zeigt eine Schaltfläche **`$ >_ Run Command`**, die die nächste Nachricht als Shell-Befehl ausführt. Standardmäßig aus. Jeder Bot-Benutzer kann damit jeden Befehl auf dem Host ausführen. Aktivieren Sie dies nur, wenn Sie allen zugelassenen Benutzern vertrauen. Unter `telegram` ist dieser Schlüssel ungültig. |
| `logging` | Objekt | nein | integrierter Standard-Logger | Benannte Logger (siehe unten) |

??? note "Was geschieht, wenn ich `shell` auslasse?"

    Sie können das Feld auslassen. Der Bot verwendet `/bin/bash`. Dasselbe gilt
    für `timeout`, `page_size` und andere optionale Felder: Ohne Angabe gelten
    die Standardwerte. Legen Sie sie nur für einen abweichenden Wert fest
    (beispielsweise `shell: /bin/sh`).

### Umfang der angezeigten Befehlsausgabe { #how-much-command-output-you-see }

Zwei Limits werden nacheinander angewendet. `max_output_bytes` ist **Ihr** Limit
und gilt zusätzlich zu einem unveränderlichen Telegram-Limit.

**1. Ihr Limit: `max_output_bytes`** (Standardwert `524288`, also 512 KB)

Während ein Befehl läuft, behält der Bot jeweils höchstens diese Menge seiner
Standard- und Fehlerausgabe. Darüber hinausgehende Daten werden verworfen, der
Befehl läuft jedoch bis zum Ende oder bis zu seinem `timeout` weiter. In diesem
Fall beginnt das Ergebnis mit `(output truncated)`.

**2. Telegrams Limit: Eine Nachricht fasst höchstens 4096 Byte**

Dieses Limit ist fest. Längere Ergebnisse teilt der Bot in mehrere Nachrichten.
Jeder Teil antwortet auf den vorherigen, sodass Reihenfolge und Zusammenhang
erhalten bleiben; die Menüschaltflächen erscheinen am letzten Teil. Wenn
möglich wird an Zeilengrenzen geteilt.

Ist das Ergebnis danach immer noch sehr lang, stoppt der Bot nach 10 Nachrichten.
Die letzte endet mit einem Hinweis wie
`(output too long; showing first N bytes)`, wobei `N` die tatsächlich
empfangene Ausgabemenge angibt.

Ein höheres `max_output_bytes` lässt den Bot mehr Ausgabe behalten, sichtbar
sind jedoch höchstens ungefähr zehn Nachrichten. Kürzen Sie sehr lange Befehle
(zum Beispiel `journalctl -u nginx | tail -n 50`) oder schreiben Sie die
vollständige Ausgabe in eine Datei auf dem Server.

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
| `api` | Zeichenfolge | nein | `https://api.telegram.org` | Basis-URL der Bot API |
| `proxy.enabled` | bool | nein | `false` | Proxy für die Telegram API verwenden |
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
        - "123456789"        # numeric user id
        - "@alice"           # or a username
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
finden Sie unter [Menü](concepts/menu.md). Jeder
[Schaltflächen](concepts/button.md)- oder [Kategorie](concepts/category.md)-Knoten:

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
          format: console   # or json
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
- [Funktionen](functions/index.md) — Bedeutung von `function`, `command`, `path` und `args`
- [CLI](cli.md) — Ihre Konfiguration validieren und ausführen
