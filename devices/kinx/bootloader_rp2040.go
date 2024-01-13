//go:build tinygo && rp2040

package main

import (
	"machine"
)

func jumpToBootloader() {
	machine.EnterBootloader()
}
