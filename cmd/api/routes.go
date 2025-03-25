package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /v1/auth/signup/oauth/google", app.signUpWithGoogle)
	mux.HandleFunc("GET /v1/auth/signup/oauth/google/callback", app.signUpWithGoogleCallback)
	
	return mux
}