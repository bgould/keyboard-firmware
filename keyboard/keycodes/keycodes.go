package keycodes 

/*
 * Keycodes based on HID Usage Keyboard/Keypad Page(0x07) plus special codes
 * http://www.usb.org/developers/devclass_docs/Hut1_12.pdf
 */
/*

#define MOD_BIT(code)   (1<<MOD_INDEX(code))
#define MOD_INDEX(code) ((code) & 0x07)
#define FN_BIT(code)    (1<<FN_INDEX(code))
#define FN_INDEX(code)  ((code) - FN0)
#define FN_MIN          FN0
#define FN_MAX          FN31

*/

type Keycode uint8

func (code Keycode) IsError() bool {
	return (ROLL_OVER <= (code) && (code) <= UNDEFINED)
}

func (code Keycode) IsAny() bool {
	return (A <= (code) && (code) <= 0xFF)
}

func (code Keycode) IsKey() bool {
	return (A <= (code) && (code) <= EXSEL)
}

func (code Keycode) IsModifier() bool {
	return LCTRL <= code && code <= RGUI
}

func (code Keycode) IsSpecial() bool {
	return ((0xA5 <= (code) && (code) <= 0xDF) || (0xE8 <= (code) && (code) <= 0xFF))
}

func (code Keycode) IsSystem() bool {
	return (PWR <= (code) && (code) <= WAKE)
}

func (code Keycode) IsConsumer() bool {
	return (MUTE <= (code) && (code) <= WFAV)
}

func (code Keycode) IsFn() bool {
	return (FN0 <= (code) && (code) <= FN31)
}

func (code Keycode) IsMouseKey() bool {
	return (MS_UP <= (code) && (code) <= MS_ACCEL2)
}

func (code Keycode) IsMouseKeyMove() bool {
	return (MS_UP <= (code) && (code) <= MS_RIGHT)
}

func (code Keycode) IsMouseKeyButton() bool {
	return (MS_BTN1 <= (code) && (code) <= MS_BTN5)
}

func (code Keycode) IsMouseKeyWheel() bool {
	return (MS_WH_UP <= (code) && (code) <= MS_WH_RIGHT)
}

func (code Keycode) IsMouseKeyAccel() bool {
	return (MS_ACCEL0 <= (code) && (code) <= MS_ACCEL2)
}

/* USB HID Keyboard/Keypad Usage(0x07) */
const (
	NO Keycode = iota
	ROLL_OVER
	POST_FAIL
	UNDEFINED
	A /* 0x04 */
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M /* 0x10 */
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
	N1
	N2
	N3 /* 0x20 */
	N4
	N5
	N6
	N7
	N8
	N9
	N0
	ENTER
	ESCAPE
	BSPACE
	TAB
	SPACE
	MINUS
	EQUAL
	LBRACKET
	RBRACKET   /* 0x30 */
	BSLASH     /* \ (and |) */
	NONUS_HASH /* Non-US # and ~ (Typically near the Enter key) */
	SCOLON     /* ; (and :) */
	QUOTE      /* ' and " */
	GRAVE      /* Grave accent and tilde */
	COMMA      /* , and < */
	DOT        /* . and > */
	SLASH      /* / and ? */
	CAPSLOCK
	F1
	F2
	F3
	F4
	F5
	F6
	F7 /* 0x40 */
	F8
	F9
	F10
	F11
	F12
	PSCREEN
	SCROLLLOCK
	PAUSE
	INSERT
	HOME
	PGUP
	DELETE
	END
	PGDOWN
	RIGHT
	LEFT /* 0x50 */
	DOWN
	UP
	NUMLOCK
	KP_SLASH
	KP_ASTERISK
	KP_MINUS
	KP_PLUS
	KP_ENTER
	KP_1
	KP_2
	KP_3
	KP_4
	KP_5
	KP_6
	KP_7
	KP_8 /* 0x60 */
	KP_9
	KP_0
	KP_DOT
	NONUS_BSLASH /* Non-US \ and | (Typically near the Left-Shift key) */
	APPLICATION
	POWER
	KP_EQUAL
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21 /* 0x70 */
	F22
	F23
	F24
	EXECUTE
	HELP
	MENU
	SELECT
	STOP
	AGAIN
	UNDO
	CUT
	COPY
	PASTE
	FIND
	_MUTE
	_VOLUP /* 0x80 */
	_VOLDOWN
	LOCKING_CAPS   /* locking Caps Lock */
	LOCKING_NUM    /* locking Num Lock */
	LOCKING_SCROLL /* locking Scroll Lock */
	KP_COMMA
	KP_EQUAL_AS400 /* equal sign on AS/400 */
	INT1
	INT2
	INT3
	INT4
	INT5
	INT6
	INT7
	INT8
	INT9
	LANG1 /* 0x90 */
	LANG2
	LANG3
	LANG4
	LANG5
	LANG6
	LANG7
	LANG8
	LANG9
	ALT_ERASE
	SYSREQ
	CANCEL
	CLEAR
	PRIOR
	RETURN
	SEPARATOR
	OUT /* 0xA0 */
	OPER
	CLEAR_AGAIN
	CRSEL
	EXSEL /* 0xA4 */
)

/* NOTE: Following code range(0xB0-DD) are shared with special codes of 8-bit keymap */
const (
	KP_00 = iota + 0xB0
	KP_000
	THOUSANDS_SEPARATOR
	DECIMAL_SEPARATOR
	CURRENCY_UNIT
	CURRENCY_SUB_UNIT
	KP_LPAREN
	KP_RPAREN
	KP_LCBRACKET /* { */
	KP_RCBRACKET /* } */
	KP_TAB
	KP_BSPACE
	KP_A
	KP_B
	KP_C
	KP_D
	KP_E /* 0xC0 */
	KP_F
	KP_XOR
	KP_HAT
	KP_PERC
	KP_LT
	KP_GT
	KP_AND
	KP_LAZYAND
	KP_OR
	KP_LAZYOR
	KP_COLON
	KP_HASH
	KP_SPACE
	KP_ATMARK
	KP_EXCLAMATION
	KP_MEM_STORE /* 0xD0 */
	KP_MEM_RECALL
	KP_MEM_CLEAR
	KP_MEM_ADD
	KP_MEM_SUB
	KP_MEM_MUL
	KP_MEM_DIV
	KP_PLUS_MINUS
	KP_CLEAR
	KP_CLEAR_ENTRY
	KP_BINARY
	KP_OCTAL
	KP_DECIMAL
	KP_HEXADECIMAL /* 0xDD */
)

/* Modifiers */
const (
	LCTRL = iota + 0xE0
	LSHIFT
	LALT
	LGUI
	RCTRL
	RSHIFT
	RALT
	RGUI /* 0xE7 */
)

/* Special keycodes for 8-bit keymap
   NOTE: 0xA5-DF and 0xE8-FF are used for internal special purpose */
const (
	/* System Control */
	SYSTEM_POWER = iota + 0xA5
	SYSTEM_SLEEP
	SYSTEM_WAKE

	/* Media Control */
	AUDIO_MUTE
	AUDIO_VOL_UP
	AUDIO_VOL_DOWN
	MEDIA_NEXT_TRACK
	MEDIA_PREV_TRACK
	MEDIA_FAST_FORWARD
	MEDIA_REWIND
	MEDIA_STOP
	MEDIA_PLAY_PAUSE
	MEDIA_EJECT
	MEDIA_SELECT
	MAIL
	CALCULATOR
	MY_COMPUTER
	WWW_SEARCH
	WWW_HOME
	WWW_BACK
	WWW_FORWARD
	WWW_STOP
	WWW_REFRESH
	WWW_FAVORITES /* 0xBC */
)

/* Jump to bootloader */
const BOOTLOADER = 0xBF

/* Fn key */
const (
	FN0 = iota + 0xC0
	FN1
	FN2
	FN3
	FN4
	FN5
	FN6
	FN7
	FN8
	FN9
	FN10
	FN11
	FN12
	FN13
	FN14
	FN15

	FN16 = iota + 0xD0
	FN17
	FN18
	FN19
	FN20
	FN21
	FN22
	FN23
	FN24
	FN25
	FN26
	FN27
	FN28
	FN29
	FN30
	FN31 /* 0xDF */
)

/**************************************/
/* 0xE0-E7 for Modifiers. DO NOT USE. */
/**************************************/

/* Mousekey */
const (
	MS_UP = iota + 0xF0
	MS_DOWN
	MS_LEFT
	MS_RIGHT
	MS_BTN1
	MS_BTN2
	MS_BTN3
	MS_BTN4
	MS_BTN5 /* 0xF8 */
	/* Mousekey wheel */
	MS_WH_UP
	MS_WH_DOWN
	MS_WH_LEFT
	MS_WH_RIGHT /* 0xFC */
	/* Mousekey accel */
	MS_ACCEL0
	MS_ACCEL1
	MS_ACCEL2 /* 0xFF */
)

/*
 * Short names for ease of definition of keymap
 */
const (
	LCTL = LCTRL
	RCTL = RCTRL
	LSFT = LSHIFT
	RSFT = RSHIFT
	ESC  = ESCAPE
	BSPC = BSPACE
	ENT  = ENTER
	DEL  = DELETE
	INS  = INSERT
	CAPS = CAPSLOCK
	CLCK = CAPSLOCK
	RGHT = RIGHT
	PGDN = PGDOWN
	PSCR = PSCREEN
	SLCK = SCROLLLOCK
	PAUS = PAUSE
	BRK  = PAUSE
	NLCK = NUMLOCK
	SPC  = SPACE
	MINS = MINUS
	EQL  = EQUAL
	GRV  = GRAVE
	RBRC = RBRACKET
	LBRC = LBRACKET
	COMM = COMMA
	BSLS = BSLASH
	SLSH = SLASH
	SCLN = SCOLON
	QUOT = QUOTE
	APP  = APPLICATION
	NUHS = NONUS_HASH
	NUBS = NONUS_BSLASH
	LCAP = LOCKING_CAPS
	LNUM = LOCKING_NUM
	LSCR = LOCKING_SCROLL
	ERAS = ALT_ERASE
	CLR  = CLEAR

	/* Japanese specific */
	ZKHK = GRAVE
	RO   = INT1
	KANA = INT2
	JYEN = INT3
	JPY  = INT3
	HENK = INT4
	MHEN = INT5

	/* Korean specific */
	HAEN = LANG1
	HANJ = LANG2

	/* Keypad */
	P1   = KP_1
	P2   = KP_2
	P3   = KP_3
	P4   = KP_4
	P5   = KP_5
	P6   = KP_6
	P7   = KP_7
	P8   = KP_8
	P9   = KP_9
	P0   = KP_0
	PDOT = KP_DOT
	PCMM = KP_COMMA
	PSLS = KP_SLASH
	PAST = KP_ASTERISK
	PMNS = KP_MINUS
	PPLS = KP_PLUS
	PEQL = KP_EQUAL
	PENT = KP_ENTER

	/* Unix function key */
	EXEC = EXECUTE
	SLCT = SELECT
	AGIN = AGAIN
	PSTE = PASTE

	/* Mousekey */
	MS_U = MS_UP
	MS_D = MS_DOWN
	MS_L = MS_LEFT
	MS_R = MS_RIGHT
	BTN1 = MS_BTN1
	BTN2 = MS_BTN2
	BTN3 = MS_BTN3
	BTN4 = MS_BTN4
	BTN5 = MS_BTN5
	WH_U = MS_WH_UP
	WH_D = MS_WH_DOWN
	WH_L = MS_WH_LEFT
	WH_R = MS_WH_RIGHT
	ACL0 = MS_ACCEL0
	ACL1 = MS_ACCEL1
	ACL2 = MS_ACCEL2

	/* Sytem Control */
	PWR  = SYSTEM_POWER
	SLEP = SYSTEM_SLEEP
	WAKE = SYSTEM_WAKE

	/* Consumer Page */
	MUTE = AUDIO_MUTE
	VOLU = AUDIO_VOL_UP
	VOLD = AUDIO_VOL_DOWN
	MNXT = MEDIA_NEXT_TRACK
	MPRV = MEDIA_PREV_TRACK
	MFFD = MEDIA_FAST_FORWARD
	MRWD = MEDIA_REWIND
	MSTP = MEDIA_STOP
	MPLY = MEDIA_PLAY_PAUSE
	EJCT = MEDIA_EJECT
	MSEL = MEDIA_SELECT
	CALC = CALCULATOR
	MYCM = MY_COMPUTER
	WSCH = WWW_SEARCH
	WHOM = WWW_HOME
	WBAK = WWW_BACK
	WFWD = WWW_FORWARD
	WSTP = WWW_STOP
	WREF = WWW_REFRESH
	WFAV = WWW_FAVORITES

	/* Jump to bootloader */
	BTLD = BOOTLOADER

	/* Transparent */
	TRANSPARENT = 1
	TRNS        = TRANSPARENT
)
