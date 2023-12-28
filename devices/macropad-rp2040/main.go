//go:build macropad_rp2040

package main

import (
	"machine"
	"machine/usb"
	"time"

	rotary_encoder "github.com/bgould/keyboard-firmware/drivers/rotary-encoder"
	"github.com/bgould/keyboard-firmware/hosts/multihost"
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	_debug          = false
	encoderInterval = 5 * time.Millisecond
)

var (

	// TODO: encoder API needs to be improved/revamped
	encoder = rotary_encoder.New(machine.ROT_A, machine.ROT_B)

	keymap = Keymap()
	mapper = &MacroPadRP2040KeyMapper{keymap}

	host   = multihost.New(usbvial.New(VialDeviceDefinition, mapper)) // , serial.New(serialer))
	matrix = keyboard.NewMatrix(1, 16, keyboard.RowReaderFunc(ReadRow))

	board = keyboard.New(machine.Serial, host, matrix, keymap)
)

func init() {
	configurePins()
	usb.Serial = vial.MagicSerialNumber("")
	encoder.Configure(rotary_encoder.Config{})
}

func main() {

	// time.Sleep(time.Second)
	serialer.Write([]byte("testing\n"))

	// time.Sleep(2 * time.Second)
	// driver := &vial.Driver{}
	// println("map key test")
	// println(driver.MapKey(0, 0, 0))

	// println("Adding encoders")
	matrix.WithEncoders(
		keyboard.EncoderPos{
			Encoder: encoder,
			PosCW:   keyboard.Pos{Row: 0, Col: encIndexCW},
			PosCCW:  keyboard.Pos{Row: 0, Col: encIndexCCW},
		},
	)

	board.SetDebug(_debug)

	// board.SetEncoders([]keyboard.Encoder{encoder}, keyboard.EncodersSubscriberFunc(encoderCallback))
	// board.SetEncoders([]keyboard.Encoder{encoder}, keyboard.EncodersSubscriberFunc(encoderCallback))

	board.SetKeyAction(keyboard.KeyActionFunc(
		func(key keycodes.Keycode, made bool) {
			if usbvial.UnlockStatus() != vial.UnlockInProgress {
				switch key {
				case keycodes.FN12:
					if made {
						switch board.ActiveLayer() {
						case 0:
							board.SetActiveLayer(1)
						case 1:
							board.SetActiveLayer(0)
						}
						// fmt.Fprintf(serialer, "layer: %d\n", board.ActiveLayer())
					}
				default:
					// fmt.Fprintf(serialer, "fn: %d %t\n", key-keycodes.FN0, made)
				}
			}
		},
	))

	go func() {
		for {
			board.Task()
			time.Sleep(500 * time.Microsecond)
		}
	}()

	for {
		time.Sleep(time.Hour)
	}

}
