package usbvial

import "github.com/bgould/keyboard-firmware/keyboard/keycodes"

const (
	// MagicSerialPrefix is a value in the serial number of HID devices
	// that the	Vial desktop app uses to identify compatible devices.
	MagicSerialPrefix = "vial:f64c2b3c"
)

// MagicSerialNumber returns a string value that the Vial desktop app
// can recognize as a Vial-compatible device based on a "magic" value
// (see MagicSerialPrefix constant in this package).  If the provided
// string `sn` is not the zero value, it is appended with a separator
// to the prefix and returned; otherwise just the prefix is returned.
func MagicSerialNumber(sn string) string {
	if sn != "" {
		return MagicSerialPrefix + ":" + sn
	}
	return MagicSerialPrefix
}

// ViaCommand represents a command from the VIA command set. VIA is a
// graphical configurator for QMK firmware. Vial is an open source
// alternative to VIA, and shares some of the same command set as its
// predecessor. Not all VIA commands are supported by this package,
// only the ones that are necessary for Vial.
type ViaCommand uint8

//go:generate go run golang.org/x/tools/cmd/stringer -type=ViaCommand

const (
	ViaCmdGetProtocolVersion       ViaCommand = 0x01
	ViaCmdGetKeyboardValue         ViaCommand = 0x02
	ViaCmdSetKeyboardValue         ViaCommand = 0x03
	ViaCmdDynamicKeymapGetKeycode  ViaCommand = 0x04
	ViaCmdDynamicKeymapSetKeycode  ViaCommand = 0x05
	ViaCmdDynamicKeymapReset       ViaCommand = 0x06
	ViaCmdLightingSetValue         ViaCommand = 0x07
	ViaCmdLightingGetValue         ViaCommand = 0x08
	ViaCmdLightingSave             ViaCommand = 0x09
	ViaCmdEepromReset              ViaCommand = 0x0A
	ViaCmdBootloaderJump           ViaCommand = 0x0B
	ViaCmdKeymapMacroGetCount      ViaCommand = 0x0C
	ViaCmdKeymapMacroGetBufferSize ViaCommand = 0x0D
	ViaCmdKeymapMacroGetBuffer     ViaCommand = 0x0E
	ViaCmdKeymapMacroSetBuffer     ViaCommand = 0x0F
	ViaCmdKeymapMacroReset         ViaCommand = 0x10
	ViaCmdKeymapGetLayerCount      ViaCommand = 0x11
	ViaCmdKeymapGetBuffer          ViaCommand = 0x12
	ViaCmdKeymapSetBuffer          ViaCommand = 0x13
	ViaCmdVialPrefix               ViaCommand = 0xFE
	ViaCmdUnhandled                ViaCommand = 0xFF
)

type ViaKeyboardValueID uint8

const (
	ViaKbdUptime            ViaKeyboardValueID = 0x01
	ViaKbdLayoutOptions                        = 0x02
	ViaKbdSwitchMatrixState                    = 0x03
	ViaKbdFirmwareVersion                      = 0x04
	ViaKbdDeviceIndication                     = 0x05
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=VialCommand

type VialCommand uint8

const (
	VialCmdGetKeyboardID    VialCommand = 0x00
	VialCmdGetSize          VialCommand = 0x01
	VialCmdGetDef           VialCommand = 0x02
	VialCmdGetEncoder       VialCommand = 0x03
	VialCmdSetEncoder       VialCommand = 0x04
	VialCmdGetUnlockStatus  VialCommand = 0x05
	VialCmdUnlockStart      VialCommand = 0x06
	VialCmdUnlockPoll       VialCommand = 0x07
	VialCmdLock             VialCommand = 0x08
	VialCmdQmkSettingsQuery VialCommand = 0x09
	VialCmdQmkSettingsGet   VialCommand = 0x0A
	VialCmdQmkSettingsSet   VialCommand = 0x0B
	VialCmdQmkSettingsReset VialCommand = 0x0C
	VialCmdDynamicEntryOp   VialCommand = 0x0D
	/* operate on tapdance, combos, etc */
)

var (
	// txb         [256]byte // FIXME ... max packet size in descriptors is 32 bytes, why is the buffer 256?
	KeyboardDef []byte // may be preferable to have a callback function to copy def to tx buffer
	// device      KeyMapper // *keyboard.Device

)

// func SetDevice(d KeyMapper) {
// 	device = d
// }

func (host *Host) keyVia(layer, kbIndex, index int) uint16 {
	if kbIndex > 0 { // TODO: support multiple keyboards?
		return 0
	}
	if host == nil || host.km == nil {
		return 0
	}
	kc := uint16(host.km.MapKey(layer, index))
	switch kc {
	default:
		kc = kc & 0x0FFF
	}
	return kc
}

func (host *Host) processPacket(rx []byte, tx []byte) bool {

	device := host.km

	txb := host.txb[:32]
	copy(txb, rx) // FIXME: probably isn't necessary to do this copy

	viaCmd := ViaCommand(rx[0])

	if viaCmd != ViaCmdVialPrefix {
		if debug {
			println("viaCmd:", viaCmd.String())
		}
	}

	switch viaCmd {

	case ViaCmdGetProtocolVersion:
		// println("usb: 0x01 - GetProtocolVersionCount")
		// GetProtocolVersionCount
		txb[2] = 0x09

	case ViaCmdGetKeyboardValue:

	case ViaCmdSetKeyboardValue:

	case ViaCmdDynamicKeymapGetKeycode:

	case ViaCmdDynamicKeymapSetKeycode:
		println("5sb: 0x05 - ", len(rx), rx[1], rx[2], rx[3], rx[4], rx[5])
		//Keys[b[1]][b[2]][b[3]] = Keycode((uint16(b[4]) << 8) + uint16(b[5]))
		// device.SetKeycodeVia(int(b[1]), int(b[2]), int(b[3]), Keycode((uint16(b[4])<<8)+uint16(b[5])))
		// device.flashCh <- true

	case ViaCmdDynamicKeymapReset: // 0x06

	case ViaCmdLightingSetValue: // 0x07

	case ViaCmdLightingGetValue: // 0x08
		txb[1] = 0x00
		txb[2] = 0x00

	case ViaCmdLightingSave: // 0x09

	case ViaCmdEepromReset: // 0x0A

	case ViaCmdBootloaderJump: // 0x0B

	case ViaCmdKeymapMacroGetCount: // 0x0C
		txb[1] = 0x10

	case ViaCmdKeymapMacroGetBufferSize: // 0x0D
		txb[1] = 0x07
		txb[2] = 0x9B

	case ViaCmdKeymapMacroGetBuffer: // 0x0E

	case ViaCmdKeymapMacroSetBuffer: // 0x0F

	case ViaCmdKeymapMacroReset: // 0x10

	case ViaCmdKeymapGetLayerCount: // 0x11
		// println("7sb: 0x11 - DynamicKeymapGetLayerCountCommand")
		// DynamicKeymapGetLayerCountCommand
		if device != nil {
			txb[1] = device.GetLayerCount()
		} else {
			txb[1] = 0x01
		}

	case ViaCmdKeymapGetBuffer: // 0x12
		if device == nil {
			println("warning: device was nil")
			break
		}
		// DynamicKeymapReadBufferCommand
		offset := (uint16(rx[1]) << 8) + uint16(rx[2])
		sz := rx[3]
		cnt := device.GetMaxKeyCount()
		// println("  offset : ", offset, "+", sz, cnt)
		for i := 0; i < int(sz/2); i++ {
			//fmt.Printf("  %02X %02X\n", b[4+i+1], b[4+i+0])
			tmp := i + int(offset)/2
			layer := tmp / (cnt * 1) // len(device.kb))
			tmp = tmp % (cnt * 1)    // len(device.kb))
			kbd := tmp / cnt
			idx := tmp % cnt
			kc := host.keyVia(layer, kbd, idx)
			txb[4+2*i+1] = uint8(kc)
			txb[4+2*i+0] = uint8(kc >> 8)
		}
		// println("done")

	case ViaCmdKeymapSetBuffer: // 0x13

	case ViaCmdVialPrefix: // 0xFE

		vialCmd := VialCommand(rx[1])
		if debug {
			println("vialCmd:", vialCmd.String())
		}

		switch vialCmd {

		case VialCmdGetKeyboardID:
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

		case VialCmdGetSize:
			// println("vial: 0x01 - retrieve keyboard definition size")
			// Retrieve keyboard definition size
			size := len(KeyboardDef)
			txb[0] = uint8(size)
			txb[1] = uint8(size >> 8)
			txb[2] = uint8(size >> 16)
			txb[3] = uint8(size >> 24)

		case VialCmdGetDef:
			// Retrieve 32-bytes block of the definition, page ID encoded within 2 bytes
			page := uint16(rx[2]) + (uint16(rx[3]) << 8)
			start := page * 32
			end := start + 32
			if end < start || int(start) >= len(KeyboardDef) {
				return false
			}
			if int(end) > len(KeyboardDef) {
				end = uint16(len(KeyboardDef))
			}
			copy(txb[:32], KeyboardDef[start:end])

		case VialCmdGetEncoder:
			if em, ok := host.km.(EncoderMapper); ok {
				layer := rx[2]
				idx := rx[3]
				ccw, cw := em.MapEncoder(int(layer), int(idx))
				println("VialCmdGetEncoder: ", layer, idx, ccw, cw)
				tx[0] = 0x0
				tx[1] = byte(ccw)
				tx[2] = 0x0
				tx[3] = byte(cw)
			}
			/*
				case vial_get_encoder: {
						uint8_t layer = msg[2];
						uint8_t idx = msg[3];
						uint16_t keycode = dynamic_keymap_get_encoder(layer, idx, 0);
						msg[0]  = keycode >> 8;
						msg[1]  = keycode & 0xFF;
						keycode = dynamic_keymap_get_encoder(layer, idx, 1);
						msg[2] = keycode >> 8;
						msg[3] = keycode & 0xFF;
						break;
				}
			*/

		case VialCmdSetEncoder:

			if es, ok := host.km.(EncoderSaver); ok {
				var kc uint16
				layer := int(rx[2])
				idx := int(rx[3])
				cw := rx[4] > 0
				kc |= uint16(rx[5]) << 8
				kc |= uint16(rx[6])
				println("VialCmdSetEncoder: ", layer, idx, cw, uint8(kc>>8), uint8(kc))
				if uint8(kc>>8) > 0 {
					// FIXME: multi-byte keycodes not yet supported
					break
				}
				es.SaveEncoder(layer, idx, cw, keycodes.Keycode(kc))
				println(" -- encoder value saved successfully")
			}

			/*
				case vial_set_encoder: {
						dynamic_keymap_set_encoder(msg[2], msg[3], msg[4], vial_keycode_firewall((msg[5] << 8) | msg[6]));
						break;
				}
			*/

		case VialCmdGetUnlockStatus:
			// println("VialCmdGetUnlockStatus")
			txb[0] = 1 // unlocked
			txb[1] = 0 // unlock_in_progress

		case VialCmdUnlockStart:
			// println("VialCmdUnlockStart: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdUnlockPoll:
			// println("VialCmdUnlockPoll: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdLock:
			// println("VialCmdLock: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdQmkSettingsQuery:
			// println("VialCmdQmkSettingsQuery")
			for i := range txb[:32] {
				txb[i] = 0xFF
			}

		case VialCmdQmkSettingsGet:
			// println("VialCmdLock: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdQmkSettingsSet:
			// println("VialCmdQmkSettingsSet: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdQmkSettingsReset:
			// println("VialCmdQmkSettingsReset: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdDynamicEntryOp:
			// println("VialCmdQmkSettingsQuery")
			txb[0] = 0x00
			txb[1] = 0x00
			txb[2] = 0x00

		default:
			println("vial: default - ", rx[1])
		}

	default:
		return false
	}

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
