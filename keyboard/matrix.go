package keyboard

import (
	"time"
)

const (
	debounceCount    = 5
	debounceInterval = 100 * time.Microsecond
)

type RowReader interface {
	ReadRow(rowIndex uint8) Row
}

type RowReaderFunc func(rowIndex uint8) Row

func (fn RowReaderFunc) ReadRow(rowIndex uint8) Row {
	return fn(rowIndex)
}

func NewMatrix(r, c uint8, io RowReader) *Matrix {
	matrix := &Matrix{
		io:         io,
		nRows:      r,
		nCols:      c,
		rows:       make([]Row, r),
		debouncing: make([]Row, r),
	}
	return matrix
}

type Matrix struct {
	io         RowReader
	nRows      uint8
	nCols      uint8
	rows       []Row
	debouncing []Row
	debounce   uint8
	ghosting   bool
	encoders   *encoders
	encStates  []encoderState
}

type EncoderPos struct {
	Encoder
	PosCW  Pos
	PosCCW Pos
}

func (pos *EncoderPos) isInRow(i uint8) bool {
	return pos.isInRowCW(i) || pos.isInRowCCW(i)
}

func (pos *EncoderPos) isInRowCW(i uint8) bool {
	return pos.PosCW.Row == i
}
func (pos *EncoderPos) isInRowCCW(i uint8) bool {
	return pos.PosCCW.Row == i
}

type encoderState struct {
	EncoderPos
	turned    time.Time
	clockwise bool
}

func (m *Matrix) WithEncoders(encs ...EncoderPos) *Matrix {
	if encs == nil || len(encs) == 0 {
		m.encoders = nil
	}
	m.encoders = &encoders{
		encoders:   make([]Encoder, len(encs)),
		values:     make([]int, len(encs)),
		subcribers: []EncodersSubscriber{EncodersSubscriberFunc(m.encoderChanged)},
	}
	m.encStates = make([]encoderState, len(encs))
	for i, enc := range encs {
		// println("adding encoder:", i, enc.PosCW.Col, enc.PosCCW.Col)
		m.encoders.encoders[i] = enc.Encoder
		m.encStates[i] = encoderState{EncoderPos: enc}
	}
	return m
}

func (m *Matrix) encoderChanged(index int, clockwise bool) {
	if index < 0 || index >= len(m.encStates) {
		return
	}
	state := &m.encStates[index]
	state.turned = time.Now()
	state.clockwise = clockwise
	// println("encoder state changed:", index, state.turned.String(), state.clockwise)
}

func (m *Matrix) Rows() uint8 {
	return m.nRows
}

func (m *Matrix) Cols() uint8 {
	return m.nCols
}

func (m *Matrix) Ghosting() bool {
	return m.ghosting
}

func (m *Matrix) WithGhosting(hasGhosting bool) *Matrix {
	m.ghosting = hasGhosting
	return m
}

//go:inline
func (m *Matrix) GetRow(row uint8) Row {
	return m.rows[row]
}

func (m *Matrix) IsOn(row uint8, col uint8) bool {
	return m.GetRow(row).IsOn(col)
}

func (m *Matrix) HasGhostInRow(row uint8) bool {
	if !m.ghosting {
		return false
	}
	r := m.GetRow(row)
	// if there are less than 2 keys down in the row, there is no ghost
	n := 0
	for i := uint8(0); i < row; i++ {
		if r.IsOn(i) {
			n++
		}
	}
	if n < 2 {
		return false
	}
	//if (r - 1&r) == 0 { // TODO: copied from TMK; evaluate to see if this works
	//	return false
	//}
	for i := uint8(0); i < m.nRows; i++ {
		if i != row && m.GetRow(i)&r > 0 {
			return true
		}
	}
	return false
}

func (m *Matrix) Scan() (changed bool) {
	const (
		encoderInterval = 5 * time.Millisecond
	)

	// loop over rows and probe the columns for each
	for i := uint8(0); i < m.nRows; i++ {

		// scan the matrix for any normal keypresses
		row := m.io.ReadRow(i)

		// update matrix state with encoder state if applicable
		if m.encoders != nil {
			m.encoders.EncodersTask()
			for _, state := range m.encStates {
				if state.isInRow(i) && time.Since(state.turned) < encoderInterval {
					var col = uint8(255)
					if state.clockwise && state.isInRowCW(i) {
						col = state.PosCW.Col
					} else if state.isInRowCCW(i) {
						col = state.PosCCW.Col
					} else {
						panic("unreachable encoder state")
					}
					row |= (1 << col)
				}
			}
		}

		// if row is changed, mark it for debouncing
		if m.debouncing[i] != row {
			//changed = true
			//m.rows[i] = row
			m.debouncing[i] = row
			m.debounce = debounceCount
		}
	}

	// if matrix is debouncing, decrement the countdown
	if m.debounce > 0 {
		// decrement the debounce counter
		m.debounce -= 1
		// if still debouncing, wait an interval before returning
		if m.debounce > 0 {
			for start := time.Now(); time.Since(start) < debounceInterval; {
			}
			return
		}
		// if debouncing is complete, update the matrix and mark as changed
		for i, row := range m.debouncing {
			m.rows[i] = row
		}
		changed = true
	}
	return
}
