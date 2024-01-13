//go:build !tinygo || (!nrf52840 && !rp2040)

package keyboard

func DefaultEnterBootloader() error {
	return ErrBootloaderNoDefault
}

var _ EnterBootloaderFunc = DefaultEnterBootloader
