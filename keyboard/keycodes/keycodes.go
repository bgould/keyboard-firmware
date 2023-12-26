package keycodes

type Keycode uint16

// func (code Keycode) IsError() bool {
// 	return (ROLL_OVER <= (code) && (code) <= UNDEFINED)
// }

func (code Keycode) IsAny() bool {
	return (A <= (code) && (code) <= 0xFF)
}

func (code Keycode) IsKey() bool {
	return code.IsBasic() // (A <= (code) && (code) <= EXSEL)
}

// func (code Keycode) IsModifier() bool {
// 	return LCTL <= code && code <= RGUI
// }

// func (code Keycode) IsSpecial() bool {
// 	return ((0xA5 <= (code) && (code) <= 0xDF) || (0xE8 <= (code) && (code) <= 0xFF))
// }

// func (code Keycode) IsSystem() bool {
// 	return (PWR <= (code) && (code) <= WAKE)
// }

// func (code Keycode) IsConsumer() bool {
// 	return (MUTE <= (code) && (code) <= WFAV)
// }

func (code Keycode) IsFn() bool {
	return (FN0 <= (code) && (code) <= FN31)
}

// func (code Keycode) IsMouseKey() bool {
// 	return (MS_UP <= (code) && (code) <= MS_ACCEL2)
// }

// func (code Keycode) IsMouseKeyMove() bool {
// 	return (MS_UP <= (code) && (code) <= MS_RIGHT)
// }

// func (code Keycode) IsMouseKeyButton() bool {
// 	return (MS_BTN1 <= (code) && (code) <= MS_BTN5)
// }

// func (code Keycode) IsMouseKeyWheel() bool {
// 	return (MS_WH_UP <= (code) && (code) <= MS_WH_RIGHT)
// }

// func (code Keycode) IsMouseKeyAccel() bool {
// 	return (MS_ACCEL0 <= (code) && (code) <= MS_ACCEL2)
// }
