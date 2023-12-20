package main

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

import (
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
	"github.com/bgould/keyboard-firmware/matrix/kinx/kintqt"
)

type VialKeyMapper struct {
}

func (v *VialKeyMapper) GetLayerCount() uint8 {
	return uint8(len(keymap))
}

func (v *VialKeyMapper) GetMaxKeyCount() int {
	return kintqt.NumRows * kintqt.NumCols
}

func (v *VialKeyMapper) MapKey(layer, idx int) (kc keycodes.Keycode) {
	if layer >= len(keymap) {
		return
	}
	row := idx / kintqt.NumCols
	col := idx % kintqt.NumCols
	if row < len(keymap[layer]) && col < len(keymap[layer][row]) {
		kc = keymap[layer][row][col]
	}
	// println(layer, idx, row, col, kc)
	return
}
