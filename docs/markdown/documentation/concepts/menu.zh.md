---
title: 菜单
description: 构建机器人在 Telegram 中显示的菜单。将按钮和分类组成树状结构，并了解“首页”和“返回”如何用于浏览菜单。
icon: material/view-list
---

# :material-view-list: 菜单

您的菜单是顶级 `menu` 键下的节点树。有两个
节点种类：

- **[分类](category.md)** 打开子菜单，包含 `items`。
- **[按钮](button.md)** 运行操作，并指定一个[函数](function.md)。

如果不熟悉这些概念，请先阅读[按钮](button.md)。完整字段列表详见
[配置 → 菜单](../configuration.md#menu)。

## :material-format-list-bulleted: 平面菜单

最简单的菜单是没有嵌套的按钮列表：

!!! example "用三个按钮制作菜单"

    ```yaml title="Three buttons, no categories"
    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"

      - name: Free memory
        type: button
        function: command
        command: "free -h"

      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

在 Telegram 中发送 `/start`。您将看到您的菜单。

## :material-folder-outline: 按分类分组

随着菜单增长，可以将相关操作放入分类。点击分类后会显示其中的 `items`。
菜单上始终显示“首页”，“返回”只会在分类内部出现。

!!! example "将按钮分组"

    ```yaml title="System and Services categories"
    menu:
      - name: System
        type: category
        icon: "💻"
        items:
          - name: Uptime
            type: button
            function: command
            command: "uptime"

          - name: Free memory
            type: button
            function: command
            command: "free -h"

      - name: Services
        type: category
        icon: "🔧"
        items:
          - name: Restart nginx
            type: button
            function: command
            command: "systemctl restart nginx"
            confirm: true
    ```

分类可以任意嵌套。每个分类必须至少包含一个
项目。

## :material-family-tree: 同级节点名称必须唯一

同一父节点下的两个节点不能共享名称（比较忽略大小写）。
这很好，因为它们属于不同的类别：

!!! example "在不同父节点下重复使用名称"

    ```yaml title="The same name under two parents"
    menu:
      - name: Web
        type: category
        items:
          - name: Restart          # 有效
            type: button
            function: command
            command: "systemctl restart nginx"
      - name: Database
        type: category
        items:
          - name: Restart          # 有效，父级不同
            type: button
            function: command
            command: "systemctl restart postgresql"
    ```

## :material-emoticon-outline: 图标

`icon` 是显示在名称之前的可选表情符号。它纯粹是装饰性的。

!!! warning "此按钮重新启动机器"

    ```yaml title="A button with an emoji icon"
    - name: Reboot
      type: button
      icon: "🔁"
      function: command
      command: "reboot"
      confirm: true
    ```

!!! tip "挑选一个简单、常见的表情符号"

    一些表情符号使一些手机显示按钮文本被切断或溢出
    按钮。如果按钮看起来被切断，请尝试使用不同的表情符号。

## :material-cellphone-text: Telegram 菜单的外观 { #how-the-telegram-menu-looks }

所有按钮都出现在消息框下方的键盘上（
使用消息框右端的小按钮显示和隐藏）。
该键盘始终跨越聊天的整个宽度，因此按钮文本永远不会
挤压或切断。

- **主页**始终是每个屏幕上的第一个按钮。点击它即可返回
  第一个屏幕。
- **返回**在分类内出现。
- **$ >_ Run Command** 当 `enable_run_command` 打开时出现（见下文）。
- 默认情况下，每行有两个项目。分类可以使用 `columns` 更改此设置。
  如果屏幕有许多项目，**上一个**和**下一个**可让您翻阅它们。
- 带有 `confirm: true` 的按钮在运行前询问“是”/“取消”。

!!! info "菜单标题被重用，输出保持"

    新的菜单标题（主页、类别、页面）取代了之前的菜单标题，因此
    聊天不会被空白页面填满。**正在运行**提示和命令输出会留在聊天中，
    因此再次打开菜单后仍能查看之前运行的内容。

!!! info "长输出以多条消息到达"

    命令输出以代码块显示。如果超过一条 Telegram 消息的长度，
    它会拆成多条消息，每条都回复上一条。
    最后一部分保留与您所在页面相同的按钮，因此**返回**
    仍然意味着离开该类别。参见
    [配置 → 可查看的命令输出量](../configuration.md#how-much-command-output-you-see)。

## :material-help-circle-outline: 确认 { #confirmation }

将 `confirm: true` 添加到任何按钮以需要第二次点击（" 您确定吗？"）
在它运行之前。将其用于任何破坏性的事情。参见
[确认](confirmation.md)为概念。

!!! warning "该按钮停止服务"

    ```yaml title="A button that asks first"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

确认提示会在一段时间后过期（默认 5 分钟）。可通过 `confirm_ttl`
修改，详见[配置 → 根字段](../configuration.md#root-fields)。

## :material-tune-variant: 每个按钮覆盖

一些全局设置可以在单个按钮上覆盖：

!!! example "给一键自己的设置"

    ```yaml title="One button with its own timeout, folder, and variables"
    - name: Long backup
      type: button
      function: command
      command: "/usr/local/bin/backup.sh"
      timeout: "10m"          # 执行时间可能超过全局超时
      workdir: "/var/backups" # 在此目录运行
      env:
        BACKUP_MODE: "full"   # 此命令的额外环境变量
    ```

请参阅以下位置的完整字段列表
[配置→菜单](../configuration.md#menu)。

## :material-view-grid-outline: 控制布局

`menu_columns` 设置每行显示的**项目**按钮数量（默认 2）。
分类可以使用 `columns` 覆盖此设置。当菜单项目超过 `page_size`
（默认 8）时会分页，并显示“上一页”或“下一页”，直到到达两端。详见
[配置 → 根字段](../configuration.md#root-fields)。

## :material-console: 运行命令

如果您在配置的根目录设置 `enable_run_command: true`，
**$ >_ Run Command** 按钮会保留在菜单中（在分类内位于“返回”之后，
在首页位于“首页”之后）。点击后发送要运行的 shell 命令。机器人会使用
与其他按钮相同的 shell、超时、工作目录和输出限制。

“首页”或“返回”会取消提示，不运行任何内容。此功能默认关闭。
详见[配置 → 根字段](../configuration.md#root-fields)。

!!! warning "此交接整机"

    启用“运行命令”后，任何允许使用机器人的人都可以在
    主机，而不仅仅是您定义的按钮。仅当您信任时才打开它
    到目前为止每个允许的用户。

## :material-timer-sand: 一次一个命令

您的命令一个接一个地运行，而不是并排运行。

!!! info "第二次点击会等待"

    如果第一个按钮仍在运行时点击第二个按钮，第二个操作会等待，随后单独运行。
    开始时会显示它的**正在运行**提示。这样可避免您自己的两个操作同时争用
    同一服务或文件。其他用户不会被您的命令阻塞；每个用户都有自己的队列。

## :material-function-variant: 点击按钮时运行的内容

每个按钮都通过 `function` 字段指向一个**函数**。按钮
上面的示例中使用内置的 `command` 函数。要了解
函数、内置函数与自定义函数以及如何添加自己的函数，请阅读
[函数](../functions/index.md)。

## :material-link-variant: 相关页面

- [按钮](button.md) — 什么是按钮
- [分类](category.md) — 子菜单节点
- [配置 → 菜单](../configuration.md#menu) — 每个字段
