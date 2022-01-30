//go:build tinygo
// +build tinygo

package serial

import (
	"io"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const debug = false

type Host struct {
	serial io.Writer
}

func New(serial io.Writer) *Host {
	return &Host{
		serial: serial,
	}
}

func (host *Host) Send(rpt *keyboard.Report) {
	host.serial.Write([]byte(rpt.String() + "\r\n"))
}

func (host *Host) LEDs() uint8 {
	return 0
}
