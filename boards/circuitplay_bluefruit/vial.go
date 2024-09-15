//go:build circuitplay_bluefruit

package circuitplay_bluefruit

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/hosts/usbvial"
)

func NewVialHost(keymap keyboard.Keymap, matrix *keyboard.Matrix) *usbvial.Host {
	return usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)
}
