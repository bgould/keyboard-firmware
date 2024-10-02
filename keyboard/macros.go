package keyboard

import (
	"os"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func (kbd *Keyboard) SetMacroDriver(m MacrosDriver) {
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
	println("running macro: ", num, key)
	m.Driver.RunMacro(num)
	// start, end, ok := m.macroNumBounds(num)
	// if ok {
	// 	xxdfprint(os.Stdout, 0x0, m.GetBytes(num))
	// }
}

type MacrosDriver interface {
	Configure()
	Count() uint8
	RunMacro(macroNum uint8) (err error)
	Task()
}

func NewDefaultMacroDriver(count uint8, bufferSize uint16) MacrosDriver {
	return &defaultMacroDriver{count: count, buffer: make([]byte, bufferSize)}
}

type defaultMacroDriver struct {
	count  uint8
	buffer []byte
}

func (m *defaultMacroDriver) Configure() {

}

func (m *defaultMacroDriver) RunMacro(macroNum uint8) (err error) {
	buf := m.macroBytes(macroNum)
	xxdfprint(os.Stdout, 0x0, buf)
	for _, b := range buf {
		switch b {
		case MacroMagicPrefix:

		default:

		}
	}
	return nil
}

var macroEmptyBuf [0]byte

func (m *defaultMacroDriver) macroBytes(macroNum uint8) []byte {
	start, end, ok := m.macroNumBounds(macroNum)
	if !ok {
		return macroEmptyBuf[:]
	}
	return m.buffer[start:end]
}

// determine bounds of specified macro in buffer
func (m *defaultMacroDriver) macroNumBounds(macroNum uint8) (start, end int, ok bool) {
	if len(m.buffer) == 0 {
		return 0, 0, false
	}
	if macroNum >= m.count {
		return 0, 0, false
	}
	for i, c, n := 0, len(m.buffer), uint8(0); i < c; i++ {
		if b := m.buffer[i]; b == 0x0 {
			if n == macroNum {
				end = i
				return start, end, true
			} else {
				n++
				start = i + 1
			}
		}
	}
	return start, end, false
}

// Task
func (m *defaultMacroDriver) Task() {

}

func (m *defaultMacroDriver) Count() uint8 {
	return m.count
}

func (m *defaultMacroDriver) VialMacroBuffer() []byte {
	return m.buffer
}

// func (m *Macros) Buffer() []byte {
// 	return m.buffer
// }

type MacroCode uint8

const (
	MacroCodeTap = iota + 1
	MacroCodeDown
	MacroCodeUp
	MacroCodeDelay
	MacroCodeVialExtTap
	MacroCodeVialExtDown
	MacroCodeVialExtUp
)

const MacroMagicPrefix uint8 = 1
