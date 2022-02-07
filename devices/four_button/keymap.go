package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func FourButtonDefaultKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		// FourButtonLayer(W, A, S, D),
		FourButtonLayer(MS_UP, MS_L, MS_D, MS_R),
	})
}

func FourButtonLayer(
	/**/ K00, /**/
	K01, K02, K03 /**/ Keycode,
) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		/*       0x0  0x1  0x2  0x3  */
		/*****************************/
		/* 0 */ {K00, K01, K02, K03},
		/*****************************/
	})
}
