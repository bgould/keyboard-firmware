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

	state backlightState
}

type backlightState struct {
	mode       BacklightMode
	level      BacklightLevel
	steps      uint8
	breathing  bool
	breathStep bool
}

func (bl *Backlight) steps() uint8 {
	if bl.Steps == 0 {
		return 1
	}
	return bl.Steps
}

func (kbd *Keyboard) SetBacklight(bl Backlight) {
	kbd.backlight = bl
	kbd.backlight.Driver.SetBacklight(bl.DefaultMode, bl.DefaultLevel)
}

func (kbd *Keyboard) BacklightEnabled() bool {
	return kbd.backlight.Driver != nil
}

func (kbd *Keyboard) processBacklight(key keycodes.Keycode, made bool) {
	if !kbd.BacklightEnabled() {
		return
	}
	kbd.backlight.ProcessKey(key, made)
}

func (bl *Backlight) ProcessKey(key keycodes.Keycode, made bool) {
	if !key.IsBacklight() { // sanity check
		return
	}
	prev := bl.state
	state := &bl.state
	state.steps = bl.steps()
	state.breathStep = bl.IncludeBreathingInSteps
	switch key {
	case keycodes.QK_BACKLIGHT_ON:
		// println("backlight on", made)
		if made {
			state.On()
		}
	case keycodes.QK_BACKLIGHT_OFF:
		// println("backlight off", made)
		if made {
			state.Off()
		}
	case keycodes.QK_BACKLIGHT_TOGGLE:
		// println("backlight toggle", made)
		if made {
			state.Toggle()
		}
	case keycodes.QK_BACKLIGHT_DOWN:
		step := state.step()
		// println("backlight down", made, step)
		if made && step != 0 {
			state.Down()
		}
	case keycodes.QK_BACKLIGHT_UP:
		// println("backlight up", made, state.Step())
		if made {
			state.Up()
		}
	case keycodes.QK_BACKLIGHT_STEP:
		// println("backlight step", made, state.Step())
		if made {
			state.Next()
		}
	case keycodes.QK_BACKLIGHT_TOGGLE_BREATHING:
		// println("backlight toggle breathing", made)
		if made {
			state.ToggleBreathing()
		}
	}
	// kbd.backlight.Driver.SetBacklight(state.mode, state.level)
	changed := bl.state != prev
	if bl.Driver != nil && changed {
		bl.Driver.SetBacklight(state.mode, state.level)
	}
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
	step := st.step()
	// println("turning on with step at", step)
	if step == 0 {
		st.setStep(1)
	}
	st.mode = BacklightOn
}

func (st *backlightState) Up() {
	st.setStep(st.step() + 1)
}

func (st *backlightState) Down() {
	st.setStep(st.step() - 1)
}

func (st *backlightState) Next() {
	curr := st.step()
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
			st.setStep(0)
		}
	} else {
		st.setStep(curr + 1)
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
func (st *backlightState) setLevel(val uint8) {
	st.level = BacklightLevel(val)
}

// SetRawLevel sets the level to an scaled by number of steps; returns the
// raw level value between 0x00-0xFF
func (st *backlightState) setStep(step uint8) uint8 {
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

// step returns the current
func (st *backlightState) step() uint8 {
	if st.level == 0 {
		return 0
	} else if st.level == 0xFF {
		return st.steps
	} else {
		return uint8(st.level) / (0xFF / st.steps)
	}
}
