package main

import (
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/devices/kint"
	"github.com/bgould/keyboard-firmware/devices/kint/kint_pe"
	"github.com/bgould/keyboard-firmware/keyboard"
)

var (
	adapter = kint_pe.NewAdapter(i2c)

	host   = configureHost()
	matrix = adapter.NewMatrix()
	keymap = kint.KinTKeymap()
	board  = keyboard.New(&SerialConsole{machine.Serial}, host, matrix, keymap)
)

func main() {

	if _debug {
		time.Sleep(3 * time.Second)
		println("initializing hardware")
	}

	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

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
	// machine.LED.Set(false)
	setLED(true)

	if _debug {
		println("starting task loop")
	}

	board.SetDebug(_debug)
	for {
		board.Task()
		time.Sleep(100 * time.Microsecond)
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
	machine.LED.Set(on)
}
