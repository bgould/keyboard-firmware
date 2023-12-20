//go:build tinygo

package usbvial

import (
	"machine"
	"machine/usb"
	"machine/usb/descriptor"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func init() {
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
			{
				Index:   usb.HID_INTERFACE,
				Handler: setupHandler,
			},
		})
}

var (
	txb [256]byte
	// Keys        [][][]Keycode // [row][col]Keycode
	Changed     bool
	Changed2    bool
	wbuf        []byte
	KeyboardDef []byte
	// dev2        *keyboard.Device
	device KeyMapper // *keyboard.Device
	// mapper KeyMapper
)

type KeyMapper interface {
	GetLayerCount() uint8
	GetMaxKeyCount() int
	MapKey(layer, idx int) keycodes.Keycode
}

func SetDevice(d KeyMapper) {
	device = d
}

func rxHandler(b []byte) {
	rxHandler2(b)
}

func rxHandler2(b []byte) bool {
	switch b[0] {
	//case 0x12, 0x0E:
	default:
		//fmt.Printf("RxHandler % X\n", b)
	}

	copy(txb[:32], b)
	switch b[0] {
	case 0x01:
		// println("usb: 0x01 - GetProtocolVersionCount")
		// GetProtocolVersionCount
		txb[2] = 0x09
	case 0x11:
		// println("7sb: 0x11 - DynamicKeymapGetLayerCountCommand")
		// DynamicKeymapGetLayerCountCommand
		if device != nil {
			txb[1] = device.GetLayerCount()
		} else {
			txb[1] = 0x01
		}
	case 0x12:
		// println("cb: 0x12 - DynamicKeymapReadBufferCommand")
		if device == nil {
			println("warning: device was nil")
			break
		}
		// DynamicKeymapReadBufferCommand
		offset := (uint16(b[1]) << 8) + uint16(b[2])
		sz := b[3]
		//fmt.Printf("  offset : %04X + %d\n", offset, sz)
		cnt := device.GetMaxKeyCount()
		// println("  offset : ", offset, "+", sz, cnt)
		// break
		for i := 0; i < int(sz/2); i++ {
			//fmt.Printf("  %02X %02X\n", b[4+i+1], b[4+i+0])
			tmp := i + int(offset)/2
			layer := tmp / (cnt * 1) // len(device.kb))
			tmp = tmp % (cnt * 1)    // len(device.kb))
			kbd := tmp / cnt
			idx := tmp % cnt
			//layer := 0
			//idx := tmp & 0xFF
			kc := keyVia(layer, kbd, idx)
			// println(layer, idx, kc)
			//fmt.Printf("  (%d, %d, %d)\n", layer, kbd, idx)
			txb[4+2*i+1] = uint8(kc)
			txb[4+2*i+0] = uint8(kc >> 8)
		}

	case 0x0D:
		// println("Dsb: 0x0D - DynamicKeymapMacroGetBufferSizeCommand")
		// DynamicKeymapMacroGetBufferSizeCommand
		txb[1] = 0x07
		txb[2] = 0x9B
	case 0x0C:
		// println("Csb: 0x0C - DynamicKeymapMacroGetCountCommand")
		// DynamicKeymapMacroGetCountCommand
		txb[1] = 0x10
	case 0x0E:
		// println("Esb: 0x0E - DynamicKeymapMacroGetBufferCommand")
		// DynamicKeymapMacroGetBufferCommand
	case 0x02:
		// println("2usb: 0x02 - id_get_keyboard_value")
		// id_get_keyboard_value
		Changed = false
		Changed2 = false
	case 0x05:
		// println("5sb: 0x05 - ", len(b), b[1], b[2], b[3], b[4], b[5])
		// fmt.Printf("XXXXXXXXX % X\n", b)
		//Keys[b[1]][b[2]][b[3]] = Keycode((uint16(b[4]) << 8) + uint16(b[5]))
		// device.SetKeycodeVia(int(b[1]), int(b[2]), int(b[3]), Keycode((uint16(b[4])<<8)+uint16(b[5])))
		// device.flashCh <- true
		//Changed = true
	case 0x08:
		// println("8sb: 0x08 - id_lighting_get_value")
		// id_lighting_get_value
		txb[1] = 0x00
		txb[2] = 0x00
	case 0xFE: // vial
		switch b[1] {
		case 0x00:
			// println("vial: 0x00 - Get keyboard ID and Vial protocol version")
			// Get keyboard ID and Vial protocol version
			const vialProtocolVersion = 0x00000006
			txb[0] = vialProtocolVersion
			txb[1] = vialProtocolVersion >> 8
			txb[2] = vialProtocolVersion >> 16
			txb[3] = vialProtocolVersion >> 24
			txb[4] = 0x9D
			txb[5] = 0xD0
			txb[6] = 0xD5
			txb[7] = 0xE1
			txb[8] = 0x87
			txb[9] = 0xF3
			txb[10] = 0x54
			txb[11] = 0xE2
		case 0x01:
			// println("vial: 0x01 - retrieve keyboard definition size")
			// Retrieve keyboard definition size
			size := len(KeyboardDef)
			txb[0] = uint8(size)
			txb[1] = uint8(size >> 8)
			txb[2] = uint8(size >> 16)
			txb[3] = uint8(size >> 24)
		case 0x02:
			// Retrieve 32-bytes block of the definition, page ID encoded within 2 bytes
			// println("vial: 0x02 - retrieve 32 byte blocks of definition")
			page := uint16(b[2]) + (uint16(b[3]) << 8)
			start := page * 32
			end := start + 32
			if end < start || int(start) >= len(KeyboardDef) {
				return false
			}
			if int(end) > len(KeyboardDef) {
				end = uint16(len(KeyboardDef))
			}
			//fmt.Printf("vial_get_def : page=%04X start=%04X end=%04X\n", page, start, end)
			copy(txb[:32], KeyboardDef[start:end])
		case 0x09:
			// println("vial: 0x09 - vial_qmk_settings_query")
			// vial_qmk_settings_query
			// 未対応
			for i := range txb[:32] {
				txb[i] = 0xFF
			}
		case 0x0D:
			// println("vial: 0x0D - vial_dynamic_entry_op")
			// vial_dynamic_entry_op
			txb[0] = 0x00
			txb[1] = 0x00
			txb[2] = 0x00
		case 0x05:
			// println("vial: 0x0D - vial_get_unlock_status")
			// vial_get_unlock_status
			txb[0] = 1 // unlocked
			txb[1] = 0 // unlock_in_progress
		default:
			// println("vial: default - ", b[1])
		}
	default:
		return false
	}
	machine.SendUSBInPacket(6, txb[:32])
	//fmt.Printf("Tx        % X\n", txb[:32])

	return true
}

// func Save() error {
// 	layers := 6
// 	keyboards := len(device.kb)

// 	cnt := device.GetMaxKeyCount()
// 	wbuf := make([]byte, 4+layers*keyboards*cnt*2)
// 	needed := int64(len(wbuf)) / machine.Flash.EraseBlockSize()
// 	if needed == 0 {
// 		needed = 1
// 	}

// 	err := machine.Flash.EraseBlocks(0, needed)
// 	if err != nil {
// 		return err
// 	}

// 	// TODO: Size should be written last
// 	sz := machine.Flash.Size()
// 	wbuf[0] = byte(sz >> 24)
// 	wbuf[1] = byte(sz >> 16)
// 	wbuf[2] = byte(sz >> 8)
// 	wbuf[3] = byte(sz >> 0)

// 	offset := 4
// 	for layer := 0; layer < layers; layer++ {
// 		for keyboard := 0; keyboard < keyboards; keyboard++ {
// 			for key := 0; key < cnt; key++ {
// 				wbuf[offset+2*key+0] = byte(device.Key(layer, keyboard, key) >> 8)
// 				wbuf[offset+2*key+1] = byte(device.Key(layer, keyboard, key))
// 			}
// 			offset += cnt * 2
// 		}
// 	}

// 	_, err = machine.Flash.WriteAt(wbuf[:], 0)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func setupHandler(setup usb.Setup) bool {
	ok := false
	if setup.BmRequestType == usb.SET_REPORT_TYPE && setup.BRequest == usb.SET_IDLE {
		machine.SendZlp()
		ok = true
	}
	return ok
}

func keyVia(layer, kbIndex, index int) uint16 {
	//fmt.Printf("    KeyVia(%d, %d, %d)\n", layer, kbIndex, index)
	if kbIndex > 0 { // TODO: support multiple keyboards?
		return 0
	}
	if device == nil {
		return 0
	}
	kc := uint16(device.MapKey(layer, index))
	switch kc {
	// case jp.MouseLeft:
	// 	kc = 0x00D1
	// case jp.MouseRight:
	// 	kc = 0x00D2
	// case jp.MouseMiddle:
	// 	kc = 0x00D3
	// case jp.MouseBack:
	// 	kc = 0x00D4
	// case jp.MouseForward:
	// 	kc = 0x00D5
	// case jp.WheelUp:
	// 	kc = 0x00D9
	// case jp.WheelDown:
	// 	kc = 0x00DA
	// case jp.KeyMediaVolumeInc:
	// 	kc = 0x00A9
	// case jp.KeyMediaVolumeDec:
	// 	kc = 0x00AA
	// case 0xFF10, 0xFF11, 0xFF12, 0xFF13, 0xFF14, 0xFF15:
	// 	// TO(x)
	// 	kc = 0x5200 | (kc & 0x000F)
	// case 0xFF00, 0xFF01, 0xFF02, 0xFF03, 0xFF04, 0xFF05:
	// 	// MO(x)
	// 	kc = 0x5220 | (kc & 0x000F)
	// case keycodes.KeyRestoreDefaultKeymap:
	// 	// restore default keymap for QMK
	// 	kc = keycodes.KeyRestoreDefaultKeymap
	default:
		kc = kc & 0x0FFF
	}
	return kc
}
