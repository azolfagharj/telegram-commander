---
title: 允许的用户
description: 选择哪些 Telegram 帐户可以打开机器人菜单。通过数字用户 ID 或用户名添加人员，然后查看其他人会得到什么。
icon: material/shield-lock
---

# :material-shield-lock: 允许的用户

允许使用机器人的 Telegram 帐户列表，按数字 `user_id` 或
`@username`。其他人都被拒绝。

未获授权的用户会收到一条包含 `user_id` 和 `username` 的消息，以便向管理员
申请访问权限。首次使用时也可以通过这种方式找到自己的 ID，详见
[在 CLI 中运行 → 第 5 步](../installation/download-and-run.md#step-5-find-your-user-id-if-needed)。

添加新用户时，让对方向机器人发送任意消息，复制机器人回复的 ID，
将该 ID 加入 `allowed_users`，然后重启机器人。对方下次发送消息时
就会看到菜单，而不是拒绝提示。

!!! warning "列表中的所有用户共享同一菜单"

    不支持按用户设置权限。您添加的任何人都可以点击所有已定义的按钮，
    因此只能加入您信任其操作这台服务器的用户。

!!! example "允许使用机器人的两个帐户"

    ```yaml title="telegram section of config.yaml"
    telegram:
      allowed_users:
        - "123456789"
        - "@alice"
    ```

## 配置

对于 `allowed_users` 和其他 Telegram 设置，请参阅
[配置 → telegram](../configuration.md#telegram)。

## 相关

- [配置文件](config-file.md) — 保存`telegram`部分
- [机器人如何连接](long-polling.md) — 没有要打开的入站端口
