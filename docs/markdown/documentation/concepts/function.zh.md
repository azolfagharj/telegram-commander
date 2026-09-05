---
title: 函数
description: 函数将按钮上的值转换为 shell 命令。使用内置命令和脚本函数，或将您自己的命令和脚本函数添加为 YAML 文件。
icon: material/function
---

# :material-function: 函数

函数是点击[按钮](button.md)时运行的内容。它接收一些[参数](parameter.md)，
并将其转换为 shell 命令。

有两种：

- **内置函数**随程序提供（`command`、`script`），无需自行创建。
- **自定义函数**是您在 `function_directory` 中通过 YAML 文件定义的额外函数。

点击按钮后，机器人会查找对应函数，填入按钮中的值，并在
[Shell](shell.md)中运行结果。

!!! example "使用内置 `command` 函数的按钮"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

## 配置

对于 `function_directory` 和引用函数的按钮字段，请参见
[配置](../configuration.md)。

## 相关

- [参数](parameter.md) — 函数需要的命名值
- [函数](../functions/index.md#two-kinds-of-function) — 深入了解内置函数和自定义函数
- [菜单](menu.md) — 按钮如何引用函数
