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

	// blockdev tinyfs.BlockDevice = machine.Flash
	// keymapfs tinyfs.Filesystem  = littlefs.New(blockdev)
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

type VialDriver struct {
	*usbvial.KeyboardDeviceDriver
}

var (
	_ vial.Handler      = (*VialDriver)(nil)
	_ vial.DeviceDriver = (*VialDriver)(nil)
)

func (d *VialDriver) Handle(rx []byte, tx []byte) (sendTx bool) {
	// println("called Handle()", rx[0], rx[1])
	switch rx[0] {
	case 0xEE:
		switch rx[1] {
		case 0x01: // set time
			var unixTime uint64
			unixTime |= uint64(rx[2]) << 56
			unixTime |= uint64(rx[3]) << 48
			unixTime |= uint64(rx[4]) << 40
			unixTime |= uint64(rx[5]) << 32
			unixTime |= uint64(rx[6]) << 24
			unixTime |= uint64(rx[7]) << 16
			unixTime |= uint64(rx[8]) << 8
			unixTime |= uint64(rx[9]) << 0
			println("\nsetting unix time", int64(unixTime))
			setUnixTime(time.Unix(int64(unixTime), 0))
			// cmd := console.CommandInfo{
			// 	Cmd:    "time",
			// 	Argv:   []string{"set", strconv.FormatInt(int64(unixTime), 10)},
			// 	Stdout: cli,
			// }
			// timeset(cmd)
		}
	}
	return false
}
