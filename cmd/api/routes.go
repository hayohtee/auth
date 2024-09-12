package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Endpoints goes here
	mux.HandleFunc("POST /users", app.createUserHandler)
	mux.HandleFunc("POST /users/login", app.loginUserHandler)

	mux.Handle("GET /admin", app.authorize(RoleAdmin, app.welcomeAdminHandler))
	mux.Handle("GET /customer", app.authorize(RoleCustomer, app.welcomeCustomerHandler))

	return mux
}
