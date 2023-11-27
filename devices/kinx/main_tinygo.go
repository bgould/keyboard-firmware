//go:build tinygo

package main

import (
	"runtime"
	"time"
)

func adjustTimeOffset(t time.Time) {
	runtime.AdjustTimeOffset(-1 * int64(time.Since(t)))
}
