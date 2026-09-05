---
title: Script d’écho
description: La fonction echo-script fournie exécute un script avec Bash, sans exiger le droit d’exécution. Elle reçoit path et args depuis un bouton.
icon: material/bash
---

# :material-bash: Script d’écho

`echo-script` exécute un script avec `bash`. Le fichier lui-même n’a donc pas
besoin du droit d’exécution. C’est l’une des
[fonctions fournies](../index.md#custom-functions) utilisables directement
depuis un bouton.

- Exécute : `bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path` (obligatoire) : chemin du fichier de script
- `args` (facultatif) : arguments supplémentaires

!!! example "Le fichier de fonction"

    ```yaml title="functions/echo-script.yaml"
    name: echo-script
    run: "bash {{.path}}{{if .args}} {{.args}}{{end}}"
    params:
      - name: path
        type: string
        required: true
        description: Path to the script file
      - name: args
        type: string
        required: false
        description: Optional script arguments
    ```

Écrivez les deux noms de paramètres directement sur le bouton.

!!! example "Exécuter un script avec Bash"

    ```yaml title="Bouton de nettoyage"
    - name: Run cleanup
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

Ce bouton exécute `bash /opt/scripts/cleanup.sh --verbose`. Sans `args`, il
exécute `bash /opt/scripts/cleanup.sh`, car la partie `{{if .args}}` est ignorée
lorsque la valeur est vide.

## Pages associées

- [`script`](../built-in/script.md) — exécuter directement un script exécutable
- [Variables de remplacement](../write-your-own/placeholders.md) — fonctionnement de `{{if .args}}`
- [Fonctions personnalisées](../index.md#custom-functions) — les cinq exemples fournis
