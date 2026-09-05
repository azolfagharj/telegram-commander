---
icon: material/tune
title: 参数
description: 参数是函数所需的命名值。使用函数声明的相同名称将每个值直接写入按钮。
---

# :material-tune: 参数

[函数](function.md)所需的命名值。例如，内置
`command` 函数需要一个名为 `command` 的参数。您提供参数
[button](button.md) 上的值。

一些参数是必需的；其他是可选的，默认值定义在
函数文件。

## 从按钮传递一个值

将参数名称作为键直接写在按钮上。内置的
`command`、`path`、`args`字段遵循此规则，自定义参数
名称的工作方式相同。

!!! example "提供自定义参数名称"

    ```yaml title="Recent service logs"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`unit` 和 `lines` 必须由 `journal-unit` 声明。数值为
`lines` 不需要加引号。

!!! warning "不要使用嵌套 `params:` 映射"

    自定义函数文件使用 `params:` 来声明其参数。一个按钮
    没有。将 `host:`、`unit:` 或任何其他值直接放在按钮上。

## 验证

[`validate`](../cli.md#validate) 检查：

- 每个必需的参数都有一个值；
- 按钮上的每个值均由其函数声明；
- 声明为 `int` 的值包含一个整数；
- 声明为 `bool` 的值包含有效的布尔值；
- 默认值与声明的参数类型匹配。

当按钮省略可选参数时，可选参数将使用其默认值。

参数名称不能与按钮设置相同：`name`、`type`、`icon`、
`id`、`function`、`confirm`、`timeout`、`workdir`、`env`、`columns` 或
`items`。允许使用名称 `command`、`path` 和 `args`。

请参阅[函数 → 从按钮传递值](../functions/index.md#passing-values-from-a-button)
以获得完整的解释和示例。

## 配置

按钮设置和参数键请参见
[配置→菜单](../configuration.md#menu)。

## 相关

- [函数](function.md) — 使用参数的内容
- [规则](../functions/write-your-own/rules.md) — 自定义函数的参数规则
