package vial

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

type KeyMapper interface {
	GetLayerCount() uint8
	GetMaxKeyCount() int
	NumRows() int
	NumCols() int
	MapKey(layer, row, col int) keycodes.Keycode
}

type KeySetter interface {
	SetKey(layer, row, col int, kc keycodes.Keycode) bool
}

type EncoderMapper interface {
	MapEncoder(layer, idx int) (ccw keycodes.Keycode, cw keycodes.Keycode)
}

type EncoderSaver interface {
	SaveEncoder(layer, idx int, clockwise bool, kc keycodes.Keycode)
}

type Handler interface {
	Handle(rx []byte, tx []byte) bool
}

type HandlerFunc func(rx []byte, tx []byte) (sendTx bool)

func (fn HandlerFunc) Handle(rx []byte, tx []byte) (sendTx bool) {
	return fn(rx, tx)
}
