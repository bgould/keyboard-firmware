package keyboard

import (
	"io"
	"strconv"
)

func xxdfprint(w io.Writer, offset uint32, b []byte) {
	var l int
	var addr = make([]byte, 16)
	var data = make([]byte, 48)
	var buf16 = make([]byte, 16)
	for i, c := 0, len(b); i < c; i += 16 {
		a := offset + uint32(i)
		bin2hex([]byte{byte(a >> 24), byte(a >> 16), byte(a >> 8), byte(a)}, addr)
		w.Write(addr)
		w.Write([]byte(": "))
		l = i + 16
		if l >= c {
			l = c
		}
		for j, n := 0, l-i; j < 16; j++ {
			data[j*3] = ' '
			data[j*3+1] = ' '
			data[j*3+2] = ' '
			if j >= n {
				buf16[j] = ' '
			} else {
				var buf [2]byte
				bin2hex([]byte{byte(b[i+j])}, buf[:])
				data[j*3+1] = buf[0]
				data[j*3+2] = buf[1]
				if !strconv.IsPrint(rune(b[i+j])) {
					buf16[j] = '.'
				} else {
					buf16[j] = b[i+j]
				}
			}
		}
		w.Write(data)
		w.Write([]byte("    "))
		w.Write(buf16)
		w.Write([]byte("\r\n"))
	}
}

func bin2hex(in []byte, out []byte) {
	const (
		chars = "0123456789ABCDEF"
	)
	for i, b := range in {
		var n = i * 2
		out[n+0] = chars[b>>4]
		out[n+1] = chars[b&15]
	}
}
