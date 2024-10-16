package vial

import "github.com/bgould/keyboard-firmware/keyboard/hsv"

type VialRGBer interface {
	VialRGBMode() uint16
	VialRGBSpeed() uint8
	VialRGBColor() hsv.Color
	VialRGBUpdate(mode uint16, speed uint8, color hsv.Color)
	VialRGBSave()
}

func (dev *Device) UseVialRGB(drv VialRGBer) {
	dev.def.UseVialRGB(drv)
}

func (def *DeviceDefinition) UseVialRGB(drv VialRGBer) {
	def.rgb = &vialrgbImpl{drv: drv}
}

const (
	VialRGBProtocolVersion = 1
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=VialRGBSetCommand

type VialRGBSetCommand uint8

const (
	VialRGBCommandSetMode VialRGBSetCommand = iota + 0x41
	VialRGBCommandDirectFastset
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=VialRGBGetCommand

type VialRGBGetCommand uint8

const (
	VialRGBCommandGetInfo VialRGBGetCommand = iota + 0x40
	VialRGBCommandGetMode
	VialRGBCommandGetSupported
	VialRGBCommandGetNumberLEDs
	VialRGBCommandGetLEDInfo
)

type vialrgbImpl struct {
	drv VialRGBer
}

var _ rgbImpl = (*vialrgbImpl)(nil)

func (rgb *vialrgbImpl) handleGetValue(rx []byte, tx []byte) {
	cmd := VialRGBGetCommand(rx[1])
	args := rx[2:]
	tx[0] = rx[0]
	tx[1] = rx[1]
	resp := tx[2:]
	switch cmd {
	case VialRGBCommandGetInfo:
		if debug {
			println("VialRGBCommandGetInfo")
		}
		resp[0] = VialRGBProtocolVersion & 0xFF
		resp[1] = VialRGBProtocolVersion >> 8
		resp[2] = 0xFF // RGB_MATRIX_MAXIMUM_BRIGHTNESS
		if debug {
			println("VialRGBCommandGetInfo", tx[2], tx[3], tx[4])
		}
	case VialRGBCommandGetMode:
		if debug {
			println("VialRGBCommandGetMode")
		}
		mode := rgb.drv.VialRGBMode()
		speed := rgb.drv.VialRGBSpeed()
		color := rgb.drv.VialRGBColor()
		resp[0] = uint8(mode)
		resp[1] = uint8(mode >> 8)
		resp[2] = speed
		resp[3] = color.H
		resp[4] = color.S
		resp[5] = color.V
	case VialRGBCommandGetSupported:
		if debug {
			println("VialRGBCommandGetSupported", args[0], args[1], args[2], args[3])
		}
		// FIXME: actually have map of modes similar to vialrgb_effects.inc
		if args[0] == 0 {
			resp[0] = 0x0 // VIALRGB_EFFECT_OFF
			resp[1] = 0x0
			resp[2] = 0x1 // VIALRGB_EFFECT_DIRECT
			resp[3] = 0x0
			resp[4] = 0x2 // VIALRGB_EFFECT_SOLID_COLOR
			resp[5] = 0x0
			resp[6] = 0x6 // VIALRGB_EFFECT_BREATHING
			resp[7] = 0x0
		} else {
			resp[0] = 0xFF
			resp[1] = 0xFF
		}
	}
}

func (rgb *vialrgbImpl) handleSetValue(rx []byte, tx []byte) {
	cmd := VialRGBSetCommand(rx[1])
	args := rx[2:]
	tx[0] = rx[0]
	tx[1] = rx[1]
	// resp := tx[2:]
	switch cmd {
	case VialRGBCommandSetMode:
		if debug {
			println("VialRGBCommandSetMode", args[0], args[1], args[2], args[3], args[4], args[5])
		}
		mode := uint16(args[0]) | uint16(args[1]<<8)
		speed := args[2]
		color := hsv.Color{H: args[3], S: args[4], V: args[5]}
		rgb.drv.VialRGBUpdate(mode, speed, color)
	case VialRGBCommandDirectFastset:
		if debug {
			println("VialRGBCommandDirectFastset", args[0], args[1], args[2], args[3], args[4], args[5])
		}
	}
}

func (rgb *vialrgbImpl) handleSave(rx []byte, tx []byte) {
	if debug {
		println("VialRGBSave")
	}
	rgb.drv.VialRGBSave()
}
