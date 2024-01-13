//go:build tinygo && nrf52840

package keyboard

import (
	"machine"
)

func DefaultEnterBootloader() error {
	machine.EnterUF2Bootloader()
	return nil
}

var _ EnterBootloaderFunc = DefaultEnterBootloader
