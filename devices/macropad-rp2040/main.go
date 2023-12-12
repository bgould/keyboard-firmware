//go:build macropad_rp2040

package main

import (
	"fmt"
	"machine"
	"time"

	rotary_encoder "github.com/bgould/keyboard-firmware/drivers/rotary-encoder"
	"github.com/bgould/keyboard-firmware/hosts/multihost"
	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const _debug = false

var (

	// TODO: encoder API needs to be improved/revamped
	encoder = rotary_encoder.New(machine.ROT_A, machine.ROT_B)

	host   = multihost.New(usbhid.New(), serial.New(serialer))
	matrix = keyboard.NewMatrix(1, 16, keyboard.RowReaderFunc(ReadRow))
	keymap = Keymap()

	board = keyboard.New(machine.Serial, host, matrix, keymap)
)

func init() {
	configurePins()
	encoder.Configure(rotary_encoder.Config{})
}

func main() {

	serialer.Write([]byte("testing\n"))

	board.SetDebug(_debug)

	board.SetEncoders(
		[]keyboard.Encoder{encoder},
		keyboard.EncodersSubscriberFunc(func(index int, clockwise bool) {
			fmt.Fprintf(serialer, "encoder: %d %t\n", index, clockwise)
		}),
	)

	board.SetKeyAction(keyboard.KeyActionFunc(
		func(key keycodes.Keycode, made bool) {
			switch key {
			case keycodes.FN12:
				if made {
					switch board.ActiveLayer() {
					case 0:
						board.SetActiveLayer(1)
					case 1:
						board.SetActiveLayer(0)
					}
					fmt.Fprintf(serialer, "layer: %d\n", board.ActiveLayer())
				}
			default:
				fmt.Fprintf(serialer, "fn: %d %t\n", key-keycodes.FN0, made)
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
