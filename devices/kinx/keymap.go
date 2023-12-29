package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/matrix/kinx"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	_______ = KC_NO
)

// FN0 -> KC_Toggle "Keypad" KC_layer on KC_key press
// FN1 -> KC_Toggle "Programming" KC_on and KC_off on KC_key up/down
// FN2 -> KC_CPU Reset KC_on key down

func initKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		// 0 - KC_Default Layer
		kinx.Layer(
			KC_FN3_, KC_F1__, KC_F2__, KC_F3__, KC_F4__, KC_F5__, KC_F6__, KC_F7__, KC_F8__ /*    */, KC_F9__, KC_F10_, KC_F11_, KC_F12_, KC_PSCR, KC_SCRL, KC_PAUS, KC_FN0_, KC_FN1_,
			KC_EQL_, KC_N1__, KC_N2__, KC_N3__, KC_N4__, KC_N5__ /*                                                          */, KC_N6__, KC_N7__, KC_N8__, KC_N9__, KC_N0__, KC_MINS,
			KC_TAB_, KC_Q___, KC_W___, KC_E___, KC_R___, KC_T___ /*                                                          */, KC_Y___, KC_U___, KC_I___, KC_O___, KC_P___, KC_BSLS,
			KC_RCTL, KC_A___, KC_S___, KC_D___, KC_F___, KC_G___ /*                                                          */, KC_H___, KC_J___, KC_K___, KC_L___, KC_SCLN, KC_QUOT,
			KC_LSFT, KC_Z___, KC_X___, KC_C___, KC_V___, KC_B___ /*                                                          */, KC_N___, KC_M___, KC_COMM, KC_DOT_, KC_SLSH, KC_RSFT,
			/*    */ KC_GRV_, KC_INS_, KC_LEFT, KC_RGHT /*                                                                            */, KC_UP__, KC_DOWN, KC_LBRC, KC_RBRC,
			/*                                        */ KC_ESC_, KC_LGUI /*                                        */, KC_LALT, KC_LCTL,
			/*                                                 */ KC_HOME /*                                        */, KC_PGUP,
			/*                               */ KC_BSPC, KC_DEL_, KC_END_ /*                                        */, KC_PGDN, KC_ENT_, KC_SPC_, /**/
		),
		// 1 - KC_Keypad Layer
		kinx.Layer(
			KC_ESC_, _______, _______, _______, KC_MPLY, KC_MPRV, KC_MNXT, _______, _______ /*    */, _______, _______, _______, _______, KC_MUTE, KC_VOLD, KC_VOLU, KC_FN0_, KC_FN1_,
			_______, _______, _______, _______, _______, _______ /*                                                          */, _______, KC_NUM_, KC_PEQL, KC_PSLS, KC_PAST, _______,
			_______, _______, _______, KC_MS_U, _______, _______ /*                                                          */, _______, KC_KP_7, KC_KP_8, KC_KP_9, KC_PMNS, _______,
			KC_CAPS, _______, KC_MS_L, KC_MS_D, KC_MS_R, _______ /*                                                          */, _______, KC_KP_4, KC_KP_5, KC_KP_6, KC_PPLS, _______,
			_______, _______, _______, _______, _______, _______ /*                                                          */, _______, KC_KP_1, KC_KP_2, KC_KP_3, KC_PENT, _______,
			/*    */ _______, KC_INS_, KC_LEFT, KC_RGHT /*                                                                            */, KC_UP__, KC_DOWN, KC_DOT_, KC_PENT,
			/*                                        */ KC_BTN1, KC_BTN2 /*                                        */, KC_BTN4, KC_BTN3,
			/*                                                 */ KC_HOME /*                                        */, KC_PGUP,
			/*                               */ KC_BSPC, KC_DEL_, KC_END /*                                         */, KC_PGDN, KC_ENT_, KC_KP_0, /**/
		),
		// 2 - KC_Programming Layer
		kinx.Layer(
			_______, _______, _______, _______, _______, _______, _______, _______, _______ /*    */, KC_FN2_, _______, _______, _______, _______, _______, _______, _______, KC_FN1_,
			_______, _______, _______, _______, _______, _______ /*                                                          */, _______, _______, _______, _______, _______, _______,
			_______, _______, _______, _______, _______, _______ /*                                                          */, _______, _______, _______, _______, _______, _______,
			_______, _______, _______, _______, _______, _______ /*                                                          */, _______, _______, _______, _______, _______, _______,
			_______, _______, _______, _______, _______, _______ /*                                                          */, _______, _______, _______, _______, _______, _______,
			/*    */ _______, _______, _______, _______ /*                                                                            */, _______, _______, _______, _______,
			/*                                        */ _______, _______ /*                                        */, _______, _______,
			/*                                                 */ _______ /*                                        */, _______,
			/*                               */ _______, _______, _______ /*                                        */, _______, _______, _______, /**/
		),
	})
}
