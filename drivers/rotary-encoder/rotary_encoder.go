//go:build tinygo

package rotary_encoder

import (
	"machine"
)

var (
	states = []int8{0, -1, 1, 0, 1, 0, 0, -1, -1, 0, 0, 1, 0, 1, -1, 0}
	// enc    = New(machine.ROT_A, machine.ROT_B)
)

func New(pinA, pinB machine.Pin) *Device {
	return &Device{pinA: pinA, pinB: pinB, oldAB: 0b00000011}
}

type Device struct {
	pinA machine.Pin
	pinB machine.Pin

	precision int

	oldAB int
	value int
}

type Config struct {
	Precision int
}

func (enc *Device) Configure(cfg Config) {
	enc.pinA.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	enc.pinA.SetInterrupt(machine.PinRising|machine.PinFalling, enc.interrupt)

	enc.pinB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	enc.pinB.SetInterrupt(machine.PinRising|machine.PinFalling, enc.interrupt)

	if cfg.Precision > 0 {
		enc.precision = cfg.Precision
	} else {
		enc.precision = 4
	}
}

func (enc *Device) interrupt(pin machine.Pin) {
	aHigh, bHigh := enc.pinA.Get(), enc.pinB.Get()
	enc.oldAB <<= 2
	if aHigh {
		enc.oldAB |= 1 << 1
	}
	if bHigh {
		enc.oldAB |= 1
	}
	enc.value += int(states[enc.oldAB&0x0f])
}

func (enc *Device) Value() int {
	return enc.value / enc.precision
}

func (enc *Device) SetValue(v int) {
	enc.value = v * enc.precision
}
