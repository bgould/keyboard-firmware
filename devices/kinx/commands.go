package main

import (
	"strconv"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/console"
)

func initCommands() console.Commands {
	return console.Commands{
		"rate": console.CommandHandlerFunc(rate),
		"time": console.Commands{
			"set": console.CommandHandlerFunc(timeset),
			"get": console.CommandHandlerFunc(timeget),
		},
		"totp": totpCommands,
	}
}

func rate(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("scan rate: "))
	cmd.Stdout.Write([]byte(strconv.Itoa(ds.scanRate)))
	cmd.Stdout.Write([]byte("\n"))
	return 0
}

func timeset(cmd console.CommandInfo) int {
	// fmt.Fprintf(cmd.Stdout, "timeset called: %v\n", cmd)
	if len(cmd.Argv) != 1 {
		cmd.Stdout.Write([]byte("Usage: " + cmd.Cmd + " <unix time in seconds>\n"))
		return -2
	}
	unixTime, err := strconv.ParseInt(cmd.Argv[0], 10, 64)
	if err != nil {
		cmd.Stdout.Write([]byte("Could not parse unix timestamp: "))
		cmd.Stdout.Write([]byte(cmd.Argv[0]))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	}
	t := time.Unix(int64(unixTime), 0)
	cmd.Stdout.Write([]byte("Setting unix timestamp: "))
	cmd.Stdout.Write([]byte(strconv.FormatInt(unixTime, 10)))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte(t.Format(time.RFC3339)))
	cmd.Stdout.Write([]byte("\n"))
	if err := board.RTCSet(t); err != nil {
		cmd.Stdout.Write([]byte("error setting unix time: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 2
	}
	return 0
}

func timeget(cmd console.CommandInfo) int {
	// fmt.Fprintf(cmd.Stdout, "timeset called: %v\n", cmd)
	// var ok bool
	now, err := board.RTCTime() //readTime()
	ok := err == nil
	cmd.Stdout.Write([]byte("RTC configured: "))
	cmd.Stdout.Write([]byte(strconv.FormatBool(ok)))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Unix timestamp: "))
	cmd.Stdout.Write([]byte(strconv.FormatInt(now.UTC().Unix(), 10)))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Local datetime: "))
	cmd.Stdout.Write([]byte(now.Local().Format(time.RFC3339)))
	cmd.Stdout.Write([]byte("\n"))
	return 0
}
