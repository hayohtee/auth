package main

import "net/http"

func (app *application) signUpWithGoogle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "", http.StatusTemporaryRedirect)
}

func (app *application) signUpWithGoogleCallback(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "", http.StatusTemporaryRedirect)
}	