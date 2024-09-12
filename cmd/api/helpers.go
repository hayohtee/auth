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

func generateJWT(userID int64, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     "http://localhost:4000",
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
		"user_id": userID,
		"role":    role,
	})

	return token.SignedString([]byte(JWTSecret))
}

func validateJWT(token string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
