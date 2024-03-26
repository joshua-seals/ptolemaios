-- Version: 1.1
-- Description: Create Sessions table with index for expiry.
-- sessions expiry is managed by alexedwards/scs automatically.
CREATE TABLE IF NOT EXISTS sessions (
	token TEXT PRIMARY KEY,
	data BYTEA NOT NULL,
	expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);

-- -- Version: 1.2
-- -- Description: ptolemaios products
-- CREATE TABLE IF NOT EXISTS ptolemaios (
-- 	ptolemaios_id INT GENERATED ALWAYS AS IDENTITY,
-- 	ptolemaios_name TEXT,
-- 	ptolemaios_description TEXT,
-- 	ptolemaios_price NUMERIC(5,2),
-- 	ptolemaios_caffeine TEXT,
-- 	ptolemaios_calories INT,
-- 	PRIMARY KEY (ptolemaios_id)
-- );

-- -- Version: 1.3
-- -- Description: Create User table, used for only admin currently
-- CREATE EXTENSION IF NOT EXISTS citext;

-- CREATE TABLE IF NOT EXISTS users (
-- id bigserial PRIMARY KEY,
-- created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(), 
-- name text NOT NULL,
-- email citext UNIQUE NOT NULL,
-- password_hash bytea NOT NULL
-- );

-- -- Version: 1.4
-- -- Description: Create Tokens for the privileged actions
-- CREATE TABLE IF NOT EXISTS tokens (
-- hash bytea PRIMARY KEY,
-- user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE, 
-- expiry timestamp(0) with time zone NOT NULL,
-- scope text NOT NULL
-- );

