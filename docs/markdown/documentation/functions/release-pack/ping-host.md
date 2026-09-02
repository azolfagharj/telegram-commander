---
title: Ping Host
description: The bundled ping-host function pings a host a few times. Set the host and optional count directly on the button.
icon: material/access-point-network
---

# :material-access-point-network: Ping Host

`ping-host` pings a host a few times and sends the result back. It is one of the
[bundled functions](../index.md#custom-functions) you can use from a button as
it is.

- Runs: `ping -c {{.count}} {{.host}}`
- `host` (required): hostname or IP
- `count` (optional, default `4`): number of packets

!!! example "The function file"

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

## Add a button

!!! example "Ping a fixed host three times"

    ```yaml title="Ping gateway button"
    - name: Ping gateway
      type: button
      function: ping-host
      host: "192.168.1.1"
      count: 3
    ```

`count` is numeric, so it does not need quotes. You can also leave it out to use
the default value `4`:

!!! example "Use the default count"

    ```yaml title="Ping DNS button"
    - name: Ping DNS
      type: button
      function: ping-host
      host: "1.1.1.1"
    ```

## Related

- [Rules](../write-your-own/rules.md) — how defaults and required values behave
- [Custom functions](../index.md#custom-functions) — all five bundled examples
