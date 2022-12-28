//go:build teensy41

package kint41

import (
	"machine"
	"runtime"

	"github.com/bgould/keyboard-firmware/keyboard"
)

func (c *Controller) NewMatrix() *keyboard.Matrix {
	return keyboard.NewMatrix(15, 7, keyboard.RowReaderFunc(c.ReadRow))
}

const (
	LEDCapsLock   = machine.D12
	LEDNumLock    = machine.D26
	LEDScrollLock = machine.D25
	LEDKeyPad     = machine.D24
)

var leds = []machine.Pin{
	LEDCapsLock,
	LEDNumLock,
	LEDScrollLock,
	LEDKeyPad,
}

var rows = []machine.Pin{
	machine.D8,
	machine.D9,
	machine.D10,
	machine.D11,
	machine.D7,
	machine.D16,
	machine.D5,
	machine.D3,
	machine.D4,
	machine.D1,
	machine.D0,
	machine.D2,
	machine.D17,
	machine.D23,
	machine.D21,
}

var columns = []machine.Pin{
	machine.D18,
	machine.D14,
	machine.D15,
	machine.D20,
	machine.D22,
	machine.D19,
	machine.D6,
}

type Controller struct {
}

func (c *Controller) Configure() {
	for _, pin := range leds {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		pin.Set(false)
	}
	for _, pin := range columns {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullUp})
	}
	for _, pin := range rows {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}

func (c *Controller) ReadRow(rowIndex uint8) (row keyboard.Row) {
	for i, pin := range rows {
		v := i != int(rowIndex)
		pin.Set(v)
	}
	delayMicros(10)
	for i, pin := range columns {
		v := pin.Get()
		if !v {
			row |= (1 << i)
		}
	}
	return row
}

func delayMicros(usecs uint32) {
	// (cycles to delay) = microseconds * (cycles per microsecond)
	var cycles = usecs * (runtime.CORE_FREQ / 1e6)
	// cycle counter increments once per cycle; busy-loop until time cycles to delay elapsed
	for start := runtime.DWT_CYCCNT.Get(); runtime.DWT_CYCCNT.Get()-start < cycles; {
	}
}
