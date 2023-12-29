package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	_______ = KC_TRANSPARENT
)

func Keymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		MacroPadRP2040Layer(
			/****************************************/
			/*   *\  ________                  /*   */
			/*   *\ |        | */ KC_VOLD, KC_VOLU, //
			/*   *\ |________|     */ KC_FN12, /*   */
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
		MacroPadRP2040Layer(
			/****************************************/
			/*   *\  ________                  /*   */
			/*   *\ |        | */ KC_BRID, KC_BRIU, //
			/*   *\ |________|     */ KC_FN12, /*   */
			/*   *\                            /*   */
			/*   */ _______, _______, _______, /*   */
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

func MacroPadRP2040Layer(
	/**/ k13, k14,
	/*     */ k00,
	k01, k02, k03,
	k04, k05, k06,
	k07, k08, k09,
	k10, k11, k12 Keycode,
) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		{
			k00, k01, k02, k03, k04, k05, k06, k07,
			k08, k09, k10, k11, k12, k13, k14, 0x0,
		},
	})
}
