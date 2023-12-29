//go:build macropad_rp2040

package main

import (
	"machine"
	"machine/usb"
	"time"

	rotary_encoder "github.com/bgould/keyboard-firmware/drivers/rotary-encoder"
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	_debug          = false
	encoderInterval = 5 * time.Millisecond
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

	keymap = Keymap()
	host   = usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)
	board  = keyboard.New(machine.Serial, host, matrix, keymap)
)

func init() {
	configurePins()
	usb.Serial = vial.MagicSerialNumber("")
	encoder.Configure(rotary_encoder.Config{})
}

func main() {

	board.SetDebug(_debug)

	board.SetKeyAction(keyboard.KeyActionFunc(
		func(key keycodes.Keycode, made bool) {
			if usbvial.UnlockStatus() != vial.UnlockInProgress {
				switch key {
				case keycodes.KC_FN12:
					if made {
						switch board.ActiveLayer() {
						case 0:
							board.SetActiveLayer(1)
						case 1:
							board.SetActiveLayer(0)
						}
						println("layer:", board.ActiveLayer())
						// fmt.Fprintf(serialer, "layer: %d\n", board.ActiveLayer())
					}
				default:
					// fmt.Fprintf(serialer, "fn: %d %t\n", key-keycodes.FN0, made)
				}
			}
		},
	))

	go func() {
		for {
			board.Task()
			time.Sleep(500 * time.Microsecond)
		}
	}()

	for {
		time.Sleep(time.Hour)
	}

}
