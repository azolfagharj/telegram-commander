---
title: Als Dienst ausführen
description: Führen Sie den Bot mit systemd im Hintergrund aus, damit er beim Systemstart und nach einem Fehler automatisch startet — als root oder als normaler Benutzer.
icon: material/server
---

# :material-server: Als Dienst ausführen

Diese Seite beschreibt die Ausführung von Telegram Commander als
systemd-Dienst. Eine schrittweise Anleitung für den ersten Start im Terminal
finden Sie unter [In der CLI ausführen](download-and-run.md).

## :material-cog-outline: systemd (root)

Erstellen Sie die Unit-Datei selbst, zum Beispiel
`/etc/systemd/system/telegram-commander.service`:

!!! warning "In diesem Beispiel wird jede Schaltfläche als root ausgeführt"

    ```ini title="/etc/systemd/system/telegram-commander.service"
    [Unit]
    Description=telegram-commander Telegram bot
    After=network-online.target
    Wants=network-online.target

    [Service]
    Type=simple
    ExecStart=/path/to/telegram-commander run --config /path/to/configfile.yaml
    Restart=on-failure
    RestartSec=5

    [Install]
    WantedBy=multi-user.target
    ```

    Ohne eine Zeile `User=` startet der Dienst als root. Daher besitzen alle
    Befehle in Ihrem Menü vollständige Rechte auf dem Rechner. Fügen Sie dem
    Abschnitt `[Service]` die Zeile `User=someone` hinzu, wenn Ihre
    Schaltflächen nicht so weitreichende Rechte benötigen.

Ersetzen Sie die Platzhalterpfade und führen Sie anschließend Folgendes aus:

!!! example "Den Dienst laden und starten"

    ```bash title="Den Dienst aktivieren, starten und überwachen"
    sudo systemctl daemon-reload
    sudo systemctl enable --now telegram-commander
    sudo systemctl status telegram-commander
    sudo journalctl -u telegram-commander -f
    ```

!!! info "Änderungen an der Konfiguration erfordern einen Neustart"

    Der Dienst hält den Bot in Betrieb, liest Ihre
    [Konfigurationsdatei](../concepts/config-file.md) jedoch nur beim Start.
    Führen Sie nach einer Änderung
    `sudo systemctl restart telegram-commander` aus.

## Verwandte Seiten

- [In der CLI ausführen](download-and-run.md) — erläuterter erster Start
- [Konfiguration](../configuration.md) — die Konfigurationsdatei
- [CLI](../cli.md) — `run`, `validate` und weitere Befehle
