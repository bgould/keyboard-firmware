//go:build macropad_rp2040
// +build macropad_rp2040

package main

import (
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
)

var (
	host   = configureHost()
	matrix = keyboard.NewMatrix(1, 16, keyboard.RowReaderFunc(ReadRow))
	keymap = Keymap()
	board  = keyboard.New(&SerialConsole{machine.Serial}, host, matrix, keymap)
)

func init() {
	configurePins()
}

func main() {
	board.SetDebug(_debug)
	for {
		board.Task()
		time.Sleep(100 * time.Microsecond)
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
