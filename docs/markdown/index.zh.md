---
title: Telegram Commander
description: 将 YAML 文件变成 Telegram 机器人，轻点一次即可在 Linux 服务器上运行命令。
icon: material/cellphone-link
hide:
  - navigation
  - toc
---

# :material-cellphone-link: 通过 Telegram 控制 Linux 服务器

<div class="hero" markdown>
<div class="hero-art" markdown>
![Telegram Commander](/images/logo-large.png){ .off-glb width="230" }
</div>
<div class="hero-text" markdown>
**轻点一下即可在服务器上运行命令，并将输出发回聊天。**

Telegram Commander 将一个简单的 YAML [配置文件](documentation/concepts/config-file.md)
变成带有[按钮](documentation/concepts/button.md)菜单的 Telegram 机器人。为按钮设置任意
终端命令，即可通过手机运行，无需编写代码。

<p class="hero-lang">用其他语言阅读本页：
<a href="/" hreflang="en" class="hero-lang__link">English</a>
<a href="/de/" hreflang="de" class="hero-lang__link">Deutsch</a>
<a href="/fr/" hreflang="fr" class="hero-lang__link">Français</a>
<a href="/es/" hreflang="es" class="hero-lang__link">Español</a>
<a href="/ru/" hreflang="ru" class="hero-lang__link">Русский</a>
<a href="/fa/" hreflang="fa" class="hero-lang__link">فارسی</a>
</p>
</div>
</div>

<div style="text-align: center" markdown="span">
[开始使用 :material-arrow-right:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[查看配置 :material-file-code-outline:](documentation/configuration.md#a-minimal-config){ .md-button }
</div>

[安装](documentation/installation/download-and-run.md) ·
[概念](documentation/concepts/config-file.md) ·
[函数](documentation/functions/index.md) ·
[配置](documentation/configuration.md) ·
[CLI](documentation/cli.md)

## :material-image-multiple-outline:{ .shots } 截图 { .split }

以下图片展示菜单、正在运行的命令、返回的输出，以及手动输入命令。
点击任意图片可查看大图。

<div style="text-align: center" markdown="span">
  ![打开系统菜单时的命令输出](/images/01.jpeg){ width="140" loading=lazy }
  ![资源和进程按钮](/images/02.jpeg){ width="140" loading=lazy }
  ![存储和软件包按钮](/images/03.jpeg){ width="140" loading=lazy }
  ![网络工具和手动输入命令](/images/04.jpeg){ width="140" loading=lazy }
  ![命令输出发送回聊天](/images/05.jpeg){ width="140" loading=lazy }
</div>

## :material-lightning-bolt:{ .bolt } 快速简单 { .split }

<div class="grid cards cols-3 center-title step-cards" markdown>

-   :material-file-document-outline:{ .middle } __编写菜单__

    ---

    :material-numeric-1-circle:{ .step } 列出按钮及其命令。

-   :material-rocket-launch:{ .middle } __启动机器人__

    ---

    :material-numeric-2-circle:{ .step } 立即运行，或将其设置为服务。

-   :material-gesture-tap-button:{ .middle } __点击并查看__

    ---

    :material-numeric-3-circle:{ .step } 点击按钮，在聊天中查看输出。

</div>

<div style="text-align: center" markdown="span">
[立即开始 :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
</div>

## :material-view-grid-outline:{ .grid-icon } 用例 { .split }

<div class="grid cards cols-4 icon-left" markdown>

- :material-restart:{ .lg } 重新启动或停止服务
- :material-docker:{ .lg } 启动和停止容器
- :material-package-down:{ .lg } 更新系统包
- :material-text-box-search-outline:{ .lg } 查看日志
- :material-harddisk:{ .lg } 检查磁盘空间
- :material-memory:{ .lg } 监控 CPU 和内存
- :material-access-point-network:{ .lg } Ping 主机和测试 URL
- :material-backup-restore:{ .lg } 进行和恢复备份
- :material-script-text:{ .lg } 运行您自己的脚本
- :material-power:{ .lg } 重新启动或关闭主机
- :material-console:{ .lg } 手动输入任何命令
- :material-all-inclusive:{ .lg } 以及几乎所有其他内容

</div>

## :material-thumb-up-outline:{ .thumb } 为什么使用它 { .split }

<div class="grid cards cols-4 center-title" markdown>

- :material-clock-fast:{ .lg .middle } __无编码__

    ---

    在一个 YAML 文件中描述菜单和命令。

    [:octicons-arrow-right-24: 配置文件](documentation/concepts/config-file.md)

- :material-cellphone-link:{ .lg .middle } __从任何地方__

    ---

    在手机上打开 Telegram 即可操作服务器，无需通过 VPN 连接主机。

    [:octicons-arrow-right-24: 机器人如何连接](documentation/concepts/long-polling.md)

- :material-lan-disconnect:{ .lg .middle } __无开放端口__

    ---

    该机器人连接到 Telegram。没有任何内容暴露在互联网上。

    [:octicons-arrow-right-24: 机器人如何连接](documentation/concepts/long-polling.md)

- :material-message-text-outline:{ .lg .middle } __聊天中输出__

    ---

    结果以消息形式返回。您不需要 SSH 会话。

    [:octicons-arrow-right-24: 您看到多少输出](documentation/configuration.md#how-much-command-output-you-see)

- :material-shield-lock:{ .lg .middle } __受控并记录__

    ---

    选择谁获得菜单，确认有风险的操作，并记录每次运行。

    [:octicons-arrow-right-24: 访问和确认](documentation/concepts/allowed-users.md)

- :material-folder-outline:{ .lg .middle } __嵌套菜单__

    ---

    将按钮分到不同分类中。“首页”始终位于顶部，“返回”可回到上一级。

    [:octicons-arrow-right-24: 菜单](documentation/concepts/menu.md)

- :material-function-variant:{ .lg .middle } __可重用函数__

    ---

    编写一次命令，然后在每个按钮上填写不同的值。

    [:octicons-arrow-right-24: 函数](documentation/functions/index.md)

- :material-cog-play-outline:{ .lg .middle } __保持运行__

    ---

    将其安装为服务，机器人随主机启动。

    [:octicons-arrow-right-24: 作为服务运行](documentation/installation/run-as-a-service.md)

</div>

## :material-file-code-outline:{ .code-icon } 一个小例子 { .split }

此配置创建了一个带有一个按钮的机器人，名为 "Uptime"。点击它运行
服务器上的 `uptime` 命令。

!!! example "此完整配置添加一个按钮"

    ```yaml title="config.yaml"
    telegram:
      bot_token: "YOUR_BOT_TOKEN" # (1)!
      allowed_users:
        - "YOUR_USER_ID" # (2)!

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime" # (3)!
    ```

    1. 在 Telegram 中通过 BotFather 创建机器人并获取令牌。
    2. 只有此处列出的帐户才能打开菜单。您可以使用数字 ID
        或 `@username`。
    3. 您可以在终端中输入的任何内容都会显示在此处。

这是一个完整的、可工作的配置。其他一切都是可选的。

!!! tip "专为小型、值得信赖的群体而设计"

    没有任何内容侦听传入连接，并且仅侦听中的帐户
    `allowed_users` 获取菜单。每个可以使用该机器人的人都可以运行
    您定义的按钮，因此请保持该列表简短。

## :material-hand-pointing-right: 准备好尝试了吗？ { .split }

<div style="text-align: center" markdown>
[立即开始 :material-rocket-launch-outline:](documentation/installation/download-and-run.md){ .md-button .md-button--primary }
[概念 :material-book-open-variant:](documentation/concepts/config-file.md){ .md-button }
[最新版本 :material-download:](https://github.com/azolfagharj/telegram-commander/releases/latest){ .md-button .md-button--primary }

[浏览源码 :fontawesome-brands-github:](https://github.com/azolfagharj/telegram-commander){ .md-button }
</div>

Telegram Commander 是免费且开源的。如果可以节省你的时间，
[考虑支持其发展](https://azolfagharj.github.io/donate/) —
它有助于保持项目的活力和维护。
