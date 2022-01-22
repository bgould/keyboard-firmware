//go:build tinygo
// +build tinygo

package usbhid

import (
	"machine/usb"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const debug = false

type Host struct {
	kbd *usb.Keyboard
}

/*
type MouseButton byte

const (
	MouseBtnLeft   MouseButton = 0x01
	MouseBtnRight  MouseButton = 0x02
	MouseBtnMiddle MouseButton = 0x04
)

type ConsumerKey uint16

const (
	ConsKeyHome       ConsumerKey = 0x0100
	ConsKeyKbdLayout  ConsumerKey = 0x0200
	ConsKeySearch     ConsumerKey = 0x0400
	ConsKeySnapshot   ConsumerKey = 0x0800
	ConsKeyVolUp      ConsumerKey = 0x1000
	ConsKeyVolDown    ConsumerKey = 0x2000
	ConsKeyPlayPause  ConsumerKey = 0x4000
	ConsKeyFastFwd    ConsumerKey = 0x8000
	ConsKeyRewind     ConsumerKey = 0x0001
	ConsKeyNextTrack  ConsumerKey = 0x0002
	ConsKeyPrevTrack  ConsumerKey = 0x0004
	ConsKeyRandomPlay ConsumerKey = 0x0008
	ConsKeyStop       ConsumerKey = 0x0010
)

func (r *Report) Keyboard(mod KeyboardModifier, keys ...byte) *Report {
	r[0] = byte(mod)
	r[1] = 0x0
	for i, c := 0, len(keys); i < 6; i++ {
		if i < c {
			r[i+2] = keys[i]
		} else {
			r[i+2] = 0x0
		}
	}
	return r
}

func (r *Report) Mouse(buttons MouseButton, x int8, y int8) *Report {
	r[0] = 0x0
	r[1] = 0x3
	r[2] = byte(buttons)
	r[3] = byte(x)
	r[4] = byte(y)
	r[5] = 0x0
	r[6] = 0x0
	r[7] = 0x0
	return r
}

func (r *Report) Consumer(key ConsumerKey) *Report {
	r[0] = 0x0
	r[1] = 0x2
	r[2] = byte(key >> 8)
	r[3] = byte(key & 0xFF)
	r[4] = 0x0
	r[5] = 0x0
	r[6] = 0x0
	r[7] = 0x0
	return r
}
*/

func New(hid *usb.HID) *Host {
	return &Host{
		kbd: hid.Keyboard(),
	}
}

func (host *Host) Send(rpt *keyboard.Report) {
	if debug {
		println(rpt[0], rpt[1], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
	}
	host.kbd.SendKeys(rpt[0], rpt[2], rpt[3], rpt[4], rpt[5], rpt[6], rpt[7])
}
