//go:build rtc_experimental

package main

import "time"

type RTC struct {
	dev interface {
		init() error
		read() (time.Time, error)
	}
}
