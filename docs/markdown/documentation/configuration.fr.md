---
icon: material/file-cog-outline
title: Configuration
description: Tous les réglages de Telegram Commander avec leur type, valeur par défaut et signification : telegram, menu, fonctions, délais, sorties et journaux.
---

# :material-file-cog-outline: Configuration

Le [fichier de configuration](concepts/config-file.md) décrit tout le bot : la
connexion Telegram, les personnes autorisées, le menu de
[boutons](concepts/button.md) et les journaux. Transmettez-le avec `--config`
aux commandes `run`, `validate`, `fmt` et `list-functions` (voir [CLI](cli.md)).

Toutes les clés utilisent `lower_snake_case`. **Les clés inconnues sont
refusées** : une faute est immédiatement signalée par
[la validation](cli.md#validate).

**Obligatoire** signifie que la validation échoue si le champ manque ou reste
vide après application des valeurs par défaut.  
Les champs **facultatifs** peuvent être omis ; la colonne Par défaut indique la
valeur alors utilisée.

Si vous débutez, commencez par
[Exécuter dans la CLI](installation/download-and-run.md), qui crée une première
configuration. Consultez [Concepts](concepts/config-file.md) pour le vocabulaire.

## :material-rocket-launch-outline: Configuration minimale { #a-minimal-config }

Seuls `telegram`, avec un jeton et un
[utilisateur autorisé](concepts/allowed-users.md), et `menu` sont obligatoires.
Tous les autres champs ont une valeur par défaut :

!!! example "Commencer avec un utilisateur et un bouton"

    ```yaml title="config.yaml (minimal)"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

Le dossier `config-examples/` de la version contient un exemple minimal et un complet.

## :material-card-bulleted-outline: Champs racine { #root-fields }

| Champ | Type | Obligatoire | Par défaut | Description |
|-------|------|-------------|------------|-------------|
| `telegram` | objet | oui | — | Réglages Telegram (voir ci-dessous) |
| `menu` | liste | oui | — | Arborescence du menu, au moins un élément |
| `function_directory` | chaîne | non | non défini | Dossier YAML des fonctions personnalisées (voir les règles ci-dessous) |
| `shell` | chaîne | non | `/bin/bash` | [Shell](concepts/shell.md) utilisé comme `shell -c "<command>"` |
| `timeout` | durée | non | `60s` | Délai maximal par défaut d’une commande |
| `max_output_bytes` | entier | non | `524288` | Sortie maximale conservée par commande (voir [Quantité de sortie affichée](#how-much-command-output-you-see)) |
| `workdir` | chaîne | non | dossier du processus | Dossier de travail par défaut |
| `env` | objet | non | vide | Variables d’environnement supplémentaires |
| `menu_columns` | entier | non | `2` | Boutons d’élément par ligne sous la zone de message |
| `page_size` | entier | non | `8` | Éléments par page avant pagination |
| `confirm_ttl` | durée | non | `5m` | Durée de validité d’une [confirmation](concepts/confirmation.md) |
| `enable_run_command` | booléen | non | `false` | Affiche **`$ >_ Run Command`** pour exécuter le message suivant comme commande shell. Cette option est désactivée par défaut. Toute personne autorisée peut alors exécuter n’importe quelle commande sur l’hôte : activez-la uniquement si vous faites confiance à tous les utilisateurs autorisés. Cette clé est invalide sous `telegram`. |
| `logging` | objet | non | journal par défaut | Journaux nommés (voir ci-dessous) |

??? note "Que se passe-t-il si `shell` est omis ?"

    Le bot utilise `/bin/bash`. Les autres champs racine facultatifs utilisent
    également leur valeur par défaut. Définissez-les uniquement pour utiliser
    une autre valeur, par exemple `shell: /bin/sh`.

### Quantité de sortie affichée { #how-much-command-output-you-see }

Deux limites s’appliquent successivement. `max_output_bytes` est **votre**
limite, en plus d’une limite Telegram non modifiable.

**1. Votre limite : `max_output_bytes`** (`524288` par défaut, soit 512 Ko)

Pendant l’exécution, le bot conserve au maximum cette quantité pour la sortie
normale et la sortie d’erreur, séparément. Le surplus est abandonné, mais la
commande continue jusqu’à sa fin ou son `timeout`. Le résultat commence alors
par `(output truncated)`.

**2. Limite Telegram : un message contient au maximum 4096 octets**

Si le résultat dépasse un message, le bot le divise en plusieurs messages liés
et ordonnés. Les boutons apparaissent sur la dernière partie. La séparation se
fait si possible entre les lignes.

Pour une sortie encore plus longue, le bot s’arrête après 10 messages et le
dernier se termine par une note comme
`(output too long; showing first N bytes)`, où `N` indique la quantité de sortie
que vous avez réellement reçue.

Augmenter `max_output_bytes` conserve plus de sortie, mais vous en voyez
toujours environ dix messages au maximum. Réduisez plutôt la sortie de la
commande, par exemple `journalctl -u nginx | tail -n 50`, ou écrivez-la dans
un fichier sur le serveur.

### Règles de `function_directory` { #function_directory-rules }

| Situation | Résultat |
|-----------|----------|
| Clé absente | Journal d’information ; fonctions intégrées uniquement |
| Clé présente mais vide (`""`) | Journal d’information ; fonctions intégrées uniquement |
| Chemin absent ou inaccessible | Erreur bloquante ; arrêt du processus |
| Dossier existant mais vide | Valide |

!!! warning "Un chemin incorrect arrête le bot"

    Si `function_directory` indique un dossier absent ou illisible, le
    programme s’arrête au lieu de démarrer sans vos fonctions personnalisées.

## :material-send-circle-outline: `telegram` { #telegram }

| Champ | Type | Obligatoire | Par défaut | Description |
|-------|------|-------------|------------|-------------|
| `bot_token` | chaîne | oui | — | Jeton du bot fourni par BotFather |
| `allowed_users` | liste de chaînes | oui | — | [Utilisateurs autorisés](concepts/allowed-users.md) |
| `api` | chaîne | non | `https://api.telegram.org` | URL de base de l’API du bot |
| `proxy.enabled` | booléen | non | `false` | Utiliser un proxy pour Telegram |
| `proxy.url` | chaîne | conditionnel | — | Obligatoire si `proxy.enabled` vaut `true` |
| `insecure` | booléen | non | `false` | Ignorer la vérification TLS (déconseillé) |

Les utilisateurs non autorisés reçoivent leur `user_id` et leur `username` afin
de demander l’accès. C’est aussi ainsi que vous trouvez votre propre identifiant
la première fois ; consultez
[Exécuter dans la CLI → Étape 5](installation/download-and-run.md#step-5-find-your-user-id-if-needed).

!!! example "Se connecter avec un proxy"

    ```yaml title="section telegram avec un proxy"
    telegram:
      bot_token: "123456789:AAExampleTokenValue"
      allowed_users:
        - "123456789"        # identifiant numérique de l’utilisateur
        - "@alice"           # ou un nom d’utilisateur
      proxy:
        enabled: true
        url: "socks5://127.0.0.1:10808"
    ```

Pour permettre les commandes shell saisies dans Telegram, ajoutez ce réglage à
la **racine** du fichier, pas sous `telegram` :

!!! tip "Ajouter un réglage racine"

    ```yaml title="Activer le bouton Run Command"
    enable_run_command: true
    ```

## :material-menu: Menu { #menu }

Cette section référence les champs. Pour une explication guidée, consultez
[Menu](concepts/menu.md). Chaque [bouton](concepts/button.md) ou
[catégorie](concepts/category.md) accepte :

| Champ | Type | Obligatoire | Description |
|-------|------|-------------|-------------|
| `name` | chaîne | oui | Nom affiché, unique parmi les voisins sans tenir compte de la casse |
| `type` | `category` \| `button` | oui | Type d’élément |
| `items` | liste | si `category` | Enfants, au moins un |
| `function` | chaîne | si `button` | Nom de la [fonction](concepts/function.md) |
| `command` | chaîne | si `function: command` | Commande shell intégrée |
| `path` | chaîne | si `function: script` | Chemin du script |
| `icon` | chaîne | non | Emoji facultatif |
| `id` | chaîne | non | Identifiant facultatif pour cet élément ; vous pouvez l’omettre |
| `confirm` | booléen | non | Demander une [confirmation](concepts/confirmation.md), `false` par défaut |
| `timeout` | durée | non | Remplace le délai global |
| `workdir` | chaîne | non | Remplace le dossier de travail |
| `env` | objet | non | Variables d’environnement du bouton |
| `columns` | entier | non | Remplace le nombre de colonnes de la catégorie |
| `args` | chaîne | non | Arguments facultatifs de `script` |
| Tout nom de paramètre déclaré | scalaire | selon la fonction | Valeur transmise, par exemple `url`, `host`, `unit` ou `lines` |

Sur un **bouton**, toute autre clé scalaire est un paramètre de fonction. Son
nom doit être déclaré par la fonction. [`validate`](cli.md#validate) refuse les
noms inconnus et vérifie les valeurs `int` et `bool`. Les chaînes, nombres et
booléens s’écrivent directement en YAML.

Sur une **catégorie**, toute clé extérieure aux champs ci-dessus est une erreur.
Les catégories n’exécutent pas de fonction et ne peuvent donc pas contenir de
paramètres.

`command`, `path` et `args` remplissent les paramètres de même nom. Écrivez les
autres [paramètres](concepts/parameter.md) directement sur le bouton, sans bloc
`params:`. Consultez
[Fonctions → Transmettre des valeurs](functions/index.md#passing-values-from-a-button).

## :material-math-log: `logging` { #logging }

Facultatif. Sans réglage, un journal console sur `stderr` au niveau `info` est utilisé.

!!! example "Écrire les journaux normaux et un fichier d’audit"

    ```yaml title="section logging avec un fichier d’audit"
    logging:
      logs:
        default:
          level: info
          format: console   # ou JSON
          output:
            - output: stderr
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

Sorties prises en charge : `stdout`, `stderr`, `file`, `discard`.

Le journal `audit` enregistre chaque commande, la personne, le bouton, le code
de sortie et la durée. Consultez [Journal d’audit](concepts/audit-log.md).

## Pages associées

- [Exécuter dans la CLI](installation/download-and-run.md) — créer et lancer une première configuration
- [Menu](concepts/menu.md) — l’arborescence en détail
- [Fonctions](functions/index.md) — signification de `function`, `command`, `path` et `args`
- [CLI](cli.md) — valider et exécuter la configuration
