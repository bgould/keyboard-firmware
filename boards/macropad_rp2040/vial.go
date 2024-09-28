//go:build macropad_rp2040

package macropad_rp2040

import (
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

func NewVialHost(keymap keyboard.Keymap, matrix *keyboard.Matrix, macros *keyboard.Macros) *usbvial.Host {
	return usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix, macros)
}
