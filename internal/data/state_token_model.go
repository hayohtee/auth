package data

import (
	"context"
	"database/sql"
	"time"
)

type StateTokenModel struct {
	db *sql.DB
}

func NewStateTokenModel(db *sql.DB) *StateTokenModel {
	return &StateTokenModel{
		db: db,
	}
}

func (s StateTokenModel) New(ttl time.Duration) (StateToken, error) {
	token, err := generateStateToken(ttl)
	if err != nil {
		return StateToken{}, err
	}
	err = s.Insert(token)
	return token, err
}

func (s StateTokenModel) Insert(token StateToken) error {
	query := `
		INSERT INTO state_tokens(hash, expiry)
		VALUES($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, token.Hash, token.Expiry)
	return err
}
