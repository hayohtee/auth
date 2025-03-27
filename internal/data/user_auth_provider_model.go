package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserAuthProviderModel struct {
	db *sql.DB
}

func (u UserAuthProviderModel) Insert(user *UserWithAuthProvider) error {
	query := `
		INSERT INTO users(name, email, email_verified, avatar_url)
		VALUES($1, $2, $3, $4)
		RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{user.User.Name, user.User.Email, user.User.EmailVerified, user.User.AvatarURL}
	if err := u.db.QueryRowContext(ctx, query, args...).Scan(&user.User.ID, &user.User.CreatedAt); err != nil {
		return err
	}

	query = `
		INSERT INTO user_auth_providers(user_id, provider, provider_id)
		VALUES($1, $2, $3)`

	args = []any{user.User.ID, user.AuthProvider.Provider, user.AuthProvider.ProviderID}
	_, err := u.db.ExecContext(ctx, query, args...)
	return err
}

func (u UserAuthProviderModel) Get(provider, providerID string) (User, error) {
	query := `
		SELECT u.id, u.name, u.email, u.email_verified, u.avatar_url, u.created_at
		FROM users u
		JOIN user_auth_providers ua ON u.id = ua.user_id
		WHERE ua.provider = $1 AND ua.provider_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := u.db.QueryRowContext(ctx, query, provider, providerID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.EmailVerified,
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
