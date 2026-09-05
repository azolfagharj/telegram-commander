---
title: CLI 参考
description: Telegram Commander 的全部命令行选项，包括运行机器人、检查和格式化配置文件，以及这些命令使用的 --config 标志。
icon: material/console-line
---

# :material-console-line: CLI 参考

!!! info "命令格式"

    ```text title="Command syntax"
    telegram-commander <command> [flags]
    ```

对于需要加载配置的命令，必须提供 `--config` / `-c`。路径可以是相对于
当前工作目录的路径，也可以是绝对路径。文件内容详见[配置](configuration.md)。

如果您刚开始使用，请参阅 [在 CLI 中运行](installation/download-and-run.md)，
其中按实际使用顺序介绍这些命令。

## :material-format-list-bulleted-square: 命令

### `run` { #run }

在前台运行机器人。

!!! info "在前台启动机器人"

    ```bash title="Run the bot"
    telegram-commander run --config /path/to/configfile.yaml
    ```

!!! note "配置更改需要重新启动"

    机器人在启动时读取一次配置。编辑文件后，重新启动
    进程（例如 `systemctl restart telegram-commander`）。

### `validate` { #validate }

离线验证配置、函数和按钮引用。

!!! info "离线或在线检查配置"

    ```bash title="Validate the config"
    telegram-commander validate --config /path/to/configfile.yaml
    telegram-commander validate --config /path/to/configfile.yaml --online
    ```

!!! note "`--online` 需要互联网"

    对于 `--online`，检查还会询问 Telegram 机器人令牌是否有效，
    所以机器必须能够连接到 Telegram。

### `version`

打印程序版本。

### `fmt`

格式化 YAML 配置文件。

!!! info "输出或保存格式化后的 YAML"

    ```bash title="Format a config file"
    telegram-commander fmt --config /path/to/configfile.yaml
    telegram-commander fmt --config /path/to/configfile.yaml -w
    ```

### `environ`

打印进程环境变量（对于调试服务单元很有用）。

### `list-functions` { #list-functions }

列出内置函数和已加载的自定义函数。可用它确认自定义函数文件是否已找到。
详见[函数](functions/index.md)。

!!! info "显示所有可用函数"

    ```bash title="List available functions"
    telegram-commander list-functions --config /path/to/configfile.yaml
    ```

### `completion`

生成 shell 补全脚本：

!!! info "选择使用的 shell"

    ```bash title="Generate a completion script"
    telegram-commander completion bash
    telegram-commander completion zsh
    telegram-commander completion fish
    telegram-commander completion powershell
    ```

### `manpage`

将手册页写入标准输出。

## 相关页面

- [配置](configuration.md) — 传递给 `--config` 的文件
- [自定义函数](functions/index.md#custom-functions) — `list-functions` 显示的内容
- [作为服务运行](installation/run-as-a-service.md) — 通过 systemd 运行 `run`
