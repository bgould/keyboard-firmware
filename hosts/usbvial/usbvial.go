//go:build tinygo

package usbvial

import (
	"machine"
	"machine/usb"
	"machine/usb/descriptor"

	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
)

const (
	debug = true
)

var (
	host *Host
)

type Host struct {
	*usbhid.Host
	dev *vial.Device
	txb [32]byte
}

func New(def vial.DeviceDefinition, keymap vial.KeyMapper) *Host {
	host = &Host{Host: usbhid.New(), dev: vial.NewDevice(def, keymap)}
	return host
}

// TODO: consider moving out of init
func init() {
	configureVialEndpoints()
}

func configureVialEndpoints() {

	descriptor.CDCHID.Configuration[2] = 0x84
	descriptor.CDCHID.Configuration[3] = 0x00
	descriptor.CDCHID.Configuration[4] = 0x04

	descriptor.CDCHID.Configuration = append(descriptor.CDCHID.Configuration, []byte{
		// 32 byte

		// Interface Descriptor
		0x09, 0x04, 0x03, 0x00, 0x02, 0x03, 0x00, 0x00, 0x00,
		// Length: 9 bytes
		// Descriptor Type: Interface (0x04)
		// Interface Number: 3
		// Alternate Setting: 0
		// Number of Endpoints: 2
		// Interface Class: 3 (HID - Human Interface Device)
		// Interface Subclass: 0
		// Interface Protocol: 0
		// Interface String Descriptor Index: 0 (No string descriptor)

		// HID Descriptor
		0x09, 0x21, 0x11, 0x01, 0x00, 0x01, 0x22, 0x22, 0x00,
		// Length: 9 bytes
		// Descriptor Type: HID (0x21)
		// HID Class Specification Release: 1.11
		// Country Code: 0 (Not localized)
		// Number of Descriptors: 1
		// Descriptor Type: Report (0x22)
		// Descriptor Length: 34 bytes (0x0022)

		// Endpoint Descriptor
		0x07, 0x05, 0x86, 0x03, 0x20, 0x00, 0x01,
		// Length: 7 bytes
		// Descriptor Type: Endpoint (0x05)
		// Endpoint Address: 0x86 (Endpoint 6, IN direction)
		// Attributes: 3 (Interrupt transfer type)
		// Maximum Packet Size: 32 bytes (0x0020)
		// Interval: 1 ms

		// Endpoint Descriptor
		0x07, 0x05, 0x07, 0x03, 0x20, 0x00, 0x01,
		// Length: 7 bytes
		// Descriptor Type: Endpoint (0x05)
		// Endpoint Address: 0x07 (Endpoint 7, OUT direction)
		// Attributes: 3 (Interrupt transfer type)
		// Maximum Packet Size: 32 bytes (0x0020)
		// Interval: 1 ms

	}...)

	descriptor.CDCHID.HID[3] = []byte{
		0x06, 0x60, 0xff, // Usage Page (Vendor-Defined 0xFF60)
		0x09, 0x61, // Usage (Vendor-Defined 0x61)
		0xa1, 0x01, // Collection (Application)
		0x09, 0x62, //   Usage (Vendor-Defined 0x62)
		0x15, 0x00, //   Logical Minimum (0)
		0x26, 0xff, 0x00, //   Logical Maximum (255)
		0x95, 0x20, //   Report Count (32)
		0x75, 0x08, //   Report Size (8)
		0x81, 0x02, //   Input (Data, Var, Abs)
		0x09, 0x63, //   Usage (Vendor-Defined 0x63)
		0x15, 0x00, //   Logical Minimum (0)
		0x26, 0xff, 0x00, //   Logical Maximum (255)
		0x95, 0x20, //   Report Count (32)
		0x75, 0x08, //   Report Size (8)
		0x91, 0x02, //   Output (Data, Var, Abs)
		0xc0, // End Collection
	}

	machine.ConfigureUSBEndpoint(descriptor.CDCHID,
		[]usb.EndpointConfig{
			{
				Index:     usb.MIDI_ENDPOINT_OUT,
				IsIn:      false,
				Type:      usb.ENDPOINT_TYPE_INTERRUPT,
				RxHandler: rxHandler,
			},
			{
				Index: usb.MIDI_ENDPOINT_IN,
				IsIn:  true,
				Type:  usb.ENDPOINT_TYPE_INTERRUPT,
			},
		},
		[]usb.SetupConfig{
			// {
			// 	Index:   usb.HID_INTERFACE,
			// 	Handler: hid.DefaultSetupHandler, // setupHandler,
			// },
		})
}

// func setupHandler(setup usb.Setup) bool {
// 	ok := false
// 	if setup.BmRequestType == usb.SET_REPORT_TYPE && setup.BRequest == usb.SET_IDLE {
// 		machine.SendZlp()
// 		ok = true
// 	}
// 	return ok
// }

func rxHandler(b []byte) {
	if host == nil {
		return
	}
	rx := b[:32]
	tx := host.txb[:32]
	if host.dev.Handle(rx, tx) {
		machine.SendUSBInPacket(6, tx)
	}
}
