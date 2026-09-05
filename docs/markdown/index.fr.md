---
title: Telegram Commander
description: Transformez un fichier YAML en bot Telegram qui exécute les commandes de votre serveur Linux d’un simple appui.
icon: material/cellphone-link
hide:
  - navigation
  - toc
---

# :material-cellphone-link: Contrôlez votre serveur Linux depuis Telegram

<div class="hero" markdown>
<div class="hero-art" markdown>
![Telegram Commander](/images/logo-large.png){ .off-glb width="230" }
</div>
<div class="hero-text" markdown>
**Un appui exécute une commande sur votre serveur et renvoie la sortie dans votre conversation.**

Telegram Commander transforme un simple [fichier de configuration](documentation/concepts/config-file.md)
YAML en bot Telegram avec un menu de [boutons](documentation/concepts/button.md).
Placez une commande de terminal derrière chaque bouton et exécutez-la depuis
votre téléphone, sans écrire de code.

<p class="hero-lang">Lire cette page dans une autre langue :
<a href="/" hreflang="en" class="hero-lang__link">English</a>
<a href="/de/" hreflang="de" class="hero-lang__link">Deutsch</a>
<a href="/es/" hreflang="es" class="hero-lang__link">Español</a>
<a href="/ru/" hreflang="ru" class="hero-lang__link">Русский</a>
<a href="/zh/" hreflang="zh" class="hero-lang__link">简体中文</a>
<a href="/fa/" hreflang="fa" class="hero-lang__link">فارسی</a>
</p>
</div>
</div>

<div style="text-align: center" markdown="span">
[Commencer :material-arrow-right:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Voir une configuration :material-file-code-outline:](documentation/configuration.md#a-minimal-config){ .md-button }
</div>

[Installation](documentation/installation/download-and-run.md) ·
[Concepts](documentation/concepts/config-file.md) ·
[Fonctions](documentation/functions/index.md) ·
[Configuration](documentation/configuration.md) ·
[CLI](documentation/cli.md)

## :material-image-multiple-outline:{ .shots } Captures d’écran { .split }

Votre menu, une commande en cours, la sortie reçue et la saisie manuelle d’une
commande. Cliquez sur une image pour l’agrandir.

<div style="text-align: center" markdown="span">
  ![Sortie de commande avec le menu Système ouvert](/images/01.jpeg){ width="140" loading=lazy }
  ![Boutons des ressources et processus](/images/02.jpeg){ width="140" loading=lazy }
  ![Boutons du stockage et des paquets](/images/03.jpeg){ width="140" loading=lazy }
  ![Outils réseau et saisie manuelle d’une commande](/images/04.jpeg){ width="140" loading=lazy }
  ![Sortie de commande renvoyée dans la conversation](/images/05.jpeg){ width="140" loading=lazy }
</div>

## :material-lightning-bolt:{ .bolt } Simple et rapide { .split }

<div class="grid cards cols-3 center-title step-cards" markdown>

-   :material-file-document-outline:{ .middle } __Écrivez votre menu__

    ---

    :material-numeric-1-circle:{ .step } Listez les boutons et leurs commandes.

-   :material-rocket-launch:{ .middle } __Démarrez le bot__

    ---

    :material-numeric-2-circle:{ .step } Lancez-le maintenant ou gardez-le actif comme service.

-   :material-gesture-tap-button:{ .middle } __Appuyez et lisez__

    ---

    :material-numeric-3-circle:{ .step } Appuyez sur un bouton et lisez la
    sortie dans la conversation.

</div>

<div style="text-align: center" markdown="span">
[Commencer maintenant :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
</div>

## :material-view-grid-outline:{ .grid-icon } Exemples d’utilisation { .split }

<div class="grid cards cols-4 icon-left" markdown>

-   :material-restart:{ .lg } Redémarrer ou arrêter un service
-   :material-docker:{ .lg } Démarrer et arrêter des conteneurs
-   :material-package-down:{ .lg } Mettre à jour les paquets
-   :material-text-box-search-outline:{ .lg } Lire les journaux
-   :material-harddisk:{ .lg } Vérifier l’espace disque
-   :material-memory:{ .lg } Surveiller le processeur et la mémoire
-   :material-access-point-network:{ .lg } Interroger des hôtes et tester des URL
-   :material-backup-restore:{ .lg } Sauvegarder et restaurer
-   :material-script-text:{ .lg } Exécuter vos scripts
-   :material-power:{ .lg } Redémarrer ou éteindre l’hôte
-   :material-console:{ .lg } Saisir une commande
-   :material-all-inclusive:{ .lg } Et presque tout le reste

</div>

## :material-thumb-up-outline:{ .thumb } Pourquoi l’utiliser { .split }

<div class="grid cards cols-4 center-title" markdown>

-   :material-clock-fast:{ .lg .middle } __Aucun code__

    ---

    Décrivez le menu et les commandes dans un fichier YAML.

    [:octicons-arrow-right-24: Fichier de configuration](documentation/concepts/config-file.md)

-   :material-cellphone-link:{ .lg .middle } __Depuis partout__

    ---

    Ouvrez Telegram sur votre téléphone et exécutez vos commandes sur le
    serveur, sans VPN vers l’hôte.

    [:octicons-arrow-right-24: Connexion du bot](documentation/concepts/long-polling.md)

-   :material-lan-disconnect:{ .lg .middle } __Aucun port ouvert__

    ---

    Le bot établit uniquement des connexions sortantes vers Telegram. Rien
    n’est exposé à Internet.

    [:octicons-arrow-right-24: Connexion du bot](documentation/concepts/long-polling.md)

-   :material-message-text-outline:{ .lg .middle } __Sortie dans la conversation__

    ---

    Le résultat revient sous forme de message, sans session SSH.

    [:octicons-arrow-right-24: Quantité de sortie affichée](documentation/configuration.md#how-much-command-output-you-see)

-   :material-shield-lock:{ .lg .middle } __Contrôlé et enregistré__

    ---

    Choisissez les utilisateurs, confirmez les actions risquées et enregistrez chaque exécution.

    [:octicons-arrow-right-24: Accès et confirmation](documentation/concepts/allowed-users.md)

-   :material-folder-outline:{ .lg .middle } __Menus imbriqués__

    ---

    Regroupez les boutons. Accueil reste en haut et Retour remonte d’un niveau.

    [:octicons-arrow-right-24: Menu](documentation/concepts/menu.md)

-   :material-function-variant:{ .lg .middle } __Fonctions réutilisables__

    ---

    Écrivez une commande une fois et fournissez des valeurs différentes sur chaque bouton.

    [:octicons-arrow-right-24: Fonctions](documentation/functions/index.md)

-   :material-cog-play-outline:{ .lg .middle } __Toujours actif__

    ---

    Installez-le comme service pour le démarrer avec l’hôte.

    [:octicons-arrow-right-24: Exécuter comme service](documentation/installation/run-as-a-service.md)

</div>

## :material-file-code-outline:{ .code-icon } Un petit exemple { .split }

Cette configuration crée un bouton « Uptime ». Il exécute la commande `uptime`.

!!! example "Cette configuration complète ajoute un bouton"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN" # (1)!
      allowed_users:
        - "YOUR_USER_ID" # (2)!

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime" # (3)!
    ```

    1.  Demandez un jeton à BotFather lors de la création du bot.
    2.  Seuls les comptes listés peuvent ouvrir le menu. Utilisez un identifiant numérique ou un `@username`.
    3.  Placez ici toute commande utilisable dans un terminal.

C’est une configuration complète et fonctionnelle. Tout le reste est facultatif.

!!! tip "Conçu pour un petit groupe de confiance"

    Rien n’écoute les connexions entrantes et seuls les comptes de
    `allowed_users` reçoivent le menu. Chaque utilisateur peut exécuter tous les
    boutons définis : gardez cette liste courte.

## :material-hand-pointing-right: Prêt à essayer ? { .split }

<div style="text-align: center" markdown>
[Commencer :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[Concepts :material-book-open-variant:](documentation/concepts/config-file.md){ .md-button }
[Dernière version :material-download:](https://github.com/azolfagharj/telegram-commander/releases/latest){ .md-button .md-button--primary }

[Parcourir le code source :fontawesome-brands-github:](https://github.com/azolfagharj/telegram-commander){ .md-button }
</div>

Telegram Commander est gratuit et open source. S’il vous fait gagner du temps,
[vous pouvez soutenir son développement](https://azolfagharj.github.io/donate/) :
cela aide à maintenir le projet.
