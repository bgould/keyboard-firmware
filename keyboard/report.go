package keyboard

import (
	"io"
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	RptKeyboard ReportType = 0x0
	RptMouse    ReportType = 0x2
	RptConsumer ReportType = 0x3
)

type ReportType byte

func (t ReportType) String() string {
	switch t {
	case RptKeyboard:
		return "Keyboard"
	case RptMouse:
		return "Mouse"
	case RptConsumer:
		return "Consumer"
	default:
		return "Unknown"
	}
}

type Report [8]byte

type KeyboardModifier byte

const (
	KbdModNone       KeyboardModifier = 0x0
	KbdModCtrlLeft   KeyboardModifier = 1 << 0
	KbdModShiftLeft  KeyboardModifier = 1 << 1
	KbdModAltLeft    KeyboardModifier = 1 << 2
	KbdModGuiLeft    KeyboardModifier = 1 << 3
	KbdModCtrlRight  KeyboardModifier = 1 << 4
	KbdModShiftRight KeyboardModifier = 1 << 5
	KbdModAltRight   KeyboardModifier = 1 << 6
	KbdModGuiRight   KeyboardModifier = 1 << 7
)

func (r *Report) Type() ReportType {
	return ReportType(r[1])
}

func (r *Report) Make(key keycodes.Keycode) {

	if key.IsConsumer() {
		r[1] = byte(RptConsumer)
		if consumer := keycode2consumer(key); r[0] == 0 && consumer > 0 {
			r[0] = 1
			r[2] = uint8(consumer)
			r[3] = uint8(consumer >> 8)
		}
		return
	}

	r[1] = byte(RptKeyboard)
	if key.IsModifier() {
		r[0] |= 1 << (key & 0x07)
		return
	}
	firstZero := 0
	for i := 2; i < 8; i++ {
		switch keycodes.Keycode(r[i]) {
		case keycodes.KC_NO:
			if firstZero == 0 {
				firstZero = i
			}
		case key:
			return
		}
	}
	if firstZero > 0 {
		r[firstZero] = byte(key)
	}
}

func (r *Report) Break(key keycodes.Keycode) {

	if key.IsConsumer() {
		r[1] = byte(RptConsumer)
		if r[0] == 0 {
			return
		}
		if c := keycode2consumer(key); r[2] == uint8(c) && r[3] == uint8(c>>8) {
			r[0] = 0
			r[2] = 0
			r[3] = 0
		}
		return
	}

	r[1] = byte(RptKeyboard)
	if key.IsModifier() {
		r[0] &= ^(1 << (key & 0x07))
		return
	}
	for i := 2; i < 8; i++ {
		if r[i] == byte(key) {
			r[i] = 0x0
		}
	}

}

func (r *Report) Keyboard(mod KeyboardModifier, keys ...byte) {
	r[0] = byte(mod)
	r[1] = byte(RptKeyboard)
	for i, c := 0, len(keys); i < 6; i++ {
		if i < c {
			r[i+2] = keys[i]
		} else {
			r[i+2] = 0x0
		}
	}
}

type MouseButton byte

const (
	MouseBtnLeft   MouseButton = 0x01
	MouseBtnRight  MouseButton = 0x02
	MouseBtnMiddle MouseButton = 0x04

	mouseButtons = 2
	mouseX       = 3
	mouseY       = 4
	mouseV       = 5
	mouseH       = 6
)

// TODO: implement wheel, accel
func (r *Report) Mouse(buttons MouseButton, x int8, y int8, v int8, h int8) {
	r[0] = 0x0
	r[1] = byte(RptMouse)
	r[mouseButtons] = byte(buttons)
	r[mouseX] = byte(x)
	r[mouseY] = byte(y)
	r[mouseV] = byte(v)
	r[mouseH] = byte(h)
	r[7] = 0x0
}

func (r *Report) String() string {
	return r.Type().String() + "[" +
		" " + hex(r[0]) +
		" " + hex(r[1]) +
		" " + hex(r[2]) +
		" " + hex(r[3]) +
		" " + hex(r[4]) +
		" " + hex(r[5]) +
		" " + hex(r[6]) +
		" " + hex(r[7]) +
		" ]"
}

func (r *Report) WriteDebug(w io.Writer) {
	for i := 0; i < 8; i++ {
		w.Write([]byte(hex(r[i])))
	}
}

func hex(b uint8) string {
	s := strconv.FormatUint(uint64(b), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}
