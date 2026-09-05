---
title: command
description: La fonction intégrée command exécute une commande shell exactement comme elle est écrite sur le bouton, avec tubes, redirections et commandes enchaînées.
icon: material/console
---

# :material-console: `command`

`command` est une [fonction intégrée](../index.md#built-in-functions). Elle
exécute une commande shell exactement comme vous l’avez écrite sur le bouton.

| Paramètre | Obligatoire | Par défaut | Signification |
|-----------|-------------|------------|---------------|
| `command` | oui | — | La commande shell à exécuter |

!!! example "Exécuter une commande telle qu’elle est écrite"

    ```yaml title="Bouton d’affichage du noyau"
    - name: Show kernel
      type: button
      function: command
      command: "uname -a"
    ```

## Tubes et commandes enchaînées

Les commandes passent par [`/bin/bash -c`](../../concepts/shell.md), les tubes,
les redirections et `&&` fonctionnent donc dans une seule valeur `command`.

!!! example "Les tubes fonctionnent dans les commandes"

    ```yaml title="Bouton des processus principaux"
    - name: Top processes
      type: button
      function: command
      command: "ps aux --sort=-%mem | head -n 10"
    ```

!!! tip "Utilisez une fonction personnalisée pour les commandes répétées"

    Si plusieurs boutons reprennent la même forme de commande, créez une
    fonction personnalisée et ne fournissez à chaque bouton que les valeurs qui
    changent. Des paramètres comme `url` ou `host` s’écrivent directement sur
    ces boutons.

## Pages associées

- [`script`](script.md) — exécuter un script plutôt qu’une commande en ligne
- [Fonctions intégrées](../index.md#built-in-functions) — les deux fonctions intégrées
- [Shell](../../concepts/shell.md) — exécution des commandes
- [Confirmation](../../concepts/confirmation.md) — demander avant une commande risquée
