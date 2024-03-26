//go:build !tinygo

package keyboard

import (
	"time"
)

func adjustTimeOffset(t time.Time) bool {
	return false
}

func usbManufacturer() string { return "" }
func usbProduct() string      { return "" }
func usbSerial() string       { return "" }

func disableInterrupts() (mask uintptr) { return 0 }
func restoreInterrupts(mask uintptr)    {}
