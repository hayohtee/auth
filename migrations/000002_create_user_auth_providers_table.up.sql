CREATE TABLE IF NOT EXISTS user_auth_providers(
    id bigserial PRIMARY KEY,
    user_id bigint REFERENCES users(id) ON DELETE CASCADE,
    provider text NOT NULL CHECK(provider IN ('google', 'apple', 'github', 'facebook')),
    provider_id text UNIQUE NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    UNIQUE(user_id, provider)
);