package keyboard

import (
	"image/color"
	"sync"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/hsv"
)

const (
	backlight_debug = false
)

type BacklightColorStrip struct {

	//
	ColorStrip ColorPixeler

	//
	Interval time.Duration

	// DoubleSyncHack is a temporary workaround for Macropad RP2040
	DoubleSyncHack bool

	mutex  sync.Mutex
	cancel func()

	state backlightState

	brightening bool
	step        uint8
	last        time.Time

	// channelLED uint8
}

type ColorPixeler interface {
	NumPixels() int
	GetPixel(pos int) (c color.RGBA)
	SetPixel(pos int, c color.RGBA)
	SyncPixels()
}

var _ BacklightDriver = (*BacklightColorStrip)(nil)

func (bl *BacklightColorStrip) Configure() {
	if bl.Interval == 0 {
		bl.Interval = 4 * time.Millisecond
	}
	for i, c := 0, bl.ColorStrip.NumPixels(); i < c; i++ {
		bl.ColorStrip.SetPixel(i, color.RGBA{})
	}
	bl.ColorStrip.SyncPixels()
}

func (bl *BacklightColorStrip) Task() {

	if time.Since(bl.last) < bl.Interval {
		return
	}

	bl.mutex.Lock()
	defer bl.mutex.Unlock()

	if bl.state.mode != BacklightBreathing {
		return
	}

	if bl.step == 0 {
		bl.brightening = !bl.brightening
	}

	brightness := bl.step
	if !bl.brightening {
		brightness = 255 - brightness
	}
	bl.set(uint8(brightness), false)

	bl.last = bl.last.Add(bl.Interval)
	bl.step++
}

func (bl *BacklightColorStrip) SetBacklight(mode BacklightMode, color hsv.Color) {

	bl.mutex.Lock()
	defer bl.mutex.Unlock()

	if mode == bl.state.mode && color == bl.state.color {
		return
	}

	bl.state.mode, bl.state.color = mode, color
	if backlight_debug {
		println("SetBacklight(): ", bl.state.mode, bl.state.color.String())
	}

	switch bl.state.mode {

	case BacklightOff:
		if backlight_debug {
			println("BacklightOff")
		}
		bl.cancelIfRunning()
		bl.set(0, backlight_debug)
		if bl.DoubleSyncHack {
			time.Sleep(500 * time.Microsecond)
			bl.set(0, backlight_debug)
		}

	case BacklightOn:
		if backlight_debug {
			println("BacklightOn")
		}
		bl.cancelIfRunning()
		bl.set(uint8(bl.state.color.V), backlight_debug)
		if bl.DoubleSyncHack {
			time.Sleep(500 * time.Microsecond)
			bl.set(uint8(bl.state.color.V), backlight_debug)
		}

	case BacklightBreathing:
		if backlight_debug {
			println("BacklightBreathing")
		}
		bl.cancelIfRunning()
		bl.brightening = false
		bl.step = 0xF
		bl.last = time.Now()

	}
}

func (bl *BacklightColorStrip) cancelIfRunning() {
	if bl.cancel != nil {
		bl.cancel()
		bl.cancel = nil
	}
}

func (bl *BacklightColorStrip) set(val uint8, debug bool) {
	hsvColor := bl.state.color
	hsvColor.V = val
	r, g, b := hsvColor.ConvertToRGB()
	// r, g, b := bl.state.color.H, bl.state.color.S, bl.state.color.V
	// println("backlight RGB", r, g, b)
	col := color.RGBA{R: r, G: g, B: b, A: 0xFF}
	for i, c := 0, bl.ColorStrip.NumPixels(); i < c; i++ {
		bl.ColorStrip.SetPixel(i, col)
		if debug {
			println("set pixel", i, "to", col.R, col.G, col.B, col.A, "for val", val)
		}
	}
	bl.ColorStrip.SyncPixels()
}
