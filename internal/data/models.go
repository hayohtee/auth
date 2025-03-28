package data

import "database/sql"

type Models struct {
	Users         UserModel
	AuthProviders UserAuthProviderModel
	Tokens        StateTokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:         UserModel{db: db},
		AuthProviders: UserAuthProviderModel{db: db},
		Tokens:        StateTokenModel{db: db},
	}
}
