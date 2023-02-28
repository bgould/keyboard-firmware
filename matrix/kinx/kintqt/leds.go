package kintqt

import "tinygo.org/x/drivers/mcp23017"

type LED uint8

const (
	LEDNumLock    LED = 7
	LEDUser1      LED = 8
	LEDUser2      LED = 9
	LEDUser3      LED = 10
	LEDUser4      LED = 11
	LEDUser5      LED = 12
	LEDUser6      LED = 13
	LEDCapsLock   LED = 14
	LEDKeypad     LED = 15
	LEDScrollLock LED = 31

	ledMask = uint32(0 |
		1<<LEDScrollLock |
		1<<LEDCapsLock |
		1<<LEDNumLock |
		1<<LEDKeypad)

	// mask for LED/GPIO output pins on both ports
	port0_ledMask = mcp23017.Pins(ledMask & 0xFFFF)
	port1_ledMask = mcp23017.Pins(ledMask >> 16)
)

type LEDs uint32

func (l *LEDs) Get(led LED) bool {
	return (*l & 1 << led) > 0
}

func (l *LEDs) Set(led LED, on bool) {
	if on {
		*l |= (1 << led) // high
	} else {
		*l &^= (1 << led) // low
	}
}

func (l *LEDs) port0state() mcp23017.Pins {
	return ^mcp23017.Pins(*l) & port0_ledMask
}

func (l *LEDs) port1state() mcp23017.Pins {
	return ^mcp23017.Pins(*l>>16) & port1_ledMask
}
