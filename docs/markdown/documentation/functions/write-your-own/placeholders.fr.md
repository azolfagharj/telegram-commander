---
title: Variables de remplacement
description: Écrivez la commande run d’une fonction avec des variables de remplacement, y compris des parties facultatives présentes uniquement lorsqu’un bouton fournit une valeur.
icon: material/code-braces
---

# :material-code-braces: Variables de remplacement

Le champ `run` est un petit modèle. Les variables placées entre doubles
accolades sont remplacées par les valeurs des paramètres avant l’exécution.

Deux formes couvrent presque tous les besoins :

- `{{.name}}` insère la valeur du paramètre nommé `name`.
- `{{if .name}} ... {{end}}` inclut la partie centrale seulement si `name`
  possède une valeur.

## Insérer une valeur

!!! example "Une valeur au milieu d’une commande"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
    ```

Un bouton avec `path: "/var/log/app.log"` exécute
`tail -n 200 /var/log/app.log`.

## Parties facultatives

Une valeur vide rend `{{if .name}}` faux et tout le bloc disparaît. Vous pouvez
ainsi ajouter une option uniquement lorsqu’elle est souhaitée.

!!! example "Créer une partie facultative"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}{{if .args}} | grep -- {{.args}}{{end}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        description: Optional text to filter for
    ```

Deux boutons l’utilisent, l’un avec filtre et l’autre sans :

!!! example "Fournir des valeurs depuis un bouton"

    ```yaml title="Boutons des journaux de l’application"
    - name: App log
      type: button
      function: tail-log
      path: "/var/log/app.log"

    - name: App errors
      type: button
      function: tail-log
      path: "/var/log/app.log"
      args: "ERROR"
    ```

**App log** exécute `tail -n 200 /var/log/app.log` et **App errors** exécute
`tail -n 200 /var/log/app.log | grep -- ERROR`.

Vous pouvez aussi choisir entre deux formes avec
`{{if .args}} ... {{else}} ... {{end}}`.

## Noms utilisables

Une variable peut utiliser tout paramètre déclaré dans `params`. Sa valeur est
celle du bouton ou le `default` du paramètre.

!!! warning "Un nom non déclaré échoue lors de l’exécution"

    [`validate`](../../cli.md#validate) vérifie uniquement la syntaxe des
    variables, pas leurs noms. Si `run` contient `{{.uri}}` sans paramètre
    `uri`, la configuration est validée mais le bouton signale une erreur dans
    la conversation.

!!! warning "Les valeurs sont insérées comme texte brut"

    Rien n’est automatiquement placé entre guillemets ni échappé. Une valeur
    contenant des espaces ou des caractères shell fait partie de la commande
    telle quelle. Préférez des valeurs fixes sur les boutons et ajoutez
    [`confirm: true`](../../concepts/confirmation.md) aux actions destructrices.

## Pages associées

- [Structure des fichiers](file-structure.md) — emplacement de `run` et `params`
- [Règles](rules.md) — ce que le chargeur accepte
- [Guide pas à pas](step-by-step.md) — essayer une fonction de bout en bout
