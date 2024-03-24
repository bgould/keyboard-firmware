package circuitplay_bluefruit

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def -package circuitplay_bluefruit vial.json

func Keymap(layers ...keyboard.Layer) keyboard.Keymap {
	if len(layers) == 0 {
		return DefaultKeymap()
	}
	return layers
}

func DefaultKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		Layer(KC_BSPC, KC_A___, MO_(01)),
		Layer(BL_TOGG, BL_STEP, MO_(01)),
	})
}

func Layer(K00, K01, K02 Keycode) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		/*       0x0  0x1       */
		/************************/
		/* 0 */ {K00, K01, K02},
		/************************/
	})
}
