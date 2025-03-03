package keyboard

import (
	"io"
	"os"
	"strings"

	"github.com/bgould/keyboard-firmware/keyboard/console"
	"tinygo.org/x/tinyfs"
)

type FsErr uint8

const (
	ErrNoFS FsErr = iota
	ErrNotAFile
)

func (err FsErr) Error() string {
	switch err {
	case ErrNoFS:
		return "not configured"
	case ErrNotAFile:
		return "not a file"
	default:
		return "unknown"
	}
}

func (kbd *Keyboard) SetFS(fs tinyfs.Filesystem) {
	kbd.fs = fs
}

func (kbd *Keyboard) FS() tinyfs.Filesystem {
	return kbd.fs
}

const (
	savedKeymapFilename    = "saved.keymap"
	savedMacrosFilename    = "saved.macros"
	savedBacklightFilename = "saved.backlight"
)

func (kbd *Keyboard) ConfigureFilesystem() (err error) {
	fs := kbd.FS()
	if fs == nil {
		return
	}
	// FIXME: swallowing error here because tinyfs keeps returning "invalid parameter"
	_ = fs.Mount()
	// if err := fs.Mount(); err != nil {
	// 	println("Could not mount LittleFS filesystem: ", err.Error(), "\r\n")
	// 	return err
	// } else {
	println("Successfully mounted LittleFS filesystem.\r\n")

	// FIXME: consolidate/standardize with block below
	// attempt to load saved keymap
	if info, err := fs.Stat(savedKeymapFilename); err != nil {
		println("unable to load ", savedKeymapFilename, ": ", err)
		// return err
	} else {
		println("Attempting to load keymap file: ", info.Name())
		_, err := kbd.LoadKeymapFromFile(info.Name())
		if err != nil {
			println("error loading keymap file: ", err) //return err
		}
	}

	// FIXME: consolidate/standardize with block above
	// attempt to load saved macros
	if info, err := fs.Stat(savedMacrosFilename); err != nil {
		println("unable to load ", savedMacrosFilename, ": ", err)
		// return err
	} else {
		println("Attempting to load macros file: ", info.Name())
		_, err := kbd.LoadMacrosFromFile(info.Name())
		if err != nil {
			println("error loading macros file: ", err) //return err
		}
	}

	// FIXME: consolidate/standardize with block above
	// attempt to load saved macros
	if info, err := fs.Stat(savedBacklightFilename); err != nil {
		println("unable to load ", savedBacklightFilename, ": ", err)
		// return err
	} else {
		println("Attempting to load backlight file: ", info.Name())
		_, err := kbd.LoadBacklightFromFile(info.Name())
		if err != nil {
			println("error loading backlight file: ", err) //return err
		}
	}
	kbd.backlight.sync = true

	return nil
}

func (kbd *Keyboard) saveToFile(filename string, obj io.WriterTo) (n int64, err error) {
	f, err := kbd.fs.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err = obj.WriteTo(f)
	return n, err
}

// LoadKeymapFromFile updates the current in-memory keymap from the filesystem
func (kbd *Keyboard) loadFromFile(filename string, obj io.ReaderFrom) (n int64, err error) {
	f, err := kbd.fs.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	if f.IsDir() {
		return 0, ErrNotAFile
	}
	n, err = obj.ReadFrom(f)
	return
}

// SaveKeymapToFile write the current in-memory keymap to the filesystem
func (kbd *Keyboard) SaveKeymapToFile(filename string) (n int64, err error) {
	return kbd.saveToFile(filename, kbd.keymap)
}

// LoadKeymapFromFile updates the current in-memory keymap from the filesystem
func (kbd *Keyboard) LoadKeymapFromFile(filename string) (n int64, err error) {
	return kbd.loadFromFile(filename, kbd.keymap)
}

// SaveMacrosToFile write the current in-memory macros to the filesystem
func (kbd *Keyboard) SaveMacrosToFile(filename string) (n int64, err error) {
	if !kbd.macros.Enabled() {
		return 0, nil
	}
	return kbd.saveToFile(filename, kbd.macros.Driver)
}

// LoadMacrosFromFile updates the current in-memory macros from the filesystem
func (kbd *Keyboard) LoadMacrosFromFile(filename string) (n int64, err error) {
	if !kbd.macros.Enabled() {
		return 0, nil
	}
	return kbd.loadFromFile(filename, kbd.macros.Driver)
}

// SaveMacrosToFile write the current in-memory macros to the filesystem
func (kbd *Keyboard) SaveBacklightToFile(filename string) (n int64, err error) {
	if !kbd.BacklightEnabled() {
		return 0, nil
	}
	return kbd.saveToFile(filename, &kbd.backlight.state)
}

// LoadMacrosFromFile updates the current in-memory macros from the filesystem
func (kbd *Keyboard) LoadBacklightFromFile(filename string) (n int64, err error) {
	if !kbd.macros.Enabled() {
		return 0, nil
	}
	return kbd.loadFromFile(filename, &kbd.backlight.state)
}

// ########################### Filesystem Commands ###########################/

func (kbd *Keyboard) addFilesystemCommands(commands console.Commands) {
	// Filesystem Commands
	commands["mount"] = console.CommandHandlerFunc(kbd.mount)
	commands["umount"] = console.CommandHandlerFunc(kbd.umount)
	commands["format"] = console.CommandHandlerFunc(kbd.format)
	commands["ls"] = console.CommandHandlerFunc(kbd.ls)
	commands["mv"] = console.CommandHandlerFunc(kbd.mv)
	commands["rm"] = console.CommandHandlerFunc(kbd.rm)
	commands["cat"] = console.CommandHandlerFunc(kbd.cat)
	commands["mkdir"] = console.CommandHandlerFunc(kbd.mkdir)
}

func (kbd *Keyboard) mount(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	if err := kbd.FS().Mount(); err != nil {
		println("Could not mount LittleFS filesystem: ", err.Error(), "\r\n")
		return 1
	} else {
		println("Successfully mounted LittleFS filesystem.\r\n")
		return 0
	}
}

func (kbd *Keyboard) format(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	if err := kbd.FS().Format(); err != nil {
		println("Could not format LittleFS filesystem: ", err.Error(), "\r\n")
		return 1
	} else {
		println("Successfully formatted LittleFS filesystem.\r\n")
		return 0
	}
}

func (kbd *Keyboard) umount(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	if err := kbd.FS().Unmount(); err != nil {
		println("Could not unmount LittleFS filesystem: ", err.Error(), "\r\n")
		return 1
	} else {
		println("Successfully unmounted LittleFS filesystem.\r\n")
		return 0
	}
}

func (kbd *Keyboard) ls(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	path := "/"
	argv := cmd.Argv
	if len(argv) > 0 {
		path = strings.TrimSpace(argv[0])
	}
	dir, err := kbd.FS().Open(path)
	if err != nil {
		println("Could not open directory", path, ":", err.Error())
		return 1
	}
	defer dir.Close()
	infos, err := dir.Readdir(0)
	_ = infos
	if err != nil {
		println("Could not read directory", path, ":", err.Error())
		return 1
	}
	for _, info := range infos {
		s := "-rwxrwxrwx"
		if info.IsDir() {
			s = "drwxrwxrwx"
		}
		println(s, info.Size(), info.Name())
		//fmt.Printf("%s %5d %s\n", s, info.Size(), info.Name())
	}
	return 0
}

func (kbd *Keyboard) mkdir(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	tgt := ""
	argv := cmd.Argv
	if len(argv) == 1 {
		tgt = strings.TrimSpace(argv[0])
	}
	if tgt == "" {
		println("Usage: mkdir <target dir>")
		return 1
	}
	err := kbd.FS().Mkdir(tgt, 0777)
	if err != nil {
		println("Could not mkdir " + tgt + ": " + err.Error())
	}
	return 0
}

func (kbd *Keyboard) rm(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	tgt := ""
	argv := cmd.Argv
	if len(argv) == 1 {
		tgt = strings.TrimSpace(argv[0])
	}
	if tgt == "" {
		println("Usage: rm <target dir>")
		return 1
	}
	err := kbd.FS().Remove(tgt)
	if err != nil {
		println("Could not rm ", tgt, ":", err.Error())
		return 1
	}
	return 0
}

func (kbd *Keyboard) mv(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	src, dst := "", ""
	argv := cmd.Argv
	if len(argv) == 2 {
		src = strings.TrimSpace(argv[0])
		dst = strings.TrimSpace(argv[1])
	}
	if src == "" || dst == "" {
		println("Usage: mv <srcfile> <destfile>")
		return 1
	}
	err := kbd.FS().Rename(src, dst)
	if err != nil {
		println("Could not mv ", src, "to", dst, ":", err.Error())
		return 1
	}
	return 0
}

func (kbd *Keyboard) cat(cmd console.CommandInfo) int {
	if kbd.FS() == nil {
		println("No filesystem available")
		return 1
	}
	tgt := ""
	argv := cmd.Argv
	if len(argv) == 1 {
		tgt = strings.TrimSpace(argv[0])
	}
	if tgt == "" {
		println("Usage: cat <target dir>")
		return 1
	}
	f, err := kbd.FS().Open(tgt)
	if err != nil {
		println("Could not open: " + err.Error())
		return 1
	}
	defer f.Close()
	if f.IsDir() {
		println("Not a file: " + tgt)
		return 1
	}
	off := 0x0
	buf := make([]byte, 64)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			println("Error reading " + tgt + ": " + err.Error())
		}
		xxdfprint(os.Stdout, uint32(off), buf[:n])
		off += n
	}
	return 0
}
