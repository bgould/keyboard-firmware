package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

// TODO: refactor/fix interfaces
// FIXME: actually a "Keymap"
func KinTKeymap() keyboard.Keymap {
	return KinesisAdvantageLayer(
		ESC, F1, F2, F3, F4, F5, F6, F7, F8 /*  */, F9, F10, F11, F12, PSCR, SLCK, BRK, TRNS, TRNS,
		EQL, N1, N2, N3, N4, N5 /*                       */, N6, N7, N8, N9, N0, MINS,
		TAB, Q, W, E, R, T /*                            */, Y, U, I, O, P, BSLS,
		LCTL, A, S, D, F, G /*                           */, H, J, K, L, SCLN, QUOT,
		LSFT, Z, X, C, V, B /*                           */, N, M, COMM, DOT, SLSH, RSFT,
		/**/ GRV, INS, LEFT, RGHT /*                     */, UP, DOWN, LBRC, RBRC,
		/*           */ ESC, LGUI /*                */, RALT, RCTL,
		/*                */ HOME /*                */, PGUP,
		/*      */ BSPC, DEL, END /*                */, PGDN, ENT, SPC, /**/
	)
}

func KinesisAdvantageLayer(
	kC0, kD0, kE0, kC1, kD1, kE1, kC2, kD2, kE2 /* */, kC3, kD3, kE3, kC4, kD4, kE4, kC5, kE5, kD5,
	k00, k10, k20, k30, k40, k50 /*                          */, k60, k70, k80, k90, kA0, kB0,
	k01, k11, k21, k31, k41, k51 /*                          */, k61, k71, k81, k91, kA1, kB1,
	k02, k12, k22, k32, k42, k52 /*                          */, k62, k72, k82, k92, kA2, kB2,
	k03, k13, k23, k33, k43, k53 /*                          */, k63, k73, k83, k93, kA3, kB3,
	/**/ k14, k24, k34, k54 /*                                    */, k64, k84, k94, kA4,
	/*                    */ k56, k55 /*                */, k96, k85,
	/*                         */ k35 /*                */, k86,
	/*               */ k36, k46, k25 /*                */, k66, k75, k65 /**/ Keycode,
) keyboard.Keymap {
	return keyboard.Keymap([][]Keycode{
		/*        0x0  0x1  0x2  0x3  0x4  0x5  0x6  */
		/*********************************************/
		/*    */ {k00, k01, k02, k03, 0x0, 0x0, 0x0},
		/*    */ {k10, k11, k12, k13, k14, 0x0, 0x0},
		/*    */ {k20, k21, k22, k23, k24, k25, 0x0},
		/*    */ {k30, k31, k32, k33, k34, k35, k36},
		/*    */ {k40, k41, k42, k43, 0x0, 0x0, k46},
		/*    */ {k50, k51, k52, k53, k54, k55, k56},
		/*    */ {k60, k61, k62, k63, k64, k65, k66},
		/*    */ {k70, k71, k72, k73, 0x0, k75, 0x0},
		/*    */ {k80, k81, k82, k83, k84, k85, k86},
		/*    */ {k90, k91, k92, k93, k94, 0x0, k96},
		/*    */ {kA0, kA1, kA2, kA3, kA4, 0x0, 0x0},
		/*    */ {kB0, kB1, kB2, kB3, 0x0, 0x0, 0x0},
		/*    */ {kC0, kC1, kC2, kC3, kC4, kC5, 0x0},
		/*    */ {kD0, kD1, kD2, kD3, kD4, kD5, 0x0},
		/*    */ {kE0, kE1, kE2, kE3, kE4, kE5, 0x0},
		/*********************************************/
	})
}
