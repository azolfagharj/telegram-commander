# Category

A menu node that opens a submenu instead of running something. A [category](category.md)
has `items` (more [buttons](button.md) or categories) instead of a
[function](function.md). Categories let you group related actions.

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

Buttons and categories together form a tree. See [Buttons and menus](../buttons.md).

## Configuration

For category fields (`type`, `items`, `columns`, and more), see
[Configuration → Buttons](../configuration.md#buttons).

## Related

- [Button](button.md) — runs a function when tapped
- [Buttons and menus](../buttons.md) — nesting, layout, and pagination
