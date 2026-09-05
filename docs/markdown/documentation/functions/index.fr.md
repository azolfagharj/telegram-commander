---
title: Qu’est-ce qu’une fonction ?
description: Une fonction transforme les valeurs d’un bouton en commande shell. Découvrez comment un bouton fournit les paramètres déclarés par une fonction.
---

# :material-function-variant: Qu’est-ce qu’une fonction ?

Une **fonction** est une recette qui transforme des valeurs nommées, ses
[paramètres](../concepts/parameter.md), en commande shell. Chaque
[bouton](../concepts/button.md) de votre menu nomme exactement une fonction dans
son champ `function`.

Imaginez une fonction comme une commande à compléter. Une fonction d’utilisation
du disque possède un emplacement vide, le chemin, que chaque bouton remplit.

## Que se passe-t-il lorsque vous appuyez sur un bouton ?

1. Le bot recherche la fonction indiquée dans le champ `function` du bouton.
2. Il collecte les valeurs écrites sur ce bouton.
3. Il construit une commande shell avec ces valeurs.
4. Il exécute la commande dans le [shell](../concepts/shell.md) et renvoie la
   sortie dans la conversation sous forme de bloc de code.

Si la fonction n’existe pas ou qu’une valeur obligatoire manque, le bot ne
démarre pas : [`validate`](../cli.md#validate) signale d’abord le problème.

## Exemple détaillé

La fonction `command` est intégrée et toujours disponible. Elle exécute la
valeur du champ `command` du bouton.

!!! example "Utiliser la fonction intégrée command"

    ```yaml title="Bouton Uptime"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

Appuyez sur **Uptime** : le bot exécute `uptime` sur le serveur et renvoie la sortie.

## Transmettre des valeurs depuis un bouton { #passing-values-from-a-button }

Écrivez chaque valeur directement sur le bouton. Deux méthodes sont possibles :

1. Utilisez les champs raccourcis `command`, `path` et `args`. Chacun remplit
   le paramètre portant le même nom.
2. Pour tout autre paramètre, utilisez son nom comme clé du bouton.

!!! example "Transmettre une URL avec son nom de paramètre"

    ```yaml title="Bouton de vérification de l’API"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

Ici, `url` correspond au paramètre `url` déclaré par `curl-url`. La même règle
s’applique à `host`, `unit` et `lines`.

!!! warning "Ne placez pas les valeurs dans `params:`"

    `params:` appartient au fichier d’une fonction personnalisée, où il déclare
    les valeurs acceptées. Sur un bouton, écrivez chaque valeur directement :

    ```yaml title="Les valeurs se placent directement sur le bouton"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

    Les valeurs YAML numériques n’ont pas besoin de guillemets.

[`validate`](../cli.md#validate) compare chaque clé aux paramètres déclarés par
la fonction. Une faute ou un nom non déclaré fait échouer la validation. Les
valeurs `int` et `bool` sont aussi vérifiées. Les valeurs facultatives utilisent
leur valeur par défaut si le bouton les omet.

## Deux types de fonction { #two-kinds-of-function }

Les fonctions sont fournies dans le programme ou proviennent d’un fichier YAML
sur votre serveur. Une fois utilisées par un bouton, elles se comportent de la
même manière.

| | Intégrée | Personnalisée |
|---|---|---|
| Origine | Fournie dans le programme | Un fichier YAML que vous écrivez |
| Fichier à créer ? | Non | Oui, un fichier par fonction |
| Noms | Réservés (`command`, `script`) | Tout nom non réservé |
| Toujours disponible ? | Oui | Seulement avec `function_directory` |
| Modifiable ? | Non | Oui, ce sont vos fichiers |

Vous pouvez mélanger les deux types dans un menu. Commencez avec des boutons
`command`, puis créez des fonctions personnalisées lorsque vous répétez une
commande en ne changeant qu’une petite partie.

### Fonctions intégrées { #built-in-functions }

Deux fonctions sont toujours chargées, même sans `function_directory`. Leurs
champs `command`, `path` et `args` sont placés directement sur le bouton.

| Fonction | Action | Obligatoire | Facultatif |
|----------|--------|-------------|------------|
| [`command`](built-in/command.md) | Exécute une commande shell telle quelle | `command` | — |
| [`script`](built-in/script.md) | Exécute un script avec des arguments | `path` | `args` |

Les deux noms sont **réservés**. Une fonction personnalisée ne peut pas les
réutiliser : le chargeur s’arrête avec une erreur comme
`function name "command" is reserved`.

### Fonctions personnalisées { #custom-functions }

Une fonction personnalisée est un fichier YAML décrivant une commande
réutilisable. Placez ces fichiers dans votre dossier et indiquez-le avec
[`function_directory`](../configuration.md#function_directory-rules).

!!! example "Indiquer le dossier contenant les fonctions"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

Au démarrage, le bot lit ce dossier et ses sous-dossiers, puis charge tous les
fichiers `.yaml` et `.yml`. Les autres fichiers sont ignorés.

L’archive de la version contient déjà un dossier `functions/` avec cinq exemples :

| Fonction | Action | Valeurs du bouton |
|----------|--------|--------------------|
| [Script d’écho](release-pack/echo-script.md) | Exécute un script avec Bash | `path`, `args` facultatif |
| [Chemin du disque](release-pack/disk-path.md) | Affiche l’utilisation du disque | `path` facultatif |
| [URL Curl](release-pack/curl-url.md) | Récupère une URL | `url` |
| [Ping d’un hôte](release-pack/ping-host.md) | Interroge un hôte | `host`, `count` facultatif |
| [Unité du journal](release-pack/journal-unit.md) | Affiche les journaux récents d’un service | `unit`, `lines` facultatif |

Pour écrire la vôtre, commencez par la
[Structure des fichiers](write-your-own/file-structure.md) ou suivez le
[guide pas à pas](write-your-own/step-by-step.md).

!!! tip "Vérifier les fonctions chargées"

    ```bash title="Lister toutes les fonctions visibles par le bot"
    ./telegram-commander list-functions --config config.yaml
    ```

    Les fonctions intégrées affichent `source=builtin` ; les fonctions
    personnalisées affichent leur fichier d’origine.

## Consignes de sécurité

!!! warning "Les boutons s’exécutent avec les droits du bot"

    Les commandes disposent des droits du compte qui exécute le bot. Si c’est
    root, comme dans la configuration de [service](../installation/run-as-a-service.md)
    par défaut, les boutons peuvent tout faire sur l’hôte. N’ajoutez que des
    [utilisateurs autorisés](../configuration.md#telegram) de confiance.

    Les paramètres sont insérés comme texte. Traitez-les comme des entrées
    shell : préférez des valeurs fixes et ajoutez
    [`confirm: true`](../concepts/confirmation.md) aux actions destructrices.

!!! info "Les longues sorties sont limitées et divisées"

    Les commandes s’arrêtent à leur `timeout` et le bot conserve au plus
    `max_output_bytes`. Une sortie dépassant un message Telegram est divisée.
    Consultez [Configuration → Quantité de sortie affichée](../configuration.md#how-much-command-output-you-see).

## Pages associées

- [`command`](built-in/command.md) — exécuter une commande shell
- [`script`](built-in/script.md) — exécuter un fichier de script
- [Guide pas à pas](write-your-own/step-by-step.md) — créer votre première fonction
- [Menu](../concepts/menu.md) — comment les boutons référencent les fonctions
- [Paramètre](../concepts/parameter.md) — les valeurs nommées d’une fonction
