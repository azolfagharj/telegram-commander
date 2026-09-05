---
title: 文件结构
description: 说明 Telegram Commander 函数文件的每个字段，包括 name、run，以及带有 type、required、default 和 description 的 params 列表。
icon: material/file-tree
---

# :material-file-tree: 文件结构

一个文件描述一个函数。将文件放在
[`function_directory`](../../configuration.md#function_directory-rules)，
包括子文件夹，并为其指定 `.yaml` 或 `.yml` 名称。文件名
本身并不重要：函数的名称来自内部的 `name` 字段。

!!! example "为每个文件定义一个自定义函数"

    ```yaml title="Custom function file"
    name: my-function          # 必填，必须唯一且不是保留名称
    run: "echo {{.args}}"      # 必填，要运行的命令
    params:                    # 可选参数列表
      - name: args             # 每个参数都必须填写
        type: string           # 可选：string（默认）、int 或 bool
        required: true         # 可选，默认为 false
        default: ""            # 未提供时使用的可选值
        description: Some text  # 给自己的可选备注
    ```

## 顶级字段

| 字段 | 必填 | 含义 |
|-------|----------|---------|
| `name` | 是 | 按钮在 `function` 字段中使用的名称 |
| `run` | 是 | 要运行的命令，使用[占位符](placeholders.md)插入值 |
| `params` | 否 | 函数接受的命名值列表 |

根本没有 `params` 的函数是有效的：`run` 是一个固定命令。

## 参数字段

| 字段 | 必填 | 默认值 | 含义 |
|-------|----------|---------|---------|
| `name` | 是 | — | `run`里面`{{.name}}`使用的名称 |
| `type` | 否 | `string` | 值类型：`string`、`int` 或 `bool` |
| `required` | 否 | `false` | 按钮必须提供一个值 |
| `default` | 否 | 空 | 未提供值时使用 |
| `description` | 否 | 空 | 给自己的注释，未在 Telegram 中显示 |

!!! info "会检查 `type`"

    声明为 `int` 的值和默认值必须是整数，声明为 `bool` 的值和默认值
    必须是有效的布尔值。无效值会导致 [`validate`](../../cli.md#validate) 失败。

!!! info "按钮键与参数名称匹配"

    直接在同名按钮上写入任何参数。名称如
    `url`、`host` 和 `lines` 与 `command`、`path` 和 `args` 一起使用。参见
    [从按钮传递值](../index.md#passing-values-from-a-button)。

## 文件夹布局

您可以随意组织文件夹：

!!! example "子文件夹也被读取"

    ```text title="functions/"
    functions/
      disk.yaml
      logs/
        nginx.yaml
        app.yml
    ```

所有三个文件均已加载。具有任何其他扩展名的文件将被跳过。

## 相关

- [规则](rules.md) — 加载程序拒绝的内容
- [占位符](placeholders.md) — 编写 `run` 命令
- [分步指南](step-by-step.md) — 构建第一个函数
