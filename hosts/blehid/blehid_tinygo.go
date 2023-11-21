//go:build tinygo && nrf

package blehid

import "machine/usb/hid"

var port = &kbport{}

var (
	keybuf   = make([]byte, 9)
	conbuf   = make([]byte, 9)
	mousebuf = make([]byte, 5)
)

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
