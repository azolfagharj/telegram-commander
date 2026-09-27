---
title: Nouvelles tentatives d’envoi
description: Une courte coupure réseau ne fait plus perdre le résultat. Le bot réessaie l’envoi avec une attente croissante, respecte le délai de Telegram en cas de limitation, et garde l’ordre des messages par conversation.
icon: material/refresh
---

# :material-refresh: Nouvelles tentatives d’envoi

Si la connexion à Telegram se coupe un instant pendant que le bot envoie un
résultat, le message n’est pas perdu. Le bot réessaie de lui-même, en
attendant un peu plus longtemps à chaque fois, jusqu’à ce que l’envoi réussisse.

Quand Telegram lui-même demande au bot de ralentir, le bot attend exactement
le délai demandé par Telegram avant de réessayer, au lieu de deviner son
propre délai.

Les messages vers une même conversation arrivent toujours dans l’ordre où
vous les avez déclenchés, même en appuyant sur plusieurs boutons pendant que
la connexion a des problèmes.

Si la coupure dure trop longtemps, le bot finit par abandonner pour ce
résultat. Rien n’apparaît dans la conversation et rien ne reste à moitié
envoyé — la tentative n’est notée que dans le journal serveur, jamais
affichée.

!!! tip "Rien à régler pour une installation normale"

    Les valeurs par défaut absorbent d’elles-mêmes une courte coupure réseau.
    Modifiez la durée pendant laquelle le bot continue d’essayer avec
    `delivery_retry_backoff`, `delivery_retry_backoff_max` et
    `delivery_retry_ttl` dans votre [fichier de configuration](config-file.md).

## Configuration

Pour `delivery_retry_backoff`, `delivery_retry_backoff_max` et
`delivery_retry_ttl`, consultez [Configuration](../configuration.md).

## Pages associées

- [Confirmation](confirmation.md) — un autre moment où le bot attend avant d’agir
- [Connexion du bot](long-polling.md) — la connexion sortante que ceci protège
- [Contrôler la sortie](../output-control.md) — comment un résultat arrive dans votre conversation
