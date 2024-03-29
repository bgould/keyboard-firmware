//go:build !teensy41

package main

import (
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/matrix/kinx/kintqt"
)

var (
	adapter = kintqt.NewAdapter(i2c)
	matrix  = adapter.NewMatrix()
)

func configureMatrix() {
	cli := board.CLI()
	if err := configureI2C(); err != nil {
		errmsg(err)
	}
	cli.WriteString("initializing matrix")
	if err := adapter.Initialize(); err != nil {
		cli.WriteString("matrix error: " + err.Error())
	} else {
		leds := kintqt.LEDs(0)
		leds.Set(kintqt.LEDCapsLock, false)
		leds.Set(kintqt.LEDNumLock, false)
		leds.Set(kintqt.LEDScrollLock, false)
		leds.Set(kintqt.LEDKeypad, keypadDefault)
		adapter.UpdateLEDs(leds)
		matrixInitialized = true
	}
}

func errmsg(err error) {
	for {
		println("error:", err)
		time.Sleep(2 * time.Second)
	}
}

func bootBlink() {
	if !matrixInitialized {
		return
	}
	for i, leds, on := 0, kintqt.LEDs(0), true; i < 10; i++ {
		on = !on
		leds.Set(kintqt.LEDKeypad, on)
		leds.Set(kintqt.LEDScrollLock, on)
		leds.Set(kintqt.LEDNumLock, on)
		leds.Set(kintqt.LEDCapsLock, on)
		adapter.UpdateLEDs(leds)
		time.Sleep(100 * time.Millisecond)
	}
	syncLEDs(keyboard.LEDs(7))
}

const keypadDefault = true

func syncLEDs(oldState keyboard.LEDs) keyboard.LEDs {
	leds := board.LEDs()
	if !matrixInitialized {
		return leds
	}
	caps := leds.Get(keyboard.LEDCapsLock)
	nlck := leds.Get(keyboard.LEDNumLock)
	slck := leds.Get(keyboard.LEDScrollLock)
	kpad := board.ActiveLayer() == 1
	if kpad {
		leds.Set(keyboard.LED(4), true)
	}
	// println(leds, caps, nlck, slck, kpad)
	if leds != oldState {
		//println("state change: ", leds, caps, nlck, slck)
		// oldState = leds
		qtleds := kintqt.LEDs(0)
		qtleds.Set(kintqt.LEDCapsLock, caps)
		qtleds.Set(kintqt.LEDNumLock, nlck)
		qtleds.Set(kintqt.LEDScrollLock, slck)
		qtleds.Set(kintqt.LEDKeypad, kpad)
		adapter.UpdateLEDs(qtleds)
	}
	return leds
}
