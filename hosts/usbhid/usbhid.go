//go:build tinygo

package usbhid

import (
	"machine"
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const debug = false

type Host struct{}

func New() *Host {
	return &Host{}
}

func (host *Host) Send(rpt keyboard.Report) {
	if debug {
		writeDebug(rpt[:])
	}
	switch rpt.Type() {
	case keyboard.RptKeyboard:
		sendKeyboardReport(rpt[0], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
	case keyboard.RptMouse:
		sendMouseReport(rpt[2], rpt[3], rpt[4], rpt[5])
	}
}

func (host *Host) LEDs() uint8 {
	return 0
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
