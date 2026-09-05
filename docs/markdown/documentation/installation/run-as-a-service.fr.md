---
title: Exécuter comme service
description: Gardez le bot actif en arrière-plan avec systemd afin qu’il démarre avec la machine et redémarre après une panne, en tant que root ou utilisateur standard.
icon: material/server
---

# :material-server: Exécuter comme service

Cette page explique comment exécuter Telegram Commander comme service systemd.
Pour un premier lancement pas à pas dans le terminal, consultez
[Exécuter dans la CLI](download-and-run.md).

## :material-cog-outline: systemd (root)

Créez vous-même le fichier d’unité, par exemple
`/etc/systemd/system/telegram-commander.service` :

!!! warning "Cet exemple exécute chaque bouton comme root"

    ```ini title="/etc/systemd/system/telegram-commander.service"
    [Unit]
    Description=telegram-commander Telegram bot
    After=network-online.target
    Wants=network-online.target

    [Service]
    Type=simple
    ExecStart=/path/to/telegram-commander run --config /path/to/configfile.yaml
    Restart=on-failure
    RestartSec=5

    [Install]
    WantedBy=multi-user.target
    ```

    Sans ligne `User=`, le service démarre comme root : toutes les commandes du
    menu disposent de tous les droits sur la machine. Ajoutez `User=someone`
    dans `[Service]` si vos boutons n’ont pas besoin de ces droits.

Remplacez les chemins fictifs, puis :

!!! example "Charger et démarrer le service"

    ```bash title="Activer, démarrer et surveiller le service"
    sudo systemctl daemon-reload
    sudo systemctl enable --now telegram-commander
    sudo systemctl status telegram-commander
    sudo journalctl -u telegram-commander -f
    ```

!!! info "Les modifications de la configuration nécessitent un redémarrage"

    Le service garde le bot actif, mais il lit votre
    [fichier de configuration](../concepts/config-file.md) une seule fois au
    démarrage. Exécutez `sudo systemctl restart telegram-commander` après une
    modification.

## Pages associées

- [Exécuter dans la CLI](download-and-run.md) — premier lancement expliqué
- [Configuration](../configuration.md) — le fichier de configuration
- [CLI](../cli.md) — `run`, `validate` et les autres commandes
