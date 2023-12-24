package main

import "github.com/bgould/keyboard-firmware/hosts/usbvial/vial"

func loadKeyboardDef() {
	vial.KeyboardDef = []byte{
		0x5D, 0x00, 0x00, 0x80, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x3D, 0x82, 0x80, 0x17, 0x1C, 0x2E, 0x8B, 0x89, 0x9F, 0x24, 0xFF, 0xE8, 0x0F, 0x33, 0xF8, 0x2A, 0x14, 0xC7, 0x6D, 0x7F, 0x2D, 0xC4, 0xC3, 0x46, 0xE8, 0x1C, 0x95, 0xE2, 0x58, 0x92, 0x71, 0xD2, 0xA9, 0xA3, 0x69, 0xE7, 0x68, 0x37, 0x14, 0x67, 0x5D, 0x60, 0xEF, 0x45, 0x0E, 0xF7, 0x03, 0x1C, 0x24, 0x1B, 0xBE, 0xF9, 0x1B, 0xB3, 0xBB, 0x99, 0xB8, 0xD2, 0xE7, 0xF8, 0xB7, 0x8F, 0x1F, 0x71, 0xFB, 0x79, 0x9A, 0x66, 0x50, 0x2C, 0xF9, 0x12, 0xFA, 0x18, 0xB0, 0x86, 0xF1, 0xB9, 0x4B, 0xAE, 0xC6, 0x3D, 0x93, 0x3F, 0xC4, 0x4B, 0x0B, 0xE7, 0xC1, 0x2B, 0x86, 0x55, 0xD1, 0xDE, 0xA5, 0x83, 0xD3, 0x69, 0xDB, 0xF4, 0xB7, 0x38, 0xC5, 0x64, 0xBF, 0xB4, 0xAD, 0x20, 0x1F, 0x5C, 0x16, 0x67, 0xF6, 0x8E, 0x03, 0x2D, 0xC9, 0xF2, 0x76, 0xEC, 0x5E, 0x5F, 0x01, 0x1D, 0xED, 0x8D, 0x65, 0x94, 0x39, 0x80, 0xC7, 0x17, 0x4E, 0xF8, 0xE7, 0xE3, 0xF8, 0x5C, 0xB0, 0x84, 0x88, 0xFF, 0x26, 0x9F, 0x2F, 0x67, 0x05, 0x6F, 0x3D, 0xB6, 0x2F, 0x70, 0xF8, 0x44, 0x40, 0x17, 0xB8, 0x1D, 0xF2, 0xEC, 0xC3, 0xC5, 0xCB, 0x8A, 0x26, 0x77, 0x82, 0x5D, 0x60, 0x0B, 0x42, 0xB1, 0x94, 0xC2, 0x12, 0xF0, 0x0C, 0x18, 0xCB, 0xAB, 0x11, 0x02, 0x59, 0x13, 0xF1, 0x91, 0x9D, 0xA1, 0x7D, 0x5C, 0xA7, 0xAF, 0x04, 0xA9, 0x0D, 0x23, 0xCD, 0x28, 0x91, 0x61, 0x6A, 0xC0, 0x18, 0xAB, 0x4C, 0x7E, 0x55, 0xD9, 0x89, 0x2A, 0x02, 0x2A, 0xB5, 0x9C, 0x40, 0x3A, 0x58, 0xAD, 0x23, 0xDE, 0x2C, 0x73, 0x48, 0xD4, 0x75, 0x63, 0x38, 0xAC, 0x85, 0xA5, 0x8B, 0x72, 0x48, 0x34, 0x0F, 0xE0, 0xB1, 0x4C, 0x97, 0x47, 0x96, 0x61, 0x25, 0xFD, 0x74, 0xFA, 0xD0, 0x83, 0x19, 0x13, 0x94, 0xD6, 0xAA, 0x31, 0xCE, 0xA0, 0x34, 0x49, 0x75, 0x36, 0xB9, 0x95, 0xC7, 0x03, 0xC7, 0xE0, 0x59, 0xC4, 0x1D, 0xB3, 0x48, 0x1E, 0xE4, 0x7A, 0x30, 0x53, 0x8F, 0xC8, 0x20, 0xA9, 0x55, 0x19, 0x2D, 0x17, 0xFB, 0xEF, 0x9F, 0xF0, 0xCE, 0x5B, 0xFB, 0x96, 0xC0, 0xA6, 0xA2, 0xED, 0xC7, 0x18, 0x19, 0xF6, 0x59, 0x5F, 0x3D, 0x6C, 0x1F, 0xD6, 0xA0, 0x1B, 0xE7, 0xBC, 0xA2, 0xD7, 0x89, 0xB6, 0x87, 0xDE, 0xD7, 0x6C, 0x0D, 0xF2, 0x44, 0x9D, 0xBC, 0x9F, 0xD9, 0x3C, 0xCE, 0x63, 0x7C, 0xC0, 0x71, 0x9E, 0xF9, 0x93, 0xDD, 0xFD, 0x16, 0xA0, 0xAB, 0xC4, 0x32, 0xCF, 0xA8, 0xB9, 0x79, 0x2F, 0xE4, 0x35, 0x10, 0x34, 0xB8, 0x81, 0xCE, 0xF3, 0xB1, 0x21, 0x88, 0xD1, 0x02, 0x58, 0x35, 0xE2, 0xAD, 0x16, 0x8F, 0x4B, 0xB1, 0xE1, 0xC6, 0xAF, 0xDD, 0xCB, 0xFB, 0x1A, 0x45, 0x82, 0x49, 0x72, 0x0E, 0xFB, 0xE8, 0x1C, 0x37, 0x6B, 0xCF, 0x4B, 0x6E, 0xA2, 0x29, 0xC1, 0x6F, 0xCF, 0x36, 0xE7, 0x7C, 0x32, 0xCF, 0x5F, 0x3F, 0x4B, 0x3B, 0x26, 0xE7, 0x0A, 0x58, 0x27, 0x71, 0x40, 0x9C, 0x55, 0x49, 0x60, 0x51, 0xE0, 0x91, 0x6A, 0xF5, 0xDC, 0x44, 0x6A, 0x41, 0x51, 0x1C, 0x68, 0x5E, 0x96, 0xC1, 0x61, 0x31, 0x84, 0x56, 0xAE, 0x2F, 0xF8, 0x54, 0x5E, 0x67, 0x47, 0x43, 0xF5, 0xED, 0xBB, 0xE8, 0xF9, 0x1F, 0x97, 0xB1, 0xAB, 0x24, 0xAF, 0x57, 0x6C, 0xB8, 0xB2, 0xDF, 0xB2, 0xD8, 0x8D, 0xA1, 0xA2, 0x5B, 0xA9, 0xA7, 0xE6, 0xB2, 0x4E, 0xFA, 0x75, 0x01, 0xFD, 0x03, 0x24, 0xE6, 0x5F, 0x36, 0xB4, 0xFB, 0xCA, 0x55, 0x20, 0x15, 0xED, 0xD6, 0x4C, 0xD8, 0x50, 0x11, 0xC0, 0xA3, 0x34, 0x77, 0x33, 0x4B, 0x3A, 0x4A, 0x7D, 0xD4, 0xBE, 0x11, 0x85, 0xFF, 0xFF, 0xCF, 0x87, 0x1F, 0x92,
	}
}
