CREATE TABLE IF NOT EXISTS user_credentials (
    id bigserial PRIMARY KEY,
    user_id bigserial REFERENCES users (id) ON DELETE CASCADE,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL
);
