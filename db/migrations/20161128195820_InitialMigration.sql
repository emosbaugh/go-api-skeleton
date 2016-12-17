
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    email varchar(128) NOT NULL,
    password varchar(128) NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    password_updated_at timestamp NOT NULL,
    PRIMARY KEY(email)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
