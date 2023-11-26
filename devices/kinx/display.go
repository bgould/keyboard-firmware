package main

import "time"

var ds DisplayState

type DisplayState struct {
	ts   time.Time
	tsOk bool

	scanRate int
}
