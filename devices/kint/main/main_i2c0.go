//go:build !feather_rp2040 && !i2c1
// +build !feather_rp2040,!i2c1

package main

var i2c = machine.I2C0
