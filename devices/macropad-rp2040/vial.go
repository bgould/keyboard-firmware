package main

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

import (
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type VialKeyMapper struct {
}

func (v *VialKeyMapper) GetLayerCount() uint8 {
	return uint8(len(keymap))
}

func (v *VialKeyMapper) GetMaxKeyCount() int {
	return 13
}

func (v *VialKeyMapper) MapKey(layer, idx int) keycodes.Keycode {
	// println("MapKey: ", layer, idx)
	if idx > 13 || idx < 0 {
		return 0
	}
	if layer >= len(keymap) {
		return 0
	}
	return keymap[layer][0][idx]
}

//      [{"x":2, "w":0.5}, "1,0", {"w":0.5}, "1,1"],
