package keyboard

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

// type Layers struct {
// 	Default uint8

// 	active uint32
// }

// func (layers *Layers) ActiveLayer() int {

// }

func (kbd *Keyboard) processLayer(key keycodes.Keycode, made bool) {
	if !key.IsLayer() { // sanity check
		return
	}
	switch {
	case keycodes.IsQkMomentary(key):
		layer := uint8(key - keycodes.QK_MOMENTARY)
		kbd.momentaryLayer(layer, made)
	case keycodes.IsQkToggleLayer(key):
		layer := uint8(key - keycodes.QK_TOGGLE_LAYER)
		kbd.toggleLayer(layer, made)
	default:
		// FIXME
		println("unsupported layer key", key, made)
	}
}

func (kbd *Keyboard) momentaryLayer(layer uint8, made bool) {
	layer &= 0x1F
	if made {
		kbd.SetActiveLayer(layer)
	} else {
		kbd.SetActiveLayer(kbd.defaultLayer)
	}
}

func (kbd *Keyboard) toggleLayer(layer uint8, made bool) {
	layer &= 0x1F
	keydown := kbd.layerToggles&(1<<layer) > 0
	if keydown && !made {
		if kbd.ActiveLayer() == layer {
			kbd.SetActiveLayer(kbd.defaultLayer)
		} else {
			kbd.SetActiveLayer(layer)
		}
	}
	if made {
		kbd.layerToggles |= (1 << layer)
	} else {
		kbd.layerToggles &= ^(1 << layer)
	}
}
