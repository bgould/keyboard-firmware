package main

import (
	"runtime"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbhid"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
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
	bootBlink()

	board.SetKeyAction(func() keyboard.KeyActionFunc {
		var previous bool
		return func(key keycodes.Keycode, made bool) {
			switch key {
			case keycodes.FN0:
				if previous && !made {
					if board.ActiveLayer() == 1 {
						board.SetActiveLayer(0)
					} else {
						board.SetActiveLayer(1)
					}
				}
				previous = made
			}
		}
	}())

	println("starting task loop")
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
	return usbhid.New()
}
