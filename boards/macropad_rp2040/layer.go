package macropad_rp2040

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func Layer(
	/**/ k13, k14,
	/*     */ k00,
	k01, k02, k03,
	k04, k05, k06,
	k07, k08, k09,
	k10, k11, k12 keycodes.Keycode,
) keyboard.Layer {
	return keyboard.Layer([][]keycodes.Keycode{
		{
			k00, k01, k02, k03, k04, k05, k06, k07,
			k08, k09, k10, k11, k12, k13, k14, 0x0,
		},
	})
}
