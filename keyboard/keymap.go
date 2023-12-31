package keyboard

import (
	"io"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type Layer [][]keycodes.Keycode

func (layer Layer) KeyAt(row, col int) keycodes.Keycode {
	return layer[row][col]
}

type Keymap []Layer

func (keymap Keymap) GetLayerCount() uint8 {
	return uint8(len(keymap))
}

func (keymap Keymap) GetMaxKeyCount() int {
	return keymap.NumRows() * keymap.NumCols()
}

func (keymap Keymap) NumRows() int {
	return len(keymap[0])
}

func (keymap Keymap) NumCols() int {
	return len(keymap[0][0])
}

func (keymap Keymap) MapKey(layer, row, col int) (kc keycodes.Keycode) {
	if uint8(layer) >= keymap.GetLayerCount() || row >= keymap.NumRows() || col >= keymap.NumCols() {
		return
	}
	kc = keymap[layer][row][col]
	// println(layer, idx, row, col, kc)
	return
}

func (keymap Keymap) SetKey(layer, row, col int, kc keycodes.Keycode) bool {
	if uint8(layer) >= keymap.GetLayerCount() || row >= keymap.NumRows() || col >= keymap.NumCols() {
		return false
	}
	keymap[layer][row][col] = kc
	// println(layer, idx, row, col, kc)
	return true
}

func (keymap Keymap) StoredSize() int {
	mapSize := keymap.GetMaxKeyCount() * int(keymap.GetLayerCount()) * 2 // 16-bit keycodes
	hdrSize := 4
	ftrSize := 4
	return hdrSize + mapSize + ftrSize
}

var _ io.WriterTo = (Keymap)(nil)

func (keymap Keymap) ZeroFill() {
	for _, layer := range keymap {
		for _, row := range layer {
			for iCol := range row {
				row[iCol] = 0x0
			}
		}
	}
}

func (keymap Keymap) WriteTo(w io.Writer) (n int64, err error) {
	var result int
	// header
	result, err = w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	n += int64(result)
	if err != nil {
		return
	}
	// loop over layers
	for _, layer := range keymap {
		for _, row := range layer {
			for _, key := range row {
				result, err = w.Write([]byte{byte(key >> 8), byte(key)})
				n += int64(result)
				if err != nil {
					return
				}
			}
		}
	}
	// footer
	result, err = w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	n += int64(result)
	return
}

var _ io.ReaderFrom = (Keymap)(nil)

func (keymap Keymap) ReadFrom(r io.Reader) (n int64, err error) {
	var result int
	var buffer [10]byte
	header := buffer[0:4]
	keybuf := buffer[4:6]
	footer := buffer[6:10]
	// header
	result, err = r.Read(header) // w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	n += int64(result)
	if err != nil {
		return
	}
	// loop over layers
	for _, layer := range keymap {
		for _, row := range layer {
			for i := range row {
				result, err = r.Read(keybuf) // .Write([]byte{byte(key >> 8), byte(key)})
				n += int64(result)
				if err != nil {
					return
				}
				row[i] = keycodes.Keycode(uint16(keybuf[0])<<8 | uint16(keybuf[1]))
			}
		}
	}
	// footer
	result, err = r.Read(footer)
	n += int64(result)
	if err != nil {
		return
	}
	return
}
