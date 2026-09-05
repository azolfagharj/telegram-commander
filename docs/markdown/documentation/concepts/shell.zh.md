---
title: Shell
description: 每个命令都通过 shell 运行，默认情况下为 /bin/bash，因此管道和重定向可以工作。在配置文件的根目录更改一次。
icon: material/console-line
---

# :material-console-line: Shell

用于运行命令的程序，默认为 `/bin/bash`。命令以
`bash -c "your command"`，因此管道和重定向工作。

您可以在[配置文件](config-file.md)中通过 `shell` 键全局更改 shell。

!!! example "选择 shell"

    ```yaml title="Root of config.yaml"
    shell: /bin/bash
    ```

每个[函数](function.md)都通过此 shell 运行命令。

## 配置

对于 `shell` 字段和相关根选项，请参见
[配置 → 根字段](../configuration.md#root-fields)。

## 相关

- [函数](function.md) — 生成 shell 要运行的命令
- [配置文件](config-file.md) — 其中设置了`shell`
