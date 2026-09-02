---
title: Step by step guide
description: Create your first Telegram Commander function from scratch, add a button for it, validate the config, and run the bot to see the result in the chat.
icon: material/format-list-numbered
---

# :material-format-list-numbered: Step by step guide

This walks through one small function from an empty folder to a working button.
It takes a few minutes and needs nothing but your config file.

## 1. Point the config at a folder

!!! example "Add the folder to your config"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

The path is relative to the config file. See
[Configuration → `function_directory` rules](../../configuration.md#function_directory-rules).

## 2. Write the function file

Create `functions/greet.yaml`. The parameter is called `args` so a button can
fill it:

!!! example "Your first function"

    ```yaml title="functions/greet.yaml"
    name: greet
    run: "echo Hello {{.args}}"
    params:
      - name: args
        required: true
        description: Name to greet
    ```

## 3. Add a button that uses it

!!! example "A button for the new function"

    ```yaml title="Say hello button"
    - name: Say hello
      type: button
      function: greet
      args: "world"
    ```

## 4. Validate and list the functions

!!! example "Check the config, then see what loaded"

    ```bash title="Validate, then list the functions"
    ./telegram-commander validate --config config.yaml
    ./telegram-commander list-functions --config config.yaml
    ```

You should see a `greet` line. Each line shows the name, where the function came
from, and how many parameters it has. See
[CLI → list-functions](../../cli.md#list-functions).

If `validate` complains, the message names the button and the value that is
missing — check [Rules](rules.md).

## 5. Run the bot and tap the button

!!! example "Start the bot in the foreground"

    ```bash title="Run and watch the output"
    ./telegram-commander run --config config.yaml
    ```

Send `/start` in Telegram, tap **Say hello**, and the bot runs `echo Hello world`
and sends `Hello world` back as a code block.

## 6. Make it reusable

Change the button's `args` value, or add a second button with a different one.
The function stays the same:

!!! example "Two buttons, one function"

    ```yaml title="Greeting buttons"
    - name: Greet world
      type: button
      function: greet
      args: "world"
    - name: Greet team
      type: button
      function: greet
      args: "team"
    ```

## Related

- [Placeholders](placeholders.md) — add optional parts to the command
- [Rules](rules.md) — what the loader accepts
- [File structure](file-structure.md) — every field explained
- [Menu](../../concepts/menu.md) — where to put the button in your menu
