package modelm

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func ANSI101DefaultLayer() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{ANSI101Layer(

		ESC, F1, F2, F3, F4, F5, F6, F7, F8, F9, F10, F11, F12, PSCR, SLCK, BRK,

		GRV, N1, N2, N3, N4, N5, N6, N7, N8, N9, N0, MINS, EQL, BSPC, INS, HOME, PGUP, NLCK, PSLS, PAST, PMNS,
		TAB, Q, W, E, R, T, Y, U, I, O, P, LBRC, RBRC, BSLS, DEL, END, PGDN, P7, P8, P9, PPLS,
		CAPS, A, S, D, F, G, H, J, K, L, SCLN, QUOT, ENT, P4, P5, P6,
		LSFT, Z, X, C, V, B, N, M, COMM, DOT, SLSH, RSFT, UP, P1, P2, P3, PENT,
		LCTL, LALT, SPC, RALT, RCTL, LEFT, DOWN, RGHT, P0, PDOT,
	)})
}

func ANSI101Layer(

	K72, K53, K54, K64, K74, K76, K78, K69, K59, K56, K46, K4B, K4C, K4F, K3F, K1E,

	K52, K42, K43, K44, K45, K55, K57, K47, K48, K49, K4A, K5A, K58, K66, K5C, K5E, K5D, K1B, K1C, K1D, K0D,
	K62, K32, K33, K34, K35, K65, K67, K37, K38, K39, K3A, K6A, K68, K26, K5B, K4E, K4D, K3B, K3C, K3D, K3E,
	K63, K22, K23, K24, K25, K75, K77, K27, K28, K29, K2A, K7A, K16, K6B, K6C, K6D,
	K61, K12, K13, K14, K15, K05, K07, K17, K18, K19, K0A, K11, K7E, K2B, K2C, K2D, K2E,
	K50, K7F, K06, K0F, K10, K0E, K0B, K0C, K7C, K7D Keycode,

) keyboard.Layer {
	return keyboard.Layer([][]Keycode{
		/*       0x0  0x1  0x2  0x3  0x4  0x5  0x6  0x7  0x8  0x9  0xA  0xB  0xC  0xD  0xE  0xF */
		/****************************************************************************************/
		/* 0 */ {0x0, 0x0, 0x0, 0x0, 0x0, K05, K06, K07, 0x0, 0x0, K0A, K0B, K0C, K0D, K0E, K0F},
		/* 1 */ {K10, K11, K12, K13, K14, K15, K16, K17, K18, K19, 0x0, K1B, K1C, K1D, K1E, 0x0},
		/* 2 */ {0x0, 0x0, K22, K23, K24, K25, K26, K27, K28, K29, K2A, K2B, K2C, K2D, K2E, 0x0},
		/* 3 */ {0x0, 0x0, K32, K33, K34, K35, 0x0, K37, K38, K39, K3A, K3B, K3C, K3D, K3E, K3F},
		/* 4 */ {0x0, 0x0, K42, K43, K44, K45, K46, K47, K48, K49, K4A, K4B, K4C, K4D, K4E, K4F},
		/* 5 */ {K50, 0x0, K52, K53, K54, K55, K56, K57, K58, K59, K5A, K5B, K5C, K5D, K5E, 0x0},
		/* 6 */ {0x0, K61, K62, K63, K64, K65, K66, K67, K68, K69, K6A, K6B, K6C, K6D, 0x0, 0x0},
		/* 7 */ {0x0, 0x0, K72, 0x0, K74, K75, K76, K77, K78, 0x0, K7A, 0x0, K7C, K7D, K7E, K7F},
		/****************************************************************************************/
	})
}
