//go:build !tinygo

// this file is currently a shim for code completion and to reduce IDE errors
// TODO: actually support non-TinyGo builds
package main

import (
	"time"

	serialhost "github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfs"
)

var (
	i2c    = (drivers.I2C)(nil)
	serial = (keyboard.Serialer)(nil)

	blockdev tinyfs.BlockDevice = nil
	keymapfs tinyfs.Filesystem  = nil
)

func adjustTimeOffset(t time.Time) {
}

func configureI2C() error {
	return nil
}

func initHost() keyboard.Host {
	return serialhost.New(nil)
}

func initFilesystem() {

}
