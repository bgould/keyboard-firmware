//go:build teensy41 && serial.usb

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

func configureHost() keyboard.Host {
	return serial.New(machine.Serial)
}
