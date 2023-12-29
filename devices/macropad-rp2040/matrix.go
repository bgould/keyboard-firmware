//go:build macropad_rp2040

package main

import (
	"machine"

	rotary_encoder "github.com/bgould/keyboard-firmware/drivers/rotary-encoder"
	"github.com/bgould/keyboard-firmware/keyboard"
)

var (
	encoder = rotary_encoder.New(machine.ROT_A, machine.ROT_B)
	encPos  = keyboard.EncoderPos{
		Encoder: encoder,
		PosCW:   keyboard.Pos{Row: 0, Col: encIndexCW},
		PosCCW:  keyboard.Pos{Row: 0, Col: encIndexCCW},
	}
	reader = keyboard.RowReaderFunc(ReadRow)
	matrix = keyboard.NewMatrix(1, 16, reader).WithEncoders(encPos)
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

const (
	encIndexCW  = 14
	encIndexCCW = 13
)

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
