package hsv

import "image/color"

type Color struct {
	H, S, V uint8
}

func (c *Color) RGBA() (r, g, b, a uint32) {
	r_, g_, b_ := c.ConvertToRGB()
	return color.RGBA{R: r_, G: g_, B: b_, A: 0x0}.RGBA()
}

func (hsv *Color) ConvertToRGB() (r, g, b uint8) {
	var p, q, t uint8
	var h, s, v, region, remainder uint16

	if hsv.S == 0 {
		r = hsv.V
		g = hsv.V
		b = hsv.V
		return
	}

	h = uint16(hsv.H)
	s = uint16(hsv.S)
	v = uint16(hsv.V)

	region = h * 6 / 255
	remainder = (h*2 - uint16(region*85)) * 3

	p = uint8((v * (255 - s)) >> 8)
	q = uint8((v * (255 - ((s * remainder) >> 8))) >> 8)
	t = uint8((v * (255 - ((s * (255 - remainder)) >> 8))) >> 8)

	switch region {
	case 6:
	case 0:
		r = uint8(v)
		g = t
		b = p
	case 1:
		r = q
		g = uint8(v)
		b = p
	case 2:
		r = p
		g = uint8(v)
		b = t
	case 3:
		r = p
		g = q
		b = uint8(v)
	case 4:
		r = t
		g = p
		b = uint8(v)
	default:
		r = uint8(v)
		g = p
		b = q
	}
	return
}

var _ color.Color = (*Color)(nil)

func (h *Color) HueInc(step uint8) (newval uint8, changed bool) {
	h.H += step
	changed = true
	newval = h.H
	return
}

func (h *Color) HueDec(step uint8) (newval uint8, changed bool) {
	h.H -= step
	changed = true
	newval = h.H
	return
}

func (h *Color) SatInc(step uint8) (newval uint8, changed bool) {
	if h.S == 255 {
		changed = false
	} else if (255 - h.S) <= step {
		h.S = 255
		changed = true
	} else {
		h.S += step
		changed = true
	}
	newval = h.S
	return
}

func (h *Color) SatDec(step uint8) (newval uint8, changed bool) {
	if h.S == 0 {
		changed = false
	} else if h.S < step {
		h.S = 0
		changed = true
	} else {
		h.S -= step
		changed = true
	}
	newval = h.S
	return
}

func (h *Color) ValInc(step uint8) (newval uint8, changed bool) {
	if h.V == 255 {
		changed = false
	} else if (255 - h.V) <= step {
		h.V = 255
		changed = true
	} else {
		h.V += step
		changed = true
	}
	newval = h.V
	return
}

func (h *Color) ValDec(step uint8) (newval uint8, changed bool) {
	if h.V == 0 {
		changed = false
	} else if h.V < step {
		h.V = 0
		changed = true
	} else {
		h.V -= step
		changed = true
	}
	newval = h.V
	return
}
