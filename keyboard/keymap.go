package keyboard

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

type Keymap [][]keycodes.Keycode

func (keymap Keymap) KeyAt(position Pos) keycodes.Keycode {
	return keymap[position.Row][position.Col]
}
