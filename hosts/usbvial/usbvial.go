//go:build tinygo

package usbvial

import (
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	debug = false
)

var (
	host *Host
)

type Host struct {
	*usbhid.Host
	km  KeyMapper
	txb [32]byte
}

func New(keymap KeyMapper) *Host {
	host = &Host{Host: usbhid.New(), km: keymap}
	return host
}

type KeyMapper interface {
	GetLayerCount() uint8
	GetMaxKeyCount() int
	NumRows() int
	NumCols() int
	MapKey(layer, idx int) keycodes.Keycode
}

type EncoderMapper interface {
	MapEncoder(layer, idx int) (ccw keycodes.Keycode, cw keycodes.Keycode)
}

type EncoderSaver interface {
	SaveEncoder(layer, idx int, clockwise bool, kc keycodes.Keycode)
}
