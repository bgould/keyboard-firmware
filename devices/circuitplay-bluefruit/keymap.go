package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func CircuitPlaygroundDefaultKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		CircuitPlaygroundLayer(BSPC, A),
	})
}

func CircuitPlaygroundLayer(K00, K01 Keycode) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		/*       0x0  0x1       */
		/************************/
		/* 0 */ {K00, K01},
		/************************/
	})
}
