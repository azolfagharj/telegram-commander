---
title: 占位符
description: 如何使用占位符编写函数的运行命令，包括仅在按钮提供值时出现的可选部分。
icon: material/code-braces
---

# :material-code-braces: 占位符

`run` 字段是一个小模板。双花括号中的占位符会在命令执行前
替换为参数值。

两种模式几乎涵盖了所有内容：

- `{{.name}}` 插入名为 `name` 的参数值。
- 仅当 `name` 有值时，`{{if .name}} ... {{end}}` 才包含中间部分。

## 插入一个值

!!! example " 命令中间一个值"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
    ```

带有 `path: "/var/log/app.log"` 的按钮运行
`tail -n 200 /var/log/app.log`。

## 可选部分

空值使 `{{if .name}}` 为假，因此整个块消失。这个
这就是仅在需要时添加标志的方式。

!!! example "用占位符制作可选部分"

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n 200 {{.path}}{{if .args}} | grep -- {{.args}}{{end}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        description: Optional text to filter for
    ```

使用它的两个按钮，一个带过滤器，一个不带：

!!! example "通过按钮提供值"

    ```yaml title="App log buttons"
    - name: App log
      type: button
      function: tail-log
      path: "/var/log/app.log"

    - name: App errors
      type: button
      function: tail-log
      path: "/var/log/app.log"
      args: "ERROR"
    ```

**应用程序日志**运行`tail -n 200 /var/log/app.log`，并且**应用程序错误**运行
`tail -n 200 /var/log/app.log | grep -- ERROR`。

您还可以使用 `{{if .args}} ... {{else}} ... {{end}}` 在两种形式之间进行选择。

## 您可以使用哪些名称

占位符可以使用您在 `params` 中声明的任何参数。其值为
按钮提供的内容或参数 `default`。

!!! warning "未声明的名称会在运行时失败"

    仅检查占位符语法
    [`validate`](../../cli.md#validate)，不是名字。如果`run`提到
    `{{.uri}}` 但没有参数称为 `uri`，配置验证良好并且
    当您点击该按钮时，该按钮会在聊天中报告错误。

!!! warning "值以纯文本形式插入"

    系统不会自动添加引号或进行转义。带空格或 shell 特殊字符的值
    会按原样成为命令的一部分，因此应优先使用按钮中的固定值，
    按钮并将 [`confirm: true`](../../concepts/confirmation.md) 添加到
    任何破坏性的东西。

## 相关

- [文件结构](file-structure.md) — `run` 和 `params` 所在的位置
- [规则](rules.md) — 加载程序接受的内容
- [分步指南](step-by-step.md) — 尝试从一端到另一端
