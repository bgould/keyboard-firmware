package keyboard

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

func (kbd *Keyboard) processLayer(key keycodes.Keycode, made bool) {
	if !key.IsLayer() { // sanity check
		return
	}
	switch {
	case keycodes.IsQkMomentary(key):
		println("momentary layer key", key, made)
	case keycodes.IsQkToggleLayer(key):
		println("toggle layer key", key, made)
	default:
		// FIXME
		println("unsupported layer key", key, made)
	}
}
