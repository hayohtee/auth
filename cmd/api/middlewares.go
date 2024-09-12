package main

import (
	"net/http"
	"strings"
)

func (app *application) authorize(role string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		data := strings.Split(authorizationHeader, " ")
		if len(data) != 2 {
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		payload, err := validateJWT(data[1])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !strings.EqualFold(payload.Role, role) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
