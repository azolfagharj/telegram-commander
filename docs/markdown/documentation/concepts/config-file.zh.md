---
icon: material/file-document-outline
title: 配置文件
description: 一个 YAML 文件包含机器人令牌、允许的用户和按钮菜单。使用 --config 将其传给 run、validate、fmt 和 list-functions。
---

# :material-file-document-outline: 配置文件

一个 YAML 文件描述所有内容：机器人令牌、可以使用机器人的用户，以及
[按钮](button.md)菜单。使用 `--config` 将其传给
读取它的命令：`run`、`validate`、`fmt` 和 `list-functions`。其他
命令（例如 `version` 和 `completion`）不采用配置文件。

!!! example "包含一个按钮的可用配置"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN"
      allowed_users:
        - "123456789"

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

## 相关

- [配置](../configuration.md) — 每个字段、默认值和验证规则
- [菜单](menu.md) — 按钮和分类树
- [CLI](../cli.md) — 使用 `--config` 传递文件
- [在 CLI 中运行](../installation/download-and-run.md) — 创建第一个配置
