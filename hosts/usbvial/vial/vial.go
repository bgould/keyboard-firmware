package vial

import (
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	// MagicSerialPrefix is a value in the serial number of HID devices
	// that the	Vial desktop app uses to identify compatible devices.
	MagicSerialPrefix = "vial:f64c2b3c"

	VialProtocolVersion = 0x09
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

//go:generate go run golang.org/x/tools/cmd/stringer -type=ViaKeyboardValueID

const (
	ViaKbdUptime            ViaKeyboardValueID = 0x01
	ViaKbdLayoutOptions     ViaKeyboardValueID = 0x02
	ViaKbdSwitchMatrixState ViaKeyboardValueID = 0x03
	ViaKbdFirmwareVersion   ViaKeyboardValueID = 0x04
	ViaKbdDeviceIndication  ViaKeyboardValueID = 0x05
)

type VialCommand uint8

//go:generate go run golang.org/x/tools/cmd/stringer -type=VialCommand

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
)

type UnlockStatus uint8

//go:generate go run golang.org/x/tools/cmd/stringer -type=UnlockStatus

const (
	Locked UnlockStatus = iota
	Unlocked
	UnlockInProgress
)

// the key override structure as it is stored in eeprom and transferred to vial-gui;
// it is deserialized into key_override_t by vial_get_key_override
type KeyOverrideEntry struct {
	Trigger         uint16
	Replacement     uint16
	Layers          uint16
	TriggerMods     uint8
	NegativeModMask uint8
	SupressedMods   uint8
	Options         uint8
}

var (
	// txb         [256]byte // FIXME ... max packet size in descriptors is 32 bytes, why is the buffer 256?
	KeyboardDef []byte // may be preferable to have a callback function to copy def to tx buffer
	// device      KeyMapper // *keyboard.Device

)

type Device struct {
	km  KeyMapper
	txb [32]byte

	unlocked bool
}

func NewDevice(mapper KeyMapper) *Device {
	return &Device{km: mapper}
}

func (h *Device) Unlocked() bool {
	return h.unlocked
}

// func SetDevice(d KeyMapper) {
// 	device = d
// }

func (host *Device) keyVia(layer, kbIndex, idx int) uint16 {
	if kbIndex > 0 { // TODO: support multiple keyboards?
		return 0
	}
	if host == nil || host.km == nil {
		return 0
	}
	numCols := host.km.NumCols()
	row := idx / numCols
	col := idx % numCols
	kc := uint16(host.km.MapKey(layer, row, col))
	switch kc {
	default:
		kc = kc & 0x0FFF
	}
	return kc
}

func (host *Device) Handle(rx []byte, tx []byte) bool {

	device := host.km

	// txb := host.txb[:32]
	// copy(txb, rx) // FIXME: probably isn't necessary to do this copy

	viaCmd := ViaCommand(rx[0])

	if viaCmd != ViaCmdVialPrefix {
		if debug {
			println("viaCmd:", strconv.FormatUint(uint64(viaCmd), 64))
		}
	}

	// println(rx[0], rx[1])

	switch viaCmd {

	case ViaCmdGetProtocolVersion: // 0x01
		tx[0] = rx[0]
		tx[1] = rx[1]
		tx[2] = VialProtocolVersion

	case ViaCmdGetKeyboardValue: // 0x02

	case ViaCmdSetKeyboardValue: // 0x03

	case ViaCmdDynamicKeymapGetKeycode: // 0x04

	case ViaCmdDynamicKeymapSetKeycode: // 0x05
		if debug {
			println("ViaCmdDynamicKeymapSetKeycode: ", rx[1], rx[2], rx[3], rx[4], rx[5])
		}
		if setter, ok := host.km.(KeySetter); ok {
			layer := int(rx[1])
			row := int(rx[2])
			col := int(rx[3])
			kc := keycodes.Keycode(uint16(rx[4])>>8 | uint16(rx[5]))
			// entry := KeyOverrideEntry{
			// 	Trigger:         uint16(rx[4])>>8 | uint16(rx[5]),
			// 	Replacement:     uint16(rx[6])>>8 | uint16(rx[7]),
			// 	Layers:          uint16(rx[8])>>8 | uint16(rx[9]),
			// 	TriggerMods:     rx[10],
			// 	NegativeModMask: rx[11],
			// 	SupressedMods:   rx[12],
			// 	Options:         rx[13],
			// }
			result := setter.SetKey(layer, row, col, kc)
			if debug {
				println("-- set keycode result: ", result)
			}
		}

	case ViaCmdDynamicKeymapReset: // 0x06

	case ViaCmdLightingSetValue: // 0x07

	case ViaCmdLightingGetValue: // 0x08
		tx[1] = 0x00
		tx[2] = 0x00

	case ViaCmdLightingSave: // 0x09

	case ViaCmdEepromReset: // 0x0A

	case ViaCmdBootloaderJump: // 0x0B

	case ViaCmdKeymapMacroGetCount: // 0x0C
		tx[1] = 0x10

	case ViaCmdKeymapMacroGetBufferSize: // 0x0D
		tx[1] = 0x07
		tx[2] = 0x9B

	case ViaCmdKeymapMacroGetBuffer: // 0x0E

	case ViaCmdKeymapMacroSetBuffer: // 0x0F

	case ViaCmdKeymapMacroReset: // 0x10

	case ViaCmdKeymapGetLayerCount: // 0x11
		if device != nil {
			tx[1] = device.GetLayerCount()
		} else {
			tx[1] = 0x01
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
			tx[4+2*i+1] = uint8(kc)
			tx[4+2*i+0] = uint8(kc >> 8)
		}
		// println("done")

	case ViaCmdKeymapSetBuffer: // 0x13

	case ViaCmdVialPrefix: // 0xFE

		vialCmd := VialCommand(rx[1])
		if debug {
			println("vialCmd:", vialCmd)
		}

		switch vialCmd {

		case VialCmdGetKeyboardID:
			// println("vial: 0x00 - Get keyboard ID and Vial protocol version")
			// Get keyboard ID and Vial protocol version
			const vialProtocolVersion = 0x00000006
			tx[0] = vialProtocolVersion
			tx[1] = vialProtocolVersion >> 8
			tx[2] = vialProtocolVersion >> 16
			tx[3] = vialProtocolVersion >> 24
			tx[4] = 0x9D
			tx[5] = 0xD0
			tx[6] = 0xD5
			tx[7] = 0xE1
			tx[8] = 0x87
			tx[9] = 0xF3
			tx[10] = 0x54
			tx[11] = 0xE2

		case VialCmdGetSize:
			// println("vial: 0x01 - retrieve keyboard definition size")
			// Retrieve keyboard definition size
			size := len(KeyboardDef)
			tx[0] = uint8(size)
			tx[1] = uint8(size >> 8)
			tx[2] = uint8(size >> 16)
			tx[3] = uint8(size >> 24)

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
			copy(tx[:32], KeyboardDef[start:end])

		case VialCmdGetEncoder:
			if em, ok := host.km.(EncoderMapper); ok {
				layer := rx[2]
				idx := rx[3]
				ccw, cw := em.MapEncoder(int(layer), int(idx))
				// println("VialCmdGetEncoder: ", layer, idx, ccw, cw)
				tx[0] = byte(ccw >> 8)
				tx[1] = byte(ccw)
				tx[2] = byte(cw >> 8)
				tx[3] = byte(cw)
			}

		case VialCmdSetEncoder:
			if es, ok := host.km.(EncoderSaver); ok {
				var kc uint16
				layer := int(rx[2])
				index := int(rx[3])
				cw := rx[4] > 0
				kc |= uint16(rx[5]) << 8
				kc |= uint16(rx[6])
				if debug {
					println("VialCmdSetEncoder: ", layer, index, cw, uint8(kc>>8), uint8(kc))
				}
				if uint8(kc>>8) > 0 {
					// FIXME: multi-byte keycodes not yet supported
					break
				}
				es.SaveEncoder(layer, index, cw, keycodes.Keycode(kc))
				if debug {
					println(" -- encoder value saved successfully")
				}
			}

		case VialCmdGetUnlockStatus:
			// println("VialCmdGetUnlockStatus")
			tx[0] = 1 // unlocked
			tx[1] = 0 // unlock_in_progress

		case VialCmdUnlockStart:
			// println("VialCmdUnlockStart: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdUnlockPoll:
			// println("VialCmdUnlockPoll: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdLock:
			// println("VialCmdLock: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdQmkSettingsQuery:
			// println("VialCmdQmkSettingsQuery")
			for i := range tx[:32] {
				tx[i] = 0xFF
			}

		case VialCmdQmkSettingsGet:
			// println("VialCmdLock: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdQmkSettingsSet:
			// println("VialCmdQmkSettingsSet: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdQmkSettingsReset:
			// println("VialCmdQmkSettingsReset: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdDynamicEntryOp:
			// println("VialCmdQmkSettingsQuery")
			tx[0] = 0x00
			tx[1] = 0x00
			tx[2] = 0x00

		default:
			if debug {
				println("vial: default - ", rx[1])
			}

		}

	default:
		if debug {
			println("vial default - ", rx[0])
		}
		if handler, ok := host.km.(Handler); ok {
			return handler.Handle(rx, tx)
		}
		return false
	}

	return true
}

// TODO: determine correct logic for this function, or if it is even necessary
// func (h *Device) keycodeFirewall(kc keycodes.Keycode) keycodes.Keycode {
// 	if kc == keycodes.PROG && !h.Unlocked() {
// 		return 0
// 	}
// 	return kc
// }

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