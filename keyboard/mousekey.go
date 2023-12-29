package keyboard

import (
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type MouseKeysConfig struct {
	MoveDelta     uint16
	MoveMaxSpeed  uint16
	MoveTimeToMax uint16

	WheelDelta     uint16
	WheelMaxSpeed  uint16
	WheelTimeToMax uint16

	Delay    time.Duration
	Interval time.Duration
}

// https://github.com/bgould/costar_tmk_keyboard/blob/master/config.h#L42-L49
func DefaultMouseKeysConfig() MouseKeysConfig {
	return MouseKeysConfig{
		Delay:          0,
		Interval:       20 * time.Millisecond,
		MoveDelta:      3,
		MoveMaxSpeed:   10,
		MoveTimeToMax:  20,
		WheelDelta:     1,
		WheelMaxSpeed:  16,
		WheelTimeToMax: 40,
	}
}

type MouseKeys struct {
	config MouseKeysConfig

	lastTimer time.Time
	repeat    uint16
	accel     uint16

	report struct {
		buttons uint8
		x       int8
		y       int8
		v       int8
		h       int8
	}

	debug bool
}

const (
	mousekeyMaxMove  = 127
	mousekeyMaxWheel = 127
)

func NewMouseKeys(config MouseKeysConfig) *MouseKeys {
	if config.MoveMaxSpeed > mousekeyMaxMove {
		config.MoveMaxSpeed = mousekeyMaxMove
	}
	if config.WheelMaxSpeed > mousekeyMaxWheel {
		config.WheelMaxSpeed = mousekeyMaxWheel
	}
	return &MouseKeys{
		config: config,
	}
}

func (mk *MouseKeys) Task(report *Report) bool {

	since := mk.config.Interval
	if mk.repeat == 0 {
		since = mk.config.Delay
	}
	if time.Since(mk.lastTimer) < since {
		return false
	}

	if mk.report.x == 0 && mk.report.y == 0 &&
		mk.report.v == 0 && mk.report.h == 0 {
		return false
	}

	if mk.repeat < 255 {
		mk.repeat++
	}

	if mk.report.x > 0 {
		mk.report.x = mk.moveUnit()
	}
	if mk.report.x < 0 {
		mk.report.x = int8(mk.moveUnit()) * -1
	}
	if mk.report.y > 0 {
		mk.report.y = mk.moveUnit()
	}
	if mk.report.y < 0 {
		mk.report.y = mk.moveUnit() * -1
	}

	// diagonal move [1/sqrt(2) = 0.7]
	if mk.report.x != 0 && mk.report.y != 0 {
		mk.report.x = int8(0.7 * float32(mk.report.x))
		mk.report.y = int8(0.7 * float32(mk.report.y))
	}

	if mk.report.x > 0 {
		mk.report.x = mk.moveUnit()
	}
	if mk.report.x < 0 {
		mk.report.x = int8(mk.moveUnit()) * -1
	}
	if mk.report.y > 0 {
		mk.report.y = mk.moveUnit()
	}
	if mk.report.y < 0 {
		mk.report.y = mk.moveUnit() * -1
	}

	report.Mouse(MouseButton(mk.report.buttons), mk.report.x, mk.report.y, mk.report.v, mk.report.h)
	return true

}

func (mk *MouseKeys) Make(code keycodes.Keycode) {
	switch code {
	case keycodes.KC_MS_UP:
		mk.report.y = mk.moveUnit() * -1
	case keycodes.KC_MS_DOWN:
		mk.report.y = mk.moveUnit()
	case keycodes.KC_MS_LEFT:
		mk.report.x = mk.moveUnit() * -1
	case keycodes.KC_MS_RIGHT:
		mk.report.x = mk.moveUnit()
	case keycodes.KC_MS_WH_UP:
		mk.report.v = mk.wheelUnit()
	case keycodes.KC_MS_WH_DOWN:
		mk.report.v = mk.wheelUnit() * -1
	case keycodes.KC_MS_WH_LEFT:
		mk.report.h = mk.wheelUnit() * -1
	case keycodes.KC_MS_WH_RIGHT:
		mk.report.h = mk.wheelUnit()
	case keycodes.KC_MS_BTN1:
		mk.report.buttons |= uint8(keycodes.KC_BTN1) & 0xD0 // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN2:
		mk.report.buttons |= uint8(keycodes.KC_BTN2) & 0xD0 // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN3:
		mk.report.buttons |= uint8(keycodes.KC_BTN3) & 0xD0 // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN4:
		mk.report.buttons |= uint8(keycodes.KC_BTN4) & 0xD0 // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN5:
		mk.report.buttons |= uint8(keycodes.KC_BTN5) & 0xD0 // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_ACCEL0:
		mk.accel |= (1 << 0)
	case keycodes.KC_MS_ACCEL1:
		mk.accel |= (1 << 1)
	case keycodes.KC_MS_ACCEL2:
		mk.accel |= (1 << 2)
	}
}

func (mk *MouseKeys) Break(code keycodes.Keycode) {
	switch code {
	case keycodes.KC_MS_UP:
		if mk.report.y < 0 {
			mk.report.y = 0
		}
	case keycodes.KC_MS_DOWN:
		if mk.report.y > 0 {
			mk.report.y = 0
		}
	case keycodes.KC_MS_LEFT:
		if mk.report.x < 0 {
			mk.report.x = 0
		}
	case keycodes.KC_MS_RIGHT:
		if mk.report.x > 0 {
			mk.report.x = 0
		}
	case keycodes.KC_MS_WH_UP:
		if mk.report.v > 0 {
			mk.report.v = 0
		}
	case keycodes.KC_MS_WH_DOWN:
		if mk.report.v < 0 {
			mk.report.v = 0
		}
	case keycodes.KC_MS_WH_LEFT:
		if mk.report.h < 0 {
			mk.report.h = 0
		}
	case keycodes.KC_MS_WH_RIGHT:
		if mk.report.h > 0 {
			mk.report.h = 0
		}
	case keycodes.KC_MS_BTN1:
		mk.report.buttons &= ^uint8(keycodes.KC_BTN1 & 0x0D) // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN2:
		mk.report.buttons &= ^uint8(keycodes.KC_BTN2 & 0x0D) // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN3:
		mk.report.buttons &= ^uint8(keycodes.KC_BTN3 & 0x0D) // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN4:
		mk.report.buttons &= ^uint8(keycodes.KC_BTN4 & 0x0D) // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_BTN5:
		mk.report.buttons &= ^uint8(keycodes.KC_BTN5 & 0x0D) // FIXME: hardcoded to QMK code
	case keycodes.KC_MS_ACCEL0:
		mk.accel &= ^uint16(1 << 0)
	case keycodes.KC_MS_ACCEL1:
		mk.accel &= ^uint16(1 << 1)
	case keycodes.KC_MS_ACCEL2:
		mk.accel &= ^uint16(1 << 2)
	}
	if mk.report.x == 0 && mk.report.y == 0 &&
		mk.report.v == 0 && mk.report.h == 0 {
		mk.repeat = 0
	}
}

func (mk *MouseKeys) moveUnit() int8 {
	var unit uint16
	if mk.accel&(1<<0) > 0 {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed) / 4
	} else if mk.accel&(1<<1) > 0 {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed) / 2
	} else if mk.accel&(1<<2) > 0 {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed)
	} else if mk.repeat == 0 {
		unit = mk.config.MoveDelta
	} else if mk.repeat >= mk.config.MoveTimeToMax {
		unit = mk.config.MoveDelta * mk.config.MoveMaxSpeed
	} else {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed * mk.repeat) / mk.config.MoveTimeToMax
	}
	if unit > mk.config.MoveMaxSpeed {
		return int8(mk.config.MoveMaxSpeed)
	} else if unit == 0 {
		return 1
	} else {
		return int8(unit)
	}
}

func (mk *MouseKeys) wheelUnit() int8 {
	var unit uint16
	if mk.accel&(1<<0) > 0 {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed) / 4
	} else if mk.accel&(1<<1) > 0 {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed) / 2
	} else if mk.accel&(1<<2) > 0 {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed)
	} else if mk.repeat == 0 {
		unit = mk.config.WheelDelta
	} else if mk.repeat >= mk.config.WheelTimeToMax {
		unit = mk.config.WheelDelta * mk.config.WheelMaxSpeed
	} else {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed * mk.repeat) / mk.config.WheelTimeToMax
	}
	if unit > mk.config.WheelMaxSpeed {
		return int8(mk.config.WheelMaxSpeed)
	} else if unit == 0 {
		return 1
	} else {
		return int8(unit)
	}
}

// func (mk *MouseKeys) WriteDebug(w io.Writer) {
// 	w.Write([]byte("mousekey [btn|x y v h](rep/acl): ["))
// 	w.Write([]byte(hex(mk.report.buttons)))
// 	w.Write([]byte("|"))
// 	w.Write([]byte(strconv.FormatInt(int64(mk.report.x), 10)))
// 	w.Write([]byte(" "))
// 	w.Write([]byte(strconv.FormatInt(int64(mk.report.y), 10)))
// 	w.Write([]byte(" "))
// 	w.Write([]byte(strconv.FormatInt(int64(mk.report.v), 10)))
// 	w.Write([]byte(" "))
// 	w.Write([]byte(strconv.FormatInt(int64(mk.report.h), 10)))
// 	w.Write([]byte("]("))
// 	w.Write([]byte(strconv.FormatInt(int64(mk.repeat), 10)))
// 	w.Write([]byte("/"))
// 	w.Write([]byte(strconv.FormatInt(int64(mk.accel), 10)))
// 	w.Write([]byte(")\n"))
// }
