//go:build macropad_rp2040

package macropad_rp2040

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/hosts/usbvial"
)

func NewVialHost(keymap keyboard.Keymap, matrix *keyboard.Matrix) *usbvial.Host {
	return usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)
}
