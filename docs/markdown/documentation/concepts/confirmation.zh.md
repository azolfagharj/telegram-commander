---
title: 确认
description: 让机器人在按钮运行之前询问。对无法撤消的任何操作启用确认，并设置“是”或“取消”提示保持有效的时间。
icon: material/help-circle-outline
---

# :material-help-circle-outline: 确认

这是[按钮](button.md)运行前可选的“确定吗？”步骤，通过
`confirm: true` 启用，适合破坏性操作。

机器人会显示 `Confirm: … ?`，并提供“是”“取消”“首页”，在分类内还会显示
“返回”。启用“运行命令”后，该按钮也会保留在确认页面。点击“是”之前不会
运行任何内容；“取消”“首页”和“返回”都会放弃该命令。

提示会在一段时间后过期（默认 5 分钟）。可在[配置文件](config-file.md)中
通过 `confirm_ttl` 修改。过期后再次点击按钮即可获得新提示。

!!! warning "将其用于任何无法撤消的事情"

    只需在手机上轻按一下菜单键，每个人都可以使用该机器人
    看到相同的菜单。将 `confirm: true` 放在停止服务的按钮上，
    删除数据，或者重启机器。

!!! example "运行前询问的按钮"

    ```yaml title="Stop nginx button"
    - name: Stop nginx
      type: button
      function: command
      command: "systemctl stop nginx"
      confirm: true
    ```

## 配置

对于按钮上的 `confirm` 和全局 `confirm_ttl`，请参见
[配置](../configuration.md)。

## 相关

- [按钮](button.md) — 可以要求确认的节点
- [菜单 → 确认](menu.md#confirmation) — 示例和行为
