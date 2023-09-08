package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/superlinkx/go-runner/command/license"
)

type Commander interface {
	Name() string
	Usage() string
	Run() error
}

var (
	ErrNoCmd           = errors.New("no command specified")
	ErrInvalidCmd      = errors.New("invalid command specified")
	ErrFailedCreateCmd = errors.New("failed to create command")
)

func ParseCLI() (Commander, error) {
	// Generate a nice usage message
	flag.Usage = usage

	// Default usage if no arguments provided
	if len(os.Args) < 2 {
		flag.Usage()
		return nil, ErrNoCmd
	}

	switch os.Args[1] {
	case license.Name:
		return license.Create(), nil
	default:
		return nil, ErrInvalidCmd
	}
}

// usage creates a pretty usage message for our main command
func usage() {
	var longestCmdLen int

	w := flag.CommandLine.Output()
	fmt.Fprint(w, "A go commandlet runner\n\nUsage: go-runner COMMAND\n\nCommands:\n")

	for _, cmd := range Commands() {
		if len(cmd.String()) > longestCmdLen {
			longestCmdLen = len(cmd.String())
		}
	}

	for cmd, usage := range CommandsUsage() {
		cmdStr := subCmd(cmd).String()
		padding := strings.Repeat(" ", longestCmdLen-len(cmdStr))
		fmt.Fprintf(w, "  %s%s    %s\n", cmdStr, padding, usage)
	}
}
