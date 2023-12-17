//go:build !tinygo || !nrf52840

package main

func jumpToBootloader() {
	println("jumpToBootloader not implemented")
}
