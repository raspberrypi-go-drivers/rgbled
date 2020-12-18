# RGBLED

[![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/rgbled.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/rgbled)
![golangci-lint](https://github.com/raspberrypi-go-drivers/led/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspberrypi-go-drivers/rgbled)](https://goreportcard.com/report/github.com/raspberrypi-go-drivers/rgbled)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This drivers allows interact with a RGB LED connected to a GPIO pin

## Documentation

For full documentation, please visit [![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/rgbled.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/rgbled)

## Quick start

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/rgbled"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	rgbLED := rgbled.NewRGBLED(18, 27, 22)
	// Switch on the led an set it to lime green
	if err := rgbLED.SetColor("#a4eb34"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	time.Sleep(3 * time.Second)
	// Reduce the brightness to 50%
	rgbLED.SetBrightness(50)
	time.Sleep(3 * time.Second)
	// Switch off the LED
	rgbLED.SetBrightness(0)
}
```
