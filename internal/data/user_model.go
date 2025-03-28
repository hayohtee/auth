package data

import (
	"context"
	"database/sql"
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
		)
		INSERT INTO user_credentials(user_id, email, password_hash)
		SELECT n.id, $2, $3 FROM n new_user
		RETURNING n.id, email, n.created_at
		`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{user.User.Name, user.Credential.Email, user.Credential.Password.Hash}
	err := u.db.QueryRowContext(ctx, query, args...).Scan(&user.User.ID, &user.User.Email, &user.User.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "user_credentials_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
