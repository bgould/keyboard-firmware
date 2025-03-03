package vial

import (
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	// MagicSerialPrefix is a value in the serial number of HID devices
	// that the	Vial desktop app uses to identify compatible devices.
	MagicSerialPrefix = "vial:f64c2b3c"

	ViaProtocolVersion = 0x09

	VialProtocolVersion = 0x00000006

	VialUnlockCounterMax   = 50
	VialUnlockHoldDuration = 2 * time.Second
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

//go:generate go run golang.org/x/tools/cmd/stringer -type=ViaKeyboardValueID -trimprefix=ViaKbd

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

type ViaChannelID uint8

const (
	ViaChannelCustom       ViaChannelID = 0x00
	ViaChannelBacklight                 = 0x01
	ViaChannelRGBLight                  = 0x02
	ViaChannelRGBMatrix                 = 0x03
	ViaChannelQMKAudio                  = 0x04
	ViaChannelQMKLEDMatrix              = 0x05
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

type DeviceDriver interface {
	GetLayerCount() uint8
	GetMatrixRowState(rowIndex int) uint32
	MapKey(layer, row, col int) keycodes.Keycode
	SetKey(layer, row, col int, kc keycodes.Keycode) bool
}

type Device struct {
	km  DeviceDriver
	txb [32]byte
	def DeviceDefinition
	rgb rgbImpl

	unlockStatus  UnlockStatus
	unlockCounter int
	unlockStart   time.Time
}

type Pos struct {
	Row uint8
	Col uint8
}

type DeviceDefinition struct {
	Name          string       `json:"name"`
	VendorID      string       `json:"vendorId"`
	ProductID     string       `json:"productId"`
	Matrix        DeviceMatrix `json:"matrix"`
	Lighting      string       `json:"lighting"`
	UnlockKeys    []Pos        `json:"-"`
	LzmaDefLength uint16       `json:"-"`
	LzmaDefWriter LzmaDefPageWriter
	rgb           rgbImpl
}

type LzmaDefPageWriter interface {
	WriteLzmaDefPage(buf []byte, page uint16) bool
}

type LzmaDefPageWriterFunc func(buf []byte, page uint16) bool

func (fn LzmaDefPageWriterFunc) WriteLzmaDefPage(buf []byte, page uint16) bool {
	return fn(buf, page)
}

type DeviceMatrix struct {
	Rows int `json:"rows"`
	Cols int `json:"cols"`
}

func NewDevice(def DeviceDefinition, driver DeviceDriver) *Device {
	unlockKeys := def.UnlockKeys
	if len(unlockKeys) == 0 {
		def.UnlockKeys = []Pos{{0, 0}}
	}
	return &Device{km: driver, def: def}
}

func (dev *Device) UnlockStatus() UnlockStatus {
	return dev.unlockStatus
}

func (dev *Device) keyVia(layer, kbIndex, idx int) uint16 {
	if kbIndex > 0 { // TODO: support multiple keyboards?
		return 0
	}
	if dev == nil || dev.km == nil {
		return 0
	}
	numCols := dev.def.Matrix.Cols //dev.km.NumCols()
	row := idx / numCols
	col := idx % numCols
	kc := uint16(dev.km.MapKey(layer, row, col))
	return kc
}

func (dev *Device) Handle(rx []byte, tx []byte) bool {

	mapper := dev.km

	// txb := host.txb[:32]
	// copy(txb, rx) // FIXME: probably isn't necessary to do this copy

	// ensure tx buffer is zero'd
	for i := range tx {
		tx[i] = 0
	}

	viaCmd := ViaCommand(rx[0])

	if viaCmd != ViaCmdVialPrefix {
		if debug {
			println("viaCmd:", viaCmd) //)strconv.FormatUint(uint64(viaCmd), 16))
		}
	}

	// println(rx[0], rx[1])

	switch viaCmd {

	case ViaCmdGetProtocolVersion: // 0x01
		tx[0] = rx[0]
		tx[1] = rx[1]
		tx[2] = ViaProtocolVersion

	case ViaCmdGetKeyboardValue: // 0x02

		if debug {
			println("getting keyboard value:", rx[1])
		}

		switch ViaKeyboardValueID(rx[1]) {
		case ViaKbdUptime:
			if debug {
				println("uptime not yet implemented")
			}

		case ViaKbdLayoutOptions:
			if debug {
				println("layout options not yet implemented")
			}

		case ViaKbdSwitchMatrixState:
			// matrix, ok := dev.km.(Matrixer)
			if debug {
				println("get keyboard value: switch matrix state", dev.unlockStatus)
			}
			tx[0] = rx[0]
			tx[1] = rx[1]
			// TODO: #if ((MATRIX_COLS / 8 + 1) * MATRIX_ROWS <= 28)
			if dev.UnlockStatus() == Unlocked {
				nRows, nCols, i := dev.def.Matrix.Rows, dev.def.Matrix.Cols, 2
				for rowIndex := 0; rowIndex < nRows; rowIndex++ {
					value := dev.km.GetMatrixRowState(rowIndex)
					switch {
					case nCols > 24:
						tx[i] = uint8(value >> 24)
						i++
						fallthrough
					case nCols > 16:
						tx[i] = uint8(value >> 16)
						i++
						fallthrough
					case nCols > 8:
						tx[i] = uint8(value >> 8)
						i++
						fallthrough
					default:
						tx[i] = uint8(value)
						i++
					}
					if debug {
						println("matrix state", rowIndex, ":", tx[i-4], tx[i-3], tx[i-2], tx[i-1])
					}
				}
			}
			// TODO: #endif

		case ViaKbdFirmwareVersion:
			if debug {
				println("firmware version not yet implemented")
			}

		case ViaKbdDeviceIndication:
			if debug {
				println("device indication not yet implemented")
			}

		}

	case ViaCmdSetKeyboardValue: // 0x03

	case ViaCmdDynamicKeymapGetKeycode: // 0x04

	case ViaCmdDynamicKeymapSetKeycode: // 0x05
		if debug {
			println("ViaCmdDynamicKeymapSetKeycode: ", rx[1], rx[2], rx[3], rx[4], rx[5])
		}
		// if setter, ok := mapper.(KeySetter); ok {
		layer := int(rx[1])
		row := int(rx[2])
		col := int(rx[3])
		kc := keycodes.Keycode(uint16(rx[4])<<8 | uint16(rx[5]))
		result := dev.km.SetKey(layer, row, col, kc)
		if debug {
			println("-- set keycode result: ", kc, result)
		}
		// }

	case ViaCmdDynamicKeymapReset: // 0x06

	case ViaCmdLightingSetValue: // 0x07
		if debug {
			println("ViaCmdLightingSetValue:", rx[1], rx[2], rx[3])
		}
		if dev.def.rgb != nil {
			dev.def.rgb.handleSetValue(rx, tx)
		}

	case ViaCmdLightingGetValue: // 0x08
		if debug {
			println("ViaCmdLightingGetValue:", VialRGBGetCommand(rx[1]), rx[2], rx[3], rx[4], rx[5])
		}
		if dev.def.rgb != nil {
			dev.def.rgb.handleGetValue(rx, tx)
		}

	case ViaCmdLightingSave: // 0x09
		if debug {
			println("ViaCmdLightingSave:", rx[1], rx[2], rx[3], rx[4], rx[5])
		}
		if dev.def.rgb != nil {
			dev.def.rgb.handleSave(rx, tx)
		}

	case ViaCmdEepromReset: // 0x0A

	case ViaCmdBootloaderJump: // 0x0B

	case ViaCmdKeymapMacroGetCount: // 0x0C
		if debug {
			println("ViaCmdKeymapMacroGetCount")
		}
		tx[0] = rx[0]
		if drv, ok := dev.km.(MacroDriver); ok {
			tx[1] = drv.GetMacroCount()
		} else {
			tx[1] = 0x0
		}

	case ViaCmdKeymapMacroGetBufferSize: // 0x0D
		if debug {
			println("ViaCmdKeymapMacroGetBufferSize")
		}
		tx[0] = rx[0]
		if drv, ok := dev.km.(MacroDriver); ok {
			bufsize := uint16(len(drv.GetMacroBuffer()))
			tx[1] = byte(bufsize >> 8)
			tx[2] = byte(bufsize)
		} else {
			tx[1] = 0x00
			tx[2] = 0x00
		}

	case ViaCmdKeymapMacroGetBuffer: // 0x0E
		if debug {
			println("ViaCmdKeymapMacroGetBuffer:", rx[1], rx[2], rx[3])
		}
		tx[0] = rx[0]
		tx[1] = rx[1]
		tx[2] = rx[2]
		tx[3] = rx[3]
		if drv, ok := dev.km.(MacroDriver); ok {
			offset := (uint16(rx[1]) << 8) + uint16(rx[2])
			size := rx[3]
			buf := drv.GetMacroBuffer()
			copy(tx[4:4+size], buf[offset:])
		}

	case ViaCmdKeymapMacroSetBuffer: // 0x0F
		if debug {
			println("ViaCmdKeymapMacroSetBuffer", rx[1], rx[2], rx[3])
		}
		tx[0] = rx[0]
		tx[1] = rx[1]
		tx[2] = rx[2]
		tx[3] = rx[3]
		if drv, ok := dev.km.(MacroDriver); ok {
			offset := (uint16(rx[1]) << 8) + uint16(rx[2])
			size := rx[3]
			buf := drv.GetMacroBuffer()
			copy(buf[offset:], rx[4:4+size])
		}

	case ViaCmdKeymapMacroReset: // 0x10
		if debug {
			println("ViaCmdKeymapMacroReset")
		}

	case ViaCmdKeymapGetLayerCount: // 0x11
		tx[1] = mapper.GetLayerCount()

	case ViaCmdKeymapGetBuffer: // 0x12
		offset := (uint16(rx[1]) << 8) + uint16(rx[2])
		sz := rx[3]
		cnt := dev.def.Matrix.Rows * dev.def.Matrix.Cols //mapper.GetMaxKeyCount()
		for i := 0; i < int(sz/2); i++ {
			tmp := i + int(offset)/2
			layer := tmp / (cnt * 1) // len(device.kb)) // TODO: support multiple "keyboards"?
			tmp = tmp % (cnt * 1)    // len(device.kb))  // TODO: support multiple "keyboards"?
			kbd := tmp / cnt
			idx := tmp % cnt
			kc := dev.keyVia(layer, kbd, idx)
			tx[4+2*i+1] = uint8(kc)
			tx[4+2*i+0] = uint8(kc >> 8)
		}

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
			tx[0] = VialProtocolVersion
			tx[1] = VialProtocolVersion >> 8
			tx[2] = VialProtocolVersion >> 16
			tx[3] = VialProtocolVersion >> 24
			tx[4] = 0x9D  // TODO
			tx[5] = 0xD0  // TODO
			tx[6] = 0xD5  // TODO
			tx[7] = 0xE1  // TODO
			tx[8] = 0x87  // TODO
			tx[9] = 0xF3  // TODO
			tx[10] = 0x54 // TODO
			tx[11] = 0xE2 // TODO

		case VialCmdGetSize:
			// Retrieve keyboard definition size
			size := dev.def.LzmaDefLength
			tx[0] = uint8(size)
			tx[1] = uint8(size >> 8)
			tx[2] = 0x0 // uint8(size >> 16) // size is uint16
			tx[3] = 0x0 // uint8(size >> 24) // size is uint16

		case VialCmdGetDef:
			// Retrieve 32-bytes block of the definition, page ID encoded within 2 bytes
			page := uint16(rx[2]) + (uint16(rx[3]) << 8)
			if !dev.def.LzmaDefWriter.WriteLzmaDefPage(tx[:32], page) {
				return false // TODO: error handling
			}

		case VialCmdGetEncoder:
			if em, ok := dev.km.(EncoderMapper); ok {
				layer := rx[2]
				idx := rx[3]
				ccwRow, ccwCol, cwRow, cwCol, ok := em.MapEncoder(int(idx))
				if !ok {
					tx[0] = 0x0
					tx[1] = 0x0
					tx[2] = 0x0
					tx[3] = 0x0
				} else {
					ccw := dev.km.MapKey(int(layer), ccwRow, ccwCol)
					cw := dev.km.MapKey(int(layer), cwRow, cwCol)
					tx[0] = byte(ccw >> 8)
					tx[1] = byte(ccw)
					tx[2] = byte(cw >> 8)
					tx[3] = byte(cw)
				}
			}

		case VialCmdSetEncoder:
			em, emOk := dev.km.(EncoderMapper)
			if emOk {
				var kc uint16
				layer := int(rx[2])
				index := int(rx[3])
				cw := rx[4] > 0
				kc |= uint16(rx[5]) << 8
				kc |= uint16(rx[6])
				if debug {
					println("VialCmdSetEncoder: ", layer, index, cw, uint8(kc>>8), uint8(kc))
				}
				var row, col int
				if ccwRow, ccwCol, cwRow, cwCol, ok := em.MapEncoder(index); ok {
					if cw {
						row = cwRow
						col = cwCol
					} else {
						row = ccwRow
						col = ccwCol
					}
				}
				if ok := dev.km.SetKey(layer, row, col, keycodes.Keycode(kc)); ok {
					if debug {
						println(" -- encoder value saved successfully")
					}
				} else {
					if debug {
						println(" -- encoder value not saved")
					}
				}
			}

		case VialCmdGetUnlockStatus:
			// TODO: implement equivalent of VIAL_INSECURE
			switch dev.UnlockStatus() {
			case Locked:
				tx[0] = byte(Locked) // locked
				tx[1] = 0            // unlock NOT in progress
			case Unlocked:
				tx[0] = byte(Unlocked) // unlocked
				tx[1] = 0              // unlock NOT in progress
			case UnlockInProgress:
				tx[0] = byte(Locked) // locked
				tx[1] = 1            // unlock in progress
			}
			unlockKeys := dev.def.UnlockKeys
			for i, pos := range unlockKeys {
				tx[2+i*2] = pos.Row
				tx[3+i*2] = pos.Col
			}
			// println("get unlock status:", tx[0])

		case VialCmdUnlockStart:
			if debug {
				println("unlock start")
			}
			dev.unlockStatus = UnlockInProgress
			dev.unlockStart = time.Now()
			// TODO: implement equivalent of VIAL_INSECURE
			// if dev.unlocker != nil {
			// 	dev.unlocker.StartUnlock()
			// }
			// dev.unlockStatus = UnlockInProgress
			// dev.unlockCounter = VialUnlockCounterMax
			// dev.unlockStarted = time.Now()
			// println("VialCmdUnlockStart: ", rx[0], rx[1], rx[2], rx[3], rx[4], rx[5], rx[6], rx[7], rx[8])

		case VialCmdUnlockPoll:
			if debug {
				println("unlock poll")
			}
			var dur = time.Since(dev.unlockStart)
			if dev.UnlockStatus() == UnlockInProgress {
				holding := true
				for _, pos := range dev.def.UnlockKeys {
					if holding {
						rowState := dev.km.GetMatrixRowState(int(pos.Row))
						holding = (rowState & (1 << pos.Col)) > 0
						if debug {
							println("rowState: ", holding)
						}
					}
				}
				if !holding {
					dev.unlockStart = time.Now()
					dur = 0
				}
				// println("unlock in progress: ", time.Since(dev.unlockStart)/time.Millisecond, holding)
				if holding && dur > VialUnlockHoldDuration {
					// println("unlocked!")
					dev.unlockStatus = Unlocked
					dev.unlockStart = time.Time{}
					dur = 0
				}
			}
			switch dev.UnlockStatus() {
			case Locked:
				tx[0] = byte(Locked) // locked
				tx[1] = 0            // unlock NOT in progress
				tx[2] = 0
			case Unlocked:
				tx[0] = byte(Unlocked) // unlocked
				tx[1] = 0              // unlock NOT in progress
				tx[2] = VialUnlockCounterMax
			case UnlockInProgress:
				tx[0] = byte(Locked) // locked
				tx[1] = 1            // unlock in progress
				tx[2] = VialUnlockCounterMax - byte((float32(dur)/float32(VialUnlockHoldDuration))*VialUnlockCounterMax)
			}
			if debug {
				println("VialCmdUnlockPoll: ", tx[0], tx[1], VialUnlockCounterMax-tx[2])
			}

		case VialCmdLock:
			dev.unlockStatus = Locked
			dev.unlockStart = time.Time{}
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
		if handler, ok := dev.km.(Handler); ok {
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
