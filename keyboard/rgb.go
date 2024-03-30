package keyboard

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

func (kbd *Keyboard) processRgb(key keycodes.Keycode, made bool) {
	// if !kbd.BacklightEnabled() {
	// 	return
	// }
	// kbd.backlight.ProcessKey(key, made)
	var hsv = &kbd.backlight.state.color
	var step uint8 = 17
	var newval, changed = uint8(0), false
	switch key {
	case keycodes.RGB_TOG:
		println("RGB toggle")
	case keycodes.RGB_MODE_FORWARD:
		println("RGB mode forward")
	case keycodes.RGB_MODE_REVERSE:
		println("RGB mode reverse")
	case keycodes.RGB_HUI:
		// println("RGB hue increase", made)
		if made {
			newval, changed = hsv.HueInc(5)
			// println("RGB hue increase:", newval, changed)
		}
	case keycodes.RGB_HUD:
		// println("RGB hue decrease", made)
		if made {
			newval, changed = hsv.HueDec(5)
		}
	case keycodes.RGB_SAI:
		// println("RGB sat increase", made)
		if made {
			newval, changed = hsv.SatInc(5)
			// println("RGB hue increase:", newval, changed)
		}
	case keycodes.RGB_SAD:
		// println("RGB sat decrease", made)
		if made {
			newval, changed = hsv.SatDec(5)
			// println("RGB hue decrease:", newval, changed)
		}
	case keycodes.RGB_VAI:
		// println("RGB val increase", made)
		if made {
			newval, changed = hsv.ValInc(step)
			// println("RGB val increase:", newval, changed)
		}
	case keycodes.RGB_VAD:
		// println("RGB val decrease", made)
		if made {
			newval, changed = hsv.ValDec(step)
			// println("RGB val decrease:", newval, changed)
		}
	// case RGB_SPI:
	// case RGB_SPD:
	// case RGB_MODE_PLAIN:
	// case RGB_MODE_BREATHE:
	// case RGB_MODE_RAINBOW:
	// case RGB_MODE_SWIRL:
	// case RGB_MODE_SNAKE:
	// case RGB_MODE_KNIGHT:
	// case RGB_MODE_XMAS:
	// case RGB_MODE_GRADIENT:
	// case RGB_MODE_RGBTEST:
	// case RGB_MODE_TWINKLE:
	default:
		if made {
			println("processRgb():", key)
		}
	}
	_, _ = newval, changed
	if changed {
		// println("backlight color: ", hsv.H, hsv.S, hsv.V)
		kbd.backlight.Sync()
	}
}
