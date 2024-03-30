//go:build circuitplay_bluefruit && backlight.gpio && !backlight.none

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

func init() {
	backlight := keyboard.Backlight{
		IncludeBreathingInSteps: true,
		Driver: &keyboard.BacklightGPIO{
			LED: machine.LED,
			PWM: machine.PWM0,
		},
		Steps: 4,
	}
	kbd.SetBacklight(backlight)
	kbd.BacklightDriver().Configure()
}
