---
title: Curl URL
description: Die mitgelieferte Funktion curl-url ruft eine URL mit curl ab. Fügen Sie die URL mit dem Schlüssel url direkt zu einer Schaltfläche hinzu.
icon: material/web
---

# :material-web: Curl URL

`curl-url` ruft eine URL mit `curl` ab. Die Funktion schlägt bei HTTP-Fehlern
fehl (`-f`) und bricht nach 30 Sekunden ab. Sie gehört zu den
[mitgelieferten Funktionen](../index.md#custom-functions), die Sie unverändert
von einer Schaltfläche aus verwenden können.

- Führt aus: `curl -fsSL --max-time 30 {{.url}}`
- `url` (erforderlich): die abzurufende URL

!!! example "Die Funktionsdatei"

    ```yaml title="functions/curl-url.yaml"
    name: curl-url
    run: "curl -fsSL --max-time 30 {{.url}}"
    params:
      - name: url
        type: string
        required: true
        description: Abzurufende URL
    ```

## Eine Schaltfläche hinzufügen

Schreiben Sie `url` direkt auf die Schaltfläche. Der Name entspricht dem von
der Funktion deklarierten Parameter.

!!! example "Einen Endpunkt prüfen"

    ```yaml title="Schaltfläche zum Prüfen der API"
    - name: API prüfen
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

Die Schaltfläche führt
`curl -fsSL --max-time 30 https://example.com/health` aus.
[`validate`](../../cli.md#validate) meldet einen Fehler, wenn `url` fehlt.

## Verwandte Themen

- [Parameter](../../concepts/parameter.md) — wie Schaltflächenwerte geprüft werden
- [Eigene Funktionen](../index.md#custom-functions) — alle fünf mitgelieferten Beispiele
