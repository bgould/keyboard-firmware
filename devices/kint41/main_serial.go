//go:build teensy41 && serial.usb

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/hosts/serial"
)

const _debug = true

func configureHost() keyboard.Host {
	return serial.New(machine.Serial)
}
