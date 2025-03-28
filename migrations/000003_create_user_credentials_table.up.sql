CREATE TABLE IF NOT EXISTS user_credentials (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL
);
