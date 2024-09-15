//go:build circuitplay_bluefruit

package circuitplay_bluefruit

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/hosts/usbvial"
)

const (
	USBManufacturer = "Adafruit"
	USBProduct      = "Circuit Playground Bluefruit"
)

type Board struct {
}

func New() *Board {
	return &Board{}
}

// Configure sets up the encoder, WS2812, and key GPIO pins. It is should be
// called once, after init has completed.
func (dev *Board) Configure() {
	slider.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	for _, pin := range keys {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}
	machine.WS2812.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (dev *Board) ReadRow(rowIndex uint8) keyboard.Row {
	switch rowIndex {
	case 0:
		v := keyboard.Row(0)
		for i := range keys {
			if keys[i].Get() {
				v |= (1 << i)
			}
		}
		if !slider.Get() {
			v |= (1 << len(keys))
		}
		return v
	default:
		return 0
	}

}

func (dev *Board) NewMatrix() *keyboard.Matrix {
	return keyboard.NewMatrix(MatrixRows, MatrixCols, dev)
}

func (dev *Board) NewVialKeyboard(layers ...keyboard.Layer) (*keyboard.Keyboard, *usbvial.Host) {
	keymap := Keymap(layers...)
	matrix := dev.NewMatrix()
	host := NewVialHost(keymap, matrix)
	kbd := keyboard.New(host, matrix, keymap)
	return kbd, host
}

const (
	MatrixRows = 1
	MatrixCols = 3
)

var (
	keys   = []machine.Pin{machine.BUTTONA, machine.BUTTONB}
	slider = machine.SLIDER
)
