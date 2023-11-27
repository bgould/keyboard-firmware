package main

import (
	"runtime"
	"strconv"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
	totp "github.com/bgould/tinytotp"
)

var (
	cli    = initConsole()
	host   = initHost()
	keymap = initKeymap()
	board  = keyboard.New(serial, host, matrix, keymap)

	matrixInitialized = false
	// keyAction = configureKeyAction()

	fn0made bool
	fn1prev uint8
	fn2made time.Time
	fn3made time.Time

	lastTotp uint64
)

func init() {
	board.SetDebug(_debug)
	board.SetKeyAction(keyboard.KeyActionFunc(keyAction))
}

func main() {

	if _debug {
		time.Sleep(3 * time.Second)
	}

	serial.Write([]byte("\r\n"))
	cli.WriteString("---------------------")
	cli.WriteString("initializing hardware")

	configureMatrix()
	initDisplay()
	initTime()

	cli.WriteString("matrix initialized: " + strconv.FormatBool(matrixInitialized))

	bootBlink()

	cli.WriteString("starting task loop")
	cli.WriteString("---------------------")
	go deviceLoop()
	for {
		runtime.Gosched()
		// time.Sleep(1 * time.Second)
	}
}

func deviceLoop() {
	var oldState keyboard.LEDs
	for last, count := time.Now(), 0; true; count++ {
		timeTask()
		if matrixInitialized {
			board.Task()
			oldState = syncLEDs(oldState)
		}
		// runtime.Gosched()
		now := time.Now()
		if now, d := time.Now(), now.Sub(last); d > time.Second {

			ds.scanRate = (count * 1000) / int(d/time.Millisecond)

			// TOTP-related functionality
			if totpKeys[0].Name != "" && totpKeys[0].Key != "" {
				ds.totpCounter = uint64(totp.TimeBasedCounter(time.Now(), totp.DefaultOpts.Period))
				if ds.totpCounter != lastTotp {
					// TODO: un-hardcode index
					ds.totpAccount = totpKeys[0].Name
					numbers, err := totp.GenerateCode(string(totpKeys[0].Key), now)
					if err != nil {
						cli.WriteString("warning: error updating TOTP - " + err.Error())
						numbers = "000000"
					}
					ds.totpNumbers = numbers
					lastTotp = ds.totpCounter
				}
			}

			// print("\r== scan:", ds.scanRate, " ==> \r")
			// println("count: ", count, " ", d/time.Millisecond, " ", )
			count = 0
			last = time.Now()
			ds.ts, ds.tsOk = last, true
			if err := showTime(ds, false); err != nil {
				cli.WriteString("warning: error updating display - " + err.Error())
			}

		}
		displayTask()
		cli.Task()
	}
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
		} else {
			board.SetActiveLayer(fn1prev)
			fn1prev = 0
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
				cpuReset()
			}
		}

	// Status output
	case keycodes.FN3:
		if !made && time.Since(fn3made) > time.Second {
			setDisplay(false)
		} else if made {
			setDisplay(true)
			fn3made = time.Now()
		}
		if err := showTime(ds, true); err != nil {
			cli.WriteString("warning: error updating display: " + err.Error())
		}
	}

}
