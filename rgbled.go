package rgbled

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/bbayszczak/raspberrypi-go-drivers/led"
)

// RGBLED is the representation of an RGB LED
type RGBLED struct {
	redLED     *led.LED
	greenLED   *led.LED
	blueLED    *led.LED
	redComp    int
	greenComp  int
	blueComp   int
	brightness int
}

// NewRGBLED creates a new RGBLED
func NewRGBLED(redPin uint8, greenPin uint8, bluePin uint8) *RGBLED {
	rgbLED := RGBLED{
		redLED:     led.NewLED(redPin),
		greenLED:   led.NewLED(greenPin),
		blueLED:    led.NewLED(bluePin),
		brightness: 100,
		redComp:    0,
		greenComp:  0,
		blueComp:   0,
	}
	rgbLED.redLED.Dimmable()
	rgbLED.greenLED.Dimmable()
	rgbLED.blueLED.Dimmable()
	return &rgbLED
}

// SetColor set the LED color using an hex code (#xxxxxx)
func (rgbLED *RGBLED) SetColor(hexCode string) error {
	var err error
	if len(hexCode) != 7 {
		return errors.New("hexCode should be in format #xxxxxx")
	}
	rgbLED.redComp, err = hexToPercent(hexCode[1:3])
	if err != nil {
		return err
	}
	rgbLED.greenComp, err = hexToPercent(hexCode[3:5])
	if err != nil {
		return err
	}
	rgbLED.blueComp, err = hexToPercent(hexCode[5:7])
	if err != nil {
		return err
	}
	if err := rgbLED.redLED.SetBrightness(getPercentage(rgbLED.redComp, rgbLED.brightness)); err != nil {
		return fmt.Errorf("impossible to change red component brightness: %s", err)
	}
	if err := rgbLED.greenLED.SetBrightness(getPercentage(rgbLED.greenComp, rgbLED.brightness)); err != nil {
		return fmt.Errorf("impossible to change green component brightness: %s", err)
	}
	if err := rgbLED.blueLED.SetBrightness(getPercentage(rgbLED.blueComp, rgbLED.brightness)); err != nil {
		return fmt.Errorf("impossible to change blue component brightness: %s", err)
	}
	return nil
}

// SetBrightness set the LED brightness
func (rgbLED *RGBLED) SetBrightness(percentage int) error {
	rgbLED.brightness = percentage
	if err := rgbLED.redLED.SetBrightness(getPercentage(rgbLED.redComp, rgbLED.brightness)); err != nil {
		return fmt.Errorf("impossible to change red component brightness: %s", err)
	}
	if err := rgbLED.greenLED.SetBrightness(getPercentage(rgbLED.greenComp, rgbLED.brightness)); err != nil {
		return fmt.Errorf("impossible to change green component brightness: %s", err)
	}
	if err := rgbLED.blueLED.SetBrightness(getPercentage(rgbLED.blueComp, rgbLED.brightness)); err != nil {
		return fmt.Errorf("impossible to change blue component brightness: %s", err)
	}
	return nil
}

func hexToPercent(hexCode string) (int, error) {
	hex, err := strconv.ParseInt(hexCode, 16, 64)
	if err != nil {
		return 0, err
	}
	percent := int(math.RoundToEven(float64(hex) / 255.0 * 100.0))
	return percent, nil
}

func getPercentage(number int, percentage int) int {
	result := float64(number) / 100.0 * float64(percentage)
	return int(math.RoundToEven(result))
}
