//go:build macropad_rp2040

package main

import (
	"machine"
	"machine/usb"
	"time"

	"github.com/bgould/keyboard-firmware/boards/macropad_rp2040"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
)

var (
	board     = macropad_rp2040.New()
	kbd, host = board.NewVialKeyboard(layers...)
)

func main() {

	// set up the hardware
	board.Configure()

	// FIXME: this should be handled in the core based on a single call to keyboard.Configure() or something
	if kbd.FS() != nil {
		println("configuring filesystem")
		kbd.ConfigureFilesystem()
	}

	// configure USB interface with Vial
	usb.Manufacturer = macropad_rp2040.USBManufacturer
	usb.Product = macropad_rp2040.USBProduct
	usb.Serial = vial.MagicSerialNumber("")
	host.Configure()

	kbd.EnableConsole(machine.Serial)

	// task loop
	for {
		kbd.Task()
		time.Sleep(500 * time.Microsecond)
	}
}
