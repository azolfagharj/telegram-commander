---
icon: material/tune
title: Parameter
description: Ein Parameter ist ein benannter Wert, den eine Funktion benötigt. Schreiben Sie jeden Wert direkt auf die Schaltfläche und verwenden Sie dabei denselben Namen, den die Funktion deklariert.
---

# :material-tune: Parameter

Ein benannter Wert, den eine [Funktion](function.md) benötigt. Die integrierte
Funktion `command` benötigt beispielsweise einen Parameter namens `command`.
Sie geben Parameterwerte auf der [Schaltfläche](button.md) an.

Einige Parameter sind erforderlich, andere sind optional und besitzen einen in
der Funktionsdatei definierten Standardwert.

## Einen Wert von einer Schaltfläche übergeben

Schreiben Sie den Parameternamen als Schlüssel direkt auf die Schaltfläche. Die
integrierten Felder `command`, `path` und `args` folgen dieser Regel; eigene
Parameternamen funktionieren genauso.

!!! example "Eigene Parameternamen angeben"

    ```yaml title="Aktuelle Dienstprotokolle"
    - name: nginx-Protokolle
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`unit` und `lines` müssen von `journal-unit` deklariert sein. Der numerische
Wert für `lines` muss nicht in Anführungszeichen stehen.

!!! warning "Verwenden Sie keine verschachtelte `params:`-Zuordnung"

    Eine Datei für eine eigene Funktion verwendet `params:`, um ihre Parameter
    zu deklarieren. Eine Schaltfläche tut dies nicht. Schreiben Sie `host:`,
    `unit:` oder jeden anderen Wert direkt auf die Schaltfläche.

## Validierung

[`validate`](../cli.md#validate) prüft Folgendes:

- Für jeden erforderlichen Parameter ist ein Wert vorhanden.
- Jeder Wert auf der Schaltfläche ist von ihrer Funktion deklariert.
- Als `int` deklarierte Werte enthalten eine Ganzzahl.
- Als `bool` deklarierte Werte enthalten einen gültigen booleschen Wert.
- Standardwerte entsprechen dem deklarierten Parametertyp.

Für ausgelassene optionale Parameter wird der jeweilige Standardwert verwendet.

Parameternamen dürfen nicht mit Schaltflächeneinstellungen übereinstimmen:
`name`, `type`, `icon`, `id`, `function`, `confirm`, `timeout`, `workdir`,
`env`, `columns` oder `items`. Die Namen `command`, `path` und `args` sind
zulässig.

Die vollständige Erklärung und Beispiele finden Sie unter
[Funktionen → Werte von einer Schaltfläche übergeben](../functions/index.md#passing-values-from-a-button).

## Konfiguration

Informationen zu Schaltflächeneinstellungen und Parameterschlüsseln finden Sie
unter [Konfiguration → Menü](../configuration.md#menu).

## Verwandte Themen

- [Funktion](function.md) — was Parameter verwendet
- [Regeln](../functions/write-your-own/rules.md) — Parameterregeln für eigene Funktionen
