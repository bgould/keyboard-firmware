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
		println("initializing hardware")
	}

	configureMatrix()

	if _debug {
		println("starting task loop")
	}
	board.SetDebug(_debug)

	go bootBlink()
	go deviceLoop()

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

func configureHost() keyboard.Host {
	return usbhid.New()
}
