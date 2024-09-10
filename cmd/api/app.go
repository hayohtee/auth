package main

import (
	"github.com/hayohtee/auth/internal/data"
	"log/slog"
)

// application holds the dependencies for the handlers, middlewares
// and helpers.
type application struct {
	logger *slog.Logger
	users  data.UserModel
}
