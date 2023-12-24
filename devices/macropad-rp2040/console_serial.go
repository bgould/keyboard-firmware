//go:build macropad_rp2040 && !console.tinyterm

package main

import "machine"

var serialer = machine.Serial

// func configureConsole() keyboard.Console {
// 	return &SerialConsole{machine.Serial}
// }

// type SerialConsole struct {
// 	machine.Serialer
// }

// func (sc *SerialConsole) Read(buf []byte) (n int, err error) {
// 	for i := range buf {
// 		buf[i], err = sc.ReadByte()
// 		if err != nil {
// 			n = i - 1
// 			return n, err
// 		}
// 	}
// 	return len(buf), nil
// }
