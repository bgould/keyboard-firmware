package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/matrix/kinx"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	____ = TRNS
)

func Keymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		kinx.Layer(
			ESC, F1, F2, F3, F4, F5, F6, F7, F8 /*  */, F9, F10, F11, F12, PSCR, SLCK, BRK, FN0, PROG,
			EQL, N1, N2, N3, N4, N5 /*                       */, N6, N7, N8, N9, N0, MINS,
			TAB, Q, W, E, R, T /*                            */, Y, U, I, O, P, BSLS,
			LCTL, A, S, D, F, G /*                           */, H, J, K, L, SCLN, QUOT,
			LSFT, Z, X, C, V, B /*                           */, N, M, COMM, DOT, SLSH, RSFT,
			/**/ GRV, INS, LEFT, RGHT /*                     */, UP, DOWN, LBRC, RBRC,
			/*                 */ ESC, LGUI /*         */, LALT, RCTL,
			/*                      */ HOME /*         */, PGUP,
			/*            */ BSPC, DEL, END /*         */, PGDN, ENT, SPC, /**/
		),
		kinx.Layer(
			____, ____, ____, ____, ____, ____, ____, ____, ____ /*  */, ____, ____, ____, ____, MUTE, VOLD, VOLU, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, ____, ____, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, ____, MS_U, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, MS_L, MS_D, MS_R, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, ____, ____, ____, ____, ____,
			/* */ ____, ____, ____, ____ /*                                                  */, ____, ____, ____, ____,
			/*                         */ ____, ____ /*                          */, ____, ____,
			/*                               */ ____ /*                          */, ____,
			/*                   */ ____, ____, ____ /*                          */, ____, ____, ____, /**/
		),
	})
}
