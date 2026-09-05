---
title: Shell
description: Chaque commande passe par un shell, /bin/bash par défaut, afin que les tubes et redirections fonctionnent. Modifiez-le à la racine de votre configuration.
icon: material/console-line
---

# :material-console-line: Shell

Le programme qui exécute votre commande, `/bin/bash` par défaut. Les commandes
s’exécutent sous la forme `bash -c "your command"`, les tubes et les
redirections fonctionnent donc.

Vous pouvez modifier le shell globalement avec la clé `shell` de votre
[fichier de configuration](config-file.md).

!!! example "Choisir le shell"

    ```yaml title="Racine de config.yaml"
    shell: /bin/bash
    ```

Chaque [fonction](function.md) exécute sa commande avec ce shell.

## Configuration

Pour le champ `shell` et les options racine associées, consultez
[Configuration → Champs racine](../configuration.md#root-fields).

## Pages associées

- [Fonction](function.md) — produit la commande exécutée par le shell
- [Fichier de configuration](config-file.md) — emplacement du réglage `shell`
