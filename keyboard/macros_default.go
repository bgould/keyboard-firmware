package keyboard

import (
	"io"
	"os"
	"time"
)

func NewDefaultMacroDriver(count uint8, bufferSize uint16) MacrosDriver {
	return &defaultMacroDriver{count: count, buffer: make([]byte, bufferSize)}
}

type defaultMacroDriver struct {
	count  uint8
	buffer []byte
	head   int
	tail   int
	curr   int
	op     MacroCode
	at     time.Time
}

func (m *defaultMacroDriver) Configure() {

}

func (m *defaultMacroDriver) RunMacro(num uint8) (err error) {
	if head, tail, ok := m.macroNumBounds(num); ok {
		m.head, m.tail, m.curr = head, tail, 0
		xxdfprint(os.Stdout, 0x0, m.running())
		for op, arg := m.next(); op != MacroCodeNone; op, arg = m.next() {
			// println("curr: ", m.curr)
			println("next: ", op, arg)
		}
		return nil
	} else {
		m.head, m.tail = 0, 0
		return MacroErrInvalidNum
	}
}

func (m *defaultMacroDriver) running() []byte {
	if m.tail > m.head && m.tail < len(m.buffer) {
		return m.buffer[m.head:m.tail]
	}
	return macroEmptyBuf[0:0]
}

func (m *defaultMacroDriver) reset() (code MacroCode, arg uint16) {
	m.head = 0
	m.tail = 0
	m.curr = 0
	return 0, 0
}

func (m *defaultMacroDriver) next() (code MacroCode, arg uint16) {
	running := m.running()
	if m.curr >= len(running) {
		return m.reset()
	}
	switch b := running[m.curr]; b {
	case 0:
		m.reset()
		return MacroCodeNone, 0
	case MacroMagicPrefix:
		m.curr++
		switch op := MacroCode(running[m.curr]); op {

		// "regular" QMK keycode actions; 1 byte arg following opcode
		case MacroCodeTap:
			fallthrough
		case MacroCodeDown:
			fallthrough
		case MacroCodeUp:
			m.curr++
			if m.curr+1 >= len(running) {
				return m.reset()
			}
			m.curr++
			return op, uint16(running[m.curr])

		// delay opcode in milliseconds; 2 byte arg following opcode
		case MacroCodeDelay:
			m.curr++
			if m.curr+2 >= len(running) {
				return m.reset()
			}
			m.curr += 2
			return op, (uint16(running[m.curr-1])) + (uint16(running[m.curr]))

		// vial extended opcodes; 2 byte arg following opcode
		case MacroCodeVialExtTap:
			fallthrough
		case MacroCodeVialExtDown:
			fallthrough
		case MacroCodeVialExtUp:
			m.curr++
			if m.curr+2 >= len(running) {
				return m.reset()
			}
			m.curr += 2
			return op, (uint16(running[m.curr]) >> 8) | (uint16(running[m.curr-1]))

		// invalid byte sequence in macro
		default:
			return m.reset()
		}

	default:
		// If the char wasn't magic, just send it
		// send_string_with_delay(data, DYNAMIC_KEYMAP_MACRO_DELAY);
		m.curr++
		return MacroCodeSend, uint16(b)
	}
}

var macroEmptyBuf [0]byte

// func (m *defaultMacroDriver) macroBytes(macroNum uint8) []byte {
// 	start, end, ok := m.macroNumBounds(macroNum)
// 	if !ok {
// 		return macroEmptyBuf[:]
// 	}
// 	return m.buffer[start:end]
// }

// determine bounds of specified macro in buffer
func (m *defaultMacroDriver) macroNumBounds(macroNum uint8) (start, end int, ok bool) {
	if len(m.buffer) == 0 {
		return 0, 0, false
	}
	if macroNum >= m.count {
		return 0, 0, false
	}
	for i, c, n := 0, len(m.buffer), uint8(0); i < c; i++ {
		if b := m.buffer[i]; b == 0x0 {
			if n == macroNum {
				end = i
				return start, end, true
			} else {
				n++
				start = i + 1
			}
		}
	}
	return start, end, false
}

// Task
func (m *defaultMacroDriver) Task() {

}

func (m *defaultMacroDriver) Count() uint8 {
	return m.count
}

func (m *defaultMacroDriver) VialMacroBuffer() []byte {
	return m.buffer
}

func (m *defaultMacroDriver) StoredSize() int {
	bufSize := len(m.buffer)
	hdrSize := 4
	ftrSize := 4
	return hdrSize + bufSize + ftrSize
}

func (m *defaultMacroDriver) ZeroFill() {
	for iCol := range m.buffer {
		m.buffer[iCol] = 0x0
	}
}

var _ io.WriterTo = (*defaultMacroDriver)(nil)

func (m *defaultMacroDriver) WriteTo(w io.Writer) (n int64, err error) {
	var result int
	// header
	result, err = w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	n += int64(result)
	if err != nil {
		return
	}
	result, err = w.Write(m.buffer)
	n += int64(result)
	if err != nil {
		return
	}
	// footer
	result, err = w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	n += int64(result)
	return
}

var _ io.ReaderFrom = (*defaultMacroDriver)(nil)

func (m *defaultMacroDriver) ReadFrom(r io.Reader) (n int64, err error) {
	var result int
	var buffer [8]byte
	header := buffer[0:4]
	footer := buffer[4:8]
	// header
	result, err = r.Read(header)
	n += int64(result)
	if err != nil {
		return
	}
	// buffer
	result, err = r.Read(m.buffer)
	n += int64(result)
	if err != nil {
		return
	}
	// footer
	result, err = r.Read(footer)
	n += int64(result)
	if err != nil {
		return
	}
	return
}
