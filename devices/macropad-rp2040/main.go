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
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
	"tinygo.org/x/drivers/encoders"
	"tinygo.org/x/drivers/ws2812"
	"tinygo.org/x/tinyfs/littlefs"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

var (
	backlight = keyboard.Backlight{
		IncludeBreathingInSteps: true,
		Driver: &keyboard.BacklightColorStrip{
			ColorStrip: keyboard.ColorStrip{
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

	backlight.Driver.SetBacklight(keyboard.BacklightOff, 0xFF)

	board.ConfigureFilesystem()
	board.EnableConsole(machine.Serial)

	board.SetKeyAction(keyboard.KeyActionFunc(
		func(key keycodes.Keycode, made bool) {
			if host.UnlockStatus() != vial.UnlockInProgress {
				switch key {
				case keycodes.KC_FN12:
					if made {
						switch board.ActiveLayer() {
						case 0:
							board.SetActiveLayer(1)
						case 1:
							board.SetActiveLayer(0)
						}
						// println("layer:", board.ActiveLayer())
						// fmt.Fprintf(serialer, "layer: %d\n", board.ActiveLayer())
					}
				default:
					// fmt.Fprintf(serialer, "fn: %d %t\n", key-keycodes.FN0, made)
				}
			}
		},
	))

	go func() {
		for {
			board.Task()
			time.Sleep(500 * time.Microsecond)
		}
	}()

	for {
		time.Sleep(time.Hour)
	}

}
