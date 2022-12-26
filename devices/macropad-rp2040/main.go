//go:build macropad_rp2040

package main

import (
	"machine"
	"time"

	rotary_encoder "github.com/bgould/keyboard-firmware/drivers/rotary-encoder"
	"github.com/bgould/keyboard-firmware/hosts/multihost"
	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = false

var (

	// TODO: encoder API needs to be improved/revamped
	encoder = rotary_encoder.New(machine.ROT_A, machine.ROT_B)

	console = configureConsole()
	host    = multihost.New(usbhid.New(), serial.New(console))
	matrix  = keyboard.NewMatrix(1, 16, keyboard.RowReaderFunc(ReadRow))
	keymap  = Keymap()

	board = keyboard.New(console, host, matrix, keymap)
)

func init() {
	configurePins()
	encoder.Configure(rotary_encoder.Config{})
}

func main() {

	console.Write([]byte("testing\n"))

	board.SetDebug(_debug)

	board.SetEncoders(
		[]keyboard.Encoder{encoder},
		keyboard.EncodersSubscriberFunc(func(index int, clockwise bool) {
			console.Write([]byte("encoder: "))
			console.Write([]byte{"0123456789"[index]})
			if clockwise {
				console.Write([]byte(" true\n"))
			} else {
				console.Write([]byte(" false\n"))
			}
		}),
	)

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
