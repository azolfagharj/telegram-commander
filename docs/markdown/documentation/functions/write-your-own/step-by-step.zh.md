---
title: 分步指南
description: 从头创建第一个 Telegram Commander 函数，为其添加按钮并验证配置，然后运行机器人，在聊天中查看结果。
icon: material/format-list-numbered
---

# :material-format-list-numbered: 分步指南

本页会从空文件夹开始，逐步创建一个小函数和可用按钮。
整个过程只需几分钟，除配置文件外无需其他准备。

## 1. 将配置指向文件夹

!!! example "将文件夹添加到您的配置"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

该路径是相对于配置文件的。参见
[配置→`function_directory`规则](../../configuration.md#function_directory-rules)。

## 2. 编写函数文件

创建 `functions/greet.yaml`。该参数名为 `args`，因此按钮可以
填写：

!!! example "您的第一个函数"

    ```yaml title="functions/greet.yaml"
    name: greet
    run: "echo Hello {{.args}}"
    params:
      - name: args
        required: true
        description: Name to greet
    ```

## 3. 添加一个使用它的按钮

!!! example "新函数的按钮"

    ```yaml title="Say hello button"
    - name: Say hello
      type: button
      function: greet
      args: "world"
    ```

## 4. 验证并列出函数

!!! example "检查配置，然后看看加载了什么"

    ```bash title="Validate, then list the functions"
    ./telegram-commander validate --config config.yaml
    ./telegram-commander list-functions --config config.yaml
    ```

您应该看到 `greet` 行。每行显示名称、函数所在位置
来自，以及它有多少个参数。参见
[CLI → 列表函数](../../cli.md#list-functions)。

如果 `validate` 报错，消息会指出相关按钮和缺少的值，详见[规则](rules.md)。

## 5. 运行机器人并点击按钮

!!! example "在前台启动机器人"

    ```bash title="Run and watch the output"
    ./telegram-commander run --config config.yaml
    ```

在 Telegram 中发送 `/start`，点击 **Say hello**，机器人就会运行 `echo Hello world`
并将 `Hello world` 作为代码块发送回来。

## 6. 使其可重复使用

更改按钮的 `args` 值，或添加另一个不同的按钮。
函数保持不变：

!!! example "两个按钮，一个函数"

    ```yaml title="Greeting buttons"
    - name: Greet world
      type: button
      function: greet
      args: "world"
    - name: Greet team
      type: button
      function: greet
      args: "team"
    ```

## 相关

- [占位符](placeholders.md) — 将可选部分添加到命令中
- [规则](rules.md) — 加载程序接受的内容
- [文件结构](file-structure.md) — 每个字段都有解释
- [菜单](../../concepts/menu.md) — 菜单中按钮的放置位置
