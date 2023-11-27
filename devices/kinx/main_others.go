//go:build !tinygo

// this file is currently a shim for code completion and to reduce IDE errors
// TODO: actually support non-TinyGo builds
package main

import (
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers"
)

var (
	i2c    = (drivers.I2C)(nil)
	serial = (keyboard.Serialer)(nil)
)

func adjustTimeOffset(t time.Time) {
}
