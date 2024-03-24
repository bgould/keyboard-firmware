//go:build keymap.custom

package main

import (
	"github.com/bgould/keyboard-firmware/boards/macropad_rp2040"
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	_______ = KC_TRANSPARENT
)

var layers = []keyboard.Layer{
	macropad_rp2040.Layer(
		/****************************************/
		/*   *\  ________                  /*   */
		/*   *\ |        | */ KC_VOLD, KC_VOLU, //
		/*   *\ |________|     */ TG_(01), /*   */
		/*   *\                            /*   */
		/*   */ KC_F1__, KC_F2__, KC_F3__, /*   */
		/*   *\                            /*   */
		/*   */ KC_F4__, KC_F5__, KC_F6__, /*   */
		/*   *\                            /*   */
		/*   */ KC_F7__, KC_F8__, KC_F9__, /*   */
		/*   *\                            /*   */
		/*   */ KC_F10_, KC_F11_, KC_F12_, /*   */
		/*   *\                            /*   */
		/****************************************/
	),
	macropad_rp2040.Layer(
		/****************************************/
		/*   *\  ________                  /*   */
		/*   *\ |        | */ BL_DOWN, BL_UP__, //
		/*   *\ |________|     */ TG_(01), /*   */
		/*   *\                            /*   */
		/*   */ KC_F1__, BL_STEP, BL_BRTG, /*   */
		/*   *\                            /*   */
		/*   */ BL_OFF_, BL_TOGG, BL_ON__, /*   */
		/*   *\                            /*   */
		/*   */ RGB_HUI, RGB_SAI, RGB_VAI, /*   */
		/*   *\                            /*   */
		/*   */ RGB_HUD, RGB_SAD, RGB_VAD, /*   */
		/*   *\                            /*   */
		/****************************************/
	),
}
