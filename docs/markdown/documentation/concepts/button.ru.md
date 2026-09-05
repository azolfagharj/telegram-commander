---
title: Кнопка
description: Нажимаемый элемент меню, который запускает функцию. Напишите каждое значение функции непосредственно на кнопке с именем параметра, объявленным функцией.
icon: material/gesture-tap-button
---
# :material-gesture-tap-button: Кнопка

Нажимаемый пункт в меню Telegram. Кнопка имеет `name` и при нажатии запускает
[функцию](function.md) на вашем сервере и отправляет результат обратно в чат.

Кнопки и [категории](category.md) вместе образуют дерево [меню](menu.md). под
ключом верхнего уровня `menu` в вашем [файле конфигурации](config-file.md).
Кнопка делает работу; категория открывает только подменю.

## Как выглядит кнопка

Кнопки появляются на клавиатуре под окном сообщения, по две в строке по
умолчанию. Имя, которое вы выбираете, представляет собой текст на ключе, поэтому
сделайте его достаточно коротким, чтобы его можно было прочитать дальше.
телефон. См. [Меню → Как выглядит меню
Telegram](menu.md#how-the-telegram-menu-looks).

## :material-format-list-checks: Части кнопки

!!! example "Каждая строка управляет одной частью кнопки"

    ```yaml title="One button, fully labelled"
    - name: Restart nginx # (1)!
      type: button # (2)!
      icon: "🔄" # (3)!
      function: command # (4)!
      command: "systemctl restart nginx" # (5)!
      confirm: true # (6)!
    ```

    1. Текст на ключе в Telegram. Оно должно быть уникальным среди своих соседей
       по
        то же самое меню.
    2. Всегда здесь `button`. Вместо этого используйте `category`, если вам
       нужно подменю.
    3. Необязательный смайлик, отображаемый перед именем. Это только украшение и
       изменения
        ничего о том, что работает.
    4. Какую [функцию ](function.md) использовать. `command` — встроенный,
       который
        запускает команду оболочки.
    5. Что запускает функция `command`. Все, что вы можете ввести в терминале
        работает здесь.
    6. Необязательно. Спросите: «Вы уверены?» перед запуском. Не включайте
       подтверждение для действий,
        которые только читают данные.

## Что происходит, когда вы нажимаете на него

1. Бот публикует короткую строку **Выполняется**, чтобы вы знали, что оно
   началось.
2. Команда запускается на машине, на которой работает бот.
3. Вывод возвращается в виде блока кода с кодом завершения и его длиной.
    взял. Длинный вывод приходит в виде нескольких сообщений подряд.
4. Вы остаетесь в том же меню, в котором находились, поэтому нажатие
   **Назад** по-прежнему оставляет его.
    категория.

## :material-code-braces: Общие кнопки

=== "Проверьте что-нибудь"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

=== "Управление сервисом"

    ```yaml title="Restart nginx button"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

=== "Читать журнал"

    ```yaml title="Nginx log button"
    - name: Nginx log
      type: button
      function: command
      command: "journalctl -u nginx -n 50 --no-pager"
    ```

=== "Запустить скрипт"

    ```yaml title="Nightly backup button"
    - name: Nightly backup
      type: button
      function: script
      path: "/usr/local/bin/backup.sh"
    ```

=== "Что-то разрушительное"

    ```yaml title="Stop nginx button"
    - name: Stop nginx
      type: button
      icon: "🛑"
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## :material-emoticon-outline: Значки

`icon` помещает перед именем смайлик. Это чисто косметически, так что вы можете
изменить или удалить его в любое время, не касаясь того, что запускает кнопка.

!!! example "Значок меняет только метку"

    ```yaml title="The same button, with and without an icon"
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

## :material-help-circle-outline: Спрашивайте, прежде чем рисковать кнопками

Добавьте `confirm: true`, и бот сначала спросит «Да» или «Отмена». Используйте
его для чего угодно который останавливает службу, удаляет данные или
перезагружает компьютер. Срок действия запроса истекает через некоторое время
(по умолчанию пять минут).

Прочтите [Подтверждение](confirmation.md), чтобы узнать, когда стоит спрашивать
и как это сделать. изменить время ожидания.

## Настройки только для одной кнопки

Большинство глобальных настроек можно изменить с помощью одной кнопки, что
удобно, когда одна работа ведет себя иначе, чем остальные:

!!! example "Переопределить настройки для одного медленного задания"

    ```yaml title="A slow job that runs somewhere else"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"
      workdir: "/var/backups"
      env:
        BACKUP_MODE: "full"
    ```

`timeout` дает этой команде больше времени на выполнение, `workdir` выбирает
каталог, в котором он работает, и `env` добавляет переменные среды только для
него.

## Значения для функции

Напишите значения функции прямо на кнопке. `command`, `path` и `args`
представляют собой поля ярлыков для параметров с такими именами.
Пользовательские имена, такие как `url`, `host`, `unit` и `lines` работают
одинаково.

!!! example "Передача пользовательских значений"

    ```yaml title="Recent Nginx logs"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

Не помещайте эти значения в `params:`. Каждый ключ должен соответствовать
параметру объявленный выбранной функцией. [`validate`](../cli.md#validate)
отчеты неизвестные имена, отсутствующие обязательные значения и недопустимые
значения `int` или `bool`.

## Конфигурация

Для каждого поля, которое принимает кнопка, см.
[Конфигурация → Меню](../configuration.md#menu).

## Похожие

<div class="grid cards cols-2" markdown>

- :material-folder-outline:{ .middle } __Категория__

    ---

    Открывает подменю вместо запуска чего-либо.[:octicons-arrow-right-24:
    Категория](category.md)

- :material-function:{ .middle } __Функция__

    ---

    Что на самом деле запускается при нажатии кнопки.

    [:octicons-arrow-right-24: Функция](function.md)

- :material-tune:{ .middle } __Параметр__

    ---

    Значения, необходимые функции от кнопки.

    [:octicons-arrow-right-24: Параметр](parameter.md)

- :material-view-list:{ .middle } __Меню__

    ---

    Постройте и организуйте все дерево.

    [:octicons-arrow-right-24: Меню](menu.md)

</div>
