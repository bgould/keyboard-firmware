//go:build tinygo

package usbvial

import (
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
)

const (
	debug = false
)

type Host struct {
	*usbhid.Host
}

func New() *Host {
	return &Host{usbhid.New()}
}
