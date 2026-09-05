---
title: Использование диска
description: Функция `disk-path` из комплекта показывает использование диска с помощью df -h. Его единственным параметром является путь, и если его опустить, он возвращается к корневой файловой системе.
icon: material/harddisk
---
# :material-harddisk: Использование диска

`disk-path` показывает использование диска для пути с `df -h`. Это один из
[функций из комплекта](../index.md#custom-functions), которые можно использовать на кнопках без изменений.

- Запускается: `df -h {{.path}}`
- `path` (необязательно, по умолчанию `/`): путь к файловой системе для
  проверки.

!!! example "Файл функции"

    ```yaml title="functions/disk-path.yaml"
    name: disk-path
    run: "df -h {{.path}}"
    params:
      - name: path
        type: string
        required: false
        default: "/"
        description: Filesystem path to check
    ```

Поскольку `path` является необязательным и имеет значение по умолчанию, кнопка
может его не указывать и всё равно работает.

!!! example "Используйте путь по умолчанию или выберите один"

    ```yaml title="Disk usage buttons"
    - name: Disk (root)
      type: button
      function: disk-path        # без поля path: используется значение "/" по умолчанию
    - name: Disk (var)
      type: button
      function: disk-path
      path: "/var"
    ```

Первая кнопка запускает `df -h /`, а вторая — `df -h /var`.

## Похожие

- [Правила](../write-your-own/rules.md) — как ведут себя значения по умолчанию и
  обязательные значения
- [Пользовательские функции](../index.md#custom-functions) — все пять примеров в
  комплекте
- [`command`](../built-in/command.md) — вместо этого напишите полную команду
