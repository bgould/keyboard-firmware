//go:build macropad_rp2040

package main

import (
	"machine"
	"time"

	rotary_encoder "github.com/bgould/keyboard-firmware/drivers/rotary-encoder"
)

var (
	enc = rotary_encoder.New(machine.ROT_A, machine.ROT_B)
)

func main() {

	enc.Configure()

	for oldValue := 0; ; {
		time.Sleep(100 * time.Microsecond) // doesn't work without this?
		if newValue := enc.Value(); newValue != oldValue {
			println("value: ", newValue)
			oldValue = newValue
		}
	}

}
