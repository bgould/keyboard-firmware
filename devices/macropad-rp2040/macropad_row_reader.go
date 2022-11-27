//go:build macropad_rp2040
// +build macropad_rp2040

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

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

func ReadRow(rowIndex uint8) keyboard.Row {
	switch rowIndex {
	case 0:
		v := keyboard.Row(0)
		for i := range keys {
			if !keys[i].Get() {
				v |= (1 << i)
			}
		}
		return v
	default:
		return 0
	}
}
