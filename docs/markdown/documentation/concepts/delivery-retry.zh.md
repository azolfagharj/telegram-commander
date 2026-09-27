---
title: 投递重试
description: 短暂网络中断不再让结果丢失。机器人会以逐渐增加的等待时间重试发送，遵守 Telegram 在限流时给出的等待时间，并保留每个聊天内的消息顺序。
icon: material/refresh
---

# :material-refresh: 投递重试

如果机器人正在发送结果时与 Telegram 的连接短暂中断，消息不会丢失。机器人
会自行重试，每次等待时间略有增加，直到发送成功。

当 Telegram 本身要求机器人放慢速度时，机器人会精确等待 Telegram 要求的时
间后再重试，而不是自行猜测等待时间。

即使您在连接出问题时点按了多个按钮，发往同一聊天的消息仍会按您触发的顺序
到达。

如果中断持续太久，机器人最终会放弃该条结果的投递。聊天中不会出现任何内容，
也不会留下半发送的状态——这次尝试只会记录在服务器日志中，从不展示给您。

!!! tip "普通设置无需配置"

    默认值足以应付短暂的网络中断。可通过配置文件中的
    `delivery_retry_backoff`、`delivery_retry_backoff_max` 和
    `delivery_retry_ttl` 调整机器人持续重试的时长，见[配置文件](config-file.md)。

## 配置

关于 `delivery_retry_backoff`、`delivery_retry_backoff_max` 和
`delivery_retry_ttl`，请参见[配置](../configuration.md)。

## 相关

- [确认](confirmation.md) — 机器人在行动前等待的另一种情形
- [机器人如何连接](long-polling.md) — 这项功能保护的正是这条对外连接
- [控制输出](../output-control.md) — 结果如何送达您的聊天
