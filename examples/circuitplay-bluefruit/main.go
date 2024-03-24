//go:build circuitplay_bluefruit

package main

import (
	"machine"
	"machine/usb"
	"time"

	"github.com/bgould/keyboard-firmware/boards/circuitplay_bluefruit"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
)

var (
	board     = circuitplay_bluefruit.New()
	kbd, host = board.NewVialKeyboard(layers...)
)

func main() {
	// set up the hardware
	board.Configure()

	// configure USB interface with Vial
	usb.Manufacturer = circuitplay_bluefruit.USBManufacturer
	usb.Product = circuitplay_bluefruit.USBProduct
	usb.Serial = vial.MagicSerialNumber("")
	host.Configure()

	time.Sleep(time.Second)

	// FIXME: this should be handled in the core based on a single call to keyboard.Configure() or something
	if kbd.FS() != nil {
		kbd.ConfigureFilesystem()
	}
	kbd.EnableConsole(machine.Serial)

	// task loop
	for {
		kbd.Task()
		time.Sleep(500 * time.Microsecond)
	}
}
