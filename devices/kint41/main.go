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
	board  = keyboard.New(machine.Serial, host, matrix, keymap)
)

func init() {
	configurePins()
}

func main() {
	board.SetDebug(_debug)
	board.SetBootloaderJump(func() {
		println("jumping to bootloader")
		delayMicros(10000)
		arm.Asm("bkpt")
	})

	go bootBlink()
	for {
		board.Task()
		metrics()
		// delayMicros(100)
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
	since := time.Since(lastSecond)
	if since >= time.Second {
		// if counter >= 200 {
		// 	since := time.Since(lastSecond)
		average = uint(float32(counter) * (float32(time.Second) / float32(since)))
		println(seconds, "-", "average:", average, "  leds", host.LEDs())
		// reset
		lastSecond = lastSecond.Add(since)
		counter = 0
		seconds++
	}
	counter++
}
