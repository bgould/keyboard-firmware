package keyboard

type Row uint32

//go:inline
func (r Row) IsOn(col uint8) bool {
	return r&(1<<col) > 0
}
