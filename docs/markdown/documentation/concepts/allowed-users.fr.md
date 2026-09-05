---
title: Utilisateurs autorisés
description: Choisissez les comptes Telegram qui peuvent ouvrir le menu du bot. Ajoutez-les avec leur identifiant numérique ou leur nom d’utilisateur.
icon: material/shield-lock
---

# :material-shield-lock: Utilisateurs autorisés

La liste des comptes Telegram autorisés à utiliser le bot, avec leur `user_id`
numérique ou leur `@username`. Tous les autres sont refusés.

Les utilisateurs non autorisés reçoivent un message contenant leur `user_id` et
leur `username` afin de demander l’accès à un administrateur. C’est également
ainsi que vous trouvez votre propre identifiant la première fois — consultez
[Exécuter dans la CLI → Étape 5](../installation/download-and-run.md#step-5-find-your-user-id-if-needed).

Ce message permet aussi d’ajouter une personne. Demandez-lui d’envoyer un
message au bot, de copier l’identifiant reçu, ajoutez cet identifiant à
`allowed_users`, puis redémarrez le bot. À son prochain message, elle recevra le
menu au lieu du refus.

!!! warning "Toutes les personnes de la liste partagent le même menu"

    Il n’existe pas d’autorisations par utilisateur. Toute personne ajoutée peut
    utiliser tous les boutons définis. Limitez donc la liste aux personnes
    auxquelles vous confiez la machine elle-même.

!!! example "Deux comptes autorisés à utiliser le bot"

    ```yaml title="section telegram de config.yaml"
    telegram:
      allowed_users:
        - "123456789"
        - "@alice"
    ```

## Configuration

Pour `allowed_users` et les autres réglages Telegram, consultez
[Configuration → telegram](../configuration.md#telegram).

## Pages associées

- [Fichier de configuration](config-file.md) — contient la section `telegram`
- [Connexion du bot](long-polling.md) — aucun port entrant à ouvrir
