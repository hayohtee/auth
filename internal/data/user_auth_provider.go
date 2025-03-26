package data

import "time"

type UserAuthProvider struct {
	ID         int64
	UserID     int64
	Provider   string
	ProviderID string
	CreatedAt  time.Time
}

type UserWithAuthProvider struct {
	User         User
	AuthProvider UserAuthProvider
}
