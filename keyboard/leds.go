package keyboard

type LED uint8

const (
	LEDNumLock    LED = 1
	LEDCapsLock   LED = 2
	LEDScrollLock LED = 3
)

type LEDs uint8

func (l *LEDs) Get(led LED) bool {
	if led == 0 {
		return false
	}
	return (*l & (1 << (led - 1))) > 0
}

func (l *LEDs) Set(led LED, on bool) {
	if on {
		*l |= (1 << led) // high
	} else {
		*l &^= (1 << led) // low
	}
}
