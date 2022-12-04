//go:build !experimental_usb
// +build !experimental_usb

package usbhid

import (
	"machine"
	machine_kb "machine/usb/hid/keyboard"
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const debug = false

var kb = machine_kb.New()

type Host struct {
	// hid *usb.HID
}

func New( /*hid *usb.HID*/ ) *Host {
	return &Host{
		// hid: hid,
	}
}

func (host *Host) Send(rpt keyboard.Report) {
	if debug {
		writeDebug(rpt[:])
	}
	switch rpt.Type() {
	case keyboard.RptKeyboard:
		kb.SendReport(rpt[0], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
	case keyboard.RptMouse:
		// if err := host.hid.SendMouseReport(rpt[2], rpt[3], rpt[4], rpt[5], rpt[6]); err != nil {
		// println("HID error:", err)
		// }
	}
}

func (host *Host) LEDs() uint8 {
	return 0
	// return host.hid.KeyboardLEDs()
}

func writeDebug(r []byte) {
	for i := 0; i < 8; i++ {
		machine.Serial.Write([]byte(hex(r[i])))
	}
	machine.Serial.WriteByte('\n')
}

func hex(b uint8) string {
	s := strconv.FormatUint(uint64(b), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}
