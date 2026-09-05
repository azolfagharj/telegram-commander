---
title: Ping d’un hôte
description: La fonction ping-host fournie interroge un hôte plusieurs fois. Définissez host et le paramètre facultatif count directement sur le bouton.
icon: material/access-point-network
---

# :material-access-point-network: Ping d’un hôte

`ping-host` interroge un hôte plusieurs fois et renvoie le résultat. C’est
l’une des [fonctions fournies](../index.md#custom-functions) utilisables
directement depuis un bouton.

- Exécute : `ping -c {{.count}} {{.host}}`
- `host` (obligatoire) : nom d’hôte ou adresse IP
- `count` (facultatif, valeur par défaut `4`) : nombre de paquets

!!! example "Le fichier de fonction"

    ```yaml title="functions/ping-host.yaml"
    name: ping-host
    run: "ping -c {{.count}} {{.host}}"
    params:
      - name: host
        type: string
        required: true
        description: Hostname or IP
      - name: count
        type: string
        required: false
        default: "4"
        description: Number of ping packets
    ```

## Ajouter un bouton

!!! example "Interroger trois fois un hôte fixe"

    ```yaml title="Bouton pour interroger la passerelle"
    - name: Ping gateway
      type: button
      function: ping-host
      host: "192.168.1.1"
      count: 3
    ```

`count` est numérique et n’a donc pas besoin de guillemets. Vous pouvez aussi
l’omettre pour utiliser la valeur par défaut `4` :

!!! example "Utiliser le nombre par défaut"

    ```yaml title="Bouton pour interroger le DNS"
    - name: Ping DNS
      type: button
      function: ping-host
      host: "1.1.1.1"
    ```

## Pages associées

- [Règles](../write-your-own/rules.md) — comportement des valeurs obligatoires et par défaut
- [Fonctions personnalisées](../index.md#custom-functions) — les cinq exemples fournis
