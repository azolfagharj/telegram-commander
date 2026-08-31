# Button

A tappable item in the Telegram menu. A [button](button.md) has a `name` and,
when tapped, runs a [function](function.md).

## Examples

!!! example "Simple button"
    ```yaml title="Button to run 'uptime' command"
    - name: Uptime
      type: button
      function: command
      command: "uptime"
    ```

You can add an emoji prefix to any button with the `icon` field. It is shown
before the name in Telegram and is purely cosmetic.

!!! example "Button with icon"
    ```yaml title="Restart nginx with a restart icon"
    - name: Restart nginx
      type: button
      icon: "🔄"
      function: command
      command: "systemctl restart nginx"
    ```

[Buttons](../buttons.md) and [categories](category.md) together form the menu
tree under the top-level `buttons` key in your [config file](config-file.md).

## Configuration

For every field of a button, see [Configuration → Buttons](../configuration.md#buttons).

## Related

- [Category](category.md) — opens a submenu instead of running something
- [Function](function.md) — what runs when a button is tapped
- [Parameter](parameter.md) — values a function needs from the button
- [Confirmation](confirmation.md) — optional "Are you sure?" step
- [Buttons and menus](../buttons.md) — build and organize your menu
