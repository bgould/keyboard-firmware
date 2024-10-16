package vial

type rgbImpl interface {
	handleGetValue(rx []byte, tx []byte)
	handleSetValue(rx []byte, tx []byte)
	handleSave(rx []byte, tx []byte)
}

type LightingMode string

const (
	LightingModeVialRGB      LightingMode = "vialrgb"
	LightingModeQMKRGBLight  LightingMode = "qmk_rgblight"
	LightingModeQMKBacklight LightingMode = "qmk_backlight"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ViaQmkRGBLightValue

type ViaQmkRGBLightValue uint8

const (
	ViaQmkRGBLightBrightness ViaQmkRGBLightValue = iota + 0x80
	ViaQmkRGBLightEffect
	ViaQmkRGBLightEffectSpeed
	ViaQmkRGBLightColor
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ViaQmkRGBMatrixValue

type ViaQmkRGBMatrixValue uint8

const (
	ViaQmkRGBMatrixBrightness ViaQmkRGBMatrixValue = iota + 1
	ViaQmkRGBMatrixEffect
	ViaQmkRGBMatrixEffectSpeed
	ViaQmkRGBMatrixColor
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=ViaQmkLEDMatrixValue

type ViaQmkLEDMatrixValue uint8

const (
	ViaQmkLEDMatrixBrightness ViaQmkLEDMatrixValue = iota + 1
	ViaQmkLEDMatrixEffect
	ViaQmkLEDMatrixEffectSpeed
)
