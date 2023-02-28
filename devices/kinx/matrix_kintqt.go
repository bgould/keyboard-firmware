//go:build !teensy41

package main

import (
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/matrix/kinx/kintqt"
)

var (
	console = &SerialConsole{machine.Serial}
	adapter = kintqt.NewAdapter(i2c)
	matrix  = adapter.NewMatrix()
)

func configureMatrix() {
	// initialize I2C bus
	err := i2c.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})
	if err != nil {
		errmsg(err)
	}
	if _debug {
		println("intializing matrix")
	}
	if err := adapter.Initialize(); err != nil {
		errmsg(err)
	}
	leds := kintqt.LEDs(0)
	adapter.UpdateLEDs(leds)
}

type SerialConsole struct {
	machine.Serialer
}

func (sc *SerialConsole) Read(buf []byte) (n int, err error) {
	for i := range buf {
		buf[i], err = sc.ReadByte()
		if err != nil {
			n = i - 1
			return n, err
		}
	}
	return len(buf), nil
}

func errmsg(err error) {
	for {
		println("error:", err)
		time.Sleep(2 * time.Second)
	}
}

func bootBlink() {
	for i, leds, on := 0, kintqt.LEDs(0), true; i < 10; i++ {
		on = !on
		leds.Set(kintqt.LEDKeypad, on)
		leds.Set(kintqt.LEDScrollLock, on)
		leds.Set(kintqt.LEDNumLock, on)
		leds.Set(kintqt.LEDCapsLock, on)
		adapter.UpdateLEDs(leds)
		time.Sleep(100 * time.Millisecond)
	}
}
