// MIT License
// 
// Copyright (c) 2023 Alyx Holms
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
