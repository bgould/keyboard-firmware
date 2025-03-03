//go:build !keycodes.alt

package keycodes

// Keycode modifiers & aliases
const (
	QK_LCTL Keycode = 0x0100
	QK_LSFT Keycode = 0x0200
	QK_LALT Keycode = 0x0400
	QK_LGUI Keycode = 0x0800
	QK_RCTL Keycode = 0x1100
	QK_RSFT Keycode = 0x1200
	QK_RALT Keycode = 0x1400
	QK_RGUI Keycode = 0x1800

	QK_RMODS_MIN Keycode = 0x1000
)

func LCTL(kc Keycode) Keycode { return (QK_LCTL | (kc)) }
func LSFT(kc Keycode) Keycode { return (QK_LSFT | (kc)) }
func LALT(kc Keycode) Keycode { return (QK_LALT | (kc)) }
func LGUI(kc Keycode) Keycode { return (QK_LGUI | (kc)) }
func LOPT(kc Keycode) Keycode { return LALT(kc) }
func LCMD(kc Keycode) Keycode { return LGUI(kc) }
func LWIN(kc Keycode) Keycode { return LGUI(kc) }
func RCTL(kc Keycode) Keycode { return (QK_RCTL | (kc)) }
func RSFT(kc Keycode) Keycode { return (QK_RSFT | (kc)) }
func RALT(kc Keycode) Keycode { return (QK_RALT | (kc)) }
func RGUI(kc Keycode) Keycode { return (QK_RGUI | (kc)) }
func ALGR(kc Keycode) Keycode { return RALT(kc) }
func ROPT(kc Keycode) Keycode { return RALT(kc) }
func RCMD(kc Keycode) Keycode { return RGUI(kc) }
func RWIN(kc Keycode) Keycode { return RGUI(kc) }

// Keycode Aliases
const (
	// FN0  = QK_KB_0
	// FN1  = QK_KB_1
	// FN2  = QK_KB_2
	// FN3  = QK_KB_3
	// FN4  = QK_KB_4
	// FN5  = QK_KB_5
	// FN6  = QK_KB_6
	// FN7  = QK_KB_7
	// FN8  = QK_KB_8
	// FN9  = QK_KB_9
	// FN10 = QK_KB_10
	// FN11 = QK_KB_11
	// FN12 = QK_KB_12
	// FN13 = QK_KB_13
	// FN14 = QK_KB_14
	// FN15 = QK_KB_15
	// FN16 = QK_KB_16
	// FN17 = QK_KB_17
	// FN18 = QK_KB_18
	// FN19 = QK_KB_19
	// FN20 = QK_KB_20
	// FN21 = QK_KB_21
	// FN22 = QK_KB_22
	// FN23 = QK_KB_23
	// FN24 = QK_KB_24
	// FN25 = QK_KB_25
	// FN26 = QK_KB_26
	// FN27 = QK_KB_27
	// FN28 = QK_KB_28
	// FN29 = QK_KB_29
	// FN30 = QK_KB_30
	// FN31 = QK_KB_31

	// PLAY = KC_MEDIA_PLAY_PAUSE
	// PREV = KC_MEDIA_PREV_TRACK
	// NEXT = KC_MEDIA_NEXT_TRACK

	// ESC_ = KC_ESC
	// NLCK = NUM
	// SLCK = SCRL

	// KPEQ = KP_EQUAL
	// KPSL = KP_SLASH
	// KPPL = KP_PLUS
	// KPMI = KP_MINUS
	// KPEN = KP_ENTER
	// KPAS = KP_ASTERISK

	KC_A___ = KC_A
	KC_B___ = KC_B
	KC_C___ = KC_C
	KC_D___ = KC_D
	KC_E___ = KC_E
	KC_F___ = KC_F
	KC_G___ = KC_G
	KC_H___ = KC_H
	KC_I___ = KC_I
	KC_J___ = KC_J
	KC_K___ = KC_K
	KC_L___ = KC_L
	KC_M___ = KC_M
	KC_N___ = KC_N
	KC_O___ = KC_O
	KC_P___ = KC_P
	KC_Q___ = KC_Q
	KC_R___ = KC_R
	KC_S___ = KC_S
	KC_T___ = KC_T
	KC_U___ = KC_U
	KC_V___ = KC_V
	KC_W___ = KC_W
	KC_X___ = KC_X
	KC_Y___ = KC_Y
	KC_Z___ = KC_Z
	KC_N1__ = KC_1
	KC_N2__ = KC_2
	KC_N3__ = KC_3
	KC_N4__ = KC_4
	KC_N5__ = KC_5
	KC_N6__ = KC_6
	KC_N7__ = KC_7
	KC_N8__ = KC_8
	KC_N9__ = KC_9
	KC_N0__ = KC_0
	KC_F1__ = KC_F1
	KC_F2__ = KC_F2
	KC_F3__ = KC_F3
	KC_F4__ = KC_F4
	KC_F5__ = KC_F5
	KC_F6__ = KC_F6
	KC_F7__ = KC_F7
	KC_F8__ = KC_F8
	KC_F9__ = KC_F9
	KC_F10_ = KC_F10
	KC_F11_ = KC_F11
	KC_F12_ = KC_F12
	KC_F13_ = KC_F13
	KC_F14_ = KC_F14
	KC_F15_ = KC_F15
	KC_F16_ = KC_F16
	KC_F17_ = KC_F17
	KC_F18_ = KC_F18
	KC_F19_ = KC_F19
	KC_F20_ = KC_F20
	KC_F21_ = KC_F21
	KC_F22_ = KC_F22
	KC_F23_ = KC_F23
	KC_F24_ = KC_F24
	KC_TAB_ = KC_TAB
	KC_DOT_ = KC_DOT
	KC_CUT_ = KC_CUT
	KC_OUT_ = KC_OUT
	KC_END_ = KC_END
	KC_UP__ = KC_UP

	KC_FN0  = QK_KB_0
	KC_FN1  = QK_KB_1
	KC_FN2  = QK_KB_2
	KC_FN3  = QK_KB_3
	KC_FN4  = QK_KB_4
	KC_FN5  = QK_KB_5
	KC_FN6  = QK_KB_6
	KC_FN7  = QK_KB_7
	KC_FN8  = QK_KB_8
	KC_FN9  = QK_KB_9
	KC_FN0_ = QK_KB_0
	KC_FN1_ = QK_KB_1
	KC_FN2_ = QK_KB_2
	KC_FN3_ = QK_KB_3
	KC_FN4_ = QK_KB_4
	KC_FN5_ = QK_KB_5
	KC_FN6_ = QK_KB_6
	KC_FN7_ = QK_KB_7
	KC_FN8_ = QK_KB_8
	KC_FN9_ = QK_KB_9
	KC_FN10 = QK_KB_10
	KC_FN11 = QK_KB_11
	KC_FN12 = QK_KB_12
	KC_FN13 = QK_KB_13
	KC_FN14 = QK_KB_14
	KC_FN15 = QK_KB_15
	KC_FN16 = QK_KB_16
	KC_FN17 = QK_KB_17
	KC_FN18 = QK_KB_18
	KC_FN19 = QK_KB_19
	KC_FN20 = QK_KB_20
	KC_FN21 = QK_KB_21
	KC_FN22 = QK_KB_22
	KC_FN23 = QK_KB_23
	KC_FN24 = QK_KB_24
	KC_FN25 = QK_KB_25
	KC_FN26 = QK_KB_26
	KC_FN27 = QK_KB_27
	KC_FN28 = QK_KB_28
	KC_FN29 = QK_KB_29
	KC_FN30 = QK_KB_30
	KC_FN31 = QK_KB_31
)

/*

	KC_A___ = KC_A
	KC_B___ = KC_B
	KC_C___ = KC_C
	KC_D___ = KC_D
	KC_E___ = KC_E
	KC_F___ = KC_F
	KC_G___ = KC_G
	KC_H___ = KC_H
	KC_I___ = KC_I
	KC_J___ = KC_J
	KC_K___ = KC_K
	KC_L___ = KC_L
	KC_M___ = KC_M
	KC_N___ = KC_N
	KC_O___ = KC_O
	KC_P___ = KC_P
	KC_Q___ = KC_Q
	KC_R___ = KC_R
	KC_S___ = KC_S
	KC_T___ = KC_T
	KC_U___ = KC_U
	KC_V___ = KC_V
	KC_W___ = KC_W
	KC_X___ = KC_X
	KC_Y___ = KC_Y
	KC_Z___ = KC_Z
	KC_N1__ = KC_N1
	KC_N2__ = KC_N2
	KC_N3__ = KC_N3
	KC_N4__ = KC_N4
	KC_N5__ = KC_N5
	KC_N6__ = KC_N6
	KC_N7__ = KC_N7
	KC_N8__ = KC_N8
	KC_N9__ = KC_N9
	KC_N0__ = KC_N0
	KC_F1__ = KC_F1
	KC_F2__ = KC_F2
	KC_F3__ = KC_F3
	KC_F4__ = KC_F4
	KC_F5__ = KC_F5
	KC_F6__ = KC_F6
	KC_F7__ = KC_F7
	KC_F8__ = KC_F8
	KC_F9__ = KC_F9
	KC_F10_ = KC_F10
	KC_F11_ = KC_F11
	KC_F12_ = KC_F12
	KC_F13_ = KC_F13
	KC_F14_ = KC_F14
	KC_F15_ = KC_F15
	KC_F16_ = KC_F16
	KC_F17_ = KC_F17
	KC_F18_ = KC_F18
	KC_F19_ = KC_F19
	KC_F20_ = KC_F20
	KC_F21_ = KC_F21
	KC_F22_ = KC_F22
	KC_F23_ = KC_F23
	KC_F24_ = KC_F24
	KC_TAB_ = KC_TAB
	KC_DOT_ = KC_DOT
	KC_CUT_ = KC_CUT
	KC_OUT__ = KC_OUT

	ENTER                                        Keycode = 0x0028
	ESCAPE                                       Keycode = 0x0029
	BACKSPACE                                    Keycode = 0x002A
	SPACE                                        Keycode = 0x002C
	MINUS                                        Keycode = 0x002D
	EQUAL                                        Keycode = 0x002E
	LEFT_BRACKET                                 Keycode = 0x002F
	RIGHT_BRACKET                                Keycode = 0x0030
	BACKSLASH                                    Keycode = 0x0031
	NONUS_HASH                                   Keycode = 0x0032
	SEMICOLON                                    Keycode = 0x0033
	QUOTE                                        Keycode = 0x0034
	GRAVE                                        Keycode = 0x0035
	COMMA                                        Keycode = 0x0036
	SLASH                                        Keycode = 0x0038
	CAPS_LOCK                                    Keycode = 0x0039

	PRINT_SCREEN                                 Keycode = 0x0046
	SCROLL_LOCK                                  Keycode = 0x0047
	PAUSE                                        Keycode = 0x0048
	INSERT                                       Keycode = 0x0049
	HOME                                         Keycode = 0x004A
	PAGE_UP                                      Keycode = 0x004B
	DELETE                                       Keycode = 0x004C
	KC_END__ = KC_END
	PAGE_DOWN                                    Keycode = 0x004E
	RIGHT                                        Keycode = 0x004F
	LEFT                                         Keycode = 0x0050
	DOWN                                         Keycode = 0x0051
	KC_UP__ = KC_UP
	NUM_LOCK                                     Keycode = 0x0053
	KP_SLASH                                     Keycode = 0x0054
	KP_ASTERISK                                  Keycode = 0x0055
	KP_MINUS                                     Keycode = 0x0056
	KP_PLUS                                      Keycode = 0x0057
	KP_ENTER                                     Keycode = 0x0058
	KP_1                                         Keycode = 0x0059
	KP_2                                         Keycode = 0x005A
	KP_3                                         Keycode = 0x005B
	KP_4                                         Keycode = 0x005C
	KP_5                                         Keycode = 0x005D
	KP_6                                         Keycode = 0x005E
	KP_7                                         Keycode = 0x005F
	KP_8                                         Keycode = 0x0060
	KP_9                                         Keycode = 0x0061
	KP_0                                         Keycode = 0x0062
	KP_DOT                                       Keycode = 0x0063
	NONUS_BACKSLASH                              Keycode = 0x0064
	APPLICATION                                  Keycode = 0x0065
	KB_POWER                                     Keycode = 0x0066
	KP_EQUAL                                     Keycode = 0x0067
EXECUTE                                      Keycode = 0x0074
	HELP                                         Keycode = 0x0075
	MENU                                         Keycode = 0x0076
	SELECT                                       Keycode = 0x0077
	STOP                                         Keycode = 0x0078
	AGAIN                                        Keycode = 0x0079
	UNDO                                         Keycode = 0x007A
	COPY                                         Keycode = 0x007C
	PASTE                                        Keycode = 0x007D
	FIND                                         Keycode = 0x007E
	KB_MUTE                                      Keycode = 0x007F
	MS_UP                                        Keycode = 0x00CD
	KB_VOLUME_UP                                 Keycode = 0x0080
	KB_VOLUME_DOWN                               Keycode = 0x0081
	LOCKING_CAPS_LOCK                            Keycode = 0x0082
	LOCKING_NUM_LOCK                             Keycode = 0x0083
	LOCKING_SCROLL_LOCK                          Keycode = 0x0084
	KP_COMMA                                     Keycode = 0x0085
	KP_EQUAL_AS400                               Keycode = 0x0086
	INTERNATIONAL_1                              Keycode = 0x0087
	INTERNATIONAL_2                              Keycode = 0x0088
	INTERNATIONAL_3                              Keycode = 0x0089
	INTERNATIONAL_4                              Keycode = 0x008A
	INTERNATIONAL_5                              Keycode = 0x008B
	INTERNATIONAL_6                              Keycode = 0x008C
	INTERNATIONAL_7                              Keycode = 0x008D
	INTERNATIONAL_8                              Keycode = 0x008E
	INTERNATIONAL_9                              Keycode = 0x008F
	LANGUAGE_1                                   Keycode = 0x0090
	LANGUAGE_2                                   Keycode = 0x0091
	LANGUAGE_3                                   Keycode = 0x0092
	LANGUAGE_4                                   Keycode = 0x0093
	LANGUAGE_5                                   Keycode = 0x0094
	LANGUAGE_6                                   Keycode = 0x0095
	LANGUAGE_7                                   Keycode = 0x0096
	LANGUAGE_8                                   Keycode = 0x0097
	LANGUAGE_9                                   Keycode = 0x0098
	ALTERNATE_ERASE                              Keycode = 0x0099
	SYSTEM_REQUEST                               Keycode = 0x009A
	CANCEL                                       Keycode = 0x009B
	CLEAR                                        Keycode = 0x009C
	PRIOR                                        Keycode = 0x009D
	RETURN                                       Keycode = 0x009E
	SEPARATOR                                    Keycode = 0x009F
	OPER                                         Keycode = 0x00A1
	CLEAR_AGAIN                                  Keycode = 0x00A2
	CRSEL                                        Keycode = 0x00A3
	EXSEL                                        Keycode = 0x00A4
	SYSTEM_POWER                                 Keycode = 0x00A5
	SYSTEM_SLEEP                                 Keycode = 0x00A6
	SYSTEM_WAKE                                  Keycode = 0x00A7
	AUDIO_MUTE                                   Keycode = 0x00A8
	AUDIO_VOL_UP                                 Keycode = 0x00A9
	AUDIO_VOL_DOWN                               Keycode = 0x00AA
	MEDIA_NEXT_TRACK                             Keycode = 0x00AB
	MEDIA_PREV_TRACK                             Keycode = 0x00AC
	MEDIA_STOP                                   Keycode = 0x00AD
	MEDIA_PLAY_PAUSE                             Keycode = 0x00AE
	MEDIA_SELECT                                 Keycode = 0x00AF
	MEDIA_EJECT                                  Keycode = 0x00B0
	MAIL                                         Keycode = 0x00B1
	CALCULATOR                                   Keycode = 0x00B2
	MY_COMPUTER                                  Keycode = 0x00B3
	WWW_SEARCH                                   Keycode = 0x00B4
	WWW_HOME                                     Keycode = 0x00B5
	WWW_BACK                                     Keycode = 0x00B6
	WWW_FORWARD                                  Keycode = 0x00B7
	WWW_STOP                                     Keycode = 0x00B8
	WWW_REFRESH                                  Keycode = 0x00B9
	WWW_FAVORITES                                Keycode = 0x00BA
	MEDIA_FAST_FORWARD                           Keycode = 0x00BB
	MEDIA_REWIND                                 Keycode = 0x00BC
	BRIGHTNESS_UP                                Keycode = 0x00BD
	BRIGHTNESS_DOWN                              Keycode = 0x00BE
	CONTROL_PANEL                                Keycode = 0x00BF
	ASSISTANT                                    Keycode = 0x00C0
	MISSION_CONTROL                              Keycode = 0x00C1
	LAUNCHPAD                                    Keycode = 0x00C2



*/

// #define IS_LAYER_KEYCODE(code) ((code) >= QK_TO && (code) <= QK_LAYER_TAP_TOGGLE_MAX)
func (code Keycode) IsLayer() bool {
	return ((code) >= QK_TO && (code) <= QK_LAYER_TAP_TOGGLE_MAX)
}

func MO_(layer uint8) Keycode {
	return QK_MOMENTARY + Keycode(layer&0x1F)
}

func TG_(layer uint8) Keycode {
	return QK_TOGGLE_LAYER + Keycode(layer&0x1F)
}
