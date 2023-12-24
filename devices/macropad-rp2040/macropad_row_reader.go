//go:build macropad_rp2040

package main

import (
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
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

type MacroPadRP2040KeyMapper struct {
	keyboard.Keymap
}

var (
	_ vial.KeyMapper     = (*MacroPadRP2040KeyMapper)(nil)
	_ vial.KeySetter     = (*MacroPadRP2040KeyMapper)(nil)
	_ vial.EncoderMapper = (*MacroPadRP2040KeyMapper)(nil)
	_ vial.EncoderSaver  = (*MacroPadRP2040KeyMapper)(nil)
)

// func (km *MacroPadRP2040KeyMapper) SaveKey(layer, row, col int, kc keycodes.Keycode) {
// 	println(
// 		"layer:", layer,
// 		"row:", row,
// 		"col:", col,
// 		"kc:", kc,
// 		// "trigger:", entry.Trigger,
// 		// "replacement:", entry.Replacement,
// 		// "layers:", entry.Layers,
// 		// "mods:", entry.TriggerMods,
// 		// "negative mask:", entry.NegativeModMask,
// 		// "supressed mods:", entry.SupressedMods,
// 		// "options:", entry.Options,
// 	)
// 	km.Keymap.SetKey(layer, row, col, kc)
// }

func (km *MacroPadRP2040KeyMapper) MapEncoder(layer, idx int) (ccw, cw keycodes.Keycode) {
	return km.MapKey(layer, 0, encIndexCCW), km.MapKey(layer, 0, encIndexCW)
}

func (km *MacroPadRP2040KeyMapper) SaveEncoder(layer, idx int, clockwise bool, kc keycodes.Keycode) {
	if layer >= len(km.Keymap) {
		return
	}
	if idx > 0 {
		return
	}
	encIdx := encIndexCCW
	if clockwise {
		encIdx = encIndexCW
	}
	km.SetKey(layer, 0, encIdx, kc)
	// km.Keymap[layer][0][encIdx] = kc
	// return km.MapKey(layer, encIndexCCW), km.MapKey(layer, encIndexCW)
}
