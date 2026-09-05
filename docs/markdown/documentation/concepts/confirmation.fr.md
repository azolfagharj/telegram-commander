---
title: Confirmation
description: Demandez une confirmation avant l’exécution d’un bouton. Activez confirm pour les actions irréversibles et réglez la durée de validité du choix Oui ou Annuler.
icon: material/help-circle-outline
---

# :material-help-circle-outline: Confirmation { #confirmation }

Une étape facultative « Êtes-vous sûr ? » avant l’exécution d’un
[bouton](button.md), activée avec `confirm: true`. Elle est utile pour les
actions destructrices.

Le bot demande (`Confirm: … ?`) avec Oui / Annuler, ainsi qu’Accueil (et Retour
dans une catégorie). Si Exécuter une commande est activé, ce bouton reste aussi
sur l’écran de confirmation. Rien ne s’exécute avant d’appuyer sur Oui ;
Annuler, Accueil et Retour laissent tous la commande intacte.

La demande expire après un délai (5 minutes par défaut). Modifiez-le avec
`confirm_ttl` dans votre [fichier de configuration](config-file.md). Après
expiration, appuyez de nouveau sur le bouton pour obtenir une nouvelle demande.

!!! warning "Utilisez-la pour toute action irréversible"

    Sur un téléphone, une touche du menu n’est qu’à un appui, et toutes les
    personnes autorisées voient le même menu. Ajoutez `confirm: true` aux
    boutons qui arrêtent un service, suppriment des données ou redémarrent la
    machine.

!!! example "Un bouton qui demande avant de s’exécuter"

    ```yaml title="Bouton d’arrêt de nginx"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## Configuration

Pour `confirm` sur les boutons et le réglage global `confirm_ttl`, consultez
[Configuration](../configuration.md).

## Pages associées

- [Bouton](button.md) — l’élément qui peut demander une confirmation
- [Menu → Confirmation](menu.md#confirmation) — exemples et comportement
