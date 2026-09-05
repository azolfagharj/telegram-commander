---
title: script
description: La fonction intégrée script exécute un fichier de script sur le serveur, avec des arguments facultatifs après le chemin. Le script doit être exécutable.
icon: material/script-text
---

# :material-script-text: `script`

`script` est une [fonction intégrée](../index.md#built-in-functions). Elle
exécute un fichier de script avec des arguments facultatifs.

| Paramètre | Obligatoire | Par défaut | Signification |
|-----------|-------------|------------|---------------|
| `path` | oui | — | Chemin du script |
| `args` | non | vide | Arguments transmis après le chemin |

!!! example "Exécuter un script avec des arguments"

    ```yaml title="Bouton du rapport nocturne"
    - name: Nightly report
      type: button
      function: script
      path: "/usr/local/bin/report.sh"
      args: "--today"
    ```

Cette configuration exécute `/usr/local/bin/report.sh --today`. Sans `args`,
elle exécute seulement `/usr/local/bin/report.sh`.

## Le script doit être exécutable

Le chemin est exécuté directement, le fichier doit donc avoir le droit
d’exécution :

!!! tip "Rendre le script exécutable une fois"

    ```bash title="Autoriser l’exécution du fichier"
    chmod +x /usr/local/bin/report.sh
    ```

Si vous ne pouvez pas modifier le fichier, exécutez-le avec un interpréteur.
L’exemple [`echo-script`](../release-pack/echo-script.md) fourni le fait en
appelant d’abord `bash`.

## Pages associées

- [`command`](command.md) — exécuter une commande en ligne plutôt qu’un fichier
- [Script d’écho](../release-pack/echo-script.md) — exécuter un script avec Bash
- [Fonctions intégrées](../index.md#built-in-functions) — les deux fonctions intégrées
- [Shell](../../concepts/shell.md) — exécution des commandes
