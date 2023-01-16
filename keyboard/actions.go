package keyboard

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

type KeyAction interface {
	KeyAction(key keycodes.Keycode, made bool)
}

type KeyActionFunc func(key keycodes.Keycode, made bool)

func (fn KeyActionFunc) KeyAction(key keycodes.Keycode, made bool) {
	fn(key, made)
}

var _ KeyActionFunc = (KeyActionFunc)(nil)
