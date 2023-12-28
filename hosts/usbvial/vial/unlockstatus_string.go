// Code generated by "stringer -type=UnlockStatus"; DO NOT EDIT.

package vial

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Locked-0]
	_ = x[Unlocked-1]
	_ = x[UnlockInProgress-2]
}

const _UnlockStatus_name = "LockedUnlockedUnlockInProgress"

var _UnlockStatus_index = [...]uint8{0, 6, 14, 30}

func (i UnlockStatus) String() string {
	if i >= UnlockStatus(len(_UnlockStatus_index)-1) {
		return "UnlockStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _UnlockStatus_name[_UnlockStatus_index[i]:_UnlockStatus_index[i+1]]
}
