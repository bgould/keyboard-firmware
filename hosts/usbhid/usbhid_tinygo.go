//go:build tinygo && !teensy41

package usbhid

import (
	"machine/usb/hid"
	"runtime/volatile"
)

var port = &kbport{}

// TODO: consider moving out of init
func init() {
	hid.SetHandler(port)
}

var (
	keybuf   = make([]byte, 9)
	conbuf   = make([]byte, 9)
	mousebuf = make([]byte, 5)

	ledState volatile.Register8
)

func sendKeyboardReport(mod, k1, k2, k3, k4, k5, k6 byte) {
	// var b [9]byte
	// b[0] = 0x02
	// b[1] = mod
	// b[2] = 0
	// b[3] = k1
	// b[4] = k2
	// b[5] = k3
	// b[6] = k4
	// b[7] = k5
	// b[8] = k6
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

	// b := []byte{0x02, mod, 0, k1, k2, k3, k4, k5, k6}
	// if port.txc {
	// 	port.buf.Put(b)
	// } else {
	// 	port.txc = true
	// 	machine.SendUSBInPacket(usb.HID_ENDPOINT_IN, b)
	// 	// hid.SendUSBPacket(b)
	// }

	//port.tx([]byte{0x02, mod, 0, k1, k2, k3, k4, k5, k6})
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
	ledState.Set(rx[1])
	return true
}
