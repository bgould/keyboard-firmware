package main

import (
	"runtime"
	"strconv"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
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

	rtcUpdate = make(chan struct{}, 1)
)

var _ vial.KeySetter = (keyboard.Keymap)(nil)

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

	go func() {
		for {
			<-rtcUpdate
			cli.WriteString("syncing time")
			rtcSync()
		}
	}()

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

type VialDriver struct {
	keyboard.Keymap
}

var _ vial.Handler = (*VialDriver)(nil)

func (d *VialDriver) Handle(rx []byte, tx []byte) (sendTx bool) {
	// println("called Handle()", rx[0], rx[1])
	switch rx[0] {
	case 0xEE:
		switch rx[1] {
		case 0x01: // set time
			var unixTime uint64
			unixTime |= uint64(rx[2]) << 56
			unixTime |= uint64(rx[3]) << 48
			unixTime |= uint64(rx[4]) << 40
			unixTime |= uint64(rx[5]) << 32
			unixTime |= uint64(rx[6]) << 24
			unixTime |= uint64(rx[7]) << 16
			unixTime |= uint64(rx[8]) << 8
			unixTime |= uint64(rx[9]) << 0
			println("\nsetting unix time", int64(unixTime))
			setUnixTime(time.Unix(int64(unixTime), 0))
			// cmd := console.CommandInfo{
			// 	Cmd:    "time",
			// 	Argv:   []string{"set", strconv.FormatInt(int64(unixTime), 10)},
			// 	Stdout: cli,
			// }
			// timeset(cmd)
		}
	}
	return false
}
