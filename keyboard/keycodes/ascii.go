package keycodes

func AsciiToKeycode(ascii uint8) (keycode Keycode, shifted, altgred, dead bool) {
	ascii &= 0x7F
	keycode = ascii_to_keycode_lut[ascii]
	shifted = (ascii_to_shift_lut[ascii/8]>>(ascii%8))&1 == 1
	altgred = (ascii_to_altgr_lut[ascii/8]>>(ascii%8))&1 == 1
	dead = (ascii_to_dead_lut[ascii/8]>>(ascii%8))&1 == 1
	return

}

/* Look-up table to convert an ASCII character to a keycode.
 */
var ascii_to_keycode_lut [128]Keycode = [...]Keycode{
	// NUL   SOH      STX      ETX      EOT      ENQ      ACK      BEL
	XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX,
	// BS    TAB      LF       VT       FF       CR       SO       SI
	KC_BSPC, KC_TAB_, KC_ENT_, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX,
	// DLE   DC1      DC2      DC3      DC4      NAK      SYN      ETB
	XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX,
	// CAN   EM       SUB      ESC      FS       GS       RS       US
	XXXXXXX, XXXXXXX, XXXXXXX, KC_ESC_, XXXXXXX, XXXXXXX, XXXXXXX, XXXXXXX,

	//       !        "        #        $        %        &        '
	KC_SPC_, KC_N1__, KC_QUOT, KC_N3__, KC_N4__, KC_N5__, KC_N7__, KC_QUOT,
	// (     )        *        +        ,        -        .        /
	KC_N9__, KC_N0__, KC_N8__, KC_EQL_, KC_COMM, KC_MINS, KC_DOT_, KC_SLSH,
	// 0     1        2        3        4        5        6        7
	KC_N0__, KC_N1__, KC_N2__, KC_N3__, KC_N4__, KC_N5__, KC_N6__, KC_N7__,
	// 8     9        :        ;        <        =        >        ?
	KC_N8__, KC_N9__, KC_SCLN, KC_SCLN, KC_COMM, KC_EQL_, KC_DOT, KC_SLSH,
	// @     A        B        C        D        E        F        G
	KC_N2__, KC_A___, KC_B___, KC_C___, KC_D___, KC_E___, KC_F___, KC_G___,
	// H     I        J        K        L        M        N        O
	KC_H___, KC_I___, KC_J___, KC_K___, KC_L___, KC_M___, KC_N___, KC_O___,
	// P     Q        R        S        T        U        V        W
	KC_P___, KC_Q___, KC_R___, KC_S___, KC_T___, KC_U___, KC_V___, KC_W___,
	// X     Y        Z        [        \        ]        ^        _
	KC_X___, KC_Y___, KC_Z___, KC_LBRC, KC_BSLS, KC_RBRC, KC_N6__, KC_MINS,
	// `     a        b        c        d        e        f        g
	KC_GRV_, KC_A___, KC_B___, KC_C___, KC_D___, KC_E___, KC_F___, KC_G___,
	// h     i        j        k        l        m        n        o
	KC_H___, KC_I___, KC_J___, KC_K___, KC_L___, KC_M___, KC_N___, KC_O___,
	// p     q        r        s        t        u        v        w
	KC_P___, KC_Q___, KC_R___, KC_S___, KC_T___, KC_U___, KC_V___, KC_W___,
	// x     y        z        {        |        }        ~        DEL
	KC_X___, KC_Y___, KC_Z___, KC_LBRC, KC_BSLS, KC_RBRC, KC_GRV_, KC_DEL_,
}

// Bit-Packed look-up table to convert an ASCII character to whether
// [Shift] needs to be sent with the keycode.
var ascii_to_shift_lut [16]uint8 = [...]uint8{
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b01111110,
	0b11110000,
	0b00000000,
	0b00101011,
	0b11111111,
	0b11111111,
	0b11111111,
	0b11100011,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00011110,
}

// Bit-Packed look-up table to convert an ASCII character to whether
// [AltGr] needs to be sent with the keycode.
var ascii_to_altgr_lut [16]uint8 = [...]uint8{
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
}

// Bit-Packed look-up table to convert an ASCII character to whether
// [Space] needs to be sent after the keycode
var ascii_to_dead_lut [16]uint8 = [...]uint8{
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
}
