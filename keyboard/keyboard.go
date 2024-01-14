package keyboard

import (
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type Host interface {
	Send(report Report)
	LEDs() uint8
}

type Event struct {
	Pos  Pos
	Made bool
	Time uint32
}

type Pos struct {
	// Layer uint8 // TBD
	Row uint8
	Col uint8
}

type Keyboard struct {
	// console Console
	matrix *Matrix
	layers Keymap
	host   Host

	prev []Row
	leds uint8

	activeLayer uint8

	encoders *encoders
	rtc      *rtc

	mouseKeys *MouseKeys

	keyReport      Report
	mouseReport    Report
	consumerReport Report

	// debug bool

	keyActionFunc KeyAction

	enterBootloader EnterBootloaderFunc
	enterCpuReset   EnterBootloaderFunc
}

func New(serial Serialer, host Host, matrix *Matrix, keymap Keymap) *Keyboard {
	return &Keyboard{
		// console:   console,
		matrix:    matrix,
		layers:    keymap,
		host:      host,
		prev:      make([]Row, matrix.Rows()),
		mouseKeys: NewMouseKeys(DefaultMouseKeysConfig()),
		encoders:  nil,
	}
}

func (kbd *Keyboard) SetKeyAction(action KeyAction) {
	kbd.keyActionFunc = action
}

// func (kbd *Keyboard) SetConsole(console Console) {
// 	kbd.console = console
// }

func (kbd *Keyboard) SetDebug(dbg bool) {
	// kbd.debug = dbg
}

func (kbd *Keyboard) SetEnterBootloaderFunc(fn EnterBootloaderFunc) {
	kbd.enterBootloader = fn
}

func (kbd *Keyboard) SetCPUResetFunc(fn EnterBootloaderFunc) {
	kbd.enterCpuReset = fn
}

func (kbd *Keyboard) LEDs() LEDs {
	// return kbd.host.LEDs()
	return LEDs(kbd.leds)
}

func (kbd *Keyboard) SetActiveLayer(index uint8) {
	kbd.activeLayer = index
}

func (kbd *Keyboard) ActiveLayer() uint8 {
	return kbd.activeLayer
}

func (kbd *Keyboard) SetEncoders(encs []Encoder, subscriber EncodersSubscriber) {
	if len(encs) == 0 {
		kbd.encoders = nil
	}
	kbd.encoders = &encoders{
		encoders:   encs,
		subcribers: []EncodersSubscriber{subscriber}, values: make([]int, len(encs)),
	}
}

func (kbd *Keyboard) Task() {
	kbd.rtc.task()
	kbd.matrix.Scan()
	for i, rows := uint8(0), kbd.matrix.Rows(); i < rows; i++ {
		row := kbd.matrix.GetRow(i)
		diff := row ^ kbd.prev[i]
		if diff == 0 {
			continue
		}
		if kbd.matrix.HasGhostInRow(i) {
			continue
		}
		for j, cols := uint8(0), kbd.matrix.Cols(); j < cols; j++ {
			mask := Row(1) << j
			if diff&mask > 0 {
				ev := Event{
					Pos:  Pos{i, j},
					Made: row&mask > 0,
				}
				kbd.processEvent(ev)
				kbd.prev[i] ^= mask
			}
		}
	}
	if newLEDs := kbd.host.LEDs(); newLEDs != kbd.leds {
		// TODO: have some sort of event notification
		// println("LED state changed -", "new: ", newLEDs, "old: ", kbd.leds)
		kbd.leds = newLEDs
	}
	if kbd.mouseKeys.Task(&kbd.mouseReport) {
		kbd.host.Send(kbd.mouseReport)
	}
	if kbd.encoders != nil {
		kbd.encoders.EncodersTask()
	}
}

func (kbd *Keyboard) processEvent(ev Event) {
	l := kbd.activeLayer
	if int(l) > len(kbd.layers) {
		l = 0
	}
	key := kbd.layers.MapKey(int(l), int(ev.Pos.Row), int(ev.Pos.Col))
	// key := kbd.layers[l].KeyAt(ev.Pos)
	for key == keycodes.KC_TRANSPARENT && l > 0 {
		l--
		key = kbd.layers.MapKey(int(l), int(ev.Pos.Row), int(ev.Pos.Col))
	}
	switch {
	case key.IsBasic() || key.IsModifier():
		kbd.processKey(key, ev.Made)
	case key.IsMouse():
		kbd.processMouseKey(key, ev.Made)
	case key.IsConsumer():
		kbd.processConsumerKey(key, ev.Made)
	case key.IsSystem():
		kbd.processSystemKey(key, ev.Made)
	case key.IsKb():
		kbd.processKb(key, ev.Made)
		// case key.IsSpecial():
		// 	kbd.processSpecialKey(key, ev.Made)
	}
}

func (kbd *Keyboard) processKey(key keycodes.Keycode, made bool) {
	if made {
		kbd.keyReport.Make(key)
	} else {
		kbd.keyReport.Break(key)
	}
	kbd.host.Send(kbd.keyReport)
}

func (kbd *Keyboard) processMouseKey(key keycodes.Keycode, made bool) {
	if made {
		kbd.mouseKeys.Make(key)
	} else {
		kbd.mouseKeys.Break(key)
	}
}

func (kbd *Keyboard) processConsumerKey(key keycodes.Keycode, made bool) {
	if made {
		kbd.consumerReport.Make(key)
	} else {
		kbd.consumerReport.Break(key)
	}
	// if kbd.debug {
	// 	kbd.console.Write([]byte("consumer report => " + kbd.consumerReport.String() + "\r\n"))
	// }
	kbd.host.Send(kbd.consumerReport)
}

func (kbd *Keyboard) processSystemKey(key keycodes.Keycode, made bool) {
	// if kbd.debug {
	// 	kbd.console.Write([]byte("system report => " + kbd.consumerReport.String() + "\r\n"))
	// }
}

// func (kbd *Keyboard) processSpecialKey(key keycodes.Keycode, made bool) {
// 	switch {
// 	case key == keycodes.QK_BOOTLOADER:
// 		if !made {
// 			break
// 		}
// 		if kbd.jumpToBootloader != nil {
// 			kbd.jumpToBootloader()
// 		}
// 		return
// 	}
// }

func (kbd *Keyboard) processKb(key keycodes.Keycode, made bool) {
	if !key.IsKb() { // sanity check
		return
	}
	// fnIndex := uint8(key - keycodes.FN0)
	if fn := kbd.keyActionFunc; fn != nil {
		// TODO: should pass *kbd to KeyAction?
		// TODO: consider error reporting
		fn.KeyAction(key, made)
	}
}

func (kbd *Keyboard) GetLayerCount() uint8 {
	return kbd.layers.GetLayerCount()
}

func (kbd *Keyboard) MapKey(layer, row, col int) keycodes.Keycode {
	return kbd.layers.MapKey(layer, row, col)
}

// TODO: Keep track of "dirty" keys and implement keypress for saving
func (kbd *Keyboard) SetKey(layer, row, col int, kc keycodes.Keycode) bool {
	return kbd.layers.SetKey(layer, row, col, kc)
}

func (kbd *Keyboard) GetMatrixRowState(idx int) uint32 {
	return uint32(kbd.matrix.GetRow(uint8(idx)))
}
