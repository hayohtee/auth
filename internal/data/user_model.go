package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserModel struct {
	db *sql.DB
}

func (u UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users(name, email, email_verified, password_hash)
		VALUES($1, $2, $3, $4)
		RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{user.Name, user.Email, user.EmailVerified, user.Password.Hash}
	if err := u.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (u UserModel) GetByEmail(email string) (User, error) {
	query := `
		SELECT id, name, email, email_verified, password_hash, avatar_url, created_at
		FROM users
		WHERE id = $1 AND password_hash IS NOT NULL`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.EmailVerified,
		&user.Password.Hash,
		&user.AvatarURL,
		&user.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return User{}, ErrRecordNotFound
		default:
			return User{}, err
		}
	}
	return user, nil
}
