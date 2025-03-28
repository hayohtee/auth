CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    name TEXT NOT NULL,
    avatar_url TEXT,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now()
);