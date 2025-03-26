CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    email citext NOT NULL,
    password bytea,
    email_verified bool NOT NULL,
    avatar_url text,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
);