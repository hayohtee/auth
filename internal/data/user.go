package data

type User struct {
	ID            int64    `json:"id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Password      password `json:"-"`
	EmailVerified bool     `json:"email_verified"`
	AvatarURL     string   `json:"avatar_url,omitempty"`
}
