---
title: 命令
description: 内置 command 函数完全按照按钮中的内容运行一条 shell 命令，包括管道、重定向和链式命令。
icon: material/console
---

# :material-console: `command`

`command`是[内置函数](../index.md#built-in-functions)。它运行单个 shell 命令
与您在按钮上写的完全一样。

| 参数 | 必填 | 默认值 | 含义 |
|-----------|----------|---------|---------|
| `command` | 是 | — | 要运行的 shell 命令 |

!!! example "按照写入的方式运行一个命令"

    ```yaml title="Show kernel button"
    - name: Show kernel
      type: button
      function: command
      command: "uname -a"
    ```

## 管道和链式命令

命令通过 [`/bin/bash -c`](../../concepts/shell.md) 运行，因此管道，
重定向和 `&&` 都在一个 `command` 值内工作。

!!! example "管道在命令内工作"

    ```yaml title="Top processes button"
    - name: Top processes
      type: button
      function: command
      command: "ps aux --sort=-%mem | head -n 10"
    ```

!!! tip "对重复命令使用自定义函数"

    如果多个按钮重复相同的命令形式，请创建自定义函数
    并只为每个按钮提供更改的值。自定义参数名称
    例如 `url` 或 `host` 可以直接写在这些按钮上。

## 相关

- [`script`](script.md) — 运行脚本文件而不是内联命令
- [内置函数](../index.md#built-in-functions) — 两个内置函数
- [Shell](../../concepts/shell.md) — 命令如何执行
- [确认](../../concepts/confirmation.md) — 在危险命令之前询问
