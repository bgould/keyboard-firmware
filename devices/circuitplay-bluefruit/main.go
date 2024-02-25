//go:build tinygo && circuitplay_bluefruit

package main

import (
	"machine"
	"machine/usb"

	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

var (
	buttons = []machine.Pin{machine.BUTTONA, machine.BUTTONB}
	slider  = machine.SLIDER
	keymap  = CircuitPlaygroundDefaultKeymap()
	matrix  = keyboard.NewMatrix(1, 3, keyboard.RowReaderFunc(ReadRow))
	host    = usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)
	board   = keyboard.New(host, matrix, keymap)

	backlight = keyboard.Backlight{
		Driver:       &keyboard.BacklightGPIO{LED: machine.LED, PWM: machine.PWM0},
		DefaultMode:  keyboard.BacklightBreathing,
		DefaultLevel: 0xFF,
	}
)

func init() {
	board.SetBacklight(backlight)
	board.SetKeyAction(keyboard.KeyActionFunc(keyAction))
}

func main() {

	backlight.Driver.Configure()

	configurePins()

	usb.Serial = vial.MagicSerialNumber("")
	host.Configure()

	for {
		board.Task()
	}

}

func configurePins() {
	slider.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	for _, pin := range buttons {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
		println("configured button", pin, pin.Get())
	}
}

func ReadRow(rowIndex uint8) keyboard.Row {
	switch rowIndex {
	case 0:
		v := keyboard.Row(0)
		for i := range buttons {
			if buttons[i].Get() {
				v |= (1 << i)
			}
		}
		if !slider.Get() {
			v |= (1 << len(buttons))
		}
		return v
	default:
		return 0
	}

}

// TODO: natively support momentary layer switching keycodes
func keyAction(key keycodes.Keycode, made bool) {
	switch key {
	// Toggle function layer on key down/up
	case keycodes.KC_FN1:
		if made {
			board.SetActiveLayer(1)
			println("layer 1 on")
		} else {
			board.SetActiveLayer(0)
			println("layer 1 off")
		}
	}
}
