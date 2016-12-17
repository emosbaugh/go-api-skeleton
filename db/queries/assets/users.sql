-- users.sql

-- name: list
SELECT email,password,created_at,updated_at,password_updated_at
FROM users;

-- name: get
SELECT email,password,created_at,updated_at,password_updated_at
FROM users
WHERE email = $1;

-- name: create
INSERT INTO users (email,password,created_at,updated_at,password_updated_at)
VALUES (:email,:password,current_timestamp,current_timestamp,current_timestamp);

-- name: update-password
UPDATE users
SET password = :password, updated_at = current_timestamp, password_updated_at = current_timestamp
WHERE email = :email;
