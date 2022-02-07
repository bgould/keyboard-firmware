package main

import (
	"device/arm"
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
)

var (
	matrix = keyboard.NewMatrix(15, 7, keyboard.RowReaderFunc(ReadRow))
	keymap = KinTKeymap()
	host   = configureHost()
)

func init() {
	configurePins()
}

func main() {
	board := keyboard.New(machine.Serial, host, matrix, keymap).
		WithDebug(_debug).
		WithJumpToBootloader(func() {
			println("jumping to bootloader")
			delayMicros(10000)
			arm.Asm("bkpt")
		})

	go bootBlink()
	for {
		board.Task()
		metrics()
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
