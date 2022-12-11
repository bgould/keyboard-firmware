//go:build host_serial
// +build host_serial

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = false

func configureHost() keyboard.Host {
	machine.UART1.Configure(machine.UARTConfig{
		TX: machine.GPIO20,
		RX: machine.GPIO21,
	})
	return serial.New(machine.UART1)
}
