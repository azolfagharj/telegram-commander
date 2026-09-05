---
title: Fonction
description: Une fonction transforme les valeurs d’un bouton en commande shell. Utilisez les fonctions intégrées command et script, ou ajoutez les vôtres sous forme de fichiers YAML.
icon: material/function
---

# :material-function: Fonction

Ce qui s’exécute lorsque vous appuyez sur un [bouton](button.md). Une
[fonction](function.md) reçoit des [paramètres](parameter.md) et les transforme
en commande shell.

Il existe deux types :

- Les **fonctions intégrées** sont fournies avec le programme (`command`,
  `script`). Vous ne les créez pas.
- Les **fonctions personnalisées** sont des fonctions supplémentaires que vous
  définissez dans des fichiers YAML placés dans `function_directory`.

Lorsque vous appuyez sur un bouton, le bot recherche la fonction, insère les
valeurs du bouton et exécute le résultat dans le [shell](shell.md).

!!! example "Un bouton utilisant la fonction intégrée `command`"

    ```yaml title="Bouton Uptime"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

## Configuration

Pour `function_directory` et les champs de bouton qui référencent des fonctions,
consultez [Configuration](../configuration.md).

## Pages associées

- [Paramètre](parameter.md) — les valeurs nommées nécessaires à une fonction
- [Fonctions](../functions/index.md#two-kinds-of-function) — fonctions intégrées et personnalisées en détail
- [Menu](menu.md) — comment les boutons référencent les fonctions
