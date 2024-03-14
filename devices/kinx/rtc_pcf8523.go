//go:build rtc.pcf8523

package main

import (
	"tinygo.org/x/drivers/pcf8523"
)

var (
	rtc = pcf8523.New(i2c)

	// rtcInit bool
	// rtcErr  error

	// rtcLast time.Time
)

func initRTC() bool {
	board.CLI().WriteString("initializing pcf8523")
	// make sure the battery takes over if power is lost
	err := rtc.SetPowerManagement(pcf8523.PowerManagement_SwitchOver_ModeStandard)
	if err != nil {
		board.CLI().WriteString("rtc error: " + err.Error())
		return false
	}
	board.SetRTC(&rtc)
	return true
}

// var (
// 	ErrNotInitialized = errors.New("not initialized")
// )

// type rtcState int

// const (
// 	rtcStateIdle rtcState = iota
// 	rtcStateConn
// 	rtcStateErr
// )

// func init() {
// 	time.Local = time.FixedZone("EST", -5*3600)
// }

// func timeTask() {
// 	if !rtcInit {
// 		return
// 	}
// 	if time.Since(rtcLast) > time.Hour {
// 		t, ok := readTime()
// 		if !ok {
// 			println("could not read current time from RTC")
// 		}
// 		runtime.AdjustTimeOffset(-1 * int64(time.Since(t)))
// 		rtcLast = t
// 	}
// }

// var newRTCTime time.Time

// func setUnixTime(t time.Time) error {
// 	if !rtcInit {
// 		return ErrNotInitialized
// 	}
// 	// if err := rtc.SetTime(t.UTC()); err != nil {
// 	// 	return err
// 	// }
// 	adjustTimeOffset(t)
// 	rtcUpdate <- struct{}{}
// 	rtcLast = t
// 	return nil
// }

// func readTime() (time.Time, bool) {
// 	if !rtcInit {
// 		return time.Now(), false
// 	}
// 	ts, err := rtc.ReadTime()
// 	if err != nil {
// 		return time.Now(), false
// 	}
// 	return ts, true
// }

// func rtcSync() {
// 	if rtcInit {
// 		println("sync time")
// 		rtc.SetTime(time.Now().UTC())
// 	}
// }
