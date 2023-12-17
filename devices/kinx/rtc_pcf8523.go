//go:build rtc.pcf8523

package main

import (
	"errors"
	"runtime"
	"time"

	"tinygo.org/x/drivers/pcf8523"
)

var (
	rtc = pcf8523.New(i2c)

	rtcInit bool
	rtcErr  error

	rtcLast time.Time
)

var (
	ErrNotInitialized = errors.New("not initialized")
)

type rtcState int

const (
	rtcStateIdle rtcState = iota
	rtcStateConn
	rtcStateErr
)

func init() {
	time.Local = time.FixedZone("EST", -5*3600)
}

func initTime() bool {
	cli.WriteString("initializing pcf8523")
	// make sure the battery takes over if power is lost
	rtcErr = rtc.SetPowerManagement(pcf8523.PowerManagement_SwitchOver_ModeStandard)
	rtcInit = rtcErr == nil
	return rtcInit
}

func timeTask() {
	if !rtcInit {
		return
	}
	if time.Since(rtcLast) > time.Hour {
		t, ok := readTime()
		if !ok {
			println("could not read current time from RTC")
		}
		runtime.AdjustTimeOffset(-1 * int64(time.Since(t)))
		rtcLast = t
	}
}

func setUnixTime(t time.Time) error {
	if !rtcInit {
		return ErrNotInitialized
	}
	if err := rtc.SetTime(t.UTC()); err != nil {
		return err
	}
	adjustTimeOffset(t)
	rtcLast = t
	return nil
}

func readTime() (time.Time, bool) {
	if !rtcInit {
		return time.Now(), false
	}
	ts, err := rtc.ReadTime()
	if err != nil {
		return time.Now(), false
	}
	return ts, true
}
