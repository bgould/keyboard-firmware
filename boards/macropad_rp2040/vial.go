//go:build macropad_rp2040

package macropad_rp2040

import (
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def -package macropad_rp2040 vial.json

func NewVialHost(keymap keyboard.Keymap, matrix *keyboard.Matrix, macros usbvial.VialMacroDriver) *usbvial.Host {
	return usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix, macros)
}

func (dev *Board) NewVialKeyboard(layers ...keyboard.Layer) (*keyboard.Keyboard, *usbvial.Host) {
	keymap := Keymap(layers...)
	matrix := dev.NewMatrix()
	macros, ok := keyboard.NewDefaultMacroDriver(32, 4096).(usbvial.VialMacroDriver)
	if !ok {
		macros = nil
	}
	host := NewVialHost(keymap, matrix, macros)
	kbd := keyboard.New(host, matrix, keymap)
	kbd.SetMacroDriver(macros)
	rgb := usbvial.NewKeyboardVialRGBer(kbd)
	host.UseVialRGB(rgb)
	return kbd, host
}
