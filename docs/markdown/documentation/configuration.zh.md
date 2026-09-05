---
icon: material/file-cog-outline
title: 配置
description: Telegram Commander 的全部设置及其类型、默认值和含义，包括 Telegram、菜单、函数目录、超时、输出限制和日志。
---

# :material-file-cog-outline: 配置

[配置文件](concepts/config-file.md)描述整个机器人：Telegram
连接、允许使用的用户、[按钮](concepts/button.md)菜单和日志。
您可以使用 `--config` 将其传递给读取它的命令 - `run`、`validate`、
`fmt` 和 `list-functions`（请参见 [CLI](cli.md)）。

所有键都使用 `lower_snake_case`。**未知键会被拒绝**，因此拼写错误会在
[验证](cli.md#validate)时立即显示。

**必填** 表示如果此字段丢失或为空，则验证失败
应用默认值。  
**可选**字段可以省略；默认列显示当时使用的内容。

刚接触本项目？请从[在 CLI 中运行](installation/download-and-run.md)开始，
其中会逐步创建第一份配置。下文术语详见[概念](concepts/config-file.md)。

## :material-rocket-launch-outline: 最小配置 { #a-minimal-config }

仅`telegram`（带有一个令牌和一个[允许的用户](concepts/allowed-users.md)）
需要 `menu`。其他一切都有默认值：

!!! example " 以一名允许的用户和一个按钮启动"

    ```yaml title="config.yaml (minimal)"
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

该版本中的 `config-examples/` 文件夹包括最小的和完整的
示例。

## :material-card-bulleted-outline: 根字段 { #root-fields }

| 字段 | 类型 | 必填 | 默认值 | 描述 |
|-------|------|----------|---------|-------------|
| `telegram` | 对象 | 是 | — | Telegram 设置（见下文） |
| `menu` | 列表 | 是 | — | 菜单树；至少一个节点 |
| `function_directory` | 字符串 | 否 | 未设置 | 自定义函数YAML目录（见下面的规则） |
| `shell` | 字符串 | 否 | `/bin/bash` | [Shell](concepts/shell.md) 用作 `shell -c "<command>"` |
| `timeout` | 时长 | 否 | `60s` | 默认命令超时 |
| `max_output_bytes` | 整数 | 否 | `524288` | 每个命令保留的最大输出（请参阅[可查看的命令输出量](#how-much-command-output-you-see)） |
| `workdir` | 字符串 | 否 | 进程cwd | 命令的默认工作目录 |
| `env` | 映射 | 否 | 空 | 命令的额外环境变量 |
| `menu_columns` | 整数 | 否 | `2` | 消息框下方每行的项目按钮 |
| `page_size` | 整数 | 否 | `8` | 分页前每页的项目数 |
| `confirm_ttl` | 时长 | 否 | `5m` | [confirmation](concepts/confirmation.md) 提示保持有效的时间有多长 |
| `enable_run_command` | 布尔值 | 否 | `false` | 显示 **`$ >_ Run Command`** 按钮，将下一条消息作为 shell 命令运行。默认关闭。任何可使用机器人的人都能在主机上运行任意命令，因此只能在信任所有允许用户时启用。此键放在 `telegram` 下无效。 |
| `logging` | 对象 | 否 | 内置默认记录器 | 命名记录器（见下文） |

??? note " 如果我省略 `shell` 会怎样？"

    你可以省略它。该机器人使用 `/bin/bash`。与 `timeout`、`page_size` 相同，
    和其他可选根字段：忽略它们并应用默认值。你只需要
    当您想要非默认值时设置它们（例如
    `shell: /bin/sh`）。

### 您看到多少命令输出 { #how-much-command-output-you-see }

两个限制依次适用。 `max_output_bytes` 是 **您的** 限制，并且
超出了您无法更改的 Telegram 限制。

**1.您的限制：`max_output_bytes`**（默认 `524288`，因此 512 KB）

当命令运行时，机器人最多保留这么多的输出（计算）
分别用于正常输出和错误输出。过去的一切都被丢弃，
但命令本身会继续运行，直到完成或达到 `timeout`。
发生这种情况时，结果以 `(output truncated)` 开头。

**2. Telegram 的限制：一条消息最多可容纳 4096 字节**

这个问题已由 Telegram 修复。如果结果比一条消息长，
机器人将其分成几条消息。每个部分都作为对
在它之前分开，因此它们保持在一起并按顺序排列，并且菜单按钮出现
在最后一部分。只要有可能，分割就会发生在行边界上，所以
线条没有被切成两半。

如果分割后结果仍然很长，机器人会在 10 后停止
消息，最后一条以注释结尾，例如
`(output too long; showing first N bytes)`，其中`N`是输出的多少
你确实收到了。

因此，提高 `max_output_bytes` 可以让机器人保持更多输出，但你仍然会看到
最多大约十条消息。对于那么长的输出，通常最好
缩短命令本身（例如 `journalctl -u nginx | tail -n 50`）或
将完整输出写入服务器上的文件。

### `function_directory` 规则 { #function_directory-rules }

| 情况 | 结果 |
|-----------|--------|
| 缺少键 | 信息日志；仅内置函数 |
| 键存在但为空 (`""`) | 信息日志；仅内置函数 |
| 键设置为不存在或不可访问的路径 | 硬错误；进程停止 |
| 路径存在但目录为空 | 好的 |

!!! warning "错误路径会阻止机器人启动"

    如果`function_directory`指向的文件夹不存在或无法访问
    阅读，程序因错误而停止，而不是在没有您的情况下启动
    自定义函数。

## :material-send-circle-outline: `telegram` { #telegram }

| 字段 | 类型 | 必填 | 默认值 | 描述 |
|-------|------|----------|---------|-------------|
| `bot_token` | 字符串 | 是 | — | 来自 BotFather 的 Bot 令牌 |
| `allowed_users` | 字符串列表 | 是 | — | [允许的用户](concepts/allowed-users.md) |
| `api` | 字符串 | 否 | `https://api.telegram.org` | 机器人 API 基本 URL |
| `proxy.enabled` | 布尔值 | 否 | `false` | 使用 Telegram API 代理 |
| `proxy.url` | 字符串 | 条件必填 | — | `proxy.enabled` 为 `true` 时必填 |
| `insecure` | 布尔值 | 否 | `false` | 跳过 TLS 验证（不推荐） |

未经授权的用户会收到一条包含 `user_id` 和 `username` 的消息，以便他们可以向管理员请求访问权限。这也是您第一次找到自己的 id 的方法 — 请参阅 CLI 中的 [Run → 步骤 5](installation/download-and-run.md#step-5-find-your-user-id-if-needed)。

!!! example " 通过代理连接"

    ```yaml title="telegram section with a proxy"
    telegram:
      bot_token: "123456789:AAExampleTokenValue"
      allowed_users:
        - "123456789"        # 用户数字 id
        - "@alice"           # 或用户名
      proxy:
        enabled: true
        url: "socks5://127.0.0.1:10808"
    ```

要让允许的用户从 Telegram 键入 shell 命令，请将其设置在文件的 **root** 下（而不是在 `telegram` 下）：

!!! tip "添加一根设置"

    ```yaml title="Enable the Run Command button"
    enable_run_command: true
    ```

## :material-menu: 菜单 { #menu }

本节提供字段参考。带示例的分步说明详见[菜单](concepts/menu.md)。
每个[按钮](concepts/button.md)或[分类](concepts/category.md)节点：

| 字段 | 类型 | 必填 | 描述 |
|-------|------|----------|-------------|
| `name` | 字符串 | 是 | 显示名称（兄弟姐妹中唯一，不区分大小写） |
| `type` | `category` \| `button` | 是 | 节点种类 |
| `items` | 列表 | 是，如果 `category` | 子节点；分类必须至少包含一个 |
| `function` | 字符串 | 是，如果 `button` | [函数](concepts/function.md)名称 |
| `command` | 字符串 | 是，如果 `function: command` | 内置`command`的Shell命令 |
| `path` | 字符串 | 是，如果 `function: script` | 内置`script`的脚本路径 |
| `icon` | 字符串 | 否 | 可选的表情符号前缀 |
| `id` | 字符串 | 否 | 该节点的可选 ID。你可以省略它 |
| `confirm` | 布尔值 | 否 | 运行前询问[确认](concepts/confirmation.md)（默认`false`） |
| `timeout` | 时长 | 否 | 覆盖全局超时 |
| `workdir` | 字符串 | 否 | 覆盖工作目录 |
| `env` | 映射 | 否 | 此按钮的额外环境 |
| `columns` | 整数 | 否 | 覆盖此类别的列 |
| `args` | 字符串 | 否 | `script` 的可选参数 |
| 任何声明的参数名称 | 标量 | 正如函数所声明的 | 传递给所选函数的值，例如 `url`、`host`、`unit` 或 `lines` |

在**按钮**上，任何其他标量键都被视为函数参数。
其名称必须与所选函数声明的参数匹配。未知
参数名称失败 [`validate`](cli.md#validate)。声明为 `int` 的值或
`bool` 也已检查。字符串、数字和布尔值可以直接写入
作为 YAML 值；数字不需要引号。

在**分类**上，上述分类字段之外的任何键都是错误的。
分类不运行函数，因此不能包含参数键。

`command`、`path` 和 `args` 是用以下参数填充参数的快捷字段
相同的名字。其他[参数](concepts/parameter.md)名称直接写
在按钮上。不要将按钮值放置在嵌套的 `params:` 映射内。参见
[函数 → 从按钮传递值](functions/index.md#passing-values-from-a-button)。

## :material-math-log: `logging` { #logging }

可选。如果省略，则使用 `stderr` 上 `info` 上的默认控制台记录器。

命名记录器：

!!! example "写入正常日志和审计文件"

    ```yaml title="logging section with an audit file"
    logging:
      logs:
        default:
          level: info
          format: console   # 或 JSON
          output:
            - output: stderr
        audit:
          level: info
          format: json
          output:
            - output: file
              file: /var/log/telegram-commander/audit.log
    ```

支持的输出：`stdout`、`stderr`、`file`、`discard`。

上面显示的 `audit` 日志记录器会记录每次命令运行（用户、按钮、
退出代码和耗时）。详见[审计日志](concepts/audit-log.md)。

## 相关页面

- [在 CLI](installation/download-and-run.md) 中运行 — 构建并运行第一个配置
- [菜单](concepts/menu.md) — 深入了解菜单树
- [函数](functions/index.md) — `function`、`command`、`path` 和 `args` 的含义
- [CLI](cli.md) — 验证并使用您的配置运行
