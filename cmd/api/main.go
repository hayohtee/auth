package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Load .env file
	if err := godotenv.Load(); err != nil {
		logger.Error("failed to load .env file", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
