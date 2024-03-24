package keyboard

import (
	// "device/arm"
	"image/color"
)

const (
	colorstrip_no_int = false
)

type ColorStrip struct {
	Writer ColorWriter
	Pixels []color.RGBA
}

type ColorWriter interface {
	WriteColors(values []color.RGBA) error
}

func (ind *ColorStrip) SetPixel(pos int, c color.RGBA) {
	pos, ok := ind.offsetPos(pos)
	if !ok {
		return
	}
	ind.Pixels[pos] = c
}

func (ind *ColorStrip) GetPixel(pos int) (c color.RGBA) {
	pos, ok := ind.offsetPos(pos)
	if !ok {
		return
	}
	return ind.Pixels[pos]
}

func (ind *ColorStrip) SyncPixels() {
	var mask uintptr
	if colorstrip_no_int {
		mask = disableInterrupts()
	}
	ind.Writer.WriteColors(ind.Pixels[:])
	if colorstrip_no_int {
		restoreInterrupts(mask)
	}
}

func (ind *ColorStrip) NumPixels() int {
	return len(ind.Pixels)
}

func (ind *ColorStrip) offsetPos(pos int) (int, bool) {
	if pos >= len(ind.Pixels) {
		return -1, false
	}
	return pos, true
}
