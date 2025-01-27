//go:build circuitplay_bluefruit

package circuitplay_bluefruit

import (
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

func NewVialHost(keymap keyboard.Keymap, matrix *keyboard.Matrix, macros usbvial.VialMacroDriver) *usbvial.Host {
	return usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix, macros)
}
