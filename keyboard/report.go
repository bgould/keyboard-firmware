package keyboard

import (
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	RptKeyboard ReportType = 0x0
	RptMouse    ReportType = 0x2
	RptConsumer ReportType = 0x3
)

type ReportType byte

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

func NewReport() *Report {
	return new(Report)
}

func (r *Report) Make(key keycodes.Keycode) {
	r[1] = byte(RptKeyboard)
	if key.IsModifier() {
		r[0] |= 1 << (key & 0x07)
		return
	}
	firstZero := 0
	for i := 2; i < 8; i++ {
		switch keycodes.Keycode(r[i]) {
		case keycodes.NO:
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

func (r *Report) Keyboard(mod KeyboardModifier, keys ...byte) *Report {
	r[0] = byte(mod)
	r[1] = byte(RptKeyboard)
	for i, c := 0, len(keys); i < 6; i++ {
		if i < c {
			r[i+2] = keys[i]
		} else {
			r[i+2] = 0x0
		}
	}
	return r
}

type MouseButton byte

const (
	MouseBtnLeft   MouseButton = 0x01
	MouseBtnRight  MouseButton = 0x02
	MouseBtnMiddle MouseButton = 0x04
)

func (r *Report) Mouse(buttons MouseButton, x int8, y int8) *Report {
	r[0] = 0x0
	r[1] = byte(RptMouse)
	r[2] = byte(buttons)
	r[3] = byte(x)
	r[4] = byte(y)
	r[5] = 0x0
	r[6] = 0x0
	r[7] = 0x0
	return r
}

type ConsumerKey uint16

const (
	ConsKeyHome       ConsumerKey = 0x0100
	ConsKeyKbdLayout  ConsumerKey = 0x0200
	ConsKeySearch     ConsumerKey = 0x0400
	ConsKeySnapshot   ConsumerKey = 0x0800
	ConsKeyVolUp      ConsumerKey = 0x1000
	ConsKeyVolDown    ConsumerKey = 0x2000
	ConsKeyPlayPause  ConsumerKey = 0x4000
	ConsKeyFastFwd    ConsumerKey = 0x8000
	ConsKeyRewind     ConsumerKey = 0x0001
	ConsKeyNextTrack  ConsumerKey = 0x0002
	ConsKeyPrevTrack  ConsumerKey = 0x0004
	ConsKeyRandomPlay ConsumerKey = 0x0008
	ConsKeyStop       ConsumerKey = 0x0010
)

func (r *Report) Consumer(key ConsumerKey) *Report {
	r[0] = 0x0
	r[1] = byte(RptConsumer)
	r[2] = byte(key >> 8)
	r[3] = byte(key & 0xFF)
	r[4] = 0x0
	r[5] = 0x0
	r[6] = 0x0
	r[7] = 0x0
	return r
}

func (r *Report) String() string {
	return "[" +
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

func hex(b uint8) string {
	s := strconv.FormatUint(uint64(b), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}
