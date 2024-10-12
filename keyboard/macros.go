package keyboard

import (
	"io"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func (kbd *Keyboard) SetMacroDriver(m MacrosDriver) {
	if rcv, ok := m.(EventReceiver); ok {
		kbd.AddEventReceiver(rcv)
	}
	kbd.macros.Driver = m
}

func (kbd *Keyboard) MacrosEnabled() bool {
	return kbd.macros.Enabled()
}

func (kbd *Keyboard) processMacro(key keycodes.Keycode, made bool) {
	if !kbd.MacrosEnabled() {
		return
	}
	if made {
		kbd.macros.ProcessKey(key, made)
	}
}

type Macros struct {
	// MacroDriver implementation to use; if nil, macros are disabled.
	Driver MacrosDriver
}

func (m *Macros) Enabled() bool {
	return m.Driver != nil && m.Driver.Count() > 0
}

func (m *Macros) ProcessKey(key keycodes.Keycode, made bool) {
	if !key.IsMacro() { // sanity check
		return
	}
	if !m.Enabled() {
		return
	}
	num := uint8(key - keycodes.QK_MACRO_0)
	m.Driver.RunMacro(num)
}

type MacrosDriver interface {
	Configure()
	Count() uint8
	RunMacro(macroNum uint8) (err error)
	Task(proc KeycodeProcessor)
	io.ReaderFrom
	io.WriterTo
	ZeroFill()
}

type KeycodeProcessor interface {
	ProcessKeycode(kc keycodes.Keycode, made bool)
	ClearKeycodes()
	Modifiers() KeyboardModifier
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=MacroCode

type MacroCode uint8

const (
	MacroCodeNone MacroCode = iota
	MacroCodeTap
	MacroCodeDown
	MacroCodeUp
	MacroCodeDelay
	MacroCodeVialExtTap
	MacroCodeVialExtDown
	MacroCodeVialExtUp
	MacroCodeSend
)

const MacroMagicPrefix uint8 = 1

type MacroError uint8

const (
	MacroErrInvalidNum MacroError = iota + 1
)

func (err MacroError) Error() string {
	switch err {
	case MacroErrInvalidNum:
		return "invalid macro index"
	default:
		return "unknown macro error"
	}
}
