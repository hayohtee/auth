package main

import (
	"flag"
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
	
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "The port to start the server on")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("DATABASE_DSN"), "The database DSN")
	
}
