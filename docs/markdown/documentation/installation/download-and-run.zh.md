---
title: 在 CLI 中运行
description: 让您的机器人从终端逐步运行，从下载版本并编写一个小配置文件到点击第一个按钮。
icon: material/console
---

# :material-console: 在 CLI 中运行

本页将从头带您启动机器人，无需事先了解本项目。如果遇到不熟悉的术语，
请查看[概念](../concepts/config-file.md)页面。

## :material-clipboard-check-outline: 开始之前

您需要 Telegram 提供两件事：

1. **机器人令牌。** 打开与 [@BotFather](https://t.me/BotFather) 的聊天，
   发送 `/newbot` 并按提示操作，然后复制获得的令牌，例如
   `123456789:AAExampleTokenValue`。
2. **数字用户 ID。** 这是一个数字，不是 `@username`。如果暂时不知道，
   首次向机器人发送消息时，它会告知您的 ID（见[第 5 步](#step-5-find-your-user-id-if-needed)）。

## :material-download: 第1步：下载

下载发布存档并解压。

!!! example "下载并打开发布文件夹"

    ```bash title="Download and extract the release"
    wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
    tar -xzf telegram-commander.tar.gz
    cd telegram-commander
    ```

在文件夹内您会发现：

- `telegram-commander-linux-amd64` 和 `telegram-commander-linux-arm64` — 程序，每种 CPU 类型一个
- `config-examples/` — 现成的[配置文件](../concepts/config-file.md)（参见[配置](../configuration.md)）
- `functions/` — 自定义[函数](../concepts/function.md)示例（参见[自定义函数](../functions/index.md#custom-functions)）

## :material-chip: 第 2 步：选择您的二进制文件

!!! info "哪个文件适合您的机器？"

    大多数服务器和 PC 都是 `amd64`（也称为 `x86_64`）。小型 ARM 板
    部分云虚拟机为`arm64`。

    如果不确定，请运行 `uname -m`：`x86_64` 表示 amd64，`aarch64` 表示
    arm64。

=== ":fontawesome-brands-linux: AMD64"

    ```bash title="Keep the amd64 program"
    mv telegram-commander-linux-amd64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-arm64
    ```

=== ":fontawesome-brands-linux: ARM64"

    ```bash title="Keep the arm64 program"
    mv telegram-commander-linux-arm64 telegram-commander
    chmod +x telegram-commander
    rm telegram-commander-linux-amd64
    ```

现在您有一个名为 `telegram-commander` 的程序。

## :material-file-cog-outline: 第 3 步：创建配置

将最小示例复制到工作文件中：

!!! example "制作可编辑配置"

    ```bash title="Copy the example config"
    cp config-examples/config.minimal.yaml ./config.yaml
    ```

打开 `config.yaml` 并替换两个占位符：

- `YOUR_BOT_TOKEN` — 来自 BotFather 的令牌
- `YOUR_USER_ID` — 您的数字 ID（或暂时保留它并参阅步骤 5）

每项设置的含义详见[配置](../configuration.md)。

## :material-file-check-outline: 第 4 步：验证

运行前务必检查配置。这可以捕获拼写错误和错误，而无需
启动机器人。

!!! success "检查配置是否有效"

    ```bash title="Validate the config"
    ./telegram-commander validate --config config.yaml
    ```

如果输出 `Valid configuration`，说明配置有效；否则会准确列出错误及其位置。
详情请参阅 [CLI](../cli.md#validate)。

## :material-account-search: 第 5 步：找到您的用户 ID（如果需要） { #step-5-find-your-user-id-if-needed }

如果您不知道您的用户 ID，请仅在 `config.yaml` 中设置令牌，将任何
现在 `allowed_users` 中的编号，然后运行机器人：

!!! info "启动一次即可查看您的用户id"

    ```bash title="Run the bot to learn your id"
    ./telegram-commander run --config config.yaml
    ```

打开 Telegram，找到机器人并发送任意消息。由于您尚未加入
[允许的用户](../concepts/allowed-users.md)，机器人会回复
`user_id` 和 `username`。将该 ID 复制到 `allowed_users`，停止机器人
`Ctrl+C`，然后再次运行。

这是访问控制的一部分，详见[配置 → telegram](../configuration.md#telegram)。

## :material-play-circle-outline: 第6步：运行

!!! example " 在终端中启动机器人"

    ```bash title="Start the bot"
    ./telegram-commander run --config config.yaml
    ```

在 Telegram 中打开您的机器人并发送 `/start`。你应该看到你的菜单。点击一个
[button](../concepts/button.md) 运行其命令。

!!! success "您的机器人已上线"

    您在 `config.yaml` 中描述的菜单现在位于您的聊天中，并且每次点击
    在这台机器上运行它的命令。

要在注销服务器后保持机器人运行，请将其设置为服务。
请参阅[作为服务运行](run-as-a-service.md)。

## :material-map-marker-path: 接下来是什么

- 添加更多[按钮](../concepts/button.md)和[类别](../concepts/category.md)：[菜单](../concepts/menu.md)
- 了解实际运行的内容：[函数](../functions/index.md)
- 查看每个命令行选项：[CLI](../cli.md)
