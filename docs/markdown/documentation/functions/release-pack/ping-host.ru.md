---
title: Пинг-хост
description: Функция `ping-host` из комплекта проверяет узел несколько раз. Установите хост и дополнительный счетчик прямо на кнопке.
icon: material/access-point-network
---
# :material-access-point-network: Пинг-хост

`ping-host` несколько раз пингует хост и отправляет результат обратно. Это один
из
[функций из комплекта](../index.md#custom-functions), которые можно использовать на кнопках без изменений.

- Запускается: `ping -c {{.count}} {{.host}}`
- `host` (обязательно): имя хоста или IP.
- `count` (необязательно, по умолчанию `4`): количество пакетов.

!!! example "Файл функции"

    ```yaml title="functions/ping-host.yaml"
    name: ping-host
    run: "ping -c {{.count}} {{.host}}"
    params:
      - name: host
        type: string
        required: true
        description: Hostname or IP
      - name: count
        type: string
        required: false
        default: "4"
        description: Number of ping packets
    ```

## Добавляем кнопку

!!! example "Пинг фиксированного хоста три раза"

    ```yaml title="Ping gateway button"
    - name: Ping gateway
      type: button
      function: ping-host
      host: "192.168.1.1"
      count: 3
    ```

`count` — числовое значение, поэтому кавычки ему не нужны. Вы также можете
оставить его для использования значение по умолчанию `4`:

!!! example "Использовать счетчик по умолчанию"

    ```yaml title="Ping DNS button"
    - name: Ping DNS
      type: button
      function: ping-host
      host: "1.1.1.1"
    ```

## Похожие

- [Правила](../write-your-own/rules.md) — как ведут себя значения по умолчанию и
  обязательные значения
- [Пользовательские функции](../index.md#custom-functions) — все пять примеров в
  комплекте
