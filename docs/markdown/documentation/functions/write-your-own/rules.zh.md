---
title: 规则
description: Telegram Commander 函数文件必须通过的检查，包括允许的名称、保留字、唯一名称、拼写错误拒绝以及默认值的使用方式。
icon: material/format-list-checks
---

# :material-format-list-checks: 规则

机器人启动和运行时会检查每个函数文件
[`validate`](../../cli.md#validate)。如果一个文件错误，则不会运行任何内容，因此
在您的菜单上线之前就会发现错误。

## 读取哪些文件

- 仅读取以 `.yaml` 或 `.yml` 结尾的文件。其他任何内容都将被忽略。
- `function_directory` 的子文件夹也会被读取。
- 一个文件必须准确描述一个函数。

## 命名

- **`name` 为必填** 并且必须匹配 `^[A-Za-z0-9._-]+$` — 字母、数字、
  `.`、`-` 和 `_`。没有空格。
- **名称是唯一的，忽略大小写。** `Deploy` 和 `deploy` 发生冲突，即使
  他们生活在不同的文件中。
- **禁止保留名称。** 您不能将函数命名为 `command` 或
  `script`；那些属于
  [内置函数](../index.md#built-in-functions)。

## 命令

- **需要 `run`。** 使用以下命令写入命令
  [占位符](placeholders.md) 值所在的位置。
- 验证配置时会检查占位符语法，因此损坏
  `{{` 早报。

## 参数

- **每个参数需要一个`name`**，与参数匹配相同的字符规则
  函数名称。
- **`type` 必须是** `string`、`int` 或 `bool`。省略它意味着 `string`。
  检查声明为 `int` 或 `bool` 的值和默认值。
- 参数名称不能是按钮设置：`name`、`type`、`icon`、`id`、
  `function`、`confirm`、`timeout`、`workdir`、`env`、`columns` 或 `items`。
  允许使用名称 `command`、`path` 和 `args`。
- 按钮上写入的每个值都必须与该按钮声明的参数匹配
  按钮的函数。

函数的 `run` 字段只能对声明的参数使用占位符。
未声明的占位符是一个错误。

## 拼写错误是错误

**未知键会被拒绝。** 例如将 `required:` 错写为 `requird:`，
验证会停止，而不是静默忽略错误。

!!! warning "此文件未加载"

    ```yaml title="functions/broken.yaml"
    name: broken
    run: "echo {{.args}}"
    params:
      - name: args
        requird: true      # 已拒绝：未知键
    ```

## 必需的值和默认值

- **没有值的必需参数**使验证失败并显示一条消息
  像`required parameter "args" for function "greet" is missing`。
- 存在但为空的值视为缺失。
- **可选参数**回退到其 `default`。如果没有 `default`，则
  变为空，这就是命令的可选部分被省略的方式。

!!! example " 必需和可选在一个文件中 "

    ```yaml title="functions/tail-log.yaml"
    name: tail-log
    run: "tail -n {{.args}} {{.path}}"
    params:
      - name: path
        required: true
        description: Log file path
      - name: args
        default: "200"
        description: Number of lines
    ```

    仅设置 `path` 的按钮运行 `tail -n 200 /var/log/app.log`。

## 相关

- [文件结构](file-structure.md) — 每个字段都有解释
- [占位符](placeholders.md) — 编写 `run` 命令
- [CLI → validate](../../cli.md#validate) — 自己运行检查
