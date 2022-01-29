package main

import (
	"device/arm"
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const _debug = false

var (
	keymap = []keyboard.Keymap{KinTKeymap()}
	matrix = keyboard.NewMatrix(15, 7, keyboard.RowReaderFunc(ReadRow))

	lastSecond = time.Now()
	counter    uint
	average    uint
	seconds    uint
)

func main() {
	configurePins()
	host := configureHost()
	board := keyboard.New(machine.Serial, host, matrix, keymap).
		WithDebug(_debug).
		WithJumpToBootloader(jumpToBootloader)
	for {
		for since := time.Since(lastSecond); since >= time.Second; since = 0 {
			average = uint(float32(counter) * (float32(time.Second) / float32(since)))
			println(seconds, "-", "average:", average)
			println(seconds, "-", "   leds:", host.LEDs())
			// reset
			lastSecond = lastSecond.Add(since)
			counter = 0
			seconds++
		}
		board.Task()
		counter++
		time.Sleep(100 * time.Microsecond)
	}
}

func jumpToBootloader() {
	delayMicros(100)
	arm.Asm("bkpt")
}
