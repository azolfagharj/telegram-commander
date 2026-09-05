---
title: 作为服务运行
description: 使用 systemd 让机器人在后台运行，使其开机启动并在故障后自动恢复；既可使用 root，也可使用普通用户。
icon: material/server
---

# :material-server: 作为服务运行

本页介绍如何将 Telegram Commander 作为 systemd 服务运行。首次在终端中
分步运行时，请参阅[在 CLI 中运行](download-and-run.md)。

## :material-cog-outline: systemd（根）

自己创建单元文件，例如`/etc/systemd/system/telegram-commander.service`：

!!! warning "此示例以 root 身份运行每个按钮"

    ```ini title="/etc/systemd/system/telegram-commander.service"
    [Unit]
    Description=telegram-commander Telegram bot
    After=network-online.target
    Wants=network-online.target

    [Service]
    Type=simple
    ExecStart=/path/to/telegram-commander run --config /path/to/configfile.yaml
    Restart=on-failure
    RestartSec=5

    [Install]
    WantedBy=multi-user.target
    ```

    如果没有 `User=` 行，服务会以 root 身份启动，因此菜单中的每条命令
    都拥有系统的完整权限。如果按钮不需要这么高的权限，请在 `[Service]`
    部分添加 `User=someone`。

替换占位符路径，然后：

!!! example "加载并启动服务"

    ```bash title="Enable, start, and watch the service"
    sudo systemctl daemon-reload
    sudo systemctl enable --now telegram-commander
    sudo systemctl status telegram-commander
    sudo journalctl -u telegram-commander -f
    ```

!!! info "配置更改需要重新启动"

    服务会让机器人持续运行，但只会在启动时读取一次
    [配置文件](../concepts/config-file.md)。编辑配置后，请运行
    `sudo systemctl restart telegram-commander`。

## 相关页面

- [在 CLI](download-and-run.md) 中运行 — 解释了第一次运行
- [配置](../configuration.md) — 配置文件
- [CLI](../cli.md) — `run`、`validate` 等
