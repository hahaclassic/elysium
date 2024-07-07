-- +goose Up
-- +goose StatementBegin

-- id is BIGINT because its telegram id
CREATE TABLE IF NOT EXISTS users (
    id BIGINT NOT NULL, 
    username VARCHAR(255) NOT NULL,
    premium BOOLEAN NOT NULL,
    language VARCHAR(3) NOT NULL,
    PRIMARY KEY ( id )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
