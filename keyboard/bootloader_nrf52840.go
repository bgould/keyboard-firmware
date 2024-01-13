//go:build tinygo && nrf52840

package keyboard

import (
	"machine"
)

func DefaultCPUReset() error {
	machine.CPUReset()
	return nil
}

func DefaultEnterBootloader() error {
	machine.EnterUF2Bootloader()
	return nil
}

var _ EnterBootloaderFunc = DefaultEnterBootloader
var _ EnterBootloaderFunc = DefaultCPUReset
