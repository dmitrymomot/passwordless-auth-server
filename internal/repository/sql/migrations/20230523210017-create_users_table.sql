
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL CHECK (email <> ''),
    verified BOOLEAN NOT NULL DEFAULT false,
    updated_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX IF NOT EXISTS users_email_idx ON users (email);


-- +migrate Down
DROP TABLE IF EXISTS users;