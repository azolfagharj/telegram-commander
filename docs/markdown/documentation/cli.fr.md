---
title: Référence CLI
description: Toutes les options de ligne de commande Telegram Commander, du lancement du bot à la validation et au formatage de la configuration avec --config.
icon: material/console-line
---

# :material-console-line: Référence CLI

!!! info "Les commandes utilisent cette forme"

    ```text title="Syntaxe de commande"
    telegram-commander <command> [flags]
    ```

`--config` / `-c` est **obligatoire** pour les commandes qui chargent la
configuration. Le chemin peut être relatif au dossier de travail actuel ou
absolu. Le contenu du fichier est décrit dans [Configuration](configuration.md).

Si vous débutez, [Exécuter dans la CLI](installation/download-and-run.md)
présente ces commandes dans leur ordre d’utilisation.

## :material-format-list-bulleted-square: Commandes

### `run` { #run }

Exécute le bot au premier plan.

!!! info "Démarrer le bot au premier plan"

    ```bash title="Démarrer le bot"
    telegram-commander run --config /path/to/configfile.yaml
    ```

!!! note "Les modifications nécessitent un redémarrage"

    Le bot lit la configuration une fois au démarrage. Après une modification,
    redémarrez le processus, par exemple avec
    `systemctl restart telegram-commander`.

### `validate` { #validate }

Valide hors ligne la configuration, les fonctions et les références des boutons.

!!! info "Vérifier la configuration hors ligne ou en ligne"

    ```bash title="Valider la configuration"
    telegram-commander validate --config /path/to/configfile.yaml
    telegram-commander validate --config /path/to/configfile.yaml --online
    ```

!!! note "`--online` nécessite Internet"

    Avec `--online`, la vérification demande aussi à Telegram si le jeton du bot
    fonctionne. La machine doit donc pouvoir joindre Telegram.

### `version`

Affiche la version du programme.

### `fmt`

Affiche proprement un fichier de configuration YAML.

!!! info "Afficher ou enregistrer le YAML formaté"

    ```bash title="Formater un fichier de configuration"
    telegram-commander fmt --config /path/to/configfile.yaml
    telegram-commander fmt --config /path/to/configfile.yaml -w
    ```

### `environ`

Affiche les variables d’environnement du processus, utile pour diagnostiquer
les unités de service.

### `list-functions` { #list-functions }

Liste les fonctions intégrées et personnalisées chargées. Utilisez cette
commande pour confirmer que vos fichiers ont été trouvés. Consultez
[Fonctions](functions/index.md).

!!! info "Afficher toutes les fonctions disponibles"

    ```bash title="Lister les fonctions disponibles"
    telegram-commander list-functions --config /path/to/configfile.yaml
    ```

### `completion`

Génère les scripts de complétion du shell :

!!! info "Choisir votre shell"

    ```bash title="Générer un script de complétion"
    telegram-commander completion bash
    telegram-commander completion zsh
    telegram-commander completion fish
    telegram-commander completion powershell
    ```

### `manpage`

Écrit une page de manuel sur la sortie standard.

## Pages associées

- [Configuration](configuration.md) — le fichier transmis avec `--config`
- [Fonctions personnalisées](functions/index.md#custom-functions) — ce qu’affiche `list-functions`
- [Exécuter comme service](installation/run-as-a-service.md) — lancer `run` sous systemd
