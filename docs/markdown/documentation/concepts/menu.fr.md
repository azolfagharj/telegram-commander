---
title: Menu
description: Construisez le menu affiché dans Telegram. Mélangez boutons et catégories dans une arborescence et découvrez les déplacements Accueil et Retour.
icon: material/view-list
---

# :material-view-list: Menu

Votre menu est une arborescence d’éléments sous la clé `menu`. Il existe deux
types d’éléments :

- Une **[catégorie](category.md)** ouvre un sous-menu et possède des `items`.
- Un **[bouton](button.md)** exécute une [fonction](function.md).

Si ces termes sont nouveaux, commencez par [Concepts](button.md). Pour tous les
champs, consultez [Configuration → Menu](../configuration.md#menu).

## :material-format-list-bulleted: Un menu simple

Le menu le plus simple est une liste de boutons sans imbrication :

!!! example "Créer un menu avec trois boutons"

    ```yaml title="Trois boutons sans catégories"
    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"

      - name: Free memory
        type: button
        function: command
        command: "free -h"

      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

Envoyez `/start` dans Telegram. Votre menu apparaît.

## :material-folder-outline: Regrouper avec des catégories

Lorsque le menu grandit, regroupez les actions liées dans des catégories. Un
appui affiche leurs `items`. Accueil reste toujours présent. Retour apparaît
uniquement dans une catégorie.

!!! example "Regrouper les boutons dans des catégories"

    ```yaml title="Catégories Système et Services"
    menu:
      - name: System
        type: category
        icon: "💻"
        items:
          - name: Uptime
            type: button
            function: command
            command: "uptime"

          - name: Free memory
            type: button
            function: command
            command: "free -h"

      - name: Services
        type: category
        icon: "🔧"
        items:
          - name: Restart nginx
            type: button
            function: command
            command: "systemctl restart nginx"
            confirm: true
    ```

Les catégories peuvent être imbriquées sans limite. Une catégorie doit contenir
au moins un élément.

## :material-family-tree: Les noms voisins doivent être uniques

Deux éléments sous le même parent ne peuvent pas avoir le même nom, sans tenir
compte de la casse. Le même nom est permis dans des catégories différentes :

!!! example "Réutiliser un nom sous des parents différents"

    ```yaml title="Le même nom sous deux parents"
    menu:
      - name: Web
        type: category
        items:
          - name: Restart          # valide
            type: button
            function: command
            command: "systemctl restart nginx"
      - name: Database
        type: category
        items:
          - name: Restart          # valide, parent différent
            type: button
            function: command
            command: "systemctl restart postgresql"
    ```

## :material-emoticon-outline: Icônes

`icon` est un emoji facultatif affiché avant le nom. Il est uniquement visuel.

!!! warning "Ce bouton redémarre la machine"

    ```yaml title="Un bouton avec une icône emoji"
    - name: Reboot
      type: button
      icon: "🔁"
      function: command
      command: "reboot"
      confirm: true
    ```

!!! tip "Choisissez un emoji simple et courant"

    Certains emoji peuvent tronquer le texte sur certains téléphones. Essayez
    un autre emoji si le bouton semble coupé.

## :material-cellphone-text: Apparence du menu Telegram { #how-the-telegram-menu-looks }

Tous les boutons apparaissent sur le clavier sous la zone de message. Ce clavier
occupe toute la largeur de la conversation afin que le texte reste lisible.

- **Accueil** est toujours le premier bouton de chaque écran. Appuyez dessus
  pour revenir au premier écran.
- **Retour** apparaît dans une catégorie.
- **$ >_ Run Command** apparaît lorsque `enable_run_command` est activé.
- Les éléments sont placés deux par ligne par défaut. Une catégorie peut
  modifier cette disposition avec `columns`. Si un écran contient beaucoup
  d’éléments, **Précédent** et **Suivant** permettent de parcourir les pages.
- Les boutons avec `confirm: true` demandent Oui / Annuler.

!!! info "Les titres du menu sont réutilisés, la sortie reste"

    Un nouveau titre de menu, qu’il s’agisse d’Accueil, d’une catégorie ou
    d’une page, remplace le précédent afin de ne pas remplir la conversation
    d’écrans vides. La ligne **Running** et la sortie restent visibles après
    la réouverture du menu.

!!! info "Les longues sorties arrivent en plusieurs messages"

    La sortie est présentée comme bloc de code. Si elle dépasse un message
    Telegram, elle est divisée en plusieurs messages, chacun répondant au
    précédent. La dernière partie conserve les boutons de la page sur laquelle
    vous vous trouviez, afin que **Retour** quitte toujours la catégorie.
    Consultez
    [Configuration → Quantité de sortie affichée](../configuration.md#how-much-command-output-you-see).

## :material-help-circle-outline: Confirmation { #confirmation }

Ajoutez `confirm: true` pour demander un second appui avant l’exécution.
Utilisez-le pour les actions destructrices. Consultez
[Confirmation](confirmation.md).

!!! warning "Ce bouton arrête un service"

    ```yaml title="Un bouton qui demande d’abord"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

La demande expire après un délai, 5 minutes par défaut. Modifiez-le avec
`confirm_ttl` ; consultez
[Configuration → Champs racine](../configuration.md#root-fields).

## :material-tune-variant: Réglages propres à un bouton

Certains réglages globaux peuvent être remplacés sur un bouton :

!!! example "Donner ses propres réglages à un bouton"

    ```yaml title="Un bouton avec son propre délai, dossier et variables"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"          # peut dépasser le délai global
      workdir: "/var/backups" # exécuter ici
      env:
        BACKUP_MODE: "full"   # variable d’environnement supplémentaire pour cette commande
    ```

Consultez la liste complète dans
[Configuration → Menu](../configuration.md#menu).

## :material-view-grid-outline: Contrôler la disposition

`menu_columns` définit le nombre de boutons par ligne, 2 par défaut. Une
catégorie peut le remplacer avec `columns`. Au-delà de `page_size` éléments,
8 par défaut, le menu est divisé en pages avec Précédent et Suivant. Consultez
[Configuration → Champs racine](../configuration.md#root-fields).

## :material-console: Exécuter une commande

Avec `enable_run_command: true` à la racine, un bouton
**$ >_ Run Command** reste dans le menu, après Retour dans une catégorie ou
après Accueil sur le premier écran. Appuyez dessus puis envoyez la commande
shell à exécuter. Le bot utilise les mêmes shell, délai, dossier de travail et
limites de sortie que les autres boutons.

Accueil ou Retour annule la demande sans rien exécuter. Cette fonction est
désactivée par défaut. Consultez
[Configuration → Champs racine](../configuration.md#root-fields).

!!! warning "Cette option donne accès à toute la machine"

    Toute personne autorisée peut alors exécuter n’importe quelle commande sur
    l’hôte, pas seulement vos boutons. Activez-la uniquement si vous faites
    entièrement confiance à chaque utilisateur autorisé.

## :material-timer-sand: Une commande à la fois

Vos commandes s’exécutent l’une après l’autre.

!!! info "Un second appui attend son tour"

    Si vous appuyez sur un second bouton pendant la première commande, il
    attend puis s’exécute seul. Sa ligne **Running** apparaît au démarrage.
    Cela évite que deux de vos commandes agissent simultanément sur le même
    service ou fichier. Les autres personnes ne sont pas bloquées par votre
    commande : chacune possède sa propre file d’attente.

## :material-function-variant: Action d’un bouton

Chaque bouton référence une **fonction** avec son champ `function`. Les exemples
utilisent la fonction intégrée `command`. Pour les fonctions intégrées,
personnalisées et leur création, consultez [Fonctions](../functions/index.md).

## :material-link-variant: Pages associées

- [Bouton](button.md) — définition d’un bouton
- [Catégorie](category.md) — éléments de sous-menu
- [Configuration → Menu](../configuration.md#menu) — tous les champs
