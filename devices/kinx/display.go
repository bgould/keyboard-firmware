package main

import "time"

var (
	ds, lastState DisplayState
)

type DisplayState struct {
	ts   time.Time
	tsOk bool

	scanRate int

	totpCounter uint64
	totpAccount string
	totpNumbers string
}
