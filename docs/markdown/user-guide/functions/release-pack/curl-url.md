---
title: Curl URL
description: The bundled curl-url function fetches a URL with curl. Add the URL directly to a button with the url key.
icon: material/web
---

# :material-web: Curl URL

`curl-url` fetches a URL with `curl`. It fails on HTTP errors (`-f`) and gives
up after 30 seconds. It is one of the
[bundled functions](../index.md#custom-functions) you can use from a button as
it is.

- Runs: `curl -fsSL --max-time 30 {{.url}}`
- `url` (required): the URL to request

!!! example "The function file"

    ```yaml title="functions/curl-url.yaml"
    name: curl-url
    run: "curl -fsSL --max-time 30 {{.url}}"
    params:
      - name: url
        type: string
        required: true
        description: URL to request
    ```

## Add a button

Write `url` directly on the button. Its name matches the parameter declared by
the function.

!!! example "Check an endpoint"

    ```yaml title="Check API button"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

The button runs `curl -fsSL --max-time 30 https://example.com/health`.
[`validate`](../../cli.md#validate) reports an error if `url` is missing.

## Related

- [Parameters](../../concepts/parameter.md) — how button values are checked
- [Custom functions](../index.md#custom-functions) — all five bundled examples
