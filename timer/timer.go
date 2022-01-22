package timer

import "time"

func Wait(duration time.Duration) {
	//start := time.Now().UnixNano()
	//for stop := start + int64(duration); time.Now().UnixNano() < stop; {
	//}
	New(duration).WaitUntilExpired()
}

type Timer struct {
	start    int64
	interval int64
}

func New(interval time.Duration) Timer {
	return Timer{
		start:    time.Now().UnixNano(),
		interval: int64(interval),
	}
}

func (t Timer) Expired() bool {
	return time.Now().UnixNano() > (t.start + t.interval)
}

func (t Timer) WaitUntilExpired() {
	for !t.Expired() {
	}
}
