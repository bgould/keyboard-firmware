//go:build tinygo

package main

import (
	"machine"
	"machine/usb"
	"runtime"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
)

var (
	serial = machine.Serial
	driver = &VialDriver{usbvial.NewKeyboardDriver(keymap, matrix)}
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
	usb.Manufacturer = "Kinesis"
	usb.Product = "Advantage2"
	usb.Serial = vial.MagicSerialNumber("")
	return usbvial.New(VialDeviceDefinition, driver)
}
