---
title: Exécuter dans la CLI
description: Lancez votre bot depuis le terminal, du téléchargement et de la petite configuration jusqu’au premier appui sur un bouton.
icon: material/console
---

# :material-console: Exécuter dans la CLI

Cette page vous guide de zéro jusqu’à un bot en fonctionnement. Aucune
expérience préalable du projet n’est nécessaire. Si un terme n’est pas clair,
consultez les pages [Concepts](../concepts/config-file.md).

## :material-clipboard-check-outline: Avant de commencer

Vous avez besoin de deux éléments fournis par Telegram :

1. **Un jeton de bot.** Ouvrez une conversation avec
   [@BotFather](https://t.me/BotFather), envoyez `/newbot`, suivez les
   instructions et copiez le jeton reçu. Il ressemble à
   `123456789:AAExampleTokenValue`.
2. **Votre identifiant utilisateur numérique.** Il s’agit d’un nombre, pas de
   votre `@username`. Si vous ne le connaissez pas, le bot vous l’indiquera
   lors de votre premier message (voir [l’étape 5](#step-5-find-your-user-id-if-needed)).

## :material-download: Étape 1 : télécharger

Téléchargez l’archive de la version et extrayez-la.

!!! example "Télécharger et ouvrir le dossier de la version"

    ```bash title="Télécharger et extraire la version"
    wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
    tar -xzf telegram-commander.tar.gz
    cd telegram-commander
    ```

Le dossier contient :

- `telegram-commander-linux-amd64` et `telegram-commander-linux-arm64` — le programme pour chaque type de processeur
- `config-examples/` — des [configurations](../concepts/config-file.md) prêtes à l’emploi (voir [Configuration](../configuration.md))
- `functions/` — des [fonctions](../concepts/function.md) personnalisées d’exemple (voir [Fonctions personnalisées](../functions/index.md#custom-functions))

## :material-chip: Étape 2 : choisir le fichier exécutable

!!! info "Quel fichier convient à votre machine ?"

    La plupart des serveurs et PC utilisent `amd64` (aussi appelé `x86_64`).
    Les petites cartes ARM et certaines machines virtuelles utilisent `arm64`.

    En cas de doute, exécutez `uname -m` : `x86_64` signifie amd64 et
    `aarch64` signifie arm64.

=== ":fontawesome-brands-linux: AMD64"

    ```bash title="Conserver le programme amd64"
    mv telegram-commander-linux-amd64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-arm64
    ```

=== ":fontawesome-brands-linux: ARM64"

    ```bash title="Conserver le programme arm64"
    mv telegram-commander-linux-arm64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-amd64
    ```

Vous disposez maintenant d’un seul programme nommé `telegram-commander`.

## :material-file-cog-outline: Étape 3 : créer la configuration

Copiez l’exemple minimal dans un fichier de travail :

!!! example "Créer une configuration modifiable"

    ```bash title="Copier l’exemple de configuration"
    cp config-examples/config.minimal.yaml ./config.yaml
    ```

Ouvrez `config.yaml` et remplacez deux valeurs :

- `YOUR_BOT_TOKEN` — le jeton de BotFather
- `YOUR_USER_ID` — votre identifiant numérique (ou laissez-le pour l’instant et consultez l’étape 5)

Pour connaître chaque réglage, consultez [Configuration](../configuration.md).

## :material-file-check-outline: Étape 4 : valider

Vérifiez toujours la configuration avant le lancement. Cela détecte les fautes
sans démarrer le bot.

!!! success "Vérifier que la configuration fonctionne"

    ```bash title="Valider la configuration"
    ./telegram-commander validate --config config.yaml
    ```

Si `Valid configuration` s’affiche, tout est prêt. Sinon, la commande indique
précisément le problème et son emplacement. Consultez la
[page CLI](../cli.md#validate).

## :material-account-search: Étape 5 : trouver votre identifiant utilisateur { #step-5-find-your-user-id-if-needed }

Si vous ne connaissez pas votre identifiant, définissez seulement le jeton dans
`config.yaml`, placez provisoirement un nombre dans `allowed_users`, puis lancez
le bot :

!!! info "Démarrer une fois pour voir votre identifiant"

    ```bash title="Lancer le bot pour connaître votre identifiant"
    ./telegram-commander run --config config.yaml
    ```

Dans Telegram, trouvez votre bot et envoyez-lui un message. Comme vous ne
figurez pas encore parmi les [utilisateurs autorisés](../concepts/allowed-users.md),
le bot répond avec votre `user_id` et votre `username`. Copiez cet identifiant
dans `allowed_users`, arrêtez le bot avec `Ctrl+C`, puis relancez-le.

Ce comportement fait partie du contrôle d’accès ; consultez
[Configuration → telegram](../configuration.md#telegram).

## :material-play-circle-outline: Étape 6 : exécuter

!!! example "Démarrer le bot dans le terminal"

    ```bash title="Démarrer le bot"
    ./telegram-commander run --config config.yaml
    ```

Ouvrez le bot dans Telegram et envoyez `/start`. Votre menu doit apparaître.
Appuyez sur un [bouton](../concepts/button.md) pour exécuter sa commande.

!!! success "Votre bot fonctionne"

    Le menu décrit dans `config.yaml` est maintenant dans la conversation et
    chaque appui exécute sa commande sur cette machine.

Pour garder le bot actif après votre déconnexion du serveur, configurez-le
comme service. Consultez [Exécuter comme service](run-as-a-service.md).

## :material-map-marker-path: Étape suivante

- Ajouter des [boutons](../concepts/button.md) et [catégories](../concepts/category.md) : [Menu](../concepts/menu.md)
- Comprendre ce qui s’exécute : [Fonctions](../functions/index.md)
- Voir toutes les options de ligne de commande : [CLI](../cli.md)
