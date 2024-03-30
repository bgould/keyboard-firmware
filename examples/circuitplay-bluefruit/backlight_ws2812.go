//go:build circuitplay_bluefruit && !backlight.gpio && !backlight.none

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
		IncludeBreathingInSteps: true,
		Driver: &keyboard.BacklightColorStrip{
			ColorStrip: &keyboard.ColorStrip{
				Writer: ws2812.NewWS2812(machine.WS2812),
				Pixels: make([]color.RGBA, 12),
			},
			Interval: 6 * time.Millisecond,
		},
		Steps: 4,
	}
	kbd.SetBacklight(backlight)
	kbd.BacklightDriver().Configure()
}
