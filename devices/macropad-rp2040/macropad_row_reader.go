//go:build macropad_rp2040

package main

import (
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

var keys = []machine.Pin{
	machine.SWITCH,
	machine.KEY1,
	machine.KEY2,
	machine.KEY3,
	machine.KEY4,
	machine.KEY5,
	machine.KEY6,
	machine.KEY7,
	machine.KEY8,
	machine.KEY9,
	machine.KEY10,
	machine.KEY11,
	machine.KEY12,
}

func configurePins() {
	for _, pin := range keys {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
}

var (
	// TODO: make this more generic and move into core library
	encTurned    time.Time
	encClockwise bool
)

func encoderCallback(index int, clockwise bool) {
	encTurned = time.Now()
	encClockwise = clockwise
	// fmt.Fprintf(serialer, "encoder: %d %t\n", index, clockwise)
}

const (
	encIndexCW  = 14
	encIndexCCW = 13
)

func ReadRow(rowIndex uint8) keyboard.Row {
	switch rowIndex {
	case 0:
		v := keyboard.Row(0)
		for i := range keys {
			if !keys[i].Get() {
				v |= (1 << i)
			}
		}
		if time.Since(encTurned) < encoderInterval {
			if encClockwise {
				v |= (1 << encIndexCW)
			} else {
				v |= (1 << encIndexCCW)
			}
		}
		return v
	default:
		return 0
	}
}
