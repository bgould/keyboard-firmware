package modelm

import (
	"github.com/bgould/keyboard-firmware/keyboard"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func ANSI101DefaultLayer() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{ANSI101Layer(

		KC_ESC_, KC_F1__, KC_F2__, KC_F3__, KC_F4__, KC_F5__, KC_F6__, KC_F7__, KC_F8__, KC_F9__, KC_F10_, KC_F11_, KC_F12_, KC_PSCR, KC_SCRL, KC_BRK_,

		KC_GRV_, KC_N1__, KC_N2__, KC_N3__, KC_N4__, KC_N5__, KC_N6__, KC_N7__, KC_N8__, KC_N9__, KC_N0__, KC_MINS, KC_EQL_, KC_BSPC, KC_INS_, KC_HOME, KC_PGUP, KC_NUM_, KC_PSLS, KC_PAST, KC_PMNS,
		KC_TAB_, KC_Q___, KC_W___, KC_E___, KC_R___, KC_T___, KC_Y___, KC_U___, KC_I___, KC_O___, KC_P___, KC_LBRC, KC_RBRC, KC_BSLS, KC_DEL_, KC_END_, KC_PGDN, KC_P7__, KC_P8__, KC_P9__, KC_PPLS,
		KC_CAPS, KC_A___, KC_S___, KC_D___, KC_F___, KC_G___, KC_H___, KC_J___, KC_K___, KC_L___, KC_SCLN, KC_QUOT, KC_ENT_, KC_P4__, KC_P5__, KC_P6__,
		KC_LSFT, KC_Z___, KC_X___, KC_C___, KC_V___, KC_B___, KC_N___, KC_M___, KC_COMM, KC_DOT_, KC_SLSH, KC_RSFT, KC_UP__, KC_P1__, KC_P2__, KC_P3__, KC_PENT,
		KC_LCTL, KC_LALT, KC_SPC_, KC_RALT, KC_RCTL, KC_LEFT, KC_DOWN, KC_RGHT, KC_P0__, KC_PDOT,
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
