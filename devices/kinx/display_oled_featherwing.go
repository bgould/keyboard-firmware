// go:build oled_featherwing

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/adafruit4650"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/notoemoji"
	"tinygo.org/x/tinyfont/notosans"
	_ "tinygo.org/x/tinyfont/notosans"
	"tinygo.org/x/tinyfont/proggy"
)

const (
	timeLayout = "03:04 PM MST"
	dateLayout = "Mon Jan _2, 2006"

	displayTimeout = 30 * time.Second // 1 * time.Minute
)

var (
	white    = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	display  = adafruit4650.New(machine.I2C0)
	onButton = machine.D5

	lastTime, lastDate string

	displayOn bool
	lastOn    time.Time
	lastInd   int16
)

func initDisplay() error {
	onButton.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	if err := display.Configure(); err != nil {
		return err
	}
	if err := display.ClearDisplay(); err != nil {
		return err
	}
	lastOn = time.Now()
	return nil
}

func displayTask() {
	if !onButton.Get() {
		lastOn = time.Now()
	}
}

func setDisplay(on bool) {
	if on {
		lastOn = time.Now()
	} else {
		lastOn = time.Now().AddDate(0, 0, -1)
	}
}

func showTime(state *DisplayState, force bool) error {

	timeout := time.Since(lastOn) > displayTimeout

	if timeout && displayOn {
		display.ClearDisplay()
		displayOn = false
	}

	tstr := state.ts.Format(timeLayout)
	dstr := state.ts.Format(dateLayout)
	secs := state.ts.Second() % 30
	indicatorH := int16(30-secs)*2 + 4

	if tstr != lastTime || dstr != lastDate || indicatorH != lastInd || force {
		if !timeout {
			display.ClearBuffer()

			// frame out the user interface
			const (
				rectX = 14
				rectY = 0
				rectH = 32
				rectW = 128
			)
			tinydraw.Rectangle(&display, rectX, rectY, rectW-rectX, rectH, white)
			tinydraw.Rectangle(&display, rectX, rectY+(rectH), rectW-rectX, rectH, white)

			// time and date
			const timeY = 0 //31
			displayRightJustified(&notosans.Notosans12pt, 124, timeY+15, tstr)
			displayRightJustified(&proggy.TinySZ8pt7b, 124, timeY+26, dstr)

			// TOTP details
			const (
				totpY = 31
				totpX = 124
			)
			account := totpKeys[0].Name
			digits := "000000"
			displayRightJustified(&notoemoji.NotoEmojiRegular12pt, totpX, totpY+16, digits)
			displayRightJustified(&proggy.TinySZ8pt7b, totpX, totpY+28, account)

			// time left indicator
			tinydraw.FilledRectangle(&display, 0, (rectH*2)-indicatorH, rectX, indicatorH, white)

			// display the buffer
			display.Display()
			displayOn = true

		}
		lastTime = tstr
		lastDate = dstr
		lastInd = indicatorH
	}
	return nil
}

func displayRightJustified(font tinyfont.Fonter, right int16, y int16, str string) {
	_, x := tinyfont.LineWidth(font, str)
	tinyfont.WriteLine(&display, font, right-int16(x), y, str, white)
}
