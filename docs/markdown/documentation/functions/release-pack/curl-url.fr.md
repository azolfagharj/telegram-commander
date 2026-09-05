---
title: URL Curl
description: La fonction curl-url fournie récupère une URL avec curl. Ajoutez l’URL directement au bouton avec la clé url.
icon: material/web
---

# :material-web: URL Curl

`curl-url` récupère une URL avec `curl`. Elle échoue en cas d’erreur HTTP
(`-f`) et abandonne après 30 secondes. C’est l’une des
[fonctions fournies](../index.md#custom-functions) utilisables directement
depuis un bouton.

- Exécute : `curl -fsSL --max-time 30 {{.url}}`
- `url` (obligatoire) : l’URL à demander

!!! example "Le fichier de fonction"

    ```yaml title="functions/curl-url.yaml"
    name: curl-url
    run: "curl -fsSL --max-time 30 {{.url}}"
    params:
      - name: url
        type: string
        required: true
        description: URL to request
    ```

## Ajouter un bouton

Écrivez `url` directement sur le bouton. Son nom correspond au paramètre
déclaré par la fonction.

!!! example "Vérifier un point d’accès"

    ```yaml title="Bouton de vérification de l’API"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

Le bouton exécute `curl -fsSL --max-time 30 https://example.com/health`.
[`validate`](../../cli.md#validate) signale une erreur si `url` manque.

## Pages associées

- [Paramètres](../../concepts/parameter.md) — vérification des valeurs de bouton
- [Fonctions personnalisées](../index.md#custom-functions) — les cinq exemples fournis
