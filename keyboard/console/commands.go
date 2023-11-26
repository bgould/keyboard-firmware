package console

import (
	"io"
	"strings"
)

const (
	CommandUnknown = -1
)

type CommandInfo struct {
	Cmd    string
	Argv   []string
	Stdout io.Writer
}

type CommandHandler interface {
	HandleCommand(cmd CommandInfo) (exitCode int)
}

type CommandHandlerFunc func(cmd CommandInfo) (exitCode int)

func (fn CommandHandlerFunc) HandleCommand(cmd CommandInfo) (exitCode int) {
	return fn(cmd)
}

type Commands map[string]CommandHandler

func (cmds Commands) HandleCommand(info CommandInfo) (exitCode int) {
	if len(info.Argv) < 1 {
		info.Stdout.Write([]byte("    available commands:\n"))
		for cmd := range cmds {
			info.Stdout.Write([]byte("      "))
			info.Stdout.Write([]byte(cmd))
			info.Stdout.Write([]byte("\n"))
		}
		return 0
	}
	sub := info.Argv[0]
	// fmt.Fprintf(info.Stdout, "handling command: %v (%s) with map %v\n", info, sub, cmds)
	cmd, ok := cmds[sub]
	if !ok {
		return CommandUnknown
	} else {
		// fmt.Fprintf(info.Stdout, "found subcommand: %v\n", cmd)
	}
	info.Cmd = strings.TrimSpace(info.Cmd + " " + sub)
	info.Argv = info.Argv[1:]
	// fmt.Fprintf(info.Stdout, "forwarding command: %v\n", info)
	return cmd.HandleCommand(info)
}

var _ CommandHandler = (Commands)(nil)
