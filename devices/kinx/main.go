package main

import (
	"runtime"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/matrix/kinx/kintqt"
)

var (
	host   = configureHost()
	keymap = Keymap()
	board  = keyboard.New(console, host, matrix, keymap)
)

func main() {

	if _debug {
		time.Sleep(3 * time.Second)
	}

	println("initializing hardware")
	configureMatrix()

	board.SetDebug(_debug)

	println("starting task loop")
	go bootBlink()
	go deviceLoop()
	go updateLEDs()

	for {
		runtime.Gosched()
		// time.Sleep(1 * time.Second)
	}
}

func deviceLoop() {
	for {
		board.Task()

		runtime.Gosched()
	}
}

func updateLEDs() {
	for {

		leds := board.LEDs()
		caps := leds.Get(keyboard.LEDCapsLock)   // (uint8(leds) & uint8(1<<(keyboard.LEDCapsLock-1)))
		nlck := leds.Get(keyboard.LEDNumLock)    // (uint8(leds) & uint8(1<<(keyboard.LEDNumLock-1)))
		slck := leds.Get(keyboard.LEDScrollLock) // (uint8(leds) & uint8(1<<(keyboard.LEDScrollLock-1)))
		// println(leds, caps, nlck, slck)

		qtleds := kintqt.LEDs(0)
		qtleds.Set(kintqt.LEDCapsLock, caps)
		qtleds.Set(kintqt.LEDNumLock, nlck)
		qtleds.Set(kintqt.LEDScrollLock, slck)
		adapter.UpdateLEDs(qtleds)

		time.Sleep(time.Millisecond)

	}
}

func configureHost() keyboard.Host {
	return usbhid.New()
}
