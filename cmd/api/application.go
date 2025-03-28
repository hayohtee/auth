package main

import (
	"log/slog"

	"github.com/hayohtee/auth/internal/data"
)

type application struct {
	config config
	logger *slog.Logger
	models data.Models
}

type config struct {
	port  int
	dsn   string
	oauth struct {
		googleClientID     string
		googleClientSecret string
	}
}
