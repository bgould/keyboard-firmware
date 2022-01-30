package main

import (
	"device/arm"
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = false

var (
	matrix = keyboard.NewMatrix(15, 7, keyboard.RowReaderFunc(ReadRow))
	keymap = KinTKeymap()
	host   = configureHost()

	board = keyboard.New(machine.Serial, host, matrix, keymap).
		WithDebug(_debug).
		WithJumpToBootloader(func() { arm.Asm("bkpt") })
)

func init() {
	configurePins()
}

func main() {
	go bootBlink()
	for {
		board.Task()
		// metrics()
		time.Sleep(100 * time.Microsecond)
	}
}

func bootBlink() {
	var on bool
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		on = !on
		for _, pin := range leds {
			pin.Set(on)
		}
	}
}

var (
	lastSecond = time.Now()
	counter    uint
	average    uint
	seconds    uint
)

func metrics() {
	for since := time.Since(lastSecond); since >= time.Second; since = 0 {
		average = uint(float32(counter) * (float32(time.Second) / float32(since)))
		println(seconds, "-", "average:", average)
		println(seconds, "-", "   leds:", host.LEDs())
		// reset
		lastSecond = lastSecond.Add(since)
		counter = 0
		seconds++
	}
	counter++
}
