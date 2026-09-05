---
title: Chemin du disque
description: La fonction disk-path fournie affiche l’utilisation du disque avec df -h. Son paramètre path est facultatif et utilise la racine par défaut.
icon: material/harddisk
---

# :material-harddisk: Chemin du disque

`disk-path` affiche l’utilisation du disque pour un chemin avec `df -h`. C’est
l’une des [fonctions fournies](../index.md#custom-functions) utilisables
directement depuis un bouton.

- Exécute : `df -h {{.path}}`
- `path` (facultatif, valeur par défaut `/`) : chemin du système de fichiers à vérifier

!!! example "Le fichier de fonction"

    ```yaml title="functions/disk-path.yaml"
    name: disk-path
    run: "df -h {{.path}}"
    params:
      - name: path
        type: string
        required: false
        default: "/"
        description: Filesystem path to check
    ```

Comme `path` est facultatif et possède une valeur par défaut, un bouton peut
l’omettre.

!!! example "Utiliser le chemin par défaut ou en choisir un"

    ```yaml title="Boutons d’utilisation du disque"
    - name: Disk (root)
      type: button
      function: disk-path        # pas de champ path : utilise "/" par défaut
    - name: Disk (var)
      type: button
      function: disk-path
      path: "/var"
    ```

Le premier bouton exécute `df -h /` et le second `df -h /var`.

## Pages associées

- [Règles](../write-your-own/rules.md) — comportement des valeurs obligatoires et par défaut
- [Fonctions personnalisées](../index.md#custom-functions) — les cinq exemples fournis
- [`command`](../built-in/command.md) — écrire une commande complète
