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
	// loop over rows and probe the columns for each
	for i := uint8(0); i < m.nRows; i++ {
		// read row, and if changed check to mark it for debouncing
		if row := m.io.ReadRow(i); m.debouncing[i] != row {
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

/*
func (m *Matrix) Print(w io.Writer) {
		fmt.Fprintf(w, " c0123456\r\n")
		fmt.Fprintf(w, "r+-------+\r\n")
		for i, row := range m.rows {
			s := fmt.Sprintf("%016b", bits.Reverse16(uint16(row)))[0:m.nCols]
			g := ""
			if m.HasGhostInRow(uint8(i)) {
				g = " <ghost"
			}
			fmt.Fprintf(w, "%X|%s|%s\r\n", byte(i), strings.ReplaceAll(s, "0", "."), g)
		}
		fmt.Fprintf(w, " +----------------+\r\n")
}
*/
