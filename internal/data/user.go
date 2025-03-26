package data

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      password  `json:"-"`
	EmailVerified bool      `json:"email_verified"`
	AvatarURL     string    `json:"avatar_url,omitempty"`
	CreatedAt     time.Time `json:"-"`
}

type password struct {
	PlainText string
	Hash      []byte
}

func (p *password) Set(plainTextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 12)
	if err != nil {
		return err
	}
	p.PlainText = plainTextPassword
	p.Hash = hash
	return nil
}

func (p *password) Matches(plainTextPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plainTextPassword)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
