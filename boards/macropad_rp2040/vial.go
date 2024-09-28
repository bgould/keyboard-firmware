//go:build macropad_rp2040

package macropad_rp2040

import (
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

func init() {
	VialDeviceDefinition.UnlockKeys = []vial.Pos{
		{Row: 0, Col: 2},
	}
}

func NewVialHost(keymap keyboard.Keymap, matrix *keyboard.Matrix) *usbvial.Host {
	return usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)
}
