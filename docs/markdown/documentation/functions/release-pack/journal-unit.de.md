---
title: Journal Unit
description: Die mitgelieferte Funktion journal-unit zeigt aktuelle journalctl-Protokolle einer systemd-Unit. Geben Sie die Unit und optional die Zeilenanzahl direkt auf der Schaltfläche an.
icon: material/text-box-search-outline
---

# :material-text-box-search-outline: Journal Unit

`journal-unit` zeigt die neuesten `journalctl`-Protokolle einer systemd-Unit.
Die Funktion gehört zu den
[mitgelieferten Funktionen](../index.md#custom-functions), die Sie unverändert
von einer Schaltfläche aus verwenden können.

- Führt aus: `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (erforderlich): Unit-Name, zum Beispiel `nginx.service`
- `lines` (optional, Standardwert `50`): Anzahl der anzuzeigenden Zeilen

!!! example "Die Funktionsdatei"

    ```yaml title="functions/journal-unit.yaml"
    name: journal-unit
    run: "journalctl -u {{.unit}} -n {{.lines}} --no-pager"
    params:
      - name: unit
        type: string
        required: true
        description: Name der systemd-Unit (zum Beispiel nginx.service)
      - name: lines
        type: string
        required: false
        default: "50"
        description: Anzahl der Protokollzeilen
    ```

## Eine Schaltfläche hinzufügen

!!! example "Aktuelle Protokolle eines Dienstes lesen"

    ```yaml title="Schaltfläche für nginx-Protokolle"
    - name: nginx-Protokolle
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`lines` ist numerisch und benötigt daher keine Anführungszeichen. Lassen Sie
das Feld weg, um den Standardwert `50` zu verwenden.

!!! example "Die standardmäßige Zeilenanzahl verwenden"

    ```yaml title="Schaltfläche für SSH-Protokolle"
    - name: SSH-Protokolle
      type: button
      function: journal-unit
      unit: "ssh.service"
    ```

## Verwandte Themen

- [Platzhalter](../write-your-own/placeholders.md) — wie `{{.unit}}` eingesetzt wird
- [Eigene Funktionen](../index.md#custom-functions) — alle fünf mitgelieferten Beispiele
