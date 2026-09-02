---
title: Zugelassene Benutzer
description: Legen Sie fest, welche Telegram-Konten das Bot-Menü öffnen dürfen. Fügen Sie Personen über ihre numerische Benutzer-ID oder ihren Benutzernamen hinzu und erfahren Sie, was alle anderen stattdessen sehen.
icon: material/shield-lock
---

# :material-shield-lock: Zugelassene Benutzer

Die Liste der Telegram-Konten, die den Bot verwenden dürfen, angegeben durch
eine numerische `user_id` oder einen `@username`. Alle anderen werden abgewiesen.

Nicht autorisierte Benutzer erhalten eine Nachricht mit ihrer `user_id` und
ihrem `username`, damit sie einen Administrator um Zugriff bitten können. So
finden Sie beim ersten Mal auch Ihre eigene ID — siehe
[In der CLI ausführen → Schritt 5](../installation/download-and-run.md#step-5-find-your-user-id-if-needed).

Mit dieser Nachricht können Sie zugleich eine Person hinzufügen. Bitten Sie die
neue Person, dem Bot eine beliebige Nachricht zu senden, die zurückgesendete ID
zu kopieren, und fügen Sie diese ID zu `allowed_users` hinzu. Starten Sie danach
den Bot neu. Bei der nächsten Nachricht erhält die Person das Menü statt der
Ablehnung.

!!! warning "Alle Personen auf der Liste nutzen dasselbe Menü"

    Es gibt keine benutzerspezifischen Berechtigungen. Jede hinzugefügte Person
    kann jede von Ihnen definierte Schaltfläche antippen. Beschränken Sie die
    Liste daher auf Personen, denen Sie auch den Rechner selbst anvertrauen.

!!! example "Zwei Konten dürfen den Bot verwenden"

    ```yaml title="Telegram-Abschnitt von config.yaml"
    telegram:
      allowed_users:
        - "123456789"
        - "@alice"
    ```

## Konfiguration

Informationen zu `allowed_users` und weiteren Telegram-Einstellungen finden Sie
unter [Konfiguration → telegram](../configuration.md#telegram).

## Verwandte Themen

- [Konfigurationsdatei](config-file.md) — enthält den Abschnitt `telegram`
- [So stellt der Bot die Verbindung her](long-polling.md) — kein eingehender Port muss geöffnet werden
