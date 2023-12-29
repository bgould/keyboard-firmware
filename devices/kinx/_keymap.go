package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/matrix/kinx"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	____ = NO
)

// FN0 -> Toggle "Keypad" layer on key press
// FN1 -> Toggle "Programming" on and off on key up/down
// FN2 -> CPU Reset on key down

func initKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		// 0 - Default Layer
		kinx.Layer(
			FN3_, F1__, F2__, F3__, F4__, F5__, F6__, F7__, F8__ /*    */, F9__, F10_, F11_, F12_, PSCR, SCRL, PAUS, FN0_, FN1_,
			EQL_, N1__, N2__, N3__, N4__, N5__ /*                                        */, N6__, N7__, N8__, N9__, N0__, MINS,
			TAB_, Q___, W___, E___, R___, T___ /*                                        */, Y___, U___, I___, O___, P___, BSLS,
			RCTL, A___, S___, D___, F___, G___ /*                                        */, H___, J___, K___, L___, SCLN, QUOT,
			LSFT, Z___, X___, C___, V___, B___ /*                                        */, N___, M___, COMM, DOT_, SLSH, RSFT,
			/* */ GRV_, INS_, LEFT, RGHT /*                                                    */, UP__, DOWN, LBRC, RBRC,
			/*                         */ ESC_, LGUI /*                            */, LALT, LCTL,
			/*                               */ HOME /*                            */, PGUP,
			/*                   */ BSPC, DEL_, END_ /*                            */, PGDN, ENT_, SPC_, /**/
		),
		// 1 - Keypad Layer
		kinx.Layer(
			ESC_, ____, ____, ____, MPLY, MPRV, MNXT, ____, ____ /*    */, ____, ____, ____, ____, MUTE, VOLD, VOLU, FN0_, FN1_,
			____, ____, ____, ____, ____, ____ /*                                        */, ____, NUM_, PEQL, PSLS, PAST, ____,
			____, ____, ____, MS_U, ____, ____ /*                                        */, ____, KP_7, KP_8, KP_9, PMNS, ____,
			CAPS, ____, MS_L, MS_D, MS_R, ____ /*                                        */, ____, KP_4, KP_5, KP_6, PPLS, ____,
			____, ____, ____, ____, ____, ____ /*                                        */, ____, KP_1, KP_2, KP_3, PENT, ____,
			/* */ ____, INS_, LEFT, RGHT /*                                                    */, UP__, DOWN, DOT_, PENT,
			/*                         */ BTN1, BTN2 /*                            */, BTN4, BTN3,
			/*                               */ HOME /*                            */, PGUP,
			/*                   */ BSPC, DEL_, END_ /*                            */, PGDN, ENT_, KP_0, /**/
		),
		// 2 - Programming Layer
		kinx.Layer(
			____, ____, ____, ____, ____, ____, ____, ____, ____ /*    */, FN2_, ____, ____, ____, ____, ____, ____, ____, FN1_,
			____, ____, ____, ____, ____, ____ /*                                        */, ____, ____, ____, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                        */, ____, ____, ____, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                        */, ____, ____, ____, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                        */, ____, ____, ____, ____, ____, ____,
			/* */ ____, ____, ____, ____ /*                                                    */, ____, ____, ____, ____,
			/*                         */ ____, ____ /*                            */, ____, ____,
			/*                               */ ____ /*                            */, ____,
			/*                   */ ____, ____, ____ /*                            */, ____, ____, ____, /**/
		),
	})
}
