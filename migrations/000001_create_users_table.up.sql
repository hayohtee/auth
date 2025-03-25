CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    name TEXT NOT NULL,
    email citext UNIQUE NOT NULL,
    password bytea,
    activated bool NOT NULL
);