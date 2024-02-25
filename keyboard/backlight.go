package keyboard

import (
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type BacklightMode uint8

type BacklightLevel uint8

const (
	BacklightOff BacklightMode = iota
	BacklightOn
	BacklightBreathing
)

type BacklightDriver interface {
	Configure()
	SetBacklight(mode BacklightMode, level BacklightLevel)
	Task()
}

type Backlight struct {

	// BacklightDriver implementation to use; if nil, backlight is disabled.
	Driver BacklightDriver

	//
	NumLevels uint8

	// Default backlight mode to use when not explicitly set
	DefaultMode BacklightMode

	// Default backlight level to use when not explicitly set
	DefaultLevel BacklightLevel

	// Whether or not backlight "breathing" is enabled
	SupportsBreathing bool
}

func (kbd *Keyboard) SetBacklight(bl Backlight) {
	kbd.backlight = bl
	kbd.blState = backlightState{mode: bl.DefaultMode, level: bl.DefaultLevel}
	kbd.backlight.Driver.SetBacklight(kbd.blState.mode, kbd.blState.level)
}

func (kbd *Keyboard) BacklightEnabled() bool {
	return kbd.backlight.Driver != nil
}

func (kbd *Keyboard) processBacklight(key keycodes.Keycode, made bool) {
	if !key.IsBacklight() { // sanity check
		return
	}
	if !kbd.BacklightEnabled() {
		return
	}
	prev := kbd.blState
	switch key {
	case keycodes.QK_BACKLIGHT_ON:
		// println("backlight on", made)
		if made {
			kbd.blState.On()
		}
	case keycodes.QK_BACKLIGHT_OFF:
		// println("backlight off", made)
		if made {
			kbd.blState.Off()
		}
	case keycodes.QK_BACKLIGHT_TOGGLE:
		// println("backlight toggle", made)
		if made {
			kbd.blState.Toggle()
		}
	case keycodes.QK_BACKLIGHT_DOWN:
		println("backlight down", made)
		// TODO: add support
	case keycodes.QK_BACKLIGHT_UP:
		println("backlight up", made)
		// TODO: add support
	case keycodes.QK_BACKLIGHT_STEP:
		println("backlight step", made)
		// TODO: add support
	case keycodes.QK_BACKLIGHT_TOGGLE_BREATHING:
		// println("backlight toggle breathing", made)
		if made {
			kbd.blState.ToggleBreathing()
		}
	}
	kbd.backlight.Driver.SetBacklight(kbd.blState.mode, kbd.blState.level)
	changed := kbd.blState != prev
	if changed {
		// println("backlight changed", kbd.blState.mode, kbd.blState.level, kbd.blState.breathing)
		kbd.backlight.Driver.SetBacklight(kbd.blState.mode, kbd.blState.level)
	} else {
		// println("backlight not changed")
	}
}

type backlightState struct {
	mode      BacklightMode
	level     BacklightLevel
	breathing bool
}

func (st *backlightState) Toggle() {
	switch st.mode {
	case BacklightOff:
		if st.breathing {
			st.mode = BacklightBreathing
		} else {
			st.mode = BacklightOn
		}
	case BacklightOn:
		st.mode = BacklightOff
		st.breathing = false
	case BacklightBreathing:
		st.mode = BacklightOff
		st.breathing = true
	}
}

func (st *backlightState) Off() {
	st.mode = BacklightOff
}

func (st *backlightState) On() {
	st.mode = BacklightOn
}

func (st *backlightState) ToggleBreathing() {
	st.breathing = !st.breathing
	if st.mode == BacklightOff {
		return
	}
	if st.breathing {
		st.mode = BacklightBreathing
	} else {
		st.mode = BacklightOn
	}
}
