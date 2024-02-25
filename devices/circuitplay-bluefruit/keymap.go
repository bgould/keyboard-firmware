package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func CircuitPlaygroundDefaultKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		CircuitPlaygroundLayer(KC_BSPC, KC_A___, KC_FN1_),
		CircuitPlaygroundLayer(BL_TOGG, BL_BRTG, KC_FN1_),
	})
}

func CircuitPlaygroundLayer(K00, K01, K02 Keycode) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		/*       0x0  0x1       */
		/************************/
		/* 0 */ {K00, K01, K02},
		/************************/
	})
}
