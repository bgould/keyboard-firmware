package usbvial

import (
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/hsv"
)

func NewKeyboardVialRGBer(kbd *keyboard.Keyboard) vial.VialRGBer {
	return &vialrgbKeyboardDriver{kbd: kbd, speed: 0xFF}
}

type vialrgbKeyboardDriver struct {
	kbd   *keyboard.Keyboard
	speed uint8
}

func (rgb *vialrgbKeyboardDriver) VialRGBMode() uint16 {
	switch rgb.kbd.BacklightMode() {
	case keyboard.BacklightOn:
		return 0x2
	case keyboard.BacklightBreathing:
		return 0x6
	default:
		return 0x0
	}
}

func (rgb *vialrgbKeyboardDriver) VialRGBSpeed() uint8 {
	return rgb.speed
}

func (rgb *vialrgbKeyboardDriver) VialRGBColor() hsv.Color {
	return rgb.kbd.BacklightColor()
}

func (rgb *vialrgbKeyboardDriver) VialRGBUpdate(mode uint16, speed uint8, color hsv.Color) {
	var blmode keyboard.BacklightMode
	switch mode {
	case 0x6:
		blmode = keyboard.BacklightBreathing
	case 0x2:
		blmode = keyboard.BacklightOn
	default:
		blmode = keyboard.BacklightOff
	}
	rgb.speed = speed
	_ = blmode
	// println("skipping update")
	rgb.kbd.BacklightUpdate(blmode, color, false)
}

func (rgb *vialrgbKeyboardDriver) VialRGBSave() {
	rgb.kbd.BacklightSave()
}
