package usbvial

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

type ViaChannelID uint8

const (
	ViaChannelCustom       ViaChannelID = 0x00
	ViaChannelBacklight                 = 0x01
	ViaChannelRGBLight                  = 0x02
	ViaChannelRGBMatrix                 = 0x03
	ViaChannelQMKAudio                  = 0x04
	ViaChannelQMKLEDMatrix              = 0x05
)
