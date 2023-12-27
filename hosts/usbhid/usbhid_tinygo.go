//go:build tinygo && !teensy41

package usbhid

import (
	"machine/usb/hid"
)

var port = &kbport{}

// TODO: consider moving out of init
func init() {
	hid.SetHandler(port)
}

var (
	keybuf   = make([]byte, 16)
	conbuf   = make([]byte, 9)
	mousebuf = make([]byte, 5)

	ledState uint8
)

func sendKeyboardReport(mod, k1, k2, k3, k4, k5, k6 byte) {
	keybuf[0] = 0x02 // REPORT_ID
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

func sendNKROKeyboardReport(mod, b1, b2, b3, b4, b5, b6 byte) {
	keybuf[0] = 0x02 // REPORT_ID
	keybuf[1] = mod
	keybuf[2] = 0
	keybuf[3] = b1
	keybuf[4] = b2
	keybuf[5] = b3
	keybuf[6] = b4
	keybuf[7] = b5
	keybuf[8] = b6
	port.tx(keybuf)
}

func sendMouseReport(buttons, x, y, wheel byte) {
	mousebuf[0] = 0x01 // REPORT_ID
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

type kbport struct {
	buf hid.RingBuffer
	txc bool
	dbg bool
}

//go:inline
func (port *kbport) tx(b []byte) {
	if port.txc {
		port.buf.Put(b)
	} else {
		port.txc = true
		hid.SendUSBPacket(b)
	}
}

func (port *kbport) TxHandler() bool {
	port.txc = false
	if b, ok := port.buf.Get(); ok {
		port.tx(b)
		return true
	}
	return false
}

func (port *kbport) RxHandler(rx []byte) bool {
	if len(rx) < 2 || rx[0] != 0x02 {
		return true
	}
	ledState = rx[1]
	return true
}
