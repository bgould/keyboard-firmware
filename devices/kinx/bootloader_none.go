//go:build !tinygo || (!nrf52840 && !rp2040)

package main

func jumpToBootloader() {
	println("jumpToBootloader not implemented")
}
