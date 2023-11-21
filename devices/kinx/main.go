package main

import (
	"runtime"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
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
	// go bootBlink()
	go deviceLoop()
	for {
		runtime.Gosched()
		// time.Sleep(1 * time.Second)
	}
}

func deviceLoop() {
	var oldState keyboard.LEDs
	for {
		board.Task()
		oldState = syncLEDs(oldState)
		runtime.Gosched()
	}
}

func configureHost() keyboard.Host {
	// return multihost.New(serial.New(machine.Serial), usbhid.New())
	// return serial.New(machine.Serial)
	return usbhid.New()
}
