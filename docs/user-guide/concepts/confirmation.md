# Confirmation

An optional "Are you sure?" step before a [button](button.md) runs, enabled
with `confirm: true`. Useful for destructive actions.

The bot asks (`Confirm: … ?`) with Yes / Cancel, plus Home (and Back when you
are inside a category). The prompt expires after a while (default 5 minutes).
Change it with `confirm_ttl` in your [config file](config-file.md).

## Example

```yaml
- name: Stop nginx
  type: button
  function: command
  command: "systemctl stop nginx"
  confirm: true
```

## Configuration

For `confirm` on buttons and global `confirm_ttl`, see
[Configuration](../configuration.md).

## Related

- [Button](button.md) — the node that can require confirmation
- [Menu → Confirmation](menu.md#confirmation) — examples and behavior
