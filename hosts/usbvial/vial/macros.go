package vial

type MacroDriver interface {
	GetMacroCount() uint8
	GetMacroBuffer() []byte
}
