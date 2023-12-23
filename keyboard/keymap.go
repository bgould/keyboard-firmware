package keyboard

import (
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type Layer [][]keycodes.Keycode

func (layer Layer) KeyAt(row, col int) keycodes.Keycode {
	return layer[row][col]
}

type Keymap []Layer

func (keymap Keymap) GetLayerCount() uint8 {
	return uint8(len(keymap))
}

func (keymap Keymap) GetMaxKeyCount() int {
	return keymap.NumRows() * keymap.NumCols()
}

func (keymap Keymap) NumRows() int {
	return len(keymap[0])
}

func (keymap Keymap) NumCols() int {
	return len(keymap[0][0])
}

func (keymap Keymap) MapKey(layer, row, col int) (kc keycodes.Keycode) {
	if uint8(layer) >= keymap.GetLayerCount() || row >= keymap.NumRows() || col >= keymap.NumCols() {
		return
	}
	// numCols := keymap.NumCols()
	// row := idx / numCols
	// col := idx % numCols
	kc = keymap[layer][row][col]
	// println(layer, idx, row, col, kc)
	return
}

func (keymap Keymap) UpdateKey(layer, row, col int, kc keycodes.Keycode) bool {
	if uint8(layer) >= keymap.GetLayerCount() || row >= keymap.NumRows() || col >= keymap.NumCols() {
		return false
	}
	keymap[layer][row][col] = kc
	// println(layer, idx, row, col, kc)
	return true
}
