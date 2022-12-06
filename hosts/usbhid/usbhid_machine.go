//go:build tinygo && usbhid_machine
// +build tinygo,usbhid_machine

package usbhid

import machine_kb "machine/usb/hid/keyboard"

var kb = machine_kb.New()

func sendKeyboardReport(mod, k1, k2, k3, k4, k5, k6 byte) {
	kb.SendReport(mod, k1, k2, k3, k4, k5, k6)
}

func sendMouseReport(buttons, x, y, wheel byte) {

}
