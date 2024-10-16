package keyboard

import (
	"io"
	"os"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const debug_macro = false

const macroDefaultTapDelayMS = 1

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
	arg    uint16
	end    time.Time
	send   struct {
		keycode keycodes.Keycode
		shifted bool
		altgred bool
		dead    bool
	}
}

var _ EventReceiver = (*defaultMacroDriver)(nil)

func (m *defaultMacroDriver) Configure() {

}

func (m *defaultMacroDriver) Count() uint8 {
	return m.count
}

func (m *defaultMacroDriver) VialMacroCount() uint8 {
	return m.Count()
}

func (m *defaultMacroDriver) VialMacroBuffer() []byte {
	return m.buffer
}

func (m *defaultMacroDriver) ReceiveEvent(ev Event) (bool, error) {
	return m.running(), nil
	// return false, nil
}

func (m *defaultMacroDriver) RunMacro(num uint8) (err error) {
	if head, tail, ok := m.macroNumBounds(num); ok {
		m.head, m.tail, m.curr = head, tail, 0
		m.op, m.arg, m.end = 0, 0, time.Now()
		if debug_macro {
			if buf, running := m.current(); running {
				xxdfprint(os.Stdout, 0x0, buf)
			}
		}
		return nil
	} else {
		m.head, m.tail = 0, 0
		return MacroErrInvalidNum
	}
}

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

func (m *defaultMacroDriver) running() bool {
	return m.tail > m.head && m.tail < len(m.buffer)
}

func (m *defaultMacroDriver) current() (buf []byte, running bool) {
	if running = m.running(); running {
		buf = m.buffer[m.head:m.tail]
	}
	return
}

// Task
func (m *defaultMacroDriver) Task(proc KeycodeProcessor) {
	if !m.running() {
		return
	}
	if time.Now().After(m.end) {
		kc := keycodes.Keycode(m.arg)
		// finish off previous operation before iterating to the next
		switch m.op {
		case MacroCodeTap, MacroCodeVialExtTap:
			proc.ProcessKeycode(kc, false)
			// println("end tap")
		case MacroCodeDelay:
			// println("end delay")
		case MacroCodeSend:
			if debug_macro {
				println("end send")
			}
			proc.ProcessKeycode(m.send.keycode, false)
			if m.send.shifted {
				proc.ProcessKeycode(keycodes.KC_LEFT_SHIFT, false)
			}
			if m.send.altgred {
				proc.ProcessKeycode(keycodes.KC_RIGHT_ALT, false)
			}
		}
		// get and execute next operation
		m.op, m.arg = m.nextOp()
		if debug_macro {
			println("next:", m.op.String(), m.arg, m.curr)
		}
		m.end = time.Now()
		kc = keycodes.Keycode(m.arg)
		switch m.op {
		case MacroCodeTap:
			// println("tapping", kc)
			proc.ProcessKeycode(kc, true)
			m.tapDelay()
		case MacroCodeDown:
			// println("down", kc)
			proc.ProcessKeycode(kc, true)
		case MacroCodeUp:
			// println("up", kc)
			proc.ProcessKeycode(kc, false)
		case MacroCodeDelay:
			// println("delay", arg)
			if m.arg > 0 {
				m.end = m.end.Add(time.Duration(m.arg) * time.Millisecond)
			}
		case MacroCodeVialExtTap:
			// println("tapping", kc)
			proc.ProcessKeycode(kc, true)
			m.tapDelay()
		case MacroCodeVialExtDown:
			// println("down", kc)
			proc.ProcessKeycode(kc, true)
		case MacroCodeVialExtUp:
			// println("up", kc)
			proc.ProcessKeycode(kc, false)
		case MacroCodeSend:
			if debug_macro {
				println("send", kc)
			}
			a := uint8(m.arg)
			m.send.keycode, m.send.shifted, m.send.altgred, m.send.dead = keycodes.AsciiToKeycode(a)
			if debug_macro {
				println("a2kc:", m.send.keycode, m.send.shifted, m.send.altgred, m.send.dead)
			}
			if m.send.shifted {
				proc.ProcessKeycode(keycodes.KC_LEFT_SHIFT, true)
			}
			if m.send.altgred {
				proc.ProcessKeycode(keycodes.KC_RIGHT_ALT, true)
			}
			proc.ProcessKeycode(m.send.keycode, true)
			m.tapDelay()
		}
	}
}

func (m *defaultMacroDriver) tapDelay() {
	m.end = m.end.Add(macroDefaultTapDelayMS * time.Millisecond)
}

func (m *defaultMacroDriver) nextOp() (code MacroCode, arg uint16) {
	buf, running := m.current()
	if !running {
		return m.reset()
	}
	if m.curr >= len(buf) {
		// println("end:", m.curr, len(running))
		return m.reset()
	}
	switch b := buf[m.curr]; b {
	case 0:
		// println("case 0")
		return m.reset()
	case MacroMagicPrefix:
		// println("magic:", m.curr)
		m.curr++
		switch op := MacroCode(buf[m.curr]); op {

		// "regular" QMK keycode actions; 1 byte arg following opcode
		case MacroCodeTap:
			fallthrough
		case MacroCodeDown:
			fallthrough
		case MacroCodeUp:
			m.curr++
			if m.curr >= len(buf) {
				println("unexpected reset 1")
				return m.reset()
			}
			arg := uint16(buf[m.curr])
			m.curr++
			return op, arg

		// delay opcode in milliseconds; 2 byte arg following opcode
		case MacroCodeDelay:
			// println("delay:", m.curr)
			m.curr++
			if m.curr+1 >= len(buf) {
				println("unexpected reset 2")
				return m.reset()
			}
			m.curr += 2
			low, high := buf[m.curr-2], buf[m.curr-1]
			// println("delay:", x, y)
			arg := (uint16(low - 1)) + (uint16(high-1) << 8)
			return op, arg

		// vial extended opcodes; 2 byte arg following opcode
		case MacroCodeVialExtTap:
			fallthrough
		case MacroCodeVialExtDown:
			fallthrough
		case MacroCodeVialExtUp:
			// println("vial ext: ", m.curr)
			m.curr++
			if m.curr+1 >= len(buf) {
				println("unexpected reset 3")
				return m.reset()
			}
			m.curr += 2
			low, high := buf[m.curr-2], buf[m.curr-1]
			// println("vial ext: ", high, low)
			arg := m.decodeKeycode((uint16(high) << 8) | (uint16(low) & 0xFF))
			return op, arg

		// invalid byte sequence in macro
		default:
			// println("default: ", b)
			return m.reset()
		}

	default:
		// println("send:", m.curr)
		// If the char wasn't magic, just send it
		// send_string_with_delay(data, DYNAMIC_KEYMAP_MACRO_DELAY);
		m.curr++
		return MacroCodeSend, uint16(b)
	}
}

func (m *defaultMacroDriver) reset() (code MacroCode, arg uint16) {
	m.head = 0
	m.tail = 0
	m.curr = 0
	return 0, 0
}

func (m *defaultMacroDriver) decodeKeycode(kc uint16) uint16 {
	// map 0xFF01 => 0x0100; 0xFF02 => 0x0200, etc
	if kc > 0xFF00 {
		return (kc & 0x00FF) << 8
	}
	return kc
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
