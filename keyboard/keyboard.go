package keyboard

import (
	"github.com/bgould/keyboard-firmware/keyboard/console"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
	"tinygo.org/x/tinyfs"
)

type Host interface {
	Send(report Report)
	LEDs() uint8
}

type Keyboard struct {
	matrix *Matrix
	keymap Keymap
	host   Host

	prev []Row
	leds uint8

	activeLayer  uint8
	defaultLayer uint8
	layerToggles uint32

	encoders *encoders
	rtc      *rtc

	mouseKeys *MouseKeys

	keyReport      Report
	mouseReport    Report
	consumerReport Report

	eventReceivers []EventReceiver
	keyActionFunc  KeyAction

	enterBootloader EnterBootloaderFunc
	enterCpuReset   EnterBootloaderFunc

	backlight Backlight
	// blState   backlightState

	macros Macros

	fs  tinyfs.Filesystem
	cli *console.Console
}

func New(host Host, matrix *Matrix, keymap Keymap) *Keyboard {
	kbd := &Keyboard{
		// console:   console,
		matrix:         matrix,
		keymap:         keymap,
		host:           host,
		prev:           make([]Row, matrix.Rows()),
		mouseKeys:      NewMouseKeys(DefaultMouseKeysConfig()),
		encoders:       nil,
		eventReceivers: make([]EventReceiver, 0),
	}
	if rcv, ok := host.(EventReceiver); ok {
		kbd.AddEventReceiver(rcv)
	}
	return kbd
}

func (kbd *Keyboard) EnableConsole(serialer Serialer, cmds ...console.Commands) {
	if serialer == nil {
		return
	}
	commands := console.Commands{}
	kbd.addDefaultCommands(commands)
	kbd.addFilesystemCommands(commands)
	for _, set := range cmds {
		for k, v := range set {
			commands[k] = v
		}
	}
	kbd.cli = console.New(serialer, commands)
}

func (kbd *Keyboard) CLI() *console.Console {
	return kbd.cli
}

func (kbd *Keyboard) SetKeyAction(action KeyAction) {
	kbd.keyActionFunc = action
}

func (kbd *Keyboard) AddEventReceiver(receiver EventReceiver) {
	kbd.eventReceivers = append(kbd.eventReceivers, receiver)
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

func (kbd *Keyboard) GetLayerCount() uint8 {
	return kbd.keymap.GetLayerCount()
}

func (kbd *Keyboard) MatrixRows() uint8 {
	return kbd.matrix.Rows()
}

func (kbd *Keyboard) MatrixCols() uint8 {
	return kbd.matrix.Cols()
}

func (kbd *Keyboard) MapKey(layer, row, col int) keycodes.Keycode {
	return kbd.keymap.MapKey(layer, row, col)
}

// TODO: Keep track of "dirty" keys and implement keypress for saving
func (kbd *Keyboard) SetKey(layer, row, col int, kc keycodes.Keycode) bool {
	return kbd.keymap.SetKey(layer, row, col, kc)
}

func (kbd *Keyboard) GetMatrixRowState(idx int) uint32 {
	return uint32(kbd.matrix.GetRow(uint8(idx)))
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
				var handled bool
				var err error
				if len(kbd.eventReceivers) > 0 {
					for _, rcv := range kbd.eventReceivers {
						if rcv == nil {
							continue
						}
						if handled, err = rcv.ReceiveEvent(ev); handled {
							// TODO: handle errors?
							_ = err
							break
						}
					}
				}
				if !handled {
					kbd.processEvent(ev)
					kbd.prev[i] ^= mask
				}
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
	kbd.backlight.Task()
	// if kbd.backlight.Driver != nil {
	// 	kbd.backlight.Driver.Task()
	// }
	if kbd.macros.Driver != nil {
		kbd.macros.Driver.Task(kbd)
	}
	if kbd.cli != nil {
		kbd.cli.Task()
	}
}

func (kbd *Keyboard) processEvent(ev Event) {
	l := kbd.activeLayer
	if int(l) > len(kbd.keymap) {
		l = 0
	}
	key := kbd.keymap.MapKey(int(l), int(ev.Pos.Row), int(ev.Pos.Col))
	// key := kbd.layers[l].KeyAt(ev.Pos)
	for key == keycodes.KC_TRANSPARENT && l > 0 {
		l--
		key = kbd.keymap.MapKey(int(l), int(ev.Pos.Row), int(ev.Pos.Col))
	}
	kbd.ProcessKeycode(key, ev.Made)
}

func (kbd *Keyboard) Modifiers() KeyboardModifier {
	return KeyboardModifier(kbd.keyReport[2])
}

func (kbd *Keyboard) ClearKeycodes() {
	kbd.keyReport.Keyboard(KbdModNone)
	kbd.host.Send(kbd.keyReport)
}

func (kbd *Keyboard) ProcessKeycode(key keycodes.Keycode, made bool) {
	switch {
	case key.IsBasic() || key.IsModifier():
		kbd.processKey(key, made)
	case key.IsMouse():
		kbd.processMouseKey(key, made)
	case key.IsConsumer():
		kbd.processConsumerKey(key, made)
	case key.IsSystem():
		kbd.processSystemKey(key, made)
	case key.IsKb():
		kbd.processKb(key, made)
	case key.IsMacro():
		kbd.processMacro(key, made)
	// case key.IsSpecial():
	// 	kbd.processSpecialKey(key, made)
	case key.IsBacklight():
		kbd.processBacklight(key, made)
	case key.IsRgb():
		kbd.processRgb(key, made)
	case key.IsLayer():
		kbd.processLayer(key, made)
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
