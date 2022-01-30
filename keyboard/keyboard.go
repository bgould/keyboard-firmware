package keyboard

import (
	"io"
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type Host interface {
	Send(report *Report)
	LEDs() uint8
}

type Event struct {
	Pos  Pos
	Made bool
	Time uint32
}

type Pos struct {
	Row uint8
	Col uint8
}

type Console interface {
	io.ReadWriter
	Buffered() int
}

type Keyboard struct {
	console Console
	matrix  *Matrix
	layers  []Layer
	host    Host

	prev []Row
	leds uint8

	keyReport      *Report
	mouseReport    *Report
	consumerReport *Report

	debug bool

	jumpToBootloader func()
}

func New(console Console, host Host, matrix *Matrix, keymap Keymap) *Keyboard {
	return &Keyboard{
		console:        console,
		matrix:         matrix,
		layers:         keymap,
		host:           host,
		prev:           make([]Row, matrix.Rows()),
		keyReport:      NewReport().Keyboard(0),
		mouseReport:    NewReport().Keyboard(0),
		consumerReport: NewReport().Keyboard(0),
	}
}

func (kbd *Keyboard) WithDebug(dbg bool) *Keyboard {
	kbd.debug = dbg
	return kbd
}

func (kbd *Keyboard) WithJumpToBootloader(fn func()) *Keyboard {
	kbd.jumpToBootloader = fn
	return kbd
}

func (kbd *Keyboard) LEDs() uint8 {
	return kbd.host.LEDs()
}

func (kbd *Keyboard) Task() {
	kbd.matrix.Scan()
	for i, rows := uint8(0), kbd.matrix.Rows(); i < rows; i++ {
		row := kbd.matrix.GetRow(i)
		diff := row ^ kbd.prev[i]
		if diff == 0 {
			continue
		}
		if kbd.matrix.HasGhostInRow(i) {
			continue
		}
		kbd.debugMatrix()
		for j, cols := uint8(0), kbd.matrix.Cols(); j < cols; j++ {
			mask := Row(1) << j
			if diff&mask > 0 {
				ev := Event{
					Pos:  Pos{i, j},
					Made: row&mask > 0,
				}
				kbd.processEvent(ev)
				kbd.prev[i] ^= mask
			}
		}
	}
}

func (kbd *Keyboard) processEvent(ev Event) {
	key := kbd.layers[0].KeyAt(ev.Pos)
	if kbd.debug {
		kbd.console.Write([]byte(
			"event => " +
				"loc: r=" + hex(ev.Pos.Row) + " c= " + hex(ev.Pos.Col) + ", " +
				"made: " + strconv.FormatBool(ev.Made) + ", " +
				"usb: " + hex(uint8(key)) + ", " +
				"key: " + strconv.FormatBool(key.IsKey()) + ", " +
				"mod: " + strconv.FormatBool(key.IsModifier()) + ", " +
				"msk: " + strconv.FormatBool(key.IsMouseKey()) + ", " +
				"cns: " + strconv.FormatBool(key.IsConsumer()) + ", " +
				"sys: " + strconv.FormatBool(key.IsSystem()) + ", " +
				"spc: " + strconv.FormatBool(key.IsSpecial()) + "\r\n"))
	}

	switch {
	case key.IsKey() || key.IsModifier():
		kbd.processKey(key, ev.Made)
	case key.IsMouseKey():
		kbd.processMouseKey(key, ev.Made)
	case key.IsConsumer():
		kbd.processConsumerKey(key, ev.Made)
	case key.IsSystem():
		kbd.processSystemKey(key, ev.Made)
	case key.IsSpecial():
		kbd.processSpecialKey(key, ev.Made)
	}
}

func (kbd *Keyboard) processKey(key keycodes.Keycode, made bool) {
	if made {
		kbd.keyReport.Make(key)
	} else {
		kbd.keyReport.Break(key)
	}
	if kbd.debug {
		kbd.console.Write([]byte("keyboard report => " + kbd.keyReport.String() + "\r\n"))
	}
	kbd.host.Send(kbd.keyReport)
}

func (kbd *Keyboard) processMouseKey(key keycodes.Keycode, made bool) {
	if kbd.debug {
		kbd.console.Write([]byte("mouse report => " + kbd.mouseReport.String() + "\r\n"))
	}
}

func (kbd *Keyboard) processConsumerKey(key keycodes.Keycode, made bool) {
	if kbd.debug {
		kbd.console.Write([]byte("consumer report => " + kbd.consumerReport.String() + "\r\n"))
	}
}

func (kbd *Keyboard) processSystemKey(key keycodes.Keycode, made bool) {
	if kbd.debug {
		kbd.console.Write([]byte("system report => " + kbd.consumerReport.String() + "\r\n"))
	}
}

func (kbd *Keyboard) processSpecialKey(key keycodes.Keycode, made bool) {
	switch key {
	case keycodes.BOOTLOADER:
		if !made {
			break
		}
		if kbd.jumpToBootloader != nil {
			if kbd.debug {
				kbd.console.Write([]byte("jumping to bootloader"))
			}
			kbd.jumpToBootloader()
			if kbd.debug {
				kbd.console.Write([]byte("notice: jump to bootloader appears to have failed"))
			}
		} else {
			if kbd.debug {
				kbd.console.Write([]byte("notice: no jumpToBootloader callback defined"))
			}
		}
		return
	}
	if kbd.debug {
		kbd.console.Write([]byte("special key => " +
			hex(uint8(key)) + ", made: " + strconv.FormatBool(made) + "\r\n"))
	}
}

func (kbd *Keyboard) debugMatrix() bool {
	if kbd.debug {
		kbd.matrix.Print(kbd.console)
		return true
	}
	return false
}