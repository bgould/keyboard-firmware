//go:build !oled_featherwing && !macropad_rp2040

package main

import (
	"time"
)

const (
	timeLayout = "3:04 PM"
	dateLayout = "Mon Jan _2, 2006"

	displayTimeout = 5 * time.Second // 1 * time.Minute
)

var (
	lastTime, lastDate string

	lastOn time.Time
)

func initDisplay() error {
	return nil
}

func setDisplay(on bool) {
	return nil
}

func showTime(state *DisplayState, force bool) error {
	tstr := state.ts.Format(timeLayout)
	dstr := state.ts.Format(dateLayout)
	if tstr != lastTime || dstr != lastDate {
		println("tick-tock, it's: " + dstr + " @ " + tstr)
		lastTime = tstr
		lastDate = dstr
	}
	return nil
}
