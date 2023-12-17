//go:build vial

package main

import keyboard "github.com/sago35/tinygo-keyboard"

func loadKeyboardDef() {
	keyboard.KeyboardDef = []byte{
		0x5D, 0x00, 0x00, 0x80, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x3D, 0x82, 0x80, 0x17, 0x1C, 0x2E, 0x8B, 0x89, 0x9F, 0x24, 0xFF, 0xE8, 0x0F, 0x33, 0xF8, 0x2A, 0x14, 0xC7, 0x6D, 0x7F, 0x2D, 0xC4, 0xC3, 0x46, 0xE8, 0x1C, 0x95, 0xE2, 0x58, 0x92, 0x71, 0xD2, 0xA9, 0xA3, 0x69, 0xE7, 0x68, 0x37, 0x14, 0x67, 0x5D, 0x60, 0xEF, 0x45, 0x0E, 0xF7, 0x03, 0x1C, 0x24, 0x1B, 0xBE, 0xF9, 0x1B, 0xB3, 0xBB, 0x99, 0xB8, 0xD2, 0xE7, 0xF8, 0xB7, 0x8F, 0x1F, 0x71, 0xFC, 0x3B, 0x68, 0x86, 0x82, 0x8C, 0xC9, 0x5E, 0xB7, 0x38, 0x3D, 0x01, 0xB9, 0x67, 0x54, 0xF4, 0x38, 0xD9, 0x95, 0x06, 0x53, 0x90, 0x37, 0xF3, 0xE1, 0x52, 0x0E, 0x5C, 0xE0, 0xBD, 0x6E, 0xB1, 0xCD, 0x88, 0x96, 0x5A, 0xEC, 0xDE, 0x7D, 0x4A, 0x3D, 0x67, 0xDF, 0xDC, 0xB5, 0x00, 0x0A, 0x15, 0x53, 0xB6, 0xEC, 0x44, 0x49, 0x86, 0xAA, 0x3C, 0x20, 0xE7, 0xF0, 0x65, 0xD4, 0x02, 0xE3, 0x79, 0x4A, 0x6C, 0x87, 0xB5, 0x06, 0xBD, 0x7A, 0x58, 0x9D, 0x42, 0x3C, 0x17, 0xB7, 0x3B, 0x2D, 0x5F, 0x06, 0xF3, 0x86, 0x2D, 0x74, 0x6C, 0x22, 0x48, 0x4B, 0x84, 0xD7, 0xF3, 0xD5, 0xFF, 0xB2, 0x0C, 0x1C, 0xED, 0x83, 0x70, 0x7F, 0x9D, 0x9D, 0xDC, 0xC3, 0x97, 0x16, 0x8C, 0x5F, 0xE2, 0x62, 0xC1, 0x60, 0x02, 0x66, 0x75, 0x56, 0x3B, 0x1D, 0xF4, 0x82, 0x14, 0x35, 0x0C, 0xA9, 0x85, 0x40, 0x48, 0x19, 0x0E, 0xFF, 0x65, 0x11, 0xF8, 0x4F, 0x6F, 0x5F, 0x5E, 0x9A, 0xFD, 0xB7, 0x5C, 0x4E, 0x81, 0xF8, 0x57, 0xCA, 0x03, 0x51, 0xF7, 0x06, 0x86, 0xF3, 0x6E, 0x7D, 0x87, 0x3A, 0x88, 0x2D, 0x65, 0x85, 0xF2, 0x46, 0xE4, 0x5A, 0xD8, 0x10, 0x85, 0xA1, 0x5C, 0x99, 0x34, 0xBC, 0x93, 0x03, 0x10, 0x4C, 0x4F, 0x28, 0xAB, 0xAC, 0x3E, 0xC8, 0xBD, 0xD2, 0x30, 0xA8, 0xB1, 0x9E, 0xD6, 0xF0, 0x56, 0x7A, 0x7D, 0x33, 0x38, 0x09, 0x41, 0xBA, 0x8B, 0x21, 0x01, 0xC2, 0x55, 0x26, 0x0D, 0xEA, 0xAD, 0xC8, 0xD4, 0x15, 0xCB, 0xA4, 0x0E, 0x50, 0x21, 0xD7, 0x83, 0xBD, 0x36, 0xAF, 0x54, 0xD3, 0x41, 0xF7, 0x06, 0x9A, 0x7B, 0x9C, 0xEF, 0xA4, 0x8D, 0x63, 0xDF, 0xA0, 0x1C, 0x8F, 0x9F, 0x9E, 0xAD, 0xF3, 0x32, 0xDC, 0x8A, 0xA6, 0xFC, 0xA6, 0x02, 0xC1, 0xDF, 0xCA, 0xF7, 0x70, 0xA3, 0x6A, 0x09, 0x7C, 0xFB, 0x5A, 0xE1, 0x17, 0x56, 0x62, 0x7B, 0xED, 0x39, 0x44, 0x74, 0x14, 0x30, 0x8E, 0xF3, 0x85, 0x1D, 0x4C, 0x28, 0x93, 0x32, 0xA6, 0xFF, 0x47, 0x65, 0xC8, 0x27, 0x5B, 0x5E, 0x19, 0xA6, 0xCD, 0x9B, 0xF2, 0x3D, 0xDF, 0xB0, 0x99, 0x4A, 0x34, 0xCD, 0xA6, 0xBE, 0x94, 0x60, 0xB8, 0x8C, 0x84, 0x46, 0x9B, 0x6D, 0x18, 0xBA, 0x74, 0x89, 0x67, 0x7F, 0xEA, 0xB0, 0xB6, 0x8A, 0xAA, 0x4B, 0xBA, 0x07, 0x1E, 0xD4, 0xC6, 0x0C, 0xC4, 0x98, 0xB8, 0x1D, 0x83, 0x41, 0xE2, 0xC8, 0x03, 0x7E, 0x32, 0xB4, 0x93, 0x20, 0x8B, 0x72, 0x14, 0x65, 0x4C, 0xCB, 0x10, 0xD0, 0x40, 0x90, 0xB1, 0x6D, 0x0F, 0xFE, 0x6C, 0x89, 0x6C, 0x99, 0x52, 0xBF, 0x3F, 0x69, 0xEC, 0x21, 0xC9, 0x30, 0xF8, 0xE8, 0x7E, 0xE2, 0x8B, 0xF1, 0x90, 0x23, 0x4C, 0x96, 0x1A, 0x0A, 0x99, 0x4C, 0x36, 0x6E, 0xF2, 0xAD, 0xC7, 0x49, 0xE3, 0x84, 0x48, 0x5D, 0x52, 0xAA, 0x84, 0xF0, 0x91, 0x04, 0x38, 0xE4, 0xCF, 0x34, 0xC3, 0xAB, 0xEA, 0x01, 0x1B, 0x50, 0x8B, 0x8C, 0xFF, 0x55, 0x10, 0xF8, 0x75, 0x6E, 0xBB, 0x93, 0xB1, 0x2F, 0x64, 0x29, 0x4E, 0x83, 0xAF, 0xF5, 0xF6, 0xA0, 0xD8, 0x19, 0xF6, 0xC9, 0x67, 0xD2, 0x8E, 0x54, 0x1D, 0xB6, 0xB8, 0xFF, 0xFF, 0x79, 0x34, 0x92, 0x00,
	}
}
