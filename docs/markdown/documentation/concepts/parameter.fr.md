---
icon: material/tune
title: Paramètre
description: Un paramètre est une valeur nommée dont une fonction a besoin. Écrivez chaque valeur directement sur le bouton avec le nom déclaré par la fonction.
---

# :material-tune: Paramètre

Une valeur nommée dont une [fonction](function.md) a besoin. Par exemple, la
fonction intégrée `command` nécessite un paramètre nommé `command`. Vous
fournissez les valeurs des paramètres sur le [bouton](button.md).

Certains paramètres sont obligatoires ; d’autres sont facultatifs et possèdent
une valeur par défaut définie dans le fichier de fonction.

## Transmettre une valeur depuis un bouton

Écrivez le nom du paramètre comme clé directement sur le bouton. Les champs
intégrés `command`, `path` et `args` suivent cette règle, tout comme les noms de
paramètres personnalisés.

!!! example "Fournir des noms de paramètres personnalisés"

    ```yaml title="Journaux récents du service"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`unit` et `lines` doivent être déclarés par `journal-unit`. La valeur numérique
de `lines` n’a pas besoin de guillemets.

!!! warning "N’utilisez pas de bloc `params:` imbriqué"

    Un fichier de fonction personnalisée utilise `params:` pour déclarer ses
    paramètres. Un bouton ne l’utilise pas. Placez `host:`, `unit:` ou toute
    autre valeur directement sur le bouton.

## Validation

[`validate`](../cli.md#validate) vérifie que :

- chaque paramètre obligatoire possède une valeur ;
- chaque valeur du bouton est déclarée par sa fonction ;
- les valeurs déclarées comme `int` contiennent un entier ;
- les valeurs déclarées comme `bool` contiennent un booléen valide ;
- les valeurs par défaut correspondent au type déclaré.

Les paramètres facultatifs utilisent leur valeur par défaut si le bouton ne les
fournit pas.

Les noms de paramètres ne peuvent pas être identiques aux réglages de bouton :
`name`, `type`, `icon`, `id`, `function`, `confirm`, `timeout`, `workdir`,
`env`, `columns` ou `items`. Les noms `command`, `path` et `args` sont autorisés.

Consultez [Fonctions → Transmettre des valeurs depuis un bouton](../functions/index.md#passing-values-from-a-button)
pour les explications et exemples complets.

## Configuration

Pour les réglages de bouton et les clés de paramètres, consultez
[Configuration → Menu](../configuration.md#menu).

## Pages associées

- [Fonction](function.md) — ce qui utilise les paramètres
- [Règles](../functions/write-your-own/rules.md) — règles des paramètres pour les fonctions personnalisées
