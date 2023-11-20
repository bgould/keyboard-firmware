//go:build tinygo

package blehid

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
	case keyboard.RptConsumer:
		sendConsumerReport(uint16(rpt[3])<<8|(uint16(rpt[2])), 0, 0, 0)
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

func sendKeyboardReport(mod, k1, k2, k3, k4, k5, k6 byte) {
	keybuf[0] = 0x02
	keybuf[1] = mod
	keybuf[2] = 0
	keybuf[3] = k1
	keybuf[4] = k2
	keybuf[5] = k3
	keybuf[6] = k4
	keybuf[7] = k5
	keybuf[8] = k6
	port.tx(keybuf)
}

func sendMouseReport(buttons, x, y, wheel byte) {
	mousebuf[0] = 0x01
	mousebuf[1] = buttons
	mousebuf[2] = x
	mousebuf[3] = y
	mousebuf[4] = wheel
	port.tx(mousebuf)
}

func sendConsumerReport(k1, k2, k3, k4 uint16) {
	conbuf[0] = 0x03 // REPORT_ID
	conbuf[1] = uint8(k1)
	conbuf[2] = uint8((k1 & 0x0300) >> 8)
	conbuf[3] = uint8(k2)
	conbuf[4] = uint8((k2 & 0x0300) >> 8)
	conbuf[5] = uint8(k3)
	conbuf[6] = uint8((k3 & 0x0300) >> 8)
	conbuf[7] = uint8(k4)
	conbuf[8] = uint8((k4 & 0x0300) >> 8)
	port.tx(conbuf)
}
