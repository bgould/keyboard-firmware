//go:build console_uart

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

func configureConsole() keyboard.Console {
	machine.UART1.Configure(machine.UARTConfig{
		TX: machine.GPIO20,
		RX: machine.GPIO21,
	})
	return machine.UART1
}

// func configureHost() keyboard.Host {
// 	machine.UART1.Configure(machine.UARTConfig{
// 		TX: machine.GPIO20,
// 		RX: machine.GPIO21,
// 	})
// 	return serial.New(machine.UART1)
// }
