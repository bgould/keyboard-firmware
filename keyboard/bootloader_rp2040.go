//go:build tinygo && rp2040

package keyboard

import (
	"machine"
)

func DefaultCPUReset() error {
	machine.CPUReset()
	return nil
}

func DefaultEnterBootloader() error {
	machine.EnterBootloader()
	return nil
}
