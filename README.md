# [`agilex-limo` module](https://app.viam.com/module/viam/agilex-limo)

This module implements the [`rdk:component:base` API](https://docs.viam.com/components/base/#api) in an `agilex` model for the [AgileX LIMO](https://global.agilex.ai/education/4) base to be used with [`viam-server`](https://docs.viam.com/). This driver supports differential, ackermann, and omni directional steering modes over the serial port.

## Configure your `agilex-limo` base

> [!NOTE]
> Before configuring your base, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the **CONFIGURE** tab of your machine’s page in [the Viam app](https://app.viam.com/).
[Add `base` / `agilex-limo` to your machine](https://docs.viam.com/build/configure/#components).

On the new component panel, copy and paste the following attribute template into your base’s attributes field:

```json
{
  "drive_mode": "<ackermann|differential|omni>",
  "serial_path": "<your-serial-path>"
}
```

> [!NOTE]
> For more information, see [Configure a Machine](https://docs.viam.com/build/configure/).

### Attributes

The following attributes are available for `viam:base:agilex-limo` bases:

| Name | Type | Inclusion | Description |
| ---- | ---- | --------- | ----------- |
| `drive_mode` | string | **Required** | LIMO [steering mode](https://docs.trossenrobotics.com/agilex_limo_docs/operation/steering_modes.html#switching-steering-modes). Options: `differential`, `ackermann`, `omni` (mecanum). |
| `serial_path` | string | Optional | The full filesystem path to the serial device, starting with <file>/dev/</file>. With your serial device connected, you can run `sudo dmesg \| grep tty` to show relevant device connection log messages, and then match the returned device name, such as `ttyTHS1`, to its device file, such as <file>/dev/ttyTHS1</file>. If you omit this attribute, Viam will attempt to automatically detect the path.<br>Default: `/dev/ttyTHS1` |

### Example configurations:

```json
{
  "drive_mode": "differential"
}
```

```json
{
  "drive_mode": "omni",
  "serial_path": "/dev/ttyTHS1"
}
```

## Next steps

- To test your base, go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your base, use one of the [available SDKs](https://docs.viam.com/build/program/).
- To view examples using a base component, explore [these tutorials](https://docs.viam.com/tutorials/).

## Local development

This module is written in Go.

To build: `make limobase`<br>
To test: `make test`

## License
Copyright 2021-2023 Viam Inc. <br>
Apache 2.0
