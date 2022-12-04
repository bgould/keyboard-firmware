//go:build tinygo && experimental_usb
// +build tinygo,experimental_usb

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

func (host *Host) Send(rpt keyboard.Report) {
	switch rpt.Type() {
	case keyboard.RptKeyboard:
		host.hid.SendKeyboardReport(rpt[0], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
	case keyboard.RptMouse:
		if err := host.hid.SendMouseReport(rpt[2], rpt[3], rpt[4], rpt[5], rpt[6]); err != nil {
			println("HID error:", err)
		}
	}
}

func (host *Host) LEDs() uint8 {
	return host.hid.KeyboardLEDs()
}
