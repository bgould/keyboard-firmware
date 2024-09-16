//go:build macropad_rp2040 && !backlight.none

package main

import (
	"image/color"
	"machine"
	"os"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/console"
	"github.com/bgould/keyboard-firmware/keyboard/hsv"
	"tinygo.org/x/drivers/ws2812"
)

func init() {
	backlight := keyboard.Backlight{
		Driver: &keyboard.BacklightColorStrip{
			ColorStrip: &keyboard.ColorStrip{
				Writer: ws2812.NewWS2812(machine.NEOPIXEL),
				Pixels: make([]color.RGBA, 12),
			},
			Interval: 6 * time.Millisecond,
		},
		Steps: 16,
	}
	kbd.SetBacklight(backlight)
	kbd.BacklightDriver().Configure()
}

func addBacklightCommands(commands console.Commands) {
	commands["save-backlight"] = console.CommandHandlerFunc(saveBacklight)
	commands["load-backlight"] = console.CommandHandlerFunc(loadBacklight)
}

func saveBacklight(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	if n, err := saveBacklightColorToFile(savedColorFilename); err != nil {
		println("Could not save backlight color: ", err.Error(), "\r\n")
		return 1
	} else {
		_ = n
		println("Successfully saved backlight color.\r\n")
		return 0
	}
}

func loadBacklight(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	if col, err := loadBacklightColorFromFile(savedColorFilename); err != nil {
		println("Could not load backlight color: ", err.Error(), "\r\n")
		return 1
	} else {
		// drv := backlightColorStrip()
		// kbd.BacklightDriver().SetBacklight(drv.GetMode(), col)
		kbd.BacklightUpdate(kbd.BacklightMode(), col)
		println("Successfully loaded backlight color.\r\n")
		return 0
	}
}

// func backlightColorStrip() *keyboard.BacklightColorStrip {
// 	return kbd.BacklightDriver().(*keyboard.BacklightColorStrip)
// }

const savedColorFilename = "backlight_color.hsv"

func loadSavedBacklightColor() (hsv.Color, error) {
	if fs := kbd.FS(); fs != nil {
		if info, err := fs.Stat(savedColorFilename); err != nil {
			return hsv.Color{}, err
			// println("unable to load ", savedColorFilename, ": ", err)
		} else {
			// println("Attempting to load backlight color file: ", info.Name())
			_, err := loadBacklightColorFromFile(info.Name())
			return hsv.Color{}, err
		}
	}
	return hsv.Color{}, nil
}

// saveBacklightColorToFile write the current in-memory keymap to the filesystem
func saveBacklightColorToFile(filename string) (n int64, err error) {
	color := kbd.BacklightColor() // backlightColorStrip().GetColor()
	f, err := kbd.FS().OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	b, err := f.Write([]byte{color.H, color.S, color.V})
	return int64(b), err
}

// loadBacklightColorFromFile updates the current in-memory state from the filesystem
func loadBacklightColorFromFile(filename string) (col hsv.Color, err error) {
	f, err := kbd.FS().Open(filename)
	if err != nil {
		return col, err
	}
	defer f.Close()
	if f.IsDir() {
		return col, keyboard.ErrNotAFile
	}
	var buf [3]byte
	b, err := f.Read(buf[:])
	if err == nil && b == len(buf) {
		col.H = buf[0]
		col.S = buf[1]
		col.V = buf[2]
	}
	return
}
