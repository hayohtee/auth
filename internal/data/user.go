package data

import (
	"errors"
	"time"

	"github.com/hayohtee/auth/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type UserWithCredential struct {
	User       User
	Credential UserCredential
}

type UserCredential struct {
	UserID   int64
	Email    string
	Password password
}

type UserAuthProvider struct {
	UserID     int64
	Email      string
	Provider   string
	ProviderID string
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

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePlainPassword(v *validator.Validator, plainPassword string) {
	v.Check(plainPassword != "", "password", "must be provided")
	v.Check(len(plainPassword) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(plainPassword) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateName(v *validator.Validator, name string) {
	v.Check(name != "", "name", "must be provided")
	v.Check(len(name) <= 500, "name", "must not be more than 500 bytes long")
}
