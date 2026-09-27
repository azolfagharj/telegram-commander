---
title: Erneute Zustellversuche
description: Ein kurzer Netzwerkausfall lässt kein Ergebnis mehr verloren gehen. Der Bot wiederholt den Sendevorgang mit wachsender Wartezeit, hält sich an Telegrams eigene Wartezeit bei Ratenbegrenzung und behält die Reihenfolge pro Chat.
icon: material/refresh
---

# :material-refresh: Erneute Zustellversuche

Fällt die Verbindung zu Telegram für einen Moment aus, während der Bot ein
Ergebnis sendet, geht die Nachricht nicht verloren. Der Bot versucht es von
selbst erneut, mit jeweils etwas längerer Wartezeit, bis der Sendevorgang
gelingt.

Bittet Telegram selbst den Bot, langsamer zu senden, wartet der Bot genau so
lange, wie Telegram verlangt, statt eine eigene Wartezeit zu schätzen.

Nachrichten an denselben Chat kommen weiterhin in der Reihenfolge an, in der
Sie sie ausgelöst haben — auch wenn Sie mehrere Schaltflächen antippen,
während die Verbindung Probleme hat.

Dauert der Ausfall zu lange, gibt der Bot irgendwann für dieses eine Ergebnis
auf. Im Chat erscheint nichts, und es bleibt nichts halb gesendet zurück — der
Versuch wird nur im Server-Log vermerkt, nie an Sie angezeigt.

!!! tip "Für ein normales Setup ist nichts einzustellen"

    Die Standardwerte fangen einen kurzen Netzwerkaussetzer von selbst auf.
    Ändern Sie, wie lange der Bot es weiter versucht, mit
    `delivery_retry_backoff`, `delivery_retry_backoff_max` und
    `delivery_retry_ttl` in Ihrer [Konfigurationsdatei](config-file.md).

## Konfiguration

Informationen zu `delivery_retry_backoff`, `delivery_retry_backoff_max` und
`delivery_retry_ttl` finden Sie unter [Konfiguration](../configuration.md).

## Verwandte Themen

- [Bestätigung](confirmation.md) — ein weiterer Moment, in dem der Bot wartet, bevor er handelt
- [Verbindung des Bots](long-polling.md) — die ausgehende Verbindung, die hierdurch abgesichert wird
- [Ausgabe steuern](../output-control.md) — wie ein Ergebnis in Ihrem Chat ankommt
