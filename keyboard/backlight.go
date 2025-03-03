package keyboard

import (
	"io"

	"github.com/bgould/keyboard-firmware/keyboard/hsv"
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
	SetBacklight(mode BacklightMode, color hsv.Color)
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

	sync bool
}

func (bl *Backlight) Configure() {
	if bl.Driver != nil {
		bl.Driver.Configure()
	}
	bl.state.color = hsv.Color{H: 128, S: 128, V: 255}
}

type backlightState struct {
	mode       BacklightMode
	color      hsv.Color
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
	kbd.backlight.Driver.SetBacklight(bl.DefaultMode, hsv.White)
}

func (kbd *Keyboard) BacklightDriver() BacklightDriver {
	return kbd.backlight.Driver
}

func (kbd *Keyboard) BacklightEnabled() bool {
	return kbd.backlight.Driver != nil
}

func (kbd *Keyboard) BacklightMode() BacklightMode {
	return kbd.backlight.state.mode
}

func (kbd *Keyboard) BacklightColor() hsv.Color {
	return kbd.backlight.state.color
}

func (kbd *Keyboard) BacklightUpdate(mode BacklightMode, color hsv.Color, force bool) {
	// println("entering BacklightUpdate", mode, color.H, color.S, color.V, force)
	prev := kbd.backlight.state
	state := &kbd.backlight.state
	state.mode = mode
	state.color = color
	changed := (kbd.backlight.state != prev)
	if changed || force {
		kbd.backlight.sync = true
		// kbd.backlight.Sync()
	}
}

func (kbd *Keyboard) BacklightSave() {
	println("Backlight save")
	kbd.backlightSave = true
	// cmdinfo := console.CommandInfo{Stdout: kbd.CLI(), Cmd: "save-backlight"}
	// kbd.saveBacklight(cmdinfo)
	// kbd.CLI().WriteString("BacklightSave: not implemented")
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
	changed := (bl.state != prev)
	if changed {
		bl.Sync()
	}
}

func (bl *Backlight) Task() {
	if bl.sync {
		bl.Sync()
		bl.sync = false
	}
	if bl.Driver != nil {
		bl.Driver.Task()
	}
}

func (bl *Backlight) Sync() {
	if bl.Driver != nil {
		// println("backlight HSV", bl.state.color.H, bl.state.color.S, bl.state.color.V)
		bl.Driver.SetBacklight(bl.state.mode, bl.state.color)
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
// func (st *backlightState) setLevel(val uint8) {
// 	st.level = BacklightLevel(val)
// }

// SetRawLevel sets the level to an scaled by number of steps; returns the
// raw level value between 0x00-0xFF
func (st *backlightState) setStep(step uint8) uint8 {
	if step == 0 {
		// st.level = 0
		st.color.V = 0
		st.mode = BacklightOff
	} else if step < (st.steps) {
		// println(st.steps)
		var intvl uint8 = 0xFF / st.steps
		// st.level = BacklightLevel(step * intvl)
		st.color.V = step * intvl
		st.breathing = false
		st.mode = BacklightOn
	} else {
		// st.level = 0xFF
		st.color.V = 0xFF
		st.breathing = false
		st.mode = BacklightOn
	}
	return st.color.V // uint8(st.level)
}

// step returns the current
func (st *backlightState) step() uint8 {
	if st.color.V == 0 {
		return 0
	} else if st.color.V == 0xFF {
		return st.steps
	} else {
		return uint8(st.color.V) / (0xFF / st.steps)
	}
}

func (m *backlightState) StoredSize() int {
	bufSize := 4
	hdrSize := 4
	ftrSize := 4
	return hdrSize + bufSize + ftrSize
}

// func (m *backlightState) ZeroFill() {
// 	for iCol := range m.buffer {
// 		m.buffer[iCol] = 0x0
// 	}
// }

var _ io.WriterTo = (*backlightState)(nil)

func (m *backlightState) WriteTo(w io.Writer) (n int64, err error) {
	var result int
	// header
	result, err = w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	n += int64(result)
	if err != nil {
		return
	}
	var buffer [4]byte
	buffer[0] = byte(m.mode)
	buffer[1] = m.color.H
	buffer[2] = m.color.S
	buffer[3] = m.color.V
	result, err = w.Write(buffer[:])
	n += int64(result)
	if err != nil {
		return
	}
	// footer
	result, err = w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	n += int64(result)
	return
}

var _ io.ReaderFrom = (*backlightState)(nil)

func (m *backlightState) ReadFrom(r io.Reader) (n int64, err error) {
	var result int
	var buffer [12]byte
	header := buffer[0:4]
	middle := buffer[4:8]
	footer := buffer[8:12]
	// header
	result, err = r.Read(header)
	n += int64(result)
	if err != nil {
		return
	}
	// buffer
	result, err = r.Read(middle)
	m.mode = BacklightMode(middle[0])
	m.color = hsv.Color{H: middle[1], S: middle[2], V: middle[3]}
	n += int64(result)
	if err != nil {
		return
	}
	// footer
	result, err = r.Read(footer)
	n += int64(result)
	if err != nil {
		return
	}
	return
}
