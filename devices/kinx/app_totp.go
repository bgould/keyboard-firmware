package main

import "github.com/bgould/keyboard-firmware/keyboard/console"

type TOTPKey struct {
	Name string
	Key  SecureString
}

type SecureString string

func (s SecureString) String() string {
	return "<redacted>"
}

var (
	totpKeys = []TOTPKey{
		{Name: "GitHub", Key: "KWNKKXLJYPHGSFCB"},
	}
	totpCommands = console.Commands{
		"get": console.CommandHandlerFunc(totpget),
	}
)
