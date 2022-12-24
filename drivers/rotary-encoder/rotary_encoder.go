//go:build tinygo

package rotary_encoder

import (
	"machine"
)

var (
	states = []int8{0, -1, 1, 0, 1, 0, 0, -1, -1, 0, 0, 1, 0, 1, -1, 0}
	enc    = New(machine.ROT_A, machine.ROT_B)
)

func New(pinA, pinB machine.Pin) *Device {
	return &Device{pinA: pinA, pinB: pinB, oldAB: 0b00000011}
}

type Device struct {
	pinA machine.Pin
	pinB machine.Pin

	oldAB int
	value int
}

func (enc *Device) Configure() {
	enc.pinA.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	enc.pinA.SetInterrupt(machine.PinRising|machine.PinFalling, enc.interrupt)

	enc.pinB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	enc.pinB.SetInterrupt(machine.PinRising|machine.PinFalling, enc.interrupt)
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
	return enc.value / 4
}

func (enc *Device) SetValue(v int) {
	enc.value = v * 4
}
