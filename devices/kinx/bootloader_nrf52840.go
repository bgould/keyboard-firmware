//go:build tinygo && nrf52840

package main

import (
	"machine"
)

func jumpToBootloader() {
	machine.EnterUF2Bootloader()
}
