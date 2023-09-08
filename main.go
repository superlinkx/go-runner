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

package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/superlinkx/go-runner/command"
	"github.com/superlinkx/go-runner/environment"
)

func main() {
	if err := setLogger(); err != nil {
		log.Fatalf("Failed to set log level: %s", err)
	} else if cmd, err := command.ParseCLI(); err != nil {
		slog.Error("Failed to parse CLI", "error", err)
		os.Exit(1)
	} else if err := cmd.Run(); err != nil {
		slog.Error("Error running command", "command", cmd.Name(), "error", err)
		os.Exit(1)
	} else {
		slog.Info("Command ran successfully", "command", cmd.Name())
	}
}

func setLogger() error {
	var logLevel slog.Level

	envLvl := os.Getenv(environment.LogLevel.Env())
	if envLvl == "" {
		logLevel = slog.LevelInfo
	} else {
		switch envLvl {
		case "debug":
			logLevel = slog.LevelDebug
		case "info":
			logLevel = slog.LevelInfo
		case "warn":
			logLevel = slog.LevelWarn
		case "error":
			logLevel = slog.LevelError
		default:
			return fmt.Errorf("invalid log level: %s", envLvl)
		}
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))
	return nil
}
