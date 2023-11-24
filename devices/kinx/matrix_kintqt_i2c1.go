//go:build tinygo && !teensy41 && i2c1 && !i2c0

package main

import (
	"machine"
)

var i2c = machine.I2C1
