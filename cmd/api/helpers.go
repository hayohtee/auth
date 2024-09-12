package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	JWTSecret = os.Getenv("JWT_SECRET")
)

type claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func generateJWT(userID int64, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	payload := &claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "http::localhost:4000",
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(JWTSecret))
}

func validateJWT(token string) (*claims, error) {
	payload := claims{}

	t, err := jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	return &payload, nil
}
