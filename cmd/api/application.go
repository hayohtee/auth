package main

type application struct {
	config config
}

type config struct {
	port int
	dsn  string
}
