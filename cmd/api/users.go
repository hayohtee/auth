package main

import (
	"encoding/json"
	"errors"
	"github.com/hayohtee/auth/internal/data"
	"github.com/hayohtee/auth/internal/validator"
	"net/http"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user := data.User{
		Name:  input.Name,
		Email: input.Email,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	v := validator.New()
	if data.ValidateUser(v, &user); !v.Valid() {
		response := map[string]any{
			"error": v.Errors,
		}

		js, err := json.Marshal(&response)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(js)
		return
	}

	err = app.users.Insert(&user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email already exists")
			response := map[string]any{
				"error": v.Errors,
			}

			js, err := json.Marshal(&response)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write(js)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		return
	}

	js, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}
