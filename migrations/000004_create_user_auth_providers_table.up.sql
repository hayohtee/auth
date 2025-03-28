CREATE TABLE IF NOT EXISTS user_auth_providers(
    id bigserial PRIMARY KEY,
    user_id bigint REFERENCES users(id) ON DELETE CASCADE,
    email citext NOT NULL,
    provider TEXT NOT NULL,
    provider_id TEXT UNIQUE NOT NULL,
    UNIQUE(user_id, provider_id)
);