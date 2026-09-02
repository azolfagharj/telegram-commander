---
title: So stellt der Bot die Verbindung her
description: Der Bot stellt eine ausgehende Verbindung zu Telegram her, sodass Sie auf Ihrem Server keinen Port öffnen müssen. Richten Sie einen Proxy ein, wenn Telegram in Ihrem Netzwerk blockiert ist.
icon: material/lan-disconnect
---

# :material-lan-disconnect: So stellt der Bot die Verbindung her

Der Bot stellt eine ausgehende Verbindung zu Telegram her. Sie öffnen keinen
Port auf Ihrem Server.

!!! tip "Nichts muss im Internet erreichbar sein"

    Da die Verbindung ausschließlich nach außen aufgebaut wird, gibt es kein
    abzusicherndes Web-Panel, keine eingehende Firewall-Regel und keine
    öffentliche Adresse, die jemand finden könnte. Der Rechner kann sich hinter
    einem Heimrouter oder einer strengen Firewall befinden und der Bot
    funktioniert trotzdem.

Wenn Telegram in Ihrem Netzwerk blockiert ist, richten Sie unter `telegram` in
Ihrer [Konfigurationsdatei](config-file.md) einen Proxy ein.

## Konfiguration

Informationen zu `api`, `proxy` und `insecure` in den Telegram-Einstellungen
finden Sie unter [Konfiguration → telegram](../configuration.md#telegram).

## Verwandte Themen

- [Zugelassene Benutzer](allowed-users.md) — wer den Bot verwenden darf
- [CLI → run](../cli.md#run) — den Bot im Vordergrund starten
- [Als Dienst ausführen](../installation/run-as-a-service.md) — den Bot dauerhaft ausführen
