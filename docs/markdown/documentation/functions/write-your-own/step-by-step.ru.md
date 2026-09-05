---
title: Пошаговое руководство
description: Создайте свою первую функцию Telegram Commander с нуля, добавьте для нее кнопку, проверьте конфигурацию и запустите бота, чтобы увидеть результат в чате.
icon: material/format-list-numbered
---
# :material-format-list-numbered: Пошаговое руководство

Это проходит через одну небольшую функцию от пустой папки до работающей кнопки.
Это займет несколько минут и не требует ничего, кроме вашего файла конфигурации.

## 1. Укажите конфиг в папке

!!! example "Добавьте папку в свою конфигурацию"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

Путь указан относительно файла конфигурации. См.
[Конфигурация → `function_directory` правила ](../../configuration.md#function_directory-rules).

## 2. Напишите файл функции

Создайте `functions/greet.yaml`. Параметр называется `args`, поэтому кнопка
может заполните его:

!!! example "Ваша первая функция"

    ```yaml title="functions/greet.yaml"
    name: greet
    run: "echo Hello {{.args}}"
    params:
      - name: args
        required: true
        description: Name to greet
    ```

## 3. Добавьте кнопку, которая его использует

!!! example "Кнопка для новой функции"

    ```yaml title="Say hello button"
    - name: Say hello
      type: button
      function: greet
      args: "world"
    ```

## 4. Проверьте и перечислите функции

!!! example "Проверьте конфигурацию, затем посмотрите, что загружено"

    ```bash title="Validate, then list the functions"
    ./telegram-commander validate --config config.yaml
    ./telegram-commander list-functions --config config.yaml
    ```

Вы должны увидеть строку `greet`. В каждой строке указано название, откуда
пришла функция и сколько у него параметров. См.
[CLI → список функций](../../cli.md#list-functions).

Если `validate` жалуется, в сообщении указывается кнопка и значение, которое
отсутствует — проверьте [Правила](rules.md).

## 5. Запустите бота и нажмите кнопку

!!! example "Запустить бота на переднем плане"

    ```bash title="Run and watch the output"
    ./telegram-commander run --config config.yaml
    ```

Отправьте `/start` в Telegram, нажмите **Поздороваться**, и бот запустит `echo
Hello world`. и отправляет `Hello world` обратно как блок кода.

## 6. Сделайте его многоразовым

Измените значение кнопки `args` или добавьте вторую кнопку с другим значением.
Функция остается прежней:

!!! example "Две кнопки, одна функция"

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

## Похожие

- [Заполнители](placeholders.md) — добавить в команду необязательные части
- [Правила](rules.md) — что принимает загрузчик
- [Структура файла](file-structure.md) — объяснение каждого поля
- [Меню](../../concepts/menu.md) — куда поставить кнопку в меню
