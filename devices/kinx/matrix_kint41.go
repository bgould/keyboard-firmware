//go:build teensy41

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/matrix/kinx/kint41"
)

var (
	console = machine.Serial
	adapter = &kint41.Controller{}
	matrix  = adapter.NewMatrix()
)

func init() {
	machine.UART8.Configure(machine.UARTConfig{
		TX: machine.UART8_TX_PIN,
		RX: machine.UART8_RX_PIN,
	})
	machine.Serial = machine.UART8
}

func configureMatrix() {
	adapter.Configure()
}
