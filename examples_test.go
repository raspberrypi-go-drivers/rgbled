package rgbled_test

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/rgbled"
	"github.com/stianeikeland/go-rpio/v4"
)

func Example() {
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
	if err := rgbLED.SetBrightness(50); err != nil {
		fmt.Printf("impossible to set LED brightness: %s\n", err)
	}
	time.Sleep(3 * time.Second)
	// Switch off the LED
	if err := rgbLED.SetBrightness(0); err != nil {
		fmt.Printf("impossible to set LED brightness: %s\n", err)
	}
}

func ExampleNewRGBLED() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	rgbLED := rgbled.NewRGBLED(18, 27, 22)
	if err := rgbLED.SetColor("#a4eb34"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExampleRGBLED_SetColor() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	rgbLED := rgbled.NewRGBLED(18, 27, 22)
	if err := rgbLED.SetColor("#a4eb34"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExampleRGBLED_SetBrightness() {
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
	// Reduce the brightness to 50%
	if err := rgbLED.SetBrightness(50); err != nil {
		fmt.Printf("impossible to set LED brightness: %s\n", err)
	}
}
