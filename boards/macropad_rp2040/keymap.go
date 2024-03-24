package macropad_rp2040

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def -package macropad_rp2040 vial.json

const (
	_______ = KC_TRANSPARENT
)

func Keymap(layers ...keyboard.Layer) keyboard.Keymap {
	if len(layers) == 0 {
		return DefaultKeymap()
	}
	return layers
}

func DefaultKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		Layer(
			/****************************************/
			/*   *\  ________                  /*   */
			/*   *\ |        | */ KC_VOLD, KC_VOLU, //
			/*   *\ |________|     */ TG_(01), /*   */
			/*   *\                            /*   */
			/*   */ KC_KP_7, KC_KP_8, KC_KP_9, /*   */
			/*   *\                            /*   */
			/*   */ KC_KP_4, KC_KP_5, KC_KP_6, /*   */
			/*   *\                            /*   */
			/*   */ KC_KP_1, KC_KP_2, KC_KP_3, /*   */
			/*   *\                            /*   */
			/*   */ KC_PDOT, KC_KP_0, KC_PENT, /*   */
			/*   *\                            /*   */
			/****************************************/
		),
		Layer(
			/****************************************/
			/*   *\  ________                  /*   */
			/*   *\ |        | */ KC_BRID, KC_BRIU, //
			/*   *\ |________|     */ TG_(01), /*   */
			/*   *\                            /*   */
			/*   */ KC_F1__, _______, _______, /*   */
			/*   *\                            /*   */
			/*   */ _______, _______, _______, /*   */
			/*   *\                            /*   */
			/*   */ _______, _______, _______, /*   */
			/*   *\                            /*   */
			/*   */ _______, _______, _______, /*   */
			/*   *\                            /*   */
			/****************************************/
		),
	})
}
