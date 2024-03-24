//go:build tinygo

package keyboard

import (
	"device/arm"
	"machine/usb"
	"runtime"
	"time"
)

func adjustTimeOffset(t time.Time) bool {
	runtime.AdjustTimeOffset(-1 * int64(time.Since(t)))
	return true
}

func usbManufacturer() string { return usb.Manufacturer }
func usbProduct() string      { return usb.Product }
func usbSerial() string       { return usb.Serial }

func disableInterrupts() (mask uintptr) {
	return arm.DisableInterrupts()
}

func restoreInterrupts(mask uintptr) {
	arm.EnableInterrupts(mask)
}
