---
title: Guide pas à pas
description: Créez votre première fonction Telegram Commander, ajoutez un bouton, validez la configuration et exécutez le bot pour voir le résultat.
icon: material/format-list-numbered
---

# :material-format-list-numbered: Guide pas à pas

Ce guide crée une petite fonction, d’un dossier vide jusqu’à un bouton
fonctionnel. Il ne prend que quelques minutes et nécessite uniquement votre
fichier de configuration.

## 1. Indiquer un dossier dans la configuration

!!! example "Ajouter le dossier à votre configuration"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

Le chemin est relatif au fichier de configuration. Consultez
[Configuration → règles de `function_directory`](../../configuration.md#function_directory-rules).

## 2. Écrire le fichier de fonction

Créez `functions/greet.yaml`. Le paramètre se nomme `args` afin qu’un bouton
puisse le remplir :

!!! example "Votre première fonction"

    ```yaml title="functions/greet.yaml"
    name: greet
    run: "echo Hello {{.args}}"
    params:
      - name: args
        required: true
        description: Name to greet
    ```

## 3. Ajouter un bouton qui l’utilise

!!! example "Un bouton pour la nouvelle fonction"

    ```yaml title="Bouton de salutation"
    - name: Say hello
      type: button
      function: greet
      args: "world"
    ```

## 4. Valider et lister les fonctions

!!! example "Vérifier la configuration puis afficher les fonctions chargées"

    ```bash title="Valider puis lister les fonctions"
    ./telegram-commander validate --config config.yaml
    ./telegram-commander list-functions --config config.yaml
    ```

Vous devez voir une ligne `greet`. Chaque ligne indique le nom, l’origine de la
fonction et son nombre de paramètres. Consultez
[CLI → list-functions](../../cli.md#list-functions).

Si `validate` signale un problème, le message nomme le bouton et la valeur
manquante. Consultez [Règles](rules.md).

## 5. Exécuter le bot et utiliser le bouton

!!! example "Démarrer le bot au premier plan"

    ```bash title="Exécuter et suivre la sortie"
    ./telegram-commander run --config config.yaml
    ```

Envoyez `/start` dans Telegram, appuyez sur **Say hello** : le bot exécute
`echo Hello world` et renvoie `Hello world` sous forme de bloc de code.

## 6. Rendre la fonction réutilisable

Modifiez la valeur `args` du bouton ou ajoutez un second bouton avec une autre
valeur. La fonction reste identique :

!!! example "Deux boutons, une fonction"

    ```yaml title="Boutons de salutation"
    - name: Greet world
      type: button
      function: greet
      args: "world"
    - name: Greet team
      type: button
      function: greet
      args: "team"
    ```

## Pages associées

- [Variables de remplacement](placeholders.md) — ajouter des parties facultatives
- [Règles](rules.md) — ce que le chargeur accepte
- [Structure des fichiers](file-structure.md) — explication de chaque champ
- [Menu](../../concepts/menu.md) — emplacement du bouton dans votre menu
