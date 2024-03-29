// This file is auto-generated by github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def, do not edit.
package circuitplay_bluefruit

import "github.com/bgould/keyboard-firmware/hosts/usbvial/vial"

var VialDeviceDefinition = vial.DeviceDefinition{
	Name:      "tinygo-circuitplay-bluefruit",
	VendorID:  "0x239A",
	ProductID: "0x8045",
	Matrix: vial.DeviceMatrix{
		Rows: 1,
		Cols: 2,
	},
	LzmaDefLength: uint16(len(vialLzmaDefinition)),
	LzmaDefWriter: vial.LzmaDefPageWriterFunc(vialWriteLzmaDefPage),
}

var vialLzmaDefinition = []byte{
	0x5D, 0x00, 0x00, 0x80, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x3D, 0x82, 0x80, 0x17, 0x1C, 0x2E, 0x8B, 0x89, 0x9F, 0x24, 0xFF, 0xE8, 0x0F, 0x34, 0xC8, 0x2A, 0x8F, 0x45, 0xBA, 0x46, 0xF9, 0xCF, 0x4D, 0x70, 0xC4, 0xC0, 0x48, 0x67, 0x7A, 0x59, 0xDB, 0xFF, 0x05, 0x1F, 0x41, 0x05, 0xC2, 0x83, 0xAF, 0x02, 0xA5, 0x51, 0xB7, 0x97, 0x4C, 0xDE, 0xF3, 0x54, 0xA7, 0x4E, 0x16, 0x0B, 0x4B, 0xFA, 0xFF, 0xD6, 0x1D, 0xA1, 0x58, 0x47, 0xFA, 0x88, 0xD7, 0x1A, 0x53, 0x9D, 0x7E, 0x4E, 0x6D, 0xC4, 0xDE, 0x1C, 0xFE, 0x1E, 0x42, 0x9B, 0x71, 0xA7, 0x38, 0x2A, 0xA0, 0x98, 0x33, 0xDD, 0x34, 0x51, 0x1B, 0x87, 0xD2, 0xAF, 0xEE, 0x2C, 0x60, 0xBB, 0x48, 0x9B, 0x38, 0x0D, 0x44, 0xF0, 0xCC, 0x0B, 0xC1, 0x70, 0x1E, 0x34, 0x92, 0xEE, 0xDC, 0xD1, 0xD2, 0xC4, 0xBA, 0xD1, 0x48, 0xD9, 0x7A, 0xE7, 0xE1, 0x5C, 0x66, 0x23, 0xA9, 0x7A, 0xBA, 0xFD, 0xDB, 0x14, 0x00, 0x59, 0x08, 0x84, 0x8A, 0xCA, 0x25, 0x45, 0x08, 0x78, 0xE9, 0x41, 0xFA, 0xAF, 0x5C, 0x43, 0x82, 0x9B, 0xEE, 0xB0, 0x88, 0x26, 0xE4, 0x83, 0xA7, 0x9C, 0x35, 0xE9, 0x70, 0xC8, 0x73, 0xFD, 0x62, 0x4B, 0xA0,
}

func vialWriteLzmaDefPage(tx []byte, page uint16) bool {
	start := page * 32
	end := start + 32
	len := uint16(len(vialLzmaDefinition))
	if end < start || start >= len {
		return false
	}
	if end > len {
		end = len
	}
	copy(tx[:32], vialLzmaDefinition[start:end])
	return true
}
