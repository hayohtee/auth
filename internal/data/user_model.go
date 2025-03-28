package data

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"
)

type UserModel struct {
	db *sql.DB
}

func (u UserModel) Insert(user *UserWithCredential) error {
	query := `
		WITH new_user AS (
			INSERT INTO users(name)
			VALUES($1)
			RETURNING id, created_at
		), inserted_credentials AS (
			INSERT INTO user_credentials(user_id, email, password_hash)
			SELECT id, $2, $3 FROM new_user
			RETURNING email
		)
		SELECT new_user.id, inserted_credentials.email, new_user.created_at
	 	FROM new_user, inserted_credentials`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{user.User.Name, user.Credential.Email, user.Credential.Password.Hash}
	err := u.db.QueryRowContext(ctx, query, args...).Scan(&user.User.ID, &user.User.Email, &user.User.CreatedAt)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), `duplicate key value violates unique constraint "user_credentials_email_key"`):
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (u UserModel) GetByEmail(email string) (UserWithCredential, error) {
	query := `
		SELECT u.id, u.name, uwc.email, uwc.password_hash, u.avatar_url, u.created_at
		FROM users u
		JOIN user_credentials uwc ON u.id = uwc.user_id
		WHERE uwc.email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user UserWithCredential
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&user.User.ID,
		&user.User.Name,
		&user.User.Email,
		&user.Credential.Password.Hash,
		&user.User.AvatarURL,
		&user.User.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return UserWithCredential{}, ErrRecordNotFound
		default:
			return UserWithCredential{}, err
		}
	}

	return user, nil
}
