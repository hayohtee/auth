package main

import "github.com/hayohtee/auth/internal/data"

type application struct {
	config config
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
