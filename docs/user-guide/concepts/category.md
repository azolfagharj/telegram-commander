# Category

A menu node that opens a submenu instead of running something. It has `items`
(more [buttons](button.md) or categories) instead of a [function](function.md).

## Example

```yaml
- name: System
  type: category
  items:
    - name: Uptime
      type: button
      function: command
      command: "uptime"
```

Buttons and categories together form a tree. See [Menu](menu.md).

## Configuration

For category fields (`type`, `items`, `columns`, and more), see
[Configuration → Menu](../configuration.md#menu).

## Related

- [Button](button.md) — runs a function when tapped
- [Menu](menu.md) — nesting, layout, and pagination
