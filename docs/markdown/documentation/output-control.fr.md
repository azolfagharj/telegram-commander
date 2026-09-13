---
title: Contrôler la sortie
description: Choisissez comment les résultats de commande arrivent dans votre conversation. Réglez output, max_output_messages et max_output_bytes à la racine, remplacez output sur un seul bouton et découvrez chaque combinaison racine/bouton avec des exemples.
icon: material/export-variant
---

# :material-export-variant: Contrôler la sortie

Lorsqu’un [bouton](concepts/button.md) a terminé, le bot doit placer le
résultat dans votre conversation. Telegram n’accepte jamais plus de 4096 octets
dans un message : un résultat long arrive donc soit en **plusieurs messages de
réponse**, soit en **un fichier `.txt`** que vous pouvez ouvrir et télécharger.

Trois champs facultatifs décident de ce qui se passe. Vous pouvez tous les
omettre, le bot se comporte quand même raisonnablement.

!!! info "La seule règle à retenir"

    Le bouton gagne. Si un bouton définit `output`, c’est lui qui s’applique.
    Si le bouton ne dit rien, la valeur racine s’applique. Si la racine ne dit
    rien non plus, le bot utilise `auto`.

<div class="grid cards cols-3" markdown>

-   :material-scissors-cutting:{ .middle } __`max_output_bytes`__

    ---

    Quantité de sortie conservée depuis la commande avant tout envoi.

    [:octicons-arrow-right-24: En savoir plus](#max_output_bytes)

-   :material-tune:{ .middle } __`output`__

    ---

    Messages ou fichier. Le seul champ que vous pouvez aussi définir sur un bouton.

    [:octicons-arrow-right-24: En savoir plus](#output)

-   :material-counter:{ .middle } __`max_output_messages`__

    ---

    En mode `auto`, nombre de messages autorisés avant l’envoi d’un fichier.

    [:octicons-arrow-right-24: En savoir plus](#max_output_messages)

</div>

!!! tip "Chaque image de cette page s’agrandit"

    Les captures de conversation sont petites volontairement. Cliquez sur l’une
    d’elles pour la voir en taille réelle.

## :material-card-bulleted-outline: Les trois champs en un coup d’œil { #the-three-fields-at-a-glance }

| Champ | Où le définir | Type | Par défaut | Ce qu’il décide |
|-------|---------------|------|------------|-----------------|
| `max_output_bytes` | racine uniquement | entier | `524288` | Combien d’octets de sortie le bot conserve par commande |
| `output` | racine **et** bouton | `auto` \| `text` \| `file` | `auto` | Si le résultat arrive en messages de conversation ou en un fichier `.txt` |
| `max_output_messages` | racine uniquement | entier, `1`–`10` | `2` | En mode `auto`, le nombre de messages autorisés avant l’envoi d’un fichier |

Les trois se placent dans votre
[fichier de configuration](concepts/config-file.md), à la **racine**, à côté de
`shell` et `timeout`, et non sous `telegram`. La liste complète des clés racine
se trouve sur la page [Configuration](configuration.md#root-fields).

## :material-scissors-cutting: `max_output_bytes`

C’est votre propre limite, et elle s’applique en premier. Pendant qu’une
commande tourne, le bot conserve au plus ce nombre d’octets de sa sortie,
compté séparément pour la sortie normale et la sortie d’erreur. Le reste est
jeté, mais la commande continue jusqu’à la fin ou jusqu’à son `timeout`.

!!! example "Conserver jusqu’à 2 Mo de sortie par commande"

    ```yaml title="config.yaml"
    max_output_bytes: 2097152 # (1)!
    ```

    1.  2 Mo au lieu des 512 Ko par défaut. Compté séparément pour la sortie
        normale et pour la sortie d’erreur.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Lorsque la limite a été atteinte, le résultat l’indique sur sa propre
    ligne, juste sous le résumé. Tout ce qui se trouve au-dessus de
    `(output truncated)` est la partie conservée par le bot ; le reste du
    journal n’a jamais quitté le serveur.

    </div>
    <div class="result-shot" markdown>

    ![Un résultat Telegram qui indique que la sortie est tronquée](/images/output-control/truncated.png){ .shot loading=lazy }

    </div>
    </div>

!!! tip "L’augmenter n’aide que si un fichier est envoyé"

    Un message de conversation ne peut jamais transporter plus de 4096 octets :
    un `max_output_bytes` plus grand ne vous parvient donc en entier que
    lorsque le résultat arrive en fichier `.txt`. Associez-le à `output: auto`
    (le mode par défaut) ou à `output: file`.

## :material-tune: `output`

`output` choisit le mode de livraison. Il accepte trois valeurs et c’est le
seul champ de sortie que vous pouvez aussi placer sur un bouton.

=== "auto"

    Des messages texte tant que le résultat est court, un fichier `.txt`
    sinon. C’est le mode par défaut, celui que vous gardez si vous n’écrivez
    rien.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Laisser le bot choisir pour chaque résultat"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: auto # (1)!
        max_output_messages: 2 # (2)!
        ```

        1.  Décider pour chaque résultat au lieu d’imposer une seule livraison.
        2.  Le point de bascule. Deux messages ou moins restent dans la
            conversation ; au-delà, le résultat devient un seul fichier.

    </div>
    <div class="result-shot" markdown>

    ![Le mode auto garde les sorties courtes en texte et envoie les plus longues en fichier](/images/output-control/auto.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

=== "text"

    Toujours des messages de conversation, quelle que soit la longueur du
    résultat. Le bot s’arrête après dix messages et ajoute une note indiquant
    que la suite a été coupée.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Tout garder dans la conversation"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: text # (1)!
        ```

        1.  `max_output_messages` est ignoré ici, car le bot ne passe jamais au
            fichier de lui-même.

    </div>
    <div class="result-shot" markdown>

    ![Un résultat de commande livré en messages Telegram](/images/output-control/message.png){ .shot loading=lazy }

    </div>
    </div>

=== "file"

    Toujours un seul fichier `.txt`, même pour un résultat d’une ligne. Le
    résumé accompagne le fichier en légende.

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "Envoyer chaque résultat en pièce jointe"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: file # (1)!
        ```

        1.  Chaque bouton envoie désormais une pièce jointe, même les plus
            rapides. Placez plutôt ce réglage sur des boutons précis si c’est
            trop.

    </div>
    <div class="result-shot" markdown>

    ![Un résultat de commande livré en fichier Telegram](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

!!! note "L’omettre revient à écrire `auto`"

    Vous n’avez jamais besoin d’écrire `output`. Une configuration qui
    fonctionnait avant l’existence de ce champ se comporte exactement comme
    avant.

## :material-counter: `max_output_messages`

Ce champ ne compte qu’en mode `auto`, et il n’existe qu’à la racine. C’est le
nombre de messages que le bot accepte d’envoyer avant de renoncer aux messages
et d’envoyer un seul fichier. Les valeurs permises vont de `1` à `10` ; la
valeur par défaut est `2`.

!!! example "Envoyer un fichier dès qu’un second message est nécessaire"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 1 # (1)!
    ```

    1.  Le réglage le plus strict. Seul un résultat qui tient dans un message
        reste dans la conversation.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un résultat qui tient dans la limite reste en messages de conversation. Dès
    qu’il faudrait un message de plus, tout le résultat arrive en un seul
    fichier `.txt`, et aucun message partiel n’est envoyé.

    </div>
    <div class="result-shot" markdown>

    ![Un résultat court en message à côté d’un résultat plus long en fichier](/images/output-control/allowance.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

!!! warning "C’est un champ réservé à la racine"

    Un bouton ne peut pas définir `max_output_messages`. Placez-le à la racine
    et laissez les boutons qui ont besoin d’autre chose définir `output: text`
    ou `output: file`.

## :material-table-arrow-right: La racine et le bouton ensemble { #root-and-button-together }

Un bouton peut porter son propre `output`. Rien d’autre à propos de la sortie
ne se règle bouton par bouton. Le mode effectif est simplement :

!!! abstract "Comment le mode effectif est choisi"

    ```text title="Ordre de priorité"
    button output  →  root output  →  auto
    ```

| `output` racine | `output` du bouton | Ce qui se passe réellement |
|-----------------|--------------------|----------------------------|
| non défini ou `auto` | non défini | `auto` |
| non défini ou `auto` | `auto` | `auto` |
| non défini ou `auto` | `text` | `text` |
| non défini ou `auto` | `file` | `file` |
| `text` | non défini | `text` |
| `text` | `auto` | `auto` |
| `text` | `text` | `text` |
| `text` | `file` | `file` |
| `file` | non défini | `file` |
| `file` | `auto` | `auto` |
| `file` | `text` | `text` |
| `file` | `file` | `file` |

Les sections ci-dessous parcourent les douze combinaisons, avec un exemple
chacune.

## :material-numeric-1-box-outline: Quand la racine est `auto` { #when-the-root-is-auto }

Cela comprend le cas où la racine ne dit rien du tout, car un `output` absent
signifie `auto`.

### La racine est `auto`, le bouton ne dit rien { #the-root-is-auto-the-button-says-nothing }

Le cas de tous les jours. Les résultats courts restent dans la conversation ;
un résultat qui demanderait plus de `max_output_messages` messages arrive en
fichier.

!!! example "Un bouton simple sous une racine simple"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 2

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    `df -h` affiche quelques lignes : le résultat tient dans un message et
    reste dans la conversation sous forme de bloc de code, avec le résumé
    au-dessus.

    </div>
    <div class="result-shot" markdown>

    ![L’utilisation du disque renvoyée en un seul message Telegram](/images/output-control/disk-usage.png){ .shot loading=lazy }

    </div>
    </div>

### La racine est `auto`, le bouton dit `auto` { #the-root-is-auto-the-button-says-auto }

Écrire `auto` sur le bouton ne change rien ici. Cela ne vaut la peine que si
vous voulez que ce bouton reste en `auto`, quoi que devienne la racine plus
tard.

!!! example "Écrire le mode par défaut sur le bouton"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
        output: auto # (1)!
    ```

    1.  Même comportement qu’en l’omettant aujourd’hui, mais ce bouton le
        garde même si la racine passe à `text` ou `file` plus tard.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Exactement le comportement automatique : du texte tant que le résultat est
    court, un fichier dès qu’il dépasse la limite autorisée. Sur un serveur
    avec de nombreux disques montés, ce même bouton peut vous envoyer un
    fichier.

    </div>
    <div class="result-shot" markdown>

    ![Le même bouton en message et en fichier](/images/output-control/disk-usage-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### La racine est `auto`, le bouton dit `text` { #the-root-is-auto-the-button-says-text }

Utilisez ceci pour un bouton dont vous voulez toujours lire la sortie sur
place, même lorsqu’elle est un peu longue : une liste d’état que vous faites
défiler plutôt que de télécharger.

!!! example "Garder un bouton dans la conversation"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Service status
        type: button
        function: command
        command: "systemctl status nginx --no-pager"
        output: text # (1)!
    ```

    1.  Jamais de fichier, quoi que dise `max_output_messages`.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Le résultat est divisé en plusieurs messages, chacun répondant au
    précédent, afin que l’ordre ne soit jamais perdu. Les boutons restent sur
    la dernière partie.

    </div>
    <div class="result-shot" markdown>

    ![Un long état de service divisé en trois messages Telegram](/images/output-control/service-status-split.png){ .shot loading=lazy }

    </div>
    </div>

### La racine est `auto`, le bouton dit `file` { #the-root-is-auto-the-button-says-file }

Le choix habituel pour les journaux et les sauvegardes : vous ne voulez presque
jamais les lire dans la conversation, et un fichier est plus facile à
conserver.

!!! example "Toujours télécharger celui-ci"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Nginx logs
        type: button
        function: command
        command: "journalctl -u nginx -n 2000 --no-pager"
        output: file # (1)!
    ```

    1.  Le reste du menu continue de se comporter automatiquement. Seul ce
        bouton est fixé au fichier.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Une pièce jointe dont la légende indique le nom du bouton, le code de
    sortie et la durée. Appuyez dessus pour lire tout le journal, ou gardez-le
    pour plus tard.

    </div>
    <div class="result-shot" markdown>

    ![Les journaux Nginx livrés en fichier Telegram](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-2-box-outline: Quand la racine est `output: text` { #when-the-root-is-output-text }

Chaque bouton reste dans la conversation, sauf s’il dit le contraire. Choisissez
ce mode si vous n’aimez pas télécharger des fichiers sur votre téléphone.

### La racine est `text`, le bouton ne dit rien { #the-root-is-text-the-button-says-nothing }

!!! example "Des messages de conversation partout"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un message court, car `uptime` affiche une seule ligne. Les résultats plus
    longs sont simplement divisés en davantage de messages, jusqu’à dix.

    </div>
    <div class="result-shot" markdown>

    ![Uptime renvoyé en un seul message Telegram](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### La racine est `text`, le bouton dit `auto` { #the-root-is-text-the-button-says-auto }

Le bouton retrouve la bascule automatique : ce seul bouton peut donc envoyer un
fichier alors que le reste du menu ne le fait jamais. Le
`max_output_messages` de la racine décide du point de bascule.

!!! example "Un bouton autorisé à envoyer un fichier quand il devient long"

    ```yaml title="config.yaml"
    output: text
    max_output_messages: 3 # (1)!

    menu:
      - name: Package list
        type: button
        function: command
        command: "dpkg -l"
        output: auto
    ```

    1.  Seul ce bouton lit cette limite, car c’est le seul en mode `auto`.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Une liste de paquets courte arrive en trois messages au maximum. Sur un
    serveur bien rempli, le même bouton dépasse la limite et envoie un seul
    fichier.

    </div>
    <div class="result-shot" markdown>

    ![Une liste de paquets en messages et en fichier](/images/output-control/package-list-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### La racine est `text`, le bouton dit `text` { #the-root-is-text-the-button-says-text }

Identique à ne rien dire, mais l’intention est écrite. Utile dans une longue
configuration où vous ne voulez pas qu’une modification ultérieure à la racine
change ce bouton.

!!! example "Fixer un bouton aux messages de conversation"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Last logins
        type: button
        function: command
        command: "last -n 20"
        output: text
    ```

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Toujours des messages de conversation, jusqu’à dix, quoi que devienne la
    racine plus tard.

    </div>
    <div class="result-shot" markdown>

    ![Les connexions récentes renvoyées en message Telegram](/images/output-control/last-logins.png){ .shot loading=lazy }

    </div>
    </div>

### La racine est `text`, le bouton dit `file` { #the-root-is-text-the-button-says-file }

!!! example "Un bouton téléchargeable dans un menu tout en messages"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Full system report
        type: button
        function: command
        command: "/usr/local/bin/report.sh"
        output: file
        timeout: "5m" # (1)!
    ```

    1.  Un rapport peut prendre du temps : ce bouton reçoit donc un délai plus
        long que le délai global.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    La seule pièce jointe d’un menu par ailleurs entièrement en messages. La
    légende indique la durée du rapport, ce qui est utile lorsqu’il tourne
    plusieurs minutes.

    </div>
    <div class="result-shot" markdown>

    ![Un rapport système complet livré en fichier Telegram](/images/output-control/system-report-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-3-box-outline: Quand la racine est `output: file` { #when-the-root-is-output-file }

Chaque résultat est un fichier, même d’une seule ligne. Cela convient à un bot
surtout utilisé pour des rapports et des archives.

### La racine est `file`, le bouton ne dit rien { #the-root-is-file-the-button-says-nothing }

!!! example "Des fichiers partout"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Même une seule ligne de sortie arrive en pièce jointe. La légende porte
    toujours le résumé : vous pouvez lire le code de sortie sans ouvrir le
    fichier.

    </div>
    <div class="result-shot" markdown>

    ![Uptime livré en petit fichier Telegram](/images/output-control/uptime-file.png){ .shot loading=lazy }

    </div>
    </div>

!!! warning "Même les résultats minuscules deviennent des téléchargements"

    Une réponse d’une ligne arrive en pièce jointe qu’il faut ouvrir. Si cela
    vous gêne, laissez la racine sur `auto` et mettez `output: file` seulement
    sur les boutons qui en ont besoin.

### La racine est `file`, le bouton dit `auto` { #the-root-is-file-the-button-says-auto }

Le bouton échappe à la règle du tout-fichier et se comporte de nouveau
normalement.

!!! example "Laisser une vérification rapide lisible"

    ```yaml title="config.yaml"
    output: file
    max_output_messages: 2

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
        output: auto # (1)!
    ```

    1.  Retour à la bascule automatique, pour ce bouton uniquement.

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    La réponse d’une ligne est courte : elle reste dans la conversation en
    message, alors que tous les autres boutons du menu envoient encore des
    fichiers.

    </div>
    <div class="result-shot" markdown>

    ![Uptime de retour dans la conversation en message](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### La racine est `file`, le bouton dit `text` { #the-root-is-file-the-button-says-text }

!!! example "Forcer un bouton à revenir dans la conversation"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Who is logged in
        type: button
        function: command
        command: "who"
        output: text
    ```

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Un message, jamais une pièce jointe, quelle que soit la longueur de la
    liste des utilisateurs connectés. Au-delà de dix messages, le bot signale
    que la suite a été coupée.

    </div>
    <div class="result-shot" markdown>

    ![Les utilisateurs connectés renvoyés en message Telegram](/images/output-control/who.png){ .shot loading=lazy }

    </div>
    </div>

### La racine est `file`, le bouton dit `file` { #the-root-is-file-the-button-says-file }

Répéter la valeur racine. C’est sans danger et cela garde le bouton correct si
vous assouplissez la racine plus tard.

!!! example "Un bouton qui doit toujours être un fichier"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Database dump log
        type: button
        function: command
        command: "cat /var/log/pg_dump.log"
        output: file
    ```

!!! success "Ce que vous voyez dans la conversation"

    <div class="result" markdown>
    <div class="result-text" markdown>

    Une pièce jointe, exactement ce que la racine demandait déjà. L’écrire sur
    le bouton fait survivre l’intention à un changement à la racine.

    </div>
    <div class="result-shot" markdown>

    ![Un journal de sauvegarde de base de données livré en fichier Telegram](/images/output-control/db-dump-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-file-document-outline: À quoi ressemble le fichier { #what-the-file-looks-like }

Lorsqu’un résultat est livré en fichier, vous recevez une pièce jointe et une
courte légende.

- **Nom** — le nom du bouton en minuscules avec des tirets, puis la date et
  l’heure en UTC, par exemple `nginx-logs-20260913-091204.txt`.
- **Légende** — le même résumé que celui affiché en haut d’un message de
  conversation : nom du bouton, code de sortie, durée, et la note de délai ou
  de troncature lorsqu’il y en a une.
- **Corps** — ce résumé de nouveau, puis une section `--- stdout ---`, puis une
  section `--- stderr ---` lorsque la commande y a écrit quelque chose.

!!! success "Dans le fichier téléchargé"

    ```text title="nginx-logs-20260913-091204.txt"
    Button: Nginx logs
    Exit: 0
    Duration: 316ms

    --- stdout ---
    2026-09-13T09:11:58 nginx: worker process started
    ...

    --- stderr ---
    2026-09-13T09:12:01 nginx: warning: duplicate server name
    ```

## :material-alert-outline: Limites que vous ne pouvez pas changer { #limits-you-cannot-change }

!!! warning "C’est Telegram qui fixe le plafond, pas le bot"

    - Un message de conversation contient au plus 4096 octets.
    - En mode `text`, le bot envoie au plus dix messages et termine par
      `(output too long; showing first N bytes)`.
    - Quel que soit le `max_output_messages` écrit, dix messages restent le
      plafond absolu.

!!! note "Si l’envoi échoue, vous recevez quand même le résultat"

    Lorsque le fichier ne peut pas être envoyé, faute de réseau ou parce que
    Telegram le refuse, le bot revient aux messages de conversation afin que le
    résultat ne soit pas perdu.

## :material-console: Exécuter une commande suit le réglage racine { #run-command-uses-the-root-setting }

Le bouton **`$ >_ Run Command`** ne fait pas partie de votre menu : il n’a donc
pas de `output` propre. Il suit toujours les `output` et
`max_output_messages` de la racine. Consultez
[Menu → Exécuter une commande](concepts/menu.md#run-command).

## :material-link-variant: Pages associées { #related }

<div class="grid cards cols-2" markdown>

-   :material-cog-outline:{ .middle } __Configuration__

    ---

    Toutes les clés racine, dont les trois champs de sortie dans un tableau.

    [:octicons-arrow-right-24: Configuration](configuration.md#root-fields)

-   :material-gesture-tap-button:{ .middle } __Bouton__

    ---

    Ce qu’un bouton peut remplacer, dont `output`.

    [:octicons-arrow-right-24: Bouton](concepts/button.md)

-   :material-view-list:{ .middle } __Menu__

    ---

    Comment les résultats côtoient le menu, et ce que fait Exécuter une commande.

    [:octicons-arrow-right-24: Menu](concepts/menu.md)

-   :material-timer-outline:{ .middle } __Shell__

    ---

    Le shell, le dossier de travail et le délai utilisés par vos commandes.

    [:octicons-arrow-right-24: Shell](concepts/shell.md)

</div>
