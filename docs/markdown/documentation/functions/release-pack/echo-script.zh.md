---
title: 回显脚本
description: 随附的 echo-script 函数通过 Bash 运行脚本文件，因此脚本不需要执行位。它从按钮获取 path 和 args。
icon: material/bash
---

# :material-bash: 回显脚本

`echo-script` 通过 `bash` 运行脚本，因此脚本文件本身不需要执行位。
它是可直接用于按钮的[随附函数](../index.md#custom-functions)之一。

- 运行：`bash {{.path}}{{if .args}} {{.args}}{{end}}`
- `path`（必需）：脚本文件的路径
- `args`（可选）：额外参数

!!! example "函数文件"

    ```yaml title="functions/echo-script.yaml"
    name: echo-script
    run: "bash {{.path}}{{if .args}} {{.args}}{{end}}"
    params:
      - name: path
        type: string
        required: true
        description: Path to the script file
      - name: args
        type: string
        required: false
        description: Optional script arguments
    ```

将两个参数名称直接写在按钮上。

!!! example "通过 Bash 运行脚本"

    ```yaml title="Run cleanup button"
    - name: Run cleanup
      type: button
      function: echo-script
      path: "/opt/scripts/cleanup.sh"
      args: "--verbose"
    ```

该按钮会运行 `bash /opt/scripts/cleanup.sh --verbose`。省略 `args` 后，
则运行 `bash /opt/scripts/cleanup.sh`，因为值为空时会跳过
`{{if .args}}` 部分。

## 相关

- [`script`](../built-in/script.md) — 直接运行可执行脚本
- [占位符](../write-your-own/placeholders.md) — `{{if .args}}` 的工作原理
- [自定义函数](../index.md#custom-functions) — 全部五个随附示例
