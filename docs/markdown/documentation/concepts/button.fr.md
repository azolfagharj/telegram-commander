---
title: Bouton
description: Un élément de menu qui exécute une fonction. Écrivez chaque valeur directement sur le bouton avec le nom du paramètre déclaré par la fonction.
icon: material/gesture-tap-button
---

# :material-gesture-tap-button: Bouton

Un élément du menu Telegram sur lequel vous pouvez appuyer. Un bouton possède
un `name` et exécute une [fonction](function.md) sur votre serveur, puis renvoie
le résultat dans la conversation.

Les boutons et les [catégories](category.md) forment l’arborescence du
[menu](menu.md) sous la clé `menu` de votre
[fichier de configuration](config-file.md). Un bouton effectue le travail ;
une catégorie ouvre seulement un sous-menu.

## Apparence d’un bouton

Les boutons apparaissent sur le clavier sous la zone de message, deux par ligne
par défaut. Le nom choisi est le texte de la touche : gardez-le assez court
pour un téléphone. Consultez
[Menu → Apparence du menu Telegram](menu.md#how-the-telegram-menu-looks).

## :material-format-list-checks: Éléments d’un bouton

!!! example "Chaque ligne contrôle une partie du bouton"

    ```yaml title="Un bouton entièrement décrit"
    - name: Restart nginx # (1)!
      type: button # (2)!
      icon: "🔄" # (3)!
      function: command # (4)!
      command: "systemctl restart nginx" # (5)!
      confirm: true # (6)!
    ```

    1.  Texte de la touche dans Telegram, unique parmi les éléments voisins.
    2.  Toujours `button`. Utilisez `category` pour un sous-menu.
    3.  Emoji facultatif affiché avant le nom, sans effet sur l’exécution.
    4.  [Fonction](function.md) utilisée. `command` exécute une commande shell.
    5.  Ce que la fonction `command` exécute.
    6.  Facultatif. Demande « Êtes-vous sûr ? » avant l’exécution. Omettez ce
        champ pour les actions qui ne font que lire des informations.

## Que se passe-t-il après un appui ?

1. Le bot publie une courte ligne **Running** pour indiquer le démarrage.
2. La commande s’exécute sur la machine où tourne le bot.
3. La sortie revient sous forme de bloc de code, avec le code de sortie et la
   durée. Une longue sortie arrive en plusieurs messages.
4. Vous restez dans le même menu ; **Retour** quitte donc toujours la catégorie.

## :material-code-braces: Boutons courants

=== "Vérifier quelque chose"

    ```yaml title="Bouton Uptime"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

=== "Gérer un service"

    ```yaml title="Bouton de redémarrage de nginx"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

=== "Lire un journal"

    ```yaml title="Bouton du journal Nginx"
    - name: Nginx log
      type: button
      function: command
      command: "journalctl -u nginx -n 50 --no-pager"
    ```

=== "Exécuter un script"

    ```yaml title="Bouton de sauvegarde nocturne"
    - name: Nightly backup
      type: button
      function: script
      path: "/usr/local/bin/backup.sh"
    ```

=== "Action destructive"

    ```yaml title="Bouton d’arrêt de nginx"
    - name: Stop nginx
      type: button
      icon: "🛑"
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## :material-emoticon-outline: Icônes

`icon` place un emoji devant le nom. Ce choix est uniquement visuel et peut
être modifié sans toucher à l’action du bouton.

!!! example "Une icône modifie seulement le libellé"

    ```yaml title="Le même bouton avec et sans icône"
    - name: Disk usage
      type: button
      function: command
      command: "df -h"

    - name: Disk usage
      type: button
      icon: "💾"
      function: command
      command: "df -h"
    ```

## :material-help-circle-outline: Demander avant une action risquée

Ajoutez `confirm: true` pour que le bot propose Oui ou Annuler. Utilisez-le
pour arrêter un service, supprimer des données ou redémarrer la machine. La
demande expire après cinq minutes par défaut.

Consultez [Confirmation](confirmation.md) pour les cas d’utilisation et le
réglage du délai.

## Réglages propres à un bouton

La plupart des réglages globaux peuvent être remplacés sur un bouton :

!!! example "Remplacer les réglages d’une tâche lente"

    ```yaml title="Une tâche lente exécutée ailleurs"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"
      workdir: "/var/backups"
      env:
        BACKUP_MODE: "full"
    ```

`timeout` accorde plus de temps à cette commande, `workdir` choisit son dossier
et `env` ajoute ses variables d’environnement.

## Valeurs de la fonction

Écrivez les valeurs directement sur le bouton. `command`, `path` et `args` sont
des raccourcis pour les paramètres de même nom. Les noms personnalisés comme
`url`, `host`, `unit` et `lines` fonctionnent de la même manière.

!!! example "Transmettre des valeurs personnalisées"

    ```yaml title="Journaux Nginx récents"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

Ne placez pas ces valeurs dans `params:`. Chaque clé doit correspondre à un
paramètre déclaré par la fonction. [`validate`](../cli.md#validate) signale les
noms inconnus, les valeurs obligatoires manquantes et les valeurs `int` ou
`bool` non valides.

## Configuration

Pour tous les champs acceptés, consultez
[Configuration → Menu](../configuration.md#menu).

## Pages associées

<div class="grid cards cols-2" markdown>

-   :material-folder-outline:{ .middle } __Catégorie__

    ---

    Ouvre un sous-menu au lieu d’exécuter une action.

    [:octicons-arrow-right-24: Catégorie](category.md)

-   :material-function:{ .middle } __Fonction__

    ---

    Ce qui s’exécute après un appui.

    [:octicons-arrow-right-24: Fonction](function.md)

-   :material-tune:{ .middle } __Paramètre__

    ---

    Une valeur nécessaire à une fonction.

    [:octicons-arrow-right-24: Paramètre](parameter.md)

-   :material-view-list:{ .middle } __Menu__

    ---

    Organisez toute l’arborescence.

    [:octicons-arrow-right-24: Menu](menu.md)

</div>
