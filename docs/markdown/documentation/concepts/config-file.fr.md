---
icon: material/file-document-outline
title: Fichier de configuration
description: Un seul fichier YAML contient le jeton de votre bot, les utilisateurs autorisés et le menu de boutons. Utilisez --config avec run, validate, fmt et list-functions.
---

# :material-file-document-outline: Fichier de configuration

Un seul fichier YAML décrit tout : le jeton de votre bot, les personnes qui
peuvent l’utiliser et le menu de [boutons](button.md). Vous le transmettez avec
`--config` aux commandes qui le lisent : `run`, `validate`, `fmt` et
`list-functions`. Les autres commandes, comme `version` et `completion`,
n’utilisent pas de fichier de configuration.

!!! example "Une configuration fonctionnelle avec un bouton"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

## Pages associées

- [Configuration](../configuration.md) — tous les champs, valeurs par défaut et règles de validation
- [Menu](menu.md) — l’arborescence des boutons et catégories
- [CLI](../cli.md) — transmettre le fichier avec `--config`
- [Exécuter dans la CLI](../installation/download-and-run.md) — créer votre première configuration
