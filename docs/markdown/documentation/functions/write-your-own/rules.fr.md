---
title: Règles
description: Les vérifications d’un fichier de fonction Telegram Commander : noms autorisés et réservés, unicité, refus des fautes et valeurs par défaut.
icon: material/format-list-checks
---

# :material-format-list-checks: Règles

Chaque fichier de fonction est vérifié au démarrage du bot et avec
[`validate`](../../cli.md#validate). Si un fichier est incorrect, rien ne
s’exécute : les erreurs sont détectées avant la mise en ligne du menu.

## Fichiers lus

- Seuls les fichiers terminés par `.yaml` ou `.yml` sont lus. Les autres sont ignorés.
- Les sous-dossiers de `function_directory` sont également lus.
- Un fichier doit décrire exactement une fonction.

## Nommage

- **`name` est obligatoire** et doit respecter `^[A-Za-z0-9._-]+$` : lettres,
  chiffres, `.`, `-` et `_`, sans espace.
- **Les noms sont uniques sans tenir compte de la casse.** `Deploy` et `deploy`
  entrent en conflit, même dans des fichiers différents.
- **Les noms réservés sont interdits.** Une fonction ne peut pas s’appeler
  `command` ou `script` ; ces noms appartiennent aux
  [fonctions intégrées](../index.md#built-in-functions).

## La commande

- **`run` est obligatoire.** Écrivez la commande avec des
  [variables de remplacement](placeholders.md) aux emplacements des valeurs.
- La syntaxe des variables est vérifiée lors de la validation. Une séquence
  `{{` incorrecte est donc signalée tôt.

## Paramètres

- **Chaque paramètre nécessite un `name`** conforme à la même règle de caractères.
- **`type` doit être** `string`, `int` ou `bool`. Sans valeur, il vaut `string`.
  Les valeurs et valeurs par défaut `int` ou `bool` sont vérifiées.
- Un nom de paramètre ne peut pas être un réglage de bouton : `name`, `type`,
  `icon`, `id`, `function`, `confirm`, `timeout`, `workdir`, `env`, `columns`
  ou `items`. `command`, `path` et `args` sont autorisés.
- Chaque valeur écrite sur un bouton doit correspondre à un paramètre déclaré
  par sa fonction.

Le champ `run` ne peut utiliser que les variables de paramètres déclarés. Une
variable non déclarée est une erreur.

## Les fautes sont des erreurs

**Les clés inconnues sont refusées.** Une faute comme `requird:` au lieu de
`required:` arrête la validation au lieu d’être ignorée.

!!! warning "Ce fichier ne se charge pas"

    ```yaml title="functions/broken.yaml"
    name: broken
    run: "echo {{.args}}"
    params:
      - name: args
        requird: true      # refusé : clé inconnue
    ```

## Valeurs obligatoires et valeurs par défaut

- L’absence de valeur pour un **paramètre obligatoire** fait échouer la
  validation avec un message comme `required parameter "args" for function "greet" is missing`.
- Une valeur présente mais vide est considérée comme manquante.
- Un **paramètre facultatif** utilise son `default`. Sans `default`, il devient
  vide, ce qui permet d’omettre les parties facultatives d’une commande.

!!! example "Obligatoire et facultatif dans un fichier"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n {{.args}} {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        default: "200"
        description: Number of lines
    ```

    Un bouton qui définit uniquement `path` exécute
    `tail -n 200 /var/log/app.log`.

## Pages associées

- [Structure des fichiers](file-structure.md) — explication de chaque champ
- [Variables de remplacement](placeholders.md) — écrire la commande `run`
- [CLI → validate](../../cli.md#validate) — exécuter vous-même les vérifications
