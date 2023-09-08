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
		logLevel = slog.LevelError
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
