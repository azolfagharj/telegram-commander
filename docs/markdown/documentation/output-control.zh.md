---
title: 控制输出
description: 决定命令结果如何到达聊天。在根级设置 output、max_output_messages 和 max_output_bytes，在单个按钮上覆盖 output，并查看根级与按钮的每一种组合及示例。
icon: material/export-variant
---

# :material-export-variant: 控制输出

[按钮](concepts/button.md)运行结束后，机器人需要把结果送到您的聊天里。
Telegram 单条消息最多只接受 4096 字节，所以较长的结果要么以**多条回复消息**
到达，要么作为**一个 `.txt` 文件**送来，您可以打开并下载它。

有三个可选字段决定结果如何送达。您可以一个都不写，机器人依然会有合理的
表现。

!!! info "只需记住一条规则"

    按钮优先。如果按钮设置了 `output`，就按它执行。如果按钮没写，就用根级
    的值。如果根级也没写，机器人使用 `auto`。

<div class="grid cards cols-3" markdown>

-   :material-scissors-cutting:{ .middle } __`max_output_bytes`__

    ---

    在发送任何内容之前，从命令保留多少输出。

    [:octicons-arrow-right-24: 了解更多](#max_output_bytes)

-   :material-tune:{ .middle } __`output`__

    ---

    消息还是文件。唯一也能写在按钮上的字段。

    [:octicons-arrow-right-24: 了解更多](#output)

-   :material-counter:{ .middle } __`max_output_messages`__

    ---

    在 `auto` 模式下，发送文件之前允许多少条消息。

    [:octicons-arrow-right-24: 了解更多](#max_output_messages)

</div>

!!! tip "本页每张图片都可以放大"

    聊天截图特意做得比较小。点击其中一张即可查看原始大小。

## :material-card-bulleted-outline: 三个字段一览

| 字段 | 写在哪里 | 类型 | 默认值 | 决定什么 |
|-------|------------------|------|---------|-----------------|
| `max_output_bytes` | 仅根级 | 整数 | `524288` | 机器人为每个命令保留多少字节的输出 |
| `output` | 根级**和**按钮 | `auto` \| `text` \| `file` | `auto` | 结果以聊天消息到达，还是作为一个 `.txt` 文件 |
| `max_output_messages` | 仅根级 | 整数，`1`–`10` | `2` | 在 `auto` 模式下，改为发送文件之前允许的消息条数 |

这三个字段都写在[配置文件](concepts/config-file.md)的**根级**，与 `shell` 和
`timeout` 并列，而不是放在 `telegram` 下面。完整的根级键列表见
[配置](configuration.md#root-fields)页面。

## :material-scissors-cutting: `max_output_bytes`

这是您自己的限制，并且最先生效。命令运行时，机器人最多保留这么多字节的
输出，普通输出与错误输出分别计算。超出的部分会被丢弃，但命令本身会继续
运行，直到结束或碰到它的 `timeout`。

!!! example "每个命令最多保留 2 MB 输出"

    ```yaml title="config.yaml"
    max_output_bytes: 2097152 # (1)!
    ```

    1.  2 MB，而不是默认的 512 KB。普通输出与错误输出分别计算。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    达到限制时，结果会在摘要下方单独一行说明这一点。`(output truncated)`
    上面的部分是机器人保留下来的内容；其余日志从未离开服务器。

    </div>
    <div class="result-shot" markdown>

    ![一条提示输出被截断的 Telegram 结果](/images/output-control/truncated.png){ .shot loading=lazy }

    </div>
    </div>

!!! tip "只有发送文件时，调大它才有意义"

    一条聊天消息永远装不下超过 4096 字节的内容，所以更大的
    `max_output_bytes` 只有在结果以 `.txt` 文件到达时才能完整送到您手里。
    请与 `output: auto`（默认）或 `output: file` 搭配使用。

## :material-tune: `output`

`output` 选择送达方式。它接受三个取值，也是唯一可以写在单个按钮上的输出
字段。

=== "auto"

    结果较短时用文本消息，较长时用一个 `.txt` 文件。这是默认行为，也是您
    什么都不写时保留的取值。

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "让机器人为每个结果自行判断"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: auto # (1)!
        max_output_messages: 2 # (2)!
        ```

        1.  按每个结果分别决定，而不是强制一种送达方式。
        2.  切换点。两条消息以内留在聊天里；更长的内容变成一个文件。

    </div>
    <div class="result-shot" markdown>

    ![auto 模式把短输出留作文本，把较长输出作为文件发送](/images/output-control/auto.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

=== "text"

    无论结果多长，始终发聊天消息。机器人在十条消息后停止，并附上一条说明
    其余内容已被截断。

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "把所有内容都留在聊天里"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: text # (1)!
        ```

        1.  这里会忽略 `max_output_messages`，因为机器人永远不会自行切换到
            文件。

    </div>
    <div class="result-shot" markdown>

    ![以 Telegram 消息送达的命令结果](/images/output-control/message.png){ .shot loading=lazy }

    </div>
    </div>

=== "file"

    始终一个 `.txt` 文件，哪怕结果只有一行。摘要会作为文件说明一同送来。

    <div class="result" markdown>
    <div class="result-text" markdown>

    !!! example "把每个结果都作为附件发送"

        ```yaml title="config.yaml"
        shell: /bin/bash
        timeout: 60s
        output: file # (1)!
        ```

        1.  现在每个按钮都会发送附件，包括那些很快的按钮。如果这样太多，
            请改为写在单个按钮上。

    </div>
    <div class="result-shot" markdown>

    ![以 Telegram 文件送达的命令结果](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

!!! note "省略它等同于写 `auto`"

    您从来不必写 `output`。在这个字段出现之前就能用的配置，行为完全不变。

## :material-counter: `max_output_messages`

这个字段只在 `auto` 模式下起作用，而且只存在于根级。它表示机器人在放弃
消息、改为发送一个文件之前，愿意发送多少条聊天消息。允许的取值是 `1` 到
`10`，默认是 `2`。

!!! example "结果一旦需要第二条消息就发文件"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 1 # (1)!
    ```

    1.  最严格的设置。只有能装进一条消息的结果才会留在聊天里。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    在允许条数以内的结果会作为聊天消息留下。一旦再多需要一条消息，整个结果
    就会改为作为一个 `.txt` 文件到达，并且不会发送任何零散的消息。

    </div>
    <div class="result-shot" markdown>

    ![一条较短的结果作为消息，旁边较长的结果作为文件](/images/output-control/allowance.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

!!! warning "这是一个仅根级的字段"

    按钮不能设置 `max_output_messages`。请把它写在根级，让需要不同行为的
    按钮设置 `output: text` 或 `output: file`。

## :material-table-arrow-right: 根级与按钮一起生效

按钮可以带上自己的 `output`。关于输出，按钮能设置的只有这一项。最终生效的
模式就是：

!!! abstract "如何选出生效的模式"

    ```text title="Order of precedence"
    button output  →  root output  →  auto
    ```

| 根级 `output` | 按钮 `output` | 实际发生什么 |
|---------------|-----------------|-----------------------|
| 未设置或 `auto` | 未设置 | `auto` |
| 未设置或 `auto` | `auto` | `auto` |
| 未设置或 `auto` | `text` | `text` |
| 未设置或 `auto` | `file` | `file` |
| `text` | 未设置 | `text` |
| `text` | `auto` | `auto` |
| `text` | `text` | `text` |
| `text` | `file` | `file` |
| `file` | 未设置 | `file` |
| `file` | `auto` | `auto` |
| `file` | `text` | `text` |
| `file` | `file` | `file` |

下面各节逐一讲解这十二种组合，每种配一个示例。

## :material-numeric-1-box-outline: 当根级是 `auto` 时

这也包括根级什么都没写的情况，因为缺少 `output` 就表示 `auto`。

### 根级是 `auto`，按钮没写

最常见的情况。较短的结果留在聊天里；需要超过 `max_output_messages` 条消息
的结果会作为文件到达。

!!! example "普通根级下的普通按钮"

    ```yaml title="config.yaml"
    output: auto
    max_output_messages: 2

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
    ```

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    `df -h` 只打印几行，所以能装进一条消息，作为代码块留在聊天里，上面带着
    摘要。

    </div>
    <div class="result-shot" markdown>

    ![磁盘用量以一条 Telegram 消息返回](/images/output-control/disk-usage.png){ .shot loading=lazy }

    </div>
    </div>

### 根级是 `auto`，按钮写 `auto`

在这里，在按钮上写 `auto` 不会改变任何行为。只有当您希望这个按钮无论根级
以后变成什么都保持 `auto` 时，才值得写上它。

!!! example "在按钮上写明默认值"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Disk usage
        type: button
        function: command
        command: "df -h"
        output: auto # (1)!
    ```

    1.  今天的行为与不写一样，但即使根级以后改成 `text` 或 `file`，这个按钮
        仍然保持原样。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    完全是自动行为：结果短时用文本，超过允许条数后用一个文件。在挂载了很多
    磁盘的服务器上，同一个按钮也可能发给您一个文件。

    </div>
    <div class="result-shot" markdown>

    ![同一个按钮，一次作为消息，一次作为文件](/images/output-control/disk-usage-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### 根级是 `auto`，按钮写 `text`

适合那些您总想直接在聊天里阅读的按钮，哪怕输出稍微长一点 —— 比如一份您愿意
滚动浏览而不是下载的状态列表。

!!! example "让一个按钮留在聊天里"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Service status
        type: button
        function: command
        command: "systemctl status nginx --no-pager"
        output: text # (1)!
    ```

    1.  无论 `max_output_messages` 是多少，都不会变成文件。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    结果被拆成多条消息，每条都回复上一条，所以顺序不会乱。按钮留在最后一
    部分上。

    </div>
    <div class="result-shot" markdown>

    ![一份较长的状态被拆成三条 Telegram 消息](/images/output-control/service-status-split.png){ .shot loading=lazy }

    </div>
    </div>

### 根级是 `auto`，按钮写 `file`

日志和备份的典型选择：您几乎从不想在聊天里读它们，而文件更方便保存。

!!! example "这一个总是下载"

    ```yaml title="config.yaml"
    output: auto

    menu:
      - name: Nginx logs
        type: button
        function: command
        command: "journalctl -u nginx -n 2000 --no-pager"
        output: file # (1)!
    ```

    1.  菜单其余部分仍然自动判断。只有这个按钮被固定为文件。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    一个附件，说明文字里带着按钮名称、退出码和耗时。点击它可以阅读整份
    日志，也可以留着以后看。

    </div>
    <div class="result-shot" markdown>

    ![Nginx 日志以 Telegram 文件送达](/images/output-control/file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-2-box-outline: 当根级是 `output: text` 时

除非按钮另有说明，否则每个按钮都留在聊天里。如果您不喜欢在手机上下载
文件，就选这个。

### 根级是 `text`，按钮没写

!!! example "到处都是聊天消息"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    一条短消息，因为 `uptime` 只打印一行。更长的结果只会被拆成更多消息，
    最多十条。

    </div>
    <div class="result-shot" markdown>

    ![运行时长以一条 Telegram 消息返回](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### 根级是 `text`，按钮写 `auto`

这个按钮重新启用自动切换，所以即使菜单里其他按钮从不发文件，它也可能发。
切换点由根级的 `max_output_messages` 决定。

!!! example "允许一个按钮在输出变长时发文件"

    ```yaml title="config.yaml"
    output: text
    max_output_messages: 3 # (1)!

    menu:
      - name: Package list
        type: button
        function: command
        command: "dpkg -l"
        output: auto
    ```

    1.  只有这个按钮会读取这个允许条数，因为只有它处于 `auto` 模式。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    较短的软件包列表最多以三条消息到达。在装满软件的服务器上，同一个按钮会
    超过允许条数，改为发送一个文件。

    </div>
    <div class="result-shot" markdown>

    ![软件包列表，一次作为消息，一次作为文件](/images/output-control/package-list-pair.png){ .shot .shot--wide loading=lazy }

    </div>
    </div>

### 根级是 `text`，按钮写 `text`

与不写一样，但把意图写了下来。在很长的配置里很有用，因为您不希望以后改动
根级时连带改变这个按钮。

!!! example "把一个按钮固定为聊天消息"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Last logins
        type: button
        function: command
        command: "last -n 20"
        output: text
    ```

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    始终是聊天消息，最多十条，无论以后把根级改成什么。

    </div>
    <div class="result-shot" markdown>

    ![最近的登录记录以一条 Telegram 消息返回](/images/output-control/last-logins.png){ .shot loading=lazy }

    </div>
    </div>

### 根级是 `text`，按钮写 `file`

!!! example "只发聊天消息的菜单里，有一个可下载的按钮"

    ```yaml title="config.yaml"
    output: text

    menu:
      - name: Full system report
        type: button
        function: command
        command: "/usr/local/bin/report.sh"
        output: file
        timeout: "5m" # (1)!
    ```

    1.  报告可能要跑一会儿，所以这个按钮的超时比全局超时更长。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    在一个原本只发消息的菜单里，这是唯一的附件。说明文字会显示报告用了多长
    时间，当它跑上几分钟时很有用。

    </div>
    <div class="result-shot" markdown>

    ![完整系统报告以 Telegram 文件送达](/images/output-control/system-report-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-numeric-3-box-outline: 当根级是 `output: file` 时

每个结果都是文件，哪怕只有一行。这适合主要用来出报告和存档的机器人。

### 根级是 `file`，按钮没写

!!! example "到处都是文件"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
    ```

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    即使只有一行输出也会作为附件到达。说明文字里仍然带着摘要，所以不打开
    文件也能看到退出码。

    </div>
    <div class="result-shot" markdown>

    ![运行时长以一个很小的 Telegram 文件送达](/images/output-control/uptime-file.png){ .shot loading=lazy }

    </div>
    </div>

!!! warning "连很小的结果也会变成下载"

    一行答案也会作为需要打开的附件到达。如果这让您觉得麻烦，请把根级留在
    `auto`，只在需要的按钮上写 `output: file`。

### 根级是 `file`，按钮写 `auto`

这个按钮跳出了只发文件的规则，重新恢复正常行为。

!!! example "让一次快速检查保持可读"

    ```yaml title="config.yaml"
    output: file
    max_output_messages: 2

    menu:
      - name: Uptime
        type: button
        function: command
        command: "uptime"
        output: auto # (1)!
    ```

    1.  只有这个按钮回到自动切换。

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    这条只有一行的答案很短，所以它作为消息留在聊天里，而菜单中其他按钮仍然
    发送文件。

    </div>
    <div class="result-shot" markdown>

    ![运行时长重新以消息回到聊天里](/images/output-control/uptime.png){ .shot loading=lazy }

    </div>
    </div>

### 根级是 `file`，按钮写 `text`

!!! example "强制一个按钮回到聊天里"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Who is logged in
        type: button
        function: command
        command: "who"
        output: text
    ```

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    无论登录用户的列表变得多长，都是消息，永远不是附件。超过十条消息时，
    机器人会注明其余内容已被截断。

    </div>
    <div class="result-shot" markdown>

    ![已登录用户以一条 Telegram 消息返回](/images/output-control/who.png){ .shot loading=lazy }

    </div>
    </div>

### 根级是 `file`，按钮写 `file`

重复根级的取值。这没有坏处，而且以后放宽根级时，这个按钮依然正确。

!!! example "一个必须始终是文件的按钮"

    ```yaml title="config.yaml"
    output: file

    menu:
      - name: Database dump log
        type: button
        function: command
        command: "cat /var/log/pg_dump.log"
        output: file
    ```

!!! success "您在聊天里看到的内容"

    <div class="result" markdown>
    <div class="result-text" markdown>

    一个附件，正如根级已经要求的那样。把它写在按钮上，可以让这个意图在根级
    改动后依然保留。

    </div>
    <div class="result-shot" markdown>

    ![数据库转储日志以 Telegram 文件送达](/images/output-control/db-dump-file.png){ .shot loading=lazy }

    </div>
    </div>

## :material-file-document-outline: 文件是什么样的

当结果以文件送达时，您会收到一个附件和一段简短的说明文字。

- **名称** —— 按钮名称转成小写并用连字符连接，然后是 UTC 的日期和时间，
  例如 `nginx-logs-20260913-091204.txt`。
- **说明文字** —— 与聊天消息顶部相同的摘要：按钮名称、退出码、耗时，以及
  出现超时或截断时的提示。
- **正文** —— 再次是那段摘要，然后是 `--- stdout ---` 部分；当命令向错误
  输出写过内容时，还有 `--- stderr ---` 部分。

!!! success "下载的文件里面"

    ```text title="nginx-logs-20260913-091204.txt"
    Button: Nginx logs
    Exit: 0
    Duration: 316ms

    --- stdout ---
    2026-09-13T09:11:58 nginx: worker process started
    ...

    --- stderr ---
    2026-09-13T09:12:01 nginx: warning: duplicate server name
    ```

## :material-alert-outline: 您无法更改的限制

!!! warning "上限由 Telegram 决定，不是机器人"

    - 一条聊天消息最多 4096 字节。
    - 在 `text` 模式下，机器人最多发送十条消息，并以
      `(output too long; showing first N bytes)` 结束。
    - 无论您把 `max_output_messages` 写成多少，十条消息仍然是硬上限。

!!! note "即使上传失败，您仍会拿到结果"

    当文件发不出去时 —— 没有网络，或者 Telegram 拒绝了它 —— 机器人会回退到
    聊天消息，这样结果就不会丢失。

## :material-console: 运行命令使用根级设置

**`$ >_ Run Command`** 按钮不属于您的菜单，因此它没有自己的 `output`。它
始终跟随根级的 `output` 和 `max_output_messages`。参见
[菜单 → 运行命令](concepts/menu.md#run-command)。

## :material-link-variant: 相关

<div class="grid cards cols-2" markdown>

-   :material-cog-outline:{ .middle } __配置__

    ---

    每个根级键，包括同一张表里的三个输出字段。

    [:octicons-arrow-right-24: 配置](configuration.md#root-fields)

-   :material-gesture-tap-button:{ .middle } __按钮__

    ---

    按钮可以覆盖哪些设置，其中包括 `output`。

    [:octicons-arrow-right-24: 按钮](concepts/button.md)

-   :material-view-list:{ .middle } __菜单__

    ---

    结果如何与菜单并存，以及“运行命令”的作用。

    [:octicons-arrow-right-24: 菜单](concepts/menu.md)

-   :material-timer-outline:{ .middle } __Shell__

    ---

    您的命令运行时使用的 shell、工作目录和超时。

    [:octicons-arrow-right-24: Shell](concepts/shell.md)

</div>
