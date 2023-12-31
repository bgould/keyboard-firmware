package keyboard

import (
	"bytes"
	"testing"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func TestKeymapZeroFill(t *testing.T) {
	keymap := KinTKeymap()
	found := false
	for _, layer := range keymap {
		for _, row_ := range layer {
			for _, key := range row_ {
				if key != KC_NO {
					found = true
				}
			}
		}
	}
	if !found {
		t.Fail()
	}
	keymap.ZeroFill()
	for _, layer := range keymap {
		for _, row_ := range layer {
			for _, key := range row_ {
				if key != KC_NO {
					t.Fail()
				}
			}
		}
	}
}

func TestKeymapRoundTrip(t *testing.T) {
	keymap := KinTKeymap()

	var buf bytes.Buffer
	n, err := keymap.WriteTo(&buf)
	if err != nil {
		t.Error(err)
	}
	println("WriteTo result:", n, "bytes")

	copied := KinTKeymap()
	copied.ZeroFill()

	n2, err := copied.ReadFrom(&buf)
	if err != nil {
		t.Error(err)
	}
	println("ReadFrom result:", n2, "bytes")

	// ensure same number of bytes were written and read
	if n != n2 {
		t.Fail()
	}

	// compare keymaps key by key and ensure they are equivalent
	for i, layer := range keymap {
		for j, row_ := range layer {
			for k, key := range row_ {
				if copied[i][j][k] != key {
					t.Fail()
				}
			}
		}
	}

}

const (
	KC_PROG = QK_BOOTLOADER
)

func KinTKeymap() Keymap {
	return Keymap([]Layer{KinesisAdvantageLayer(
		KC_ESC_, KC_F1__, KC_F2__, KC_F3__, KC_F4__, KC_F5__, KC_F6__, KC_F7__, KC_F8__ /*  */, KC_F9__, KC_F10_, KC_F11_, KC_F12_, KC_PSCR, KC_SCRL, KC_BRK_, KC_FN0_, KC_PROG,
		KC_EQL_, KC_N1__, KC_N2__, KC_N3__, KC_N4__, KC_N5__ /*                                                        */, KC_N6__, KC_N7__, KC_N8__, KC_N9__, KC_N0__, KC_MINS,
		KC_TAB_, KC_Q___, KC_W___, KC_E___, KC_R___, KC_T___ /*                                                        */, KC_Y___, KC_U___, KC_I___, KC_O___, KC_P___, KC_BSLS,
		KC_LCTL, KC_A___, KC_S___, KC_D___, KC_F___, KC_G___ /*                                                        */, KC_H___, KC_J___, KC_K___, KC_L___, KC_SCLN, KC_QUOT,
		KC_LSFT, KC_Z___, KC_X___, KC_C___, KC_V___, KC_B___ /*                                                        */, KC_N___, KC_M___, KC_COMM, KC_DOT_, KC_SLSH, KC_RSFT,
		/*    */ KC_GRV_, KC_INS_, KC_LEFT, KC_RGHT /*                                                                          */, KC_UP__, KC_DOWN, KC_LBRC, KC_RBRC,
		/*                                        */ KC_ESC_, KC_LGUI /*                                      */, KC_LALT, KC_RCTL,
		/*                                                 */ KC_HOME /*                                      */, KC_PGUP,
		/*                               */ KC_BSPC, KC_DEL_, KC_END_ /*                                      */, KC_PGDN, KC_ENT_, KC_SPC_, /**/
	), KinesisAdvantageLayer(
		KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS /*  */, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS,
		KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS /*                                                        */, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS,
		KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS /*                                                        */, KC_TRNS, KC_TRNS, KC_MS_U, KC_TRNS, KC_TRNS, KC_TRNS,
		KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS /*                                                        */, KC_TRNS, KC_MS_L, KC_MS_D, KC_MS_R, KC_TRNS, KC_TRNS,
		KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS /*                                                        */, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS,
		/*    */ KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS /*                                                                          */, KC_TRNS, KC_TRNS, KC_TRNS, KC_TRNS,
		/*                                        */ KC_TRNS, KC_TRNS /*                                      */, KC_TRNS, KC_TRNS,
		/*                                                 */ KC_TRNS /*                                      */, KC_TRNS,
		/*                               */ KC_TRNS, KC_TRNS, KC_TRNS /*                                      */, KC_TRNS, KC_TRNS, KC_TRNS, /**/
	)})
}

func KinesisAdvantageLayer(
	kC0, kD0, kE0, kC1, kD1, kE1, kC2, kD2, kE2 /* */, kC3, kD3, kE3, kC4, kD4, kE4, kC5, kE5, kD5,
	k00, k10, k20, k30, k40, k50 /*                          */, k60, k70, k80, k90, kA0, kB0,
	k01, k11, k21, k31, k41, k51 /*                          */, k61, k71, k81, k91, kA1, kB1,
	k02, k12, k22, k32, k42, k52 /*                          */, k62, k72, k82, k92, kA2, kB2,
	k03, k13, k23, k33, k43, k53 /*                          */, k63, k73, k83, k93, kA3, kB3,
	/**/ k14, k24, k34, k54 /*                                        */, k64, k84, k94, kA4,
	/*                    */ k56, k55 /*                */, k96, k85,
	/*                         */ k35 /*                */, k86,
	/*               */ k36, k46, k25 /*                */, k66, k75, k65 /**/ Keycode,
) Layer {
	return Layer([][]Keycode{
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
