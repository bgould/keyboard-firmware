package main

import (
	"machine"
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

	// keyAction = configureKeyAction()

	fn0made bool
	fn1prev uint8
	fn2made time.Time
	fn3made time.Time
)

func init() {
	board.SetDebug(_debug)
	board.SetKeyAction(keyboard.KeyActionFunc(keyAction))
}

func main() {

	if _debug {
		time.Sleep(3 * time.Second)
	}

	println("initializing hardware")
	configureMatrix()
	initDisplay()
	initTime()

	bootBlink()

	println("starting task loop")
	go deviceLoop()
	for {
		runtime.Gosched()
		// time.Sleep(1 * time.Second)
	}
}

func deviceLoop() {
	var oldState keyboard.LEDs
	for last, count := time.Now(), 0; true; count++ {
		board.Task()
		timeTask()
		oldState = syncLEDs(oldState)
		// runtime.Gosched()
		if d := time.Since(last); d > time.Second {
			println("count: ", count, " ", d/time.Millisecond, " ", (count*1000)/int(d/time.Millisecond))
			count = 0
			last = time.Now()
			ds.ts, ds.tsOk = last, true
			// ds.ts, ds.tsOk = readTime()
			// ds.ts, ds.tsOk = last, true
			if err := showTime(&ds, false); err != nil {
				println("warning: error updating display", err)
			}
		}
		displayTask()
	}
}

func configureHost() keyboard.Host {
	return usbhid.New()
}

// func configureKeyAction() keyboard.KeyActionFunc {
// return func(key keycodes.Keycode, made bool) {
func keyAction(key keycodes.Keycode, made bool) {
	switch key {

	// Toggle keypad layer on keypress
	case keycodes.FN0:
		if fn0made && !made {
			if board.ActiveLayer() == 1 {
				board.SetActiveLayer(0)
			} else {
				board.SetActiveLayer(1)
			}
		}
		fn0made = made

	// Toggle programming layer on key down/up
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

	// Handle "reset" press
	case keycodes.FN2:
		if made {
			fn2made = time.Now()
		} else {
			if time.Since(fn2made) > 2*time.Second {
				jumpToBootloader()
			} else {
				machine.CPUReset()
			}
		}

	case keycodes.FN3:
		if !made && time.Since(fn3made) > time.Second {
			setDisplay(false)
		} else if made {
			setDisplay(true)
			fn3made = time.Now()
		}
		if err := showTime(&ds, true); err != nil {
			println("warning: error updating display", err)
		}
	}

}

// }

var ds DisplayState

type DisplayState struct {
	ts   time.Time
	tsOk bool
}
