# AIO

This package provides drivers for [Analog Input/Output (AIO)](https://en.wikipedia.org/wiki/Analog-to-digital_converter)
devices. It is normally used by connecting an adaptor such as [BeagleBone](https://gobot.io/documentation/platforms/beaglebone/)
that supports the needed interfaces for analog devices.

## Getting Started

## Installing

```sh
go get -d -u gobot.io/x/gobot/v2/...
```

## Hardware Support

Gobot has a extensible system for connecting to hardware devices. The following AIO devices are currently supported:

- Analog Sensor
- Analog Actuator
- Grove Light Sensor
- Grove Rotary Dial
- Grove Sound Sensor
- Grove Temperature Sensor
- Temperature Sensor (supports linear and NTC thermistor in normal and inverse mode)

More drivers are coming soon...
