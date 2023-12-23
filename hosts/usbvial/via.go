package usbvial

type ViaChannelID uint8

const (
	ViaChannelCustom       ViaChannelID = 0x00
	ViaChannelBacklight                 = 0x01
	ViaChannelRGBLight                  = 0x02
	ViaChannelRGBMatrix                 = 0x03
	ViaChannelQMKAudio                  = 0x04
	ViaChannelQMKLEDMatrix              = 0x05
)
