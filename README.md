# AgileX LIMO Driver

This module implements the base driver for the [AgileX LIMO](https://global.agilex.ai/products/limo) base to be used with the [viam-server](https://docs.viam.com/). This driver supports differential, ackermann, and omni directional steering modes over the serial port.

Documentation: https://docs.viam.com/components/base/agilex-limo/

## Build and Run

The module is avaiable for configuration and use via the Viam Registry at https://app.viam.com/module/viam/agilex-limo.

For local development, this module is written in Go.

To build: `make limobase`<br>
To test: `make test`

## Configuration

```javascript
{
  "drive_mode": "", //Required. valid values are "differential", "ackermann", or "omni"
  "serial_path": "" //Optional. serial path for communication. Defaults to "/dev/ttyTHS1"
}
```

Example:

```javascript
{
  "drive_mode": "differential"
}
```
