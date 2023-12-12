//go:build tinygo

package main

import (
	"machine"
	"runtime"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
)

var (
	serial = machine.Serial
)

func adjustTimeOffset(t time.Time) {
	runtime.AdjustTimeOffset(-1 * int64(time.Since(t)))
}

func configureI2C() error {
	return i2c.Configure(machine.I2CConfig{
		Frequency: 1 * machine.MHz,
	})
}

func initHost() keyboard.Host {
	return usbhid.New()
}
