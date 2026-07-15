package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))
	args := os.Args[1:]

	if len(args) < 1 {
		slog.Error("Not enough arguments")
	}

	fileToLookAt := args[0]
	slog.Debug("Looking at file", "file", fileToLookAt)
}
