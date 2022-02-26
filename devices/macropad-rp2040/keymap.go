package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func Keymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{MacroPadRP2040Layer(
		/*       */ NLCK,
		KP_7, KP_8, KP_9,
		KP_4, KP_5, KP_6,
		KP_1, KP_2, KP_3,
		PDOT, KP_0, PENT,
	)})
}

func MacroPadRP2040Layer(
	/*     */ k00,
	k01, k02, k03,
	k04, k05, k06,
	k07, k08, k09,
	k10, k11, k12 Keycode,
) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		{
			k00, k01, k02, k03, k04, k05, k06, k07,
			k08, k09, k10, k11, k12, 0x0, 0x0, 0x0,
		},
	})
}
