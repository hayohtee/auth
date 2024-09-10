CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT                        NOT NULL,
    email         citext UNIQUE               NOT NULL,
    password_hash bytea                       NOT NULL,
    created_at    TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);