# AgileX LIMO Driver

This module implements the base driver for the [AgileX LIMO](https://global.agilex.ai/education/4) base to be used with [`viam-server`](https://docs.viam.com/). This driver supports differential, ackermann, and omni directional steering modes over the serial port.

## Build and Run

To use this module, follow these instructions to [add a module from the Viam Registry](https://docs.viam.com/registry/configure/#add-a-modular-resource-from-the-viam-registry) and select the `viam:base:agilex-limo` model from the [`agilex-limo` module](https://app.viam.com/module/viam/agilex-limo).

## Configure your AgileX LIMO base

> [!NOTE]  
> Before configuring your base, you must [create a robot](https://docs.viam.com/manage/fleet/robots/#add-a-new-robot).

Navigate to the **Config** tab of your robot’s page in [the Viam app](https://app.viam.com/). Click on the **Components** subtab and click **Create component**. Select the `base` type, then select the `agilex-limo` model. Enter a name for your base and click **Create**.

On the new component panel, copy and paste the following attribute template into your base’s **Attributes** box. 
```json
{
  "drive_mode": "<ackermann|differential|omni>",
  "serial_path": "<your-serial-path>"
}
```

> [!NOTE]  
> For more information, see [Configure a Robot](https://docs.viam.com/manage/configuration/).

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

## Next Steps

- To test your base, go to the [**Control** tab](https://docs.viam.com/manage/fleet/robots/#control).
- To write code against your base, use one of the [available SDKs](https://docs.viam.com/program/).
- To view examples using a base component, explore [these tutorials](https://docs.viam.com/tutorials/).

## Local Development

This module is written in Go.

To build: `make limobase`<br>
To test: `make test`


## License
Copyright 2021-2023 Viam Inc. <br>
Apache 2.0
