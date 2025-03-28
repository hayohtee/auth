package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/auth/signup/oauth/google", app.signUpWithGoogle)
	mux.HandleFunc("GET /v1/auth/signup/oauth/google/callback", app.signUpWithGoogleCallback)

	mux.HandleFunc("POST /v1/auth/signup", app.signUpUserHandler)
	mux.HandleFunc("POST /v1/auth/login", app.loginUserHandler)

	return mux
}
