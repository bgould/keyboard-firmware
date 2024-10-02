//go:build macropad_rp2040

package macropad_rp2040

import (
	"machine"

	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers/encoders"
)

const (
	USBManufacturer = "Adafruit"
	USBProduct      = "MacroPad RP2040"
)

type Board struct {
	encoder *encoders.QuadratureDevice
}

func New() *Board {
	return &Board{
		encoder: encoders.NewQuadratureViaInterrupt(machine.ROT_A, machine.ROT_B),
	}
}

// Configure sets up the encoder, WS2812, and key GPIO pins. It is should be
// called once, after init has completed.
func (dev *Board) Configure() {
	for _, pin := range keys {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
	machine.NEOPIXEL.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dev.encoder.Configure(encoders.QuadratureConfig{})
}

func (dev *Board) ReadRow(rowIndex uint8) keyboard.Row {
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

func (dev *Board) NewMatrix() *keyboard.Matrix {
	return keyboard.NewMatrix(MatrixRows, MatrixCols, dev).WithEncoders(
		keyboard.EncoderPos{
			Encoder: dev.encoder,
			PosCW:   keyboard.Pos{Row: 0, Col: encIndexCW},
			PosCCW:  keyboard.Pos{Row: 0, Col: encIndexCCW},
		},
	)
}

func (dev *Board) NewVialKeyboard(layers ...keyboard.Layer) (*keyboard.Keyboard, *usbvial.Host) {
	keymap := Keymap(layers...)
	matrix := dev.NewMatrix()
	macros, ok := keyboard.NewDefaultMacroDriver(32, 4096).(usbvial.VialMacroDriver)
	if !ok {
		macros = nil
	}
	host := NewVialHost(keymap, matrix, macros)
	kbd := keyboard.New(host, matrix, keymap)
	kbd.SetEventReceiver(host)
	kbd.SetMacroDriver(macros)
	return kbd, host
}

const (
	MatrixRows = 1
	MatrixCols = 16
)

const (
	encIndexCW  = 14
	encIndexCCW = 13
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
