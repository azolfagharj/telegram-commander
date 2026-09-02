---
title: Telegram Commander
description: Verwandeln Sie eine YAML-Datei in einen Telegram-Bot, der Befehle auf Ihrem Linux-Server ausführt — Schritt für Schritt mit einem Antippen.
icon: material/cellphone-link
hide:
  - navigation
  - toc
---

# :material-cellphone-link: Ihren Linux-Server über Telegram steuern

<div class="hero" markdown>
<div class="hero-art" markdown>
![Telegram Commander](/images/logo-large.png){ .off-glb width="230" }
</div>
<div class="hero-text" markdown>
**Ein Antippen führt einen Befehl auf Ihrem Server aus und sendet die Ausgabe an Ihren Chat zurück.**

Telegram Commander verwandelt eine einfache YAML-[Konfigurationsdatei](documentation/concepts/config-file.md)
in einen Telegram-Bot mit einem Menü aus [Schaltflächen](documentation/concepts/button.md).
Legen Sie einen beliebigen Terminalbefehl hinter eine Schaltfläche und führen
Sie ihn von Ihrem Smartphone aus. Sie müssen keinen Code schreiben.
</div>
</div>

<div style="text-align: center" markdown="span">
[Erste Schritte :material-arrow-right:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Konfiguration ansehen :material-file-code-outline:](documentation/configuration.md#a-minimal-config){ .md-button }
</div>

[Installation](documentation/installation/download-and-run.md) ·
[Grundlagen](documentation/concepts/config-file.md) ·
[Funktionen](documentation/functions/index.md) ·
[Konfiguration](documentation/configuration.md) ·
[CLI](documentation/cli.md)

## :material-image-multiple-outline:{ .shots } Screenshots { .split }

Ihr Menü, ein laufender Befehl, die zurückgesendete Ausgabe und die manuelle
Eingabe eines Befehls. Klicken Sie ein Bild an, um es in voller Größe zu sehen.

<div style="text-align: center" markdown="span">
  ![Befehlsausgabe bei geöffnetem System-Menü](/images/01.jpeg){ width="140" loading=lazy }
  ![Schaltflächen für Ressourcen und Prozesse](/images/02.jpeg){ width="140" loading=lazy }
  ![Schaltflächen für Speicher und Pakete](/images/03.jpeg){ width="140" loading=lazy }
  ![Netzwerkwerkzeuge und manuelle Befehlseingabe](/images/04.jpeg){ width="140" loading=lazy }
  ![An den Chat zurückgesendete Befehlsausgabe](/images/05.jpeg){ width="140" loading=lazy }
</div>

## :material-lightning-bolt:{ .bolt } Schnell und einfach { .split }

<div class="grid cards cols-3 center-title step-cards" markdown>

-   :material-file-document-outline:{ .middle } __Menü schreiben__

    ---

    :material-numeric-1-circle:{ .step } Schaltflächen und zugehörige Befehle auflisten.

-   :material-rocket-launch:{ .middle } __Bot starten__

    ---

    :material-numeric-2-circle:{ .step } Sofort ausführen oder dauerhaft als Dienst betreiben.

-   :material-gesture-tap-button:{ .middle } __Antippen und lesen__

    ---

    :material-numeric-3-circle:{ .step } Eine Schaltfläche antippen und die Ausgabe im Chat lesen.

</div>

<div style="text-align: center" markdown="span">
[Jetzt starten :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
</div>

## :material-view-grid-outline:{ .grid-icon } Anwendungsfälle { .split }

<div class="grid cards cols-4 icon-left" markdown>

-   :material-restart:{ .lg } Einen Dienst neu starten oder stoppen
-   :material-docker:{ .lg } Container starten und stoppen
-   :material-package-down:{ .lg } Systempakete aktualisieren
-   :material-text-box-search-outline:{ .lg } Protokolle und Journale lesen
-   :material-harddisk:{ .lg } Speicherplatz prüfen
-   :material-memory:{ .lg } CPU und Arbeitsspeicher überwachen
-   :material-access-point-network:{ .lg } Hosts pingen und URLs testen
-   :material-backup-restore:{ .lg } Sicherungen erstellen und wiederherstellen
-   :material-script-text:{ .lg } Eigene Skripte ausführen
-   :material-power:{ .lg } Host neu starten oder herunterfahren
-   :material-console:{ .lg } Einen beliebigen Befehl manuell eingeben
-   :material-all-inclusive:{ .lg } Und fast alles andere

</div>

## :material-thumb-up-outline:{ .thumb } Ihre Vorteile { .split }

<div class="grid cards cols-4 center-title" markdown>

-   :material-clock-fast:{ .lg .middle } __Keine Programmierung__

    ---

    Menü und Befehle in einer YAML-Datei beschreiben.

    [:octicons-arrow-right-24: Konfigurationsdatei](documentation/concepts/config-file.md)

-   :material-cellphone-link:{ .lg .middle } __Von überall__

    ---

    Telegram auf dem Smartphone öffnen und den Server bedienen. Kein VPN zum Host erforderlich.

    [:octicons-arrow-right-24: So stellt der Bot die Verbindung her](documentation/concepts/long-polling.md)

-   :material-lan-disconnect:{ .lg .middle } __Keine offenen Ports__

    ---

    Der Bot verbindet sich nach außen mit Telegram. Nichts ist im Internet erreichbar.

    [:octicons-arrow-right-24: So stellt der Bot die Verbindung her](documentation/concepts/long-polling.md)

-   :material-message-text-outline:{ .lg .middle } __Ausgabe im Chat__

    ---

    Das Ergebnis kommt als Nachricht zurück. Eine SSH-Sitzung ist nicht erforderlich.

    [:octicons-arrow-right-24: Umfang der angezeigten Ausgabe](documentation/configuration.md#how-much-command-output-you-see)

-   :material-shield-lock:{ .lg .middle } __Kontrolliert und protokolliert__

    ---

    Zugriffsberechtigte Personen auswählen, riskante Aktionen bestätigen und jede Ausführung protokollieren.

    [:octicons-arrow-right-24: Zugriff und Bestätigung](documentation/concepts/allowed-users.md)

-   :material-folder-outline:{ .lg .middle } __Verschachtelte Menüs__

    ---

    Schaltflächen in Kategorien gruppieren. Home bleibt oben; Back führt eine Ebene zurück.

    [:octicons-arrow-right-24: Menü](documentation/concepts/menu.md)

-   :material-function-variant:{ .lg .middle } __Wiederverwendbare Funktionen__

    ---

    Einen Befehl einmal schreiben und auf jeder Schaltfläche andere Werte einsetzen.

    [:octicons-arrow-right-24: Funktionen](documentation/functions/index.md)

-   :material-cog-play-outline:{ .lg .middle } __Dauerhaft verfügbar__

    ---

    Als Dienst installieren, damit der Bot mit dem Host startet.

    [:octicons-arrow-right-24: Als Dienst ausführen](documentation/installation/run-as-a-service.md)

</div>

## :material-file-code-outline:{ .code-icon } Ein kleines Beispiel { .split }

Diese Konfiguration erstellt einen Bot mit einer Schaltfläche namens „Laufzeit“.
Beim Antippen wird der Befehl `uptime` auf dem Server ausgeführt.

!!! example "Diese vollständige Konfiguration fügt eine Schaltfläche hinzu"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN" # (1)!
      allowed_users:
        - "YOUR_USER_ID" # (2)!

    menu:
      - name: Laufzeit
        type: button
        function: command
        command: "uptime" # (3)!
    ```

    1.  Fordern Sie beim Erstellen des Bots in Telegram ein Token von BotFather an.
    2.  Nur die hier aufgeführten Konten können das Menü öffnen. Sie können eine numerische ID oder einen `@username` verwenden.
    3.  Hier können Sie alles eintragen, was sich in einem Terminal eingeben lässt.

Dies ist eine vollständige, funktionierende Konfiguration. Alles Weitere ist optional.

!!! tip "Für eine kleine, vertrauenswürdige Gruppe entwickelt"

    Nichts wartet auf eingehende Verbindungen, und nur Konten in
    `allowed_users` erhalten ein Menü. Jeder Bot-Benutzer kann Ihre definierten
    Schaltflächen ausführen. Halten Sie die Liste daher kurz.

## :material-hand-pointing-right: Bereit zum Ausprobieren? { .split }

<div style="text-align: center" markdown>
[Jetzt starten :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Grundlagen :material-book-open-variant:](documentation/concepts/config-file.md){ .md-button }
[Neueste Version :material-download:](https://github.com/azolfagharj/telegram-commander/releases/latest){ .md-button .md-button--primary }

[Quellcode ansehen :fontawesome-brands-github:](https://github.com/azolfagharj/telegram-commander){ .md-button }
</div>

Telegram Commander ist kostenlos und quelloffen. Wenn es Ihnen Zeit spart,
[können Sie die Entwicklung unterstützen](https://azolfagharj.github.io/donate/) —
so bleibt das Projekt aktiv und gepflegt.
