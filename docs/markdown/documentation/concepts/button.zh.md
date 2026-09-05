---
title: 按钮
description: 点击后运行函数的菜单项。使用函数声明的参数名称，直接在按钮上填写各项值。
icon: material/gesture-tap-button
---

# :material-gesture-tap-button: 按钮

按钮是 Telegram 菜单中可点击的项目。按钮带有 `name`，点击后会在服务器上
运行[函数](function.md)，并将结果发回聊天。

按钮和[分类](category.md)共同组成[菜单](menu.md)树，位于
[配置文件](config-file.md)的顶级 `menu` 键下。按钮负责执行操作；
分类只负责打开子菜单。

## 按钮是什么样的

按钮显示在消息框下方的键盘上，默认情况下每行两个。
您选择的名称就是按键上的文字，因此应保持简短，方便在手机上阅读。
详见[菜单 → Telegram 菜单的外观](menu.md#how-the-telegram-menu-looks)。

## :material-format-list-checks: 按钮的组成部分

!!! example " 每行控制按钮的一部分"

    ```yaml title="One button, fully labelled"
    - name: Restart nginx # (1)!
      type: button # (2)!
      icon: "🔄" # (3)!
      function: command # (4)!
      command: "systemctl restart nginx" # (5)!
      confirm: true # (6)!
    ```

    1. Telegram 按键上显示的文字。在同一菜单的相邻节点中必须唯一。
    2. 此处始终为 `button`。当您需要子菜单时，请使用 `category`。
    3. 名称前显示的可选表情符号。它仅用于装饰，不会改变运行内容。
    4. 要使用的[函数](function.md)。`command` 是运行 shell 命令的内置函数。
    5. `command` 函数运行什么。您可以在终端中输入的任何内容
        在这里工作。
    6. 可选。运行前询问“确定吗？”。仅用于读取信息的操作可省略。

## 点击它时会发生什么

1. 机器人会发布一条简短的 **正在运行** 行，以便您知道它已开始。
2. 该命令在运行机器人的计算机上运行。
3. 输出以代码块形式返回，并包含退出代码和耗时。较长的输出会拆成连续多条消息。
4. 您仍停留在原来的菜单中，因此**返回**仍会离开当前分类。

## :material-code-braces: 常用按钮

=== "检查一些东西"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

=== "管理服务"

    ```yaml title="Restart nginx button"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

=== "阅读日志"

    ```yaml title="Nginx log button"
    - name: Nginx log
      type: button
      function: command
      command: "journalctl -u nginx -n 50 --no-pager"
    ```

=== "运行脚本"

    ```yaml title="Nightly backup button"
    - name: Nightly backup
      type: button
      function: script
      path: "/usr/local/bin/backup.sh"
    ```

=== "破坏性的东西"

    ```yaml title="Stop nginx button"
    - name: Stop nginx
      type: button
      icon: "🛑"
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## :material-emoticon-outline: 图标

`icon` 在名字前面放置一个表情符号。它纯粹是装饰性的，所以你可以
随时更改或删除它，而无需触摸按钮运行的内容。

!!! example " 图标仅更改标签"

    ```yaml title="The same button, with and without an icon"
    - name: Disk usage
      type: button
      function: command
      command: "df -h"

    - name: Disk usage
      type: button
      icon: "💾"
      function: command
      command: "df -h"
    ```

## :material-help-circle-outline: 在有风险的按钮之前询问

添加 `confirm: true`，机器人首先询问“是”或“取消”。用它来做任何事
停止服务、删除数据或重新启动计算机。提示已过期
一段时间后（默认五分钟）。

请阅读[确认](confirmation.md)，了解何时需要确认以及如何
更改等待时间。

## 一键设置

大多数全局设置都可以在单个按钮上覆盖，这在以下情况下很方便
一项工作的行为与其他工作不同：

!!! example " 覆盖一项慢速作业的设置"

    ```yaml title="A slow job that runs somewhere else"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"
      workdir: "/var/backups"
      env:
        BACKUP_MODE: "full"
    ```

`timeout` 给这个命令更长的时间来完成，`workdir` 选择
它运行的目录，`env` 专门为其添加了环境变量。

## 函数的值

直接在按钮上填写函数值。`command`、`path` 和 `args`
是具有这些名称的参数的快捷方式字段。自定义名称，例如
`url`、`host`、`unit` 和 `lines` 的工作方式相同。

!!! example "传递自定义值"

    ```yaml title="Recent Nginx logs"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

不要将这些值放入 `params:` 中。每个键必须匹配一个参数
由所选函数声明。 [`validate`](../cli.md#validate) 报告
未知名称、缺少所需值以及无效的 `int` 或 `bool` 值。

## 配置

对于按钮接受的每个字段，请参阅
[配置→菜单](../configuration.md#menu)。

## 相关

<div class="grid cards cols-2" markdown>

- :material-folder-outline:{ .middle } __类别__

    ---

    打开子菜单而不是运行某些内容。

    [:octicons-arrow-right-24: 分类](category.md)

- :material-function:{ .middle } __函数__

    ---

    点击按钮时实际运行的内容。

    [:octicons-arrow-right-24: 函数](function.md)

- :material-tune:{ .middle } __参数__

    ---

    函数需要从按钮获取的值。

    [:octicons-arrow-right-24: 参数](parameter.md)

- :material-view-list:{ .middle } __菜单__

    ---

    构建并组织整个树。

    [:octicons-arrow-right-24: 菜单](menu.md)

</div>
