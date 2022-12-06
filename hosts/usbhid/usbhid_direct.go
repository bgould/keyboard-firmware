//go:build tinygo && !usbhid_machine && !usbhid_experimental
// +build tinygo,!usbhid_machine,!usbhid_experimental

package usbhid

import (
	"machine/usb/hid"
)

var port = &kbport{}

// TODO: consider moving out of init
func init() {
	hid.SetHandler(port)
}

func sendKeyboardReport(mod, k1, k2, k3, k4, k5, k6 byte) {
	var b [9]byte
	b[0] = 0x02
	b[1] = mod
	b[2] = 0
	b[3] = k1
	b[4] = k2
	b[5] = k3
	b[6] = k4
	b[7] = k5
	b[8] = k6
	port.tx(b[:])
}

func sendMouseReport(buttons, x, y, wheel byte) {
	var b [5]byte
	b[0] = 0x01
	b[1] = buttons
	b[2] = x
	b[3] = y
	b[4] = wheel
	port.tx(b[:])
}

type kbport struct {
	buf     hid.RingBuffer
	waitTxc bool
	dbg     bool
}

func (kb *kbport) tx(b []byte) {
	if kb.waitTxc {
		kb.buf.Put(b)
	} else {
		kb.waitTxc = true
		hid.SendUSBPacket(b)
	}
}

func (kb *kbport) Handler() bool {
	kb.waitTxc = false
	if b, ok := kb.buf.Get(); ok {
		kb.tx(b)
		// kb.waitTxc = true
		// hid.SendUSBPacket(b)
		return true
	}
	return false
}
