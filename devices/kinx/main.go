package main

import (
	"runtime"
	"strconv"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

var (
	host   = initHost()
	keymap = initKeymap()
	board  = keyboard.New(host, matrix, keymap)

	matrixInitialized = false

	fn2made time.Time
	fn3made time.Time

	lastTotp uint64

	// rtcUpdate = make(chan struct{}, 1)
)

func init() {
	// TODO: make configurable
	time.Local = time.FixedZone("EST", -5*3600)

	board.SetKeyAction(keyboard.KeyActionFunc(keyAction))
	board.SetEnterBootloaderFunc(keyboard.DefaultEnterBootloader)
	board.SetCPUResetFunc(keyboard.DefaultCPUReset)

	initFilesystem()
	// initRTC()
}

func main() {

	if _debug {
		time.Sleep(3 * time.Second)
	}
	serial.Write([]byte("\r\n"))

	board.ConfigureFilesystem()

	board.EnableConsole(serial, initCommands())
	cli := board.CLI()
	cli.WriteString("---------------------")
	cli.WriteString("initializing hardware")

	configureMatrix()
	initDisplay()
	initRTC()

	cli.WriteString("matrix initialized: " + strconv.FormatBool(matrixInitialized))

	bootBlink()

	// go func() {
	// 	for {
	// 		<-rtcUpdate
	// 		cli.WriteString("syncing time")
	// 		rtcSync()
	// 	}
	// }()

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
		// timeTask()
		if matrixInitialized {
			board.Task()
			oldState = syncLEDs(oldState)
		}
		now := time.Now()
		if d := now.Sub(last); d > time.Second {
			ds.scanRate = (count * 1000) / int(d/time.Millisecond)
			// print("\r== scan:", ds.scanRate, " ==> \r")
			// println("count: ", count, " ", d/time.Millisecond, " ", )
			count = 0
			last = now
			ds.ts, ds.tsOk = last, true
		}
		totptask()
	}
}

// func configureKeyAction() keyboard.KeyActionFunc {
// return func(key keycodes.Keycode, made bool) {
func keyAction(key keycodes.Keycode, made bool) {
	switch key {

	// Handle "reset" press
	case keycodes.KC_FN2:
		if made {
			fn2made = time.Now()
		} else {
			if time.Since(fn2made) > 2*time.Second {
				board.EnterBootloader()
			} else {
				board.CPUReset()
			}
		}

	// Status output
	case keycodes.KC_FN3:
		if !made && time.Since(fn3made) > time.Second {
			setDisplay(false)
		} else if made {
			setDisplay(true)
			fn3made = time.Now()
		}
		if err := showTime(ds, true); err != nil {
			board.CLI().WriteString("warning: error updating display: " + err.Error())
		}
	}

}
