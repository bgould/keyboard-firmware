//go:build tinygo

package usbvial

import (
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type VialMacroDriver interface {
	VialMacroCount() uint8
	VialMacroBuffer() []byte
}

func NewKeyboard(def vial.DeviceDefinition, keymap keyboard.Keymap, matrix *keyboard.Matrix, macros VialMacroDriver) *Host {
	host = &Host{
		Host: usbhid.New(),
		dev: vial.NewDevice(def, &KeyboardDeviceDriver{
			keymap: keymap,
			matrix: matrix,
			macros: macros,
		}),
	}
	return host
}

type KeyboardDeviceDriver struct {
	keymap keyboard.Keymap
	matrix *keyboard.Matrix
	macros VialMacroDriver
}

func NewKeyboardDriver(keymap keyboard.Keymap, matrix *keyboard.Matrix) *KeyboardDeviceDriver {
	return &KeyboardDeviceDriver{
		keymap: keymap,
		matrix: matrix,
	}
}

var (
	_ vial.DeviceDriver  = (*KeyboardDeviceDriver)(nil)
	_ vial.EncoderMapper = (*KeyboardDeviceDriver)(nil)
	_ vial.MacroDriver   = (*KeyboardDeviceDriver)(nil)
)

func (kbd *KeyboardDeviceDriver) GetLayerCount() uint8 {
	return kbd.keymap.GetLayerCount()
}

func (kbd *KeyboardDeviceDriver) MapKey(layer, row, col int) keycodes.Keycode {
	return kbd.keymap.MapKey(layer, row, col)
}

// TODO: Keep track of "dirty" keys and implement keypress for saving
func (kbd *KeyboardDeviceDriver) SetKey(layer, row, col int, kc keycodes.Keycode) bool {
	return kbd.keymap.SetKey(layer, row, col, kc)
}

func (kbd *KeyboardDeviceDriver) GetMatrixRowState(idx int) uint32 {
	return uint32(kbd.matrix.GetRow(uint8(idx)))
}

func (kbd *KeyboardDeviceDriver) MapEncoder(idx int) (ccwRow, ccwCol, cwRow, cwCol int, ok bool) {
	if ccw, cw, ok := kbd.matrix.MapEncoder(idx); ok {
		return int(ccw.Row), int(ccw.Col), int(cw.Row), int(cw.Col), true
	}
	return
}

func (kbd *KeyboardDeviceDriver) GetMacroCount() uint8 {
	if kbd.macros == nil {
		return 0
	}
	return kbd.macros.VialMacroCount()
}

func (kbd *KeyboardDeviceDriver) GetMacroBuffer() []byte {
	if kbd.macros == nil {
		return []byte{}
	}
	return kbd.macros.VialMacroBuffer()
}
