---
title: Journal 单元
description: 随附的 journal-unit 函数显示 systemd 单元最近的 journalctl 日志。直接在按钮上设置单元和可选行数。
icon: material/text-box-search-outline
---

# :material-text-box-search-outline: Journal 单元

`journal-unit` 显示 systemd 单元最新的 `journalctl` 日志。它是可直接用于
按钮的[随附函数](../index.md#custom-functions)之一。

- 运行：`journalctl -u {{.unit}} -n {{.lines}} --no-pager`
- `unit`（必填）：单位名称，例如`nginx.service`
- `lines`（可选，默认`50`）：显示多少行

!!! example "函数文件"

    ```yaml title="functions/journal-unit.yaml"
    name: journal-unit
    run: "journalctl -u {{.unit}} -n {{.lines}} --no-pager"
    params:
      - name: unit
        type: string
        required: true
        description: Systemd unit name (for example nginx.service)
      - name: lines
        type: string
        required: false
        default: "50"
        description: Number of log lines
    ```

## 添加一个按钮

!!! example "读取一项服务的最近日志"

    ```yaml title="Nginx logs button"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`lines` 是数字，因此不需要引号。省略它即可使用默认值 `50`。

!!! example "使用默认行数"

    ```yaml title="SSH logs button"
    - name: SSH logs
      type: button
      function: journal-unit
      unit: "ssh.service"
    ```

## 相关

- [占位符](../write-your-own/placeholders.md) — 如何填写 `{{.unit}}`
- [自定义函数](../index.md#custom-functions) — 全部五个随附示例
