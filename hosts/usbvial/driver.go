package usbvial

import (
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func NewKeyboard(def vial.DeviceDefinition, keymap keyboard.Keymap, matrix *keyboard.Matrix) *Host {
	host = &Host{
		Host: usbhid.New(),
		dev: vial.NewDevice(def, &KeyboardDeviceDriver{
			keymap: keymap,
			matrix: matrix,
		}),
	}
	return host
}

type KeyboardDeviceDriver struct {
	keymap keyboard.Keymap
	matrix *keyboard.Matrix
}

var (
	_ vial.DeviceDriver  = (*KeyboardDeviceDriver)(nil)
	_ vial.EncoderMapper = (*KeyboardDeviceDriver)(nil)
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
