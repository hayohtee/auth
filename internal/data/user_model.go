package data

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{
		db: db,
	}
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

func (u UserModel) InsertForAuthProvider(user *UserWithAuthProvider) error {
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
