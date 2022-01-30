//go:build serial.usb
// +build serial.usb

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = true

func configureHost() keyboard.Host {
	return serial.New(machine.Serial)
}
