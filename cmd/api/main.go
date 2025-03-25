package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
