// Code generated by "stringer -type=MacroCode"; DO NOT EDIT.

package keyboard

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MacroCodeNone-0]
	_ = x[MacroCodeTap-1]
	_ = x[MacroCodeDown-2]
	_ = x[MacroCodeUp-3]
	_ = x[MacroCodeDelay-4]
	_ = x[MacroCodeVialExtTap-5]
	_ = x[MacroCodeVialExtDown-6]
	_ = x[MacroCodeVialExtUp-7]
	_ = x[MacroCodeSend-8]
}

const _MacroCode_name = "MacroCodeNoneMacroCodeTapMacroCodeDownMacroCodeUpMacroCodeDelayMacroCodeVialExtTapMacroCodeVialExtDownMacroCodeVialExtUpMacroCodeSend"

var _MacroCode_index = [...]uint8{0, 13, 25, 38, 49, 63, 82, 102, 120, 133}

func (i MacroCode) String() string {
	if i >= MacroCode(len(_MacroCode_index)-1) {
		return "MacroCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MacroCode_name[_MacroCode_index[i]:_MacroCode_index[i+1]]
}
