package main

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	JWTSecret = os.Getenv("JWT_SECRET")
)

func generateJWT(userID int64, name, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     "http://localhost:4000",
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
		"user_id": userID,
		"name":    name,
		"email":   email,
	})

	return token.SignedString([]byte(JWTSecret))
}
