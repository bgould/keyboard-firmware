package kint_pe

import "tinygo.org/x/drivers/mcp23017"

const (
	ledKeypad     = 7
	ledNumLock    = 14
	ledCapsLock   = 15
	ledScrollLock = 31

	ledMask = uint32(0 |
		1<<ledScrollLock |
		1<<ledCapsLock |
		1<<ledNumLock |
		1<<ledKeypad)
)

type LEDs uint16

func (l *LEDs) set(port0, port1 *mcp23017.Device) error {
	return nil
}
