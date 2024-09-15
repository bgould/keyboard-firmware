//go:build tinygo

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/hosts/usbhid"
	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

var (
	pins   = []machine.Pin{machine.D10, machine.D11, machine.D12, machine.D13}
	host   = usbhid.New()
	matrix = keyboard.NewMatrix(1, 4, keyboard.RowReaderFunc(ReadRow))
	keymap = []keyboard.Layer{
		Layer(
			/*         */ KC_UP__, /*         */
			/**/ KC_LEFT, KC_DOWN, KC_RGHT, /**/
		),
	}
	board = keyboard.New(host, matrix, keymap)
)

func main() {
	ConfigurePins()
	host.Configure()
	for {
		board.Task()
	}
}

func ConfigurePins() {
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
}

func ReadRow(rowIndex uint8) keyboard.Row {
	switch rowIndex {
	case 0:
		v := keyboard.Row(0)
		for i := range pins {
			if !pins[i].Get() {
				v |= (1 << i)
			}
		}
		return v
	default:
		return 0
	}
}

func Layer(
	/**/ K00, /**/
	K01, K02, K03 /**/ Keycode,
) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		/*       0x0  0x1  0x2  0x3  */
		/*****************************/
		/* 0 */ {K00, K01, K02, K03},
		/*****************************/
	})
}
