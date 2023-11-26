//go:build !rtc.pcf8523

package main

import (
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

func init() {
	time.Local = time.FixedZone("EST", -5*3600)
}

func initTime() bool {
	return true
}

func timeTask() {
}

func readTime() (time.Time, bool) {
	return time.Now(), rtcInit
}

func setUnixTime(t time.Time) error {
	runtime.AdjustTimeOffset(-1 * int64(time.Since(t)))
	return nil
}
