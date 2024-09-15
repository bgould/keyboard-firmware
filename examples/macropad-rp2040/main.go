//go:build macropad_rp2040

package main

import (
	"machine"
	"machine/usb"
	"time"

	"github.com/bgould/keyboard-firmware/boards/macropad_rp2040"
	"github.com/bgould/keyboard-firmware/keyboard/hosts/usbvial/vial"
)

var (
	board     = macropad_rp2040.New()
	kbd, host = board.NewVialKeyboard(layers...)
)

func main() {
	// set up the hardware
	board.Configure()

	// configure USB interface with Vial
	usb.Manufacturer = macropad_rp2040.USBManufacturer
	usb.Product = macropad_rp2040.USBProduct
	usb.Serial = vial.MagicSerialNumber("")
	host.Configure()

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
