//go:build tinygo
// +build tinygo

package serial

import (
	"fmt"
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
	fmt.Fprintf(
		host.serial,
		"[ %02X %02X %02X %02X %02X %02X %02X %02X ]\n",
		rpt[0], rpt[1], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7],
	)
}
