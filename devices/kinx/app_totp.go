package main

import (
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/console"
	totp "github.com/bgould/tinytotp"
)

var (
	// defaultTotpAccount string
	// defaultTotpKey     SecureString

	defaultTotpAccount string       // = "keyboard tester"
	defaultTotpKey     SecureString // = "KWNKKXLJYPHGSFCB"
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

func totptask() {
	// TOTP-related functionality
	if totpKeys[0].Name != "" && totpKeys[0].Key != "" {
		ds.totpCounter = uint64(totp.TimeBasedCounter(time.Now(), totp.DefaultOpts.Period))
		if ds.totpCounter != lastTotp {
			// TODO: un-hardcode index
			ds.totpAccount = totpKeys[0].Name
			numbers, err := totp.GenerateCode(string(totpKeys[0].Key), time.Now())
			if err != nil {
				cli.WriteString("warning: error updating TOTP - " + err.Error())
				numbers = "000000"
			}
			ds.totpNumbers = numbers
			lastTotp = ds.totpCounter
		}
	}
	if err := showTime(ds, false); err != nil {
		cli.WriteString("warning: error updating display - " + err.Error())
	}
	displayTask()
}
