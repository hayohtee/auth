package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	// ErrDuplicateEmail is a custom error that is returned when there
	// is a duplicate email in the database.
	ErrDuplicateEmail = errors.New("duplicate email")

	// ErrRecordNotFound is a custom error that is returned when looking
	// for a specific record that is not in the database.
	ErrRecordNotFound = errors.New("record not found")
)

// UserModel is a type which wraps around a sql.DB connection pool
// and provide methods for creating and managing users to and from
// the database.
type UserModel struct {
	DB *sql.DB
}

// Insert a new user record to the database.
func (u UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users(name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	args := []any{user.Name, user.Email, user.Password.hash}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

// GetByEmail retrieve the User details from the database based
// on the user's email address.
func (u UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT id, created_at, name, email, password_hash
		FROM users
		WHERE email = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
