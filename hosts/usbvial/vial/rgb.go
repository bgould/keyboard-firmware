package vial

type rgbImpl interface {
	handleGetValue(rx []byte, tx []byte)
	handleSetValue(rx []byte, tx []byte)
	handleSave(rx []byte, tx []byte)
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=VialRGBSetCommand

type LightingMode string

const (
	LightingModeVialRGB      LightingMode = "vialrgb"
	LightingModeQMKRGBLight  LightingMode = "qmk_rgblight"
	LightingModeQMKBacklight LightingMode = "qmk_backlight"
)

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

/* Start at 0x40 in order to not conflict with existing "enum via_lighting_value",
   even though they likely wouldn't be enabled together with vialrgb
enum {
    vialrgb_set_mode = 0x41,
    vialrgb_direct_fastset = 0x42,
};

enum {
    vialrgb_get_info = 0x40,
    vialrgb_get_mode = 0x41,
    vialrgb_get_supported = 0x42,
    vialrgb_get_number_leds = 0x43,
    vialrgb_get_led_info = 0x44,
};

void vialrgb_get_value(uint8_t *data, uint8_t length);
void vialrgb_set_value(uint8_t *data, uint8_t length);
void vialrgb_save(uint8_t *data, uint8_t length);

#if defined(VIALRGB_ENABLE) && !defined(RGB_MATRIX_ENABLE)
#error VIALRGB_ENABLE=yes requires RGB_MATRIX_ENABLE=yes
#endif
*/
