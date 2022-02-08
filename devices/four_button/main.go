package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = true

var (
	pins   = []machine.Pin{machine.D23, machine.D7, machine.D22, machine.D21}
	layers = FourButtonDefaultKeymap()
	matrix = keyboard.NewMatrix(1, 4, keyboard.RowReaderFunc(ReadRow))
)

func main() {

	configurePins()
	host := configureHost()

	board := keyboard.New(machine.Serial, host, matrix, layers)
	board.SetDebug(_debug)
	for {
		board.Task()
	}

}

func configurePins() {
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullUp})
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
