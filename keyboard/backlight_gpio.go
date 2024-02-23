//go:build tinygo

package keyboard

import (
	"context"
	"machine"
	"sync"
	"time"
)

type BacklightGPIO struct {
	LED machine.Pin

	PWM *machine.PWM

	mutex  sync.Mutex
	cancel func()

	state backlightState

	channelLED uint8
}

func (bl *BacklightGPIO) Configure() {
	println("BacklightGPIO.Configure()")
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

func (bl *BacklightGPIO) SetBacklight(mode BacklightMode, level BacklightLevel) {

	bl.mutex.Lock()
	defer bl.mutex.Unlock()

	bl.state.mode, bl.state.level = mode, level
	// println("SetBacklight(): ", bl.state.mode, bl.state.level)

	switch bl.state.mode {

	case BacklightOff:
		bl.cancelIfRunning()
		bl.PWM.Set(bl.channelLED, 0)

	case BacklightOn:
		bl.cancelIfRunning()
		bl.PWM.Set(bl.channelLED, bl.PWM.Top())

	case BacklightBreathing:
		bl.cancelIfRunning()
		ctx, cancel := context.WithCancel(context.Background())
		bl.cancel = cancel
		go bl.breathe(ctx, 4*time.Millisecond)

	}
}

func (bl *BacklightGPIO) cancelIfRunning() {
	if bl.cancel != nil {
		bl.cancel()
		bl.cancel = nil
	}
}

func (bl *BacklightGPIO) breathe(ctx context.Context, interval time.Duration) {
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
			time.Sleep(interval)
		}
	}
}
