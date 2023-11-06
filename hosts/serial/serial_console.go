//go:build tinygo

package serial

import (
	"machine"
)

func DefaultConsole() *Console {
	return &Console{machine.Serial}
}

type Console struct {
	machine.Serialer
}

func (sc *Console) Read(buf []byte) (n int, err error) {
	for i := range buf {
		buf[i], err = sc.ReadByte()
		if err != nil {
			n = i - 1
			return n, err
		}
	}
	return len(buf), nil
}
