package main

import (
	"time"

	"github.com/bgould/keyboard-firmware/devices/kinx/totp"
	"github.com/bgould/keyboard-firmware/keyboard/console"
)

var (
	// defaultTotpAccount string
	// defaultTotpKey     SecureString

	defaultTotpAccount string       = "keyboard tester"
	defaultTotpKey     SecureString = "KWNKKXLJYPHGSFCB"
)

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
		{Name: defaultTotpAccount, Key: defaultTotpKey},
	}
	totpCommands = console.Commands{
		"get": console.CommandHandlerFunc(totpget),
	}
)

func totpget(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("totp attempt"))
	cmd.Stdout.Write([]byte(time.Now().Local().String()))
	if code, err := totp.GenerateCode(string(totpKeys[0].Key), time.Now()); err != nil {
		cmd.Stdout.Write([]byte("error generating TOTP: " + err.Error()))
		return 1
	} else {
		cmd.Stdout.Write([]byte("generated code: " + code + "\n"))
		return 0
	}
}
