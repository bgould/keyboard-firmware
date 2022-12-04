package usbhid

import (
	"machine"
	"machine/usb/hid"
	"strconv"
)

var port *kbport

type kbport struct {
	buf     *hid.RingBuffer
	waitTxc bool
	dbg     bool
}

// TODO: consider moving out of init
func init() {
	if port == nil {
		port = newKeyboard()
		hid.SetHandler(port)
	}
}

// Port returns the USB hid-keyboard port.
func Port() *kbport {
	return port
}

func newKeyboard() *kbport {
	return &kbport{
		buf: hid.NewRingBuffer(),
		dbg: true,
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

func (kb *kbport) tx(b []byte) {
	if kb.waitTxc {
		kb.buf.Put(b)
	} else {
		kb.waitTxc = true
		writeDebug(b)
		hid.SendUSBPacket(b)
	}
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

/*
func (kb *kbport) ready() bool {
	return true
}

func (kb *kbport) sendKey(consumer bool, b []byte) bool {
	kb.tx(b)
	return true
}

func (kb *kbport) keyboardSendKeys(consumer bool) bool {
	var b [9]byte
	b[0] = 0x02
	b[1] = kb.mod
	b[2] = 0x02
	b[3] = kb.key[0]
	b[4] = kb.key[1]
	b[5] = kb.key[2]
	b[6] = kb.key[3]
	b[7] = kb.key[4]
	b[8] = kb.key[5]
	return kb.sendKey(consumer, b[:])
}
*/
