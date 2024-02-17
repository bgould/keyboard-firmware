//go:build !tinygo || (!nrf52840 && !rp2040)

package keyboard

func DefaultCPUReset() error {
	return ErrBootloaderNoDefault
}

func DefaultEnterBootloader() error {
	return ErrBootloaderNoDefault
}
