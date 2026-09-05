---
title: Unité du journal
description: La fonction journal-unit fournie affiche les journaux journalctl récents d’une unité systemd. Définissez unit et le nombre facultatif de lignes sur le bouton.
icon: material/text-box-search-outline
---

# :material-text-box-search-outline: Unité du journal

`journal-unit` affiche les journaux `journalctl` les plus récents d’une unité
systemd. C’est l’une des [fonctions fournies](../index.md#custom-functions)
utilisables directement depuis un bouton.

- Exécute : `journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit` (obligatoire) : nom de l’unité, par exemple `nginx.service`
- `lines` (facultatif, valeur par défaut `50`) : nombre de lignes à afficher

!!! example "Le fichier de fonction"

    ```yaml title="functions/journal-unit.yaml"
    name: journal-unit
    run: "journalctl -u {{.unit}} -n {{.lines}} --no-pager"
    params:
      - name: unit
        type: string
        required: true
        description: Systemd unit name (for example nginx.service)
      - name: lines
        type: string
        required: false
        default: "50"
        description: Number of log lines
    ```

## Ajouter un bouton

!!! example "Lire les journaux récents d’un service"

    ```yaml title="Bouton des journaux Nginx"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`lines` est numérique et n’a donc pas besoin de guillemets. Omettez-le pour
utiliser la valeur par défaut `50`.

!!! example "Utiliser le nombre de lignes par défaut"

    ```yaml title="Bouton des journaux SSH"
    - name: SSH logs
      type: button
      function: journal-unit
      unit: "ssh.service"
    ```

## Pages associées

- [Variables de remplacement](../write-your-own/placeholders.md) — remplissage de `{{.unit}}`
- [Fonctions personnalisées](../index.md#custom-functions) — les cinq exemples fournis
