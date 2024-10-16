package keyboard

import (
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/console"
)

// var (
// 	commands = console.Commands{}
// )

func (kbd *Keyboard) addDefaultCommands(commands console.Commands) {
	commands["status"] = console.CommandHandlerFunc(kbd.status)
	commands["layer"] = console.CommandHandlerFunc(kbd.layer)
	commands["save"] = console.CommandHandlerFunc(kbd.save)
	commands["load"] = console.CommandHandlerFunc(kbd.load)
	commands["save-macros"] = console.CommandHandlerFunc(kbd.saveMacros)
	commands["load-macros"] = console.CommandHandlerFunc(kbd.loadMacros)
	commands["save-backlight"] = console.CommandHandlerFunc(kbd.saveBacklight)
	commands["load-backlight"] = console.CommandHandlerFunc(kbd.loadBacklight)
	// commands["macros"] = console.CommandHandlerFunc(kbd.xxdmacros)
}

// func initConsole() *console.Console {
// 	return console.New(machine.Serial, commands)
// }

func (kbd *Keyboard) status(cmd console.CommandInfo) int {

	cmd.Stdout.Write([]byte("\n[USB]\n-----\n"))
	cmd.Stdout.Write([]byte("Manufacturer: "))
	cmd.Stdout.Write([]byte(usbManufacturer()))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Product:      "))
	cmd.Stdout.Write([]byte(usbProduct()))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Serial:       "))
	cmd.Stdout.Write([]byte(usbSerial()))
	cmd.Stdout.Write([]byte("\n"))

	/*
		var dataStart, dataEnd [8]byte
		st, en := machine.FlashDataStart(), machine.FlashDataEnd()
		bin2hex([]byte{byte(st >> 24), byte(st >> 16), byte(st >> 8), byte(st)}, dataStart[:])
		bin2hex([]byte{byte(en >> 24), byte(en >> 16), byte(en >> 8), byte(en)}, dataEnd[:])

		cmd.Stdout.Write([]byte("\n[Device]\n--------\n"))
		cmd.Stdout.Write([]byte("Serial Number: "))
		cmd.Stdout.Write([]byte(serialNumber[:]))
		cmd.Stdout.Write([]byte("\n"))
		cmd.Stdout.Write([]byte("Scan Rate:     "))
		cmd.Stdout.Write([]byte(strconv.Itoa(scanRate)))
		cmd.Stdout.Write([]byte("\n"))
		cmd.Stdout.Write([]byte("Flash Start:   "))
		cmd.Stdout.Write(dataStart[:])
		cmd.Stdout.Write([]byte("\n"))
		cmd.Stdout.Write([]byte("Flash End:     "))
		cmd.Stdout.Write(dataEnd[:])
		cmd.Stdout.Write([]byte("\n"))
	*/

	cmd.Stdout.Write([]byte("\n"))

	return 0
}

func (kbd *Keyboard) save(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Saving keymap ...\n"))
	if n, err := kbd.SaveKeymapToFile(savedKeymapFilename); err != nil {
		cmd.Stdout.Write([]byte("Error saving keymap: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Wrote "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Keymap saved successfully.\n"))
		return 0
	}
}

func (kbd *Keyboard) load(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Loading keymap ...\n"))
	if n, err := kbd.LoadKeymapFromFile(savedKeymapFilename); err != nil {
		cmd.Stdout.Write([]byte("Error loading keymap: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Read "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Keymap loaded successfully.\n"))
		return 0
	}
}

func (kbd *Keyboard) saveMacros(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Saving macros...\n"))
	if n, err := kbd.SaveMacrosToFile(savedMacrosFilename); err != nil {
		cmd.Stdout.Write([]byte("Error saving macros: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Wrote "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Macros saved successfully.\n"))
		return 0
	}
}

func (kbd *Keyboard) loadMacros(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Loading macros...\n"))
	if n, err := kbd.LoadMacrosFromFile(savedMacrosFilename); err != nil {
		cmd.Stdout.Write([]byte("Error loading macros: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Read "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Macros loaded successfully.\n"))
		return 0
	}
}

func (kbd *Keyboard) layer(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Active Layer: "))
	cmd.Stdout.Write([]byte(strconv.Itoa(int(kbd.ActiveLayer()))))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Default Layer: "))
	cmd.Stdout.Write([]byte(strconv.Itoa(int(kbd.defaultLayer))))
	cmd.Stdout.Write([]byte("\n"))
	return 0
}

func (kbd *Keyboard) saveBacklight(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Saving backlight...\n"))
	if n, err := kbd.SaveBacklightToFile(savedBacklightFilename); err != nil {
		cmd.Stdout.Write([]byte("Error saving backlight: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Wrote "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Backlight saved successfully.\n"))
		return 0
	}
}

func (kbd *Keyboard) loadBacklight(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Loading backlight...\n"))
	if n, err := kbd.LoadBacklightFromFile(savedBacklightFilename); err != nil {
		cmd.Stdout.Write([]byte("Error loading backlight: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Read "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Backlight loaded successfully.\n"))
		kbd.backlight.sync = true
		return 0
	}
}

// func (kbd *Keyboard) xxdmacros(cmd console.CommandInfo) int {
// 	cmd.Stdout.Write([]byte("Macro count: "))
// 	cmd.Stdout.Write([]byte(strconv.Itoa(int(kbd.macros.count))))
// 	cmd.Stdout.Write([]byte("\n"))
// 	cmd.Stdout.Write([]byte("Macro buffer length: "))
// 	cmd.Stdout.Write([]byte(strconv.Itoa(len(kbd.macros.buffer))))
// 	cmd.Stdout.Write([]byte("\n"))
// 	xxdfprint(cmd.Stdout, 0x0, kbd.macros.buffer[:128])
// 	cmd.Stdout.Write([]byte("\n"))
// 	return 0
// }

// func configureFilesystem() {
// 	if err := board.FS().Mount(); err != nil {
// 		println("Could not mount LittleFS filesystem: ", err.Error(), "\r\n")
// 	} else {
// 		println("Successfully mounted LittleFS filesystem.\r\n")
// 		// fs_mounted = true

// 		if info, err := board.FS().Stat(savedKeymapFilename); err != nil {
// 			println("unable to load ", savedKeymapFilename, ": ", err)
// 		} else {
// 			println("Attempting to load keymap file: ", info.Name())
// 			board.LoadKeymapFromFile(info.Name())
// 		}
// 	}
// }
