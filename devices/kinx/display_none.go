//go:build !oled_featherwing

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

func displayTask() {
}

func setDisplay(on bool) {
}

func showTime(state DisplayState, force bool) error {
	tstr := state.ts.Format(timeLayout)
	dstr := state.ts.Format(dateLayout)
	compare := state
	compare.scanRate = 0
	if state != lastState { // TODO: normalize scan rate so conditional statement below isn't necessary?
		if tstr != lastTime || dstr != lastDate {
			cli.WriteString(
				"Time: " + dstr + " @ " + tstr + "; " +
					"TOTP Account: " + state.totpAccount + "; " +
					"TOTP Numbers: " + state.totpNumbers,
			)
			lastTime = tstr
			lastDate = dstr
		}
		lastState = state
	}
	return nil
}
