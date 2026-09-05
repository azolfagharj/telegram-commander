---
title: 分类
description: 将按钮整理到子菜单中。分类通过 items 包含更多按钮或分类，让较长的菜单在手机上仍然简洁易用。
icon: material/folder-outline
---

# :material-folder-outline: 分类

分类是一种打开子菜单而不运行命令的菜单节点。它使用 `items`
包含更多[按钮](button.md)或分类，而不使用[函数](function.md)。

!!! example "包含一个按钮的分类"

    ```yaml title="A System category"
    - name: System
      type: category
      items:
        - name: Uptime
          type: button
          function: command
          command: "uptime"
    ```

按钮和分类共同组成一棵树。详见[菜单](menu.md)。

## 配置

分类字段（`type`、`items`、`columns` 等）详见
[配置→菜单](../configuration.md#menu)。

## 相关

- [按钮](button.md) — 点击时运行函数
- [菜单](menu.md) — 嵌套、布局和分页
