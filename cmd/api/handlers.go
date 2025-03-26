package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (app *application) signUpWithGoogle(w http.ResponseWriter, r *http.Request) {
	token, err := app.stateTokenModel.New(2 * time.Hour)
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
