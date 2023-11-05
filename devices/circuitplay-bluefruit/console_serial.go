//go:build !console.tinyterm

package main

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

func configureConsole() keyboard.Console {
	return &SerialConsole{machine.Serial}
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
