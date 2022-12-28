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
		setLED(true)
		// machine.LED.Set(true)
		println("error:", err)
		time.Sleep(time.Second)
		setLED(false)
		// machine.LED.Set(false)
		time.Sleep(time.Second)
	}
}

func setLED(on bool) {
	// machine.LED.Set(on)
}
