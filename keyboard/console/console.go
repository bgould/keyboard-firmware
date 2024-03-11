package console

import (
	"io"
	"strconv"
	"strings"
)

const consoleBufLen = 64

const (
	prompt = "==> "
	result = "<-- "
)

// const storageBufLen = 512
type Serialer interface {
	io.Writer
	io.ByteReader
	io.ByteWriter
	Buffered() int
}

func New(serial Serialer, handler CommandHandler) *Console {
	return &Console{
		serial:  serial,
		handler: handler,
	}
}

type Console struct {
	serial Serialer

	state  taskState
	input  [consoleBufLen]byte
	cursor int
	// store [storageBufLen]byte

	handler CommandHandler
}

type taskState int

const (
	taskStateInput = iota
	taskStateEscape
	taskStateEscBrc
	taskStateCSI
)

func (c *Console) Task() {

	// c.prompt()

	// for i := 0; ; {
	if c.serial.Buffered() > 0 {
		data, _ := c.serial.ReadByte()
		// if debug {
		// fmt.Printf("\rdata: %x\r\n\r", data)
		// c.prompt()
		// c.serial.Write(c.input[:c.cursor])
		// }
		switch c.state {
		case taskStateInput:
			switch data {
			case 0x8:
				fallthrough
			case 0x7f: // this is probably wrong... works on my machine tho :)
				// backspace
				if c.cursor > 0 {
					c.cursor -= 1
					c.serial.Write([]byte{0x8, 0x20, 0x8})
				}
			case 12:
				// Ctrl+L (form feed)
				c.serial.Write([]byte("\r\n"))
				c.prompt()
				c.cursor = 0
				return
			case 13:
				// return key (carriage return)
				c.serial.Write([]byte("\r\n"))
				c.runCommand(string(c.input[:c.cursor]))
				c.prompt()

				c.cursor = 0
				return
			case 27:
				// escape
				c.state = taskStateEscape
			default:
				// anything else, just echo the character if it is printable
				if strconv.IsPrint(rune(data)) {
					if c.cursor < (consoleBufLen - 1) {
						c.serial.WriteByte(data)
						c.input[c.cursor] = data
						c.cursor++
					}
				}
			}
		case taskStateEscape:
			switch data {
			case 0x5b:
				c.state = taskStateEscBrc
			default:
				c.state = taskStateInput
			}
		default:
			// TODO: handle escape sequences
			c.state = taskStateInput
		}
	}
	// }
}

func (c *Console) WriteString(s string) (n int, err error) {
	c.serial.Write([]byte("\r"))
	n, err = c.serial.Write([]byte(s))
	c.serial.Write([]byte("\r\n"))
	c.prompt()
	return
}

func (c *Console) Write(b []byte) (n int, err error) {
	// c.serial.Write([]byte("\r"))
	n, err = c.serial.Write(b)
	// c.serial.Write([]byte("\r\n"))
	// c.prompt()
	return
}

func (c *Console) runCommand(line string) {
	if c.handler == nil {
		return
	}
	if line == "" {
		return
	}

	argv := strings.SplitN(strings.TrimSpace(line), " ", -1)
	info := CommandInfo{Argv: argv, Stdout: c}
	exit := c.handler.HandleCommand(info)
	if exit != 0 {
		if exit == CommandUnknown {
			c.serialPrint([]byte(result))
			c.serialPrint([]byte("unknown command: "))
			c.serialPrint([]byte(line))
		} else {
			c.serialPrint([]byte(result))
			c.serialPrint([]byte("exit code: "))
			c.serialPrint([]byte(strconv.Itoa(exit)))
		}
		c.serialPrint([]byte("\n"))
	}
	// }
	// cmdfn(argv)
}

func (c *Console) prompt() {
	c.serialPrint([]byte(prompt))
}

func (c *Console) serialPrint(buf []byte) {
	c.serial.Write(buf)
}
