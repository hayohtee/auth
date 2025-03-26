package main

import "github.com/hayohtee/auth/internal/data"

type application struct {
	config          config
	stateTokenModel *data.StateTokenModel
}

type config struct {
	port  int
	dsn   string
	oauth struct {
		googleClientID     string
		googleClientSecret string
	}
}
