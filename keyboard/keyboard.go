package keyboard

import (
	"fmt"
	"io"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type Host interface {
	Send(report *Report)
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
}

type Keyboard struct {
	console Console
	matrix  *Matrix
	layers  []Keymap
	host    Host

	leds   uint8
	prev   []Row
	debug  bool
	report *Report

	jumpToBootloader func()
}

func New(console Console, host Host, matrix *Matrix, layers []Keymap) *Keyboard {
	return &Keyboard{
		console: console,
		matrix:  matrix,
		layers:  layers,
		host:    host,
		prev:    make([]Row, matrix.Rows()),
		report:  NewReport().Keyboard(0),
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
		fmt.Fprintf(kbd.console,
			"event => loc: r=%X c=%X, made: %t, usb: %02X, mod: %t, key: %t\r\n",
			ev.Pos.Row, ev.Pos.Col, ev.Made, key, key.IsModifier(), key.IsKey(),
		)
	}
	if key == keycodes.BOOTLOADER {
		if ev.Made && kbd.jumpToBootloader != nil {
			kbd.jumpToBootloader()
		}
	}
	if ev.Made {
		kbd.report.Make(key)
	} else {
		kbd.report.Break(key)
	}
	if kbd.debug {
		fmt.Fprintf(kbd.console, "report => %s\r\n", kbd.report.String())
	}
	kbd.host.Send(kbd.report)
}

func (kbd *Keyboard) debugMatrix() bool {
	if kbd.debug {
		kbd.matrix.Print(kbd.console)
		return true
	}
	return false
}
