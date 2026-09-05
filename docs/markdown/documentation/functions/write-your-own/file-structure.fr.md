---
title: Structure des fichiers
description: Tous les champs d’un fichier de fonction Telegram Commander, de name et run à la liste params avec type, required, default et description.
icon: material/file-tree
---

# :material-file-tree: Structure des fichiers

Un fichier décrit une fonction. Placez-le n’importe où dans votre
[`function_directory`](../../configuration.md#function_directory-rules), y
compris dans un sous-dossier, avec l’extension `.yaml` ou `.yml`. Le nom du
fichier n’a pas d’importance : le nom de la fonction vient du champ `name`.

!!! example "Définir une fonction personnalisée par fichier"

    ```yaml title="Fichier de fonction personnalisée"
    name: my-function          # obligatoire, doit être unique et non réservé
    run: "echo {{.args}}"      # obligatoire, commande à exécuter
    params:                    # liste facultative de paramètres
      - name: args             # obligatoire pour chaque paramètre
        type: string           # facultatif : string (par défaut), int ou bool
        required: true         # facultatif, false par défaut
        default: ""            # valeur facultative si elle n’est pas fournie
        description: Some text  # note facultative pour vous-même
    ```

## Champs de premier niveau

| Champ | Obligatoire | Signification |
|-------|-------------|---------------|
| `name` | oui | Nom utilisé par les boutons dans leur champ `function` |
| `run` | oui | Commande à exécuter, avec des [variables de remplacement](placeholders.md) |
| `params` | non | Liste des valeurs nommées acceptées par la fonction |

Une fonction sans `params` est valide : `run` est alors une commande fixe.

## Champs d’un paramètre

| Champ | Obligatoire | Par défaut | Signification |
|-------|-------------|------------|---------------|
| `name` | oui | — | Nom utilisé dans `{{.name}}` au sein de `run` |
| `type` | non | `string` | Type : `string`, `int` ou `bool` |
| `required` | non | `false` | Un bouton doit fournir une valeur |
| `default` | non | vide | Valeur utilisée si aucune n’est fournie |
| `description` | non | vide | Note personnelle, non affichée dans Telegram |

!!! info "Le `type` est vérifié"

    Les valeurs et valeurs par défaut déclarées comme `int` doivent contenir un
    entier. Celles déclarées comme `bool` doivent contenir un booléen valide.
    Sinon, [`validate`](../../cli.md#validate) échoue.

!!! info "Les clés du bouton correspondent aux noms des paramètres"

    Écrivez chaque paramètre directement sur un bouton avec le même nom. Des
    noms comme `url`, `host` et `lines` fonctionnent avec `command`, `path` et
    `args`. Consultez [Transmettre des valeurs depuis un bouton](../index.md#passing-values-from-a-button).

## Organisation du dossier

Vous pouvez organiser le dossier comme vous le souhaitez :

!!! example "Les sous-dossiers sont également lus"

    ```text title="functions/"
    functions/
      disk.yaml
      logs/
        nginx.yaml
        app.yml
    ```

Les trois fichiers sont chargés. Les autres extensions sont ignorées.

## Pages associées

- [Règles](rules.md) — ce que le chargeur refuse
- [Variables de remplacement](placeholders.md) — écrire la commande `run`
- [Guide pas à pas](step-by-step.md) — créer votre première fonction
