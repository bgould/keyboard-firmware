package blehid

import (
	"github.com/bgould/keyboard-firmware/keyboard"
)

const debug = false

type Host struct {
	cfg      HostConfig
	port     bleKeyboard
	keybuf   [9]byte
	conbuf   [9]byte
	mousebuf [5]byte
	ledState keyboard.LEDs
}

type HostConfig struct {
	Name         string
	Manufacturer string
	ModelNumber  string
}

func New(config HostConfig) *Host {
	host := &Host{cfg: config}
	host.port.connect()
	return host
}

func (host *Host) Send(rpt keyboard.Report) {
	switch rpt.Type() {
	case keyboard.RptKeyboard:
		host.sendKeyboardReport(rpt[0], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
	case keyboard.RptMouse:
		host.sendMouseReport(rpt[2], rpt[3], rpt[4], rpt[5])
	case keyboard.RptConsumer:
		host.sendConsumerReport(uint16(rpt[3])<<8|(uint16(rpt[2])), 0, 0, 0)
	}
}

func (host *Host) LEDs() uint8 {
	return uint8(host.ledState)
}

func (host *Host) sendKeyboardReport(mod, k1, k2, k3, k4, k5, k6 byte) {
	host.keybuf[0] = 0x02
	host.keybuf[1] = mod
	host.keybuf[2] = 0
	host.keybuf[3] = k1
	host.keybuf[4] = k2
	host.keybuf[5] = k3
	host.keybuf[6] = k4
	host.keybuf[7] = k5
	host.keybuf[8] = k6
	host.port.tx(host.keybuf[:])
}

func (host *Host) sendMouseReport(buttons, x, y, wheel byte) {
	host.mousebuf[0] = 0x01
	host.mousebuf[1] = buttons
	host.mousebuf[2] = x
	host.mousebuf[3] = y
	host.mousebuf[4] = wheel
	host.port.tx(host.mousebuf[:])
}

func (host *Host) sendConsumerReport(k1, k2, k3, k4 uint16) {
	host.conbuf[0] = 0x03 // REPORT_ID
	host.conbuf[1] = uint8(k1)
	host.conbuf[2] = uint8((k1 & 0x0300) >> 8)
	host.conbuf[3] = uint8(k2)
	host.conbuf[4] = uint8((k2 & 0x0300) >> 8)
	host.conbuf[5] = uint8(k3)
	host.conbuf[6] = uint8((k3 & 0x0300) >> 8)
	host.conbuf[7] = uint8(k4)
	host.conbuf[8] = uint8((k4 & 0x0300) >> 8)
	host.port.tx(host.conbuf[:])
}
