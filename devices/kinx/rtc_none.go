//go:build false && !rtc.pcf8523

package main

var (
	rtcInit bool
)

func initRTC() bool {
	cli.WriteString("RTC not configured")
	return true
}

// func timeTask() {
// }

// func readTime() (time.Time, bool) {
// 	return time.Now(), rtcInit
// }

// func setUnixTime(t time.Time) error {
// 	return nil
// }

// func rtcSync() {}
