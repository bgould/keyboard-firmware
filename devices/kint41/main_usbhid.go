//go:build teensy41 && experimental_usb && !serial.usb

package main

import (
	"machine"
	"machine/usb"

	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = false

func init() {
	machine.UART8.Configure(machine.UARTConfig{
		TX: machine.UART8_TX_PIN,
		RX: machine.UART8_RX_PIN,
	})
	machine.Serial = machine.UART8
}

func configureHost() keyboard.Host {
	hid := &usb.HID{}
	hid.Configure(usb.HIDConfig{})
	return usbhid.New(hid)
}
