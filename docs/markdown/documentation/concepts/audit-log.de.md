---
title: Audit-Protokoll
description: Protokollieren Sie jeden vom Bot ausgeführten Befehl, einschließlich der Person, die die Schaltfläche angetippt hat, der verwendeten Schaltfläche, des Exit-Codes und der Ausführungsdauer.
icon: material/history
---

# :material-history: Audit-Protokoll

Ein separater Protokollstrom, der jeden ausgeführten Befehl erfasst: wer ihn
ausgeführt hat, welche [Schaltfläche](button.md) verwendet wurde, den Exit-Code
und die Ausführungsdauer. Er wird unter `logging` in Ihrer
[Konfigurationsdatei](config-file.md) eingerichtet.

In diesem Protokoll können Sie später nachsehen, wer um drei Uhr morgens einen
Dienst neu gestartet hat oder ob die Sicherung der letzten Nacht tatsächlich
abgeschlossen wurde. Da es sich um einen eigenen
Protokollstrom handelt, können Sie ihn aufbewahren, nachdem das normale
Anwendungsprotokoll bereits rotiert wurde.

Jede Ausführung fügt eine Zeile hinzu. Zusammen beantworten die Zeilen vier Fragen:

| Frage | Was die Zeile aussagt |
|---|---|
| Wer? | Das Telegram-Konto, das die Schaltfläche angetippt hat |
| Was? | Die verwendete Schaltfläche |
| War es erfolgreich? | Der vom Befehl zurückgegebene Exit-Code |
| Wie lange? | Die Zeit bis zum Abschluss des Befehls |

!!! example "Audit-Protokoll in eine eigene Datei schreiben"

    ```yaml title="logging-Abschnitt von config.yaml"
    logging:
      logs:
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

## Konfiguration

Das vollständige `logging`-Schema und die unterstützten Ausgaben finden Sie
unter [Konfiguration → logging](../configuration.md#logging).

## Verwandte Themen

- [Konfigurationsdatei](config-file.md) — hier wird die Protokollierung definiert
- [Funktion](function.md) — was beim Antippen einer Schaltfläche ausgeführt wird
