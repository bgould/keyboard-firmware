// Code generated by "stringer -type=ViaQmkLEDMatrixValue"; DO NOT EDIT.

package vial

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ViaQmkLEDMatrixBrightness-1]
	_ = x[ViaQmkLEDMatrixEffect-2]
	_ = x[ViaQmkLEDMatrixEffectSpeed-3]
}

const _ViaQmkLEDMatrixValue_name = "ViaQmkLEDMatrixBrightnessViaQmkLEDMatrixEffectViaQmkLEDMatrixEffectSpeed"

var _ViaQmkLEDMatrixValue_index = [...]uint8{0, 25, 46, 72}

func (i ViaQmkLEDMatrixValue) String() string {
	i -= 1
	if i >= ViaQmkLEDMatrixValue(len(_ViaQmkLEDMatrixValue_index)-1) {
		return "ViaQmkLEDMatrixValue(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ViaQmkLEDMatrixValue_name[_ViaQmkLEDMatrixValue_index[i]:_ViaQmkLEDMatrixValue_index[i+1]]
}
