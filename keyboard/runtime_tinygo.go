//go:build tinygo

package keyboard

import (
	"runtime"
	"time"
)

func adjustTimeOffset(t time.Time) bool {
	runtime.AdjustTimeOffset(-1 * int64(time.Since(t)))
	return true
}
