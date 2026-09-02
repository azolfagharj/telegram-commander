---
title: Ping Host
description: Die mitgelieferte Funktion ping-host pingt einen Host mehrmals. Geben Sie den Host und optional die Anzahl direkt auf der Schaltfläche an.
icon: material/access-point-network
---

# :material-access-point-network: Ping Host

`ping-host` pingt einen Host mehrmals und sendet das Ergebnis zurück. Die
Funktion gehört zu den
[mitgelieferten Funktionen](../index.md#custom-functions), die Sie unverändert
von einer Schaltfläche aus verwenden können.

- Führt aus: `ping -c {{.count}} {{.host}}`
- `host` (erforderlich): Hostname oder IP-Adresse
- `count` (optional, Standardwert `4`): Anzahl der Pakete

!!! example "Die Funktionsdatei"

    ```yaml title="functions/ping-host.yaml"
    name: ping-host
    run: "ping -c {{.count}} {{.host}}"
    params:
      - name: host
        type: string
        required: true
        description: Hostname oder IP-Adresse
      - name: count
        type: string
        required: false
        default: "4"
        description: Anzahl der Ping-Pakete
    ```

## Eine Schaltfläche hinzufügen

!!! example "Einen festen Host dreimal pingen"

    ```yaml title="Schaltfläche zum Pingen des Gateways"
    - name: Gateway pingen
      type: button
      function: ping-host
      host: "192.168.1.1"
      count: 3
    ```

`count` ist numerisch und benötigt daher keine Anführungszeichen. Sie können
das Feld auch weglassen, um den Standardwert `4` zu verwenden:

!!! example "Den Standardwert für die Anzahl verwenden"

    ```yaml title="Schaltfläche zum Pingen von DNS"
    - name: DNS pingen
      type: button
      function: ping-host
      host: "1.1.1.1"
    ```

## Verwandte Themen

- [Regeln](../write-your-own/rules.md) — Verhalten von Standardwerten und erforderlichen Werten
- [Eigene Funktionen](../index.md#custom-functions) — alle fünf mitgelieferten Beispiele
