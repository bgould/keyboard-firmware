package keyboard

import "io"

type Serialer interface {
	io.Writer
	io.ByteReader
	io.ByteWriter
	Buffered() int
}
