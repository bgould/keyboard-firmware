package keyboard

import "time"

type RTC interface {
	// Init() error
	ReadTime() (t time.Time, err error)
	SetTime(t time.Time) (err error)
}

func (kbd *Keyboard) SetRTC(dev RTC) {
	if dev == nil {
		kbd.rtc = nil
		return
	}
	// TODO: "uninitialize" when replacing?
	kbd.rtc = &rtc{dev: dev, last: time.Time{}, init: true, update: make(chan time.Time, 1)}
}

func (kbd *Keyboard) RTCInitialized() bool {
	return kbd.rtc.initialized()
}

func (kbd *Keyboard) RTCTime() (time.Time, error) {
	if !kbd.rtc.initialized() {
		return time.Time{}, ErrRTCNotSet
	}
	return kbd.rtc.dev.ReadTime()
}

func (kbd *Keyboard) RTCSet(t time.Time) (err error) {
	if !kbd.rtc.initialized() {
		return ErrRTCNotSet
	}
	select {
	case kbd.rtc.update <- t:
		return nil
	default:
		return ErrRTCUpdating
	}
}

type RTCError int

const (
	ErrRTCNotSet RTCError = iota
	ErrRTCNoInit
	ErrRTCUpdating
)

func (err RTCError) Error() string {
	switch err {
	case ErrRTCNotSet:
		return "RTC: not set"
	case ErrRTCNoInit:
		return "RTC: not initalized"
	case ErrRTCUpdating:
		return "RTC: pending update"
	default:
		return "RTC: unknown"
	}
}

type rtc struct {
	dev RTC

	init bool
	last time.Time

	update chan time.Time
}

func (rtc *rtc) initialized() bool {
	return rtc != nil && rtc.dev != nil && rtc.init
}

func (rtc *rtc) task() {
	if !rtc.initialized() {
		return
	}
	select {
	case t := <-rtc.update:
		println("updating time", t.String())
		rtc.last = time.Time{}
		rtc.dev.SetTime(t)
	default:
		// fallthrough, no pending update
	}
	// TODO: maybe should track time skew to notify/warn if clock might need update
	if time.Since(rtc.last) > time.Hour {
		t, err := rtc.dev.ReadTime() //readTime()
		if err != nil {
			println("could not read current time from RTC", err.Error())
		}
		adjustTimeOffset(t)
		rtc.last = t
	}
}

// func (kbd *Keyboard) RTCInit() error {
// 	if kbd.rtc.initialized() {
// 		return nil
// 	}
// 	if kbd.rtc == nil || kbd.rtc.dev == nil {
// 		return ErrRTCNotSet
// 	}
// 	if err := kbd.rtc.dev.Init(); err != nil {
// 		return err
// 	}
// 	kbd.rtc.init = true
// 	return nil
// }
