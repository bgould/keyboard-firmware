//go:build !serial.usb
// +build !serial.usb

package main

import (
	"machine/usb"

	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = false

func configureHost() keyboard.Host {
	hid := &usb.HID{}
	hid.Configure(usb.HIDConfig{})
	return usbhid.New(hid)
}
