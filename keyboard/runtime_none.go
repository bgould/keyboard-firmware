//go:build !tinygo

package keyboard

import (
	"time"
)

func adjustTimeOffset(t time.Time) bool {
	return false
}
