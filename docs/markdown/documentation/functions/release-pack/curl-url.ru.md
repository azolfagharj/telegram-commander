---
title: Запрос URL-адрес
description: Функция `curl-url` из комплекта извлекает URL-адрес с помощью Curl. Добавьте URL-адрес непосредственно к кнопке с помощью ключа URL-адреса.
icon: material/web
---
# :material-web: Запрос URL

`curl-url` извлекает URL-адрес с `curl`. Он терпит неудачу из-за ошибок HTTP
(`-f`) и выдает вверх через 30 секунд. Это один из
[функций из комплекта](../index.md#custom-functions), которые можно использовать на кнопках без изменений.

- Запускается: `curl -fsSL --max-time 30 {{.url}}`
- `url` (обязательно): URL-адрес для запроса.

!!! example "Файл функции"

    ```yaml title="functions/curl-url.yaml"
    name: curl-url
    run: "curl -fsSL --max-time 30 {{.url}}"
    params:
      - name: url
        type: string
        required: true
        description: URL to request
    ```

## Добавляем кнопку

Напишите `url` прямо на кнопке. Его имя соответствует параметру, объявленному
функция.

!!! example "Проверить конечную точку"

    ```yaml title="Check API button"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

Кнопка имеет название `curl -fsSL --max-time 30 https://example.com/health`.
[`validate`](../../cli.md#validate) сообщает об ошибке, если `url` отсутствует.

## Похожие

- [Параметры](../../concepts/parameter.md) — как проверяются значения кнопок
- [Пользовательские функции](../index.md#custom-functions) — все пять примеров в
  комплекте
