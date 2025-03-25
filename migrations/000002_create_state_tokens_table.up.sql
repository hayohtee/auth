CREATE TABLE IF NOT EXISTS state_tokens(
    hash bytea PRIMARY KEY,
    expiry timestamp(0) with time zone NOT NULL
);