---
title: URL Curl
description: 随附的 curl-url 函数使用 curl 获取 URL。使用 url 键直接在按钮中指定 URL。
icon: material/web
---

# :material-web: URL Curl

`curl-url` 使用 `curl` 获取 URL。遇到 HTTP 错误时会失败（`-f`），
并在 30 秒后停止。它是可直接用于按钮的[随附函数](../index.md#custom-functions)之一。

- 运行：`curl -fsSL --max-time 30 {{.url}}`
- `url`（必填）：请求的URL

!!! example "函数文件"

    ```yaml title="functions/curl-url.yaml"
    name: curl-url
    run: "curl -fsSL --max-time 30 {{.url}}"
    params:
      - name: url
        type: string
        required: true
        description: URL to request
    ```

## 添加一个按钮

直接在按钮上写入 `url`。它的名称与函数声明的参数一致。

!!! example "检查端点"

    ```yaml title="Check API button"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

该按钮运行 `curl -fsSL --max-time 30 https://example.com/health`。
如果 `url` 缺失，[`validate`](../../cli.md#validate) 会报告错误。

## 相关

- [参数](../../concepts/parameter.md) — 如何检查按钮值
- [自定义函数](../index.md#custom-functions) — 全部五个随附示例
