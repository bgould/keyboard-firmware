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

	// Steps is the number of "on" states in between the off and on states.
	// The default zero value is treated the same as a value of 1.
	// A value 1 means that only on/off are supported; any value higher than that
	// will be used to automatically calculate equal "steps" in between the off
	// and fully on state.
	Steps uint8

	// Default backlight mode to use when not explicitly set
	DefaultMode BacklightMode

	// Default backlight level to use when not explicitly set
	DefaultLevel BacklightLevel

	// Whether or not backlight "breathing" is enabled
	SupportsBreathing bool

	// Whether or not to include "breathing" mode when iterating backlight steps
	IncludeBreathingInSteps bool
}

func (bl *Backlight) steps() uint8 {
	if bl.Steps == 0 {
		return 1
	}
	return bl.Steps
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
	kbd.blState.steps = kbd.backlight.steps()
	kbd.blState.breathStep = kbd.backlight.IncludeBreathingInSteps
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
		step := kbd.blState.Step()
		// println("backlight down", made, step)
		if made && step != 0 {
			kbd.blState.Down()
		}
	case keycodes.QK_BACKLIGHT_UP:
		// println("backlight up", made, kbd.blState.Step())
		if made {
			kbd.blState.Up()
		}
	case keycodes.QK_BACKLIGHT_STEP:
		// println("backlight step", made, kbd.blState.Step())
		if made {
			kbd.blState.Next()
		}
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
	mode       BacklightMode
	level      BacklightLevel
	steps      uint8
	breathing  bool
	breathStep bool
}

func (st *backlightState) Toggle() {
	switch st.mode {
	case BacklightOff:
		if st.breathing {
			st.mode = BacklightBreathing
		} else {
			st.On()
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
	step := st.Step()
	println("turning on with step at", step)
	if step == 0 {
		st.SetStep(1)
	}
	st.mode = BacklightOn
}

func (st *backlightState) Up() {
	st.SetStep(st.Step() + 1)
}

func (st *backlightState) Down() {
	st.SetStep(st.Step() - 1)
}

func (st *backlightState) Next() {
	curr := st.Step()
	mode := st.mode
	if st.mode == BacklightBreathing {
		st.mode = BacklightOn
		st.breathing = false
	}
	if curr >= (st.steps) {
		if st.breathStep && mode != BacklightBreathing {
			st.mode = BacklightBreathing
			st.breathing = true
		} else {
			st.SetStep(0)
		}
	} else {
		st.SetStep(curr + 1)
	}
}

func (st *backlightState) ToggleBreathing() {
	st.breathing = !st.breathing
	if st.mode == BacklightOff {
		st.breathing = true
		// return
	}
	if st.breathing {
		st.mode = BacklightBreathing
	} else {
		st.mode = BacklightOn
	}
}

// SetRawLevel sets the level to an unscaled by number of steps
func (st *backlightState) SetLevel(val uint8) {
	st.level = BacklightLevel(val)
}

// SetRawLevel sets the level to an scaled by number of steps; returns the
// raw level value between 0x00-0xFF
func (st *backlightState) SetStep(step uint8) uint8 {
	if step == 0 {
		st.level = 0
		st.mode = BacklightOff
	} else if step < (st.steps) {
		// println(st.steps)
		var intvl uint8 = 0xFF / st.steps
		st.level = BacklightLevel(step * intvl)
		st.breathing = false
		st.mode = BacklightOn
	} else {
		st.level = 0xFF
		st.breathing = false
		st.mode = BacklightOn
	}
	return uint8(st.level)
}

// Step returns the current
func (st *backlightState) Step() uint8 {
	if st.level == 0 {
		return 0
	} else if st.level == 0xFF {
		return st.steps
	} else {
		return uint8(st.level) / (0xFF / st.steps)
	}
}
