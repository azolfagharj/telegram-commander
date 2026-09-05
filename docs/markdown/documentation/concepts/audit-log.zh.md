---
title: 审计日志
description: 记录机器人运行的每个命令，包括谁点击了按钮、哪个按钮、退出代码以及花费了多长时间。
icon: material/history
---

# :material-history: 审计日志

独立的日志流会记录每次命令运行：由谁运行、使用了哪个[按钮](button.md)、
退出代码以及耗时。它在[配置文件](config-file.md)的 `logging` 下配置。

以后需要确认谁在凌晨三点重启了服务，或昨晚的备份是否确实完成时，
可以查看这份记录。由于它使用独立的日志流，即使普通应用日志轮转后，
仍可继续保留审计记录。

每次运行都会添加一行，并在它们之间回答四个问题：

| 问题 | 这条线告诉你什么 |
|---|---|
| 谁？ | 点击该按钮的 Telegram 账号 |
| 什么？ | 他们使用的按钮 |
| 有效吗？ | 命令返回的退出代码 |
| 多长时间？ | 命令完成所需的时间 |

!!! example "将审计日志写入独立文件"

    ```yaml title="logging section of config.yaml"
    logging:
      logs:
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

## 配置

有关完整的 `logging` 架构和支持的输出，请参阅
[配置 → logging](../configuration.md#logging)。

## 相关

- [配置文件](config-file.md) — 定义日志记录的位置
- [函数](function.md) — 点击按钮时运行的内容
