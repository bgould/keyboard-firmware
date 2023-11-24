package main

import (
	"machine"
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

	board.SetKeyAction(configureKeyAction())

	println("starting task loop")
	go deviceLoop()
	for {
		// runtime.Gosched()
		time.Sleep(1 * time.Second)
	}
}

func deviceLoop() {
	var oldState keyboard.LEDs
	for last, count := time.Now(), 0; true; count++ {
		board.Task()
		oldState = syncLEDs(oldState)
		// runtime.Gosched()
		if d := time.Since(last); d > time.Second {
			println("count: ", count, " ", d/time.Millisecond, " ", (count*1000)/int(d/time.Millisecond))
			count = 0
			last = time.Now()
		}
	}
}

func configureHost() keyboard.Host {
	return usbhid.New()
}

func configureKeyAction() keyboard.KeyActionFunc {
	var fn0made bool
	var fn1prev uint8
	return func(key keycodes.Keycode, made bool) {
		switch key {
		case keycodes.FN0:
			if fn0made && !made {
				if board.ActiveLayer() == 1 {
					board.SetActiveLayer(0)
				} else {
					board.SetActiveLayer(1)
				}
			}
			fn0made = made
		case keycodes.FN1:
			if made {
				fn1prev = board.ActiveLayer()
				board.SetActiveLayer(2)
				println("programming layer on - ", board.ActiveLayer())
			} else {
				board.SetActiveLayer(fn1prev)
				fn1prev = 0
				println("programming layer off - ", board.ActiveLayer())
			}
			if fn1prev == 2 {
				fn1prev = 0
			}
		case keycodes.FN2:
			if made {
				machine.CPUReset()
			}
		}
	}
}
