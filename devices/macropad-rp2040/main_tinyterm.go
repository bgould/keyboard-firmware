//go:build macropad.tinyterm
// +build macropad.tinyterm

package main

import (
	"image/color"
	"machine"

	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers/sh1106"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/fonts/proggy"
)

const _debug = false

var (
	display  = sh1106.NewSPI(machine.SPI1, machine.OLED_DC, machine.OLED_RST, machine.OLED_CS)
	terminal = tinyterm.NewTerminal(&displayer{&display})
	font     = &proggy.TinySZ8pt7b
	// font = &freemono.Bold12pt7b
)

func configureHost() keyboard.Host {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 48000000,
	})
	display.Configure(sh1106.Config{
		Width:  128,
		Height: 64,
	})
	display.ClearDisplay()
	terminal.Configure(&tinyterm.Config{
		Font:       font,
		FontHeight: 8,
		FontOffset: 6,
	})
	return serial.New(terminal)
}

type displayer struct {
	*sh1106.Device
}

func (d *displayer) FillRectangle(x, y, w, h int16, c color.RGBA) error {
	tinydraw.FilledRectangle(d.Device, x, y, w, h, c)
	return nil
}