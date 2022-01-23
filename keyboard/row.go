package keyboard

type Row uint32

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
