//go:build tinygo && i2c1

package main

import (
	"machine"
)

var i2c = machine.I2C1
