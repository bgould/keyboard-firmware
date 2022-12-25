//go:build macropad_rp2040

package main

import (
	"machine"
	"time"

	rotary_encoder "github.com/bgould/keyboard-firmware/drivers/rotary-encoder"
	"github.com/bgould/keyboard-firmware/keyboard"
)

var (
	// TODO: encoder API needs to be improved/revamped
	encoder = rotary_encoder.New(machine.ROT_A, machine.ROT_B)

	host   = configureHost()
	matrix = keyboard.NewMatrix(1, 16, keyboard.RowReaderFunc(ReadRow))
	keymap = Keymap()
	board  = keyboard.New(&SerialConsole{machine.Serial}, host, matrix, keymap)
)

func init() {
	configurePins()
	encoder.Configure(rotary_encoder.Config{})
}

func main() {

	board.SetDebug(_debug)

	board.SetEncoders([]keyboard.Encoder{encoder}, keyboard.EncodersSubscriberFunc(
		func(index int, clockwise bool) {
			println("encoder changed: ", index, clockwise)
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

type SerialConsole struct {
	machine.Serialer
}

func (sc *SerialConsole) Read(buf []byte) (n int, err error) {
	for i := range buf {
		buf[i], err = sc.ReadByte()
		if err != nil {
			n = i - 1
			return n, err
		}
	}
	return len(buf), nil
}
