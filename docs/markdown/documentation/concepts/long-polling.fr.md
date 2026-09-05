---
title: Connexion du bot
description: Le bot établit une connexion sortante vers Telegram, sans ouvrir de port sur votre serveur. Configurez un proxy si Telegram est bloqué sur votre réseau.
icon: material/lan-disconnect
---

# :material-lan-disconnect: Connexion du bot

Le bot établit uniquement des connexions sortantes vers Telegram. Vous n’ouvrez
aucun port sur votre serveur.

!!! tip "Rien à exposer sur Internet"

    Comme la connexion sort uniquement de votre réseau, il n’y a ni panneau web
    à sécuriser, ni règle de pare-feu entrante à ajouter, ni adresse publique à
    découvrir. La machine peut se trouver derrière un routeur domestique ou un
    pare-feu strict et le bot continue de fonctionner.

Si Telegram est bloqué sur votre réseau, définissez un proxy sous `telegram`
dans votre [fichier de configuration](config-file.md).

## Configuration

Pour `api`, `proxy` et `insecure` dans les réglages Telegram, consultez
[Configuration → telegram](../configuration.md#telegram).

## Pages associées

- [Utilisateurs autorisés](allowed-users.md) — qui peut utiliser le bot
- [CLI → run](../cli.md#run) — démarrer le bot au premier plan
- [Exécuter comme service](../installation/run-as-a-service.md) — garder le bot actif
