package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/matrix/kinx"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	____ = NO

	PLAY = MEDIA_PLAY_PAUSE
	PREV = MEDIA_PREV_TRACK
	NEXT = MEDIA_NEXT_TRACK

	ESC_ = ESC
)

// FN0 -> Toggle "Keypad" layer on key press
// FN1 -> Toggle "Programming" on and off on key up/down
// FN2 -> CPU Reset on key down

func initKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		// 0 - Default Layer
		kinx.Layer(
			FN3, F1, F2, F3, F4, F5, F6, F7, F8 /*  */, F9, F10, F11, F12, PSCR, SLCK, BRK, FN0, FN1,
			EQL, N1, N2, N3, N4, N5 /*                       */, N6, N7, N8, N9, N0, MINS,
			TAB, Q, W, E, R, T /*                            */, Y, U, I, O, P, BSLS,
			RCTL, A, S, D, F, G /*                           */, H, J, K, L, SCLN, QUOT,
			LSFT, Z, X, C, V, B /*                           */, N, M, COMM, DOT, SLSH, RSFT,
			/**/ GRV, INS, LEFT, RGHT /*                     */, UP, DOWN, LBRC, RBRC,
			/*                 */ ESC, LGUI /*         */, LALT, LCTL,
			/*                      */ HOME /*         */, PGUP,
			/*            */ BSPC, DEL, END /*         */, PGDN, ENT, SPC, /**/
		),
		// 1 - Keypad Layer
		kinx.Layer(
			ESC_, ____, ____, ____, PLAY, PREV, NEXT, ____, ____ /*  */, ____, ____, ____, ____, MUTE, VOLD, VOLU, FN0, FN1,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, NLCK, KP_EQUAL, KP_SLASH, KP_ASTERISK, ____,
			____, ____, ____, MS_U, ____, ____ /*                                      */, ____, KP_7, KP_8, KP_9, KP_MINUS, ____,
			CAPS, ____, MS_L, MS_D, MS_R, ____ /*                                      */, ____, KP_4, KP_5, KP_6, KP_PLUS, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, KP_1, KP_2, KP_3, KP_ENTER, ____,
			/* */ ____, INS, LEFT, RGHT /*                                                   */, UP, DOWN, KP_DOT, KP_ENTER,
			/*                         */ BTN1, BTN2 /*                          */, BTN4, BTN3,
			/*                               */ HOME /*                          */, PGUP,
			/*                    */ BSPC, DEL, END /*      i                    */, PGDN, ENT, KP_0, /**/
		),
		// 2 - Programming Layer
		kinx.Layer(
			____, ____, ____, ____, ____, ____, ____, ____, ____ /*   */, FN2, ____, ____, ____, ____, ____, ____, ____, FN1,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, ____, ____, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, ____, ____, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, ____, ____, ____, ____, ____,
			____, ____, ____, ____, ____, ____ /*                                      */, ____, ____, ____, ____, ____, ____,
			/* */ ____, ____, ____, ____ /*                                                  */, ____, ____, ____, ____,
			/*                         */ ____, ____ /*                          */, ____, ____,
			/*                               */ ____ /*                          */, ____,
			/*                   */ ____, ____, ____ /*                          */, ____, ____, ____, /**/
		),
	})
}
