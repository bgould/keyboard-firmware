//go:build tinygo

package main

import (
	"machine"
)

func cpuReset() {
	machine.CPUReset()
}
