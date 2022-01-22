package keyboard

import (
	"fmt"
	"io"
	"math/bits"
	"strings"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
	"github.com/bgould/keyboard-firmware/timer"
)

type Keymap [8][16]keycodes.Keycode

func (keymap *Keymap) KeyAt(position Pos) keycodes.Keycode {
	return keymap[position.Row][position.Col]
}

const (
	DebounceMS = 4
	MatrixRows = 8
	MatrixCols = 16
)

type Row uint16

//go:inline
func (r Row) IsOn(col uint8) bool {
	return r&(1<<col) > 0
}

type RowReader interface {
	ReadRow(rowIndex uint8) Row
}

type RowReaderFunc func(rowIndex uint8) Row

func (fn RowReaderFunc) ReadRow(rowIndex uint8) Row {
	return fn(rowIndex)
}

func NewMatrix(io RowReader) *Matrix {
	matrix := &Matrix{io: io}
	return matrix
}

type Matrix struct {
	io         RowReader
	rows       [MatrixRows]Row
	debouncing [MatrixRows]Row
	debounce   uint8
}

func (m *Matrix) Rows() uint8 {
	return MatrixRows
}

func (m *Matrix) Cols() uint8 {
	return MatrixCols
}

//go:inline
func (m *Matrix) GetRow(row uint8) Row {
	return m.rows[row&7]
}

func (m *Matrix) IsOn(row uint8, col uint8) bool {
	return m.GetRow(row).IsOn(col)
}

func (m *Matrix) HasGhostInRow(row uint8) bool {
	r := m.GetRow(row)
	// if there are less than 2 keys down in the row, there is no ghost
	n := 0
	for i := uint8(0); i < MatrixCols; i++ {
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
	for i := uint8(0); i < MatrixRows; i++ {
		if i != row && m.GetRow(i)&r > 0 {
			return true
		}
	}
	return false
}

func (m *Matrix) Scan() (changed bool) {
	// loop over rows and probe the columns for each
	for i := uint8(0); i < MatrixRows; i++ {
		// read row, and if changed check to mark it for debouncing
		if row := m.io.ReadRow(i); m.debouncing[i] != row {
			//changed = true
			//m.rows[i] = row
			m.debouncing[i] = row
			m.debounce = DebounceMS
		}
	}
	// if matrix is debouncing, decrement the countdown
	if m.debounce > 0 {
		// decrement the debounce counter
		m.debounce -= 1
		// if still debouncing, wait an interval before returning
		if m.debounce > 0 {
			timer.Wait(1 * time.Millisecond)
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

func (m *Matrix) Print(w io.Writer) {
	fmt.Fprintf(w, "  0123456789ABCDEF\r\n")
	fmt.Fprintf(w, " +----------------+\r\n")
	for i, row := range m.rows {
		s := fmt.Sprintf("%016b", bits.Reverse16(uint16(row)))
		g := ""
		if m.HasGhostInRow(uint8(i)) {
			g = " <ghost"
		}
		fmt.Fprintf(w, "%X|%s|%s\r\n", byte(i), strings.ReplaceAll(s, "0", "."), g)
	}
	fmt.Fprintf(w, " +----------------+\r\n")
}
