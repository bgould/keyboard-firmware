package vial

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

// type KeyMapper interface {
// 	GetLayerCount() uint8
// 	NumRows() int
// 	NumCols() int
// 	GetMaxKeyCount() int
// 	MapKey(layer, row, col int) keycodes.Keycode
// }

// type Matrixer interface {
// 	NumRows() int
// 	NumCols() int
// }

// type KeySetter interface {
// 	SetKey(layer, row, col int, kc keycodes.Keycode) bool
// }

type EncoderMapper interface {
	MapEncoder(idx int) (ccwRow, ccwCol int, cwRow, cwCol int, ok bool)
	// (ccw keycodes.Keycode, cw keycodes.Keycode)
}

type EncoderSetter interface {
	SetEncoder(layer, idx int, clockwise bool, kc keycodes.Keycode)
}

type Handler interface {
	Handle(rx []byte, tx []byte) bool
}

type HandlerFunc func(rx []byte, tx []byte) (sendTx bool)

func (fn HandlerFunc) Handle(rx []byte, tx []byte) (sendTx bool) {
	return fn(rx, tx)
}
