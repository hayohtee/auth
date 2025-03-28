package data

import "database/sql"

type Models struct {
	Users  UserModel
	Tokens StateTokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:  UserModel{db: db},
		Tokens: StateTokenModel{db: db},
	}
}
