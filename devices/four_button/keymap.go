package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

// TODO: refactor/fix interfaces
// FIXME: actually a "Keymap"
func FourButtonDefaultLayer() keyboard.Keymap {
	return FourButtonKeymap(W, A, S, D)
}

// FIXME: actually a "Layer"
func FourButtonKeymap(
	/**/ K00, /**/
	K01, K02, K03 /**/ Keycode,
) keyboard.Keymap {
	return keyboard.Keymap([][]Keycode{
		/*       0x0  0x1  0x2  0x3  */
		/*****************************/
		/* 0 */ {K00, K01, K02, K03},
		/*****************************/
	})
}
