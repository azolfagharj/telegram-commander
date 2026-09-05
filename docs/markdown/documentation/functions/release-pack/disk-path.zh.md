---
title: 磁盘路径
description: 随附的 disk-path 函数通过 df -h 显示磁盘使用情况。唯一参数是 path，省略时默认为根文件系统。
icon: material/harddisk
---

# :material-harddisk: 磁盘路径

`disk-path` 使用 `df -h` 显示指定路径的磁盘使用情况。它是可直接用于按钮的
[随附函数](../index.md#custom-functions)之一。

- 运行：`df -h {{.path}}`
- `path`（可选，默认`/`）：要检查的文件系统路径

!!! example "函数文件"

    ```yaml title="functions/disk-path.yaml"
    name: disk-path
    run: "df -h {{.path}}"
    params:
      - name: path
        type: string
        required: false
        default: "/"
        description: Filesystem path to check
    ```

由于 `path` 可选且有默认值，按钮省略它后仍可正常工作。

!!! example " 使用默认路径或选择一个"

    ```yaml title="Disk usage buttons"
    - name: Disk (root)
      type: button
      function: disk-path        # 没有 path 字段：使用默认值 "/"
    - name: Disk (var)
      type: button
      function: disk-path
      path: "/var"
    ```

第一个按钮运行 `df -h /`，第二个按钮运行 `df -h /var`。

## 相关

- [规则](../write-your-own/rules.md) — 默认值和必填值的行为方式
- [自定义函数](../index.md#custom-functions) — 全部五个随附示例
- [`command`](../built-in/command.md) — 改为编写完整命令
