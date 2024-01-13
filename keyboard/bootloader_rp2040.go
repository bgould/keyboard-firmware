//go:build tinygo && rp2040

package keyboard

import (
	"machine"
)

func DefaultEnterBootloader() error {
	machine.EnterBootloader()
	return nil
}

var _ EnterBootloaderFunc = DefaultEnterBootloader
