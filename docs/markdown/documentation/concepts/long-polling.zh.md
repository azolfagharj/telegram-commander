---
title: 机器人如何连接
description: 该机器人连接到 Telegram，因此您永远不会在服务器上打开端口。如果 Telegram 在您的网络上被阻止，请设置代理。
icon: material/lan-disconnect
---

# :material-lan-disconnect: 机器人如何连接

机器人会主动连接 Telegram，因此无需在服务器上开放端口。

!!! tip "没有任何内容可以暴露在互联网上"

    由于连接始终由机器人向外发起，因此没有需要保护的 Web 管理面板，
    无需添加入站防火墙规则，也没有可供他人查找的公开地址。服务器即使位于
    家用路由器或严格的防火墙后，机器人仍可正常工作。

如果您的网络无法访问 Telegram，请在[配置文件](config-file.md)的
`telegram` 下设置代理。

## 配置

对于 Telegram 设置下的 `api`、`proxy` 和 `insecure`，请参阅
[配置 → telegram](../configuration.md#telegram)。

## 相关

- [允许的用户](allowed-users.md) — 谁可以使用该机器人
- [CLI → run](../cli.md#run) — 在前台启动机器人
- [作为服务运行](../installation/run-as-a-service.md) — 保持机器人运行
