package multihost

import (
	"github.com/bgould/keyboard-firmware/keyboard"
)

const debug = false

type Host struct {
	delegates []keyboard.Host
}

func New(delegates ...keyboard.Host) *Host {
	return &Host{delegates}
}

func (host *Host) Send(rpt keyboard.Report) {
	for _, delegate := range host.delegates {
		delegate.Send(rpt)
	}
}

func (host *Host) LEDs() uint8 {
	return 0
}
