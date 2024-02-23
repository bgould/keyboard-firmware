package keyboard

type BootloaderError int

const (
	ErrBootloaderNotSet BootloaderError = iota + 1
	ErrBootloaderNoDefault
)

type EnterBootloaderFunc func() error

func (err BootloaderError) Error() string {
	switch err {
	case ErrBootloaderNotSet:
		return "EnterBootloader: not set"
	case ErrBootloaderNoDefault:
		return "EnterBootloader: no default"
	default:
		return "EnterBootloader: unknown"
	}
}

func (kbd *Keyboard) EnterBootloader() error {
	if kbd.enterBootloader == nil {
		return ErrBootloaderNotSet
	}
	return kbd.enterBootloader()
}

func (kbd *Keyboard) CPUReset() error {
	if kbd.enterCpuReset == nil {
		return ErrBootloaderNotSet
	}
	return kbd.enterCpuReset()
}

var _ EnterBootloaderFunc = DefaultEnterBootloader
var _ EnterBootloaderFunc = DefaultCPUReset

func (kbd *Keyboard) SetEnterBootloaderFunc(fn EnterBootloaderFunc) {
	kbd.enterBootloader = fn
}

func (kbd *Keyboard) SetCPUResetFunc(fn EnterBootloaderFunc) {
	kbd.enterCpuReset = fn
}
