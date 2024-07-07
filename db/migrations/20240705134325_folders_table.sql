-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS folders (
    id UUID NOT NULL,
    name TEXT NOT NULL,
    PRIMARY KEY ( id )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS folders;
-- +goose StatementEnd
