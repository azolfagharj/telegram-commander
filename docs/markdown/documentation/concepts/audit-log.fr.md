---
title: Journal d’audit
description: Conservez une trace de chaque commande exécutée par le bot, avec la personne, le bouton, le code de sortie et la durée.
icon: material/history
---

# :material-history: Journal d’audit

Un flux de journal distinct qui enregistre chaque commande exécutée : la
personne, le [bouton](button.md), le code de sortie et la durée. Il se configure
sous `logging` dans votre [fichier de configuration](config-file.md).

C’est le registre à consulter pour savoir qui a redémarré un service à trois
heures du matin ou si la sauvegarde de la nuit s’est terminée. Comme ce flux est
séparé, vous pouvez le conserver après la rotation du journal ordinaire.

Chaque exécution ajoute une ligne. Ensemble, elles répondent à quatre questions :

| Question | Ce que la ligne indique |
|---|---|
| Qui ? | Le compte Telegram qui a appuyé sur le bouton |
| Quoi ? | Le bouton utilisé |
| Réussite ? | Le code de sortie renvoyé par la commande |
| Combien de temps ? | La durée d’exécution de la commande |

!!! example "Écrire le journal d’audit dans son propre fichier"

    ```yaml title="section logging de config.yaml"
    logging:
      logs:
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

## Configuration

Pour le schéma `logging` complet et les sorties prises en charge, consultez
[Configuration → logging](../configuration.md#logging).

## Pages associées

- [Fichier de configuration](config-file.md) — emplacement de la journalisation
- [Fonction](function.md) — ce qui s’exécute après l’appui sur un bouton
