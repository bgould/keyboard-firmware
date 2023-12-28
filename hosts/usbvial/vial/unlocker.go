package vial

type Unlocker interface {
	// UnlockStatus() UnlockStatus
	UnlockKeyPos() []Pos
	// StartUnlock()
}
