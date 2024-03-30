//go:build tinygo && nrf

package keyboard

import (
	"context"
	"machine"
	"sync"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/hsv"
)

type BacklightGPIO struct {

	//
	LED machine.Pin

	//
	PWM *machine.PWM

	//
	Interval time.Duration

	mutex  sync.Mutex
	cancel func()

	state backlightState

	brightening bool
	step        uint8
	last        time.Time

	channelLED uint8
}

func (bl *BacklightGPIO) Configure() {
	if bl.Interval == 0 {
		bl.Interval = 4 * time.Millisecond
	}
	err := bl.PWM.Configure(machine.PWMConfig{})
	if err != nil {
		println("failed to configure PWM")
		return
	}
	bl.channelLED, err = bl.PWM.Channel(bl.LED)
	if err != nil {
		println("failed to configure LED PWM channel")
		return
	}
}

func (bl *BacklightGPIO) Task() {

	if time.Since(bl.last) < bl.Interval {
		return
	}

	bl.mutex.Lock()
	defer bl.mutex.Unlock()

	if bl.state.mode != BacklightBreathing {
		return
	}

	if bl.step == 0 {
		bl.brightening = !bl.brightening
	}

	var brightness uint32 = uint32(bl.step)
	if !bl.brightening {
		brightness = 256 - brightness
	}
	bl.PWM.Set(bl.channelLED, bl.PWM.Top()*brightness/256)
	bl.last = bl.last.Add(bl.Interval)
	bl.step++
}

func (bl *BacklightGPIO) SetBacklight(mode BacklightMode, color hsv.Color) {

	bl.mutex.Lock()
	defer bl.mutex.Unlock()

	if mode == bl.state.mode && color == bl.state.color {
		return
	}

	bl.state.mode, bl.state.color = mode, color
	// println("SetBacklight(): ", bl.state.mode, bl.state.level)

	switch bl.state.mode {

	case BacklightOff:
		// println("BacklightOff")
		bl.cancelIfRunning()
		bl.PWM.Set(bl.channelLED, 0)

	case BacklightOn:
		// println("BacklightOn")
		bl.cancelIfRunning()
		bl.PWM.Set(bl.channelLED, bl.PWM.Top())

	case BacklightBreathing:
		// println("BacklightBreathing")
		bl.cancelIfRunning()
		bl.brightening = false
		bl.step = 0xF
		bl.last = time.Now()
		// if false {
		// ctx, cancel := context.WithCancel(context.Background())
		// bl.cancel = cancel
		// go bl.breathe(ctx)
		// } else {

		// }
	}
}

func (bl *BacklightGPIO) breathe(ctx context.Context) {
	for i, brightening := uint8(0xF), false; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			if i == 0 {
				brightening = !brightening
				continue
			}
			var brightness uint32 = uint32(i)
			if !brightening {
				brightness = 256 - brightness
			}
			bl.PWM.Set(bl.channelLED, bl.PWM.Top()*brightness/256)
			time.Sleep(bl.Interval)
		}
	}
}

func (bl *BacklightGPIO) cancelIfRunning() {
	if bl.cancel != nil {
		bl.cancel()
		bl.cancel = nil
	}
}
