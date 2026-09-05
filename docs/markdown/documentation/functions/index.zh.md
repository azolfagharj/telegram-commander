---
title: 什么是函数？
description: 函数将按钮上写入的值转换为一个 shell 命令。了解按钮如何提供每个函数声明的参数。
---

# :material-function-variant: 什么是函数？

**函数**是一个配方，可以转换一些命名值（其
[parameters](../concepts/parameter.md)) 转换为 shell 命令。每个
菜单中的每个[按钮](../concepts/button.md)都通过 `function` 字段指定一个函数。

可以把函数看作填空式命令。“磁盘使用”函数有一个空位——路径——
您可以在每个按钮上填入不同的值。

## 点击按钮时会发生什么

1. 机器人查找按钮的 `function` 字段中指定的函数。
2. 它收集该按钮上写的值。
3. 它根据这些值构建一个 shell 命令。
4. 它在 [shell](../concepts/shell.md) 中运行命令并发送输出
   作为代码块返回聊天。

如果该函数不存在，或者缺少它所需的值，则机器人永远不会
开始：[`validate`](../cli.md#validate) 首先报告问题。

## 一个有效的例子

`command` 函数内置且始终可用。它会运行按钮 `command` 字段中的内容。

!!! example "使用内置 command 函数"

    ```yaml title="Uptime button"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

点击 **正常运行时间**，机器人在服务器上运行 `uptime` 并将输出发回。

## 从按钮传递值 { #passing-values-from-a-button }

将每个值直接写在按钮上。有两种方法可以做到这一点：

1. 使用快捷字段 `command`、`path` 和 `args`。每个填充
   同名参数。
2. 对于任何其他参数，使用其名称作为按钮上的键。

!!! example " 通过参数名称传递 URL"

    ```yaml title="Check API button"
    - name: Check API
      type: button
      function: curl-url
      url: "https://example.com/health"
    ```

这里 `url` 与 `curl-url` 声明的 `url` 参数匹配。相同的规则
适用于 `host`、`unit` 和 `lines` 等名称。

!!! warning "不要将值放入 `params:`" 中

    `params:` 属于自定义函数文件，它在其中声明值
    该函数接受。在按钮上直接写入每个值：

    ```yaml title="Values belong directly on the button"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

    数字 YAML 值不需要引号。

[`validate`](../cli.md#validate) 根据参数检查每个键
由所选函数声明。拼写错误或未声明的名称失败
验证。它还检查声明为 `int` 或 `bool` 的值。可选值
当按钮将它们排除在外时，请使用它们的默认值。

## 两种函数 { #two-kinds-of-function }

函数要么在程序内部提供，要么来自您保留的 YAML 文件
服务器。一旦按钮使用它们，它们的行为方式就相同。

| | 内置 | 自定义 |
|---|---|---|
| 它来自哪里 | 在程序内发货 | 您编写的 YAML 文件 |
| 你创建一个文件吗？ | 否 | 是的，每个函数一个文件 |
| 名称 | 保留（`command`、`script`） | 任何未保留的名称 |
| 随时可用？ | 是 | 仅当您设置 `function_directory` 时 |
| 可编辑吗？ | 否 | 是的，它们是您的文件 |

您可以在同一菜单中自由混合这两种类型。大多数菜单以 `command` 开头
仅按钮；当您发现自己重复以下操作时，请转向自定义函数
相同的命令，但有一点小小的改变。

### 内置函数 { #built-in-functions }

始终加载两个函数，即使您没有设置
`function_directory`。他们的 `command`、`path` 和 `args` 字段直接继续
一个按钮。

| 函数 | 作用 | 必填 | 可选 |
|----------|--------------|----------|----------|
| [`command`](built-in/command.md) | 按照编写的方式运行一个 shell 命令 | `command` | — |
| [`script`](built-in/script.md) | 运行带参数的脚本文件 | `path` | `args` |

两个名称都是**保留**。自定义函数文件可能无法重用它们：
加载程序因错误（例如 `function name "command" is reserved`）而停止，并且
机器人不启动。

### 自定义函数 { #custom-functions }

自定义函数是描述一个可重用命令的单个 YAML 文件。
将这些文件保存在您自己的文件夹中并指向
[`function_directory`](../configuration.md#function_directory-rules) 在它。

!!! example "告诉机器人您的函数文件所在的位置"

    ```yaml title="config.yaml"
    function_directory: "./functions"
    ```

机器人在启动时读取该文件夹（包括子文件夹）并选取
其中的每个 `.yaml` 和 `.yml` 文件。其他文件将被忽略。

发布存档已包含一个 `functions/` 文件夹，其中包含五个示例
您可以按原样使用：

| 函数 | 作用 | 按钮值 |
|----------|--------------|---------------|
| [回显脚本](release-pack/echo-script.md) | 通过 Bash 运行脚本 | `path`，可选`args` |
| [磁盘路径](release-pack/disk-path.md) | 显示磁盘使用情况 | 可选`path` |
| [Curl URL](release-pack/curl-url.md) | 获取一个 URL | `url` |
| [Ping 主机](release-pack/ping-host.md) | 对主机执行 ping 操作 | `host`，可选`count` |
| [Journal 单元](release-pack/journal-unit.md) | 显示最近的服务日志 | `unit`，可选`lines` |

要编写自己的内容，请从
[文件结构](write-your-own/file-structure.md)或遵循
[分步指南](write-your-own/step-by-step.md)。

!!! tip "检查加载了什么"

    ```bash title="List every function the bot can see"
    ./telegram-commander list-functions --config config.yaml
    ```

    内置函数显示 `source=builtin`；自定义函数会显示其来源文件。

## 安全注意事项

!!! warning "按钮使用机器人的权限运行"

    命令使用运行机器人的账号权限。如果该账号是 root（默认的
    [服务](../installation/run-as-a-service.md)设置），按钮就能在主机上
    执行任何操作。只能添加您信任的[允许用户](../configuration.md#telegram)。

    参数值会作为文本插入命令。应将其视为 shell 输入：优先使用按钮中的
    固定值，并为破坏性操作添加
    [`confirm: true`](../concepts/confirmation.md)。

!!! info "较长的输出会被截断和拆分"

    命令会在达到 `timeout` 时停止，机器人最多保留 `max_output_bytes`
    的输出。超过一条 Telegram 消息的内容会拆成多条消息。详见
    [配置 → 可查看的命令输出量](../configuration.md#how-much-command-output-you-see)。

## 相关

- [`command`](built-in/command.md) — 运行一个 shell 命令
- [`script`](built-in/script.md) — 运行脚本文件
- [分步指南](write-your-own/step-by-step.md) — 构建您的第一个自定义函数
- [菜单](../concepts/menu.md) — 按钮如何引用函数
- [参数](../concepts/parameter.md) — 函数需要的命名值
