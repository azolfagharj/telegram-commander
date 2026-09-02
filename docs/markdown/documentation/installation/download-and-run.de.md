---
title: In der CLI ausführen
description: Bringen Sie Ihren Bot Schritt für Schritt im Terminal zum Laufen — vom Herunterladen der Version und Schreiben einer kleinen Konfigurationsdatei bis zum Antippen der ersten Schaltfläche.
icon: material/console
---

# :material-console: In der CLI ausführen

Diese Seite führt Sie von den ersten Schritten bis zu einem laufenden Bot.
Vorkenntnisse mit dem Projekt sind nicht erforderlich. Wenn ein Begriff unklar
ist, sehen Sie auf den Seiten zu den
[Grundlagen](../concepts/config-file.md) nach.

## :material-clipboard-check-outline: Voraussetzungen

Sie benötigen zwei Dinge von Telegram:

1. **Ein Bot-Token.** Öffnen Sie einen Chat mit
   [@BotFather](https://t.me/BotFather), senden Sie `/newbot`, folgen Sie den
   Anweisungen und kopieren Sie das erhaltene Token. Es sieht etwa so aus:
   `123456789:AAExampleTokenValue`.
2. **Ihre numerische Benutzer-ID.** Dies ist eine Zahl, nicht Ihr `@username`.
   Wenn Sie sie nicht kennen, ist das kein Problem — der Bot teilt Ihnen Ihre
   ID bei Ihrer ersten Nachricht mit (siehe [Schritt 5](#step-5-find-your-user-id-if-needed)).

## :material-download: Schritt 1: Herunterladen

Laden Sie das Release-Archiv herunter und entpacken Sie es.

!!! example "Das Release herunterladen und den Ordner öffnen"

    ```bash title="Das Release herunterladen und entpacken"
    wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
    tar -xzf telegram-commander.tar.gz
    cd telegram-commander
    ```

Im Ordner finden Sie:

- `telegram-commander-linux-amd64` und `telegram-commander-linux-arm64` — das Programm, jeweils für einen CPU-Typ
- `config-examples/` — einsatzbereite [Konfigurationsdateien](../concepts/config-file.md) (siehe [Konfiguration](../configuration.md))
- `functions/` — Beispiele für eigene [Funktionen](../concepts/function.md) (siehe [Eigene Funktionen](../functions/index.md#custom-functions))

## :material-chip: Schritt 2: Passende Binärdatei auswählen

!!! info "Welche Datei passt zu Ihrem Rechner?"

    Die meisten Server und PCs verwenden `amd64` (auch `x86_64` genannt).
    Kleine ARM-Boards und einige Cloud-VMs verwenden `arm64`.

    Wenn Sie unsicher sind, führen Sie `uname -m` aus: `x86_64` bedeutet amd64,
    `aarch64` bedeutet arm64.

=== ":fontawesome-brands-linux: AMD64"

    ```bash title="Das amd64-Programm behalten"
    mv telegram-commander-linux-amd64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-arm64
    ```

=== ":fontawesome-brands-linux: ARM64"

    ```bash title="Das arm64-Programm behalten"
    mv telegram-commander-linux-arm64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-amd64
    ```

Nun besitzen Sie ein einzelnes Programm namens `telegram-commander`.

## :material-file-cog-outline: Schritt 3: Konfiguration erstellen

Kopieren Sie das Minimalbeispiel in eine Arbeitsdatei:

!!! example "Eine bearbeitbare Konfiguration erstellen"

    ```bash title="Die Beispielkonfiguration kopieren"
    cp config-examples/config.minimal.yaml ./config.yaml
    ```

Öffnen Sie `config.yaml` und ersetzen Sie zwei Platzhalter:

- `YOUR_BOT_TOKEN` — das Token von BotFather
- `YOUR_USER_ID` — Ihre numerische ID (oder lassen Sie den Wert zunächst stehen und lesen Sie Schritt 5)

Die Bedeutung aller Einstellungen wird unter
[Konfiguration](../configuration.md) erläutert.

## :material-file-check-outline: Schritt 4: Validieren

Prüfen Sie die Konfiguration immer vor dem Start. So werden Tippfehler und
Probleme erkannt, ohne den Bot zu starten.

!!! success "Prüfen, ob die Konfiguration funktioniert"

    ```bash title="Die Konfiguration validieren"
    ./telegram-commander validate --config config.yaml
    ```

Wenn `Valid configuration` ausgegeben wird, ist alles in Ordnung. Andernfalls
werden die Fehler und ihre Positionen genau aufgelistet. Weitere Informationen
finden Sie auf der [CLI-Seite](../cli.md#validate).

## :material-account-search: Schritt 5: Benutzer-ID ermitteln (falls erforderlich) { #step-5-find-your-user-id-if-needed }

Wenn Sie Ihre Benutzer-ID nicht kennen, tragen Sie in `config.yaml` zunächst
nur das Token ein, verwenden Sie vorläufig eine beliebige Zahl in
`allowed_users` und starten Sie dann den Bot:

!!! info "Einmal starten, um Ihre Benutzer-ID zu sehen"

    ```bash title="Den Bot starten, um Ihre ID zu ermitteln"
    ./telegram-commander run --config config.yaml
    ```

Öffnen Sie Telegram, suchen Sie Ihren Bot und senden Sie ihm eine beliebige
Nachricht. Da Sie noch nicht zu den
[zugelassenen Benutzern](../concepts/allowed-users.md) gehören, antwortet der
Bot mit Ihrer `user_id` und Ihrem `username`. Kopieren Sie diese ID nach
`allowed_users`, beenden Sie den Bot mit `Ctrl+C` und starten Sie ihn erneut.

Dieses Verhalten ist Teil der Zugriffskontrolle; siehe
[Konfiguration → telegram](../configuration.md#telegram).

## :material-play-circle-outline: Schritt 6: Ausführen

!!! example "Den Bot im Terminal starten"

    ```bash title="Den Bot starten"
    ./telegram-commander run --config config.yaml
    ```

Öffnen Sie Ihren Bot in Telegram und senden Sie `/start`. Nun sollte Ihr Menü
erscheinen. Tippen Sie eine [Schaltfläche](../concepts/button.md) an, um ihren
Befehl auszuführen.

!!! success "Ihr Bot ist aktiv"

    Das in `config.yaml` beschriebene Menü befindet sich nun in Ihrem Chat, und
    jedes Antippen führt den zugehörigen Befehl auf diesem Rechner aus.

Um den Bot auch nach dem Abmelden vom Server weiterlaufen zu lassen, richten
Sie ihn als Dienst ein. Siehe [Als Dienst ausführen](run-as-a-service.md).

## :material-map-marker-path: Nächste Schritte

- Weitere [Schaltflächen](../concepts/button.md) und [Kategorien](../concepts/category.md) hinzufügen: [Menü](../concepts/menu.md)
- Verstehen, was tatsächlich ausgeführt wird: [Funktionen](../functions/index.md)
- Alle Befehlszeilenoptionen anzeigen: [CLI](../cli.md)
