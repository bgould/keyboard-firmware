//go:build (tinygo && i2c1) || (tinygo && feather_rp2040)
// +build tinygo,i2c1 tinygo,feather_rp2040

package main

import (
	"machine"
)

var i2c = machine.I2C1
