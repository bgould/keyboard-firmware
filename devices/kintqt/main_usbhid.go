//go:build (nrf52840 || rp2040) && !host_tinyterm && !host_serial

package main

import (
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = false

func configureHost() keyboard.Host {
	return usbhid.New()
}