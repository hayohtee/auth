package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Endpoints goes here

	return mux
}
