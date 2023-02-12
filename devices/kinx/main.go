package main

import (
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/hosts/usbhid"
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

	for {
		board.Task()
		time.Sleep(100 * time.Microsecond)
	}

}

func configureHost() keyboard.Host {
	return usbhid.New()
}
