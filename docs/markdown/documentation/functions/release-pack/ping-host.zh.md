---
title: Ping 主机
description: 随附的 ping-host 函数会多次 ping 主机。直接在按钮上设置主机和可选次数。
icon: material/access-point-network
---

# :material-access-point-network: Ping 主机

`ping-host` 会对主机执行数次 ping，并将结果发回。它是可直接用于按钮的
[随附函数](../index.md#custom-functions)之一。

- 运行：`ping -c {{.count}} {{.host}}`
- `host`（必填）：主机名或IP
- `count`（可选，默认`4`）：数据包数量

!!! example "函数文件"

    ```yaml title="functions/ping-host.yaml"
    name: ping-host
    run: "ping -c {{.count}} {{.host}}"
    params:
      - name: host
        type: string
        required: true
        description: Hostname or IP
      - name: count
        type: string
        required: false
        default: "4"
        description: Number of ping packets
    ```

## 添加一个按钮

!!! example "Ping固定主机3次"

    ```yaml title="Ping gateway button"
    - name: Ping gateway
      type: button
      function: ping-host
      host: "192.168.1.1"
      count: 3
    ```

`count` 是数字，因此不需要引号。也可以省略它以使用默认值 `4`：

!!! example "使用默认计数"

    ```yaml title="Ping DNS button"
    - name: Ping DNS
      type: button
      function: ping-host
      host: "1.1.1.1"
    ```

## 相关

- [规则](../write-your-own/rules.md) — 默认值和必填值的行为方式
- [自定义函数](../index.md#custom-functions) — 全部五个随附示例
