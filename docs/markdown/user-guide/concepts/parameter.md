---
icon: material/tune
title: Parameter
description: A parameter is a named value a function needs. Write each value directly on the button using the same name declared by the function.
---

# :material-tune: Parameter

A named value a [function](function.md) needs. For example, the built-in
`command` function needs one parameter called `command`. You provide parameter
values on the [button](button.md).

Some parameters are required; others are optional with a default defined in the
function file.

## Pass a value from a button

Write the parameter name as a key directly on the button. The built-in
`command`, `path`, and `args` fields follow this rule, and custom parameter
names work the same way.

!!! example "Supply custom parameter names"

    ```yaml title="Recent service logs"
    - name: Nginx logs
      type: button
      function: journal-unit
      unit: "nginx.service"
      lines: 100
    ```

`unit` and `lines` must be declared by `journal-unit`. The numeric value for
`lines` does not need quotes.

!!! warning "Do not use a nested `params:` map"

    A custom function file uses `params:` to declare its parameters. A button
    does not. Put `host:`, `unit:`, or any other value directly on the button.

## Validation

[`validate`](../cli.md#validate) checks that:

- every required parameter has a value;
- every value on the button is declared by its function;
- values declared as `int` contain an integer;
- values declared as `bool` contain a valid boolean;
- defaults match the declared parameter type.

Optional parameters use their default when the button leaves them out.

Parameter names cannot be the same as button settings: `name`, `type`, `icon`,
`id`, `function`, `confirm`, `timeout`, `workdir`, `env`, `columns`, or
`items`. The names `command`, `path`, and `args` are allowed.

See [Functions → Passing values from a button](../functions/index.md#passing-values-from-a-button)
for the full explanation and examples.

## Configuration

For button settings and parameter keys, see
[Configuration → Menu](../configuration.md#menu).

## Related

- [Function](function.md) — what uses parameters
- [Rules](../functions/write-your-own/rules.md) — parameter rules for custom functions
