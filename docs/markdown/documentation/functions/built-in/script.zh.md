---
title: 脚本
description: 内置脚本函数在服务器上运行脚本文件，并在路径后添加可选参数。该脚本必须是可执行的。
icon: material/script-text
---

# :material-script-text: `script`

`script` 是[内置函数](../index.md#built-in-functions)。它运行脚本文件，
并可附加参数。

| 参数 | 必填 | 默认值 | 含义 |
|-----------|----------|---------|---------|
| `path` | 是 | — | 脚本的路径 |
| `args` | 否 | 空 | 路径之后传递的参数 |

!!! example "使用参数运行脚本"

    ```yaml title="Nightly report button"
    - name: Nightly report
      type: button
      function: script
      path: "/usr/local/bin/report.sh"
      args: "--today"
    ```

这会运行 `/usr/local/bin/report.sh --today`。如果省略 `args`，则只运行
`/usr/local/bin/report.sh`。

## 脚本必须是可执行的

该路径是直接运行的，因此该文件需要执行位：

!!! tip "使脚本可执行一次"

    ```bash title="Allow the file to run"
    chmod +x /usr/local/bin/report.sh
    ```

如果无法修改该文件，可改为通过解释器运行。发行包中的
[`echo-script`](../release-pack/echo-script.md) 示例会先调用 `bash`，
正是采用这种方式。

## 相关

- [`command`](command.md) — 运行内联命令而不是文件
- [回显脚本](../release-pack/echo-script.md) — 通过 Bash 运行脚本
- [内置函数](../index.md#built-in-functions) — 两个内置函数
- [Shell](../../concepts/shell.md) — 命令如何执行
