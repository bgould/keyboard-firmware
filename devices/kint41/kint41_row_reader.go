package main

import (
	"device/arm"
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/timer"
)

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

func configurePins() {
	for _, pin := range columns {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullUp})
	}
	for _, pin := range rows {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}

func ReadRow(rowIndex uint8) (row keyboard.Row) {
	for i, pin := range rows {
		v := i != int(rowIndex)
		pin.Set(v)
	}
	delayForSelect()
	for i, pin := range columns {
		v := pin.Get()
		if !v {
			row |= (1 << i)
		}
	}
	return row
}

func delayMicros(usecs uint32) {
	timer.Wait(time.Duration(usecs) * time.Microsecond)
}

func delayForSelect() {
	for i := 0; i < 1000; i++ {
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
		arm.Asm("nop")
	}
}
