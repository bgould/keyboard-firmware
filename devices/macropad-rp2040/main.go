//go:build macropad_rp2040

package main

import (
	"image/color"
	"machine"
	"machine/usb"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers/encoders"
	"tinygo.org/x/drivers/ws2812"
	"tinygo.org/x/tinyfs/littlefs"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

var (
	backlight = keyboard.Backlight{
		Driver: &keyboard.BacklightColorStrip{
			ColorStrip: &keyboard.ColorStrip{
				Writer: ws2812.NewWS2812(machine.NEOPIXEL),
				Pixels: make([]color.RGBA, 12),
			},
			Interval: 6 * time.Millisecond,
		},
		Steps: 4,
	}
	keymap = Keymap()
	host   = usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)
	board  = keyboard.New(host, matrix, keymap)
)

func init() {
	initFilesystem()
	configurePins()

	usb.Manufacturer = "Adafruit"
	usb.Product = "MacroPad RP2040"
	usb.Serial = vial.MagicSerialNumber("")

	encoder.Configure(encoders.QuadratureConfig{})
	host.Configure()

	board.SetBacklight(backlight)
}

func initFilesystem() {
	lfs := littlefs.New(machine.Flash)
	lfs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	board.SetFS(lfs)
}

func main() {
	if backlight.Driver != nil {
		backlight.Driver.SetBacklight(keyboard.BacklightOff, 0xFF)
	}
	board.ConfigureFilesystem()
	board.EnableConsole(machine.Serial)
	for {
		board.Task()
		time.Sleep(500 * time.Microsecond)
	}
}
