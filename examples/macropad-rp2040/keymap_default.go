//go:build !keymap.custom

package main

import "github.com/bgould/keyboard-firmware/keyboard"

// If the layers slice is empty, the default keymap for the board is loaded
var layers = []keyboard.Layer{}
