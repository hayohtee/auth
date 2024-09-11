package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Endpoints goes here
	mux.HandleFunc("POST /users", app.createUserHandler)
	mux.HandleFunc("POST /users/login", app.loginUserHandler)

	return mux
}
