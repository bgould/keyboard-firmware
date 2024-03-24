//go:build macropad_rp2040 && !backlight.none

package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers/ws2812"
)

func init() {
	backlight := keyboard.Backlight{
		Driver: &keyboard.BacklightColorStrip{
			ColorStrip: &keyboard.ColorStrip{
				Writer: ws2812.NewWS2812(machine.NEOPIXEL),
				Pixels: make([]color.RGBA, 12),
			},
			Interval: 6 * time.Millisecond,
		},
		Steps: 8,
	}
	kbd.SetBacklight(backlight)
	kbd.BacklightDriver().Configure()
}
