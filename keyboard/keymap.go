package keyboard

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

type Layer [][]keycodes.Keycode

func (layer Layer) KeyAt(position Pos) keycodes.Keycode {
	return layer[position.Row][position.Col]
}

type Keymap []Layer
