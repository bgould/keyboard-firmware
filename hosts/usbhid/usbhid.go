//go:build tinygo
// +build tinygo

package usbhid

import (
	"machine/usb"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const debug = false

type Host struct {
	hid *usb.HID
}

func New(hid *usb.HID) *Host {
	return &Host{
		hid: hid,
	}
}

func (host *Host) Send(rpt *keyboard.Report) {
	if debug {
		println(rpt[0], rpt[1], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
	}
	host.hid.SendKeys(rpt[0], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
}

func (host *Host) LEDs() uint8 {
	return host.hid.KeyboardLEDs()
}
