package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hayohtee/auth/internal/data"
	"github.com/hayohtee/auth/internal/validator"
)

func (app *application) signUpWithGoogle(w http.ResponseWriter, r *http.Request) {
	token, err := app.models.Tokens.New(2 * time.Hour)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	authEndpoint := "https://accounts.google.com/o/oauth2/v2/auth"
	url := fmt.Sprintf("%s?client_id=%s&response_type=code&scope=openid profile&state=%s&redirect_uri=http://localhost:4000/v1/auth/signup/oauth/google/callback", authEndpoint, app.config.oauth.googleClientID, token.PlainText)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (app *application) signUpWithGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	fmt.Println(code)

	endPoint := "https://oauth2.googleapis.com/token"

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	body := url.Values{}
	body.Add("code", code)
	body.Add("client_id", app.config.oauth.googleClientID)
	body.Add("client_secret", app.config.oauth.googleClientSecret)
	body.Add("redirect_uri", "http://localhost:4000/v1/auth/signup/oauth/google/callback")
	body.Add("grant_type", "authorization_code")

	bodyReader := strings.NewReader(body.Encode())

	req, err := http.NewRequest(http.MethodPost, endPoint, bodyReader)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// var data struct {
	// 	AccessToken string `json:"access_token"`
	// 	ExpiresIn   int64  `json:"expires_in"`
	// 	IDToken     string `json:"id_token"`
	// 	Scope       string `json:"scope"`
	// }

	// if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }
	//

	js, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func (app *application) signUpUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(input.Email != "", "email", "must be provided")
	v.Check(validator.Matches(input.Email, validator.EmailRX), "email", "must be a valid email address")
	v.Check(input.Password != "", "password", "must be provided")
	v.Check(len(input.Password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(input.Password) <= 72, "password", "must not be more than 72 bytes long")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user := data.UserWithCredential{
		User: data.User{
			Name: input.Name,
		},
		Credential: data.UserCredential{
			Email: input.Email,
		},
	}

	if err := user.Credential.Password.Set(input.Password); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.models.Users.Insert(&user); err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email already exist")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	data := envelope{"user": user.User}
	if err := app.writeJSON(w, http.StatusCreated, data, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {

}
