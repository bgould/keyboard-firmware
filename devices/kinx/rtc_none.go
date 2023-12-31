//go:build !rtc.pcf8523

package main

import (
	"time"
)

var (
	rtcInit bool
)

func init() {
	time.Local = time.FixedZone("EST", -5*3600)
}

func initTime() bool {
	cli.WriteString("RTC not configured")
	return true
}

func timeTask() {
}

func readTime() (time.Time, bool) {
	return time.Now(), rtcInit
}

func setUnixTime(t time.Time) error {
	return nil
}

func rtcSync() {}
